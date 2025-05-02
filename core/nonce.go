package core

import (
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/FastLane-Labs/atlas-sdk-go/types"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const (
	getUserNextNonceFunction            = "getUserNextNonce"
	getUserNextNonSeqNonceAfterFunction = "getUserNextNonSeqNonceAfter"
	getDAppNextNonceFunction            = "getDAppNextNonce"
)

func (sdk *AtlasSdk) SetUserNonce(chainId uint64, version *string, userOp *types.UserOperation) error {
	nonce, err := sdk.GetUserNextNonce(chainId, version, userOp.GetFrom(), userOp.GetCallConfig())
	if err != nil {
		return err
	}

	userOp.SetNonce(nonce)

	return nil
}

func (sdk *AtlasSdk) GetUserNextNonce(chainId uint64, version *string, user common.Address, callConfig uint32) (*big.Int, error) {
	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	atlasVerificationAddr, err := config.GetAtlasVerificationAddress(chainId, version)
	if err != nil {
		return nil, err
	}

	atlasVerificationAbi, err := contract.GetAtlasVerificationAbi(version)
	if err != nil {
		return nil, err
	}

	v := config.GetVersion(version)

	mu, ok := sdk.noncesMu[chainId][v]
	if !ok {
		// This can happen if a chain config has been overridden
		mu = &sdk.mu
	}

	var pData []byte

	mu.Lock()
	defer mu.Unlock()

	userNoncesSequential, err := utils.FlagUserNoncesSequential(callConfig, version)
	if err != nil {
		return nil, err
	}

	if userNoncesSequential {
		pData, err = atlasVerificationAbi.Pack(getUserNextNonceFunction, user, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack %s: %w", getUserNextNonceFunction, err)
		}
	} else {
		if _, ok := sdk.userLastNonSequentialNonce[chainId]; !ok {
			sdk.userLastNonSequentialNonce[chainId] = make(map[string]map[common.Address]*big.Int)
		}

		if _, ok := sdk.userLastNonSequentialNonce[chainId][v]; !ok {
			sdk.userLastNonSequentialNonce[chainId][v] = make(map[common.Address]*big.Int)
		}

		lastNonce := sdk.userLastNonSequentialNonce[chainId][v][user]

		if lastNonce == nil {
			pData, err = atlasVerificationAbi.Pack(getUserNextNonceFunction, user, false)
			if err != nil {
				return nil, fmt.Errorf("failed to pack %s: %w", getUserNextNonceFunction, err)
			}
		} else {
			pData, err = atlasVerificationAbi.Pack(getUserNextNonSeqNonceAfterFunction, user, lastNonce)
			if err != nil {
				return nil, fmt.Errorf("failed to pack %s: %w", getUserNextNonSeqNonceAfterFunction, err)
			}
		}
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:   &atlasVerificationAddr,
			Data: pData,
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", getUserNextNonceFunction, err)
	}

	_nonce, err := atlasVerificationAbi.Unpack(getUserNextNonceFunction, bData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", getUserNextNonceFunction, err)
	}

	nonce, ok := _nonce[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to cast %s: %w", getUserNextNonceFunction, err)
	}

	if !userNoncesSequential {
		sdk.userLastNonSequentialNonce[chainId][v][user] = nonce
	}

	return nonce, nil
}

func (sdk *AtlasSdk) GetDAppNextNonce(chainId uint64, version *string, dApp common.Address, callConfig uint32) (*big.Int, error) {
	dappNoncesSequential, err := utils.FlagDappNoncesSequential(callConfig, version)
	if err != nil {
		return nil, err
	}

	if !dappNoncesSequential {
		// Nonce not needed for non-sequential dapp calls
		return new(big.Int).Set(common.Big0), nil
	}

	ethClient, err := sdk.getEthClient(chainId)
	if err != nil {
		return nil, err
	}

	atlasVerificationAddr, err := config.GetAtlasVerificationAddress(chainId, version)
	if err != nil {
		return nil, err
	}

	atlasVerificationAbi, err := contract.GetAtlasVerificationAbi(version)
	if err != nil {
		return nil, err
	}

	pData, err := atlasVerificationAbi.Pack(getDAppNextNonceFunction, dApp)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s: %w", getDAppNextNonceFunction, err)
	}

	ctx, cancel := NewContextWithNetworkDeadline()
	defer cancel()

	bData, err := ethClient.CallContract(
		ctx,
		ethereum.CallMsg{
			To:   &atlasVerificationAddr,
			Data: pData,
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", getDAppNextNonceFunction, err)
	}

	_nonce, err := atlasVerificationAbi.Unpack(getDAppNextNonceFunction, bData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", getDAppNextNonceFunction, err)
	}

	nonce, ok := _nonce[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to cast %s: %w", getDAppNextNonceFunction, err)
	}

	return nonce, nil
}
