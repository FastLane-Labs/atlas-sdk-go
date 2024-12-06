package core

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AtlasSdk struct {
	ethClient map[uint64]*ethclient.Client

	userLastNonSequentialNonce map[uint64]map[string]map[common.Address]*big.Int
	noncesMu                   map[uint64]map[string]*sync.Mutex

	mu sync.Mutex
}

func NewAtlasSdk(ethClient []*ethclient.Client, chainOverrides map[uint64]map[string]*config.ChainConfig) (*AtlasSdk, error) {
	for chainId, chainConf := range chainOverrides {
		for version, conf := range chainConf {
			if err := config.OverrideChainConfig(chainId, &version, conf); err != nil {
				return nil, err
			}
		}
	}

	sdk := &AtlasSdk{
		ethClient:                  make(map[uint64]*ethclient.Client),
		userLastNonSequentialNonce: make(map[uint64]map[string]map[common.Address]*big.Int),
		noncesMu:                   make(map[uint64]map[string]*sync.Mutex),
	}

	for _, client := range ethClient {
		ctx, cancel := NewContextWithNetworkDeadline()
		defer cancel()

		chainId, err := client.ChainID(ctx)
		if err != nil {
			return nil, err
		}

		chainIdUint64 := chainId.Uint64()
		sdk.ethClient[chainIdUint64] = client

		for _, version := range config.GetAllVersions() {
			if _, ok := sdk.userLastNonSequentialNonce[chainIdUint64]; !ok {
				sdk.userLastNonSequentialNonce[chainIdUint64] = make(map[string]map[common.Address]*big.Int)
			}

			sdk.userLastNonSequentialNonce[chainIdUint64][version] = make(map[common.Address]*big.Int)

			if _, ok := sdk.noncesMu[chainIdUint64]; !ok {
				sdk.noncesMu[chainIdUint64] = make(map[string]*sync.Mutex)
			}

			sdk.noncesMu[chainIdUint64][version] = &sync.Mutex{}
		}
	}

	return sdk, nil
}

func (sdk *AtlasSdk) getEthClient(chainId uint64) (*ethclient.Client, error) {
	client, ok := sdk.ethClient[chainId]
	if !ok {
		return nil, fmt.Errorf("ethereum client not found for chain id %d", chainId)
	}

	return client, nil
}
