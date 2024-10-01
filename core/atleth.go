package core

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (sdk *AtlasSdk) GetBalanceOfBondedAtlEth(chainId uint64, account common.Address) (*big.Int, error) {
	contract, ok := sdk.atlasContract[chainId]
	if !ok {
		return nil, errors.New("atlas contract not found")
	}

	return contract.BalanceOfBonded(nil, account)
}
