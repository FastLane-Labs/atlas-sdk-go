package types

import (
	"errors"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// External representation of a dApp operation,
type DAppOperationRaw struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Nonce         *hexutil.Big   `json:"nonce"`
	Deadline      *hexutil.Big   `json:"deadline"`
	Control       common.Address `json:"control"`
	Bundler       common.Address `json:"bundler"`
	UserOpHash    common.Hash    `json:"userOpHash"`
	CallChainHash common.Hash    `json:"callChainHash"`
	Signature     hexutil.Bytes  `json:"signature"`
}

func (d *DAppOperationRaw) Sanitize() {
	if d.Nonce == nil {
		d.Nonce = (*hexutil.Big)(big.NewInt(0))
	}

	if d.Deadline == nil {
		d.Deadline = (*hexutil.Big)(big.NewInt(0))
	}
}

func (d *DAppOperationRaw) Decode() *DAppOperation {
	d.Sanitize()

	return &DAppOperation{
		From:          d.From,
		To:            d.To,
		Nonce:         d.Nonce.ToInt(),
		Deadline:      d.Deadline.ToInt(),
		Control:       d.Control,
		Bundler:       d.Bundler,
		UserOpHash:    d.UserOpHash,
		CallChainHash: d.CallChainHash,
		Signature:     d.Signature,
	}
}

// Internal representation of a dApp operation
type DAppOperation struct {
	From          common.Address
	To            common.Address
	Nonce         *big.Int
	Deadline      *big.Int
	Control       common.Address
	Bundler       common.Address
	UserOpHash    common.Hash
	CallChainHash common.Hash
	Signature     []byte
}

func (d *DAppOperation) Sanitize() {
	if d.Nonce == nil {
		d.Nonce = big.NewInt(0)
	}

	if d.Deadline == nil {
		d.Deadline = big.NewInt(0)
	}
}

func (d *DAppOperation) EncodeToRaw() *DAppOperationRaw {
	d.Sanitize()

	return &DAppOperationRaw{
		From:          d.From,
		To:            d.To,
		Nonce:         (*hexutil.Big)(d.Nonce),
		Deadline:      (*hexutil.Big)(d.Deadline),
		Control:       d.Control,
		Bundler:       d.Bundler,
		UserOpHash:    d.UserOpHash,
		CallChainHash: d.CallChainHash,
		Signature:     d.Signature,
	}
}

func (d *DAppOperation) Hash(chainId uint64) (common.Hash, error) {
	eip712Domain, err := config.GetEip712Domain(chainId)
	if err != nil {
		return common.Hash{}, err
	}

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       d.toTypedDataTypes(),
		PrimaryType: "DAppOperation",
		Domain:      *eip712Domain,
		Message:     d.toTypedDataMessage(),
	})

	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(hash), nil
}

func (d *DAppOperation) ValidateSignature(chainId uint64) error {
	if len(d.Signature) != 65 {
		return errors.New("invalid signature length")
	}

	dAppOpHash, err := d.Hash(chainId)
	if err != nil {
		return err
	}

	signer, err := utils.RecoverSigner(dAppOpHash.Bytes(), d.Signature)
	if err != nil {
		return err
	}

	if signer != d.From {
		return errors.New("invalid signature")
	}

	return nil
}

func (d *DAppOperation) toTypedDataTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"DAppOperation": []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "nonce", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "control", Type: "address"},
			{Name: "bundler", Type: "address"},
			{Name: "userOpHash", Type: "bytes32"},
			{Name: "callChainHash", Type: "bytes32"},
		},
	}
}

func (d *DAppOperation) toTypedDataMessage() apitypes.TypedDataMessage {
	d.Sanitize()

	return apitypes.TypedDataMessage{
		"from":          d.From.Hex(),
		"to":            d.To.Hex(),
		"nonce":         new(big.Int).Set(d.Nonce),
		"deadline":      new(big.Int).Set(d.Deadline),
		"control":       d.Control.Hex(),
		"bundler":       d.Bundler.Hex(),
		"userOpHash":    d.UserOpHash.Hex(),
		"callChainHash": d.CallChainHash.Hex(),
	}
}
