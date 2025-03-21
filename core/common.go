package core

import (
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CallChainHash(userOp types.UserOperation, solverOps types.SolverOperations) (common.Hash, error) {
	userOpAbiEncoded, err := userOp.AbiEncode()
	if err != nil {
		return common.Hash{}, err
	}

	solverOpsAbiEncoded, err := solverOps.AbiEncode()
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(
		userOpAbiEncoded,
		solverOpsAbiEncoded,
	), nil
}
