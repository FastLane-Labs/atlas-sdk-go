package config

import (
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
				Multicall3:        common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
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
				Atlas:             common.HexToAddress("0x57FA2aBf1dc109C5F7ea2FB6A72358D2c624971d"),
				AtlasVerification: common.HexToAddress("0xA462C35C43355928F114144AD20AddD6Bb09b52f"),
				Sorter:            common.HexToAddress("0x0cb4cCc7C853BA2CeA8bc8cB2ECF8142dF67BF79"),
				Simulator:         common.HexToAddress("0x7f9227d40590D473D9FdD855C506f2D6400687Cb"),
				Multicall3:        common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(137),
				VerifyingContract: "0xA462C35C43355928F114144AD20AddD6Bb09b52f",
			},
		},

		// Polygon Amoy
		80002: {
			Contract: &Contract{
				Atlas:             common.HexToAddress("0x282BdDFF5e58793AcAb65438b257Dbd15A8745C9"),
				AtlasVerification: common.HexToAddress("0x3b7B38362bB7E2F000Cd2432343F3483F785F435"),
				Sorter:            common.HexToAddress("0xa55051bd82eFeA1dD487875C84fE9c016859659B"),
				Simulator:         common.HexToAddress("0x3efbaBE0ee916A4677D281c417E895a3e7411Ac2"),
				Multicall3:        common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
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
				Multicall3:        common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			},
			Eip712Domain: &apitypes.TypedDataDomain{
				Name:              "AtlasVerification",
				Version:           "1.0",
				ChainId:           math.NewHexOrDecimal256(56),
				VerifyingContract: "0xae631aCDC436b9Dfd75C5629F825330d91459445",
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

func OverrideChainConfig(chainId uint64, config *ChainConfig) {
	chainConfig[chainId] = config
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
