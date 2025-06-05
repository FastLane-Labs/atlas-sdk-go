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
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	simUserOperationFunction         = "simUserOperation"
	simSolverCallFunction            = "simSolverCall"
	minGasBuffer                     = uint64(2_500_000)
	simGasSuggestedBuffer            = uint64(170_000)
	simGasSuggestedBufferMonad       = uint64(300_000)
	estimateMetacallGasLimitFunction = "estimateMetacallGasLimit"
)

var (
	maxUint96, _ = new(big.Int).SetString("79228162514264337593543950335", 10)

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
		30: "MetacallGasLimitTooLow",
		31: "MetacallGasLimitTooHigh",
		32: "DAppGasLimitMismatch",
		33: "SolverGasLimitMismatch",
		34: "BundlerSurchargeRateMismatch",
		35: "ExPostBidsAndMultipleSuccessfulSolversNotSupportedTogether",
		36: "InvertsBidValueAndMultipleSuccessfulSolversNotSupportedTogether",
		37: "NeedSolversForMultipleSuccessfulSolvers",
		38: "SolverCannotBeAuctioneerForMultipleSuccessfulSolvers",
		39: "CannotRequireFulfillmentForMultipleSuccessfulSolvers",
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
	_version := *version
	if config.IsMonad(chainId) {
		_version = config.ToMonadVersion(version)
	}

	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return &UserOperationSimulationError{err: err}
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, &_version)
	if err != nil {
		return &UserOperationSimulationError{err: err}
	}

	simulatorAbi, err := contract.GetSimulatorAbi(&_version)
	if err != nil {
		return &UserOperationSimulationError{err: err}
	}

	pData, err := simulatorAbi.Pack(simUserOperationFunction, userOp.ToParams())
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to pack %s: %w", simUserOperationFunction, err)}
	}

	var (
		gasLimit uint64
		v1_5     = config.AtlasV_1_5
		v1_6_1   = config.AtlasV_1_6_1
	)

	gte_1_5, err := config.IsVersionAtLeast(&_version, &v1_5)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to check version: %w", err)}
	}

	gte_1_6_1, err := config.IsVersionAtLeast(&_version, &v1_6_1)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to check version: %w", err)}
	}

	if gte_1_6_1 {
		gasLimit = 0 // Infinite gas, the simulator will figure it out
	} else if gte_1_5 {
		_gasLimit, err := sdk.EstimateMetacallGasLimit(chainId, &_version, userOp, types.SolverOperations{}, userOp.GetMaxFeePerGas())
		if err != nil {
			return &UserOperationSimulationError{err: fmt.Errorf("failed to estimate metacall gas limit: %w", err)}
		}

		buffer := simGasSuggestedBuffer
		if config.IsMonad(chainId) {
			buffer = simGasSuggestedBufferMonad
		}

		gasLimit = _gasLimit + buffer
	} else {
		gasLimit = userOp.GetGas().Uint64() + minGasBuffer
	}

	gasPrice := new(big.Int).Set(userOp.GetMaxFeePerGas())

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:        &simulatorAddr,
			Gas:       gasLimit,
			GasFeeCap: gasPrice,
			Value:     new(big.Int).Set(userOp.GetValue()),
			Data:      pData,
		},
		nil)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf(
			"failed to call %s: %w, simulatorAddr %s, pData %s, version %s, userOp %s, gasLimit %d",
			simUserOperationFunction,
			err,
			simulatorAddr.Hex(),
			hex.EncodeToString(pData),
			_version,
			userOp.EncodeToRaw(),
			gasLimit,
		)}
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
			Data:             hex.EncodeToString(pData) + fmt.Sprintf(", simulatorAddr %s, version %s, gasLimit %d, gasPrice %s, rawReturnData %s", simulatorAddr.Hex(), _version, gasLimit, gasPrice.String(), hex.EncodeToString(bData)),
		}
	}

	return nil
}

func (sdk *AtlasSdk) SimulateSolverOperation(chainId uint64, version *string, userOp *types.UserOperation, solverOp *types.SolverOperation, allowTracing bool) (*big.Int, *SolverOperationSimulationError) {
	_version := *version
	if config.IsMonad(chainId) {
		_version = config.ToMonadVersion(version)
	}

	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, &_version)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	simulatorAbi, err := contract.GetSimulatorAbi(&_version)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	dAppOp, err := generateDappOperationForSimulator(&_version, userOp, solverOp)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	pData, err := simulatorAbi.Pack(simSolverCallFunction, userOp.ToParams(), solverOp, dAppOp)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to pack %s: %w", simSolverCallFunction, err)}
	}

	gasPrice := new(big.Int).Set(userOp.GetMaxFeePerGas())
	if solverOp.MaxFeePerGas.Cmp(userOp.GetMaxFeePerGas()) > 0 {
		gasPrice.Set(solverOp.MaxFeePerGas)
	}

	var (
		gasLimit uint64
		v1_5     = config.AtlasV_1_5
		v1_6_1   = config.AtlasV_1_6_1
	)

	gte_1_5, err := config.IsVersionAtLeast(&_version, &v1_5)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to check version: %w", err)}
	}

	gte_1_6_1, err := config.IsVersionAtLeast(&_version, &v1_6_1)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to check version: %w", err)}
	}

	if gte_1_6_1 {
		gasLimit = 0 // Infinite gas, the simulator will figure it out
	} else if gte_1_5 {
		_gasLimit, err := sdk.EstimateMetacallGasLimit(chainId, &_version, userOp, types.SolverOperations{solverOp}, gasPrice)
		if err != nil {
			return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to estimate metacall gas limit: %w", err)}
		}

		buffer := simGasSuggestedBuffer
		if config.IsMonad(chainId) {
			buffer = simGasSuggestedBufferMonad
		}

		gasLimit = _gasLimit + buffer
	} else {
		gasBuffer := minGasBuffer

		solverGasLimit, err := sdk.GetDAppSolverGasLimit(chainId, &_version, userOp.GetControl())
		if err != nil {
			return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to get solver gas limit: %w", err)}
		}

		if uint64(solverGasLimit) > gasBuffer {
			gasBuffer = uint64(solverGasLimit)
		}

		gasLimit = userOp.GetGas().Uint64() + solverOp.Gas.Uint64() + gasBuffer
	}

	var (
		bData       []byte
		traceResult callFrame
		callMsg     = ethereum.CallMsg{
			To:        &simulatorAddr,
			Gas:       gasLimit,
			GasFeeCap: gasPrice,
			Value:     new(big.Int).Set(userOp.GetValue()),
			Data:      pData,
		}
	)

	exPostBids, err := utils.FlagExPostBids(userOp.GetCallConfig(), &_version)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: err}
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	if !exPostBids || !allowTracing {
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
			Data:          hex.EncodeToString(pData) + fmt.Sprintf(", simulatorAddr %s, version %s, gasLimit %d, gasPrice %s, rawReturnData %s", simulatorAddr.Hex(), _version, gasLimit, gasPrice.String(), hex.EncodeToString(bData)),
		}
	}

	if !exPostBids {
		// If ex post bids are not enabled, we can directly return the bid amount
		return solverOp.BidAmount, nil
	}

	exPostBidAmount, err := sdk.GetSolverBidAmountFromTrace(chainId, &_version, &traceResult)
	if err != nil {
		return nil, &SolverOperationSimulationError{err: fmt.Errorf("failed to get solver bid amount from trace: %w pData %s", err, hex.EncodeToString(pData))}
	}

	return exPostBidAmount, nil
}

func generateDappOperationForSimulator(version *string, userOp *types.UserOperation, solverOp *types.SolverOperation) (*types.DAppOperation, error) {
	trustedOpHash, err := utils.FlagTrustedOpHash(userOp.GetCallConfig(), version)
	if err != nil {
		return nil, err
	}

	userOpHash, err := userOp.Hash(trustedOpHash)
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

	verifyCallChainHash, err := utils.FlagVerifyCallChainHash(userOp.GetCallConfig(), version)
	if err != nil {
		return nil, err
	}

	if verifyCallChainHash {
		callChainHash, err := CallChainHash(userOp, []*types.SolverOperation{solverOp})
		if err != nil {
			return nil, fmt.Errorf("failed to calculate callChainHash: %w", err)
		}
		dAppOp.CallChainHash = callChainHash
	}

	return dAppOp, nil
}

func (sdk *AtlasSdk) EstimateMetacallGasLimit(chainId uint64, version *string, userOp *types.UserOperation, solverOps types.SolverOperations, gasPrice *big.Int) (uint64, error) {
	minVersion := config.AtlasV_1_5
	if minSupport, err := config.IsVersionAtLeast(version, &minVersion); err != nil || !minSupport {
		return 0, fmt.Errorf("metacall gas limit estimation is only supported for Atlas v1.5 and above")
	}

	simulatorAbi, err := contract.GetSimulatorAbi(version)
	if err != nil {
		return 0, fmt.Errorf("failed to get simulator abi: %w", err)
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, version)
	if err != nil {
		return 0, fmt.Errorf("failed to get simulator address: %w", err)
	}

	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return 0, fmt.Errorf("failed to get eth client: %w", err)
	}

	_solverOps := make([]types.SolverOperation, len(solverOps))
	for i, op := range solverOps {
		_solverOps[i] = *op
	}

	pData, err := simulatorAbi.Pack(estimateMetacallGasLimitFunction, userOp.ToParams(), _solverOps)
	if err != nil {
		return 0, fmt.Errorf("failed to pack %s: %w", estimateMetacallGasLimitFunction, err)
	}

	overrides := map[common.Address]map[string]interface{}{
		common.HexToAddress("0x0000000000000000000000000000000000000000"): {
			"balance": (*hexutil.Big)(maxUint96),
		},
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	var result string
	err = ethClient.Client().CallContext(ctx, &result, "eth_call", map[string]interface{}{
		"to":       simulatorAddr.Hex(),
		"gasPrice": (*hexutil.Big)(gasPrice),
		"data":     hexutil.Encode(pData),
	}, "latest", overrides)
	if err != nil {
		return 0, fmt.Errorf("failed to call %s: %w - pData %s - simulatorAddr %s", estimateMetacallGasLimitFunction, err, hex.EncodeToString(pData), simulatorAddr.Hex())
	}

	bData, err := hexutil.Decode(result)
	if err != nil {
		return 0, fmt.Errorf("failed to decode result: %w", err)
	}

	gasLimit, err := simulatorAbi.Unpack(estimateMetacallGasLimitFunction, bData)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack %s: %w", estimateMetacallGasLimitFunction, err)
	}

	if len(gasLimit) != 1 {
		return 0, fmt.Errorf("expected 1 gas limit, got %d", len(gasLimit))
	}

	return gasLimit[0].(*big.Int).Uint64(), nil
}
