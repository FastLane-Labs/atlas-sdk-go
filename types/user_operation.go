package types

import (
	"errors"
	"math/big"
	"math/rand/v2"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	userOpSolType, _ = abi.NewType("tuple", "struct UserOperation", []abi.ArgumentMarshaling{
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256", InternalType: "uint256"},
		{Name: "nonce", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "dapp", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "callConfig", Type: "uint32", InternalType: "uint32"},
		{Name: "sessionKey", Type: "address", InternalType: "address"},
		{Name: "data", Type: "bytes", InternalType: "bytes"},
		{Name: "signature", Type: "bytes", InternalType: "bytes"},
	})

	userOpArgs = abi.Arguments{
		{Type: userOpSolType, Name: "userOperation"},
	}
)

// External representation of a user operation
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
	SessionKey   common.Address `json:"sessionKey"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (u *UserOperationRaw) Decode() *UserOperation {
	return &UserOperation{
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

// Internal representation of a user operation
type UserOperation struct {
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
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

func (u *UserOperation) Sanitize() {
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

func (u *UserOperation) EncodeToRaw() *UserOperationRaw {
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
		SessionKey:   u.SessionKey,
		Data:         hexutil.Bytes(u.Data),
		Signature:    hexutil.Bytes(u.Signature),
	}
}

func (u *UserOperation) Hash(trusted bool, chainId uint64) (common.Hash, error) {
	eip712Domain, err := config.GetEip712Domain(chainId)
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

func (u *UserOperation) AbiEncode() ([]byte, error) {
	u.Sanitize()
	return userOpArgs.Pack(&u)
}

func (u *UserOperation) ValidateSignature(chainId uint64) error {
	if len(u.Signature) != 65 {
		return errors.New("invalid signature length")
	}

	userOpHash, err := u.Hash(false, chainId)
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

func (u *UserOperation) toTypedDataTypes(trusted bool) apitypes.Types {
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
			{Name: "sessionKey", Type: "address"},
			{Name: "data", Type: "bytes"},
		}
	}

	return t
}

func (u *UserOperation) toTypedDataMessage(trusted bool) apitypes.TypedDataMessage {
	u.Sanitize()

	if trusted {
		return apitypes.TypedDataMessage{
			"from":       u.From.Hex(),
			"to":         u.To.Hex(),
			"dapp":       u.Dapp.Hex(),
			"control":    u.Control.Hex(),
			"callConfig": big.NewInt(int64(u.CallConfig)),
			"sessionKey": u.SessionKey.Hex(),
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
		"sessionKey":   u.SessionKey.Hex(),
		"data":         u.Data,
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
	Hints []common.Address `json:"hints,omitempty"`

	Value *hexutil.Big   `json:"value,omitempty"`
	Data  hexutil.Bytes  `json:"data,omitempty"`
	From  common.Address `json:"from,omitempty"`
}

func NewUserOperationPartialRaw(chainId uint64, userOp *UserOperation, hints []common.Address) (*UserOperationPartialRaw, error) {
	userOpHash, err := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), chainId)
	if err != nil {
		return nil, err
	}

	userOpPartial := &UserOperationPartialRaw{
		ChainId:      (*hexutil.Big)(big.NewInt(int64(chainId))),
		UserOpHash:   userOpHash,
		To:           userOp.To,
		Gas:          (*hexutil.Big)(userOp.Gas),
		MaxFeePerGas: (*hexutil.Big)(userOp.MaxFeePerGas),
		Deadline:     (*hexutil.Big)(userOp.Deadline),
		Dapp:         userOp.Dapp,
		Control:      userOp.Control,
	}

	if len(hints) > 0 {
		//randomize hints
		rand.Shuffle(len(hints), func(i, j int) { hints[i], hints[j] = hints[j], hints[i] })
		userOpPartial.Hints = hints
	} else {
		userOpPartial.Data = hexutil.Bytes(userOp.Data)
		userOpPartial.From = userOp.From
		userOpPartial.Value = (*hexutil.Big)(userOp.Value)
	}

	return userOpPartial, nil
}

type UserOperationWithHintsRaw struct {
	ChainId       *hexutil.Big      `json:"chainId"`
	UserOperation *UserOperationRaw `json:"userOperation"`
	Hints         []common.Address  `json:"hints"`
}

func NewUserOperationWithHintsRaw(chainId uint64, userOp *UserOperation, hints []common.Address) *UserOperationWithHintsRaw {
	return &UserOperationWithHintsRaw{
		ChainId:       (*hexutil.Big)(big.NewInt(int64(chainId))),
		UserOperation: userOp.EncodeToRaw(),
		Hints:         hints,
	}
}

func (uop *UserOperationWithHintsRaw) Decode() (uint64, *UserOperation, []common.Address) {
	return uop.ChainId.ToInt().Uint64(), uop.UserOperation.Decode(), uop.Hints
}
