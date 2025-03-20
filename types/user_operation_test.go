package types

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateUserOperation() *UserOperation {
	return &UserOperation{
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

func TestNewUserOperationPartialRawWithHints(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	hints := map[string]interface{}{
		"single": []string{"a"},                     // shouldn't be shuffled
		"empty":  []int{},                           // shouldn't be shuffled
		"multi":  []string{"a", "b", "c", "d", "e"}, // should be shuffled
		"nums":   []int{1, 2, 3, 4, 5},              // should be shuffled
		"addresses": []common.Address{
			common.HexToAddress("0x1"),
			common.HexToAddress("0x2"),
			common.HexToAddress("0x3"),
			common.HexToAddress("0x4"),
			common.HexToAddress("0x5"),
		}, // should be shuffled
		"bigints": []*big.Int{
			big.NewInt(1),
			big.NewInt(2),
			big.NewInt(3),
			big.NewInt(4),
			big.NewInt(5),
		}, // should be shuffled
		"static": "not-a-slice", // shouldn't be touched
	}

	// Create a copy of original slices for comparison
	originalMulti := make([]string, len(hints["multi"].([]string)))
	copy(originalMulti, hints["multi"].([]string))
	originalNums := make([]int, len(hints["nums"].([]int)))
	copy(originalNums, hints["nums"].([]int))
	originalAddresses := make([]common.Address, len(hints["addresses"].([]common.Address)))
	copy(originalAddresses, hints["addresses"].([]common.Address))
	originalBigints := make([]*big.Int, len(hints["bigints"].([]*big.Int)))
	copy(originalBigints, hints["bigints"].([]*big.Int))

	// Run multiple times to ensure we get different orders
	shuffledAtLeastOnce := false
	for i := 0; i < 10; i++ {
		result, err := NewUserOperationPartialRaw(0, nil, userOp, hints)
		if err != nil {
			t.Fatalf("NewUserOperationPartialRaw() error = %v", err)
		}

		// Verify single element slice remains unchanged
		if !reflect.DeepEqual(result.Hints["single"], []string{"a"}) {
			t.Error("single element slice was modified")
		}

		// Verify empty slice remains unchanged
		if !reflect.DeepEqual(result.Hints["empty"], []int{}) {
			t.Error("empty slice was modified")
		}

		// Verify non-slice value remains unchanged
		if result.Hints["static"] != "not-a-slice" {
			t.Error("non-slice value was modified")
		}

		// Check if multi-element slices were shuffled
		if !reflect.DeepEqual(result.Hints["multi"], originalMulti) {
			shuffledAtLeastOnce = true
		}
		if !reflect.DeepEqual(result.Hints["nums"], originalNums) {
			shuffledAtLeastOnce = true
		}
		if !reflect.DeepEqual(result.Hints["addresses"], originalAddresses) {
			shuffledAtLeastOnce = true
		}
		if !reflect.DeepEqual(result.Hints["bigints"], originalBigints) {
			shuffledAtLeastOnce = true
		}
	}

	if !shuffledAtLeastOnce {
		t.Error("multi-element slices were never shuffled after 10 attempts")
	}
}
