package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateUserOperation() *UserOperationLegacy {
	return &UserOperationLegacy{
		From:         common.HexToAddress("0x1"),
		To:           common.HexToAddress("0x2"),
		Deadline:     big.NewInt(100),
		Gas:          big.NewInt(200),
		Nonce:        big.NewInt(300),
		MaxFeePerGas: big.NewInt(400),
		Value:        big.NewInt(500),
		Dapp:         common.HexToAddress("0x3"),
		Control:      common.HexToAddress("0x4"),
		CallConfig:   600,
		SessionKey:   common.HexToAddress("0x5"),
		Data:         []byte("data"),
		Signature:    []byte("signature"),
	}
}

// func generateUserOperationV15() *UserOperationV15 {
// 	return &UserOperationV15{
// 		From:         common.HexToAddress("0x1"),
// 		To:           common.HexToAddress("0x2"),
// 		Deadline:     big.NewInt(100),
// 		Gas:          big.NewInt(200),
// 		Nonce:        big.NewInt(300),
// 		MaxFeePerGas: big.NewInt(400),
// 		Value:        big.NewInt(500),
// 		Dapp:         common.HexToAddress("0x3"),
// 		Control:      common.HexToAddress("0x4"),
// 		CallConfig:   600,
// 		DappGasLimit: 600_000,
// 		SessionKey:   common.HexToAddress("0x5"),
// 		Data:         []byte("data"),
// 		Signature:    []byte("signature"),
// 	}
// }

func TestUserOperationHashDefault(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.HexToHash("0x021a7f3f62347f1f3d1163aa8eb9fc965e87556aede03c7182ec05bc60311b64")

	result, err := userOp.Hash(false, 0, nil)
	if err != nil {
		t.Errorf("UserOperation.Hash() error = %v", err)
	}

	if result != want {
		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
	}
}

func TestUserOperationHashTrusted(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.HexToHash("0x96aa1212cae2645ba1b8bf8014abccdfe9a60c16f86e21f82753d4cecc0b6089")

	result, err := userOp.Hash(true, 0, nil)
	if err != nil {
		t.Errorf("UserOperation.Hash() error = %v", err)
	}

	if result != want {
		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
	}
}

func TestUserOperationAbiEncode(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.FromHex("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000001f400000000000000000000000000000000000000000000000000000000000000c80000000000000000000000000000000000000000000000000000000000000190000000000000000000000000000000000000000000000000000000000000012c0000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000258000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000004646174610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000097369676e61747572650000000000000000000000000000000000000000000000")

	result, err := userOp.AbiEncode()
	if err != nil {
		t.Errorf("UserOperation.AbiEncode() error = %v", err)
	}

	if string(result) != string(want) {
		t.Errorf("UserOperation.AbiEncode() = %v, want %v", result, want)
	}
}

func TestUserOperationValidateSignature(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	userOp.From = common.HexToAddress("0xB764B6545d283C0E547952763F8a843394295da1")
	userOp.Signature = common.FromHex("0x63e05429d1f5253ceebddf5f709c33d211592798cc4f89af302ade417e1de0173dd2c50d1bccc996e68839491a3539c3400c8b0721c29c236a3027f1dc274e151b")

	if err := userOp.ValidateSignature(0, nil); err != nil {
		t.Errorf("DAppOperation.checkSignature() error = %v", err)
	}
}

// func TestUserOperationV15HashDefault(t *testing.T) {
// 	t.Parallel()

// 	userOp := generateUserOperationV15()
// 	want := common.HexToHash("0xbe49dd84d9a2d5713920fa3884a0c52ee40eab8e33b7547f8d4e206ad77d7ce2")

// 	result, err := userOp.Hash(false, 0, nil)
// 	if err != nil {
// 		t.Errorf("UserOperation.Hash() error = %v", err)
// 	}

// 	if result != want {
// 		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
// 	}
// }

// func TestUserOperationV15HashTrusted(t *testing.T) {
// 	t.Parallel()

// 	userOp := generateUserOperationV15()
// 	want := common.HexToHash("0x006f92454adbab861e749a27a6e84b65b33f7293ff3ced2b1e48c0eb15647e8e")

// 	result, err := userOp.Hash(true, 0, nil)
// 	if err != nil {
// 		t.Errorf("UserOperation.Hash() error = %v", err)
// 	}

// 	if result != want {
// 		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
// 	}
// }

// func TestUserOperationV15UserOpHash(t *testing.T) {
// 	t.Parallel()

// 	userOp := UserOperationV15{
// 		From:         common.HexToAddress("0xfc8b8974fc3adb8281a6c4c38d7cc895769a8568"),
// 		To:           common.HexToAddress("0x5f4f2a8961ef043817100e512286ff5096ae0042"),
// 		Value:        big.NewInt(0),
// 		Gas:          big.NewInt(20000),
// 		MaxFeePerGas: big.NewInt(5000000),
// 		Nonce:        big.NewInt(1),
// 		Deadline:     big.NewInt(27771705),
// 		Dapp:         common.HexToAddress("0x0e3009d01e85ac49d164e453ec81283eaaf46fb5"),
// 		Control:      common.HexToAddress("0x0e3009d01e85ac49d164e453ec81283eaaf46fb5"),
// 		CallConfig:   532548,
// 		DappGasLimit: 2000000,
// 		SessionKey:   common.HexToAddress("0x30d995248f48f101d18b21ece539fb862d7b4487"),
// 		Data:         []byte("0x1ad6fbc3"),
// 		Signature:    common.FromHex("0x85a2bc106a41bfb5055d62c8afe7f649af5fcea06c40e4964e0f08eb85d805eb1ebd0be19740aa7841c5084e03ab03554fa9aa46aa315efff266d72b805178191c"),
// 	}

// 	want := common.HexToHash("0x8836c01721b1efdf5cb088797df2d40a3e4437b0cf7c0cea1edae8ecd0d7b33f")

// 	result, err := userOp.Hash(true, 0, nil)
// 	if err != nil {
// 		t.Errorf("UserOperation.Hash() error = %v", err)
// 	}

// 	if result != want {
// 		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
// 	}
// }
