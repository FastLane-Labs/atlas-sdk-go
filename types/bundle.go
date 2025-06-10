package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// External representation of a bundle
type BundleRaw struct {
	ChainId          *hexutil.Big        `json:"chainId"`
	UserOperation    *UserOperationRaw   `json:"userOperation"`
	SolverOperations SolverOperationsRaw `json:"solverOperations"`
	DAppOperation    *DAppOperationRaw   `json:"dAppOperation"`
}

func (b *BundleRaw) Decode() (*Bundle, error) {
	userOp, err := b.UserOperation.Decode(b.ChainId.ToInt().Uint64())
	if err != nil {
		return nil, err
	}

	return &Bundle{
		ChainId:          b.ChainId.ToInt().Uint64(),
		UserOperation:    userOp,
		SolverOperations: b.SolverOperations.Decode(),
		DAppOperation:    b.DAppOperation.Decode(),
	}, nil
}

// Internal representation of a bundle
type Bundle struct {
	ChainId          uint64
	UserOperation    *UserOperation
	SolverOperations SolverOperations
	DAppOperation    *DAppOperation
}

func (b *Bundle) Sanitize() {
	b.UserOperation.Sanitize()
	b.DAppOperation.Sanitize()
	b.SolverOperations.Sanitize()
}

func (b *Bundle) EncodeToRaw() *BundleRaw {
	b.Sanitize()

	return &BundleRaw{
		ChainId:          (*hexutil.Big)(big.NewInt(int64(b.ChainId))),
		UserOperation:    b.UserOperation.EncodeToRaw(),
		SolverOperations: b.SolverOperations.EncodeToRaw(),
		DAppOperation:    b.DAppOperation.EncodeToRaw(),
	}
}
