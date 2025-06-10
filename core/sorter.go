package core

import (
	"fmt"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum"
)

const (
	sortBidsFunction = "sortBids"
)

func (sdk *AtlasSdk) SortSolverOperations(chainId uint64, version *string, userOp *types.UserOperation, solverOps types.SolverOperations) (types.SolverOperations, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	sorterAddr, err := config.GetSorterAddress(chainId, version)
	if err != nil {
		return nil, err
	}

	sorterAbi, err := contract.GetSorterAbi(version)
	if err != nil {
		return nil, err
	}

	userOp.Sanitize()
	solverOps.Sanitize()

	pData, err := sorterAbi.Pack(sortBidsFunction, userOp.ToParams(), solverOps)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s: %w", sortBidsFunction, err)
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:   &sorterAddr,
			Data: pData,
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", sortBidsFunction, err)
	}

	_sortedSolverOps, err := sorterAbi.Unpack(sortBidsFunction, bData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", sortBidsFunction, err)
	}

	sortedSolverOps, ok := _sortedSolverOps[0].(types.SolverOperations)
	if !ok {
		return nil, fmt.Errorf("failed to cast %s: %w", sortBidsFunction, err)
	}

	return sortedSolverOps, nil
}
