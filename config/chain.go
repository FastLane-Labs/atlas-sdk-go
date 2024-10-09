package config

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Contract struct {
	Atlas             common.Address
	AtlasVerification common.Address
	Sorter            common.Address
	Simulator         common.Address
	Multicall3        common.Address
}

type ChainConfig struct {
	Contract     *Contract
	Eip712Domain *apitypes.TypedDataDomain
}

var (
	DEFAULT_MULTICALL3 = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

	chainConfig = map[uint64]*ChainConfig{
		// Unit tests
		0: {
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(1),
				VerifyingContract: "0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA",
			},
		},

		// Ethereum Sepolia
		11155111: {
			Contract: &Contract{
				Atlas:             common.HexToAddress("0x9EE12d2fed4B43F4Be37F69930CcaD9B65133482"),
				AtlasVerification: common.HexToAddress("0xB6F66a1b7cec02324D83c8DEA192818cA23A08B3"),
				Sorter:            common.HexToAddress("0xFE3c655d4D305Ac7f1c2F6306C79397560Afea0C"),
				Simulator:         common.HexToAddress("0xc3ab39ebd49D80bc36208545021224BAF6d2Bdb0"),
				Multicall3:        DEFAULT_MULTICALL3,
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(11155111),
				VerifyingContract: "0xB6F66a1b7cec02324D83c8DEA192818cA23A08B3",
			},
		},

		// Polygon Mainnet
		137: {
			Contract: &Contract{
				Atlas:             common.HexToAddress("0x912AceADa1b9c9B378894D0610C5684167710FDD"),
				AtlasVerification: common.HexToAddress("0x2fBF38a38D753E4ce398000CCC552Efa50702e1e"),
				Sorter:            common.HexToAddress("0xFac7bf300E7eb17A2eD0Be67b60f5FeDd2E28E90"),
				Simulator:         common.HexToAddress("0x1244E4B8D93D2A72692Bf3600f7f5a494e24895a"),
				Multicall3:        DEFAULT_MULTICALL3,
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(137),
				VerifyingContract: "0x2fBF38a38D753E4ce398000CCC552Efa50702e1e",
			},
		},

		// Polygon Amoy
		80002: {
			Contract: &Contract{
				Atlas:             common.HexToAddress("0x282BdDFF5e58793AcAb65438b257Dbd15A8745C9"),
				AtlasVerification: common.HexToAddress("0x3b7B38362bB7E2F000Cd2432343F3483F785F435"),
				Sorter:            common.HexToAddress("0xa55051bd82eFeA1dD487875C84fE9c016859659B"),
				Simulator:         common.HexToAddress("0x3efbaBE0ee916A4677D281c417E895a3e7411Ac2"),
				Multicall3:        DEFAULT_MULTICALL3,
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(80002),
				VerifyingContract: "0x3b7B38362bB7E2F000Cd2432343F3483F785F435",
			},
		},

		// BSC Mainnet
		56: {
			Contract: &Contract{
				Atlas:             common.HexToAddress("0xD72D821dA82964c0546a5501347a3959808E072f"),
				AtlasVerification: common.HexToAddress("0xae631aCDC436b9Dfd75C5629F825330d91459445"),
				Sorter:            common.HexToAddress("0xb47387995e866908B25b49e8BaC7e499170461A6"),
				Simulator:         common.HexToAddress("0xAb665f032e6A20Ef7D43FfD4E92a2f4fd6d5771e"),
				Multicall3:        DEFAULT_MULTICALL3,
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(56),
				VerifyingContract: "0xae631aCDC436b9Dfd75C5629F825330d91459445",
			},
		},

		// BSC testnet
		97: {
			Contract: &Contract{
				Atlas:             common.HexToAddress("0x164d3f6Bd3e78220f59e33729aA0E473C57EB067"),
				AtlasVerification: common.HexToAddress("0xa4445464B090D92CD2BD1c6F487F5f6284B26F7B"),
				Sorter:            common.HexToAddress("0xc41EEF5317a0e477Fef550f792AF5b5Bb9e13b34"),
				Simulator:         common.HexToAddress("0x91113eBDFD0c14BFAF48affD400800Bf851F6D11"),
				Multicall3:        DEFAULT_MULTICALL3,
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(97),
				VerifyingContract: "0xa4445464B090D92CD2BD1c6F487F5f6284B26F7B",
			},
		},
	}
)

func GetChainConfig(chainId uint64) (*ChainConfig, error) {
	if chainConfig[chainId] == nil {
		return nil, fmt.Errorf("chain config not found for chain id %d", chainId)
	}

	return chainConfig[chainId], nil
}

func OverrideChainConfig(chainId uint64, config *ChainConfig) error {
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

	if common.IsHexAddress(config.Eip712Domain.VerifyingContract) {
		return errors.New("eip712 domain verifying contract is required")
	}

	chainConfig[chainId] = config
	return nil
}

func GetAtlasAddress(chainId uint64) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Atlas, nil
}

func GetAtlasVerificationAddress(chainId uint64) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.AtlasVerification, nil
}

func GetSorterAddress(chainId uint64) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Sorter, nil
}

func GetSimulatorAddress(chainId uint64) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Simulator, nil
}

func GetMulticall3Address(chainId uint64) (common.Address, error) {
	chainConf, err := GetChainConfig(chainId)
	if err != nil {
		return common.Address{}, err
	}

	return chainConf.Contract.Multicall3, nil
}

func GetEip712Domain(chainId uint64) (*apitypes.TypedDataDomain, error) {
	chainConf, err := GetChainConfig(chainId)
	if err != nil {
		return nil, err
	}

	return chainConf.Eip712Domain, nil
}
