package core

import (
	"context"
	"math/big"
	"testing"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
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

func TestEstimateMetacallGasLimitLocal(t *testing.T) {
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
