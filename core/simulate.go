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

func (sdk *AtlasSdk) SimulateUserOperation(chainId uint64, userOp *types.UserOperation) error {
	pData, err := contract.SimulatorAbi.Pack("simUserOperation", *userOp)
	if err != nil {
		return fmt.Errorf("failed to pack simUserOperation: %w", err)
	}

	if _, ok := sdk.ethClient[chainId]; !ok {
		return fmt.Errorf("no ethClient for chainId %d", chainId)
	}

	simulatorAddr, ok := sdk.simulatorAddress[chainId]
	if !ok {
		return fmt.Errorf("no simulator Address for chainId %d", chainId)
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
		return fmt.Errorf("failed to call simUserOperation: %w", err)
	}

	validOp, err := contract.SimulatorAbi.Unpack("simUserOperation", bData)
	if err != nil {
		return fmt.Errorf("failed to unpack simUserOperation: %w", err)
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		validCallResult := validOp[2].(*big.Int)
		return fmt.Errorf(
			"simUserOperation failed with result %d and outcome %s and pData %s",
			result,
			hex.EncodeToString(validCallResult.Bytes()),
			hex.EncodeToString(pData),
		)
	}

	return nil
}

func (sdk *AtlasSdk) SimulateSolverOperation(chainId uint64, userOp *types.UserOperation, solverOp *types.SolverOperation) error {
	userOpHash, err := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), chainId)
	if err != nil {
		return fmt.Errorf("failed to hash userOp: %w", err)
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
			return fmt.Errorf("failed to calculate callChainHash: %w", err)
		}
		dAppOp.CallChainHash = callChainHash
	}

	pData, err := contract.SimulatorAbi.Pack("simSolverCall", *userOp, *solverOp, *dAppOp)
	if err != nil {
		return fmt.Errorf("failed to pack simSolverCall: %w", err)
	}

	gasPrice := new(big.Int).Set(userOp.MaxFeePerGas)
	if solverOp.MaxFeePerGas.Cmp(userOp.MaxFeePerGas) > 0 {
		gasPrice.Set(solverOp.MaxFeePerGas)
	}

	if _, ok := sdk.ethClient[chainId]; !ok {
		return fmt.Errorf("no ethClient for chainId %d", chainId)
	}

	simulatorAddr, ok := sdk.simulatorAddress[chainId]
	if !ok {
		return fmt.Errorf("no simulator Address for chainId %d", chainId)
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
		return fmt.Errorf("failed to call simSolverCall: %w", err)
	}

	validOp, err := contract.SimulatorAbi.Unpack("simSolverCall", bData)
	if err != nil {
		return fmt.Errorf("failed to unpack simSolverCall: %w pData %s", err, hex.EncodeToString(pData))
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		solverOutcomeResult := validOp[2].(*big.Int)
		return fmt.Errorf(
			"simSolverCall failed with result %d and outcome %s and pData %s",
			result,
			hex.EncodeToString(solverOutcomeResult.Bytes()),
			hex.EncodeToString(pData),
		)
	}

	return nil
}
