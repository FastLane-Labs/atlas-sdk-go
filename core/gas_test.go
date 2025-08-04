package core

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gethEthClient "github.com/ethereum/go-ethereum/ethclient"
)

const (
	testRpcUrl = "https://rpc.berachain.com"
)

var (
	gasTestConfigVersion = config.AtlasV_1_6_2
	testDAppControlAddr  = common.HexToAddress("0xfBc81A39459E0D82EC31B4e585f7A318AFAdB49B")
)

// Test configurations for different chains
type TestChainConfig struct {
	name            string
	rpcUrl          string
	dAppControlAddr common.Address
	atlasVersion    string
}

var testChainConfigs = []TestChainConfig{
	{
		name:            "berachain",
		rpcUrl:          "https://rpc.berachain.com",
		dAppControlAddr: common.HexToAddress("0xfBc81A39459E0D82EC31B4e585f7A318AFAdB49B"),
		atlasVersion:    config.AtlasV_1_6_2,
	},
	{
		name:            "polygon",
		rpcUrl:          "https://polygon-rpc.com",
		dAppControlAddr: common.HexToAddress("0x0F69afd3595e2BF74A8F8d53E3322F22040825F3"),
		atlasVersion:    config.AtlasV_1_6_2,
	},
	{
		name:            "monad_testnet",
		rpcUrl:          "https://testnet-rpc.monad.xyz",
		dAppControlAddr: common.HexToAddress("0x26E07BA2c7074A06AEAb50403Dd535F627A2c5D6"),
		atlasVersion:    config.AtlasV_1_6_3_MONAD,
	},
}

// SwapTokenInfo represents the Monad-specific data structure
type SwapTokenInfo struct {
	InputToken            common.Address
	InputAmount           *big.Int
	OutputToken           common.Address
	OutputMin             *big.Int
	BidTokenIsOutputToken bool
	Target                common.Address
	SwapData              []byte
}

// generateMonadSwapData creates properly formatted data for Monad dApp control
func generateMonadSwapData() []byte {
	// Define the ABI for SwapTokenInfo struct
	swapTokenInfoABI, err := abi.NewType("tuple", "SwapTokenInfo", []abi.ArgumentMarshaling{
		{Name: "inputToken", Type: "address"},
		{Name: "inputAmount", Type: "uint256"},
		{Name: "outputToken", Type: "address"},
		{Name: "outputMin", Type: "uint256"},
		{Name: "bidTokenIsOutputToken", Type: "bool"},
		{Name: "target", Type: "address"},
		{Name: "swapData", Type: "bytes"},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create SwapTokenInfo ABI type: %v", err))
	}

	// Create the full ABI for the function parameters
	arguments := abi.Arguments{
		{Type: swapTokenInfoABI, Name: "swapInfo"},
		{Type: abi.Type{T: abi.AddressTy}, Name: "address2"},
		{Type: abi.Type{T: abi.UintTy, Size: 256}, Name: "uint256"},
	}

	// Create sample SwapTokenInfo data
	swapInfo := SwapTokenInfo{
		InputToken:            common.HexToAddress("0xA0b86a33E6441a8f7B0C3D1B4d5C6D2A9f7E8D3C"),
		InputAmount:           big.NewInt(1000000000000000000), // 1 ETH
		OutputToken:           common.HexToAddress("0xB1c87a44F7552b9C4E6D3F5A8B2E1D6F9C8A5B2E"),
		OutputMin:             big.NewInt(500000000000000000), // 0.5 ETH minimum
		BidTokenIsOutputToken: true,
		Target:                common.HexToAddress("0xC2d98b55F8663c0F7E4F6B7D9C3F2A1E8D5C6B4F"),
		SwapData:              []byte("mock_swap_data"),
	}

	// Encode the parameters
	data, err := arguments.Pack(
		swapInfo,
		common.HexToAddress("0xD3e09c66A9774d1F8F5A7C8E0D4A3B2F9E6D7C5A"),
		big.NewInt(12345),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to encode SwapTokenInfo: %v", err))
	}

	// Add a dummy 4-byte function selector at the beginning
	functionSelector := []byte{0x12, 0x34, 0x56, 0x78}
	return append(functionSelector, data...)
}

// Test excluded chain behavior
func TestEstimateMetacallGasLimitExcludedChain(t *testing.T) {
	t.Skip("Skipped in CI because RPC calls are unreliable on public RPCs, run locally with good RPCs instead")

	// Test Unichain which is in the excluded list (chain ID 130)
	ethClient, err := gethEthClient.Dial("https://unichain.drpc.org")
	if err != nil {
		t.Skipf("Skipping excluded chain test: failed to create eth client: %v", err)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		t.Skipf("Skipping excluded chain test: failed to get chain id: %v", err)
	}

	t.Logf("Unichain chain ID: %d", chainId.Uint64())

	sdk, err := NewAtlasSdk([]interface{}{ethClient}, nil)
	if err != nil {
		t.Skipf("Skipping excluded chain test: failed to create atlas sdk: %v", err)
	}

	// Unichain dApp control address and Atlas version
	dAppControlAddr := common.HexToAddress("0x892F8f6779ca6927c1A6Cc74319e03d2abEf18D5")
	version := config.AtlasV_1_6_2

	// Create a user operation with proper parameters
	userOp := generateUserOperationWithParams(chainId.Uint64(), version, 100000, 100, 600)

	// Get the dApp config from the dApp control contract
	t.Log("Getting dApp config...")
	dAppConfig, err := sdk.GetDAppConfig(chainId.Uint64(), &version, userOp, dAppControlAddr)
	if err != nil {
		t.Fatalf("Failed to get dApp config: %v", err)
	}

	// Set the proper values from dApp config
	userOp.SetControl(dAppControlAddr)
	userOp.SetDappGasLimit(dAppConfig.DappGasLimit)
	userOp.SetSolverGasLimit(dAppConfig.SolverGasLimit)
	userOp.SetCallConfig(dAppConfig.CallConfig)

	t.Logf("dApp config: DappGasLimit=%d, SolverGasLimit=%d, CallConfig=%d",
		dAppConfig.DappGasLimit, dAppConfig.SolverGasLimit, dAppConfig.CallConfig)

	// Create empty solver operations
	solverOps := types.SolverOperations{}

	// Test that EstimateMetacallGasLimitLocal returns the expected error
	t.Log("Testing EstimateMetacallGasLimitLocal on excluded chain...")
	localLimit, err := sdk.EstimateMetacallGasLimitLocal(
		chainId.Uint64(),
		&version,
		userOp,
		solverOps,
	)

	// We expect this to fail with a specific error message
	if err == nil {
		t.Fatalf("Expected EstimateMetacallGasLimitLocal to fail for excluded chain ID %d, but it succeeded with result: %d",
			chainId.Uint64(), localLimit)
	}

	expectedErrorSubstring := fmt.Sprintf("metacall gas limit estimation offchain is not supported for chainId %d", chainId.Uint64())
	if !strings.Contains(err.Error(), expectedErrorSubstring) {
		t.Fatalf("Expected error to contain '%s', but got: %v", expectedErrorSubstring, err)
	}

	t.Logf("✅ Correctly rejected excluded chain ID %d with error: %v", chainId.Uint64(), err)

	// Also verify that the regular EstimateMetacallGasLimit (which tries local first, then fallback)
	// should fallback to the on-chain method and succeed with proper dApp config values
	t.Log("Testing EstimateMetacallGasLimit fallback behavior...")
	onChainLimit, err := sdk.EstimateMetacallGasLimit(
		chainId.Uint64(),
		&version,
		userOp,
		solverOps,
		big.NewInt(1000000000),
	)

	if err != nil {
		t.Fatalf("Expected EstimateMetacallGasLimit to succeed via fallback to on-chain with proper dApp config, but got error: %v", err)
	}

	t.Logf("✅ EstimateMetacallGasLimit successfully fell back to on-chain estimation: %d", onChainLimit)

	// Verify the result is reasonable (should be greater than 0)
	if onChainLimit <= 0 {
		t.Fatalf("Expected positive gas limit, got: %d", onChainLimit)
	}
}

func TestEstimateMetacallGasLimitLocal(t *testing.T) {
	t.Skip("Skipped in CI because RPC calls are unreliable on public RPCs, run locally with good RPCs instead")

	ethClient, err := gethEthClient.Dial(testRpcUrl)
	if err != nil {
		t.Fatalf("failed to create eth client: %v", err)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	sdk, err := NewAtlasSdk([]interface{}{ethClient}, nil)
	if err != nil {
		t.Fatalf("failed to create atlas sdk: %v", err)
	}

	userOp := generateUserOperation(chainId.Uint64(), gasTestConfigVersion)

	dAppConfig, err := sdk.GetDAppConfig(chainId.Uint64(), &gasTestConfigVersion, userOp, testDAppControlAddr)
	if err != nil {
		t.Fatalf("failed to get dapp config: %v", err)
	}

	userOp.SetControl(testDAppControlAddr)
	userOp.SetDappGasLimit(dAppConfig.DappGasLimit)
	userOp.SetSolverGasLimit(dAppConfig.SolverGasLimit)
	userOp.SetCallConfig(dAppConfig.CallConfig)

	solverOps := types.SolverOperations{
		generateSolverOperation(),
		generateSolverOperation(),
		generateSolverOperation(),
	}

	onChainLimit, err := sdk.estimateMetacallGasLimit(
		chainId.Uint64(),
		&gasTestConfigVersion,
		userOp,
		solverOps,
		big.NewInt(1000000000),
	)
	if err != nil {
		t.Fatalf("failed to estimate metacall gas limit: %v", err)
	}

	localLimit, err := sdk.EstimateMetacallGasLimitLocal(
		chainId.Uint64(),
		&gasTestConfigVersion,
		userOp,
		solverOps,
	)
	if err != nil {
		t.Fatalf("failed to estimate metacall gas limit: %v", err)
	}

	t.Logf("onChainLimit: %d", onChainLimit)
	t.Logf("localLimit: %d", localLimit)

	if onChainLimit != localLimit {
		t.Fatalf("onChainLimit and localLimit are not equal: %d != %d", onChainLimit, localLimit)
	}
}

// Helper functions for generating varied test operations

func generateUserOperationWithParams(chainId uint64, atlasVersion string, gasLimit int64, dataSize int, callConfig uint32) *types.UserOperation {
	atlasAddr, err := config.GetAtlasAddress(chainId, &atlasVersion)
	if err != nil {
		panic(err)
	}

	// Generate data of specified size
	var data []byte
	if strings.Contains(atlasVersion, "monad") {
		// For Monad chains, use properly formatted SwapTokenInfo data
		data = generateMonadSwapData()
	} else {
		// For other chains, use simple byte data
		data = make([]byte, dataSize)
		for i := range data {
			data[i] = byte(i % 256)
		}
	}

	userOp, err := types.NewUserOperation(
		chainId,
		types.UserOperationsParams{
			From:         common.HexToAddress("0x1111111111111111111111111111111111111111"),
			To:           atlasAddr,
			Deadline:     big.NewInt(100),
			Gas:          big.NewInt(gasLimit),
			Nonce:        big.NewInt(300),
			MaxFeePerGas: big.NewInt(400),
			Value:        big.NewInt(500),
			Dapp:         common.HexToAddress("0x3333333333333333333333333333333333333333"),
			Control:      common.HexToAddress("0x4444444444444444444444444444444444444444"),
			CallConfig:   callConfig,
			SessionKey:   common.HexToAddress("0x5555555555555555555555555555555555555555"),
			Data:         data,
			Signature:    []byte("signature"),
		},
	)
	if err != nil {
		panic(err)
	}

	return userOp
}

func generateSolverOperationWithParams(gasLimit int64, dataSize int, bidAmount int64) *types.SolverOperation {
	// Generate data of specified size
	data := make([]byte, dataSize)
	for i := range data {
		data[i] = byte((i + 42) % 256) // Different pattern than user op data
	}

	return &types.SolverOperation{
		From:         common.HexToAddress("0x1111111111111111111111111111111111111111"),
		To:           common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Value:        big.NewInt(100),
		Gas:          big.NewInt(gasLimit),
		MaxFeePerGas: big.NewInt(300),
		Deadline:     big.NewInt(400),
		Solver:       common.HexToAddress("0x3333333333333333333333333333333333333333"),
		Control:      common.HexToAddress("0x4444444444444444444444444444444444444444"),
		UserOpHash:   common.HexToHash("0x9999999999999999999999999999999999999999999999999999999999999999"),
		BidToken:     common.HexToAddress("0x5555555555555555555555555555555555555555"),
		BidAmount:    big.NewInt(bidAmount),
		Data:         data,
		Signature:    []byte("signature"),
	}
}

func generateSolverOperations(count int, gasLimit int64, dataSize int) types.SolverOperations {
	ops := make(types.SolverOperations, count)
	for i := 0; i < count; i++ {
		// Vary bid amounts to make operations different
		bidAmount := int64(500 + i*100)
		ops[i] = generateSolverOperationWithParams(gasLimit, dataSize, bidAmount)
	}
	return ops
}

// Comprehensive test cases across multiple chains
func TestEstimateMetacallGasLimitComprehensive(t *testing.T) {
	t.Skip("Skipped in CI because RPC calls are unreliable on public RPCs, run locally with good RPCs instead")

	testCases := []struct {
		name           string
		userGasLimit   int64
		userDataSize   int
		solverCount    int
		solverGasLimit int64
		solverDataSize int
	}{
		// Basic variations
		{"minimal_user_no_solvers", 100, 0, 0, 0, 0},
		{"minimal_user_one_solver", 100, 0, 1, 100, 0},
		{"small_data_one_solver", 200, 32, 1, 200, 32},
		{"medium_data_one_solver", 500, 1024, 1, 500, 1024},

		// Different solver counts
		{"two_solvers", 200, 4, 2, 200, 4},
		{"three_solvers", 200, 4, 3, 200, 4},
		{"five_solvers", 200, 4, 5, 200, 4},
		{"ten_solvers", 200, 4, 10, 200, 4},

		// Different gas limits
		{"low_gas_user", 50, 4, 3, 200, 4},
		{"high_gas_user", 1000, 4, 3, 200, 4},
		{"low_gas_solvers", 200, 4, 3, 50, 4},
		{"high_gas_solvers", 200, 4, 3, 1000, 4},
		{"mixed_gas_limits", 750, 4, 3, 300, 4},

		// Different data sizes
		{"empty_data", 200, 0, 3, 200, 0},
		{"tiny_data", 200, 4, 3, 200, 4},
		{"small_data", 200, 64, 3, 200, 64},
		{"medium_data", 200, 512, 3, 200, 512},
		{"large_data", 200, 2048, 3, 200, 2048},
		{"huge_data", 200, 8192, 3, 200, 8192},

		// Asymmetric configurations
		{"large_user_small_solvers", 1000, 1024, 5, 100, 32},
		{"small_user_large_solvers", 100, 32, 5, 1000, 1024},
		{"many_small_solvers", 200, 100, 15, 50, 10},
		{"few_large_solvers", 200, 100, 2, 800, 2048},

		// Edge cases
		{"zero_gas_user", 0, 0, 1, 200, 4},
		{"single_byte_data", 200, 1, 3, 200, 1},
		{"max_reasonable_solvers", 200, 64, 20, 200, 64},
	}

	// Test across different chain configurations
	for _, chainConfig := range testChainConfigs {
		t.Run(chainConfig.name, func(t *testing.T) {
			ethClient, err := gethEthClient.Dial(chainConfig.rpcUrl)
			if err != nil {
				t.Skipf("Skipping %s tests: failed to create eth client: %v", chainConfig.name, err)
			}

			chainId, err := ethClient.ChainID(context.Background())
			if err != nil {
				t.Skipf("Skipping %s tests: failed to get chain id: %v", chainConfig.name, err)
			}

			sdk, err := NewAtlasSdk([]interface{}{ethClient}, nil)
			if err != nil {
				t.Skipf("Skipping %s tests: failed to create atlas sdk: %v", chainConfig.name, err)
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					userOp := generateUserOperationWithParams(chainId.Uint64(), chainConfig.atlasVersion, tc.userGasLimit, tc.userDataSize, 600)

					dAppConfig, err := sdk.GetDAppConfig(chainId.Uint64(), &chainConfig.atlasVersion, userOp, chainConfig.dAppControlAddr)
					if err != nil {
						t.Fatalf("failed to get dapp config: %v", err)
					}

					// Use the actual dAppConfig values
					userOp.SetControl(chainConfig.dAppControlAddr)
					userOp.SetDappGasLimit(dAppConfig.DappGasLimit)
					userOp.SetSolverGasLimit(dAppConfig.SolverGasLimit)
					userOp.SetCallConfig(dAppConfig.CallConfig)

					solverOps := generateSolverOperations(tc.solverCount, tc.solverGasLimit, tc.solverDataSize)

					onChainLimit, err := sdk.estimateMetacallGasLimit(
						chainId.Uint64(),
						&chainConfig.atlasVersion,
						userOp,
						solverOps,
						big.NewInt(1000000000),
					)
					if err != nil {
						t.Fatalf("failed to estimate metacall gas limit (on-chain): %v", err)
					}

					localLimit, err := sdk.EstimateMetacallGasLimitLocal(
						chainId.Uint64(),
						&chainConfig.atlasVersion,
						userOp,
						solverOps,
					)
					if err != nil {
						t.Fatalf("failed to estimate metacall gas limit (local): %v", err)
					}

					t.Logf("Chain: %s, Test case: %s", chainConfig.name, tc.name)
					t.Logf("  DApp config: dappGas=%d, solverGas=%d, callConfig=%d",
						dAppConfig.DappGasLimit, dAppConfig.SolverGasLimit, dAppConfig.CallConfig)
					t.Logf("  User: gas=%d, dataSize=%d", tc.userGasLimit, tc.userDataSize)
					t.Logf("  Solvers: count=%d, gas=%d, dataSize=%d", tc.solverCount, tc.solverGasLimit, tc.solverDataSize)
					t.Logf("  OnChain: %d, Local: %d, Diff: %d", onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))

					if onChainLimit != localLimit {
						t.Errorf("Gas limit mismatch for %s/%s: onChain=%d, local=%d, diff=%d",
							chainConfig.name, tc.name, onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))
					}
				})
			}
		})
	}
}

// Test with mixed solver configurations
func TestEstimateMetacallGasLimitRandomVariations(t *testing.T) {
	t.Skip("Skipped in CI because RPC calls are unreliable on public RPCs, run locally with good RPCs instead")

	// Test variations with mixed solver configurations
	mixedTestCases := []struct {
		name          string
		userGasLimit  int64
		userDataSize  int
		solverConfigs []struct {
			gas      int64
			dataSize int
		}
	}{
		{
			name:         "mixed_solver_sizes",
			userGasLimit: 300,
			userDataSize: 128,
			solverConfigs: []struct {
				gas      int64
				dataSize int
			}{
				{100, 32},
				{500, 1024},
				{200, 64},
			},
		},
		{
			name:         "extreme_mixed_solvers",
			userGasLimit: 1000,
			userDataSize: 2048,
			solverConfigs: []struct {
				gas      int64
				dataSize int
			}{
				{50, 0},
				{2000, 4096},
				{100, 16},
				{800, 512},
				{300, 256},
			},
		},
		{
			name:         "many_tiny_solvers",
			userGasLimit: 500,
			userDataSize: 256,
			solverConfigs: []struct {
				gas      int64
				dataSize int
			}{
				{10, 1}, {20, 2}, {30, 3}, {40, 4}, {50, 5},
				{60, 6}, {70, 7}, {80, 8}, {90, 9}, {100, 10},
			},
		},
	}

	// Test on berachain (primary test chain)
	chainConfig := testChainConfigs[0] // berachain

	ethClient, err := gethEthClient.Dial(chainConfig.rpcUrl)
	if err != nil {
		t.Fatalf("failed to create eth client: %v", err)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	sdk, err := NewAtlasSdk([]interface{}{ethClient}, nil)
	if err != nil {
		t.Fatalf("failed to create atlas sdk: %v", err)
	}

	for _, tc := range mixedTestCases {
		t.Run(tc.name, func(t *testing.T) {
			userOp := generateUserOperationWithParams(chainId.Uint64(), chainConfig.atlasVersion, tc.userGasLimit, tc.userDataSize, 600)

			dAppConfig, err := sdk.GetDAppConfig(chainId.Uint64(), &chainConfig.atlasVersion, userOp, chainConfig.dAppControlAddr)
			if err != nil {
				t.Fatalf("failed to get dapp config: %v", err)
			}

			// Use the actual dAppConfig values
			userOp.SetControl(chainConfig.dAppControlAddr)
			userOp.SetDappGasLimit(dAppConfig.DappGasLimit)
			userOp.SetSolverGasLimit(dAppConfig.SolverGasLimit)
			userOp.SetCallConfig(dAppConfig.CallConfig)

			// Generate mixed solver operations
			solverOps := make(types.SolverOperations, len(tc.solverConfigs))
			for i, config := range tc.solverConfigs {
				solverOps[i] = generateSolverOperationWithParams(config.gas, config.dataSize, int64(500+i*100))
			}

			onChainLimit, err := sdk.estimateMetacallGasLimit(
				chainId.Uint64(),
				&chainConfig.atlasVersion,
				userOp,
				solverOps,
				big.NewInt(1000000000),
			)
			if err != nil {
				t.Fatalf("failed to estimate metacall gas limit (on-chain): %v", err)
			}

			localLimit, err := sdk.EstimateMetacallGasLimitLocal(
				chainId.Uint64(),
				&chainConfig.atlasVersion,
				userOp,
				solverOps,
			)
			if err != nil {
				t.Fatalf("failed to estimate metacall gas limit (local): %v", err)
			}

			// Log detailed information
			t.Logf("Test case: %s", tc.name)
			t.Logf("  DApp config: dappGas=%d, solverGas=%d, callConfig=%d",
				dAppConfig.DappGasLimit, dAppConfig.SolverGasLimit, dAppConfig.CallConfig)
			t.Logf("  User: gas=%d, dataSize=%d", tc.userGasLimit, tc.userDataSize)
			solverInfo := make([]string, len(tc.solverConfigs))
			for i, config := range tc.solverConfigs {
				solverInfo[i] = fmt.Sprintf("(gas=%d,data=%d)", config.gas, config.dataSize)
			}
			t.Logf("  Solvers: %s", strings.Join(solverInfo, ", "))
			t.Logf("  OnChain: %d, Local: %d, Diff: %d", onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))

			if onChainLimit != localLimit {
				t.Errorf("Gas limit mismatch for %s: onChain=%d, local=%d, diff=%d",
					tc.name, onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))
			}
		})
	}
}

// Test different Atlas versions (v1.6.2 and above support local estimation)
func TestEstimateMetacallGasLimitVersions(t *testing.T) {
	t.Skip("Skipped in CI because RPC calls are unreliable on public RPCs, run locally with good RPCs instead")

	testConfig := struct {
		userGasLimit   int64
		userDataSize   int
		solverCount    int
		solverGasLimit int64
		solverDataSize int
	}{
		userGasLimit:   300,
		userDataSize:   256,
		solverCount:    3,
		solverGasLimit: 400,
		solverDataSize: 128,
	}

	// Test each chain configuration which has different versions
	for _, chainConfig := range testChainConfigs {
		t.Run(fmt.Sprintf("chain_%s_version_%s", chainConfig.name, chainConfig.atlasVersion), func(t *testing.T) {
			ethClient, err := gethEthClient.Dial(chainConfig.rpcUrl)
			if err != nil {
				t.Skipf("Skipping %s tests: failed to create eth client: %v", chainConfig.name, err)
			}

			chainId, err := ethClient.ChainID(context.Background())
			if err != nil {
				t.Skipf("Skipping %s tests: failed to get chain id: %v", chainConfig.name, err)
			}

			sdk, err := NewAtlasSdk([]interface{}{ethClient}, nil)
			if err != nil {
				t.Skipf("Skipping %s tests: failed to create atlas sdk: %v", chainConfig.name, err)
			}

			userOp := generateUserOperationWithParams(chainId.Uint64(), chainConfig.atlasVersion, testConfig.userGasLimit, testConfig.userDataSize, 600)

			dAppConfig, err := sdk.GetDAppConfig(chainId.Uint64(), &chainConfig.atlasVersion, userOp, chainConfig.dAppControlAddr)
			if err != nil {
				t.Fatalf("failed to get dapp config for version %s: %v", chainConfig.atlasVersion, err)
			}

			// Use the actual dAppConfig values
			userOp.SetControl(chainConfig.dAppControlAddr)
			userOp.SetDappGasLimit(dAppConfig.DappGasLimit)
			userOp.SetSolverGasLimit(dAppConfig.SolverGasLimit)
			userOp.SetCallConfig(dAppConfig.CallConfig)

			solverOps := generateSolverOperations(testConfig.solverCount, testConfig.solverGasLimit, testConfig.solverDataSize)

			onChainLimit, err := sdk.estimateMetacallGasLimit(
				chainId.Uint64(),
				&chainConfig.atlasVersion,
				userOp,
				solverOps,
				big.NewInt(1000000000),
			)
			if err != nil {
				t.Fatalf("failed to estimate metacall gas limit (on-chain) for version %s: %v", chainConfig.atlasVersion, err)
			}

			localLimit, err := sdk.EstimateMetacallGasLimitLocal(
				chainId.Uint64(),
				&chainConfig.atlasVersion,
				userOp,
				solverOps,
			)
			if err != nil {
				t.Fatalf("failed to estimate metacall gas limit (local) for version %s: %v", chainConfig.atlasVersion, err)
			}

			t.Logf("Chain: %s, Version: %s", chainConfig.name, chainConfig.atlasVersion)
			t.Logf("  DApp config: dappGas=%d, solverGas=%d, callConfig=%d",
				dAppConfig.DappGasLimit, dAppConfig.SolverGasLimit, dAppConfig.CallConfig)
			t.Logf("  OnChain: %d, Local: %d, Diff: %d", onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))

			if onChainLimit != localLimit {
				t.Errorf("Gas limit mismatch for %s/%s: onChain=%d, local=%d, diff=%d",
					chainConfig.name, chainConfig.atlasVersion, onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))
			}
		})
	}
}

// Stress test with extreme configurations
func TestEstimateMetacallGasLimitStressTest(t *testing.T) {
	t.Skip("Skipped in CI because RPC calls are unreliable on public RPCs, run locally with good RPCs instead")

	stressTestCases := []struct {
		name           string
		userGasLimit   int64
		userDataSize   int
		solverCount    int
		solverGasLimit int64
		solverDataSize int
	}{
		// Extreme cases that might reveal edge cases in calculation
		{"zero_everything", 0, 0, 0, 0, 0},
		{"max_data_sizes", 1000, 16384, 1, 1000, 16384},
		{"many_zero_gas_solvers", 500, 100, 50, 0, 10},
		{"one_massive_solver", 200, 64, 1, 10000, 32768},
		{"asymmetric_extreme", 10, 32768, 25, 5000, 1},
		{"high_gas_limits", 1000, 1000, 5, 1000, 1000},
		{"minimal_with_large_count", 1, 1, 100, 1, 1},
	}

	// Test on berachain (primary test chain)
	chainConfig := testChainConfigs[0] // berachain

	ethClient, err := gethEthClient.Dial(chainConfig.rpcUrl)
	if err != nil {
		t.Fatalf("failed to create eth client: %v", err)
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	sdk, err := NewAtlasSdk([]interface{}{ethClient}, nil)
	if err != nil {
		t.Fatalf("failed to create atlas sdk: %v", err)
	}

	for _, tc := range stressTestCases {
		t.Run(tc.name, func(t *testing.T) {
			userOp := generateUserOperationWithParams(chainId.Uint64(), chainConfig.atlasVersion, tc.userGasLimit, tc.userDataSize, 600)

			dAppConfig, err := sdk.GetDAppConfig(chainId.Uint64(), &chainConfig.atlasVersion, userOp, chainConfig.dAppControlAddr)
			if err != nil {
				t.Fatalf("failed to get dapp config: %v", err)
			}

			// Use the actual dAppConfig values
			userOp.SetControl(chainConfig.dAppControlAddr)
			userOp.SetDappGasLimit(dAppConfig.DappGasLimit)
			userOp.SetSolverGasLimit(dAppConfig.SolverGasLimit)
			userOp.SetCallConfig(dAppConfig.CallConfig)

			solverOps := generateSolverOperations(tc.solverCount, tc.solverGasLimit, tc.solverDataSize)

			onChainLimit, err := sdk.estimateMetacallGasLimit(
				chainId.Uint64(),
				&chainConfig.atlasVersion,
				userOp,
				solverOps,
				big.NewInt(1000000000),
			)
			if err != nil {
				t.Logf("On-chain estimation failed for %s (expected for extreme cases): %v", tc.name, err)
				return // Skip if on-chain fails with extreme parameters
			}

			localLimit, err := sdk.EstimateMetacallGasLimitLocal(
				chainId.Uint64(),
				&chainConfig.atlasVersion,
				userOp,
				solverOps,
			)
			if err != nil {
				t.Logf("Local estimation failed for %s (expected for extreme cases): %v", tc.name, err)
				return // Skip if local fails with extreme parameters
			}

			t.Logf("Stress test: %s", tc.name)
			t.Logf("  DApp config: dappGas=%d, solverGas=%d, callConfig=%d",
				dAppConfig.DappGasLimit, dAppConfig.SolverGasLimit, dAppConfig.CallConfig)
			t.Logf("  Config: user(gas=%d,data=%d), solvers(count=%d,gas=%d,data=%d)",
				tc.userGasLimit, tc.userDataSize, tc.solverCount, tc.solverGasLimit, tc.solverDataSize)
			t.Logf("  OnChain: %d, Local: %d, Diff: %d", onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))

			if onChainLimit != localLimit {
				t.Errorf("Gas limit mismatch for stress test %s: onChain=%d, local=%d, diff=%d",
					tc.name, onChainLimit, localLimit, int64(onChainLimit)-int64(localLimit))
			}
		})
	}
}
