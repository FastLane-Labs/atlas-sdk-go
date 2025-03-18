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

type UserOperationV15 struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Dapp         common.Address
	Control      common.Address
	CallConfig   uint32
	DappGasLimit uint32
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

func (u *UserOperationV15) GetFrom() common.Address {
	return u.From
}

func (u *UserOperationV15) GetTo() common.Address {
	return u.To
}

func (u *UserOperationV15) GetValue() *big.Int {
	return u.Value
}

func (u *UserOperationV15) GetGas() *big.Int {
	return u.Gas
}

func (u *UserOperationV15) GetMaxFeePerGas() *big.Int {
	return u.MaxFeePerGas
}

func (u *UserOperationV15) GetNonce() *big.Int {
	return u.Nonce
}

func (u *UserOperationV15) GetDeadline() *big.Int {
	return u.Deadline
}

func (u *UserOperationV15) GetDapp() common.Address {
	return u.Dapp
}

func (u *UserOperationV15) GetControl() common.Address {
	return u.Control
}

func (u *UserOperationV15) GetCallConfig() uint32 {
	return u.CallConfig
}

func (u *UserOperationV15) GetSessionKey() common.Address {
	return u.SessionKey
}

func (u *UserOperationV15) GetData() []byte {
	return u.Data
}

func (u *UserOperationV15) GetSignature() []byte {
	return u.Signature
}

func (u *UserOperationV15) SetNonce(nonce *big.Int) {
	u.Nonce = nonce
}

func (u *UserOperationV15) GetDappGasLimit() uint32 {
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

func (u *UserOperationV15) Sanitize() {
	if u.Value == nil {
		u.Value = big.NewInt(0)
	}

	if u.Gas == nil {
		u.Gas = big.NewInt(0)
	}

	if u.MaxFeePerGas == nil {
		u.MaxFeePerGas = big.NewInt(0)
	}

	if u.Nonce == nil {
		u.Nonce = big.NewInt(0)
	}

	if u.Deadline == nil {
		u.Deadline = big.NewInt(0)
	}
}

func (u *UserOperationV15) toTypedDataMessage(trusted bool) apitypes.TypedDataMessage {
	u.Sanitize()

	if trusted {
		return apitypes.TypedDataMessage{
			"from":         u.From.Hex(),
			"to":           u.To.Hex(),
			"dapp":         u.Dapp.Hex(),
			"control":      u.Control.Hex(),
			"callConfig":   big.NewInt(int64(u.CallConfig)),
			"dappGasLimit": big.NewInt(int64(u.DappGasLimit)),
			"sessionKey":   u.SessionKey.Hex(),
		}
	}

	return apitypes.TypedDataMessage{
		"from":         u.From.Hex(),
		"to":           u.To.Hex(),
		"value":        new(big.Int).Set(u.Value),
		"gas":          new(big.Int).Set(u.Gas),
		"maxFeePerGas": new(big.Int).Set(u.MaxFeePerGas),
		"nonce":        new(big.Int).Set(u.Nonce),
		"deadline":     new(big.Int).Set(u.Deadline),
		"dapp":         u.Dapp.Hex(),
		"control":      u.Control.Hex(),
		"callConfig":   big.NewInt(int64(u.CallConfig)),
		"dappGasLimit": big.NewInt(int64(u.DappGasLimit)),
		"sessionKey":   u.SessionKey.Hex(),
		"data":         u.Data,
	}
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
		DappGasLimit: (*hexutil.Big)(big.NewInt(int64(u.DappGasLimit))),
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

func (u *UserOperationV15) Hash(trusted bool, chainId uint64, version *string) (common.Hash, error) {
	eip712Domain, err := config.GetEip712Domain(chainId, version)
	if err != nil {
		return common.Hash{}, err
	}

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       u.toTypedDataTypes(trusted),
		PrimaryType: "UserOperation",
		Domain:      *eip712Domain,
		Message:     u.toTypedDataMessage(trusted),
	})

	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(hash), nil
}
