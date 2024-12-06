package core

import (
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const (
	balanceOfBondedFunction = "balanceOfBonded"
)

func (sdk *AtlasSdk) GetBalanceOfBondedAtlEth(chainId uint64, version *string, account common.Address) (*big.Int, error) {
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

	pData, err := atlasAbi.Pack(balanceOfBondedFunction, account)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s: %w", balanceOfBondedFunction, err)
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:   &atlasAddr,
			Data: pData,
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", balanceOfBondedFunction, err)
	}

	_balanceOfBonded, err := atlasAbi.Unpack(balanceOfBondedFunction, bData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", balanceOfBondedFunction, err)
	}

	balanceOfBonded, ok := _balanceOfBonded[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to cast %s: %w", balanceOfBondedFunction, err)
	}

	return balanceOfBonded, nil
}
