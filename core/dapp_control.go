package core

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol"
	dappcontrol_1_5 "github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol/1.5"
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

func (sdk *AtlasSdk) getDAppControlV15Contract(chainId uint64, dAppControlAddr common.Address) (*dappcontrol_1_5.DAppControl, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	dAppControlContract, err := dappcontrol_1_5.NewDAppControl(dAppControlAddr, ethClient)
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

func (sdk *AtlasSdk) GetDAppConfig(chainId uint64, version *string, userOp types.UserOperation, dAppControlAddr common.Address) (types.DAppConfig, error) {
	if *version == config.AtlasV_1_5 {
		dAppControlContract, err := sdk.getDAppControlV15Contract(chainId, dAppControlAddr)
		if err != nil {
			return nil, err
		}

		callOpts, cancel := NewCallOptsWithNetworkDeadline()
		defer cancel()

		dappGasLimit := userOp.GetDappGasLimit()
		if dappGasLimit == 0 {
			return nil, errors.New("dappGasLimit is 0")
		}

		dAppConfig, err := dAppControlContract.GetDAppConfig(callOpts, dappcontrol_1_5.UserOperation{
			From:         userOp.GetFrom(),
			To:           userOp.GetTo(),
			Value:        userOp.GetValue(),
			Gas:          userOp.GetGas(),
			MaxFeePerGas: userOp.GetMaxFeePerGas(),
			Nonce:        userOp.GetNonce(),
			Deadline:     userOp.GetDeadline(),
			Dapp:         userOp.GetDapp(),
			Control:      userOp.GetControl(),
			CallConfig:   userOp.GetCallConfig(),
			SessionKey:   userOp.GetSessionKey(),
			DappGasLimit: userOp.GetDappGasLimit(),
			Data:         userOp.GetData(),
			Signature:    userOp.GetSignature(),
		})
		if err != nil {
			return nil, err
		}

		return &types.DAppConfigV15{
			DAppConfigLegacy: types.DAppConfigLegacy{
				To:             dAppConfig.To,
				CallConfig:     dAppConfig.CallConfig,
				BidToken:       dAppConfig.BidToken,
				SolverGasLimit: dAppConfig.SolverGasLimit,
			},
			DappGasLimit: dAppConfig.DappGasLimit,
		}, nil
	}

	dAppControlContract, err := sdk.getDAppControlContract(chainId, dAppControlAddr)
	if err != nil {
		return nil, err
	}

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	dAppConfig, err := dAppControlContract.GetDAppConfig(callOpts, dappcontrol.UserOperation{
		From:         userOp.GetFrom(),
		To:           userOp.GetTo(),
		Value:        userOp.GetValue(),
		Gas:          userOp.GetGas(),
		MaxFeePerGas: userOp.GetMaxFeePerGas(),
		Nonce:        userOp.GetNonce(),
		Deadline:     userOp.GetDeadline(),
		Dapp:         userOp.GetDapp(),
		Control:      userOp.GetControl(),
		CallConfig:   userOp.GetCallConfig(),
		SessionKey:   userOp.GetSessionKey(),
		Data:         userOp.GetData(),
		Signature:    userOp.GetSignature(),
	})
	if err != nil {
		return nil, err
	}

	return &types.DAppConfigLegacy{
		To:             dAppConfig.To,
		CallConfig:     dAppConfig.CallConfig,
		BidToken:       dAppConfig.BidToken,
		SolverGasLimit: dAppConfig.SolverGasLimit,
	}, nil
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

func (sdk *AtlasSdk) GetDAppGasLimit(chainId uint64, version *string, dAppControlAddr common.Address) (uint32, error) {
	minVersion := config.AtlasV_1_5
	if minSupport, err := config.IsVersionAtLeast(version, &minVersion); err != nil || !minSupport {
		return 0, fmt.Errorf("GetDAppGasLimit is only supported for Atlas v1.5 and above")
	}

	dAppControlContract, err := sdk.getDAppControlV15Contract(chainId, dAppControlAddr)
	if err != nil {
		return 0, err
	}

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	dAppGasLimit, err := dAppControlContract.GetDAppGasLimit(callOpts)
	if err != nil {
		return 0, err
	}

	return dAppGasLimit, nil
}
