// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shmonad

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

// Policy is an auto generated low-level Go binding around an user-defined struct.
type Policy struct {
	EscrowDuration *big.Int
	Active         bool
}

// ShMonadMetaData contains all meta data concerning the ShMonad contract.
var ShMonadMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"sponsoredExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unbondingTask\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addPolicyAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agentBoostYieldFromBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromReleaseAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inUnderlying\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agentExecuteWithSponsor\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"actualPayorCost\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"agentTransferFromBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromReleaseAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inUnderlying\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agentTransferToUnbonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromReleaseAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inUnderlying\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agentWithdrawFromBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromReleaseAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inUnderlying\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfBonded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfUnbonding\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchHold\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchRelease\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bond\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bondRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bondedTotalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"boostYield\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"boostYield\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimAndRebond\",\"inputs\":[{\"name\":\"fromPolicyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toPolicyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bondRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimAndRedeem\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimAsTask\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPolicy\",\"inputs\":[{\"name\":\"escrowDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"policyERC20Wrapper\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndBond\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bondRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shMonToBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disablePolicy\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHoldAmount\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicy\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPolicy\",\"components\":[{\"name\":\"escrowDuration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyAgents\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hold\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPolicyAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"policyBalanceAvailable\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inUnderlying\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"balanceAvailable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"policyCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePolicyAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBondedBalance\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minBonded\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"maxTopUpPerPeriod\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"topUpPeriodDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"topUpAvailable\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inUnderlying\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"amountAvailable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbond\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMinBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"unbondBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondWithTask\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMinBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"unbondBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unbondingCompleteBlock\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddPolicyAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentBoostYieldFromBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentExecuteWithSponsor\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actualPayorCost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentTransferFromBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentTransferToUnbonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentWithdrawFromBonded\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Bond\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreatePolicy\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"escrowDuration\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisablePolicy\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovePolicyAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetTopUp\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"minBonded\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"maxTopUpPerPeriod\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"topUpPeriodDuration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unbond\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expectedUnbondBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AgentSelfUnbondingDisallowed\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxDeposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxMint\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBondedForHold\",\"inputs\":[{\"name\":\"bonded\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holdRequested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[{\"name\":\"bonded\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"unbonding\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"held\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InsufficientNativeTokenSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientUnbondedBalance\",\"inputs\":[{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientUnbondingBalance\",\"inputs\":[{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientUnheldBondedBalance\",\"inputs\":[{\"name\":\"bonded\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"held\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSponsoredExecutorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTaskManagerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MsgDotValueExceedsMsgValueArg\",\"inputs\":[{\"name\":\"msgDotValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValueArg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MsgGasLimitTooLow\",\"inputs\":[{\"name\":\"gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPolicyAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PolicyAgentAlreadyExists\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PolicyAgentNotFound\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"agent\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PolicyInactive\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"PolicyNeedsAtLeastOneAgent\",\"inputs\":[{\"name\":\"policyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TopUpPeriodDurationTooShort\",\"inputs\":[{\"name\":\"requestedPeriodDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minPeriodDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"UnbondingPeriodIncomplete\",\"inputs\":[{\"name\":\"unbondingCompleteBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// ShMonadABI is the input ABI used to generate the binding from.
// Deprecated: Use ShMonadMetaData.ABI instead.
var ShMonadABI = ShMonadMetaData.ABI

// ShMonad is an auto generated Go binding around an Ethereum contract.
type ShMonad struct {
	ShMonadCaller     // Read-only binding to the contract
	ShMonadTransactor // Write-only binding to the contract
	ShMonadFilterer   // Log filterer for contract events
}

// ShMonadCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShMonadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShMonadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShMonadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShMonadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShMonadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShMonadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShMonadSession struct {
	Contract     *ShMonad          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShMonadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShMonadCallerSession struct {
	Contract *ShMonadCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ShMonadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShMonadTransactorSession struct {
	Contract     *ShMonadTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ShMonadRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShMonadRaw struct {
	Contract *ShMonad // Generic contract binding to access the raw methods on
}

// ShMonadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShMonadCallerRaw struct {
	Contract *ShMonadCaller // Generic read-only contract binding to access the raw methods on
}

// ShMonadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShMonadTransactorRaw struct {
	Contract *ShMonadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShMonad creates a new instance of ShMonad, bound to a specific deployed contract.
func NewShMonad(address common.Address, backend bind.ContractBackend) (*ShMonad, error) {
	contract, err := bindShMonad(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShMonad{ShMonadCaller: ShMonadCaller{contract: contract}, ShMonadTransactor: ShMonadTransactor{contract: contract}, ShMonadFilterer: ShMonadFilterer{contract: contract}}, nil
}

// NewShMonadCaller creates a new read-only instance of ShMonad, bound to a specific deployed contract.
func NewShMonadCaller(address common.Address, caller bind.ContractCaller) (*ShMonadCaller, error) {
	contract, err := bindShMonad(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShMonadCaller{contract: contract}, nil
}

// NewShMonadTransactor creates a new write-only instance of ShMonad, bound to a specific deployed contract.
func NewShMonadTransactor(address common.Address, transactor bind.ContractTransactor) (*ShMonadTransactor, error) {
	contract, err := bindShMonad(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShMonadTransactor{contract: contract}, nil
}

// NewShMonadFilterer creates a new log filterer instance of ShMonad, bound to a specific deployed contract.
func NewShMonadFilterer(address common.Address, filterer bind.ContractFilterer) (*ShMonadFilterer, error) {
	contract, err := bindShMonad(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShMonadFilterer{contract: contract}, nil
}

// bindShMonad binds a generic wrapper to an already deployed contract.
func bindShMonad(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShMonadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShMonad *ShMonadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShMonad.Contract.ShMonadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShMonad *ShMonadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShMonad.Contract.ShMonadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShMonad *ShMonadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShMonad.Contract.ShMonadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShMonad *ShMonadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShMonad.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShMonad *ShMonadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShMonad.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShMonad *ShMonadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShMonad.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ShMonad *ShMonadCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ShMonad *ShMonadSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ShMonad.Contract.DOMAINSEPARATOR(&_ShMonad.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ShMonad *ShMonadCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ShMonad.Contract.DOMAINSEPARATOR(&_ShMonad.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShMonad *ShMonadCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShMonad *ShMonadSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ShMonad.Contract.Allowance(&_ShMonad.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ShMonad.Contract.Allowance(&_ShMonad.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ShMonad *ShMonadCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ShMonad *ShMonadSession) Asset() (common.Address, error) {
	return _ShMonad.Contract.Asset(&_ShMonad.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ShMonad *ShMonadCallerSession) Asset() (common.Address, error) {
	return _ShMonad.Contract.Asset(&_ShMonad.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShMonad *ShMonadCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShMonad *ShMonadSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOf(&_ShMonad.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOf(&_ShMonad.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_ShMonad *ShMonadCaller) BalanceOfBonded(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "balanceOfBonded", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_ShMonad *ShMonadSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOfBonded(&_ShMonad.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOfBonded(&_ShMonad.CallOpts, account)
}

// BalanceOfBonded0 is a free data retrieval call binding the contract method 0xdc6d2a7e.
//
// Solidity: function balanceOfBonded(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCaller) BalanceOfBonded0(opts *bind.CallOpts, policyID uint64, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "balanceOfBonded0", policyID, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBonded0 is a free data retrieval call binding the contract method 0xdc6d2a7e.
//
// Solidity: function balanceOfBonded(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadSession) BalanceOfBonded0(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOfBonded0(&_ShMonad.CallOpts, policyID, account)
}

// BalanceOfBonded0 is a free data retrieval call binding the contract method 0xdc6d2a7e.
//
// Solidity: function balanceOfBonded(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) BalanceOfBonded0(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOfBonded0(&_ShMonad.CallOpts, policyID, account)
}

// BalanceOfUnbonding is a free data retrieval call binding the contract method 0x51d3c81a.
//
// Solidity: function balanceOfUnbonding(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCaller) BalanceOfUnbonding(opts *bind.CallOpts, policyID uint64, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "balanceOfUnbonding", policyID, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnbonding is a free data retrieval call binding the contract method 0x51d3c81a.
//
// Solidity: function balanceOfUnbonding(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadSession) BalanceOfUnbonding(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOfUnbonding(&_ShMonad.CallOpts, policyID, account)
}

// BalanceOfUnbonding is a free data retrieval call binding the contract method 0x51d3c81a.
//
// Solidity: function balanceOfUnbonding(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) BalanceOfUnbonding(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.BalanceOfUnbonding(&_ShMonad.CallOpts, policyID, account)
}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_ShMonad *ShMonadCaller) BondedTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "bondedTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_ShMonad *ShMonadSession) BondedTotalSupply() (*big.Int, error) {
	return _ShMonad.Contract.BondedTotalSupply(&_ShMonad.CallOpts)
}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_ShMonad *ShMonadCallerSession) BondedTotalSupply() (*big.Int, error) {
	return _ShMonad.Contract.BondedTotalSupply(&_ShMonad.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.ConvertToAssets(&_ShMonad.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.ConvertToAssets(&_ShMonad.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.ConvertToShares(&_ShMonad.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.ConvertToShares(&_ShMonad.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShMonad *ShMonadCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShMonad *ShMonadSession) Decimals() (uint8, error) {
	return _ShMonad.Contract.Decimals(&_ShMonad.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShMonad *ShMonadCallerSession) Decimals() (uint8, error) {
	return _ShMonad.Contract.Decimals(&_ShMonad.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ShMonad *ShMonadCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ShMonad *ShMonadSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ShMonad.Contract.Eip712Domain(&_ShMonad.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ShMonad *ShMonadCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ShMonad.Contract.Eip712Domain(&_ShMonad.CallOpts)
}

// GetHoldAmount is a free data retrieval call binding the contract method 0xda92ea0d.
//
// Solidity: function getHoldAmount(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCaller) GetHoldAmount(opts *bind.CallOpts, policyID uint64, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "getHoldAmount", policyID, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHoldAmount is a free data retrieval call binding the contract method 0xda92ea0d.
//
// Solidity: function getHoldAmount(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadSession) GetHoldAmount(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.GetHoldAmount(&_ShMonad.CallOpts, policyID, account)
}

// GetHoldAmount is a free data retrieval call binding the contract method 0xda92ea0d.
//
// Solidity: function getHoldAmount(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) GetHoldAmount(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.GetHoldAmount(&_ShMonad.CallOpts, policyID, account)
}

// GetPolicy is a free data retrieval call binding the contract method 0x6d738773.
//
// Solidity: function getPolicy(uint64 policyID) view returns((uint48,bool))
func (_ShMonad *ShMonadCaller) GetPolicy(opts *bind.CallOpts, policyID uint64) (Policy, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "getPolicy", policyID)

	if err != nil {
		return *new(Policy), err
	}

	out0 := *abi.ConvertType(out[0], new(Policy)).(*Policy)

	return out0, err

}

// GetPolicy is a free data retrieval call binding the contract method 0x6d738773.
//
// Solidity: function getPolicy(uint64 policyID) view returns((uint48,bool))
func (_ShMonad *ShMonadSession) GetPolicy(policyID uint64) (Policy, error) {
	return _ShMonad.Contract.GetPolicy(&_ShMonad.CallOpts, policyID)
}

// GetPolicy is a free data retrieval call binding the contract method 0x6d738773.
//
// Solidity: function getPolicy(uint64 policyID) view returns((uint48,bool))
func (_ShMonad *ShMonadCallerSession) GetPolicy(policyID uint64) (Policy, error) {
	return _ShMonad.Contract.GetPolicy(&_ShMonad.CallOpts, policyID)
}

// GetPolicyAgents is a free data retrieval call binding the contract method 0x6653bd41.
//
// Solidity: function getPolicyAgents(uint64 policyID) view returns(address[])
func (_ShMonad *ShMonadCaller) GetPolicyAgents(opts *bind.CallOpts, policyID uint64) ([]common.Address, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "getPolicyAgents", policyID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPolicyAgents is a free data retrieval call binding the contract method 0x6653bd41.
//
// Solidity: function getPolicyAgents(uint64 policyID) view returns(address[])
func (_ShMonad *ShMonadSession) GetPolicyAgents(policyID uint64) ([]common.Address, error) {
	return _ShMonad.Contract.GetPolicyAgents(&_ShMonad.CallOpts, policyID)
}

// GetPolicyAgents is a free data retrieval call binding the contract method 0x6653bd41.
//
// Solidity: function getPolicyAgents(uint64 policyID) view returns(address[])
func (_ShMonad *ShMonadCallerSession) GetPolicyAgents(policyID uint64) ([]common.Address, error) {
	return _ShMonad.Contract.GetPolicyAgents(&_ShMonad.CallOpts, policyID)
}

// IsPolicyAgent is a free data retrieval call binding the contract method 0x6d20d833.
//
// Solidity: function isPolicyAgent(uint64 policyID, address agent) view returns(bool)
func (_ShMonad *ShMonadCaller) IsPolicyAgent(opts *bind.CallOpts, policyID uint64, agent common.Address) (bool, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "isPolicyAgent", policyID, agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPolicyAgent is a free data retrieval call binding the contract method 0x6d20d833.
//
// Solidity: function isPolicyAgent(uint64 policyID, address agent) view returns(bool)
func (_ShMonad *ShMonadSession) IsPolicyAgent(policyID uint64, agent common.Address) (bool, error) {
	return _ShMonad.Contract.IsPolicyAgent(&_ShMonad.CallOpts, policyID, agent)
}

// IsPolicyAgent is a free data retrieval call binding the contract method 0x6d20d833.
//
// Solidity: function isPolicyAgent(uint64 policyID, address agent) view returns(bool)
func (_ShMonad *ShMonadCallerSession) IsPolicyAgent(policyID uint64, agent common.Address) (bool, error) {
	return _ShMonad.Contract.IsPolicyAgent(&_ShMonad.CallOpts, policyID, agent)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ShMonad *ShMonadCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ShMonad *ShMonadSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxDeposit(&_ShMonad.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxDeposit(&_ShMonad.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ShMonad *ShMonadCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ShMonad *ShMonadSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxMint(&_ShMonad.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxMint(&_ShMonad.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ShMonad *ShMonadCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ShMonad *ShMonadSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxRedeem(&_ShMonad.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxRedeem(&_ShMonad.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ShMonad *ShMonadCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ShMonad *ShMonadSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxWithdraw(&_ShMonad.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ShMonad.Contract.MaxWithdraw(&_ShMonad.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShMonad *ShMonadCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShMonad *ShMonadSession) Name() (string, error) {
	return _ShMonad.Contract.Name(&_ShMonad.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShMonad *ShMonadCallerSession) Name() (string, error) {
	return _ShMonad.Contract.Name(&_ShMonad.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ShMonad *ShMonadCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ShMonad *ShMonadSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ShMonad.Contract.Nonces(&_ShMonad.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ShMonad.Contract.Nonces(&_ShMonad.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShMonad *ShMonadCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShMonad *ShMonadSession) Owner() (common.Address, error) {
	return _ShMonad.Contract.Owner(&_ShMonad.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShMonad *ShMonadCallerSession) Owner() (common.Address, error) {
	return _ShMonad.Contract.Owner(&_ShMonad.CallOpts)
}

// PolicyBalanceAvailable is a free data retrieval call binding the contract method 0x7d05b18e.
//
// Solidity: function policyBalanceAvailable(uint64 policyID, address account, bool inUnderlying) view returns(uint256 balanceAvailable)
func (_ShMonad *ShMonadCaller) PolicyBalanceAvailable(opts *bind.CallOpts, policyID uint64, account common.Address, inUnderlying bool) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "policyBalanceAvailable", policyID, account, inUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PolicyBalanceAvailable is a free data retrieval call binding the contract method 0x7d05b18e.
//
// Solidity: function policyBalanceAvailable(uint64 policyID, address account, bool inUnderlying) view returns(uint256 balanceAvailable)
func (_ShMonad *ShMonadSession) PolicyBalanceAvailable(policyID uint64, account common.Address, inUnderlying bool) (*big.Int, error) {
	return _ShMonad.Contract.PolicyBalanceAvailable(&_ShMonad.CallOpts, policyID, account, inUnderlying)
}

// PolicyBalanceAvailable is a free data retrieval call binding the contract method 0x7d05b18e.
//
// Solidity: function policyBalanceAvailable(uint64 policyID, address account, bool inUnderlying) view returns(uint256 balanceAvailable)
func (_ShMonad *ShMonadCallerSession) PolicyBalanceAvailable(policyID uint64, account common.Address, inUnderlying bool) (*big.Int, error) {
	return _ShMonad.Contract.PolicyBalanceAvailable(&_ShMonad.CallOpts, policyID, account, inUnderlying)
}

// PolicyCount is a free data retrieval call binding the contract method 0xde54d429.
//
// Solidity: function policyCount() view returns(uint64)
func (_ShMonad *ShMonadCaller) PolicyCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "policyCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PolicyCount is a free data retrieval call binding the contract method 0xde54d429.
//
// Solidity: function policyCount() view returns(uint64)
func (_ShMonad *ShMonadSession) PolicyCount() (uint64, error) {
	return _ShMonad.Contract.PolicyCount(&_ShMonad.CallOpts)
}

// PolicyCount is a free data retrieval call binding the contract method 0xde54d429.
//
// Solidity: function policyCount() view returns(uint64)
func (_ShMonad *ShMonadCallerSession) PolicyCount() (uint64, error) {
	return _ShMonad.Contract.PolicyCount(&_ShMonad.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewDeposit(&_ShMonad.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewDeposit(&_ShMonad.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewMint(&_ShMonad.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewMint(&_ShMonad.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewRedeem(&_ShMonad.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewRedeem(&_ShMonad.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewWithdraw(&_ShMonad.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ShMonad.Contract.PreviewWithdraw(&_ShMonad.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShMonad *ShMonadCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShMonad *ShMonadSession) Symbol() (string, error) {
	return _ShMonad.Contract.Symbol(&_ShMonad.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShMonad *ShMonadCallerSession) Symbol() (string, error) {
	return _ShMonad.Contract.Symbol(&_ShMonad.CallOpts)
}

// TopUpAvailable is a free data retrieval call binding the contract method 0x97c76115.
//
// Solidity: function topUpAvailable(uint64 policyID, address account, bool inUnderlying) view returns(uint256 amountAvailable)
func (_ShMonad *ShMonadCaller) TopUpAvailable(opts *bind.CallOpts, policyID uint64, account common.Address, inUnderlying bool) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "topUpAvailable", policyID, account, inUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TopUpAvailable is a free data retrieval call binding the contract method 0x97c76115.
//
// Solidity: function topUpAvailable(uint64 policyID, address account, bool inUnderlying) view returns(uint256 amountAvailable)
func (_ShMonad *ShMonadSession) TopUpAvailable(policyID uint64, account common.Address, inUnderlying bool) (*big.Int, error) {
	return _ShMonad.Contract.TopUpAvailable(&_ShMonad.CallOpts, policyID, account, inUnderlying)
}

// TopUpAvailable is a free data retrieval call binding the contract method 0x97c76115.
//
// Solidity: function topUpAvailable(uint64 policyID, address account, bool inUnderlying) view returns(uint256 amountAvailable)
func (_ShMonad *ShMonadCallerSession) TopUpAvailable(policyID uint64, account common.Address, inUnderlying bool) (*big.Int, error) {
	return _ShMonad.Contract.TopUpAvailable(&_ShMonad.CallOpts, policyID, account, inUnderlying)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ShMonad *ShMonadCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ShMonad *ShMonadSession) TotalAssets() (*big.Int, error) {
	return _ShMonad.Contract.TotalAssets(&_ShMonad.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ShMonad *ShMonadCallerSession) TotalAssets() (*big.Int, error) {
	return _ShMonad.Contract.TotalAssets(&_ShMonad.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShMonad *ShMonadCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShMonad *ShMonadSession) TotalSupply() (*big.Int, error) {
	return _ShMonad.Contract.TotalSupply(&_ShMonad.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShMonad *ShMonadCallerSession) TotalSupply() (*big.Int, error) {
	return _ShMonad.Contract.TotalSupply(&_ShMonad.CallOpts)
}

// UnbondingCompleteBlock is a free data retrieval call binding the contract method 0xd9306102.
//
// Solidity: function unbondingCompleteBlock(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCaller) UnbondingCompleteBlock(opts *bind.CallOpts, policyID uint64, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShMonad.contract.Call(opts, &out, "unbondingCompleteBlock", policyID, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingCompleteBlock is a free data retrieval call binding the contract method 0xd9306102.
//
// Solidity: function unbondingCompleteBlock(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadSession) UnbondingCompleteBlock(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.UnbondingCompleteBlock(&_ShMonad.CallOpts, policyID, account)
}

// UnbondingCompleteBlock is a free data retrieval call binding the contract method 0xd9306102.
//
// Solidity: function unbondingCompleteBlock(uint64 policyID, address account) view returns(uint256)
func (_ShMonad *ShMonadCallerSession) UnbondingCompleteBlock(policyID uint64, account common.Address) (*big.Int, error) {
	return _ShMonad.Contract.UnbondingCompleteBlock(&_ShMonad.CallOpts, policyID, account)
}

// AddPolicyAgent is a paid mutator transaction binding the contract method 0x462fff96.
//
// Solidity: function addPolicyAgent(uint64 policyID, address agent) returns()
func (_ShMonad *ShMonadTransactor) AddPolicyAgent(opts *bind.TransactOpts, policyID uint64, agent common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "addPolicyAgent", policyID, agent)
}

// AddPolicyAgent is a paid mutator transaction binding the contract method 0x462fff96.
//
// Solidity: function addPolicyAgent(uint64 policyID, address agent) returns()
func (_ShMonad *ShMonadSession) AddPolicyAgent(policyID uint64, agent common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.AddPolicyAgent(&_ShMonad.TransactOpts, policyID, agent)
}

// AddPolicyAgent is a paid mutator transaction binding the contract method 0x462fff96.
//
// Solidity: function addPolicyAgent(uint64 policyID, address agent) returns()
func (_ShMonad *ShMonadTransactorSession) AddPolicyAgent(policyID uint64, agent common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.AddPolicyAgent(&_ShMonad.TransactOpts, policyID, agent)
}

// AgentBoostYieldFromBonded is a paid mutator transaction binding the contract method 0x882b9318.
//
// Solidity: function agentBoostYieldFromBonded(uint64 policyID, address from, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactor) AgentBoostYieldFromBonded(opts *bind.TransactOpts, policyID uint64, from common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "agentBoostYieldFromBonded", policyID, from, amount, fromReleaseAmount, inUnderlying)
}

// AgentBoostYieldFromBonded is a paid mutator transaction binding the contract method 0x882b9318.
//
// Solidity: function agentBoostYieldFromBonded(uint64 policyID, address from, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadSession) AgentBoostYieldFromBonded(policyID uint64, from common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentBoostYieldFromBonded(&_ShMonad.TransactOpts, policyID, from, amount, fromReleaseAmount, inUnderlying)
}

// AgentBoostYieldFromBonded is a paid mutator transaction binding the contract method 0x882b9318.
//
// Solidity: function agentBoostYieldFromBonded(uint64 policyID, address from, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactorSession) AgentBoostYieldFromBonded(policyID uint64, from common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentBoostYieldFromBonded(&_ShMonad.TransactOpts, policyID, from, amount, fromReleaseAmount, inUnderlying)
}

// AgentExecuteWithSponsor is a paid mutator transaction binding the contract method 0xe4cf2418.
//
// Solidity: function agentExecuteWithSponsor(uint64 policyID, address payor, address recipient, uint256 msgValue, uint256 gasLimit, address callTarget, bytes callData) payable returns(uint128 actualPayorCost, bool success, bytes returnData)
func (_ShMonad *ShMonadTransactor) AgentExecuteWithSponsor(opts *bind.TransactOpts, policyID uint64, payor common.Address, recipient common.Address, msgValue *big.Int, gasLimit *big.Int, callTarget common.Address, callData []byte) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "agentExecuteWithSponsor", policyID, payor, recipient, msgValue, gasLimit, callTarget, callData)
}

// AgentExecuteWithSponsor is a paid mutator transaction binding the contract method 0xe4cf2418.
//
// Solidity: function agentExecuteWithSponsor(uint64 policyID, address payor, address recipient, uint256 msgValue, uint256 gasLimit, address callTarget, bytes callData) payable returns(uint128 actualPayorCost, bool success, bytes returnData)
func (_ShMonad *ShMonadSession) AgentExecuteWithSponsor(policyID uint64, payor common.Address, recipient common.Address, msgValue *big.Int, gasLimit *big.Int, callTarget common.Address, callData []byte) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentExecuteWithSponsor(&_ShMonad.TransactOpts, policyID, payor, recipient, msgValue, gasLimit, callTarget, callData)
}

// AgentExecuteWithSponsor is a paid mutator transaction binding the contract method 0xe4cf2418.
//
// Solidity: function agentExecuteWithSponsor(uint64 policyID, address payor, address recipient, uint256 msgValue, uint256 gasLimit, address callTarget, bytes callData) payable returns(uint128 actualPayorCost, bool success, bytes returnData)
func (_ShMonad *ShMonadTransactorSession) AgentExecuteWithSponsor(policyID uint64, payor common.Address, recipient common.Address, msgValue *big.Int, gasLimit *big.Int, callTarget common.Address, callData []byte) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentExecuteWithSponsor(&_ShMonad.TransactOpts, policyID, payor, recipient, msgValue, gasLimit, callTarget, callData)
}

// AgentTransferFromBonded is a paid mutator transaction binding the contract method 0xb1f3a9a8.
//
// Solidity: function agentTransferFromBonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactor) AgentTransferFromBonded(opts *bind.TransactOpts, policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "agentTransferFromBonded", policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentTransferFromBonded is a paid mutator transaction binding the contract method 0xb1f3a9a8.
//
// Solidity: function agentTransferFromBonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadSession) AgentTransferFromBonded(policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentTransferFromBonded(&_ShMonad.TransactOpts, policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentTransferFromBonded is a paid mutator transaction binding the contract method 0xb1f3a9a8.
//
// Solidity: function agentTransferFromBonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactorSession) AgentTransferFromBonded(policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentTransferFromBonded(&_ShMonad.TransactOpts, policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentTransferToUnbonded is a paid mutator transaction binding the contract method 0x70ca267f.
//
// Solidity: function agentTransferToUnbonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactor) AgentTransferToUnbonded(opts *bind.TransactOpts, policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "agentTransferToUnbonded", policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentTransferToUnbonded is a paid mutator transaction binding the contract method 0x70ca267f.
//
// Solidity: function agentTransferToUnbonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadSession) AgentTransferToUnbonded(policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentTransferToUnbonded(&_ShMonad.TransactOpts, policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentTransferToUnbonded is a paid mutator transaction binding the contract method 0x70ca267f.
//
// Solidity: function agentTransferToUnbonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactorSession) AgentTransferToUnbonded(policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentTransferToUnbonded(&_ShMonad.TransactOpts, policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentWithdrawFromBonded is a paid mutator transaction binding the contract method 0x87084d17.
//
// Solidity: function agentWithdrawFromBonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactor) AgentWithdrawFromBonded(opts *bind.TransactOpts, policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "agentWithdrawFromBonded", policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentWithdrawFromBonded is a paid mutator transaction binding the contract method 0x87084d17.
//
// Solidity: function agentWithdrawFromBonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadSession) AgentWithdrawFromBonded(policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentWithdrawFromBonded(&_ShMonad.TransactOpts, policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// AgentWithdrawFromBonded is a paid mutator transaction binding the contract method 0x87084d17.
//
// Solidity: function agentWithdrawFromBonded(uint64 policyID, address from, address to, uint256 amount, uint256 fromReleaseAmount, bool inUnderlying) returns()
func (_ShMonad *ShMonadTransactorSession) AgentWithdrawFromBonded(policyID uint64, from common.Address, to common.Address, amount *big.Int, fromReleaseAmount *big.Int, inUnderlying bool) (*types.Transaction, error) {
	return _ShMonad.Contract.AgentWithdrawFromBonded(&_ShMonad.TransactOpts, policyID, from, to, amount, fromReleaseAmount, inUnderlying)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ShMonad *ShMonadTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ShMonad *ShMonadSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Approve(&_ShMonad.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ShMonad *ShMonadTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Approve(&_ShMonad.TransactOpts, spender, value)
}

// BatchHold is a paid mutator transaction binding the contract method 0xba7c0e40.
//
// Solidity: function batchHold(uint64 policyID, address[] accounts, uint256[] amounts) returns()
func (_ShMonad *ShMonadTransactor) BatchHold(opts *bind.TransactOpts, policyID uint64, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "batchHold", policyID, accounts, amounts)
}

// BatchHold is a paid mutator transaction binding the contract method 0xba7c0e40.
//
// Solidity: function batchHold(uint64 policyID, address[] accounts, uint256[] amounts) returns()
func (_ShMonad *ShMonadSession) BatchHold(policyID uint64, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.BatchHold(&_ShMonad.TransactOpts, policyID, accounts, amounts)
}

// BatchHold is a paid mutator transaction binding the contract method 0xba7c0e40.
//
// Solidity: function batchHold(uint64 policyID, address[] accounts, uint256[] amounts) returns()
func (_ShMonad *ShMonadTransactorSession) BatchHold(policyID uint64, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.BatchHold(&_ShMonad.TransactOpts, policyID, accounts, amounts)
}

// BatchRelease is a paid mutator transaction binding the contract method 0x9292430b.
//
// Solidity: function batchRelease(uint64 policyID, address[] accounts, uint256[] amounts) returns()
func (_ShMonad *ShMonadTransactor) BatchRelease(opts *bind.TransactOpts, policyID uint64, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "batchRelease", policyID, accounts, amounts)
}

// BatchRelease is a paid mutator transaction binding the contract method 0x9292430b.
//
// Solidity: function batchRelease(uint64 policyID, address[] accounts, uint256[] amounts) returns()
func (_ShMonad *ShMonadSession) BatchRelease(policyID uint64, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.BatchRelease(&_ShMonad.TransactOpts, policyID, accounts, amounts)
}

// BatchRelease is a paid mutator transaction binding the contract method 0x9292430b.
//
// Solidity: function batchRelease(uint64 policyID, address[] accounts, uint256[] amounts) returns()
func (_ShMonad *ShMonadTransactorSession) BatchRelease(policyID uint64, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.BatchRelease(&_ShMonad.TransactOpts, policyID, accounts, amounts)
}

// Bond is a paid mutator transaction binding the contract method 0xf6201239.
//
// Solidity: function bond(uint64 policyID, address bondRecipient, uint256 shares) returns()
func (_ShMonad *ShMonadTransactor) Bond(opts *bind.TransactOpts, policyID uint64, bondRecipient common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "bond", policyID, bondRecipient, shares)
}

// Bond is a paid mutator transaction binding the contract method 0xf6201239.
//
// Solidity: function bond(uint64 policyID, address bondRecipient, uint256 shares) returns()
func (_ShMonad *ShMonadSession) Bond(policyID uint64, bondRecipient common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Bond(&_ShMonad.TransactOpts, policyID, bondRecipient, shares)
}

// Bond is a paid mutator transaction binding the contract method 0xf6201239.
//
// Solidity: function bond(uint64 policyID, address bondRecipient, uint256 shares) returns()
func (_ShMonad *ShMonadTransactorSession) Bond(policyID uint64, bondRecipient common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Bond(&_ShMonad.TransactOpts, policyID, bondRecipient, shares)
}

// BoostYield is a paid mutator transaction binding the contract method 0x2eac4115.
//
// Solidity: function boostYield(uint256 shares, address from) returns()
func (_ShMonad *ShMonadTransactor) BoostYield(opts *bind.TransactOpts, shares *big.Int, from common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "boostYield", shares, from)
}

// BoostYield is a paid mutator transaction binding the contract method 0x2eac4115.
//
// Solidity: function boostYield(uint256 shares, address from) returns()
func (_ShMonad *ShMonadSession) BoostYield(shares *big.Int, from common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.BoostYield(&_ShMonad.TransactOpts, shares, from)
}

// BoostYield is a paid mutator transaction binding the contract method 0x2eac4115.
//
// Solidity: function boostYield(uint256 shares, address from) returns()
func (_ShMonad *ShMonadTransactorSession) BoostYield(shares *big.Int, from common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.BoostYield(&_ShMonad.TransactOpts, shares, from)
}

// BoostYield0 is a paid mutator transaction binding the contract method 0xfb2c2845.
//
// Solidity: function boostYield() payable returns()
func (_ShMonad *ShMonadTransactor) BoostYield0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "boostYield0")
}

// BoostYield0 is a paid mutator transaction binding the contract method 0xfb2c2845.
//
// Solidity: function boostYield() payable returns()
func (_ShMonad *ShMonadSession) BoostYield0() (*types.Transaction, error) {
	return _ShMonad.Contract.BoostYield0(&_ShMonad.TransactOpts)
}

// BoostYield0 is a paid mutator transaction binding the contract method 0xfb2c2845.
//
// Solidity: function boostYield() payable returns()
func (_ShMonad *ShMonadTransactorSession) BoostYield0() (*types.Transaction, error) {
	return _ShMonad.Contract.BoostYield0(&_ShMonad.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x503914db.
//
// Solidity: function claim(uint64 policyID, uint256 shares) returns()
func (_ShMonad *ShMonadTransactor) Claim(opts *bind.TransactOpts, policyID uint64, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "claim", policyID, shares)
}

// Claim is a paid mutator transaction binding the contract method 0x503914db.
//
// Solidity: function claim(uint64 policyID, uint256 shares) returns()
func (_ShMonad *ShMonadSession) Claim(policyID uint64, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Claim(&_ShMonad.TransactOpts, policyID, shares)
}

// Claim is a paid mutator transaction binding the contract method 0x503914db.
//
// Solidity: function claim(uint64 policyID, uint256 shares) returns()
func (_ShMonad *ShMonadTransactorSession) Claim(policyID uint64, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Claim(&_ShMonad.TransactOpts, policyID, shares)
}

// ClaimAndRebond is a paid mutator transaction binding the contract method 0x947e05b7.
//
// Solidity: function claimAndRebond(uint64 fromPolicyID, uint64 toPolicyID, address bondRecipient, uint256 shares) returns()
func (_ShMonad *ShMonadTransactor) ClaimAndRebond(opts *bind.TransactOpts, fromPolicyID uint64, toPolicyID uint64, bondRecipient common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "claimAndRebond", fromPolicyID, toPolicyID, bondRecipient, shares)
}

// ClaimAndRebond is a paid mutator transaction binding the contract method 0x947e05b7.
//
// Solidity: function claimAndRebond(uint64 fromPolicyID, uint64 toPolicyID, address bondRecipient, uint256 shares) returns()
func (_ShMonad *ShMonadSession) ClaimAndRebond(fromPolicyID uint64, toPolicyID uint64, bondRecipient common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.ClaimAndRebond(&_ShMonad.TransactOpts, fromPolicyID, toPolicyID, bondRecipient, shares)
}

// ClaimAndRebond is a paid mutator transaction binding the contract method 0x947e05b7.
//
// Solidity: function claimAndRebond(uint64 fromPolicyID, uint64 toPolicyID, address bondRecipient, uint256 shares) returns()
func (_ShMonad *ShMonadTransactorSession) ClaimAndRebond(fromPolicyID uint64, toPolicyID uint64, bondRecipient common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.ClaimAndRebond(&_ShMonad.TransactOpts, fromPolicyID, toPolicyID, bondRecipient, shares)
}

// ClaimAndRedeem is a paid mutator transaction binding the contract method 0x09805805.
//
// Solidity: function claimAndRedeem(uint64 policyID, uint256 shares) returns(uint256 assets)
func (_ShMonad *ShMonadTransactor) ClaimAndRedeem(opts *bind.TransactOpts, policyID uint64, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "claimAndRedeem", policyID, shares)
}

// ClaimAndRedeem is a paid mutator transaction binding the contract method 0x09805805.
//
// Solidity: function claimAndRedeem(uint64 policyID, uint256 shares) returns(uint256 assets)
func (_ShMonad *ShMonadSession) ClaimAndRedeem(policyID uint64, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.ClaimAndRedeem(&_ShMonad.TransactOpts, policyID, shares)
}

// ClaimAndRedeem is a paid mutator transaction binding the contract method 0x09805805.
//
// Solidity: function claimAndRedeem(uint64 policyID, uint256 shares) returns(uint256 assets)
func (_ShMonad *ShMonadTransactorSession) ClaimAndRedeem(policyID uint64, shares *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.ClaimAndRedeem(&_ShMonad.TransactOpts, policyID, shares)
}

// ClaimAsTask is a paid mutator transaction binding the contract method 0xef3bdd25.
//
// Solidity: function claimAsTask(uint64 policyID, uint256 shares, address account) returns()
func (_ShMonad *ShMonadTransactor) ClaimAsTask(opts *bind.TransactOpts, policyID uint64, shares *big.Int, account common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "claimAsTask", policyID, shares, account)
}

// ClaimAsTask is a paid mutator transaction binding the contract method 0xef3bdd25.
//
// Solidity: function claimAsTask(uint64 policyID, uint256 shares, address account) returns()
func (_ShMonad *ShMonadSession) ClaimAsTask(policyID uint64, shares *big.Int, account common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.ClaimAsTask(&_ShMonad.TransactOpts, policyID, shares, account)
}

// ClaimAsTask is a paid mutator transaction binding the contract method 0xef3bdd25.
//
// Solidity: function claimAsTask(uint64 policyID, uint256 shares, address account) returns()
func (_ShMonad *ShMonadTransactorSession) ClaimAsTask(policyID uint64, shares *big.Int, account common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.ClaimAsTask(&_ShMonad.TransactOpts, policyID, shares, account)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x321677f2.
//
// Solidity: function createPolicy(uint48 escrowDuration) returns(uint64 policyID, address policyERC20Wrapper)
func (_ShMonad *ShMonadTransactor) CreatePolicy(opts *bind.TransactOpts, escrowDuration *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "createPolicy", escrowDuration)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x321677f2.
//
// Solidity: function createPolicy(uint48 escrowDuration) returns(uint64 policyID, address policyERC20Wrapper)
func (_ShMonad *ShMonadSession) CreatePolicy(escrowDuration *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.CreatePolicy(&_ShMonad.TransactOpts, escrowDuration)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x321677f2.
//
// Solidity: function createPolicy(uint48 escrowDuration) returns(uint64 policyID, address policyERC20Wrapper)
func (_ShMonad *ShMonadTransactorSession) CreatePolicy(escrowDuration *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.CreatePolicy(&_ShMonad.TransactOpts, escrowDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) payable returns(uint256)
func (_ShMonad *ShMonadTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) payable returns(uint256)
func (_ShMonad *ShMonadSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Deposit(&_ShMonad.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) payable returns(uint256)
func (_ShMonad *ShMonadTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Deposit(&_ShMonad.TransactOpts, assets, receiver)
}

// DepositAndBond is a paid mutator transaction binding the contract method 0x08915200.
//
// Solidity: function depositAndBond(uint64 policyID, address bondRecipient, uint256 shMonToBond) payable returns()
func (_ShMonad *ShMonadTransactor) DepositAndBond(opts *bind.TransactOpts, policyID uint64, bondRecipient common.Address, shMonToBond *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "depositAndBond", policyID, bondRecipient, shMonToBond)
}

// DepositAndBond is a paid mutator transaction binding the contract method 0x08915200.
//
// Solidity: function depositAndBond(uint64 policyID, address bondRecipient, uint256 shMonToBond) payable returns()
func (_ShMonad *ShMonadSession) DepositAndBond(policyID uint64, bondRecipient common.Address, shMonToBond *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.DepositAndBond(&_ShMonad.TransactOpts, policyID, bondRecipient, shMonToBond)
}

// DepositAndBond is a paid mutator transaction binding the contract method 0x08915200.
//
// Solidity: function depositAndBond(uint64 policyID, address bondRecipient, uint256 shMonToBond) payable returns()
func (_ShMonad *ShMonadTransactorSession) DepositAndBond(policyID uint64, bondRecipient common.Address, shMonToBond *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.DepositAndBond(&_ShMonad.TransactOpts, policyID, bondRecipient, shMonToBond)
}

// DisablePolicy is a paid mutator transaction binding the contract method 0x4d28a646.
//
// Solidity: function disablePolicy(uint64 policyID) returns()
func (_ShMonad *ShMonadTransactor) DisablePolicy(opts *bind.TransactOpts, policyID uint64) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "disablePolicy", policyID)
}

// DisablePolicy is a paid mutator transaction binding the contract method 0x4d28a646.
//
// Solidity: function disablePolicy(uint64 policyID) returns()
func (_ShMonad *ShMonadSession) DisablePolicy(policyID uint64) (*types.Transaction, error) {
	return _ShMonad.Contract.DisablePolicy(&_ShMonad.TransactOpts, policyID)
}

// DisablePolicy is a paid mutator transaction binding the contract method 0x4d28a646.
//
// Solidity: function disablePolicy(uint64 policyID) returns()
func (_ShMonad *ShMonadTransactorSession) DisablePolicy(policyID uint64) (*types.Transaction, error) {
	return _ShMonad.Contract.DisablePolicy(&_ShMonad.TransactOpts, policyID)
}

// Hold is a paid mutator transaction binding the contract method 0xaba9598b.
//
// Solidity: function hold(uint64 policyID, address account, uint256 amount) returns()
func (_ShMonad *ShMonadTransactor) Hold(opts *bind.TransactOpts, policyID uint64, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "hold", policyID, account, amount)
}

// Hold is a paid mutator transaction binding the contract method 0xaba9598b.
//
// Solidity: function hold(uint64 policyID, address account, uint256 amount) returns()
func (_ShMonad *ShMonadSession) Hold(policyID uint64, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Hold(&_ShMonad.TransactOpts, policyID, account, amount)
}

// Hold is a paid mutator transaction binding the contract method 0xaba9598b.
//
// Solidity: function hold(uint64 policyID, address account, uint256 amount) returns()
func (_ShMonad *ShMonadTransactorSession) Hold(policyID uint64, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Hold(&_ShMonad.TransactOpts, policyID, account, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address deployer) returns()
func (_ShMonad *ShMonadTransactor) Initialize(opts *bind.TransactOpts, deployer common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "initialize", deployer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address deployer) returns()
func (_ShMonad *ShMonadSession) Initialize(deployer common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Initialize(&_ShMonad.TransactOpts, deployer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address deployer) returns()
func (_ShMonad *ShMonadTransactorSession) Initialize(deployer common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Initialize(&_ShMonad.TransactOpts, deployer)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) payable returns(uint256)
func (_ShMonad *ShMonadTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) payable returns(uint256)
func (_ShMonad *ShMonadSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Mint(&_ShMonad.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) payable returns(uint256)
func (_ShMonad *ShMonadTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Mint(&_ShMonad.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ShMonad *ShMonadTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ShMonad *ShMonadSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ShMonad.Contract.Permit(&_ShMonad.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ShMonad *ShMonadTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ShMonad.Contract.Permit(&_ShMonad.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ShMonad *ShMonadTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ShMonad *ShMonadSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Redeem(&_ShMonad.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ShMonad *ShMonadTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Redeem(&_ShMonad.TransactOpts, shares, receiver, owner)
}

// Release is a paid mutator transaction binding the contract method 0xcc56a82a.
//
// Solidity: function release(uint64 policyID, address account, uint256 amount) returns()
func (_ShMonad *ShMonadTransactor) Release(opts *bind.TransactOpts, policyID uint64, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "release", policyID, account, amount)
}

// Release is a paid mutator transaction binding the contract method 0xcc56a82a.
//
// Solidity: function release(uint64 policyID, address account, uint256 amount) returns()
func (_ShMonad *ShMonadSession) Release(policyID uint64, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Release(&_ShMonad.TransactOpts, policyID, account, amount)
}

// Release is a paid mutator transaction binding the contract method 0xcc56a82a.
//
// Solidity: function release(uint64 policyID, address account, uint256 amount) returns()
func (_ShMonad *ShMonadTransactorSession) Release(policyID uint64, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Release(&_ShMonad.TransactOpts, policyID, account, amount)
}

// RemovePolicyAgent is a paid mutator transaction binding the contract method 0x13788fe5.
//
// Solidity: function removePolicyAgent(uint64 policyID, address agent) returns()
func (_ShMonad *ShMonadTransactor) RemovePolicyAgent(opts *bind.TransactOpts, policyID uint64, agent common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "removePolicyAgent", policyID, agent)
}

// RemovePolicyAgent is a paid mutator transaction binding the contract method 0x13788fe5.
//
// Solidity: function removePolicyAgent(uint64 policyID, address agent) returns()
func (_ShMonad *ShMonadSession) RemovePolicyAgent(policyID uint64, agent common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.RemovePolicyAgent(&_ShMonad.TransactOpts, policyID, agent)
}

// RemovePolicyAgent is a paid mutator transaction binding the contract method 0x13788fe5.
//
// Solidity: function removePolicyAgent(uint64 policyID, address agent) returns()
func (_ShMonad *ShMonadTransactorSession) RemovePolicyAgent(policyID uint64, agent common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.RemovePolicyAgent(&_ShMonad.TransactOpts, policyID, agent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShMonad *ShMonadTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShMonad *ShMonadSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShMonad.Contract.RenounceOwnership(&_ShMonad.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShMonad *ShMonadTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShMonad.Contract.RenounceOwnership(&_ShMonad.TransactOpts)
}

// SetMinBondedBalance is a paid mutator transaction binding the contract method 0xb1e986a6.
//
// Solidity: function setMinBondedBalance(uint64 policyID, uint128 minBonded, uint128 maxTopUpPerPeriod, uint32 topUpPeriodDuration) returns()
func (_ShMonad *ShMonadTransactor) SetMinBondedBalance(opts *bind.TransactOpts, policyID uint64, minBonded *big.Int, maxTopUpPerPeriod *big.Int, topUpPeriodDuration uint32) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "setMinBondedBalance", policyID, minBonded, maxTopUpPerPeriod, topUpPeriodDuration)
}

// SetMinBondedBalance is a paid mutator transaction binding the contract method 0xb1e986a6.
//
// Solidity: function setMinBondedBalance(uint64 policyID, uint128 minBonded, uint128 maxTopUpPerPeriod, uint32 topUpPeriodDuration) returns()
func (_ShMonad *ShMonadSession) SetMinBondedBalance(policyID uint64, minBonded *big.Int, maxTopUpPerPeriod *big.Int, topUpPeriodDuration uint32) (*types.Transaction, error) {
	return _ShMonad.Contract.SetMinBondedBalance(&_ShMonad.TransactOpts, policyID, minBonded, maxTopUpPerPeriod, topUpPeriodDuration)
}

// SetMinBondedBalance is a paid mutator transaction binding the contract method 0xb1e986a6.
//
// Solidity: function setMinBondedBalance(uint64 policyID, uint128 minBonded, uint128 maxTopUpPerPeriod, uint32 topUpPeriodDuration) returns()
func (_ShMonad *ShMonadTransactorSession) SetMinBondedBalance(policyID uint64, minBonded *big.Int, maxTopUpPerPeriod *big.Int, topUpPeriodDuration uint32) (*types.Transaction, error) {
	return _ShMonad.Contract.SetMinBondedBalance(&_ShMonad.TransactOpts, policyID, minBonded, maxTopUpPerPeriod, topUpPeriodDuration)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ShMonad *ShMonadTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ShMonad *ShMonadSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Transfer(&_ShMonad.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ShMonad *ShMonadTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Transfer(&_ShMonad.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ShMonad *ShMonadTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ShMonad *ShMonadSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.TransferFrom(&_ShMonad.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ShMonad *ShMonadTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.TransferFrom(&_ShMonad.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShMonad *ShMonadTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShMonad *ShMonadSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.TransferOwnership(&_ShMonad.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShMonad *ShMonadTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.TransferOwnership(&_ShMonad.TransactOpts, newOwner)
}

// Unbond is a paid mutator transaction binding the contract method 0xa45df497.
//
// Solidity: function unbond(uint64 policyID, uint256 shares, uint256 newMinBalance) returns(uint256 unbondBlock)
func (_ShMonad *ShMonadTransactor) Unbond(opts *bind.TransactOpts, policyID uint64, shares *big.Int, newMinBalance *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "unbond", policyID, shares, newMinBalance)
}

// Unbond is a paid mutator transaction binding the contract method 0xa45df497.
//
// Solidity: function unbond(uint64 policyID, uint256 shares, uint256 newMinBalance) returns(uint256 unbondBlock)
func (_ShMonad *ShMonadSession) Unbond(policyID uint64, shares *big.Int, newMinBalance *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Unbond(&_ShMonad.TransactOpts, policyID, shares, newMinBalance)
}

// Unbond is a paid mutator transaction binding the contract method 0xa45df497.
//
// Solidity: function unbond(uint64 policyID, uint256 shares, uint256 newMinBalance) returns(uint256 unbondBlock)
func (_ShMonad *ShMonadTransactorSession) Unbond(policyID uint64, shares *big.Int, newMinBalance *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.Unbond(&_ShMonad.TransactOpts, policyID, shares, newMinBalance)
}

// UnbondWithTask is a paid mutator transaction binding the contract method 0xbe358022.
//
// Solidity: function unbondWithTask(uint64 policyID, uint256 shares, uint256 newMinBalance) payable returns(uint256 unbondBlock)
func (_ShMonad *ShMonadTransactor) UnbondWithTask(opts *bind.TransactOpts, policyID uint64, shares *big.Int, newMinBalance *big.Int) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "unbondWithTask", policyID, shares, newMinBalance)
}

// UnbondWithTask is a paid mutator transaction binding the contract method 0xbe358022.
//
// Solidity: function unbondWithTask(uint64 policyID, uint256 shares, uint256 newMinBalance) payable returns(uint256 unbondBlock)
func (_ShMonad *ShMonadSession) UnbondWithTask(policyID uint64, shares *big.Int, newMinBalance *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.UnbondWithTask(&_ShMonad.TransactOpts, policyID, shares, newMinBalance)
}

// UnbondWithTask is a paid mutator transaction binding the contract method 0xbe358022.
//
// Solidity: function unbondWithTask(uint64 policyID, uint256 shares, uint256 newMinBalance) payable returns(uint256 unbondBlock)
func (_ShMonad *ShMonadTransactorSession) UnbondWithTask(policyID uint64, shares *big.Int, newMinBalance *big.Int) (*types.Transaction, error) {
	return _ShMonad.Contract.UnbondWithTask(&_ShMonad.TransactOpts, policyID, shares, newMinBalance)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ShMonad *ShMonadTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ShMonad.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ShMonad *ShMonadSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Withdraw(&_ShMonad.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ShMonad *ShMonadTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ShMonad.Contract.Withdraw(&_ShMonad.TransactOpts, assets, receiver, owner)
}

// ShMonadAddPolicyAgentIterator is returned from FilterAddPolicyAgent and is used to iterate over the raw logs and unpacked data for AddPolicyAgent events raised by the ShMonad contract.
type ShMonadAddPolicyAgentIterator struct {
	Event *ShMonadAddPolicyAgent // Event containing the contract specifics and raw log

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
func (it *ShMonadAddPolicyAgentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadAddPolicyAgent)
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
		it.Event = new(ShMonadAddPolicyAgent)
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
func (it *ShMonadAddPolicyAgentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadAddPolicyAgentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadAddPolicyAgent represents a AddPolicyAgent event raised by the ShMonad contract.
type ShMonadAddPolicyAgent struct {
	PolicyID uint64
	Agent    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddPolicyAgent is a free log retrieval operation binding the contract event 0x536699e73f1d293b0cdae5c5bd2880028ce7a620ca9d2d7fd725b99bcca4e93e.
//
// Solidity: event AddPolicyAgent(uint64 indexed policyID, address indexed agent)
func (_ShMonad *ShMonadFilterer) FilterAddPolicyAgent(opts *bind.FilterOpts, policyID []uint64, agent []common.Address) (*ShMonadAddPolicyAgentIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "AddPolicyAgent", policyIDRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadAddPolicyAgentIterator{contract: _ShMonad.contract, event: "AddPolicyAgent", logs: logs, sub: sub}, nil
}

// WatchAddPolicyAgent is a free log subscription operation binding the contract event 0x536699e73f1d293b0cdae5c5bd2880028ce7a620ca9d2d7fd725b99bcca4e93e.
//
// Solidity: event AddPolicyAgent(uint64 indexed policyID, address indexed agent)
func (_ShMonad *ShMonadFilterer) WatchAddPolicyAgent(opts *bind.WatchOpts, sink chan<- *ShMonadAddPolicyAgent, policyID []uint64, agent []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "AddPolicyAgent", policyIDRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadAddPolicyAgent)
				if err := _ShMonad.contract.UnpackLog(event, "AddPolicyAgent", log); err != nil {
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

// ParseAddPolicyAgent is a log parse operation binding the contract event 0x536699e73f1d293b0cdae5c5bd2880028ce7a620ca9d2d7fd725b99bcca4e93e.
//
// Solidity: event AddPolicyAgent(uint64 indexed policyID, address indexed agent)
func (_ShMonad *ShMonadFilterer) ParseAddPolicyAgent(log types.Log) (*ShMonadAddPolicyAgent, error) {
	event := new(ShMonadAddPolicyAgent)
	if err := _ShMonad.contract.UnpackLog(event, "AddPolicyAgent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadAgentBoostYieldFromBondedIterator is returned from FilterAgentBoostYieldFromBonded and is used to iterate over the raw logs and unpacked data for AgentBoostYieldFromBonded events raised by the ShMonad contract.
type ShMonadAgentBoostYieldFromBondedIterator struct {
	Event *ShMonadAgentBoostYieldFromBonded // Event containing the contract specifics and raw log

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
func (it *ShMonadAgentBoostYieldFromBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadAgentBoostYieldFromBonded)
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
		it.Event = new(ShMonadAgentBoostYieldFromBonded)
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
func (it *ShMonadAgentBoostYieldFromBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadAgentBoostYieldFromBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadAgentBoostYieldFromBonded represents a AgentBoostYieldFromBonded event raised by the ShMonad contract.
type ShMonadAgentBoostYieldFromBonded struct {
	PolicyID uint64
	From     common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentBoostYieldFromBonded is a free log retrieval operation binding the contract event 0x59faf099c17e5201749d4140de5a4173bd83a409d233cdeb392e044764a1b1d8.
//
// Solidity: event AgentBoostYieldFromBonded(uint64 indexed policyID, address indexed from, uint256 amount)
func (_ShMonad *ShMonadFilterer) FilterAgentBoostYieldFromBonded(opts *bind.FilterOpts, policyID []uint64, from []common.Address) (*ShMonadAgentBoostYieldFromBondedIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "AgentBoostYieldFromBonded", policyIDRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadAgentBoostYieldFromBondedIterator{contract: _ShMonad.contract, event: "AgentBoostYieldFromBonded", logs: logs, sub: sub}, nil
}

// WatchAgentBoostYieldFromBonded is a free log subscription operation binding the contract event 0x59faf099c17e5201749d4140de5a4173bd83a409d233cdeb392e044764a1b1d8.
//
// Solidity: event AgentBoostYieldFromBonded(uint64 indexed policyID, address indexed from, uint256 amount)
func (_ShMonad *ShMonadFilterer) WatchAgentBoostYieldFromBonded(opts *bind.WatchOpts, sink chan<- *ShMonadAgentBoostYieldFromBonded, policyID []uint64, from []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "AgentBoostYieldFromBonded", policyIDRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadAgentBoostYieldFromBonded)
				if err := _ShMonad.contract.UnpackLog(event, "AgentBoostYieldFromBonded", log); err != nil {
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

// ParseAgentBoostYieldFromBonded is a log parse operation binding the contract event 0x59faf099c17e5201749d4140de5a4173bd83a409d233cdeb392e044764a1b1d8.
//
// Solidity: event AgentBoostYieldFromBonded(uint64 indexed policyID, address indexed from, uint256 amount)
func (_ShMonad *ShMonadFilterer) ParseAgentBoostYieldFromBonded(log types.Log) (*ShMonadAgentBoostYieldFromBonded, error) {
	event := new(ShMonadAgentBoostYieldFromBonded)
	if err := _ShMonad.contract.UnpackLog(event, "AgentBoostYieldFromBonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadAgentExecuteWithSponsorIterator is returned from FilterAgentExecuteWithSponsor and is used to iterate over the raw logs and unpacked data for AgentExecuteWithSponsor events raised by the ShMonad contract.
type ShMonadAgentExecuteWithSponsorIterator struct {
	Event *ShMonadAgentExecuteWithSponsor // Event containing the contract specifics and raw log

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
func (it *ShMonadAgentExecuteWithSponsorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadAgentExecuteWithSponsor)
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
		it.Event = new(ShMonadAgentExecuteWithSponsor)
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
func (it *ShMonadAgentExecuteWithSponsorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadAgentExecuteWithSponsorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadAgentExecuteWithSponsor represents a AgentExecuteWithSponsor event raised by the ShMonad contract.
type ShMonadAgentExecuteWithSponsor struct {
	PolicyID        uint64
	Payor           common.Address
	Agent           common.Address
	Recipient       common.Address
	MsgValue        *big.Int
	GasLimit        *big.Int
	ActualPayorCost *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAgentExecuteWithSponsor is a free log retrieval operation binding the contract event 0x9740c7400cd7fcf275c3a44d18ca3bb10d14c08f18614b7ac08f69c10361d38d.
//
// Solidity: event AgentExecuteWithSponsor(uint64 indexed policyID, address indexed payor, address indexed agent, address recipient, uint256 msgValue, uint256 gasLimit, uint256 actualPayorCost)
func (_ShMonad *ShMonadFilterer) FilterAgentExecuteWithSponsor(opts *bind.FilterOpts, policyID []uint64, payor []common.Address, agent []common.Address) (*ShMonadAgentExecuteWithSponsorIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "AgentExecuteWithSponsor", policyIDRule, payorRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadAgentExecuteWithSponsorIterator{contract: _ShMonad.contract, event: "AgentExecuteWithSponsor", logs: logs, sub: sub}, nil
}

// WatchAgentExecuteWithSponsor is a free log subscription operation binding the contract event 0x9740c7400cd7fcf275c3a44d18ca3bb10d14c08f18614b7ac08f69c10361d38d.
//
// Solidity: event AgentExecuteWithSponsor(uint64 indexed policyID, address indexed payor, address indexed agent, address recipient, uint256 msgValue, uint256 gasLimit, uint256 actualPayorCost)
func (_ShMonad *ShMonadFilterer) WatchAgentExecuteWithSponsor(opts *bind.WatchOpts, sink chan<- *ShMonadAgentExecuteWithSponsor, policyID []uint64, payor []common.Address, agent []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "AgentExecuteWithSponsor", policyIDRule, payorRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadAgentExecuteWithSponsor)
				if err := _ShMonad.contract.UnpackLog(event, "AgentExecuteWithSponsor", log); err != nil {
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

// ParseAgentExecuteWithSponsor is a log parse operation binding the contract event 0x9740c7400cd7fcf275c3a44d18ca3bb10d14c08f18614b7ac08f69c10361d38d.
//
// Solidity: event AgentExecuteWithSponsor(uint64 indexed policyID, address indexed payor, address indexed agent, address recipient, uint256 msgValue, uint256 gasLimit, uint256 actualPayorCost)
func (_ShMonad *ShMonadFilterer) ParseAgentExecuteWithSponsor(log types.Log) (*ShMonadAgentExecuteWithSponsor, error) {
	event := new(ShMonadAgentExecuteWithSponsor)
	if err := _ShMonad.contract.UnpackLog(event, "AgentExecuteWithSponsor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadAgentTransferFromBondedIterator is returned from FilterAgentTransferFromBonded and is used to iterate over the raw logs and unpacked data for AgentTransferFromBonded events raised by the ShMonad contract.
type ShMonadAgentTransferFromBondedIterator struct {
	Event *ShMonadAgentTransferFromBonded // Event containing the contract specifics and raw log

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
func (it *ShMonadAgentTransferFromBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadAgentTransferFromBonded)
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
		it.Event = new(ShMonadAgentTransferFromBonded)
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
func (it *ShMonadAgentTransferFromBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadAgentTransferFromBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadAgentTransferFromBonded represents a AgentTransferFromBonded event raised by the ShMonad contract.
type ShMonadAgentTransferFromBonded struct {
	PolicyID uint64
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentTransferFromBonded is a free log retrieval operation binding the contract event 0x4e03b940b5a8784d593bb4747b48911bd4ab47cc890e28cfeda002a6d6bed227.
//
// Solidity: event AgentTransferFromBonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) FilterAgentTransferFromBonded(opts *bind.FilterOpts, policyID []uint64, from []common.Address, to []common.Address) (*ShMonadAgentTransferFromBondedIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "AgentTransferFromBonded", policyIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadAgentTransferFromBondedIterator{contract: _ShMonad.contract, event: "AgentTransferFromBonded", logs: logs, sub: sub}, nil
}

// WatchAgentTransferFromBonded is a free log subscription operation binding the contract event 0x4e03b940b5a8784d593bb4747b48911bd4ab47cc890e28cfeda002a6d6bed227.
//
// Solidity: event AgentTransferFromBonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) WatchAgentTransferFromBonded(opts *bind.WatchOpts, sink chan<- *ShMonadAgentTransferFromBonded, policyID []uint64, from []common.Address, to []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "AgentTransferFromBonded", policyIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadAgentTransferFromBonded)
				if err := _ShMonad.contract.UnpackLog(event, "AgentTransferFromBonded", log); err != nil {
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

// ParseAgentTransferFromBonded is a log parse operation binding the contract event 0x4e03b940b5a8784d593bb4747b48911bd4ab47cc890e28cfeda002a6d6bed227.
//
// Solidity: event AgentTransferFromBonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) ParseAgentTransferFromBonded(log types.Log) (*ShMonadAgentTransferFromBonded, error) {
	event := new(ShMonadAgentTransferFromBonded)
	if err := _ShMonad.contract.UnpackLog(event, "AgentTransferFromBonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadAgentTransferToUnbondedIterator is returned from FilterAgentTransferToUnbonded and is used to iterate over the raw logs and unpacked data for AgentTransferToUnbonded events raised by the ShMonad contract.
type ShMonadAgentTransferToUnbondedIterator struct {
	Event *ShMonadAgentTransferToUnbonded // Event containing the contract specifics and raw log

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
func (it *ShMonadAgentTransferToUnbondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadAgentTransferToUnbonded)
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
		it.Event = new(ShMonadAgentTransferToUnbonded)
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
func (it *ShMonadAgentTransferToUnbondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadAgentTransferToUnbondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadAgentTransferToUnbonded represents a AgentTransferToUnbonded event raised by the ShMonad contract.
type ShMonadAgentTransferToUnbonded struct {
	PolicyID uint64
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentTransferToUnbonded is a free log retrieval operation binding the contract event 0x3653bb505b47c5ee335029e29814fa882c06bb9cb4ab1a55be7e919a1750338b.
//
// Solidity: event AgentTransferToUnbonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) FilterAgentTransferToUnbonded(opts *bind.FilterOpts, policyID []uint64, from []common.Address, to []common.Address) (*ShMonadAgentTransferToUnbondedIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "AgentTransferToUnbonded", policyIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadAgentTransferToUnbondedIterator{contract: _ShMonad.contract, event: "AgentTransferToUnbonded", logs: logs, sub: sub}, nil
}

// WatchAgentTransferToUnbonded is a free log subscription operation binding the contract event 0x3653bb505b47c5ee335029e29814fa882c06bb9cb4ab1a55be7e919a1750338b.
//
// Solidity: event AgentTransferToUnbonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) WatchAgentTransferToUnbonded(opts *bind.WatchOpts, sink chan<- *ShMonadAgentTransferToUnbonded, policyID []uint64, from []common.Address, to []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "AgentTransferToUnbonded", policyIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadAgentTransferToUnbonded)
				if err := _ShMonad.contract.UnpackLog(event, "AgentTransferToUnbonded", log); err != nil {
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

// ParseAgentTransferToUnbonded is a log parse operation binding the contract event 0x3653bb505b47c5ee335029e29814fa882c06bb9cb4ab1a55be7e919a1750338b.
//
// Solidity: event AgentTransferToUnbonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) ParseAgentTransferToUnbonded(log types.Log) (*ShMonadAgentTransferToUnbonded, error) {
	event := new(ShMonadAgentTransferToUnbonded)
	if err := _ShMonad.contract.UnpackLog(event, "AgentTransferToUnbonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadAgentWithdrawFromBondedIterator is returned from FilterAgentWithdrawFromBonded and is used to iterate over the raw logs and unpacked data for AgentWithdrawFromBonded events raised by the ShMonad contract.
type ShMonadAgentWithdrawFromBondedIterator struct {
	Event *ShMonadAgentWithdrawFromBonded // Event containing the contract specifics and raw log

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
func (it *ShMonadAgentWithdrawFromBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadAgentWithdrawFromBonded)
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
		it.Event = new(ShMonadAgentWithdrawFromBonded)
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
func (it *ShMonadAgentWithdrawFromBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadAgentWithdrawFromBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadAgentWithdrawFromBonded represents a AgentWithdrawFromBonded event raised by the ShMonad contract.
type ShMonadAgentWithdrawFromBonded struct {
	PolicyID uint64
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentWithdrawFromBonded is a free log retrieval operation binding the contract event 0x1fd76891c6c6397aab11d2503d2906f4c0d799713801ba670a58a031d24814ae.
//
// Solidity: event AgentWithdrawFromBonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) FilterAgentWithdrawFromBonded(opts *bind.FilterOpts, policyID []uint64, from []common.Address, to []common.Address) (*ShMonadAgentWithdrawFromBondedIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "AgentWithdrawFromBonded", policyIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadAgentWithdrawFromBondedIterator{contract: _ShMonad.contract, event: "AgentWithdrawFromBonded", logs: logs, sub: sub}, nil
}

// WatchAgentWithdrawFromBonded is a free log subscription operation binding the contract event 0x1fd76891c6c6397aab11d2503d2906f4c0d799713801ba670a58a031d24814ae.
//
// Solidity: event AgentWithdrawFromBonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) WatchAgentWithdrawFromBonded(opts *bind.WatchOpts, sink chan<- *ShMonadAgentWithdrawFromBonded, policyID []uint64, from []common.Address, to []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "AgentWithdrawFromBonded", policyIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadAgentWithdrawFromBonded)
				if err := _ShMonad.contract.UnpackLog(event, "AgentWithdrawFromBonded", log); err != nil {
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

// ParseAgentWithdrawFromBonded is a log parse operation binding the contract event 0x1fd76891c6c6397aab11d2503d2906f4c0d799713801ba670a58a031d24814ae.
//
// Solidity: event AgentWithdrawFromBonded(uint64 indexed policyID, address indexed from, address indexed to, uint256 amount)
func (_ShMonad *ShMonadFilterer) ParseAgentWithdrawFromBonded(log types.Log) (*ShMonadAgentWithdrawFromBonded, error) {
	event := new(ShMonadAgentWithdrawFromBonded)
	if err := _ShMonad.contract.UnpackLog(event, "AgentWithdrawFromBonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShMonad contract.
type ShMonadApprovalIterator struct {
	Event *ShMonadApproval // Event containing the contract specifics and raw log

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
func (it *ShMonadApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadApproval)
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
		it.Event = new(ShMonadApproval)
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
func (it *ShMonadApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadApproval represents a Approval event raised by the ShMonad contract.
type ShMonadApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShMonad *ShMonadFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ShMonadApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadApprovalIterator{contract: _ShMonad.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShMonad *ShMonadFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShMonadApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadApproval)
				if err := _ShMonad.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShMonad *ShMonadFilterer) ParseApproval(log types.Log) (*ShMonadApproval, error) {
	event := new(ShMonadApproval)
	if err := _ShMonad.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadBondIterator is returned from FilterBond and is used to iterate over the raw logs and unpacked data for Bond events raised by the ShMonad contract.
type ShMonadBondIterator struct {
	Event *ShMonadBond // Event containing the contract specifics and raw log

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
func (it *ShMonadBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadBond)
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
		it.Event = new(ShMonadBond)
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
func (it *ShMonadBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadBond represents a Bond event raised by the ShMonad contract.
type ShMonadBond struct {
	PolicyID uint64
	Account  common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBond is a free log retrieval operation binding the contract event 0x92041dc62baaf5ab79bf2ce62189d81390117160481113e1a622c8a26aad1b02.
//
// Solidity: event Bond(uint64 indexed policyID, address indexed account, uint256 amount)
func (_ShMonad *ShMonadFilterer) FilterBond(opts *bind.FilterOpts, policyID []uint64, account []common.Address) (*ShMonadBondIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Bond", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadBondIterator{contract: _ShMonad.contract, event: "Bond", logs: logs, sub: sub}, nil
}

// WatchBond is a free log subscription operation binding the contract event 0x92041dc62baaf5ab79bf2ce62189d81390117160481113e1a622c8a26aad1b02.
//
// Solidity: event Bond(uint64 indexed policyID, address indexed account, uint256 amount)
func (_ShMonad *ShMonadFilterer) WatchBond(opts *bind.WatchOpts, sink chan<- *ShMonadBond, policyID []uint64, account []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Bond", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadBond)
				if err := _ShMonad.contract.UnpackLog(event, "Bond", log); err != nil {
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

// ParseBond is a log parse operation binding the contract event 0x92041dc62baaf5ab79bf2ce62189d81390117160481113e1a622c8a26aad1b02.
//
// Solidity: event Bond(uint64 indexed policyID, address indexed account, uint256 amount)
func (_ShMonad *ShMonadFilterer) ParseBond(log types.Log) (*ShMonadBond, error) {
	event := new(ShMonadBond)
	if err := _ShMonad.contract.UnpackLog(event, "Bond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the ShMonad contract.
type ShMonadClaimIterator struct {
	Event *ShMonadClaim // Event containing the contract specifics and raw log

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
func (it *ShMonadClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadClaim)
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
		it.Event = new(ShMonadClaim)
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
func (it *ShMonadClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadClaim represents a Claim event raised by the ShMonad contract.
type ShMonadClaim struct {
	PolicyID uint64
	Account  common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0xfe8aa29cf039ed63090c19c8df7227928debdffaaabc3980055f519185105ef1.
//
// Solidity: event Claim(uint64 indexed policyID, address indexed account, uint256 amount)
func (_ShMonad *ShMonadFilterer) FilterClaim(opts *bind.FilterOpts, policyID []uint64, account []common.Address) (*ShMonadClaimIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Claim", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadClaimIterator{contract: _ShMonad.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0xfe8aa29cf039ed63090c19c8df7227928debdffaaabc3980055f519185105ef1.
//
// Solidity: event Claim(uint64 indexed policyID, address indexed account, uint256 amount)
func (_ShMonad *ShMonadFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *ShMonadClaim, policyID []uint64, account []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Claim", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadClaim)
				if err := _ShMonad.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0xfe8aa29cf039ed63090c19c8df7227928debdffaaabc3980055f519185105ef1.
//
// Solidity: event Claim(uint64 indexed policyID, address indexed account, uint256 amount)
func (_ShMonad *ShMonadFilterer) ParseClaim(log types.Log) (*ShMonadClaim, error) {
	event := new(ShMonadClaim)
	if err := _ShMonad.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadCreatePolicyIterator is returned from FilterCreatePolicy and is used to iterate over the raw logs and unpacked data for CreatePolicy events raised by the ShMonad contract.
type ShMonadCreatePolicyIterator struct {
	Event *ShMonadCreatePolicy // Event containing the contract specifics and raw log

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
func (it *ShMonadCreatePolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadCreatePolicy)
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
		it.Event = new(ShMonadCreatePolicy)
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
func (it *ShMonadCreatePolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadCreatePolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadCreatePolicy represents a CreatePolicy event raised by the ShMonad contract.
type ShMonadCreatePolicy struct {
	PolicyID       uint64
	Creator        common.Address
	EscrowDuration *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCreatePolicy is a free log retrieval operation binding the contract event 0x369b9ef7936f8128d1eff4409b83d784498ec057c569447a736deab847ed7de9.
//
// Solidity: event CreatePolicy(uint64 indexed policyID, address indexed creator, uint48 escrowDuration)
func (_ShMonad *ShMonadFilterer) FilterCreatePolicy(opts *bind.FilterOpts, policyID []uint64, creator []common.Address) (*ShMonadCreatePolicyIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "CreatePolicy", policyIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadCreatePolicyIterator{contract: _ShMonad.contract, event: "CreatePolicy", logs: logs, sub: sub}, nil
}

// WatchCreatePolicy is a free log subscription operation binding the contract event 0x369b9ef7936f8128d1eff4409b83d784498ec057c569447a736deab847ed7de9.
//
// Solidity: event CreatePolicy(uint64 indexed policyID, address indexed creator, uint48 escrowDuration)
func (_ShMonad *ShMonadFilterer) WatchCreatePolicy(opts *bind.WatchOpts, sink chan<- *ShMonadCreatePolicy, policyID []uint64, creator []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "CreatePolicy", policyIDRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadCreatePolicy)
				if err := _ShMonad.contract.UnpackLog(event, "CreatePolicy", log); err != nil {
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

// ParseCreatePolicy is a log parse operation binding the contract event 0x369b9ef7936f8128d1eff4409b83d784498ec057c569447a736deab847ed7de9.
//
// Solidity: event CreatePolicy(uint64 indexed policyID, address indexed creator, uint48 escrowDuration)
func (_ShMonad *ShMonadFilterer) ParseCreatePolicy(log types.Log) (*ShMonadCreatePolicy, error) {
	event := new(ShMonadCreatePolicy)
	if err := _ShMonad.contract.UnpackLog(event, "CreatePolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ShMonad contract.
type ShMonadDepositIterator struct {
	Event *ShMonadDeposit // Event containing the contract specifics and raw log

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
func (it *ShMonadDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadDeposit)
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
		it.Event = new(ShMonadDeposit)
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
func (it *ShMonadDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadDeposit represents a Deposit event raised by the ShMonad contract.
type ShMonadDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ShMonad *ShMonadFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*ShMonadDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadDepositIterator{contract: _ShMonad.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ShMonad *ShMonadFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ShMonadDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadDeposit)
				if err := _ShMonad.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ShMonad *ShMonadFilterer) ParseDeposit(log types.Log) (*ShMonadDeposit, error) {
	event := new(ShMonadDeposit)
	if err := _ShMonad.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadDisablePolicyIterator is returned from FilterDisablePolicy and is used to iterate over the raw logs and unpacked data for DisablePolicy events raised by the ShMonad contract.
type ShMonadDisablePolicyIterator struct {
	Event *ShMonadDisablePolicy // Event containing the contract specifics and raw log

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
func (it *ShMonadDisablePolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadDisablePolicy)
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
		it.Event = new(ShMonadDisablePolicy)
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
func (it *ShMonadDisablePolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadDisablePolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadDisablePolicy represents a DisablePolicy event raised by the ShMonad contract.
type ShMonadDisablePolicy struct {
	PolicyID uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDisablePolicy is a free log retrieval operation binding the contract event 0x724f90fbdf12cba6f20f1055f6d0bd3cdb9d3b5b8a1f4c758b92b18ee9f30fc8.
//
// Solidity: event DisablePolicy(uint64 indexed policyID)
func (_ShMonad *ShMonadFilterer) FilterDisablePolicy(opts *bind.FilterOpts, policyID []uint64) (*ShMonadDisablePolicyIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "DisablePolicy", policyIDRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadDisablePolicyIterator{contract: _ShMonad.contract, event: "DisablePolicy", logs: logs, sub: sub}, nil
}

// WatchDisablePolicy is a free log subscription operation binding the contract event 0x724f90fbdf12cba6f20f1055f6d0bd3cdb9d3b5b8a1f4c758b92b18ee9f30fc8.
//
// Solidity: event DisablePolicy(uint64 indexed policyID)
func (_ShMonad *ShMonadFilterer) WatchDisablePolicy(opts *bind.WatchOpts, sink chan<- *ShMonadDisablePolicy, policyID []uint64) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "DisablePolicy", policyIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadDisablePolicy)
				if err := _ShMonad.contract.UnpackLog(event, "DisablePolicy", log); err != nil {
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

// ParseDisablePolicy is a log parse operation binding the contract event 0x724f90fbdf12cba6f20f1055f6d0bd3cdb9d3b5b8a1f4c758b92b18ee9f30fc8.
//
// Solidity: event DisablePolicy(uint64 indexed policyID)
func (_ShMonad *ShMonadFilterer) ParseDisablePolicy(log types.Log) (*ShMonadDisablePolicy, error) {
	event := new(ShMonadDisablePolicy)
	if err := _ShMonad.contract.UnpackLog(event, "DisablePolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ShMonad contract.
type ShMonadEIP712DomainChangedIterator struct {
	Event *ShMonadEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ShMonadEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadEIP712DomainChanged)
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
		it.Event = new(ShMonadEIP712DomainChanged)
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
func (it *ShMonadEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadEIP712DomainChanged represents a EIP712DomainChanged event raised by the ShMonad contract.
type ShMonadEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ShMonad *ShMonadFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ShMonadEIP712DomainChangedIterator, error) {

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ShMonadEIP712DomainChangedIterator{contract: _ShMonad.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ShMonad *ShMonadFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ShMonadEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadEIP712DomainChanged)
				if err := _ShMonad.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ShMonad *ShMonadFilterer) ParseEIP712DomainChanged(log types.Log) (*ShMonadEIP712DomainChanged, error) {
	event := new(ShMonadEIP712DomainChanged)
	if err := _ShMonad.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ShMonad contract.
type ShMonadInitializedIterator struct {
	Event *ShMonadInitialized // Event containing the contract specifics and raw log

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
func (it *ShMonadInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadInitialized)
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
		it.Event = new(ShMonadInitialized)
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
func (it *ShMonadInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadInitialized represents a Initialized event raised by the ShMonad contract.
type ShMonadInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ShMonad *ShMonadFilterer) FilterInitialized(opts *bind.FilterOpts) (*ShMonadInitializedIterator, error) {

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ShMonadInitializedIterator{contract: _ShMonad.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ShMonad *ShMonadFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ShMonadInitialized) (event.Subscription, error) {

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadInitialized)
				if err := _ShMonad.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ShMonad *ShMonadFilterer) ParseInitialized(log types.Log) (*ShMonadInitialized, error) {
	event := new(ShMonadInitialized)
	if err := _ShMonad.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShMonad contract.
type ShMonadOwnershipTransferredIterator struct {
	Event *ShMonadOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShMonadOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadOwnershipTransferred)
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
		it.Event = new(ShMonadOwnershipTransferred)
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
func (it *ShMonadOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadOwnershipTransferred represents a OwnershipTransferred event raised by the ShMonad contract.
type ShMonadOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShMonad *ShMonadFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShMonadOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadOwnershipTransferredIterator{contract: _ShMonad.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShMonad *ShMonadFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShMonadOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadOwnershipTransferred)
				if err := _ShMonad.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShMonad *ShMonadFilterer) ParseOwnershipTransferred(log types.Log) (*ShMonadOwnershipTransferred, error) {
	event := new(ShMonadOwnershipTransferred)
	if err := _ShMonad.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadRemovePolicyAgentIterator is returned from FilterRemovePolicyAgent and is used to iterate over the raw logs and unpacked data for RemovePolicyAgent events raised by the ShMonad contract.
type ShMonadRemovePolicyAgentIterator struct {
	Event *ShMonadRemovePolicyAgent // Event containing the contract specifics and raw log

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
func (it *ShMonadRemovePolicyAgentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadRemovePolicyAgent)
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
		it.Event = new(ShMonadRemovePolicyAgent)
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
func (it *ShMonadRemovePolicyAgentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadRemovePolicyAgentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadRemovePolicyAgent represents a RemovePolicyAgent event raised by the ShMonad contract.
type ShMonadRemovePolicyAgent struct {
	PolicyID uint64
	Agent    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovePolicyAgent is a free log retrieval operation binding the contract event 0xd1dd9f27582bbce7226a4f323da8f0e775e31a715f1779e55da4e895c027f83a.
//
// Solidity: event RemovePolicyAgent(uint64 indexed policyID, address indexed agent)
func (_ShMonad *ShMonadFilterer) FilterRemovePolicyAgent(opts *bind.FilterOpts, policyID []uint64, agent []common.Address) (*ShMonadRemovePolicyAgentIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "RemovePolicyAgent", policyIDRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadRemovePolicyAgentIterator{contract: _ShMonad.contract, event: "RemovePolicyAgent", logs: logs, sub: sub}, nil
}

// WatchRemovePolicyAgent is a free log subscription operation binding the contract event 0xd1dd9f27582bbce7226a4f323da8f0e775e31a715f1779e55da4e895c027f83a.
//
// Solidity: event RemovePolicyAgent(uint64 indexed policyID, address indexed agent)
func (_ShMonad *ShMonadFilterer) WatchRemovePolicyAgent(opts *bind.WatchOpts, sink chan<- *ShMonadRemovePolicyAgent, policyID []uint64, agent []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "RemovePolicyAgent", policyIDRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadRemovePolicyAgent)
				if err := _ShMonad.contract.UnpackLog(event, "RemovePolicyAgent", log); err != nil {
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

// ParseRemovePolicyAgent is a log parse operation binding the contract event 0xd1dd9f27582bbce7226a4f323da8f0e775e31a715f1779e55da4e895c027f83a.
//
// Solidity: event RemovePolicyAgent(uint64 indexed policyID, address indexed agent)
func (_ShMonad *ShMonadFilterer) ParseRemovePolicyAgent(log types.Log) (*ShMonadRemovePolicyAgent, error) {
	event := new(ShMonadRemovePolicyAgent)
	if err := _ShMonad.contract.UnpackLog(event, "RemovePolicyAgent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadSetTopUpIterator is returned from FilterSetTopUp and is used to iterate over the raw logs and unpacked data for SetTopUp events raised by the ShMonad contract.
type ShMonadSetTopUpIterator struct {
	Event *ShMonadSetTopUp // Event containing the contract specifics and raw log

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
func (it *ShMonadSetTopUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadSetTopUp)
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
		it.Event = new(ShMonadSetTopUp)
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
func (it *ShMonadSetTopUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadSetTopUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadSetTopUp represents a SetTopUp event raised by the ShMonad contract.
type ShMonadSetTopUp struct {
	PolicyID            uint64
	Account             common.Address
	MinBonded           *big.Int
	MaxTopUpPerPeriod   *big.Int
	TopUpPeriodDuration uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTopUp is a free log retrieval operation binding the contract event 0xcbb3320930c2934da510e3107d0acc2f5494678603ee08c1e646ccf545f9b10e.
//
// Solidity: event SetTopUp(uint64 indexed policyID, address indexed account, uint128 minBonded, uint128 maxTopUpPerPeriod, uint32 topUpPeriodDuration)
func (_ShMonad *ShMonadFilterer) FilterSetTopUp(opts *bind.FilterOpts, policyID []uint64, account []common.Address) (*ShMonadSetTopUpIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "SetTopUp", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadSetTopUpIterator{contract: _ShMonad.contract, event: "SetTopUp", logs: logs, sub: sub}, nil
}

// WatchSetTopUp is a free log subscription operation binding the contract event 0xcbb3320930c2934da510e3107d0acc2f5494678603ee08c1e646ccf545f9b10e.
//
// Solidity: event SetTopUp(uint64 indexed policyID, address indexed account, uint128 minBonded, uint128 maxTopUpPerPeriod, uint32 topUpPeriodDuration)
func (_ShMonad *ShMonadFilterer) WatchSetTopUp(opts *bind.WatchOpts, sink chan<- *ShMonadSetTopUp, policyID []uint64, account []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "SetTopUp", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadSetTopUp)
				if err := _ShMonad.contract.UnpackLog(event, "SetTopUp", log); err != nil {
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

// ParseSetTopUp is a log parse operation binding the contract event 0xcbb3320930c2934da510e3107d0acc2f5494678603ee08c1e646ccf545f9b10e.
//
// Solidity: event SetTopUp(uint64 indexed policyID, address indexed account, uint128 minBonded, uint128 maxTopUpPerPeriod, uint32 topUpPeriodDuration)
func (_ShMonad *ShMonadFilterer) ParseSetTopUp(log types.Log) (*ShMonadSetTopUp, error) {
	event := new(ShMonadSetTopUp)
	if err := _ShMonad.contract.UnpackLog(event, "SetTopUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShMonad contract.
type ShMonadTransferIterator struct {
	Event *ShMonadTransfer // Event containing the contract specifics and raw log

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
func (it *ShMonadTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadTransfer)
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
		it.Event = new(ShMonadTransfer)
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
func (it *ShMonadTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadTransfer represents a Transfer event raised by the ShMonad contract.
type ShMonadTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShMonad *ShMonadFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShMonadTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadTransferIterator{contract: _ShMonad.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShMonad *ShMonadFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShMonadTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadTransfer)
				if err := _ShMonad.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShMonad *ShMonadFilterer) ParseTransfer(log types.Log) (*ShMonadTransfer, error) {
	event := new(ShMonadTransfer)
	if err := _ShMonad.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the ShMonad contract.
type ShMonadUnbondIterator struct {
	Event *ShMonadUnbond // Event containing the contract specifics and raw log

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
func (it *ShMonadUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadUnbond)
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
		it.Event = new(ShMonadUnbond)
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
func (it *ShMonadUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadUnbond represents a Unbond event raised by the ShMonad contract.
type ShMonadUnbond struct {
	PolicyID            uint64
	Account             common.Address
	Amount              *big.Int
	ExpectedUnbondBlock *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0xbfb5ab35030371340c9b234059b781190459a96c96b0f75c371bb0f5df190d54.
//
// Solidity: event Unbond(uint64 indexed policyID, address indexed account, uint256 amount, uint256 expectedUnbondBlock)
func (_ShMonad *ShMonadFilterer) FilterUnbond(opts *bind.FilterOpts, policyID []uint64, account []common.Address) (*ShMonadUnbondIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Unbond", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadUnbondIterator{contract: _ShMonad.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0xbfb5ab35030371340c9b234059b781190459a96c96b0f75c371bb0f5df190d54.
//
// Solidity: event Unbond(uint64 indexed policyID, address indexed account, uint256 amount, uint256 expectedUnbondBlock)
func (_ShMonad *ShMonadFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *ShMonadUnbond, policyID []uint64, account []common.Address) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Unbond", policyIDRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadUnbond)
				if err := _ShMonad.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0xbfb5ab35030371340c9b234059b781190459a96c96b0f75c371bb0f5df190d54.
//
// Solidity: event Unbond(uint64 indexed policyID, address indexed account, uint256 amount, uint256 expectedUnbondBlock)
func (_ShMonad *ShMonadFilterer) ParseUnbond(log types.Log) (*ShMonadUnbond, error) {
	event := new(ShMonadUnbond)
	if err := _ShMonad.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShMonadWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ShMonad contract.
type ShMonadWithdrawIterator struct {
	Event *ShMonadWithdraw // Event containing the contract specifics and raw log

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
func (it *ShMonadWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShMonadWithdraw)
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
		it.Event = new(ShMonadWithdraw)
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
func (it *ShMonadWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShMonadWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShMonadWithdraw represents a Withdraw event raised by the ShMonad contract.
type ShMonadWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ShMonad *ShMonadFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*ShMonadWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ShMonad.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ShMonadWithdrawIterator{contract: _ShMonad.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ShMonad *ShMonadFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ShMonadWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ShMonad.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShMonadWithdraw)
				if err := _ShMonad.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ShMonad *ShMonadFilterer) ParseWithdraw(log types.Log) (*ShMonadWithdraw, error) {
	event := new(ShMonadWithdraw)
	if err := _ShMonad.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
