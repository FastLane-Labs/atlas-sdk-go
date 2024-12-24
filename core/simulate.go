package core

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"slices"
	"strings"

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

func (sdk *AtlasSdk) SimulateUserOperation(chainId uint64, version *string, userOp *types.UserOperation) *UserOperationSimulationError {
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
			Gas:       userOp.Gas.Uint64() + 1500000, // Add gas for validateCalls and others
			GasFeeCap: new(big.Int).Set(userOp.MaxFeePerGas),
			Value:     new(big.Int).Set(userOp.Value),
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

func (sdk *AtlasSdk) SimulateSolverOperation(chainId uint64, version *string, userOp *types.UserOperation, solverOp *types.SolverOperation) *SolverOperationSimulationError {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return &SolverOperationSimulationError{err: err}
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, version)
	if err != nil {
		return &SolverOperationSimulationError{err: err}
	}

	simulatorAbi, err := contract.GetSimulatorAbi(version)
	if err != nil {
		return &SolverOperationSimulationError{err: err}
	}

	userOpHash, err := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), chainId, version)
	if err != nil {
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to hash userOp: %w", err)}
	}

	dAppOp := &types.DAppOperation{
		From:          common.Address{},
		To:            userOp.To,
		Nonce:         big.NewInt(0),
		Deadline:      userOp.Deadline,
		Control:       userOp.Control,
		Bundler:       common.Address{},
		UserOpHash:    userOpHash,
		CallChainHash: common.Hash{},
		Signature:     []byte(""),
	}

	if utils.FlagVerifyCallChainHash(userOp.CallConfig) {
		callChainHash, err := CallChainHash(userOp, []*types.SolverOperation{solverOp})
		if err != nil {
			return &SolverOperationSimulationError{err: fmt.Errorf("failed to calculate callChainHash: %w", err)}
		}
		dAppOp.CallChainHash = callChainHash
	}

	pData, err := simulatorAbi.Pack(simSolverCallFunction, userOp, solverOp, dAppOp)
	if err != nil {
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to pack %s: %w", simSolverCallFunction, err)}
	}

	gasPrice := new(big.Int).Set(userOp.MaxFeePerGas)
	if solverOp.MaxFeePerGas.Cmp(userOp.MaxFeePerGas) > 0 {
		gasPrice.Set(solverOp.MaxFeePerGas)
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:        &simulatorAddr,
			Gas:       userOp.Gas.Uint64() + solverOp.Gas.Uint64() + 1500000, // Add gas for validateCalls and others
			GasFeeCap: gasPrice,
			Value:     new(big.Int).Set(userOp.Value),
			Data:      pData,
		},
		nil,
	)
	if err != nil {
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to call %s: %w", simSolverCallFunction, err)}
	}

	validOp, err := simulatorAbi.Unpack(simSolverCallFunction, bData)
	if err != nil {
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to unpack %s: %w pData %s", simSolverCallFunction, err, hex.EncodeToString(pData))}
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		solverOutcomeResult := validOp[2].(*big.Int)
		return &SolverOperationSimulationError{
			Result:        result,
			SolverOutcome: solverOutcomeResult.Uint64(),
			Data:          hex.EncodeToString(pData),
		}
	}

	return nil
}
func (sdk *AtlasSdk) SimulateMetacall(
	chainId uint64,
	version *string,
	from common.Address,
	userOp *types.UserOperation,
	solverOps []*types.SolverOperation,
	dAppOp *types.DAppOperation,
) (*big.Int, error) {
	supportedVersions := []string{"1.2", "1.3"}
	if version == nil || !slices.Contains(supportedVersions, *version) {
		return nil, fmt.Errorf("unsupported version %s, supported versions are %s", *version, strings.Join(supportedVersions, ", "))
	}

	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get eth client: %w", err)
	}

	atlasAddr, err := config.GetAtlasAddress(chainId, version)
	if err != nil {
		return nil, fmt.Errorf("failed to get atlas address: %w", err)
	}

	atlasAbi, err := contract.GetAtlasAbi(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get atlas abi: %w", err)
	}

	pData, err := atlasAbi.Pack("metacall", userOp, solverOps, dAppOp)
	if err != nil {
		return nil, fmt.Errorf("failed to pack metacall: %w", err)
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	callMsg := ethereum.CallMsg{
		To:   &atlasAddr,
		Data: pData,
	}

	var traceResult interface{}
	err = ethClient.Client().CallContext(ctx, &traceResult, "debug_traceCall", callMsg, "latest", map[string]interface{}{
		"tracer": "callTracer",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to trace call: %w", err)
	}

	var parsedBidAmount *big.Int
	parsed := false
	for _, logEntry := range traceResult.(map[string]interface{})["logs"].([]interface{}) {
		eventLog := logEntry.(map[string]interface{})
		if eventLog["name"] == "SolverTxResult" {
			inputs := eventLog["inputs"].([]interface{})
			for _, input := range inputs {
				inputMap := input.(map[string]interface{})
				if inputMap["name"] == "bidAmount" {
					amountStr := inputMap["value"].(string)
					bidAmount, ok := new(big.Int).SetString(amountStr, 10)
					if !ok {
						return nil, fmt.Errorf("failed to parse bid amount")
					}
					parsedBidAmount = bidAmount
					parsed = true
					break
				}
			}
		}
		if parsed {
			break
		}
	}

	if !parsed {
		return nil, fmt.Errorf("SolverTxResult event not found in trace logs")
	}

	return parsedBidAmount, nil
}
