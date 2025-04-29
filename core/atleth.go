package core

import (
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	balanceOfBondedFunction = "balanceOfBonded"
)

// Returns bonded shMonad for Monad chains
func (sdk *AtlasSdk) GetBalanceOfBondedAtlEth(chainId uint64, version *string, account common.Address) (*big.Int, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	var (
		_abi        *abi.ABI
		destination common.Address
		args        []interface{}
	)

	if config.IsMonad(chainId) {
		_version := config.ToMonadVersion(version)

		monadConfig, err := sdk.GetMonadConfig(chainId, &_version)
		if err != nil {
			return nil, err
		}

		shMonadAbi, err := contract.GetShMonadAbi()
		if err != nil {
			return nil, err
		}

		_abi = shMonadAbi
		destination = monadConfig.ShMonad
		args = []interface{}{monadConfig.PolicyId, account}
	} else {
		atlasAddr, err := config.GetAtlasAddress(chainId, version)
		if err != nil {
			return nil, err
		}

		atlasAbi, err := contract.GetAtlasAbi(version)
		if err != nil {
			return nil, err
		}

		_abi = atlasAbi
		destination = atlasAddr
		args = []interface{}{account}
	}

	pData, err := _abi.Pack(balanceOfBondedFunction, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s: %w", balanceOfBondedFunction, err)
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:   &destination,
			Data: pData,
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", balanceOfBondedFunction, err)
	}

	_balanceOfBonded, err := _abi.Unpack(balanceOfBondedFunction, bData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", balanceOfBondedFunction, err)
	}

	balanceOfBonded, ok := _balanceOfBonded[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to cast %s: %w", balanceOfBondedFunction, err)
	}

	return balanceOfBonded, nil
}
