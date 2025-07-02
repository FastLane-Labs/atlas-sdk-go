package contract

import (
	"errors"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/ethereum/go-ethereum/accounts/abi"

	atlas_1_0 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.0"
	atlas_1_1 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.1"
	atlas_1_2 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.2"
	atlas_1_3 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.3"
	atlas_1_5 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.5"
	atlas_1_5_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.5-monad"
	atlas_1_6 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.6"
	atlas_1_6_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/atlas/1.6-monad"

	atlasverification_1_0 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.0"
	atlasverification_1_1 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.1"
	atlasverification_1_2 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.2"
	atlasverification_1_3 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.3"
	atlasverification_1_5 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.5"
	atlasverification_1_5_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.5-monad"
	atlasverification_1_6 "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.6"
	atlasverification_1_6_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification/1.6-monad"

	simulator_1_0 "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.0"
	simulator_1_1 "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.1"
	simulator_1_2 "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.2"
	simulator_1_3 "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.3"
	simulator_1_5 "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.5"
	simulator_1_5_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.5-monad"
	simulator_1_6 "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.6"
	simulator_1_6_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/simulator/1.6-monad"

	sorter_1_0 "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.0"
	sorter_1_1 "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.1"
	sorter_1_2 "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.2"
	sorter_1_3 "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.3"
	sorter_1_5 "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.5"
	sorter_1_5_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.5-monad"
	sorter_1_6 "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.6"
	sorter_1_6_monad "github.com/FastLane-Labs/atlas-sdk-go/contract/sorter/1.6-monad"

	dappcontrol_1_5 "github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol/1.5"
	dappcontrol_1_6 "github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol/1.6"
	dappcontrol_legacy "github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol/legacy"

	"github.com/FastLane-Labs/atlas-sdk-go/contract/shmonad"
)

var (
	atlas_1_0_Abi, _       = atlas_1_0.AtlasMetaData.GetAbi()
	atlas_1_1_Abi, _       = atlas_1_1.AtlasMetaData.GetAbi()
	atlas_1_2_Abi, _       = atlas_1_2.AtlasMetaData.GetAbi()
	atlas_1_3_Abi, _       = atlas_1_3.AtlasMetaData.GetAbi()
	atlas_1_5_Abi, _       = atlas_1_5.AtlasMetaData.GetAbi()
	atlas_1_5_monad_Abi, _ = atlas_1_5_monad.AtlasMetaData.GetAbi()
	atlas_1_6_Abi, _       = atlas_1_6.AtlasMetaData.GetAbi()
	atlas_1_6_monad_Abi, _ = atlas_1_6_monad.AtlasMetaData.GetAbi()

	atlasverification_1_0_Abi, _       = atlasverification_1_0.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_1_Abi, _       = atlasverification_1_1.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_2_Abi, _       = atlasverification_1_2.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_3_Abi, _       = atlasverification_1_3.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_5_Abi, _       = atlasverification_1_5.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_5_monad_Abi, _ = atlasverification_1_5_monad.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_6_Abi, _       = atlasverification_1_6.AtlasVerificationMetaData.GetAbi()
	atlasverification_1_6_monad_Abi, _ = atlasverification_1_6_monad.AtlasVerificationMetaData.GetAbi()

	simulator_1_0_Abi, _       = simulator_1_0.SimulatorMetaData.GetAbi()
	simulator_1_1_Abi, _       = simulator_1_1.SimulatorMetaData.GetAbi()
	simulator_1_2_Abi, _       = simulator_1_2.SimulatorMetaData.GetAbi()
	simulator_1_3_Abi, _       = simulator_1_3.SimulatorMetaData.GetAbi()
	simulator_1_5_Abi, _       = simulator_1_5.SimulatorMetaData.GetAbi()
	simulator_1_5_monad_Abi, _ = simulator_1_5_monad.SimulatorMetaData.GetAbi()
	simulator_1_6_Abi, _       = simulator_1_6.SimulatorMetaData.GetAbi()
	simulator_1_6_monad_Abi, _ = simulator_1_6_monad.SimulatorMetaData.GetAbi()

	sorter_1_0_Abi, _       = sorter_1_0.SorterMetaData.GetAbi()
	sorter_1_1_Abi, _       = sorter_1_1.SorterMetaData.GetAbi()
	sorter_1_2_Abi, _       = sorter_1_2.SorterMetaData.GetAbi()
	sorter_1_3_Abi, _       = sorter_1_3.SorterMetaData.GetAbi()
	sorter_1_5_Abi, _       = sorter_1_5.SorterMetaData.GetAbi()
	sorter_1_5_monad_Abi, _ = sorter_1_5_monad.SorterMetaData.GetAbi()
	sorter_1_6_Abi, _       = sorter_1_6.SorterMetaData.GetAbi()
	sorter_1_6_monad_Abi, _ = sorter_1_6_monad.SorterMetaData.GetAbi()

	dappcontrol_legacy_Abi, _ = dappcontrol_legacy.DAppControlMetaData.GetAbi()
	dappcontrol_1_5_Abi, _    = dappcontrol_1_5.DAppControlMetaData.GetAbi()
	dappcontrol_1_6_Abi, _    = dappcontrol_1_6.DAppControlMetaData.GetAbi()

	shmonad_Abi, _ = shmonad.ShMonadMetaData.GetAbi()
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
	case config.AtlasV_1_5:
		return atlas_1_5_Abi, nil
	case config.AtlasV_1_5_MONAD:
		return atlas_1_5_monad_Abi, nil
	case config.AtlasV_1_6, config.AtlasV_1_6_1, config.AtlasV_1_6_2, config.AtlasV_1_7:
		return atlas_1_6_Abi, nil
	case config.AtlasV_1_6_1_MONAD:
		return atlas_1_6_monad_Abi, nil
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
	case config.AtlasV_1_5:
		return atlasverification_1_5_Abi, nil
	case config.AtlasV_1_5_MONAD:
		return atlasverification_1_5_monad_Abi, nil
	case config.AtlasV_1_6, config.AtlasV_1_6_1, config.AtlasV_1_6_2, config.AtlasV_1_7:
		return atlasverification_1_6_Abi, nil
	case config.AtlasV_1_6_1_MONAD:
		return atlasverification_1_6_monad_Abi, nil
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
	case config.AtlasV_1_5:
		return simulator_1_5_Abi, nil
	case config.AtlasV_1_5_MONAD:
		return simulator_1_5_monad_Abi, nil
	case config.AtlasV_1_6, config.AtlasV_1_6_1, config.AtlasV_1_6_2, config.AtlasV_1_7:
		return simulator_1_6_Abi, nil
	case config.AtlasV_1_6_1_MONAD:
		return simulator_1_6_monad_Abi, nil
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
	case config.AtlasV_1_5:
		return sorter_1_5_Abi, nil
	case config.AtlasV_1_5_MONAD:
		return sorter_1_5_monad_Abi, nil
	case config.AtlasV_1_6, config.AtlasV_1_6_1, config.AtlasV_1_6_2, config.AtlasV_1_7:
		return sorter_1_6_Abi, nil
	case config.AtlasV_1_6_1_MONAD:
		return sorter_1_6_monad_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetDAppControlAbi(version *string) (*abi.ABI, error) {
	switch config.GetVersion(version) {
	case config.AtlasV_1_0, config.AtlasV_1_1, config.AtlasV_1_2, config.AtlasV_1_3:
		return dappcontrol_legacy_Abi, nil
	case config.AtlasV_1_5, config.AtlasV_1_5_MONAD:
		return dappcontrol_1_5_Abi, nil
	case config.AtlasV_1_6, config.AtlasV_1_6_1, config.AtlasV_1_6_2, config.AtlasV_1_7:
		return dappcontrol_1_6_Abi, nil
	case config.AtlasV_1_6_1_MONAD:
		return dappcontrol_1_6_Abi, nil
	}
	return nil, errInvalidAtlasVersion
}

func GetShMonadAbi() (*abi.ABI, error) {
	return shmonad_Abi, nil
}
