package types

import (
	"errors"
	"math/big"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	solverOpStructArgs = []abi.ArgumentMarshaling{
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "solver", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "userOpHash", Type: "bytes32", InternalType: "bytes32"},
		{Name: "bidToken", Type: "address", InternalType: "address"},
		{Name: "bidAmount", Type: "uint256", InternalType: "uint256"},
		{Name: "data", Type: "bytes", InternalType: "bytes"},
		{Name: "signature", Type: "bytes", InternalType: "bytes"},
	}

	solverOpSolType, _      = abi.NewType("tuple", "struct SolverOperation", solverOpStructArgs)
	solverOpArraySolType, _ = abi.NewType("tuple[]", "struct SolverOperation[]", solverOpStructArgs)

	solverOpArgs = abi.Arguments{
		{Type: solverOpSolType, Name: "solverOperation"},
	}

	solverOpArrayArgs = abi.Arguments{
		{Type: solverOpArraySolType, Name: "solverOperations"},
	}
)

// External representation of a solver operation,
// the relay receives and broadcasts solver operations in this format
type SolverOperationRaw struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Solver       common.Address `json:"solver"`
	Control      common.Address `json:"control"`
	UserOpHash   common.Hash    `json:"userOpHash"`
	BidToken     common.Address `json:"bidToken"`
	BidAmount    *hexutil.Big   `json:"bidAmount"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (s *SolverOperationRaw) Sanitize() {
	if s.Value == nil {
		s.Value = (*hexutil.Big)(big.NewInt(0))
	}

	if s.Gas == nil {
		s.Gas = (*hexutil.Big)(big.NewInt(0))
	}

	if s.MaxFeePerGas == nil {
		s.MaxFeePerGas = (*hexutil.Big)(big.NewInt(0))
	}

	if s.Deadline == nil {
		s.Deadline = (*hexutil.Big)(big.NewInt(0))
	}

	if s.BidAmount == nil {
		s.BidAmount = (*hexutil.Big)(big.NewInt(0))
	}
}

func (s *SolverOperationRaw) Decode() *SolverOperation {
	s.Sanitize()

	return &SolverOperation{
		From:         s.From,
		To:           s.To,
		Value:        s.Value.ToInt(),
		Gas:          s.Gas.ToInt(),
		MaxFeePerGas: s.MaxFeePerGas.ToInt(),
		Deadline:     s.Deadline.ToInt(),
		Solver:       s.Solver,
		Control:      s.Control,
		UserOpHash:   s.UserOpHash,
		BidToken:     s.BidToken,
		BidAmount:    s.BidAmount.ToInt(),
		Data:         s.Data,
		Signature:    s.Signature,
	}
}

// Internal representation of a solver operation
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   common.Hash
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}

func (s *SolverOperation) Sanitize() {
	if s.Value == nil {
		s.Value = big.NewInt(0)
	}

	if s.Gas == nil {
		s.Gas = big.NewInt(0)
	}

	if s.MaxFeePerGas == nil {
		s.MaxFeePerGas = big.NewInt(0)
	}

	if s.Deadline == nil {
		s.Deadline = big.NewInt(0)
	}

	if s.BidAmount == nil {
		s.BidAmount = big.NewInt(0)
	}
}

func (s *SolverOperation) EncodeToRaw() *SolverOperationRaw {
	s.Sanitize()

	return &SolverOperationRaw{
		From:         s.From,
		To:           s.To,
		Value:        (*hexutil.Big)(s.Value),
		Gas:          (*hexutil.Big)(s.Gas),
		MaxFeePerGas: (*hexutil.Big)(s.MaxFeePerGas),
		Deadline:     (*hexutil.Big)(s.Deadline),
		Solver:       s.Solver,
		Control:      s.Control,
		UserOpHash:   s.UserOpHash,
		BidToken:     s.BidToken,
		BidAmount:    (*hexutil.Big)(s.BidAmount),
		Data:         hexutil.Bytes(s.Data),
		Signature:    hexutil.Bytes(s.Signature),
	}
}

func (s *SolverOperation) Hash(chainId uint64, version *string) (common.Hash, error) {
	eip712Domain, err := config.GetEip712Domain(chainId, version)
	if err != nil {
		return common.Hash{}, err
	}

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       s.toTypedDataTypes(),
		PrimaryType: "SolverOperation",
		Domain:      *eip712Domain,
		Message:     s.toTypedDataMessage(),
	})

	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(hash), nil
}

func (s *SolverOperation) AbiEncode() ([]byte, error) {
	s.Sanitize()
	return solverOpArgs.Pack(&s)
}

func (s *SolverOperation) ValidateSignature(chainId uint64, version *string) error {
	if len(s.Signature) != 65 {
		return errors.New("invalid signature length")
	}

	solverOpHash, err := s.Hash(chainId, version)
	if err != nil {
		return err
	}

	signer, err := utils.RecoverSigner(solverOpHash.Bytes(), s.Signature)
	if err != nil {
		return err
	}

	if signer != s.From {
		return errors.New("invalid signature")
	}

	return nil
}

func (s *SolverOperation) toTypedDataTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"SolverOperation": []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "gas", Type: "uint256"},
			{Name: "maxFeePerGas", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "solver", Type: "address"},
			{Name: "control", Type: "address"},
			{Name: "userOpHash", Type: "bytes32"},
			{Name: "bidToken", Type: "address"},
			{Name: "bidAmount", Type: "uint256"},
			{Name: "data", Type: "bytes"},
		},
	}
}

func (s *SolverOperation) toTypedDataMessage() apitypes.TypedDataMessage {
	s.Sanitize()

	return apitypes.TypedDataMessage{
		"from":         s.From.Hex(),
		"to":           s.To.Hex(),
		"value":        new(big.Int).Set(s.Value),
		"gas":          new(big.Int).Set(s.Gas),
		"maxFeePerGas": new(big.Int).Set(s.MaxFeePerGas),
		"deadline":     new(big.Int).Set(s.Deadline),
		"solver":       s.Solver.Hex(),
		"control":      s.Control.Hex(),
		"userOpHash":   s.UserOpHash.Hex(),
		"bidToken":     s.BidToken.Hex(),
		"bidAmount":    new(big.Int).Set(s.BidAmount),
		"data":         s.Data,
	}
}

// Array of solver operations
type SolverOperations []*SolverOperation
type SolverOperationsRaw []*SolverOperationRaw

func (s SolverOperations) AbiEncode() ([]byte, error) {
	solverOps := make([]SolverOperation, len(s))
	for i, solverOp := range s {
		solverOps[i] = *solverOp
	}
	return solverOpArrayArgs.Pack(solverOps)
}

func (s SolverOperations) EncodeToRaw() SolverOperationsRaw {
	solverOps := make([]*SolverOperationRaw, len(s))
	for i, solverOp := range s {
		solverOps[i] = solverOp.EncodeToRaw()
	}
	return solverOps
}

func (s SolverOperations) Sanitize() {
	for _, solverOp := range s {
		solverOp.Sanitize()
	}
}

func (s SolverOperationsRaw) Decode() SolverOperations {
	solverOps := make([]*SolverOperation, len(s))
	for i, solverOp := range s {
		solverOps[i] = solverOp.Decode()
	}
	return solverOps
}
