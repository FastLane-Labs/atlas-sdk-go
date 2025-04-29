package core

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/blockchain-rpc-go/eth"
	"github.com/FastLane-Labs/blockchain-rpc-go/rpc"
	"github.com/ethereum/go-ethereum/common"
	gethEthClient "github.com/ethereum/go-ethereum/ethclient"
	gethRpc "github.com/ethereum/go-ethereum/rpc"
)

type AtlasSdk struct {
	ethClient map[uint64]eth.IEthClient

	userLastNonSequentialNonce map[uint64]map[string]map[common.Address]*big.Int
	noncesMu                   map[uint64]map[string]*sync.Mutex

	monadConfig map[uint64]map[string]*MonadConfig

	mu sync.Mutex
}

func NewAtlasSdk(rpcClients []interface{}, chainOverrides map[uint64]map[string]*config.ChainConfig) (*AtlasSdk, error) {
	if err := config.InitChainConfig(); err != nil {
		return nil, err
	}

	for chainId, chainConf := range chainOverrides {
		for version, conf := range chainConf {
			if err := config.OverrideChainConfig(chainId, &version, conf); err != nil {
				return nil, err
			}
		}
	}

	sdk := &AtlasSdk{
		ethClient:                  make(map[uint64]eth.IEthClient),
		userLastNonSequentialNonce: make(map[uint64]map[string]map[common.Address]*big.Int),
		noncesMu:                   make(map[uint64]map[string]*sync.Mutex),
		monadConfig:                make(map[uint64]map[string]*MonadConfig),
	}

	for _, rpcClient := range rpcClients {
		var client *eth.EthClient

		switch v := rpcClient.(type) {
		case *eth.EthClient:
			client = v
		case *rpc.RpcClient:
			client = eth.NewClient(v)
		case *gethRpc.Client:
			client = eth.NewClient(v)
		case *gethEthClient.Client:
			client = eth.NewClient(v.Client())
		default:
			return nil, fmt.Errorf("unsupported rpc client type: %T", v)
		}

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

		if config.IsMonad(chainIdUint64) {
			for _, version := range config.GetAllMonadVersions() {
				monadConfig, err := sdk.LoadMonadConfig(chainIdUint64, &version)
				if err != nil {
					return nil, err
				}

				if sdk.monadConfig[chainIdUint64] == nil {
					sdk.monadConfig[chainIdUint64] = make(map[string]*MonadConfig)
				}

				sdk.monadConfig[chainIdUint64][version] = monadConfig
			}
		}
	}

	return sdk, nil
}

func (sdk *AtlasSdk) getEthClient(chainId uint64) (eth.IEthClient, error) {
	client, ok := sdk.ethClient[chainId]
	if !ok {
		return nil, fmt.Errorf("ethereum client not found for chain id %d", chainId)
	}

	return client, nil
}
