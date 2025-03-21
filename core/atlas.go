package core

import (
	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	metacallFunction = "metacall"
)

func (sdk *AtlasSdk) Metacall(chainId uint64, version *string, transactOpts *bind.TransactOpts, userOp types.UserOperation, solverOps types.SolverOperations, dAppOp *types.DAppOperation, gasRefundBeneficiary *common.Address) (*gethTypes.Transaction, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	atlasAddr, err := config.GetAtlasAddress(chainId, version)
	if err != nil {
		return nil, err
	}

	atlasAbi, err := contract.GetAtlasAbi(version)
	if err != nil {
		return nil, err
	}

	userOp.Sanitize()
	solverOps.Sanitize()
	dAppOp.Sanitize()

	var (
		contract   = bind.NewBoundContract(atlasAddr, *atlasAbi, ethClient, ethClient, ethClient)
		_solverOps = make([]types.SolverOperation, len(solverOps))
		params     []interface{}
	)

	for i, op := range solverOps {
		_solverOps[i] = *op
	}

	switch config.GetVersion(version) {
	case config.AtlasV_1_0, config.AtlasV_1_1:
		params = append(params, userOp, _solverOps, dAppOp)

	default:
		var _gasRefundBeneficiary common.Address
		if gasRefundBeneficiary != nil {
			_gasRefundBeneficiary = *gasRefundBeneficiary
		}

		params = append(params, userOp, _solverOps, dAppOp, _gasRefundBeneficiary)
	}

	return contract.Transact(transactOpts, metacallFunction, params...)
}
