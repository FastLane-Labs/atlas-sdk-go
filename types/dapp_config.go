package types

import "github.com/ethereum/go-ethereum/common"

type DAppConfig struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
}
