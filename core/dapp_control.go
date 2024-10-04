package core

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

func (sdk *AtlasSdk) getDAppControlContract(chainId uint64, dAppControlAddr common.Address) (*dappcontrol.DAppControl, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	dAppControlContract, err := dappcontrol.NewDAppControl(dAppControlAddr, ethClient)
	if err != nil {
		return nil, err
	}

	return dAppControlContract, nil
}

func (sdk *AtlasSdk) GetDAppCallConfig(chainId uint64, dAppControlAddr common.Address) (uint32, error) {
	dAppControlContract, err := sdk.getDAppControlContract(chainId, dAppControlAddr)
	if err != nil {
		return 0, err
	}

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	callConfig, err := dAppControlContract.CALLCONFIG(callOpts)
	if err != nil {
		return 0, err
	}

	return callConfig, nil
}

func (sdk *AtlasSdk) GetDAppConfig(chainId uint64, userOp *types.UserOperation, dAppControlAddr common.Address) (*dappcontrol.DAppConfig, error) {
	dAppControlContract, err := sdk.getDAppControlContract(chainId, dAppControlAddr)
	if err != nil {
		return nil, err
	}

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	dAppConfig, err := dAppControlContract.GetDAppConfig(callOpts, dappcontrol.UserOperation(*userOp))
	if err != nil {
		return nil, err
	}

	return &dAppConfig, nil
}

func (sdk *AtlasSdk) GetDAppSolverGasLimit(chainId uint64, dAppControlAddr common.Address) (*big.Int, error) {
	dAppControlContract, err := sdk.getDAppControlContract(chainId, dAppControlAddr)
	if err != nil {
		return nil, err
	}

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	solverGasLimit, err := dAppControlContract.GetSolverGasLimit(callOpts)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetUint64(uint64(solverGasLimit)), nil
}
