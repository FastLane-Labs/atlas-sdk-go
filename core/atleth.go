package core

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	balanceOfBondedFunctionAtlEth = "balanceOfBonded"
	balanceOfBondedFunctionShMon  = "balanceOfBonded0"
)

// Returns bonded shMonad for Monad chains
func (sdk *AtlasSdk) GetBalanceOfBondedAtlEth(chainId uint64, version *string, account common.Address) (*big.Int, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	// Get the actual version that will be used (with fallback logic)
	chainConf, err := config.GetChainConfig(chainId, version)
	if err != nil {
		return nil, err
	}

	var (
		_abi        *abi.ABI
		destination common.Address
		function    string
		args        []interface{}
	)

	// Try to determine if this is a Monad version by checking if MonadConfig exists
	actualVersion, err := config.GetVersionFromAtlasAddress(chainId, chainConf.Contract.Atlas)
	isMonadVersion := err == nil && config.IsMonad(chainId) && strings.HasSuffix(actualVersion, "-monad")

	if isMonadVersion {
		monadConfig, err := sdk.GetMonadConfig(chainId, &actualVersion)
		if err != nil {
			return nil, err
		}

		shMonadAbi, err := contract.GetShMonadAbi()
		if err != nil {
			return nil, err
		}

		_abi = shMonadAbi
		function = balanceOfBondedFunctionShMon
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
		function = balanceOfBondedFunctionAtlEth
		destination = atlasAddr
		args = []interface{}{account}
	}

	pData, err := _abi.Pack(function, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s: %w", function, err)
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
		return nil, fmt.Errorf("failed to call %s: %w", function, err)
	}

	_balanceOfBonded, err := _abi.Unpack(function, bData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", function, err)
	}

	balanceOfBonded, ok := _balanceOfBonded[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to cast %s: %w", function, err)
	}

	return balanceOfBonded, nil
}
