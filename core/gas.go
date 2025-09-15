package core

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	USER_OP_STATIC_LENGTH     = uint64(608)
	DAPP_OP_LENGTH            = uint64(352)
	_BASE_TX_GAS_USED         = uint64(21_000)
	_BID_FIND_OVERHEAD        = uint64(5_000)
	_SOLVER_OP_STATIC_LENGTH  = uint64(608)
	_EXECUTE_SOLVER_OVERHEAD  = uint64(45_000)
	_EXTRA_CALLDATA_LENGTH    = uint64(238)
	_PRE_EXECUTE_METACALL_GAS = uint64(150_000)
	_POST_SETTLE_METACALL_GAS = uint64(70_000)
	_GAS_PER_CALLDATA_BYTE    = uint64(8)
)

var (
	excludedOffchainChainIds = map[uint64]struct{}{
		42161:    {}, // Arbitrum One
		421614:   {}, // Arbitrum Sepolia
		42170:    {}, // Arbitrum Nova
		98866:    {}, // Plume
		98867:    {}, // Plume Testnet
		10:       {}, // Optimism
		11155420: {}, // Optimism Sepolia
		// 8453:     {}, // Base
		// 84532:    {}, // Base Sepolia
		130:  {}, // Unichain
		1301: {}, // Unichain Testnet
	}
)

func (sdk *AtlasSdk) EstimateMetacallGasLimit(chainId uint64, version *string, userOp *types.UserOperation, solverOps types.SolverOperations, gasPrice *big.Int) (uint64, error) {
	// Try local calculation, fallback to RPC call
	if r, err := sdk.EstimateMetacallGasLimitLocal(chainId, version, userOp, solverOps); err == nil {
		return r, nil
	}

	return sdk.estimateMetacallGasLimit(chainId, version, userOp, solverOps, gasPrice)
}

func (sdk *AtlasSdk) estimateMetacallGasLimit(chainId uint64, version *string, userOp *types.UserOperation, solverOps types.SolverOperations, gasPrice *big.Int) (uint64, error) {
	minVersion := config.AtlasV_1_5
	if minSupport, err := config.IsVersionAtLeast(version, &minVersion); err != nil || !minSupport {
		return 0, fmt.Errorf("metacall gas limit estimation is only supported for Atlas v1.5 and above")
	}

	simulatorAbi, err := contract.GetSimulatorAbi(version)
	if err != nil {
		return 0, fmt.Errorf("failed to get simulator abi: %w", err)
	}

	simulatorAddr, err := config.GetSimulatorAddress(chainId, version)
	if err != nil {
		return 0, fmt.Errorf("failed to get simulator address: %w", err)
	}

	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return 0, fmt.Errorf("failed to get eth client: %w", err)
	}

	_solverOps := make([]types.SolverOperation, len(solverOps))
	for i, op := range solverOps {
		_solverOps[i] = *op
	}

	pData, err := simulatorAbi.Pack(estimateMetacallGasLimitFunction, userOp.ToParams(), _solverOps)
	if err != nil {
		return 0, fmt.Errorf("failed to pack %s: %w", estimateMetacallGasLimitFunction, err)
	}

	overrides := map[common.Address]map[string]interface{}{
		common.HexToAddress("0x0000000000000000000000000000000000000000"): {
			"balance": (*hexutil.Big)(maxUint96),
		},
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	var result string
	err = ethClient.Client().CallContext(ctx, &result, "eth_call", map[string]interface{}{
		"to":       simulatorAddr.Hex(),
		"gasPrice": (*hexutil.Big)(gasPrice),
		"data":     hexutil.Encode(pData),
	}, "latest", overrides)
	if err != nil {
		return 0, fmt.Errorf("failed to call %s: %w - pData %s - simulatorAddr %s", estimateMetacallGasLimitFunction, err, hex.EncodeToString(pData), simulatorAddr.Hex())
	}

	bData, err := hexutil.Decode(result)
	if err != nil {
		return 0, fmt.Errorf("failed to decode result: %w", err)
	}

	gasLimit, err := simulatorAbi.Unpack(estimateMetacallGasLimitFunction, bData)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack %s: %w", estimateMetacallGasLimitFunction, err)
	}

	if len(gasLimit) != 1 {
		return 0, fmt.Errorf("expected 1 gas limit, got %d", len(gasLimit))
	}

	return gasLimit[0].(*big.Int).Uint64(), nil
}

func (sdk *AtlasSdk) EstimateMetacallGasLimitLocal(chainId uint64, version *string, userOp *types.UserOperation, solverOps types.SolverOperations) (uint64, error) {
	minVersion := config.AtlasV_1_6_2
	if minSupport, err := config.IsVersionAtLeast(version, &minVersion); err != nil || !minSupport {
		return 0, fmt.Errorf("metacall gas limit estimation offchain is only supported for Atlas v1.6.2 and above")
	}

	if _, ok := excludedOffchainChainIds[chainId]; ok {
		return 0, fmt.Errorf("metacall gas limit estimation offchain is not supported for chainId %d", chainId)
	}

	// Calculations are based on: https://github.com/FastLane-Labs/atlas/blob/atlas-v1.6.2/src/contracts/helpers/Simulator.sol#L53

	nonSolverCalldataLength := uint64(len(userOp.GetData())) + USER_OP_STATIC_LENGTH + DAPP_OP_LENGTH + _EXTRA_CALLDATA_LENGTH
	solverOpsLen := uint64(len(solverOps))

	var (
		solverDataLenSum       uint64
		allSolversExecutionGas uint64
	)

	for _, solverOp := range solverOps {
		solverDataLenSum += uint64(len(solverOp.Data))
		allSolversExecutionGas += min(solverOp.Gas.Uint64(), uint64(userOp.GetSolverGasLimit()))
	}

	metacallCalldataGas := (solverDataLenSum + (_SOLVER_OP_STATIC_LENGTH * solverOpsLen)) * _GAS_PER_CALLDATA_BYTE
	metacallCalldataGas += nonSolverCalldataLength * _GAS_PER_CALLDATA_BYTE

	executeSolverOverhead := _EXECUTE_SOLVER_OVERHEAD * solverOpsLen
	if config.IsMonad(chainId) {
		// On Monad, the _EXECUTE_SOLVER_OVERHEAD per solverOp is not included in the suggested gas limit.
		executeSolverOverhead = 0
	}

	metacallExecutionGas := _BASE_TX_GAS_USED + _PRE_EXECUTE_METACALL_GAS + _POST_SETTLE_METACALL_GAS + userOp.GetGas().Uint64() + uint64(userOp.GetDappGasLimit()) + allSolversExecutionGas + executeSolverOverhead

	exPostBids, err := utils.FlagExPostBids(userOp.GetCallConfig(), version)
	if err != nil {
		return 0, fmt.Errorf("failed to get exPostBids: %w", err)
	}

	if exPostBids {
		metacallExecutionGas += solverOpsLen*(_BID_FIND_OVERHEAD+_EXECUTE_SOLVER_OVERHEAD) + allSolversExecutionGas
	}

	return metacallCalldataGas + metacallExecutionGas, nil
}
