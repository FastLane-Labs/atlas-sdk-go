package core

import (
	"context"
	"fmt"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"
)

const (
	shMonadFunction  = "SHMONAD"
	policyIdFunction = "POLICY_ID"
)

type MonadConfig struct {
	ShMonad  common.Address
	PolicyId uint64
}

func (sdk *AtlasSdk) GetMonadConfig(chainId uint64, version *string) (*MonadConfig, error) {
	v := config.GetVersion(version)

	monadConfig, ok := sdk.monadConfig[chainId][v]
	if !ok {
		return nil, fmt.Errorf("monad config not found for chain id %d and version %s", chainId, v)
	}

	return monadConfig, nil
}

func (sdk *AtlasSdk) LoadMonadConfig(chainId uint64, version *string) (*MonadConfig, error) {
	var (
		shMonad  common.Address
		policyId uint64
		g, _     = errgroup.WithContext(context.Background())
	)

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

	g.Go(func() error {
		pData, err := atlasAbi.Pack(shMonadFunction)
		if err != nil {
			return fmt.Errorf("failed to pack %s: %w", shMonadFunction, err)
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
			return fmt.Errorf("failed to call %s: %w", shMonadFunction, err)
		}

		_shMonad, err := atlasAbi.Unpack(shMonadFunction, bData)
		if err != nil {
			return fmt.Errorf("failed to unpack %s: %w", shMonadFunction, err)
		}

		shMonad_, ok := _shMonad[0].(common.Address)
		if !ok {
			return fmt.Errorf("failed to cast %s: %w", shMonadFunction, err)
		}

		shMonad = shMonad_
		return nil
	})

	g.Go(func() error {
		pData, err := atlasAbi.Pack(policyIdFunction)
		if err != nil {
			return fmt.Errorf("failed to pack %s: %w", policyIdFunction, err)
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
			return fmt.Errorf("failed to call %s: %w", policyIdFunction, err)
		}

		_policyId, err := atlasAbi.Unpack(policyIdFunction, bData)
		if err != nil {
			return fmt.Errorf("failed to unpack %s: %w", policyIdFunction, err)
		}

		policyId_, ok := _policyId[0].(uint64)
		if !ok {
			return fmt.Errorf("failed to cast %s: %w", policyIdFunction, err)
		}

		policyId = policyId_
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &MonadConfig{ShMonad: shMonad, PolicyId: policyId}, nil
}
