package core

import (
	"errors"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
)

func (sdk *AtlasSdk) SortSolverOperations(chainId uint64, userOp *types.UserOperation, solverOps types.SolverOperations) (types.SolverOperations, error) {
	contract, ok := sdk.sorterContract[chainId]
	if !ok {
		return nil, errors.New("sorter contract not found")
	}

	if len(solverOps) == 0 {
		return solverOps, nil
	}

	userOp.Sanitize()

	fmtSolverOps := make([]sorter.SolverOperation, len(solverOps))
	for i, solverOp := range solverOps {
		solverOp.Sanitize()
		fmtSolverOps[i] = sorter.SolverOperation{
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

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	fmtSortedSolverOps, err := contract.SortBids(callOpts, sorter.UserOperation(*userOp), fmtSolverOps)
	if err != nil {
		return nil, err
	}

	sortedSolverOps := make(types.SolverOperations, len(fmtSortedSolverOps))
	for i, fmtSortedSolverOp := range fmtSortedSolverOps {
		sortedSolverOps[i] = &types.SolverOperation{
			From:         fmtSortedSolverOp.From,
			To:           fmtSortedSolverOp.To,
			Value:        fmtSortedSolverOp.Value,
			Gas:          fmtSortedSolverOp.Gas,
			MaxFeePerGas: fmtSortedSolverOp.MaxFeePerGas,
			Deadline:     fmtSortedSolverOp.Deadline,
			Solver:       fmtSortedSolverOp.Solver,
			Control:      fmtSortedSolverOp.Control,
			UserOpHash:   fmtSortedSolverOp.UserOpHash,
			BidToken:     fmtSortedSolverOp.BidToken,
			BidAmount:    fmtSortedSolverOp.BidAmount,
			Data:         fmtSortedSolverOp.Data,
			Signature:    fmtSortedSolverOp.Signature,
		}
	}

	return sortedSolverOps, nil
}
