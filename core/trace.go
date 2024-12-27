package core

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/contract"
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
	traceCallMethod     = "debug_traceCall"
	solverTxResultEvent = "SolverTxResult"
)

var (
	unsupportedTracingVersion = map[string]struct{}{
		"1.0": {},
		"1.1": {},
	}
)

func (sdk *AtlasSdk) GetSolverBidAmountFromTrace(chainId uint64, version *string, trace *callFrame) (*big.Int, error) {
	if _, ok := unsupportedTracingVersion[*version]; ok {
		return nil, fmt.Errorf("unsupported Atlas version %s for tracing", *version)
	}

	atlasAddr, err := config.GetAtlasAddress(chainId, version)
	if err != nil {
		return nil, err
	}

	atlasAbi, err := contract.GetAtlasAbi(version)
	if err != nil {
		return nil, err
	}

	solverTxResultEvent, ok := atlasAbi.Events[solverTxResultEvent]
	if !ok {
		return nil, fmt.Errorf("solverTxResult event not found in Atlas ABI")
	}

	solverTxResult, err := inspectTraceLogs(trace, atlasAddr, solverTxResultEvent.ID)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	err = solverTxResultEvent.Inputs.UnpackIntoMap(result, solverTxResult.Data)
	if err != nil {
		return nil, err
	}

	bidAmount, ok := result["bidAmount"].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("bidAmount not found in result")
	}

	return bidAmount, nil
}

func inspectTraceLogs(trace *callFrame, emittingAddress common.Address, eventHash common.Hash) (*callLog, error) {
	for _, log := range trace.Logs {
		if log.Address == emittingAddress && log.Topics[0] == eventHash {
			return &log, nil
		}
	}

	if len(trace.Calls) > 0 {
		for _, call := range trace.Calls {
			if log, _ := inspectTraceLogs(&call, emittingAddress, eventHash); log != nil {
				return log, nil
			}
		}
	}

	return nil, errors.New("event not found in trace")
}
