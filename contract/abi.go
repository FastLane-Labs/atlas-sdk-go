package contract

import (
	"errors"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.0.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.0.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.1.0"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.0.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.0.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.1.0"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.0.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.0.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.1.0"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.0.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.0.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.1.0"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol"
)

var (
	atlas_1_0_0_Abi, _ = atlas_1_0_0.AtlasMetaData.GetAbi()
	atlas_1_0_1_Abi, _ = atlas_1_0_1.AtlasMetaData.GetAbi()
	atlas_1_1_0_Abi, _ = atlas_1_1_0.AtlasMetaData.GetAbi()

	atlasverification_1_0_0_Abi, _ = atlasverification_1_0_0.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_0_1_Abi, _ = atlasverification_1_0_1.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_1_0_Abi, _ = atlasverification_1_1_0.AtlasVerificationMetaData.GetAbi()

	simulator_1_0_0_Abi, _ = simulator_1_0_0.SimulatorMetaData.GetAbi()
	simulator_1_0_1_Abi, _ = simulator_1_0_1.SimulatorMetaData.GetAbi()
	simulator_1_1_0_Abi, _ = simulator_1_1_0.SimulatorMetaData.GetAbi()

	sorter_1_0_0_Abi, _ = sorter_1_0_0.SorterMetaData.GetAbi()
	sorter_1_0_1_Abi, _ = sorter_1_0_1.SorterMetaData.GetAbi()
	sorter_1_1_0_Abi, _ = sorter_1_1_0.SorterMetaData.GetAbi()

	dappcontrol_Abi, _ = dappcontrol.DAppControlMetaData.GetAbi()
)

var (
	errInvalidAtlasVersion = errors.New("invalid atlas version")
)

func GetAtlasAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0_0:
		return atlas_1_0_0_Abi, nil
	case config.AtlasV_1_0_1:
		return atlas_1_0_1_Abi, nil
	case config.AtlasV_1_1_0:
		return atlas_1_1_0_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetAtlasVerificationAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0_0:
		return atlasverification_1_0_0_Abi, nil
	case config.AtlasV_1_0_1:
		return atlasverification_1_0_1_Abi, nil
	case config.AtlasV_1_1_0:
		return atlasverification_1_1_0_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetSimulatorAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0_0:
		return simulator_1_0_0_Abi, nil
	case config.AtlasV_1_0_1:
		return simulator_1_0_1_Abi, nil
	case config.AtlasV_1_1_0:
		return simulator_1_1_0_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetSorterAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0_0:
		return sorter_1_0_0_Abi, nil
	case config.AtlasV_1_0_1:
		return sorter_1_0_1_Abi, nil
	case config.AtlasV_1_1_0:
		return sorter_1_1_0_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetDAppControlAbi() *abi.ABI {
	return dappcontrol_Abi
}
