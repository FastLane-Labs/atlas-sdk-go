package core

import (
	"errors"
	"math/big"
	"sync"

	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func (sdk *AtlasSdk) SetUserNonce(chainId uint64, userOp *types.UserOperation) error {
	nonce, err := sdk.GetUserNextNonce(chainId, userOp.From, userOp.CallConfig)
	if err != nil {
		return err
	}

	userOp.Nonce = nonce

	return nil
}

func (sdk *AtlasSdk) getUserNextNonce(chainId uint64, user common.Address, callConfig uint32) (*big.Int, error) {
	contract, ok := sdk.atlasVerificationContract[chainId]
	if !ok {
		return nil, errors.New("atlasVerification contract not found")
	}

	return contract.GetUserNextNonce(nil, user, utils.FlagUserNoncesSequential(callConfig))
}

func (sdk *AtlasSdk) getDAppNextNonce(chainId uint64, dApp common.Address, callConfig uint32) (*big.Int, error) {
	if !utils.FlagDappNoncesSequential(callConfig) {
		// Nonce not needed for non-sequential dapp calls
		return new(big.Int).Set(common.Big0), nil
	}

	contract, ok := sdk.atlasVerificationContract[chainId]
	if !ok {
		return nil, errors.New("atlasVerification contract not found")
	}

	return contract.GetDAppNextNonce(nil, dApp)
}

func (sdk *AtlasSdk) getNextNonce(
	chainId uint64,
	nonceMap map[uint64]map[common.Address]*big.Int,
	account common.Address,
	callConfig uint32,
	retrievalFunc func(uint64, common.Address, uint32) (*big.Int, error),
) (*big.Int, error) {
	if record, ok := nonceMap[chainId]; !ok || record[account] == nil {
		nonceMap[chainId] = make(map[common.Address]*big.Int)
		lastNonce, err := retrievalFunc(chainId, account, callConfig)
		if err != nil {
			return nil, err
		}
		nonceMap[chainId][account] = lastNonce
	}

	sdk.mu.Lock()
	defer sdk.mu.Unlock()

	return new(big.Int).Set(nonceMap[chainId][account].Add(nonceMap[chainId][account], common.Big1)), nil
}

func (sdk *AtlasSdk) getNonceMutex(chainId uint64, account common.Address) *sync.Mutex {
	sdk.mu.Lock()
	if _, ok := sdk.noncesMu[chainId]; !ok {
		sdk.noncesMu[chainId] = make(map[common.Address]*sync.Mutex)
	}
	if _, ok := sdk.noncesMu[chainId][account]; !ok {
		sdk.noncesMu[chainId][account] = &sync.Mutex{}
	}
	sdk.mu.Unlock()

	return sdk.noncesMu[chainId][account]
}

func (sdk *AtlasSdk) GetUserNextNonce(chainId uint64, user common.Address, callConfig uint32) (*big.Int, error) {
	nonceMu := sdk.getNonceMutex(chainId, user)
	nonceMu.Lock()
	defer nonceMu.Unlock()

	var nonceMap map[uint64]map[common.Address]*big.Int

	if utils.FlagUserNoncesSequential(callConfig) {
		nonceMap = sdk.userLastSequentialNonce
	} else {
		nonceMap = sdk.userLastNonSequentialNonce
	}

	return sdk.getNextNonce(chainId, nonceMap, user, callConfig, sdk.getUserNextNonce)
}

func (sdk *AtlasSdk) GetDAppNextNonce(chainId uint64, dApp common.Address, callConfig uint32) (*big.Int, error) {
	if !utils.FlagDappNoncesSequential(callConfig) {
		// Nonce not needed for non-sequential dapp calls
		return new(big.Int).Set(common.Big0), nil
	}

	nonceMu := sdk.getNonceMutex(chainId, dApp)
	nonceMu.Lock()
	defer nonceMu.Unlock()

	return sdk.getNextNonce(chainId, sdk.dAppLastSequentialNonce, dApp, callConfig, sdk.getDAppNextNonce)
}
