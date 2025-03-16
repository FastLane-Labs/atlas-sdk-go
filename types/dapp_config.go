package types

import "github.com/ethereum/go-ethereum/common"

type DAppConfig interface {
	GetTo() common.Address
	GetCallConfig() uint32
	GetBidToken() common.Address
	GetSolverGasLimit() uint32
	GetDappGasLimit() uint32
}

var (
	_ DAppConfig = &DAppConfigLegacy{}
	_ DAppConfig = &DAppConfigV15{}
)

type DAppConfigLegacy struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
}

type DAppConfigV15 struct {
	DAppConfigLegacy
	DappGasLimit uint32
}

func (c *DAppConfigV15) GetDappGasLimit() uint32 {
	return c.DappGasLimit
}

func (c *DAppConfigLegacy) GetBidToken() common.Address {
	return c.BidToken
}

func (c *DAppConfigLegacy) GetSolverGasLimit() uint32 {
	return c.SolverGasLimit
}

func (c *DAppConfigLegacy) GetCallConfig() uint32 {
	return c.CallConfig
}

func (c *DAppConfigLegacy) GetTo() common.Address {
	return c.To
}

func (c *DAppConfigLegacy) GetDappGasLimit() uint32 {
	return 0
}
