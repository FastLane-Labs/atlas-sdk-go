package contract

import (
	"errors"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.2"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.3"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.2"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.3"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.2"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.3"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.0"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.1"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.2"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.3"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol"
)

var (
	atlas_1_0_Abi, _ = atlas_1_0.AtlasMetaData.GetAbi()
	atlas_1_1_Abi, _ = atlas_1_1.AtlasMetaData.GetAbi()
	atlas_1_2_Abi, _ = atlas_1_2.AtlasMetaData.GetAbi()
	atlas_1_3_Abi, _ = atlas_1_3.AtlasMetaData.GetAbi()

	atlasverification_1_0_Abi, _ = atlasverification_1_0.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_1_Abi, _ = atlasverification_1_1.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_2_Abi, _ = atlasverification_1_2.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_3_Abi, _ = atlasverification_1_3.AtlasVerificationMetaData.GetAbi()

	simulator_1_0_Abi, _ = simulator_1_0.SimulatorMetaData.GetAbi()
	simulator_1_1_Abi, _ = simulator_1_1.SimulatorMetaData.GetAbi()
	simulator_1_2_Abi, _ = simulator_1_2.SimulatorMetaData.GetAbi()
	simulator_1_3_Abi, _ = simulator_1_3.SimulatorMetaData.GetAbi()

	sorter_1_0_Abi, _ = sorter_1_0.SorterMetaData.GetAbi()
	sorter_1_1_Abi, _ = sorter_1_1.SorterMetaData.GetAbi()
	sorter_1_2_Abi, _ = sorter_1_2.SorterMetaData.GetAbi()
	sorter_1_3_Abi, _ = sorter_1_3.SorterMetaData.GetAbi()

	dappcontrol_Abi, _ = dappcontrol.DAppControlMetaData.GetAbi()
)

var (
	errInvalidAtlasVersion = errors.New("invalid atlas version")
)

func GetAtlasAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0:
		return atlas_1_0_Abi, nil
	case config.AtlasV_1_1:
		return atlas_1_1_Abi, nil
	case config.AtlasV_1_2:
		return atlas_1_2_Abi, nil
	case config.AtlasV_1_3:
		return atlas_1_3_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetAtlasVerificationAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0:
		return atlasverification_1_0_Abi, nil
	case config.AtlasV_1_1:
		return atlasverification_1_1_Abi, nil
	case config.AtlasV_1_2:
		return atlasverification_1_2_Abi, nil
	case config.AtlasV_1_3:
		return atlasverification_1_3_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetSimulatorAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0:
		return simulator_1_0_Abi, nil
	case config.AtlasV_1_1:
		return simulator_1_1_Abi, nil
	case config.AtlasV_1_2:
		return simulator_1_2_Abi, nil
	case config.AtlasV_1_3:
		return simulator_1_3_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetSorterAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0:
		return sorter_1_0_Abi, nil
	case config.AtlasV_1_1:
		return sorter_1_1_Abi, nil
	case config.AtlasV_1_2:
		return sorter_1_2_Abi, nil
	case config.AtlasV_1_3:
		return sorter_1_3_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetDAppControlAbi() *abi.ABI {
	return dappcontrol_Abi
}
