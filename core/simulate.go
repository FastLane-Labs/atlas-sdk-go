package core

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const (
	simUserOperationFunction = "simUserOperation"
	simSolverCallFunction    = "simSolverCall"
	minGasBuffer             = uint64(1_500_000)
)

var (
	SimulationResult = map[uint8]string{
		0: "Unknown",
		1: "VerificationSimFail",
		2: "PreOpsSimFail",
		3: "UserOpSimFail",
		4: "SolverSimFail",
		5: "AllocateValueSimFail",
		6: "PostOpsSimFail",
		7: "SimulationPassed",
	}

	ValidCallsResult = map[uint8]string{
		0:  "Valid",
		1:  "UserFromInvalid",
		2:  "UserSignatureInvalid",
		3:  "DAppSignatureInvalid",
		4:  "UserNonceInvalid",
		5:  "InvalidDAppNonce",
		6:  "UnknownAuctioneerNotAllowed",
		7:  "InvalidAuctioneer",
		8:  "InvalidBundler",
		9:  "InvertBidValueCannotBeExPostBids",
		10: "GasPriceHigherThanMax",
		11: "TxValueLowerThanCallValue",
		12: "TooManySolverOps",
		13: "UserDeadlineReached",
		14: "DAppDeadlineReached",
		15: "ExecutionEnvEmpty",
		16: "NoSolverOp",
		17: "InvalidSequence",
		18: "OpHashMismatch",
		19: "DeadlineMismatch",
		20: "InvalidControl",
		21: "InvalidSolverGasLimit",
		22: "InvalidCallConfig",
		23: "CallConfigMismatch",
		24: "DAppToInvalid",
		25: "UserToInvalid",
		26: "ControlMismatch",
		27: "InvalidCallChainHash",
		28: "DAppNotEnabled",
		29: "BothUserAndDAppNoncesCannotBeSequential",
	}

	SolverOutcome = map[uint64]string{
		1 << 0:  "InvalidSignature",
		1 << 1:  "InvalidUserHash",
		1 << 2:  "DeadlinePassedAlt",
		1 << 3:  "GasPriceBelowUsersAlt",
		1 << 4:  "InvalidTo",
		1 << 5:  "UserOutOfGas",
		1 << 6:  "AlteredControl",
		1 << 7:  "AltOpHashMismatch",
		1 << 8:  "DeadlinePassed",
		1 << 9:  "GasPriceOverCap",
		1 << 10: "InvalidSolver",
		1 << 11: "InvalidBidToken",
		1 << 12: "PerBlockLimit",
		1 << 13: "InsufficientEscrow",
		1 << 14: "GasPriceBelowUsers",
		1 << 15: "CallValueTooHigh",
		1 << 16: "PreSolverFailed",
		1 << 17: "SolverOpReverted",
		1 << 18: "PostSolverFailed",
		1 << 19: "BidNotPaid",
		1 << 20: "InvertedBidExceedsCeiling",
		1 << 21: "BalanceNotReconciled",
		1 << 22: "CallbackNotCalled",
		1 << 23: "EVMError",
	}
)

type UserOperationSimulationError struct {
	err              error
	Result           uint8
	ValidCallsResult uint8
	Data             string
}

func (e *UserOperationSimulationError) ResultString() string {
	return SimulationResult[e.Result]
}

func (e *UserOperationSimulationError) ValidCallsResultString() string {
	return ValidCallsResult[e.ValidCallsResult]
}

func (e *UserOperationSimulationError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return fmt.Sprintf(
		"simUserOperation failed with result %d (%s) and validCallsResult %d (%s) and pData %s",
		e.Result,
		e.ResultString(),
		e.ValidCallsResult,
		e.ValidCallsResultString(),
		e.Data,
	)
}

type SolverOperationSimulationError struct {
	err           error
	Result        uint8
	SolverOutcome uint64
	Data          string
}

func (e *SolverOperationSimulationError) ResultString() string {
	return SimulationResult[e.Result]
}

func (e *SolverOperationSimulationError) SolverOutcomeString() string {
	outcome, ok := SolverOutcome[e.SolverOutcome]
	if !ok {
		return fmt.Sprintf("UnknownSolverOutcome %d", e.SolverOutcome)
	}
	return outcome
}

func (e *SolverOperationSimulationError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return fmt.Sprintf(
		"simSolverCall failed with result %d (%s) and outcome %d (%s) and pData %s",
		e.Result,
		e.ResultString(),
		e.SolverOutcome,
		e.SolverOutcomeString(),
		e.Data,
	)
}

func (sdk *AtlasSdk) SimulateUserOperation(chainId uint64, version *string, userOp types.UserOperation) *UserOperationSimulationError {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return &UserOperationSimulationError{err: err}
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, version)
	if err != nil {
		return &UserOperationSimulationError{err: err}
	}

	simulatorAbi, err := contract.GetSimulatorAbi(version)
	if err != nil {
		return &UserOperationSimulationError{err: err}
	}

	pData, err := simulatorAbi.Pack(simUserOperationFunction, userOp)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to pack %s: %w", simUserOperationFunction, err)}
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:        &simulatorAddr,
			Gas:       userOp.GetGas().Uint64() + minGasBuffer,
			GasFeeCap: new(big.Int).Set(userOp.GetMaxFeePerGas()),
			Value:     new(big.Int).Set(userOp.GetValue()),
			Data:      pData,
		},
		nil)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to call %s: %w", simUserOperationFunction, err)}
	}

	validOp, err := simulatorAbi.Unpack(simUserOperationFunction, bData)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to unpack %s: %w", simUserOperationFunction, err)}
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		validCallResult := validOp[2].(*big.Int)
		return &UserOperationSimulationError{
			Result:           result,
			ValidCallsResult: uint8(validCallResult.Uint64()),
			Data:             hex.EncodeToString(pData),
		}
	}

	return nil
}

func (sdk *AtlasSdk) SimulateSolverOperation(chainId uint64, version *string, userOp types.UserOperation, solverOp *types.SolverOperation, allowTracing bool) (*big.Int, *SolverOperationSimulationError) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, version)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	simulatorAbi, err := contract.GetSimulatorAbi(version)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	dAppOp, err := generateDappOperationForSimulator(chainId, version, userOp, solverOp)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	pData, err := simulatorAbi.Pack(simSolverCallFunction, userOp, solverOp, dAppOp)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to pack %s: %w", simSolverCallFunction, err)}
	}

	gasPrice := new(big.Int).Set(userOp.GetMaxFeePerGas())
	if solverOp.MaxFeePerGas.Cmp(userOp.GetMaxFeePerGas()) > 0 {
		gasPrice.Set(solverOp.MaxFeePerGas)
	}

	gasBuffer := minGasBuffer

	solverGasLimit, err := sdk.GetDAppSolverGasLimit(chainId, userOp.GetControl())
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to get solver gas limit: %w", err)}
	}

	if solverGasLimit.Uint64() > gasBuffer {
		gasBuffer = solverGasLimit.Uint64()
	}

	var (
		bData       []byte
		traceResult callFrame
		callMsg     = ethereum.CallMsg{
			To:        &simulatorAddr,
			Gas:       userOp.GetGas().Uint64() + solverOp.Gas.Uint64() + gasBuffer,
			GasFeeCap: gasPrice,
			Value:     new(big.Int).Set(userOp.GetValue()),
			Data:      pData,
		}
	)

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	if !utils.FlagExPostBids(userOp.GetCallConfig()) || !allowTracing {
		bData, err = ethClient.CallContract(ctx, callMsg, nil)
		if err != nil {
			return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to call %s: %w", simSolverCallFunction, err)}
		}
	} else {
		err = ethClient.Client().CallContext(
			ctx,
			&traceResult,
			traceCallMethod,
			toCallArg(callMsg),
			"latest",
			map[string]interface{}{
				"tracer": "callTracer",
				"tracerConfig": map[string]interface{}{
					"withLog": true,
				},
			},
		)
		if err != nil {
			return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to trace %s: %w", simSolverCallFunction, err)}
		}

		if traceResult.Error != "" {
			return nil, &SolverOperationSimulationError{err: fmt.Errorf("tracing failed for %s: %s", simSolverCallFunction, traceResult.Error)}
		}

		bData = traceResult.Output
	}

	validOp, err := simulatorAbi.Unpack(simSolverCallFunction, bData)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to unpack %s: %w pData %s", simSolverCallFunction, err, hex.EncodeToString(pData))}
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		solverOutcomeResult := validOp[2].(*big.Int)
		return nil, &SolverOperationSimulationError{
			Result:        result,
			SolverOutcome: solverOutcomeResult.Uint64(),
			Data:          hex.EncodeToString(pData),
		}
	}

	if !utils.FlagExPostBids(userOp.GetCallConfig()) {
		// If ex post bids are not enabled, we can directly return the bid amount
		return solverOp.BidAmount, nil
	}

	exPostBidAmount, err := sdk.GetSolverBidAmountFromTrace(chainId, version, &traceResult)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to get solver bid amount from trace: %w pData %s", err, hex.EncodeToString(pData))}
	}

	return exPostBidAmount, nil
}

func generateDappOperationForSimulator(chainId uint64, version *string, userOp types.UserOperation, solverOp *types.SolverOperation) (*types.DAppOperation, error) {
	userOpHash, err := userOp.Hash(utils.FlagTrustedOpHash(userOp.GetCallConfig()), chainId, version)
	if err != nil {
		return nil, fmt.Errorf("failed to hash userOp: %w", err)
	}

	dAppOp := &types.DAppOperation{
		From:          common.Address{},
		To:            userOp.GetTo(),
		Nonce:         big.NewInt(0),
		Deadline:      userOp.GetDeadline(),
		Control:       userOp.GetControl(),
		Bundler:       common.Address{},
		UserOpHash:    userOpHash,
		CallChainHash: common.Hash{},
		Signature:     []byte(""),
	}

	if utils.FlagVerifyCallChainHash(userOp.GetCallConfig()) {
		callChainHash, err := CallChainHash(userOp, []*types.SolverOperation{solverOp})
		if err != nil {
			return nil, fmt.Errorf("failed to calculate callChainHash: %w", err)
		}
		dAppOp.CallChainHash = callChainHash
	}

	return dAppOp, nil
}
