package core

import (
	"errors"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	callConfigFunction              = "CALL_CONFIG"
	getDAppConfigFunction           = "getDAppConfig"
	getSolverGasLimitFunction       = "getSolverGasLimit"
	getDAppGasLimitFunction         = "getDAppGasLimit"
	getBundlerSurchargeRateFunction = "getBundlerSurchargeRate"
)

var (
	ErrFunctionSupportedFrom_1_5 = errors.New("function supported from Atlas v1.5 and above")
	ErrFunctionSupportedFrom_1_6 = errors.New("function supported from Atlas v1.6 and above")
)

func (sdk *AtlasSdk) getDAppControlContract(chainId uint64, version *string, dAppControlAddr common.Address) (*bind.BoundContract, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	dAppControlAbi, err := contract.GetDAppControlAbi(version)
	if err != nil {
		return nil, err
	}

	return bind.NewBoundContract(dAppControlAddr, *dAppControlAbi, ethClient, ethClient, ethClient), nil
}

func (sdk *AtlasSdk) GetDAppCallConfig(chainId uint64, version *string, dAppControlAddr common.Address) (uint32, error) {
	dAppControlContract, err := sdk.getDAppControlContract(chainId, version, dAppControlAddr)
	if err != nil {
		return 0, err
	}

	var callConfig uint32

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	if err := dAppControlContract.Call(callOpts, &[]interface{}{&callConfig}, callConfigFunction); err != nil {
		return 0, err
	}

	return callConfig, nil
}

func (sdk *AtlasSdk) GetDAppConfig(chainId uint64, version *string, userOp *types.UserOperation, dAppControlAddr common.Address) (*types.DAppConfig, error) {
	dAppControlContract, err := sdk.getDAppControlContract(chainId, version, dAppControlAddr)
	if err != nil {
		return nil, err
	}

	var out []interface{}

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	if err := dAppControlContract.Call(callOpts, &out, getDAppConfigFunction, userOp.ToParams()); err != nil {
		return nil, err
	}

	dAppConfig := *abi.ConvertType(out[0], new(types.DAppConfig)).(*types.DAppConfig)

	return &dAppConfig, nil
}

func (sdk *AtlasSdk) GetDAppSolverGasLimit(chainId uint64, version *string, dAppControlAddr common.Address) (uint32, error) {
	dAppControlContract, err := sdk.getDAppControlContract(chainId, version, dAppControlAddr)
	if err != nil {
		return 0, err
	}

	var solverGasLimit uint32

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	if err := dAppControlContract.Call(callOpts, &[]interface{}{&solverGasLimit}, getSolverGasLimitFunction); err != nil {
		return 0, err
	}

	return solverGasLimit, nil
}

func (sdk *AtlasSdk) GetDAppGasLimit(chainId uint64, version *string, dAppControlAddr common.Address) (uint32, error) {
	minSupport := config.AtlasV_1_5

	gte_1_5, err := config.IsVersionAtLeast(version, &minSupport)
	if err != nil {
		return 0, err
	}

	if !gte_1_5 {
		return 0, ErrFunctionSupportedFrom_1_5
	}

	dAppControlContract, err := sdk.getDAppControlContract(chainId, version, dAppControlAddr)
	if err != nil {
		return 0, err
	}

	var dAppGasLimit uint32

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	if err := dAppControlContract.Call(callOpts, &[]interface{}{&dAppGasLimit}, getDAppGasLimitFunction); err != nil {
		return 0, err
	}

	return dAppGasLimit, nil
}

func (sdk *AtlasSdk) GetBundlerSurchargeRate(chainId uint64, version *string, dAppControlAddr common.Address) (*big.Int, error) {
	minSupport := config.AtlasV_1_6

	gte_1_6, err := config.IsVersionAtLeast(version, &minSupport)
	if err != nil {
		return nil, err
	}

	if !gte_1_6 {
		return nil, ErrFunctionSupportedFrom_1_6
	}

	dAppControlContract, err := sdk.getDAppControlContract(chainId, version, dAppControlAddr)
	if err != nil {
		return nil, err
	}

	var bundlerSurchargeRate *big.Int

	callOpts, cancel := NewCallOptsWithNetworkDeadline()
	defer cancel()

	if err := dAppControlContract.Call(callOpts, &[]interface{}{&bundlerSurchargeRate}, getBundlerSurchargeRateFunction); err != nil {
		return nil, err
	}

	return bundlerSurchargeRate, nil
}
