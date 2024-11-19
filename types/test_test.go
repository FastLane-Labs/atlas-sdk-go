package types

import (
	"os"
	"testing"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	unitTestChainConfig = &config.ChainConfig{
		Contract: &config.Contract{
			Atlas:             common.HexToAddress("0x1"),
			AtlasVerification: common.HexToAddress("0x2"),
			Sorter:            common.HexToAddress("0x3"),
			Simulator:         common.HexToAddress("0x4"),
		},
		Eip712Domain: &apitypes.TypedDataDomain{
			Name:              "AtlasVerification",
			Version:           "1.0",
			ChainId:           math.NewHexOrDecimal256(1),
			VerifyingContract: "0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA",
		},
	}
)

func TestMain(m *testing.M) {
	version := config.AtlasVLatest
	if err := config.OverrideChainConfig(0, &version, unitTestChainConfig); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
