package core

import (
	"errors"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
)

func (sdk *AtlasSdk) Metacall(chainId uint64, transactOpts *bind.TransactOpts, userOp *types.UserOperation, solverOps types.SolverOperations, dAppOp *types.DAppOperation) (*gethTypes.Transaction, error) {
	contract, ok := sdk.atlasContract[chainId]
	if !ok {
		return nil, errors.New("atlas contract not found")
	}

	userOp.Sanitize()
	solverOps.Sanitize()
	dAppOp.Sanitize()

	_dAppOp := atlas.DAppOperation{
		From:          dAppOp.From,
		To:            dAppOp.To,
		Nonce:         dAppOp.Nonce,
		Deadline:      dAppOp.Deadline,
		Control:       dAppOp.Control,
		Bundler:       dAppOp.Bundler,
		UserOpHash:    dAppOp.UserOpHash,
		CallChainHash: dAppOp.CallChainHash,
		Signature:     dAppOp.Signature,
	}

	_solverOps := make([]atlas.SolverOperation, len(solverOps))
	for i, solverOp := range solverOps {
		_solverOps[i] = atlas.SolverOperation{
			From:         solverOp.From,
			To:           solverOp.To,
			Value:        solverOp.Value,
			Gas:          solverOp.Gas,
			MaxFeePerGas: solverOp.MaxFeePerGas,
			Deadline:     solverOp.Deadline,
			Solver:       solverOp.Solver,
			Control:      solverOp.Control,
			UserOpHash:   solverOp.UserOpHash,
			BidToken:     solverOp.BidToken,
			BidAmount:    solverOp.BidAmount,
			Data:         solverOp.Data,
			Signature:    solverOp.Signature,
		}
	}

	tx, err := contract.Metacall(transactOpts, atlas.UserOperation(*userOp), _solverOps, _dAppOp)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
