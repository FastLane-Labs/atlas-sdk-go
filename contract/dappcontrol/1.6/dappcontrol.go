// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappcontrol_1_6

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
	_ = abi.ConvertType
)

// CallConfig is an auto generated low-level Go binding around an user-defined struct.
type CallConfig struct {
	UserNoncesSequential      bool
	DappNoncesSequential      bool
	RequirePreOps             bool
	TrackPreOpsReturnData     bool
	TrackUserReturnData       bool
	DelegateUser              bool
	RequirePreSolver          bool
	RequirePostSolver         bool
	ZeroSolvers               bool
	ReuseUserOp               bool
	UserAuctioneer            bool
	SolverAuctioneer          bool
	UnknownAuctioneer         bool
	VerifyCallChainHash       bool
	ForwardReturnData         bool
	RequireFulfillment        bool
	TrustedOpHash             bool
	InvertBidValue            bool
	ExPostBids                bool
	MultipleSuccessfulSolvers bool
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To                   common.Address
	CallConfig           uint32
	BidToken             common.Address
	SolverGasLimit       uint32
	DappGasLimit         uint32
	BundlerSurchargeRate *big.Int
}

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
	From                 common.Address
	To                   common.Address
	Value                *big.Int
	Gas                  *big.Int
	MaxFeePerGas         *big.Int
	Nonce                *big.Int
	Deadline             *big.Int
	Dapp                 common.Address
	Control              common.Address
	CallConfig           uint32
	DappGasLimit         uint32
	SolverGasLimit       uint32
	BundlerSurchargeRate *big.Int
	SessionKey           common.Address
	Data                 []byte
	Signature            []byte
}

// DAppControlMetaData contains all meta data concerning the DAppControl contract.
var DAppControlMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ATLAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ATLAS_VERIFICATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CALL_CONFIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTROL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SOURCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptGovernance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocateValueCall\",\"inputs\":[{\"name\":\"solved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBidFormat\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dappGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bundlerSurchargeRate\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBidValue\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBundlerSurchargeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCallConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCallConfig\",\"components\":[{\"name\":\"userNoncesSequential\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dappNoncesSequential\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePreOps\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trackPreOpsReturnData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trackUserReturnData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"delegateUser\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePreSolver\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requirePostSolver\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"zeroSolvers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reuseUserOp\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"userAuctioneer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"solverAuctioneer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"unknownAuctioneer\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifyCallChainHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"forwardReturnData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"requireFulfillment\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"trustedOpHash\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"invertBidValue\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exPostBids\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"multipleSuccessfulSolvers\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppConfig\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dappGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bundlerSurchargeRate\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"dConfig\",\"type\":\"tuple\",\"internalType\":\"structDAppConfig\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dappGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bundlerSurchargeRate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppSignatory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSolverGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postSolverCall\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"preOpsCall\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dappGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bundlerSurchargeRate\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"preSolverCall\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requireSequentialDAppNonces\",\"inputs\":[],\"outputs\":[{\"name\":\"isSequential\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireSequentialUserNonces\",\"inputs\":[],\"outputs\":[{\"name\":\"isSequential\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferGovernance\",\"inputs\":[{\"name\":\"newGovernance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userDelegated\",\"inputs\":[],\"outputs\":[{\"name\":\"delegated\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GovernanceTransferStarted\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceTransferred\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlteredControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSolver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvertBidValueCannotBeExPostBids\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeDelegatecalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDelegatecall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAtlas\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPhase\",\"inputs\":[]}]",
}

// DAppControlABI is the input ABI used to generate the binding from.
// Deprecated: Use DAppControlMetaData.ABI instead.
var DAppControlABI = DAppControlMetaData.ABI

// DAppControl is an auto generated Go binding around an Ethereum contract.
type DAppControl struct {
	DAppControlCaller     // Read-only binding to the contract
	DAppControlTransactor // Write-only binding to the contract
	DAppControlFilterer   // Log filterer for contract events
}

// DAppControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAppControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAppControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAppControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAppControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAppControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAppControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAppControlSession struct {
	Contract     *DAppControl      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAppControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAppControlCallerSession struct {
	Contract *DAppControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DAppControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAppControlTransactorSession struct {
	Contract     *DAppControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DAppControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAppControlRaw struct {
	Contract *DAppControl // Generic contract binding to access the raw methods on
}

// DAppControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAppControlCallerRaw struct {
	Contract *DAppControlCaller // Generic read-only contract binding to access the raw methods on
}

// DAppControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAppControlTransactorRaw struct {
	Contract *DAppControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAppControl creates a new instance of DAppControl, bound to a specific deployed contract.
func NewDAppControl(address common.Address, backend bind.ContractBackend) (*DAppControl, error) {
	contract, err := bindDAppControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAppControl{DAppControlCaller: DAppControlCaller{contract: contract}, DAppControlTransactor: DAppControlTransactor{contract: contract}, DAppControlFilterer: DAppControlFilterer{contract: contract}}, nil
}

// NewDAppControlCaller creates a new read-only instance of DAppControl, bound to a specific deployed contract.
func NewDAppControlCaller(address common.Address, caller bind.ContractCaller) (*DAppControlCaller, error) {
	contract, err := bindDAppControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAppControlCaller{contract: contract}, nil
}

// NewDAppControlTransactor creates a new write-only instance of DAppControl, bound to a specific deployed contract.
func NewDAppControlTransactor(address common.Address, transactor bind.ContractTransactor) (*DAppControlTransactor, error) {
	contract, err := bindDAppControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAppControlTransactor{contract: contract}, nil
}

// NewDAppControlFilterer creates a new log filterer instance of DAppControl, bound to a specific deployed contract.
func NewDAppControlFilterer(address common.Address, filterer bind.ContractFilterer) (*DAppControlFilterer, error) {
	contract, err := bindDAppControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAppControlFilterer{contract: contract}, nil
}

// bindDAppControl binds a generic wrapper to an already deployed contract.
func bindDAppControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAppControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAppControl *DAppControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAppControl.Contract.DAppControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAppControl *DAppControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAppControl.Contract.DAppControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAppControl *DAppControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAppControl.Contract.DAppControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAppControl *DAppControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAppControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAppControl *DAppControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAppControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAppControl *DAppControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAppControl.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_DAppControl *DAppControlCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_DAppControl *DAppControlSession) ATLAS() (common.Address, error) {
	return _DAppControl.Contract.ATLAS(&_DAppControl.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_DAppControl *DAppControlCallerSession) ATLAS() (common.Address, error) {
	return _DAppControl.Contract.ATLAS(&_DAppControl.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_DAppControl *DAppControlCaller) ATLASVERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "ATLAS_VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_DAppControl *DAppControlSession) ATLASVERIFICATION() (common.Address, error) {
	return _DAppControl.Contract.ATLASVERIFICATION(&_DAppControl.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_DAppControl *DAppControlCallerSession) ATLASVERIFICATION() (common.Address, error) {
	return _DAppControl.Contract.ATLASVERIFICATION(&_DAppControl.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_DAppControl *DAppControlCaller) CALLCONFIG(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "CALL_CONFIG")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_DAppControl *DAppControlSession) CALLCONFIG() (uint32, error) {
	return _DAppControl.Contract.CALLCONFIG(&_DAppControl.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_DAppControl *DAppControlCallerSession) CALLCONFIG() (uint32, error) {
	return _DAppControl.Contract.CALLCONFIG(&_DAppControl.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_DAppControl *DAppControlCaller) CONTROL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "CONTROL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_DAppControl *DAppControlSession) CONTROL() (common.Address, error) {
	return _DAppControl.Contract.CONTROL(&_DAppControl.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_DAppControl *DAppControlCallerSession) CONTROL() (common.Address, error) {
	return _DAppControl.Contract.CONTROL(&_DAppControl.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_DAppControl *DAppControlCaller) SOURCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "SOURCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_DAppControl *DAppControlSession) SOURCE() (common.Address, error) {
	return _DAppControl.Contract.SOURCE(&_DAppControl.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_DAppControl *DAppControlCallerSession) SOURCE() (common.Address, error) {
	return _DAppControl.Contract.SOURCE(&_DAppControl.CallOpts)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8ed54772.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) view returns(address bidToken)
func (_DAppControl *DAppControlCaller) GetBidFormat(opts *bind.CallOpts, userOp UserOperation) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getBidFormat", userOp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidFormat is a free data retrieval call binding the contract method 0x8ed54772.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) view returns(address bidToken)
func (_DAppControl *DAppControlSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _DAppControl.Contract.GetBidFormat(&_DAppControl.CallOpts, userOp)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8ed54772.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) view returns(address bidToken)
func (_DAppControl *DAppControlCallerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _DAppControl.Contract.GetBidFormat(&_DAppControl.CallOpts, userOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(uint256)
func (_DAppControl *DAppControlCaller) GetBidValue(opts *bind.CallOpts, solverOp SolverOperation) (*big.Int, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getBidValue", solverOp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(uint256)
func (_DAppControl *DAppControlSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _DAppControl.Contract.GetBidValue(&_DAppControl.CallOpts, solverOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(uint256)
func (_DAppControl *DAppControlCallerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _DAppControl.Contract.GetBidValue(&_DAppControl.CallOpts, solverOp)
}

// GetBundlerSurchargeRate is a free data retrieval call binding the contract method 0x228ed68a.
//
// Solidity: function getBundlerSurchargeRate() view returns(uint24)
func (_DAppControl *DAppControlCaller) GetBundlerSurchargeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getBundlerSurchargeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBundlerSurchargeRate is a free data retrieval call binding the contract method 0x228ed68a.
//
// Solidity: function getBundlerSurchargeRate() view returns(uint24)
func (_DAppControl *DAppControlSession) GetBundlerSurchargeRate() (*big.Int, error) {
	return _DAppControl.Contract.GetBundlerSurchargeRate(&_DAppControl.CallOpts)
}

// GetBundlerSurchargeRate is a free data retrieval call binding the contract method 0x228ed68a.
//
// Solidity: function getBundlerSurchargeRate() view returns(uint24)
func (_DAppControl *DAppControlCallerSession) GetBundlerSurchargeRate() (*big.Int, error) {
	return _DAppControl.Contract.GetBundlerSurchargeRate(&_DAppControl.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_DAppControl *DAppControlCaller) GetCallConfig(opts *bind.CallOpts) (CallConfig, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getCallConfig")

	if err != nil {
		return *new(CallConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CallConfig)).(*CallConfig)

	return out0, err

}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_DAppControl *DAppControlSession) GetCallConfig() (CallConfig, error) {
	return _DAppControl.Contract.GetCallConfig(&_DAppControl.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_DAppControl *DAppControlCallerSession) GetCallConfig() (CallConfig, error) {
	return _DAppControl.Contract.GetCallConfig(&_DAppControl.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0xe514a79d.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32,uint32,uint128) dConfig)
func (_DAppControl *DAppControlCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0xe514a79d.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32,uint32,uint128) dConfig)
func (_DAppControl *DAppControlSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _DAppControl.Contract.GetDAppConfig(&_DAppControl.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0xe514a79d.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32,uint32,uint128) dConfig)
func (_DAppControl *DAppControlCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _DAppControl.Contract.GetDAppConfig(&_DAppControl.CallOpts, userOp)
}

// GetDAppGasLimit is a free data retrieval call binding the contract method 0xde25244e.
//
// Solidity: function getDAppGasLimit() view returns(uint32)
func (_DAppControl *DAppControlCaller) GetDAppGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getDAppGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDAppGasLimit is a free data retrieval call binding the contract method 0xde25244e.
//
// Solidity: function getDAppGasLimit() view returns(uint32)
func (_DAppControl *DAppControlSession) GetDAppGasLimit() (uint32, error) {
	return _DAppControl.Contract.GetDAppGasLimit(&_DAppControl.CallOpts)
}

// GetDAppGasLimit is a free data retrieval call binding the contract method 0xde25244e.
//
// Solidity: function getDAppGasLimit() view returns(uint32)
func (_DAppControl *DAppControlCallerSession) GetDAppGasLimit() (uint32, error) {
	return _DAppControl.Contract.GetDAppGasLimit(&_DAppControl.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_DAppControl *DAppControlCaller) GetDAppSignatory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getDAppSignatory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_DAppControl *DAppControlSession) GetDAppSignatory() (common.Address, error) {
	return _DAppControl.Contract.GetDAppSignatory(&_DAppControl.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_DAppControl *DAppControlCallerSession) GetDAppSignatory() (common.Address, error) {
	return _DAppControl.Contract.GetDAppSignatory(&_DAppControl.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_DAppControl *DAppControlCaller) GetSolverGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getSolverGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_DAppControl *DAppControlSession) GetSolverGasLimit() (uint32, error) {
	return _DAppControl.Contract.GetSolverGasLimit(&_DAppControl.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_DAppControl *DAppControlCallerSession) GetSolverGasLimit() (uint32, error) {
	return _DAppControl.Contract.GetSolverGasLimit(&_DAppControl.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_DAppControl *DAppControlCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_DAppControl *DAppControlSession) Governance() (common.Address, error) {
	return _DAppControl.Contract.Governance(&_DAppControl.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_DAppControl *DAppControlCallerSession) Governance() (common.Address, error) {
	return _DAppControl.Contract.Governance(&_DAppControl.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_DAppControl *DAppControlCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_DAppControl *DAppControlSession) PendingGovernance() (common.Address, error) {
	return _DAppControl.Contract.PendingGovernance(&_DAppControl.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_DAppControl *DAppControlCallerSession) PendingGovernance() (common.Address, error) {
	return _DAppControl.Contract.PendingGovernance(&_DAppControl.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_DAppControl *DAppControlCaller) RequireSequentialDAppNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "requireSequentialDAppNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_DAppControl *DAppControlSession) RequireSequentialDAppNonces() (bool, error) {
	return _DAppControl.Contract.RequireSequentialDAppNonces(&_DAppControl.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_DAppControl *DAppControlCallerSession) RequireSequentialDAppNonces() (bool, error) {
	return _DAppControl.Contract.RequireSequentialDAppNonces(&_DAppControl.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_DAppControl *DAppControlCaller) RequireSequentialUserNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "requireSequentialUserNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_DAppControl *DAppControlSession) RequireSequentialUserNonces() (bool, error) {
	return _DAppControl.Contract.RequireSequentialUserNonces(&_DAppControl.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_DAppControl *DAppControlCallerSession) RequireSequentialUserNonces() (bool, error) {
	return _DAppControl.Contract.RequireSequentialUserNonces(&_DAppControl.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_DAppControl *DAppControlCaller) UserDelegated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "userDelegated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_DAppControl *DAppControlSession) UserDelegated() (bool, error) {
	return _DAppControl.Contract.UserDelegated(&_DAppControl.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_DAppControl *DAppControlCallerSession) UserDelegated() (bool, error) {
	return _DAppControl.Contract.UserDelegated(&_DAppControl.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_DAppControl *DAppControlTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_DAppControl *DAppControlSession) AcceptGovernance() (*types.Transaction, error) {
	return _DAppControl.Contract.AcceptGovernance(&_DAppControl.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_DAppControl *DAppControlTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _DAppControl.Contract.AcceptGovernance(&_DAppControl.TransactOpts)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x1e61c250.
//
// Solidity: function allocateValueCall(bool solved, address bidToken, uint256 bidAmount, bytes data) returns()
func (_DAppControl *DAppControlTransactor) AllocateValueCall(opts *bind.TransactOpts, solved bool, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "allocateValueCall", solved, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x1e61c250.
//
// Solidity: function allocateValueCall(bool solved, address bidToken, uint256 bidAmount, bytes data) returns()
func (_DAppControl *DAppControlSession) AllocateValueCall(solved bool, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _DAppControl.Contract.AllocateValueCall(&_DAppControl.TransactOpts, solved, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x1e61c250.
//
// Solidity: function allocateValueCall(bool solved, address bidToken, uint256 bidAmount, bytes data) returns()
func (_DAppControl *DAppControlTransactorSession) AllocateValueCall(solved bool, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _DAppControl.Contract.AllocateValueCall(&_DAppControl.TransactOpts, solved, bidToken, bidAmount, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_DAppControl *DAppControlTransactor) PostSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "postSolverCall", solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_DAppControl *DAppControlSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _DAppControl.Contract.PostSolverCall(&_DAppControl.TransactOpts, solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_DAppControl *DAppControlTransactorSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _DAppControl.Contract.PostSolverCall(&_DAppControl.TransactOpts, solverOp, returnData)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0xdfb90761.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) payable returns(bytes)
func (_DAppControl *DAppControlTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0xdfb90761.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) payable returns(bytes)
func (_DAppControl *DAppControlSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _DAppControl.Contract.PreOpsCall(&_DAppControl.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0xdfb90761.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,uint32,uint24,address,bytes,bytes) userOp) payable returns(bytes)
func (_DAppControl *DAppControlTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _DAppControl.Contract.PreOpsCall(&_DAppControl.TransactOpts, userOp)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_DAppControl *DAppControlTransactor) PreSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "preSolverCall", solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_DAppControl *DAppControlSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _DAppControl.Contract.PreSolverCall(&_DAppControl.TransactOpts, solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_DAppControl *DAppControlTransactorSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _DAppControl.Contract.PreSolverCall(&_DAppControl.TransactOpts, solverOp, returnData)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_DAppControl *DAppControlTransactor) TransferGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "transferGovernance", newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_DAppControl *DAppControlSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _DAppControl.Contract.TransferGovernance(&_DAppControl.TransactOpts, newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_DAppControl *DAppControlTransactorSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _DAppControl.Contract.TransferGovernance(&_DAppControl.TransactOpts, newGovernance)
}

// DAppControlGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the DAppControl contract.
type DAppControlGovernanceTransferStartedIterator struct {
	Event *DAppControlGovernanceTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAppControlGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAppControlGovernanceTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAppControlGovernanceTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAppControlGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAppControlGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAppControlGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the DAppControl contract.
type DAppControlGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_DAppControl *DAppControlFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*DAppControlGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _DAppControl.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &DAppControlGovernanceTransferStartedIterator{contract: _DAppControl.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_DAppControl *DAppControlFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *DAppControlGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _DAppControl.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAppControlGovernanceTransferStarted)
				if err := _DAppControl.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferStarted is a log parse operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_DAppControl *DAppControlFilterer) ParseGovernanceTransferStarted(log types.Log) (*DAppControlGovernanceTransferStarted, error) {
	event := new(DAppControlGovernanceTransferStarted)
	if err := _DAppControl.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAppControlGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the DAppControl contract.
type DAppControlGovernanceTransferredIterator struct {
	Event *DAppControlGovernanceTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAppControlGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAppControlGovernanceTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAppControlGovernanceTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAppControlGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAppControlGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAppControlGovernanceTransferred represents a GovernanceTransferred event raised by the DAppControl contract.
type DAppControlGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_DAppControl *DAppControlFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*DAppControlGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _DAppControl.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &DAppControlGovernanceTransferredIterator{contract: _DAppControl.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_DAppControl *DAppControlFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *DAppControlGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _DAppControl.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAppControlGovernanceTransferred)
				if err := _DAppControl.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_DAppControl *DAppControlFilterer) ParseGovernanceTransferred(log types.Log) (*DAppControlGovernanceTransferred, error) {
	event := new(DAppControlGovernanceTransferred)
	if err := _DAppControl.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
