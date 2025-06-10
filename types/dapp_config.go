package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DAppConfig struct {
	To                   common.Address `json:"to"`
	CallConfig           uint32         `json:"callConfig"`
	BidToken             common.Address `json:"bidToken"`
	SolverGasLimit       uint32         `json:"solverGasLimit"`
	DappGasLimit         uint32         `json:"dappGasLimit"`         // From Atlas v1.5
	BundlerSurchargeRate *big.Int       `json:"bundlerSurchargeRate"` // From Atlas v1.6
}
