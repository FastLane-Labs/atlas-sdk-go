package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	userOpSolTypeLegacy, _ = abi.NewType("tuple", "struct UserOperation", []abi.ArgumentMarshaling{
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

	userOpArgsLegacy = abi.Arguments{
		{Type: userOpSolTypeLegacy, Name: "userOperation"},
	}
)

func (u *UserOperation) toTypedDataTypesLegacy(trusted bool) apitypes.Types {
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

func (u *UserOperation) toTypedDataMessageLegacy(trusted bool) apitypes.TypedDataMessage {
	u.Sanitize()

	if trusted {
		return apitypes.TypedDataMessage{
			"from":       u.from.Hex(),
			"to":         u.to.Hex(),
			"dapp":       u.dapp.Hex(),
			"control":    u.control.Hex(),
			"callConfig": big.NewInt(int64(u.callConfig)),
			"sessionKey": u.sessionKey.Hex(),
		}
	}

	return apitypes.TypedDataMessage{
		"from":         u.from.Hex(),
		"to":           u.to.Hex(),
		"value":        new(big.Int).Set(u.value),
		"gas":          new(big.Int).Set(u.gas),
		"maxFeePerGas": new(big.Int).Set(u.maxFeePerGas),
		"nonce":        new(big.Int).Set(u.nonce),
		"deadline":     new(big.Int).Set(u.deadline),
		"dapp":         u.dapp.Hex(),
		"control":      u.control.Hex(),
		"callConfig":   big.NewInt(int64(u.callConfig)),
		"sessionKey":   u.sessionKey.Hex(),
		"data":         u.data,
	}
}
