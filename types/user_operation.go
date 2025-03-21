package types

import (
	"math/big"
	"math/rand/v2"
	"reflect"

	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	_ UserOperation = &UserOperationLegacy{}
	_ UserOperation = &UserOperationV15{}
)

type UserOperation interface {
	GetTo() common.Address
	GetFrom() common.Address
	GetValue() *big.Int
	GetGas() *big.Int
	GetMaxFeePerGas() *big.Int
	GetNonce() *big.Int
	GetDeadline() *big.Int
	GetDapp() common.Address
	GetControl() common.Address
	GetCallConfig() uint32
	GetDappGasLimit() uint32
	GetSessionKey() common.Address
	GetData() []byte
	GetSignature() []byte

	SetNonce(nonce *big.Int)

	Sanitize()
	EncodeToRaw() *UserOperationRaw
	Hash(trusted bool, chainId uint64, version *string) (common.Hash, error)
	AbiEncode() ([]byte, error)
	ValidateSignature(chainId uint64, version *string) error
	toTypedDataTypes(trusted bool) apitypes.Types
	toTypedDataMessage(trusted bool) apitypes.TypedDataMessage
}

type UserOperationRaw struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Nonce        *hexutil.Big   `json:"nonce"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	CallConfig   *hexutil.Big   `json:"callConfig"`
	DappGasLimit *hexutil.Big   `json:"dappGasLimit,omitempty"`
	SessionKey   common.Address `json:"sessionKey"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (u *UserOperationRaw) Sanitize() {
	if u.Value == nil {
		u.Value = (*hexutil.Big)(big.NewInt(0))
	}

	if u.Gas == nil {
		u.Gas = (*hexutil.Big)(big.NewInt(0))
	}

	if u.MaxFeePerGas == nil {
		u.MaxFeePerGas = (*hexutil.Big)(big.NewInt(0))
	}

	if u.Nonce == nil {
		u.Nonce = (*hexutil.Big)(big.NewInt(0))
	}

	if u.Deadline == nil {
		u.Deadline = (*hexutil.Big)(big.NewInt(0))
	}

	if u.CallConfig == nil {
		u.CallConfig = (*hexutil.Big)(big.NewInt(0))
	}
}

func (u *UserOperationRaw) Decode() UserOperation {
	u.Sanitize()

	if u.DappGasLimit != nil {
		return &UserOperationV15{
			From:         u.From,
			To:           u.To,
			Value:        u.Value.ToInt(),
			Gas:          u.Gas.ToInt(),
			MaxFeePerGas: u.MaxFeePerGas.ToInt(),
			Nonce:        u.Nonce.ToInt(),
			Deadline:     u.Deadline.ToInt(),
			Dapp:         u.Dapp,
			Control:      u.Control,
			CallConfig:   uint32(u.CallConfig.ToInt().Uint64()),
			DappGasLimit: uint32(u.DappGasLimit.ToInt().Uint64()),
			SessionKey:   u.SessionKey,
			Data:         u.Data,
			Signature:    u.Signature,
		}
	} else {
		return &UserOperationLegacy{
			From:         u.From,
			To:           u.To,
			Value:        u.Value.ToInt(),
			Gas:          u.Gas.ToInt(),
			MaxFeePerGas: u.MaxFeePerGas.ToInt(),
			Nonce:        u.Nonce.ToInt(),
			Deadline:     u.Deadline.ToInt(),
			Dapp:         u.Dapp,
			Control:      u.Control,
			CallConfig:   uint32(u.CallConfig.ToInt().Uint64()),
			SessionKey:   u.SessionKey,
			Data:         u.Data,
			Signature:    u.Signature,
		}
	}
}

type UserOperationPartialRaw struct {
	ChainId *hexutil.Big `json:"chainId"`

	UserOpHash   common.Hash    `json:"userOpHash"`
	To           common.Address `json:"to"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`

	//Exactly one of 1. Hints  2. (Value, Data, From) must be set
	Hints map[string]interface{} `json:"hints,omitempty"`

	Value *hexutil.Big   `json:"value,omitempty"`
	Data  hexutil.Bytes  `json:"data,omitempty"`
	From  common.Address `json:"from,omitempty"`
}

func NewUserOperationPartialRaw(chainId uint64, version *string, userOp UserOperation, hints map[string]interface{}) (*UserOperationPartialRaw, error) {
	userOpHash, err := userOp.Hash(utils.FlagTrustedOpHash(userOp.GetCallConfig(), version), chainId, version)
	if err != nil {
		return nil, err
	}

	userOpPartial := &UserOperationPartialRaw{
		ChainId:      (*hexutil.Big)(big.NewInt(int64(chainId))),
		UserOpHash:   userOpHash,
		To:           userOp.GetTo(),
		Gas:          (*hexutil.Big)(userOp.GetGas()),
		MaxFeePerGas: (*hexutil.Big)(userOp.GetMaxFeePerGas()),
		Deadline:     (*hexutil.Big)(userOp.GetDeadline()),
		Dapp:         userOp.GetDapp(),
		Control:      userOp.GetControl(),
	}

	if len(hints) > 0 {
		// Randomize slices/arrays
		for key, hint := range hints {
			v := reflect.ValueOf(hint)
			if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
				randomizeSlice(v)
				hints[key] = v.Interface()
			}
		}
		userOpPartial.Hints = hints
	} else {
		userOpPartial.Data = hexutil.Bytes(userOp.GetData())
		userOpPartial.From = userOp.GetFrom()
		userOpPartial.Value = (*hexutil.Big)(userOp.GetValue())
	}

	return userOpPartial, nil
}

type UserOperationWithHintsRaw struct {
	ChainId       *hexutil.Big           `json:"chainId"`
	UserOperation *UserOperationRaw      `json:"userOperation"`
	Hints         map[string]interface{} `json:"hints"`
}

func NewUserOperationWithHintsRaw(chainId uint64, userOp UserOperation, hints map[string]interface{}) *UserOperationWithHintsRaw {
	return &UserOperationWithHintsRaw{
		ChainId:       (*hexutil.Big)(big.NewInt(int64(chainId))),
		UserOperation: userOp.EncodeToRaw(),
		Hints:         hints,
	}
}

func (uop *UserOperationWithHintsRaw) Decode() (uint64, UserOperation, map[string]interface{}) {
	return uop.ChainId.ToInt().Uint64(), uop.UserOperation.Decode(), uop.Hints
}

func randomizeSlice(v reflect.Value) {
	if v.Len() > 1 {
		rand.Shuffle(v.Len(), func(i, j int) {
			vi, vj := v.Index(i), v.Index(j)
			tmp := reflect.New(vi.Type()).Elem()
			tmp.Set(vi)
			vi.Set(vj)
			vj.Set(tmp)
		})
	}
}
