package core

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AtlasSdk struct {
	ethClient map[uint64]*ethclient.Client

	atlasContract             map[uint64]*atlas.Atlas
	atlasVerificationContract map[uint64]*atlasverification.AtlasVerification
	simulatorContract         map[uint64]*simulator.Simulator
	sorterContract            map[uint64]*sorter.Sorter

	userLastNonSequentialNonce map[uint64]map[common.Address]*big.Int
	noncesMu                   sync.Mutex

	atlasAddress             map[uint64]common.Address
	atlasVerificationAddress map[uint64]common.Address
	simulatorAddress         map[uint64]common.Address
	sorterAddress            map[uint64]common.Address
}

func NewAtlasSdk(ethClient []*ethclient.Client, chainOverrides map[uint64]*config.ChainConfig) (*AtlasSdk, error) {
	for chainId, chainConf := range chainOverrides {
		err := config.OverrideChainConfig(chainId, chainConf)
		if err != nil {
			return nil, err
		}
	}

	sdk := &AtlasSdk{
		ethClient:                  make(map[uint64]*ethclient.Client),
		atlasContract:              make(map[uint64]*atlas.Atlas),
		atlasVerificationContract:  make(map[uint64]*atlasverification.AtlasVerification),
		simulatorContract:          make(map[uint64]*simulator.Simulator),
		sorterContract:             make(map[uint64]*sorter.Sorter),
		userLastNonSequentialNonce: make(map[uint64]map[common.Address]*big.Int),
		atlasAddress:               make(map[uint64]common.Address),
		atlasVerificationAddress:   make(map[uint64]common.Address),
		simulatorAddress:           make(map[uint64]common.Address),
		sorterAddress:              make(map[uint64]common.Address),
	}

	for _, client := range ethClient {
		ctx, cancel := NewContextWithNetworkDeadline()
		defer cancel()

		chainId, err := client.ChainID(ctx)
		if err != nil {
			return nil, err
		}

		chainConf, err := config.GetChainConfig(chainId.Uint64())
		if err != nil {
			return nil, err
		}

		atlasContract, err := atlas.NewAtlas(chainConf.Contract.Atlas, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create atlas contract: %w", err)
		}

		atlasVerificationContract, err := atlasverification.NewAtlasVerification(chainConf.Contract.AtlasVerification, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create atlas verification contract: %w", err)
		}

		simulatorContract, err := simulator.NewSimulator(chainConf.Contract.Simulator, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create simulator contract: %w", err)
		}

		sorterContract, err := sorter.NewSorter(chainConf.Contract.Sorter, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create sorter contract: %w", err)
		}

		chainIdUint64 := chainId.Uint64()
		sdk.ethClient[chainIdUint64] = client
		sdk.atlasContract[chainIdUint64] = atlasContract
		sdk.atlasVerificationContract[chainIdUint64] = atlasVerificationContract
		sdk.simulatorContract[chainIdUint64] = simulatorContract
		sdk.sorterContract[chainIdUint64] = sorterContract
		sdk.userLastNonSequentialNonce[chainIdUint64] = make(map[common.Address]*big.Int)
		sdk.atlasAddress[chainIdUint64] = chainConf.Contract.Atlas
		sdk.atlasVerificationAddress[chainIdUint64] = chainConf.Contract.AtlasVerification
		sdk.simulatorAddress[chainIdUint64] = chainConf.Contract.Simulator
		sdk.sorterAddress[chainIdUint64] = chainConf.Contract.Sorter
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

func (sdk *AtlasSdk) GetAtlasContract(chainId uint64) (*atlas.Atlas, error) {
	contract, ok := sdk.atlasContract[chainId]
	if !ok {
		return nil, fmt.Errorf("atlas contract not found for chain id %d", chainId)
	}

	return contract, nil
}

func (sdk *AtlasSdk) GetAtlasVerificationContract(chainId uint64) (*atlasverification.AtlasVerification, error) {
	contract, ok := sdk.atlasVerificationContract[chainId]
	if !ok {
		return nil, fmt.Errorf("atlas verification contract not found for chain id %d", chainId)
	}

	return contract, nil
}

func (sdk *AtlasSdk) GetSimulatorContract(chainId uint64) (*simulator.Simulator, error) {
	contract, ok := sdk.simulatorContract[chainId]
	if !ok {
		return nil, fmt.Errorf("simulator contract not found for chain id %d", chainId)
	}

	return contract, nil
}

func (sdk *AtlasSdk) GetSorterContract(chainId uint64) (*sorter.Sorter, error) {
	contract, ok := sdk.sorterContract[chainId]
	if !ok {
		return nil, fmt.Errorf("sorter contract not found for chain id %d", chainId)
	}

	return contract, nil
}
