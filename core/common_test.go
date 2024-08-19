package core

import (
	"math/big"
	"testing"

	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

func generateUserOperation() *types.UserOperation {
	return &types.UserOperation{
		From:         common.HexToAddress("0x1"),
		To:           common.HexToAddress("0x2"),
		Deadline:     big.NewInt(100),
		Gas:          big.NewInt(200),
		Nonce:        big.NewInt(300),
		MaxFeePerGas: big.NewInt(400),
		Value:        big.NewInt(500),
		Dapp:         common.HexToAddress("0x3"),
		Control:      common.HexToAddress("0x4"),
		CallConfig:   600,
		SessionKey:   common.HexToAddress("0x5"),
		Data:         []byte("data"),
		Signature:    []byte("signature"),
	}
}

func generateSolverOperation() *types.SolverOperation {
	return &types.SolverOperation{
		From:         common.HexToAddress("0x1"),
		To:           common.HexToAddress("0x2"),
		Value:        big.NewInt(100),
		Gas:          big.NewInt(200),
		MaxFeePerGas: big.NewInt(300),
		Deadline:     big.NewInt(400),
		Solver:       common.HexToAddress("0x3"),
		Control:      common.HexToAddress("0x4"),
		UserOpHash:   common.HexToHash("0x9999999999999999999999999999999999999999999999999999999999999999"),
		BidToken:     common.HexToAddress("0x5"),
		BidAmount:    big.NewInt(500),
		Data:         []byte("data"),
		Signature:    []byte("signature"),
	}
}

func TestCallChainHashWithSolverOperations(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	solverOp := generateSolverOperation()
	solverOps := []*types.SolverOperation{solverOp, solverOp, solverOp}

	want := common.HexToHash("0x8a71f907fe61688772ede6e7bb91efa992fe86c27917862adf533984dd56a2b8")

	result, err := CallChainHash(userOp, solverOps)
	if err != nil {
		t.Errorf("CallChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("CallChainHash() = %v, want %v", result, want)
	}
}

func TestCallChainHashWithoutSolverOperations(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	solverOps := []*types.SolverOperation{}

	want := common.HexToHash("0x1feca496343f60c6fd5bfa97ec935fed62285b814ef720ac633dabb1c6e25777")

	result, err := CallChainHash(userOp, solverOps)
	if err != nil {
		t.Errorf("CallChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("CallChainHash() = %v, want %v", result, want)
	}
}
