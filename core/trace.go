package core

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type callLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    hexutil.Bytes  `json:"data"`
	// Position of the log relative to subcalls within the same trace
	// See https://github.com/ethereum/go-ethereum/pull/28389 for details
	Position hexutil.Uint `json:"position"`
}

type callFrame struct {
	From   common.Address  `json:"from"`
	To     *common.Address `json:"to"`
	Output hexutil.Bytes   `json:"output"`
	Error  string          `json:"error"`
	Calls  []callFrame     `json:"calls"`
	Logs   []callLog       `json:"logs"`
}

const (
	traceCallMethod        = "debug_traceCall"
	bidFindSuccessfulError = "BidFindSuccessful"
)

func (sdk *AtlasSdk) GetSolverBidAmountFromTrace(chainId uint64, version *string, trace *callFrame) (*big.Int, error) {
	atlasAddr, err := config.GetAtlasAddress(chainId, version)
	if err != nil {
		return nil, err
	}

	atlasAbi, err := contract.GetAtlasAbi(version)
	if err != nil {
		return nil, err
	}

	bidFindSuccessfulError, ok := atlasAbi.Errors[bidFindSuccessfulError]
	if !ok {
		return nil, fmt.Errorf("bidFindSuccessful error not found in Atlas ABI")
	}

	return inspectTraceExPostBidAmount(trace, atlasAddr, bidFindSuccessfulError)
}

func inspectTraceExPostBidAmount(trace *callFrame, atlasAddr common.Address, bidFindSuccessfulError abi.Error) (*big.Int, error) {
	if trace.From == atlasAddr && trace.To != nil && *trace.To == atlasAddr {
		if bytes.Equal(bidFindSuccessfulError.ID[:4], trace.Output[:4]) {
			res, err := bidFindSuccessfulError.Inputs.UnpackValues(trace.Output[4:])
			if err == nil {
				return res[0].(*big.Int), nil
			}
		}
	}

	if len(trace.Calls) > 0 {
		for _, call := range trace.Calls {
			if bidAmount, _ := inspectTraceExPostBidAmount(&call, atlasAddr, bidFindSuccessfulError); bidAmount != nil {
				return bidAmount, nil
			}
		}
	}

	return nil, errors.New("bid amount not found in trace")
}
