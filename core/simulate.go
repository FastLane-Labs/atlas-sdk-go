package core

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

	SolverOutcome = map[uint8]string{
		0:  "InvalidSignature",
		1:  "InvalidUserHash",
		2:  "DeadlinePassedAlt",
		3:  "GasPriceBelowUsersAlt",
		4:  "InvalidTo",
		5:  "UserOutOfGas",
		6:  "AlteredControl",
		7:  "AltOpHashMismatch",
		8:  "DeadlinePassed",
		9:  "GasPriceOverCap",
		10: "InvalidSolver",
		11: "InvalidBidToken",
		12: "PerBlockLimit",
		13: "InsufficientEscrow",
		14: "GasPriceBelowUsers",
		15: "CallValueTooHigh",
		16: "PreSolverFailed",
		17: "SolverOpReverted",
		18: "PostSolverFailed",
		19: "BidNotPaid",
		20: "InvertedBidExceedsCeiling",
		21: "BalanceNotReconciled",
		22: "CallbackNotCalled",
		23: "EVMError",
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
	SolverOutcome uint8
	Data          string
}

func (e *SolverOperationSimulationError) ResultString() string {
	return SimulationResult[e.Result]
}

func (e *SolverOperationSimulationError) SolverOutcomeString() string {
	return SolverOutcome[e.SolverOutcome]
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

func (sdk *AtlasSdk) SimulateUserOperation(chainId uint64, userOp *types.UserOperation) *UserOperationSimulationError {
	pData, err := contract.SimulatorAbi.Pack("simUserOperation", *userOp)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to pack simUserOperation: %w", err)}
	}

	if _, ok := sdk.ethClient[chainId]; !ok {
		return &UserOperationSimulationError{err: fmt.Errorf("no ethClient for chainId %d", chainId)}
	}

	simulatorAddr, ok := sdk.simulatorAddress[chainId]
	if !ok {
		return &UserOperationSimulationError{err: fmt.Errorf("no simulator Address for chainId %d", chainId)}
	}

	bData, err := sdk.ethClient[chainId].CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:        &simulatorAddr,
			Gas:       userOp.Gas.Uint64() + 1500000, // Add gas for validateCalls and others
			GasFeeCap: new(big.Int).Set(userOp.MaxFeePerGas),
			Value:     new(big.Int).Set(userOp.Value),
			Data:      pData,
		},
		nil)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to call simUserOperation: %w", err)}
	}

	validOp, err := contract.SimulatorAbi.Unpack("simUserOperation", bData)
	if err != nil {
		return &UserOperationSimulationError{err: fmt.Errorf("failed to unpack simUserOperation: %w", err)}
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

func (sdk *AtlasSdk) SimulateSolverOperation(chainId uint64, userOp *types.UserOperation, solverOp *types.SolverOperation) *SolverOperationSimulationError {
	userOpHash, err := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), chainId)
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

	pData, err := contract.SimulatorAbi.Pack("simSolverCall", *userOp, *solverOp, *dAppOp)
	if err != nil {
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to pack simSolverCall: %w", err)}
	}

	gasPrice := new(big.Int).Set(userOp.MaxFeePerGas)
	if solverOp.MaxFeePerGas.Cmp(userOp.MaxFeePerGas) > 0 {
		gasPrice.Set(solverOp.MaxFeePerGas)
	}

	if _, ok := sdk.ethClient[chainId]; !ok {
		return &SolverOperationSimulationError{err: fmt.Errorf("no ethClient for chainId %d", chainId)}
	}

	simulatorAddr, ok := sdk.simulatorAddress[chainId]
	if !ok {
		return &SolverOperationSimulationError{err: fmt.Errorf("no simulator Address for chainId %d", chainId)}
	}

	bData, err := sdk.ethClient[chainId].CallContract(
		context.Background(),
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
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to call simSolverCall: %w", err)}
	}

	validOp, err := contract.SimulatorAbi.Unpack("simSolverCall", bData)
	if err != nil {
		return &SolverOperationSimulationError{err: fmt.Errorf("failed to unpack simSolverCall: %w pData %s", err, hex.EncodeToString(pData))}
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		solverOutcomeResult := validOp[2].(*big.Int)
		return &SolverOperationSimulationError{
			Result:        result,
			SolverOutcome: uint8(solverOutcomeResult.Uint64()),
			Data:          hex.EncodeToString(pData),
		}
	}

	return nil
}
