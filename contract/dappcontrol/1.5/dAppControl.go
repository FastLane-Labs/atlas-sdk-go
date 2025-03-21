// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappcontrol

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
	UserNoncesSequential  bool
	DappNoncesSequential  bool
	RequirePreOps         bool
	TrackPreOpsReturnData bool
	TrackUserReturnData   bool
	DelegateUser          bool
	RequirePreSolver      bool
	RequirePostSolver     bool
	ZeroSolvers           bool
	ReuseUserOp           bool
	UserAuctioneer        bool
	SolverAuctioneer      bool
	UnknownAuctioneer     bool
	VerifyCallChainHash   bool
	ForwardReturnData     bool
	RequireFulfillment    bool
	TrustedOpHash         bool
	InvertBidValue        bool
	ExPostBids            bool
}

// Condition is an auto generated low-level Go binding around an user-defined struct.
type Condition struct {
	Antecedent common.Address
	Context    []byte
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
	DappGasLimit   uint32
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

// SwapData is an auto generated low-level Go binding around an user-defined struct.
type SwapData struct {
	TokenUserBuys       common.Address
	AmountUserBuys      *big.Int
	TokenUserSells      common.Address
	AmountUserSells     *big.Int
	AuctionBaseCurrency common.Address
}

// SwapIntent is an auto generated low-level Go binding around an user-defined struct.
type SwapIntent struct {
	TokenUserBuys       common.Address
	AmountUserBuys      *big.Int
	TokenUserSells      common.Address
	AmountUserSells     *big.Int
	AuctionBaseCurrency common.Address
	Conditions          []Condition
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
	DappGasLimit uint32
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

// DappcontrolMetaData contains all meta data concerning the Dappcontrol contract.
var DappcontrolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlteredControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvertBidValueCannotBeExPostBids\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeDelegatecalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatecall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongPhase\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATLAS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ATLAS_VERIFICATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALL_CONFIG\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTROL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_USER_CONDITIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_CONDITION_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"solved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"allocateValueCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dappGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getBidFormat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"getBidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"userNoncesSequential\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"dappNoncesSequential\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePreOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackPreOpsReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackUserReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"delegateUser\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePreSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePostSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"zeroSolvers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reuseUserOp\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"solverAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"unknownAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"verifyCallChainHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forwardReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireFulfillment\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trustedOpHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"invertBidValue\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exPostBids\",\"type\":\"bool\"}],\"internalType\":\"structCallConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dappGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getDAppConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"solverGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dappGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDAppGasLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDAppSignatory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSolverGasLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"postSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dappGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"preOpsCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"preSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequentialDAppNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequential\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequentialUserNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequential\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"antecedent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"}],\"internalType\":\"structSwapData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"delegated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DappcontrolABI is the input ABI used to generate the binding from.
// Deprecated: Use DappcontrolMetaData.ABI instead.
var DappcontrolABI = DappcontrolMetaData.ABI

// Dappcontrol is an auto generated Go binding around an Ethereum contract.
type Dappcontrol struct {
	DappcontrolCaller     // Read-only binding to the contract
	DappcontrolTransactor // Write-only binding to the contract
	DappcontrolFilterer   // Log filterer for contract events
}

// DappcontrolCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappcontrolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappcontrolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappcontrolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappcontrolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappcontrolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappcontrolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappcontrolSession struct {
	Contract     *Dappcontrol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappcontrolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappcontrolCallerSession struct {
	Contract *DappcontrolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DappcontrolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappcontrolTransactorSession struct {
	Contract     *DappcontrolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DappcontrolRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappcontrolRaw struct {
	Contract *Dappcontrol // Generic contract binding to access the raw methods on
}

// DappcontrolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappcontrolCallerRaw struct {
	Contract *DappcontrolCaller // Generic read-only contract binding to access the raw methods on
}

// DappcontrolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappcontrolTransactorRaw struct {
	Contract *DappcontrolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappcontrol creates a new instance of Dappcontrol, bound to a specific deployed contract.
func NewDappcontrol(address common.Address, backend bind.ContractBackend) (*Dappcontrol, error) {
	contract, err := bindDappcontrol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dappcontrol{DappcontrolCaller: DappcontrolCaller{contract: contract}, DappcontrolTransactor: DappcontrolTransactor{contract: contract}, DappcontrolFilterer: DappcontrolFilterer{contract: contract}}, nil
}

// NewDappcontrolCaller creates a new read-only instance of Dappcontrol, bound to a specific deployed contract.
func NewDappcontrolCaller(address common.Address, caller bind.ContractCaller) (*DappcontrolCaller, error) {
	contract, err := bindDappcontrol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappcontrolCaller{contract: contract}, nil
}

// NewDappcontrolTransactor creates a new write-only instance of Dappcontrol, bound to a specific deployed contract.
func NewDappcontrolTransactor(address common.Address, transactor bind.ContractTransactor) (*DappcontrolTransactor, error) {
	contract, err := bindDappcontrol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappcontrolTransactor{contract: contract}, nil
}

// NewDappcontrolFilterer creates a new log filterer instance of Dappcontrol, bound to a specific deployed contract.
func NewDappcontrolFilterer(address common.Address, filterer bind.ContractFilterer) (*DappcontrolFilterer, error) {
	contract, err := bindDappcontrol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappcontrolFilterer{contract: contract}, nil
}

// bindDappcontrol binds a generic wrapper to an already deployed contract.
func bindDappcontrol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DappcontrolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dappcontrol *DappcontrolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dappcontrol.Contract.DappcontrolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dappcontrol *DappcontrolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dappcontrol.Contract.DappcontrolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dappcontrol *DappcontrolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dappcontrol.Contract.DappcontrolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dappcontrol *DappcontrolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dappcontrol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dappcontrol *DappcontrolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dappcontrol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dappcontrol *DappcontrolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dappcontrol.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_Dappcontrol *DappcontrolCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_Dappcontrol *DappcontrolSession) ATLAS() (common.Address, error) {
	return _Dappcontrol.Contract.ATLAS(&_Dappcontrol.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) ATLAS() (common.Address, error) {
	return _Dappcontrol.Contract.ATLAS(&_Dappcontrol.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_Dappcontrol *DappcontrolCaller) ATLASVERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "ATLAS_VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_Dappcontrol *DappcontrolSession) ATLASVERIFICATION() (common.Address, error) {
	return _Dappcontrol.Contract.ATLASVERIFICATION(&_Dappcontrol.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) ATLASVERIFICATION() (common.Address, error) {
	return _Dappcontrol.Contract.ATLASVERIFICATION(&_Dappcontrol.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_Dappcontrol *DappcontrolCaller) CALLCONFIG(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "CALL_CONFIG")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_Dappcontrol *DappcontrolSession) CALLCONFIG() (uint32, error) {
	return _Dappcontrol.Contract.CALLCONFIG(&_Dappcontrol.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_Dappcontrol *DappcontrolCallerSession) CALLCONFIG() (uint32, error) {
	return _Dappcontrol.Contract.CALLCONFIG(&_Dappcontrol.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_Dappcontrol *DappcontrolCaller) CONTROL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "CONTROL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_Dappcontrol *DappcontrolSession) CONTROL() (common.Address, error) {
	return _Dappcontrol.Contract.CONTROL(&_Dappcontrol.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) CONTROL() (common.Address, error) {
	return _Dappcontrol.Contract.CONTROL(&_Dappcontrol.CallOpts)
}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_Dappcontrol *DappcontrolCaller) MAXUSERCONDITIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "MAX_USER_CONDITIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_Dappcontrol *DappcontrolSession) MAXUSERCONDITIONS() (*big.Int, error) {
	return _Dappcontrol.Contract.MAXUSERCONDITIONS(&_Dappcontrol.CallOpts)
}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_Dappcontrol *DappcontrolCallerSession) MAXUSERCONDITIONS() (*big.Int, error) {
	return _Dappcontrol.Contract.MAXUSERCONDITIONS(&_Dappcontrol.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_Dappcontrol *DappcontrolCaller) SOURCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "SOURCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_Dappcontrol *DappcontrolSession) SOURCE() (common.Address, error) {
	return _Dappcontrol.Contract.SOURCE(&_Dappcontrol.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) SOURCE() (common.Address, error) {
	return _Dappcontrol.Contract.SOURCE(&_Dappcontrol.CallOpts)
}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_Dappcontrol *DappcontrolCaller) USERCONDITIONGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "USER_CONDITION_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_Dappcontrol *DappcontrolSession) USERCONDITIONGASLIMIT() (*big.Int, error) {
	return _Dappcontrol.Contract.USERCONDITIONGASLIMIT(&_Dappcontrol.CallOpts)
}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_Dappcontrol *DappcontrolCallerSession) USERCONDITIONGASLIMIT() (*big.Int, error) {
	return _Dappcontrol.Contract.USERCONDITIONGASLIMIT(&_Dappcontrol.CallOpts)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8fdbb011.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_Dappcontrol *DappcontrolCaller) GetBidFormat(opts *bind.CallOpts, userOp UserOperation) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getBidFormat", userOp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidFormat is a free data retrieval call binding the contract method 0x8fdbb011.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_Dappcontrol *DappcontrolSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _Dappcontrol.Contract.GetBidFormat(&_Dappcontrol.CallOpts, userOp)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8fdbb011.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_Dappcontrol *DappcontrolCallerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _Dappcontrol.Contract.GetBidFormat(&_Dappcontrol.CallOpts, userOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_Dappcontrol *DappcontrolCaller) GetBidValue(opts *bind.CallOpts, solverOp SolverOperation) (*big.Int, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getBidValue", solverOp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_Dappcontrol *DappcontrolSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _Dappcontrol.Contract.GetBidValue(&_Dappcontrol.CallOpts, solverOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_Dappcontrol *DappcontrolCallerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _Dappcontrol.Contract.GetBidValue(&_Dappcontrol.CallOpts, solverOp)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Dappcontrol *DappcontrolCaller) GetCallConfig(opts *bind.CallOpts) (CallConfig, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getCallConfig")

	if err != nil {
		return *new(CallConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CallConfig)).(*CallConfig)

	return out0, err

}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Dappcontrol *DappcontrolSession) GetCallConfig() (CallConfig, error) {
	return _Dappcontrol.Contract.GetCallConfig(&_Dappcontrol.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_Dappcontrol *DappcontrolCallerSession) GetCallConfig() (CallConfig, error) {
	return _Dappcontrol.Contract.GetCallConfig(&_Dappcontrol.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x779be937.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32,uint32) dConfig)
func (_Dappcontrol *DappcontrolCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0x779be937.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32,uint32) dConfig)
func (_Dappcontrol *DappcontrolSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _Dappcontrol.Contract.GetDAppConfig(&_Dappcontrol.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x779be937.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32,uint32) dConfig)
func (_Dappcontrol *DappcontrolCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _Dappcontrol.Contract.GetDAppConfig(&_Dappcontrol.CallOpts, userOp)
}

// GetDAppGasLimit is a free data retrieval call binding the contract method 0xde25244e.
//
// Solidity: function getDAppGasLimit() view returns(uint32)
func (_Dappcontrol *DappcontrolCaller) GetDAppGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getDAppGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDAppGasLimit is a free data retrieval call binding the contract method 0xde25244e.
//
// Solidity: function getDAppGasLimit() view returns(uint32)
func (_Dappcontrol *DappcontrolSession) GetDAppGasLimit() (uint32, error) {
	return _Dappcontrol.Contract.GetDAppGasLimit(&_Dappcontrol.CallOpts)
}

// GetDAppGasLimit is a free data retrieval call binding the contract method 0xde25244e.
//
// Solidity: function getDAppGasLimit() view returns(uint32)
func (_Dappcontrol *DappcontrolCallerSession) GetDAppGasLimit() (uint32, error) {
	return _Dappcontrol.Contract.GetDAppGasLimit(&_Dappcontrol.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_Dappcontrol *DappcontrolCaller) GetDAppSignatory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getDAppSignatory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_Dappcontrol *DappcontrolSession) GetDAppSignatory() (common.Address, error) {
	return _Dappcontrol.Contract.GetDAppSignatory(&_Dappcontrol.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) GetDAppSignatory() (common.Address, error) {
	return _Dappcontrol.Contract.GetDAppSignatory(&_Dappcontrol.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_Dappcontrol *DappcontrolCaller) GetSolverGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "getSolverGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_Dappcontrol *DappcontrolSession) GetSolverGasLimit() (uint32, error) {
	return _Dappcontrol.Contract.GetSolverGasLimit(&_Dappcontrol.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_Dappcontrol *DappcontrolCallerSession) GetSolverGasLimit() (uint32, error) {
	return _Dappcontrol.Contract.GetSolverGasLimit(&_Dappcontrol.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Dappcontrol *DappcontrolCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Dappcontrol *DappcontrolSession) Governance() (common.Address, error) {
	return _Dappcontrol.Contract.Governance(&_Dappcontrol.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) Governance() (common.Address, error) {
	return _Dappcontrol.Contract.Governance(&_Dappcontrol.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Dappcontrol *DappcontrolCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Dappcontrol *DappcontrolSession) PendingGovernance() (common.Address, error) {
	return _Dappcontrol.Contract.PendingGovernance(&_Dappcontrol.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Dappcontrol *DappcontrolCallerSession) PendingGovernance() (common.Address, error) {
	return _Dappcontrol.Contract.PendingGovernance(&_Dappcontrol.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_Dappcontrol *DappcontrolCaller) RequireSequentialDAppNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "requireSequentialDAppNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_Dappcontrol *DappcontrolSession) RequireSequentialDAppNonces() (bool, error) {
	return _Dappcontrol.Contract.RequireSequentialDAppNonces(&_Dappcontrol.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_Dappcontrol *DappcontrolCallerSession) RequireSequentialDAppNonces() (bool, error) {
	return _Dappcontrol.Contract.RequireSequentialDAppNonces(&_Dappcontrol.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_Dappcontrol *DappcontrolCaller) RequireSequentialUserNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "requireSequentialUserNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_Dappcontrol *DappcontrolSession) RequireSequentialUserNonces() (bool, error) {
	return _Dappcontrol.Contract.RequireSequentialUserNonces(&_Dappcontrol.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_Dappcontrol *DappcontrolCallerSession) RequireSequentialUserNonces() (bool, error) {
	return _Dappcontrol.Contract.RequireSequentialUserNonces(&_Dappcontrol.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_Dappcontrol *DappcontrolCaller) UserDelegated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dappcontrol.contract.Call(opts, &out, "userDelegated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_Dappcontrol *DappcontrolSession) UserDelegated() (bool, error) {
	return _Dappcontrol.Contract.UserDelegated(&_Dappcontrol.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_Dappcontrol *DappcontrolCallerSession) UserDelegated() (bool, error) {
	return _Dappcontrol.Contract.UserDelegated(&_Dappcontrol.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Dappcontrol *DappcontrolTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Dappcontrol *DappcontrolSession) AcceptGovernance() (*types.Transaction, error) {
	return _Dappcontrol.Contract.AcceptGovernance(&_Dappcontrol.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Dappcontrol *DappcontrolTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Dappcontrol.Contract.AcceptGovernance(&_Dappcontrol.TransactOpts)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x1e61c250.
//
// Solidity: function allocateValueCall(bool solved, address bidToken, uint256 bidAmount, bytes data) returns()
func (_Dappcontrol *DappcontrolTransactor) AllocateValueCall(opts *bind.TransactOpts, solved bool, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "allocateValueCall", solved, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x1e61c250.
//
// Solidity: function allocateValueCall(bool solved, address bidToken, uint256 bidAmount, bytes data) returns()
func (_Dappcontrol *DappcontrolSession) AllocateValueCall(solved bool, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _Dappcontrol.Contract.AllocateValueCall(&_Dappcontrol.TransactOpts, solved, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x1e61c250.
//
// Solidity: function allocateValueCall(bool solved, address bidToken, uint256 bidAmount, bytes data) returns()
func (_Dappcontrol *DappcontrolTransactorSession) AllocateValueCall(solved bool, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _Dappcontrol.Contract.AllocateValueCall(&_Dappcontrol.TransactOpts, solved, bidToken, bidAmount, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_Dappcontrol *DappcontrolTransactor) PostSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "postSolverCall", solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_Dappcontrol *DappcontrolSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _Dappcontrol.Contract.PostSolverCall(&_Dappcontrol.TransactOpts, solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_Dappcontrol *DappcontrolTransactorSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _Dappcontrol.Contract.PostSolverCall(&_Dappcontrol.TransactOpts, solverOp, returnData)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x4754f399.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_Dappcontrol *DappcontrolTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x4754f399.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_Dappcontrol *DappcontrolSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _Dappcontrol.Contract.PreOpsCall(&_Dappcontrol.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x4754f399.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_Dappcontrol *DappcontrolTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _Dappcontrol.Contract.PreOpsCall(&_Dappcontrol.TransactOpts, userOp)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_Dappcontrol *DappcontrolTransactor) PreSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "preSolverCall", solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_Dappcontrol *DappcontrolSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _Dappcontrol.Contract.PreSolverCall(&_Dappcontrol.TransactOpts, solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_Dappcontrol *DappcontrolTransactorSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _Dappcontrol.Contract.PreSolverCall(&_Dappcontrol.TransactOpts, solverOp, returnData)
}

// Swap is a paid mutator transaction binding the contract method 0x4a9de849.
//
// Solidity: function swap((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_Dappcontrol *DappcontrolTransactor) Swap(opts *bind.TransactOpts, swapIntent SwapIntent) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "swap", swapIntent)
}

// Swap is a paid mutator transaction binding the contract method 0x4a9de849.
//
// Solidity: function swap((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_Dappcontrol *DappcontrolSession) Swap(swapIntent SwapIntent) (*types.Transaction, error) {
	return _Dappcontrol.Contract.Swap(&_Dappcontrol.TransactOpts, swapIntent)
}

// Swap is a paid mutator transaction binding the contract method 0x4a9de849.
//
// Solidity: function swap((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_Dappcontrol *DappcontrolTransactorSession) Swap(swapIntent SwapIntent) (*types.Transaction, error) {
	return _Dappcontrol.Contract.Swap(&_Dappcontrol.TransactOpts, swapIntent)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_Dappcontrol *DappcontrolTransactor) TransferGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _Dappcontrol.contract.Transact(opts, "transferGovernance", newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_Dappcontrol *DappcontrolSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _Dappcontrol.Contract.TransferGovernance(&_Dappcontrol.TransactOpts, newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_Dappcontrol *DappcontrolTransactorSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _Dappcontrol.Contract.TransferGovernance(&_Dappcontrol.TransactOpts, newGovernance)
}

// DappcontrolGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the Dappcontrol contract.
type DappcontrolGovernanceTransferStartedIterator struct {
	Event *DappcontrolGovernanceTransferStarted // Event containing the contract specifics and raw log

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
func (it *DappcontrolGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappcontrolGovernanceTransferStarted)
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
		it.Event = new(DappcontrolGovernanceTransferStarted)
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
func (it *DappcontrolGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappcontrolGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappcontrolGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the Dappcontrol contract.
type DappcontrolGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Dappcontrol *DappcontrolFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*DappcontrolGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Dappcontrol.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &DappcontrolGovernanceTransferStartedIterator{contract: _Dappcontrol.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Dappcontrol *DappcontrolFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *DappcontrolGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Dappcontrol.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappcontrolGovernanceTransferStarted)
				if err := _Dappcontrol.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
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
func (_Dappcontrol *DappcontrolFilterer) ParseGovernanceTransferStarted(log types.Log) (*DappcontrolGovernanceTransferStarted, error) {
	event := new(DappcontrolGovernanceTransferStarted)
	if err := _Dappcontrol.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappcontrolGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the Dappcontrol contract.
type DappcontrolGovernanceTransferredIterator struct {
	Event *DappcontrolGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *DappcontrolGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappcontrolGovernanceTransferred)
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
		it.Event = new(DappcontrolGovernanceTransferred)
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
func (it *DappcontrolGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappcontrolGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappcontrolGovernanceTransferred represents a GovernanceTransferred event raised by the Dappcontrol contract.
type DappcontrolGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Dappcontrol *DappcontrolFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*DappcontrolGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Dappcontrol.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &DappcontrolGovernanceTransferredIterator{contract: _Dappcontrol.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Dappcontrol *DappcontrolFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *DappcontrolGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Dappcontrol.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappcontrolGovernanceTransferred)
				if err := _Dappcontrol.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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
func (_Dappcontrol *DappcontrolFilterer) ParseGovernanceTransferred(log types.Log) (*DappcontrolGovernanceTransferred, error) {
	event := new(DappcontrolGovernanceTransferred)
	if err := _Dappcontrol.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
