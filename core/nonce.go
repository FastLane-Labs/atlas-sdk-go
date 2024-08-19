package core

import (
	"errors"
	"math/big"

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

func (sdk *AtlasSdk) GetUserNextNonce(chainId uint64, user common.Address, callConfig uint32) (*big.Int, error) {
	contract, ok := sdk.atlasVerificationContract[chainId]
	if !ok {
		return nil, errors.New("atlasVerification contract not found")
	}

	var (
		nonce *big.Int
		err   error
	)

	sdk.noncesMu.Lock()
	defer sdk.noncesMu.Unlock()

	if utils.FlagUserNoncesSequential(callConfig) {
		nonce, err = contract.GetUserNextNonce(nil, user, true)
		if err != nil {
			return nil, err
		}
	} else {
		if _, ok := sdk.userLastNonSequentialNonce[chainId]; !ok {
			sdk.userLastNonSequentialNonce[chainId] = make(map[common.Address]*big.Int)
		}

		lastNonce := sdk.userLastNonSequentialNonce[chainId][user]

		if lastNonce == nil {
			nonce, err = contract.GetUserNextNonce(nil, user, false)
			if err != nil {
				return nil, err
			}
		} else {
			nonce, err = contract.GetUserNextNonSeqNonceAfter(nil, user, lastNonce)
			if err != nil {
				return nil, err
			}
		}

		sdk.userLastNonSequentialNonce[chainId][user] = nonce
	}

	return nonce, nil
}
