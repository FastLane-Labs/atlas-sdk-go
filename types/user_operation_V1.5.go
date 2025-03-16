package types

import (
	"errors"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type UserOperationV15 struct {
	UserOperationLegacy
	DappGasLimit *big.Int
}

func (u *UserOperationV15) GetDappGasLimit() *big.Int {
	return u.DappGasLimit
}

func (u *UserOperationV15) toTypedDataTypes(trusted bool) apitypes.Types {
	t := apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}

	if trusted {
		t["UserOperation"] = []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "dapp", Type: "address"},
			{Name: "control", Type: "address"},
			{Name: "callConfig", Type: "uint32"},
			{Name: "dappGasLimit", Type: "uint256"},
			{Name: "sessionKey", Type: "address"},
		}
	} else {
		t["UserOperation"] = []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "gas", Type: "uint256"},
			{Name: "maxFeePerGas", Type: "uint256"},
			{Name: "nonce", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "dapp", Type: "address"},
			{Name: "control", Type: "address"},
			{Name: "callConfig", Type: "uint32"},
			{Name: "dappGasLimit", Type: "uint256"},
			{Name: "sessionKey", Type: "address"},
			{Name: "data", Type: "bytes"},
		}
	}

	return t
}

func (u *UserOperationV15) EncodeToRaw() *UserOperationRaw {
	u.Sanitize()

	return &UserOperationRaw{
		From:         u.From,
		To:           u.To,
		Value:        (*hexutil.Big)(u.Value),
		Gas:          (*hexutil.Big)(u.Gas),
		MaxFeePerGas: (*hexutil.Big)(u.MaxFeePerGas),
		Nonce:        (*hexutil.Big)(u.Nonce),
		Deadline:     (*hexutil.Big)(u.Deadline),
		Dapp:         u.Dapp,
		Control:      u.Control,
		CallConfig:   (*hexutil.Big)(big.NewInt(int64(u.CallConfig))),
		DappGasLimit: (*hexutil.Big)(u.DappGasLimit),
		SessionKey:   u.SessionKey,
		Data:         hexutil.Bytes(u.Data),
		Signature:    hexutil.Bytes(u.Signature),
	}
}

func (u *UserOperationV15) AbiEncode() ([]byte, error) {
	u.Sanitize()
	return userOpArgsV15.Pack(&u)
}

func (u *UserOperationV15) ValidateSignature(chainId uint64, version *string) error {
	if len(u.Signature) != 65 {
		return errors.New("invalid signature length")
	}

	userOpHash, err := u.Hash(false, chainId, version)
	if err != nil {
		return err
	}

	signer, err := utils.RecoverSigner(userOpHash.Bytes(), u.Signature)
	if err != nil {
		return err
	}

	if signer != u.From {
		return errors.New("invalid signature")
	}

	return nil
}
