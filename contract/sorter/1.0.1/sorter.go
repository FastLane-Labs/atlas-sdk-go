// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sorter_1_0_1

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SolverOperation is an auto generated low-level Go binding around an user-defined struct.
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Dapp         common.Address
	Control      common.Address
	CallConfig   uint32
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

// SorterMetaData contains all meta data concerning the Sorter contract.
var SorterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ATLAS\",\"outputs\":[{\"internalType\":\"contractIAtlas\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFICATION\",\"outputs\":[{\"internalType\":\"contractIAtlasVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"}],\"name\":\"sortBids\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SorterABI is the input ABI used to generate the binding from.
// Deprecated: Use SorterMetaData.ABI instead.
var SorterABI = SorterMetaData.ABI

// Sorter is an auto generated Go binding around an Ethereum contract.
type Sorter struct {
	SorterCaller     // Read-only binding to the contract
	SorterTransactor // Write-only binding to the contract
	SorterFilterer   // Log filterer for contract events
}

// SorterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SorterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SorterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SorterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SorterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SorterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SorterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SorterSession struct {
	Contract     *Sorter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SorterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SorterCallerSession struct {
	Contract *SorterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SorterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SorterTransactorSession struct {
	Contract     *SorterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SorterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SorterRaw struct {
	Contract *Sorter // Generic contract binding to access the raw methods on
}

// SorterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SorterCallerRaw struct {
	Contract *SorterCaller // Generic read-only contract binding to access the raw methods on
}

// SorterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SorterTransactorRaw struct {
	Contract *SorterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSorter creates a new instance of Sorter, bound to a specific deployed contract.
func NewSorter(address common.Address, backend bind.ContractBackend) (*Sorter, error) {
	contract, err := bindSorter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sorter{SorterCaller: SorterCaller{contract: contract}, SorterTransactor: SorterTransactor{contract: contract}, SorterFilterer: SorterFilterer{contract: contract}}, nil
}

// NewSorterCaller creates a new read-only instance of Sorter, bound to a specific deployed contract.
func NewSorterCaller(address common.Address, caller bind.ContractCaller) (*SorterCaller, error) {
	contract, err := bindSorter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SorterCaller{contract: contract}, nil
}

// NewSorterTransactor creates a new write-only instance of Sorter, bound to a specific deployed contract.
func NewSorterTransactor(address common.Address, transactor bind.ContractTransactor) (*SorterTransactor, error) {
	contract, err := bindSorter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SorterTransactor{contract: contract}, nil
}

// NewSorterFilterer creates a new log filterer instance of Sorter, bound to a specific deployed contract.
func NewSorterFilterer(address common.Address, filterer bind.ContractFilterer) (*SorterFilterer, error) {
	contract, err := bindSorter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SorterFilterer{contract: contract}, nil
}

// bindSorter binds a generic wrapper to an already deployed contract.
func bindSorter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SorterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sorter *SorterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sorter.Contract.SorterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sorter *SorterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sorter.Contract.SorterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sorter *SorterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sorter.Contract.SorterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sorter *SorterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sorter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sorter *SorterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sorter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sorter *SorterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sorter.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_Sorter *SorterCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sorter.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_Sorter *SorterSession) ATLAS() (common.Address, error) {
	return _Sorter.Contract.ATLAS(&_Sorter.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_Sorter *SorterCallerSession) ATLAS() (common.Address, error) {
	return _Sorter.Contract.ATLAS(&_Sorter.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Sorter *SorterCaller) VERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sorter.contract.Call(opts, &out, "VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Sorter *SorterSession) VERIFICATION() (common.Address, error) {
	return _Sorter.Contract.VERIFICATION(&_Sorter.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Sorter *SorterCallerSession) VERIFICATION() (common.Address, error) {
	return _Sorter.Contract.VERIFICATION(&_Sorter.CallOpts)
}

// SortBids is a free data retrieval call binding the contract method 0x3caf7efa.
//
// Solidity: function sortBids((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps) view returns((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[])
func (_Sorter *SorterCaller) SortBids(opts *bind.CallOpts, userOp UserOperation, solverOps []SolverOperation) ([]SolverOperation, error) {
	var out []interface{}
	err := _Sorter.contract.Call(opts, &out, "sortBids", userOp, solverOps)

	if err != nil {
		return *new([]SolverOperation), err
	}

	out0 := *abi.ConvertType(out[0], new([]SolverOperation)).(*[]SolverOperation)

	return out0, err

}

// SortBids is a free data retrieval call binding the contract method 0x3caf7efa.
//
// Solidity: function sortBids((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps) view returns((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[])
func (_Sorter *SorterSession) SortBids(userOp UserOperation, solverOps []SolverOperation) ([]SolverOperation, error) {
	return _Sorter.Contract.SortBids(&_Sorter.CallOpts, userOp, solverOps)
}

// SortBids is a free data retrieval call binding the contract method 0x3caf7efa.
//
// Solidity: function sortBids((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps) view returns((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[])
func (_Sorter *SorterCallerSession) SortBids(userOp UserOperation, solverOps []SolverOperation) ([]SolverOperation, error) {
	return _Sorter.Contract.SortBids(&_Sorter.CallOpts, userOp, solverOps)
}
