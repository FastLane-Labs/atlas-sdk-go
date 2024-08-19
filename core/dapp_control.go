package core

import (
	"github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

func (sdk *AtlasSdk) GetDAppCallConfig(chainId uint64, dAppControlAddr common.Address) (uint32, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return 0, err
	}

	dAppControlContract, err := dappcontrol.NewDAppControl(dAppControlAddr, ethClient)
	if err != nil {
		return 0, err
	}

	callConfig, err := dAppControlContract.CALLCONFIG(nil)
	if err != nil {
		return 0, err
	}

	return callConfig, nil
}

func (sdk *AtlasSdk) GetDAppConfig(chainId uint64, userOp *types.UserOperation, dAppControlAddr common.Address) (*dappcontrol.DAppConfig, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	dAppControlContract, err := dappcontrol.NewDAppControl(dAppControlAddr, ethClient)
	if err != nil {
		return nil, err
	}

	dAppConfig, err := dAppControlContract.GetDAppConfig(nil, dappcontrol.UserOperation(*userOp))
	if err != nil {
		return nil, err
	}

	return &dAppConfig, nil
}
