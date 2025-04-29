package config

import (
	"errors"
	"fmt"
	"math/big"
	"slices"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Contract struct {
	Atlas             common.Address `json:"atlas"`
	AtlasVerification common.Address `json:"atlasVerification"`
	Sorter            common.Address `json:"sorter"`
	Simulator         common.Address `json:"simulator"`
	Multicall3        common.Address `json:"multicall3"`
}

type ChainConfig struct {
	Contract     *Contract                 `json:"contracts"`
	Eip712Domain *apitypes.TypedDataDomain `json:"eip712Domain"`
}

const (
	AtlasV_1_0       = "1.0"
	AtlasV_1_1       = "1.1"
	AtlasV_1_2       = "1.2"
	AtlasV_1_3       = "1.3"
	AtlasV_1_5       = "1.5"
	AtlasV_1_5_MONAD = "1.5-monad"
	AtlasVLatest     = AtlasV_1_5
)

var (
	DEFAULT_MULTICALL3 = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	chainConfig        = map[uint64]map[string]*ChainConfig{} // Indexed by [chainId][atlasVersion]
	initOnce           sync.Once
	mu                 sync.RWMutex

	allVersions      = []string{AtlasV_1_0, AtlasV_1_1, AtlasV_1_2, AtlasV_1_3, AtlasV_1_5, AtlasV_1_5_MONAD}
	allMonadVersions = []string{AtlasV_1_5_MONAD}
)

func IsVersionAtLeast(version *string, minVersion *string) (bool, error) {
	if version == nil || minVersion == nil {
		return false, errors.New("version or minVersion is nil")
	}

	idx1 := slices.Index(allVersions, *version)
	if idx1 == -1 {
		return false, fmt.Errorf("version not found in allVersions")
	}

	idx2 := slices.Index(allVersions, *minVersion)
	if idx2 == -1 {
		return false, fmt.Errorf("minVersion not found in allVersions")
	}

	return idx1 >= idx2, nil
}

func IsVersionLatest(version *string) (bool, error) {
	if version == nil {
		return false, fmt.Errorf("version is nil")
	}

	return *version == AtlasVLatest, nil
}

func InitChainConfig() error {
	var initErr error

	initOnce.Do(func() {
		mu.RLock()

		if len(chainConfig) > 0 {
			// This can happen when chain config has been overridden before initChainConfig was called
			mu.RUnlock()
			return
		}

		mu.RUnlock()

		remoteConfig, err := downloadChainConfig()
		if err != nil {
			initErr = err
			return
		}

		mu.Lock()
		defer mu.Unlock()

		for chainId, config := range remoteConfig {
			chainConfig[chainId] = config

			if IsMonad(chainId) {
				for version := range config {
					monadVersion := ToMonadVersion(&version)
					chainConfig[chainId][monadVersion] = config[version]
				}
			}
		}
	})

	return initErr
}

func GetAllVersions() []string {
	v := make([]string, len(allVersions))
	copy(v, allVersions)
	return v
}

func GetAllMonadVersions() []string {
	v := make([]string, len(allMonadVersions))
	copy(v, allMonadVersions)
	return v
}

func GetVersion(version *string) string {
	if version == nil {
		return AtlasVLatest
	}

	return *version
}

func GetVersionFromAtlasAddress(chainId uint64, atlasAddr common.Address) (string, error) {
	chainConfigInitialized := len(chainConfig) > 0
	if !chainConfigInitialized {
		if err := InitChainConfig(); err != nil {
			return "", err
		}
	}

	mu.RLock()
	defer mu.RUnlock()

	for _, version := range allVersions {
		chainConf, ok := chainConfig[chainId][version]
		if !ok {
			continue
		}

		if chainConf.Contract.Atlas == atlasAddr {
			return version, nil
		}
	}

	return "", fmt.Errorf("atlas address not found for chain id %d", chainId)
}

func GetChainConfig(chainId uint64, version *string) (*ChainConfig, error) {
	chainConfigInitialized := len(chainConfig) > 0
	if !chainConfigInitialized {
		if err := InitChainConfig(); err != nil {
			return nil, err
		}
	}

	v := GetVersion(version)

	if IsMonad(chainId) {
		v = ToMonadVersion(&v)
	}

	mu.RLock()
	defer mu.RUnlock()

	_chainConfig, ok := chainConfig[chainId][v]
	if !ok {
		return nil, fmt.Errorf("chain config not found for chain id %d and version %s", chainId, v)
	}

	return _chainConfig, nil
}

func OverrideChainConfig(chainId uint64, version *string, config *ChainConfig) error {
	if config.Contract == nil {
		return errors.New("contract config is required")
	}

	if config.Contract.Atlas == (common.Address{}) {
		return errors.New("atlas contract address is required")
	}

	if config.Contract.AtlasVerification == (common.Address{}) {
		return errors.New("atlas verification contract address is required")
	}

	if config.Contract.Sorter == (common.Address{}) {
		return errors.New("sorter contract address is required")
	}

	if config.Contract.Simulator == (common.Address{}) {
		return errors.New("simulator contract address is required")
	}

	if config.Contract.Multicall3 == (common.Address{}) {
		config.Contract.Multicall3 = DEFAULT_MULTICALL3
	}

	if config.Eip712Domain == nil {
		return errors.New("eip712 domain is required")
	}

	if len(config.Eip712Domain.Name) == 0 {
		return errors.New("eip712 domain name is required")
	}

	if len(config.Eip712Domain.Version) == 0 {
		return errors.New("eip712 domain version is required")
	}

	if config.Eip712Domain.ChainId == nil {
		return errors.New("eip712 domain chain id is required")
	}

	if !common.IsHexAddress(config.Eip712Domain.VerifyingContract) {
		return errors.New("eip712 domain verifying contract is invalid")
	}

	v := GetVersion(version)

	mu.Lock()
	defer mu.Unlock()

	if chainConfig[chainId] == nil {
		chainConfig[chainId] = make(map[string]*ChainConfig)
	}

	chainConfig[chainId][v] = config
	return nil
}

func GetAtlasAddress(chainId uint64, version *string) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId, version)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Atlas, nil
}

func GetAtlasVerificationAddress(chainId uint64, version *string) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId, version)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.AtlasVerification, nil
}

func GetSorterAddress(chainId uint64, version *string) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId, version)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Sorter, nil
}

func GetSimulatorAddress(chainId uint64, version *string) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId, version)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Simulator, nil
}

func GetMulticall3Address(chainId uint64, version *string) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId, version)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Multicall3, nil
}

func GetEip712Domain(chainId uint64, version *string) (*apitypes.TypedDataDomain, error) {
	chainConf, err := GetChainConfig(chainId, version)
	if err != nil {
		return nil, err
	}

	_chainId := new(big.Int).Set((*big.Int)(chainConf.Eip712Domain.ChainId))

	eip712Domain := &apitypes.TypedDataDomain{
		Name:              chainConf.Eip712Domain.Name,
		Version:           chainConf.Eip712Domain.Version,
		ChainId:           (*math.HexOrDecimal256)(_chainId),
		VerifyingContract: chainConf.Eip712Domain.VerifyingContract,
	}

	return eip712Domain, nil
}
