package contract

import (
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlas"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/atlasverification"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/dappcontrol"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/simulator"
	"github.com/FastLane-Labs/atlas-sdk-go/contract/sorter"
)

var (
	AtlasAbi, _             = atlas.AtlasMetaData.GetAbi()
	AtlasVerificationAbi, _ = atlasverification.AtlasVerificationMetaData.GetAbi()
	SimulatorAbi, _         = simulator.SimulatorMetaData.GetAbi()
	SorterAbi, _            = sorter.SorterMetaData.GetAbi()
	DAppControlAbi, _       = dappcontrol.DAppControlMetaData.GetAbi()
)
