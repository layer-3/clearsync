// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilynex_v3_pool

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

// ILynexV3PoolMetaData contains all meta data concerning the ILynexV3Pool contract.
var ILynexV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeIncentive\",\"inputs\":[],\"outputs\":[{\"name\":\"virtualPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"bottomTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collect\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bottomTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount0Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dataStorageOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flash\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getInnerCumulatives\",\"inputs\":[{\"name\":\"bottomTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"innerTickCumulative\",\"type\":\"int56\",\"internalType\":\"int56\"},{\"name\":\"innerSecondsSpentPerLiquidity\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"innerSecondsSpent\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimepoints\",\"inputs\":[{\"name\":\"secondsAgos\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[{\"name\":\"tickCumulatives\",\"type\":\"int56[]\",\"internalType\":\"int56[]\"},{\"name\":\"secondsPerLiquidityCumulatives\",\"type\":\"uint160[]\",\"internalType\":\"uint160[]\"},{\"name\":\"volatilityCumulatives\",\"type\":\"uint112[]\",\"internalType\":\"uint112[]\"},{\"name\":\"volumePerAvgLiquiditys\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalState\",\"inputs\":[],\"outputs\":[{\"name\":\"price\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"fee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"timepointIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"communityFeeToken0\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"communityFeeToken1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"unlocked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"price\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidityCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"cooldownInSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLiquidityPerTick\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bottomTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidityActual\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"liquidityAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastLiquidityAddTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"innerFeeGrowth0Token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"innerFeeGrowth1Token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fees0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"fees1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setCommunityFee\",\"inputs\":[{\"name\":\"communityFee0\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"communityFee1\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIncentive\",\"inputs\":[{\"name\":\"virtualPoolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidityCooldown\",\"inputs\":[{\"name\":\"newLiquidityCooldown\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTickSpacing\",\"inputs\":[{\"name\":\"newTickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zeroToOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"limitSqrtPrice\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapSupportingFeeOnInputTokens\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zeroToOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"limitSqrtPrice\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tickSpacing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tickTable\",\"inputs\":[{\"name\":\"wordPosition\",\"type\":\"int16\",\"internalType\":\"int16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ticks\",\"inputs\":[{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liquidityTotal\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"liquidityDelta\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"outerFeeGrowth0Token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outerFeeGrowth1Token\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outerTickCumulative\",\"type\":\"int56\",\"internalType\":\"int56\"},{\"name\":\"outerSecondsPerLiquidity\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"outerSecondsSpent\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timepoints\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"blockTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tickCumulative\",\"type\":\"int56\",\"internalType\":\"int56\"},{\"name\":\"secondsPerLiquidityCumulative\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"volatilityCumulative\",\"type\":\"uint88\",\"internalType\":\"uint88\"},{\"name\":\"averageTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"volumePerLiquidityCumulative\",\"type\":\"uint144\",\"internalType\":\"uint144\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalFeeGrowth0Token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalFeeGrowth1Token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bottomTick\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"liquidityAmount\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Collect\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bottomTick\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"amount0\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommunityFee\",\"inputs\":[{\"name\":\"communityFee0New\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"communityFee1New\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Flash\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paid0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paid1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Incentive\",\"inputs\":[{\"name\":\"virtualPoolAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialize\",\"inputs\":[{\"name\":\"price\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityCooldown\",\"inputs\":[{\"name\":\"liquidityCooldown\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bottomTick\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"topTick\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"liquidityAmount\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"price\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"liquidity\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TickSpacing\",\"inputs\":[{\"name\":\"newTickSpacing\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false}]",
}

// ILynexV3PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ILynexV3PoolMetaData.ABI instead.
var ILynexV3PoolABI = ILynexV3PoolMetaData.ABI

// ILynexV3Pool is an auto generated Go binding around an Ethereum contract.
type ILynexV3Pool struct {
	ILynexV3PoolCaller     // Read-only binding to the contract
	ILynexV3PoolTransactor // Write-only binding to the contract
	ILynexV3PoolFilterer   // Log filterer for contract events
}

// ILynexV3PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILynexV3PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexV3PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILynexV3PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexV3PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILynexV3PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexV3PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILynexV3PoolSession struct {
	Contract     *ILynexV3Pool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILynexV3PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILynexV3PoolCallerSession struct {
	Contract *ILynexV3PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ILynexV3PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILynexV3PoolTransactorSession struct {
	Contract     *ILynexV3PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ILynexV3PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILynexV3PoolRaw struct {
	Contract *ILynexV3Pool // Generic contract binding to access the raw methods on
}

// ILynexV3PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILynexV3PoolCallerRaw struct {
	Contract *ILynexV3PoolCaller // Generic read-only contract binding to access the raw methods on
}

// ILynexV3PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILynexV3PoolTransactorRaw struct {
	Contract *ILynexV3PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILynexV3Pool creates a new instance of ILynexV3Pool, bound to a specific deployed contract.
func NewILynexV3Pool(address common.Address, backend bind.ContractBackend) (*ILynexV3Pool, error) {
	contract, err := bindILynexV3Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILynexV3Pool{ILynexV3PoolCaller: ILynexV3PoolCaller{contract: contract}, ILynexV3PoolTransactor: ILynexV3PoolTransactor{contract: contract}, ILynexV3PoolFilterer: ILynexV3PoolFilterer{contract: contract}}, nil
}

// NewILynexV3PoolCaller creates a new read-only instance of ILynexV3Pool, bound to a specific deployed contract.
func NewILynexV3PoolCaller(address common.Address, caller bind.ContractCaller) (*ILynexV3PoolCaller, error) {
	contract, err := bindILynexV3Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolCaller{contract: contract}, nil
}

// NewILynexV3PoolTransactor creates a new write-only instance of ILynexV3Pool, bound to a specific deployed contract.
func NewILynexV3PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ILynexV3PoolTransactor, error) {
	contract, err := bindILynexV3Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolTransactor{contract: contract}, nil
}

// NewILynexV3PoolFilterer creates a new log filterer instance of ILynexV3Pool, bound to a specific deployed contract.
func NewILynexV3PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ILynexV3PoolFilterer, error) {
	contract, err := bindILynexV3Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolFilterer{contract: contract}, nil
}

// bindILynexV3Pool binds a generic wrapper to an already deployed contract.
func bindILynexV3Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILynexV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexV3Pool *ILynexV3PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexV3Pool.Contract.ILynexV3PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexV3Pool *ILynexV3PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.ILynexV3PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexV3Pool *ILynexV3PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.ILynexV3PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexV3Pool *ILynexV3PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexV3Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexV3Pool *ILynexV3PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexV3Pool *ILynexV3PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.contract.Transact(opts, method, params...)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address virtualPool)
func (_ILynexV3Pool *ILynexV3PoolCaller) ActiveIncentive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "activeIncentive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address virtualPool)
func (_ILynexV3Pool *ILynexV3PoolSession) ActiveIncentive() (common.Address, error) {
	return _ILynexV3Pool.Contract.ActiveIncentive(&_ILynexV3Pool.CallOpts)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address virtualPool)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) ActiveIncentive() (common.Address, error) {
	return _ILynexV3Pool.Contract.ActiveIncentive(&_ILynexV3Pool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCaller) DataStorageOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "dataStorageOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolSession) DataStorageOperator() (common.Address, error) {
	return _ILynexV3Pool.Contract.DataStorageOperator(&_ILynexV3Pool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) DataStorageOperator() (common.Address, error) {
	return _ILynexV3Pool.Contract.DataStorageOperator(&_ILynexV3Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolSession) Factory() (common.Address, error) {
	return _ILynexV3Pool.Contract.Factory(&_ILynexV3Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Factory() (common.Address, error) {
	return _ILynexV3Pool.Contract.Factory(&_ILynexV3Pool.CallOpts)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_ILynexV3Pool *ILynexV3PoolCaller) GetInnerCumulatives(opts *bind.CallOpts, bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "getInnerCumulatives", bottomTick, topTick)

	outstruct := new(struct {
		InnerTickCumulative           *big.Int
		InnerSecondsSpentPerLiquidity *big.Int
		InnerSecondsSpent             uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InnerTickCumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpentPerLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpent = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_ILynexV3Pool *ILynexV3PoolSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _ILynexV3Pool.Contract.GetInnerCumulatives(&_ILynexV3Pool.CallOpts, bottomTick, topTick)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _ILynexV3Pool.Contract.GetInnerCumulatives(&_ILynexV3Pool.CallOpts, bottomTick, topTick)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_ILynexV3Pool *ILynexV3PoolCaller) GetTimepoints(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "getTimepoints", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                []*big.Int
		SecondsPerLiquidityCumulatives []*big.Int
		VolatilityCumulatives          []*big.Int
		VolumePerAvgLiquiditys         []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulatives = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolatilityCumulatives = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolumePerAvgLiquiditys = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_ILynexV3Pool *ILynexV3PoolSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _ILynexV3Pool.Contract.GetTimepoints(&_ILynexV3Pool.CallOpts, secondsAgos)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _ILynexV3Pool.Contract.GetTimepoints(&_ILynexV3Pool.CallOpts, secondsAgos)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint16 communityFeeToken0, uint16 communityFeeToken1, bool unlocked)
func (_ILynexV3Pool *ILynexV3PoolCaller) GlobalState(opts *bind.CallOpts) (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint16
	CommunityFeeToken1 uint16
	Unlocked           bool
}, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "globalState")

	outstruct := new(struct {
		Price              *big.Int
		Tick               *big.Int
		Fee                uint16
		TimepointIndex     uint16
		CommunityFeeToken0 uint16
		CommunityFeeToken1 uint16
		Unlocked           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TimepointIndex = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken0 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken1 = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint16 communityFeeToken0, uint16 communityFeeToken1, bool unlocked)
func (_ILynexV3Pool *ILynexV3PoolSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint16
	CommunityFeeToken1 uint16
	Unlocked           bool
}, error) {
	return _ILynexV3Pool.Contract.GlobalState(&_ILynexV3Pool.CallOpts)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint16 communityFeeToken0, uint16 communityFeeToken1, bool unlocked)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint16
	CommunityFeeToken1 uint16
	Unlocked           bool
}, error) {
	return _ILynexV3Pool.Contract.GlobalState(&_ILynexV3Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_ILynexV3Pool *ILynexV3PoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_ILynexV3Pool *ILynexV3PoolSession) Liquidity() (*big.Int, error) {
	return _ILynexV3Pool.Contract.Liquidity(&_ILynexV3Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Liquidity() (*big.Int, error) {
	return _ILynexV3Pool.Contract.Liquidity(&_ILynexV3Pool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32 cooldownInSeconds)
func (_ILynexV3Pool *ILynexV3PoolCaller) LiquidityCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "liquidityCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32 cooldownInSeconds)
func (_ILynexV3Pool *ILynexV3PoolSession) LiquidityCooldown() (uint32, error) {
	return _ILynexV3Pool.Contract.LiquidityCooldown(&_ILynexV3Pool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32 cooldownInSeconds)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) LiquidityCooldown() (uint32, error) {
	return _ILynexV3Pool.Contract.LiquidityCooldown(&_ILynexV3Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_ILynexV3Pool *ILynexV3PoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_ILynexV3Pool *ILynexV3PoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _ILynexV3Pool.Contract.MaxLiquidityPerTick(&_ILynexV3Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _ILynexV3Pool.Contract.MaxLiquidityPerTick(&_ILynexV3Pool.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 liquidityAmount, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_ILynexV3Pool *ILynexV3PoolCaller) Positions(opts *bind.CallOpts, key [32]byte) (struct {
	LiquidityAmount           *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "positions", key)

	outstruct := new(struct {
		LiquidityAmount           *big.Int
		LastLiquidityAddTimestamp uint32
		InnerFeeGrowth0Token      *big.Int
		InnerFeeGrowth1Token      *big.Int
		Fees0                     *big.Int
		Fees1                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastLiquidityAddTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.InnerFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InnerFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 liquidityAmount, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_ILynexV3Pool *ILynexV3PoolSession) Positions(key [32]byte) (struct {
	LiquidityAmount           *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _ILynexV3Pool.Contract.Positions(&_ILynexV3Pool.CallOpts, key)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 liquidityAmount, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Positions(key [32]byte) (struct {
	LiquidityAmount           *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _ILynexV3Pool.Contract.Positions(&_ILynexV3Pool.CallOpts, key)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_ILynexV3Pool *ILynexV3PoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_ILynexV3Pool *ILynexV3PoolSession) TickSpacing() (*big.Int, error) {
	return _ILynexV3Pool.Contract.TickSpacing(&_ILynexV3Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) TickSpacing() (*big.Int, error) {
	return _ILynexV3Pool.Contract.TickSpacing(&_ILynexV3Pool.CallOpts)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 wordPosition) view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolCaller) TickTable(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "tickTable", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 wordPosition) view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolSession) TickTable(wordPosition int16) (*big.Int, error) {
	return _ILynexV3Pool.Contract.TickTable(&_ILynexV3Pool.CallOpts, wordPosition)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 wordPosition) view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) TickTable(wordPosition int16) (*big.Int, error) {
	return _ILynexV3Pool.Contract.TickTable(&_ILynexV3Pool.CallOpts, wordPosition)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_ILynexV3Pool *ILynexV3PoolCaller) Ticks(opts *bind.CallOpts, tick *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "ticks", tick)

	outstruct := new(struct {
		LiquidityTotal           *big.Int
		LiquidityDelta           *big.Int
		OuterFeeGrowth0Token     *big.Int
		OuterFeeGrowth1Token     *big.Int
		OuterTickCumulative      *big.Int
		OuterSecondsPerLiquidity *big.Int
		OuterSecondsSpent        uint32
		Initialized              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OuterTickCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsPerLiquidity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsSpent = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_ILynexV3Pool *ILynexV3PoolSession) Ticks(tick *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _ILynexV3Pool.Contract.Ticks(&_ILynexV3Pool.CallOpts, tick)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Ticks(tick *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _ILynexV3Pool.Contract.Ticks(&_ILynexV3Pool.CallOpts, tick)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_ILynexV3Pool *ILynexV3PoolCaller) Timepoints(opts *bind.CallOpts, index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "timepoints", index)

	outstruct := new(struct {
		Initialized                   bool
		BlockTimestamp                uint32
		TickCumulative                *big.Int
		SecondsPerLiquidityCumulative *big.Int
		VolatilityCumulative          *big.Int
		AverageTick                   *big.Int
		VolumePerLiquidityCumulative  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulative = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VolatilityCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AverageTick = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VolumePerLiquidityCumulative = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_ILynexV3Pool *ILynexV3PoolSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _ILynexV3Pool.Contract.Timepoints(&_ILynexV3Pool.CallOpts, index)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _ILynexV3Pool.Contract.Timepoints(&_ILynexV3Pool.CallOpts, index)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolSession) Token0() (common.Address, error) {
	return _ILynexV3Pool.Contract.Token0(&_ILynexV3Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Token0() (common.Address, error) {
	return _ILynexV3Pool.Contract.Token0(&_ILynexV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolSession) Token1() (common.Address, error) {
	return _ILynexV3Pool.Contract.Token1(&_ILynexV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) Token1() (common.Address, error) {
	return _ILynexV3Pool.Contract.Token1(&_ILynexV3Pool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolCaller) TotalFeeGrowth0Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "totalFeeGrowth0Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _ILynexV3Pool.Contract.TotalFeeGrowth0Token(&_ILynexV3Pool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _ILynexV3Pool.Contract.TotalFeeGrowth0Token(&_ILynexV3Pool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolCaller) TotalFeeGrowth1Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexV3Pool.contract.Call(opts, &out, "totalFeeGrowth1Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _ILynexV3Pool.Contract.TotalFeeGrowth1Token(&_ILynexV3Pool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_ILynexV3Pool *ILynexV3PoolCallerSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _ILynexV3Pool.Contract.TotalFeeGrowth1Token(&_ILynexV3Pool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactor) Burn(opts *bind.TransactOpts, bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "burn", bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Burn(&_ILynexV3Pool.TransactOpts, bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Burn(&_ILynexV3Pool.TransactOpts, bottomTick, topTick, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "collect", recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_ILynexV3Pool *ILynexV3PoolSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Collect(&_ILynexV3Pool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Collect(&_ILynexV3Pool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_ILynexV3Pool *ILynexV3PoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Flash(&_ILynexV3Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Flash(&_ILynexV3Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 price) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactor) Initialize(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "initialize", price)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 price) returns()
func (_ILynexV3Pool *ILynexV3PoolSession) Initialize(price *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Initialize(&_ILynexV3Pool.TransactOpts, price)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 price) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) Initialize(price *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Initialize(&_ILynexV3Pool.TransactOpts, price)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_ILynexV3Pool *ILynexV3PoolTransactor) Mint(opts *bind.TransactOpts, sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "mint", sender, recipient, bottomTick, topTick, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_ILynexV3Pool *ILynexV3PoolSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Mint(&_ILynexV3Pool.TransactOpts, sender, recipient, bottomTick, topTick, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Mint(&_ILynexV3Pool.TransactOpts, sender, recipient, bottomTick, topTick, amount, data)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xf0b9cf59.
//
// Solidity: function setCommunityFee(uint16 communityFee0, uint16 communityFee1) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactor) SetCommunityFee(opts *bind.TransactOpts, communityFee0 uint16, communityFee1 uint16) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "setCommunityFee", communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xf0b9cf59.
//
// Solidity: function setCommunityFee(uint16 communityFee0, uint16 communityFee1) returns()
func (_ILynexV3Pool *ILynexV3PoolSession) SetCommunityFee(communityFee0 uint16, communityFee1 uint16) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetCommunityFee(&_ILynexV3Pool.TransactOpts, communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xf0b9cf59.
//
// Solidity: function setCommunityFee(uint16 communityFee0, uint16 communityFee1) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) SetCommunityFee(communityFee0 uint16, communityFee1 uint16) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetCommunityFee(&_ILynexV3Pool.TransactOpts, communityFee0, communityFee1)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactor) SetIncentive(opts *bind.TransactOpts, virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "setIncentive", virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_ILynexV3Pool *ILynexV3PoolSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetIncentive(&_ILynexV3Pool.TransactOpts, virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetIncentive(&_ILynexV3Pool.TransactOpts, virtualPoolAddress)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactor) SetLiquidityCooldown(opts *bind.TransactOpts, newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "setLiquidityCooldown", newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_ILynexV3Pool *ILynexV3PoolSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetLiquidityCooldown(&_ILynexV3Pool.TransactOpts, newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetLiquidityCooldown(&_ILynexV3Pool.TransactOpts, newLiquidityCooldown)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactor) SetTickSpacing(opts *bind.TransactOpts, newTickSpacing *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "setTickSpacing", newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_ILynexV3Pool *ILynexV3PoolSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetTickSpacing(&_ILynexV3Pool.TransactOpts, newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SetTickSpacing(&_ILynexV3Pool.TransactOpts, newTickSpacing)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "swap", recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_ILynexV3Pool *ILynexV3PoolSession) Swap(recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Swap(&_ILynexV3Pool.TransactOpts, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) Swap(recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.Swap(&_ILynexV3Pool.TransactOpts, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactor) SwapSupportingFeeOnInputTokens(opts *bind.TransactOpts, sender common.Address, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.contract.Transact(opts, "swapSupportingFeeOnInputTokens", sender, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_ILynexV3Pool *ILynexV3PoolSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SwapSupportingFeeOnInputTokens(&_ILynexV3Pool.TransactOpts, sender, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountSpecified, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_ILynexV3Pool *ILynexV3PoolTransactorSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountSpecified *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _ILynexV3Pool.Contract.SwapSupportingFeeOnInputTokens(&_ILynexV3Pool.TransactOpts, sender, recipient, zeroToOne, amountSpecified, limitSqrtPrice, data)
}

// ILynexV3PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ILynexV3Pool contract.
type ILynexV3PoolBurnIterator struct {
	Event *ILynexV3PoolBurn // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolBurn)
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
		it.Event = new(ILynexV3PoolBurn)
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
func (it *ILynexV3PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolBurn represents a Burn event raised by the ILynexV3Pool contract.
type ILynexV3PoolBurn struct {
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*ILynexV3PoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolBurnIterator{contract: _ILynexV3Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolBurn, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolBurn)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseBurn(log types.Log) (*ILynexV3PoolBurn, error) {
	event := new(ILynexV3PoolBurn)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the ILynexV3Pool contract.
type ILynexV3PoolCollectIterator struct {
	Event *ILynexV3PoolCollect // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolCollect)
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
		it.Event = new(ILynexV3PoolCollect)
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
func (it *ILynexV3PoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolCollect represents a Collect event raised by the ILynexV3Pool contract.
type ILynexV3PoolCollect struct {
	Owner      common.Address
	Recipient  common.Address
	BottomTick *big.Int
	TopTick    *big.Int
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*ILynexV3PoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolCollectIterator{contract: _ILynexV3Pool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolCollect, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolCollect)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseCollect(log types.Log) (*ILynexV3PoolCollect, error) {
	event := new(ILynexV3PoolCollect)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolCommunityFeeIterator is returned from FilterCommunityFee and is used to iterate over the raw logs and unpacked data for CommunityFee events raised by the ILynexV3Pool contract.
type ILynexV3PoolCommunityFeeIterator struct {
	Event *ILynexV3PoolCommunityFee // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolCommunityFee)
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
		it.Event = new(ILynexV3PoolCommunityFee)
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
func (it *ILynexV3PoolCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolCommunityFee represents a CommunityFee event raised by the ILynexV3Pool contract.
type ILynexV3PoolCommunityFee struct {
	CommunityFee0New uint16
	CommunityFee1New uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommunityFee is a free log retrieval operation binding the contract event 0x370966829959865419a97fc8708e1d348a92142c2cfec7299e264677970174bc.
//
// Solidity: event CommunityFee(uint16 communityFee0New, uint16 communityFee1New)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterCommunityFee(opts *bind.FilterOpts) (*ILynexV3PoolCommunityFeeIterator, error) {

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolCommunityFeeIterator{contract: _ILynexV3Pool.contract, event: "CommunityFee", logs: logs, sub: sub}, nil
}

// WatchCommunityFee is a free log subscription operation binding the contract event 0x370966829959865419a97fc8708e1d348a92142c2cfec7299e264677970174bc.
//
// Solidity: event CommunityFee(uint16 communityFee0New, uint16 communityFee1New)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchCommunityFee(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolCommunityFee) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolCommunityFee)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
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

// ParseCommunityFee is a log parse operation binding the contract event 0x370966829959865419a97fc8708e1d348a92142c2cfec7299e264677970174bc.
//
// Solidity: event CommunityFee(uint16 communityFee0New, uint16 communityFee1New)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseCommunityFee(log types.Log) (*ILynexV3PoolCommunityFee, error) {
	event := new(ILynexV3PoolCommunityFee)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the ILynexV3Pool contract.
type ILynexV3PoolFeeIterator struct {
	Event *ILynexV3PoolFee // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolFee)
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
		it.Event = new(ILynexV3PoolFee)
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
func (it *ILynexV3PoolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolFee represents a Fee event raised by the ILynexV3Pool contract.
type ILynexV3PoolFee struct {
	Fee uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterFee(opts *bind.FilterOpts) (*ILynexV3PoolFeeIterator, error) {

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolFeeIterator{contract: _ILynexV3Pool.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolFee) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolFee)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseFee(log types.Log) (*ILynexV3PoolFee, error) {
	event := new(ILynexV3PoolFee)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the ILynexV3Pool contract.
type ILynexV3PoolFlashIterator struct {
	Event *ILynexV3PoolFlash // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolFlash)
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
		it.Event = new(ILynexV3PoolFlash)
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
func (it *ILynexV3PoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolFlash represents a Flash event raised by the ILynexV3Pool contract.
type ILynexV3PoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*ILynexV3PoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolFlashIterator{contract: _ILynexV3Pool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolFlash)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseFlash(log types.Log) (*ILynexV3PoolFlash, error) {
	event := new(ILynexV3PoolFlash)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolIncentiveIterator is returned from FilterIncentive and is used to iterate over the raw logs and unpacked data for Incentive events raised by the ILynexV3Pool contract.
type ILynexV3PoolIncentiveIterator struct {
	Event *ILynexV3PoolIncentive // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolIncentive)
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
		it.Event = new(ILynexV3PoolIncentive)
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
func (it *ILynexV3PoolIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolIncentive represents a Incentive event raised by the ILynexV3Pool contract.
type ILynexV3PoolIncentive struct {
	VirtualPoolAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterIncentive is a free log retrieval operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterIncentive(opts *bind.FilterOpts, virtualPoolAddress []common.Address) (*ILynexV3PoolIncentiveIterator, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolIncentiveIterator{contract: _ILynexV3Pool.contract, event: "Incentive", logs: logs, sub: sub}, nil
}

// WatchIncentive is a free log subscription operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchIncentive(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolIncentive, virtualPoolAddress []common.Address) (event.Subscription, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolIncentive)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Incentive", log); err != nil {
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

// ParseIncentive is a log parse operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseIncentive(log types.Log) (*ILynexV3PoolIncentive, error) {
	event := new(ILynexV3PoolIncentive)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Incentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the ILynexV3Pool contract.
type ILynexV3PoolInitializeIterator struct {
	Event *ILynexV3PoolInitialize // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolInitialize)
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
		it.Event = new(ILynexV3PoolInitialize)
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
func (it *ILynexV3PoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolInitialize represents a Initialize event raised by the ILynexV3Pool contract.
type ILynexV3PoolInitialize struct {
	Price *big.Int
	Tick  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*ILynexV3PoolInitializeIterator, error) {

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolInitializeIterator{contract: _ILynexV3Pool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolInitialize) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolInitialize)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseInitialize(log types.Log) (*ILynexV3PoolInitialize, error) {
	event := new(ILynexV3PoolInitialize)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolLiquidityCooldownIterator is returned from FilterLiquidityCooldown and is used to iterate over the raw logs and unpacked data for LiquidityCooldown events raised by the ILynexV3Pool contract.
type ILynexV3PoolLiquidityCooldownIterator struct {
	Event *ILynexV3PoolLiquidityCooldown // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolLiquidityCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolLiquidityCooldown)
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
		it.Event = new(ILynexV3PoolLiquidityCooldown)
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
func (it *ILynexV3PoolLiquidityCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolLiquidityCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolLiquidityCooldown represents a LiquidityCooldown event raised by the ILynexV3Pool contract.
type ILynexV3PoolLiquidityCooldown struct {
	LiquidityCooldown uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityCooldown is a free log retrieval operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterLiquidityCooldown(opts *bind.FilterOpts) (*ILynexV3PoolLiquidityCooldownIterator, error) {

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolLiquidityCooldownIterator{contract: _ILynexV3Pool.contract, event: "LiquidityCooldown", logs: logs, sub: sub}, nil
}

// WatchLiquidityCooldown is a free log subscription operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchLiquidityCooldown(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolLiquidityCooldown) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolLiquidityCooldown)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
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

// ParseLiquidityCooldown is a log parse operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseLiquidityCooldown(log types.Log) (*ILynexV3PoolLiquidityCooldown, error) {
	event := new(ILynexV3PoolLiquidityCooldown)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ILynexV3Pool contract.
type ILynexV3PoolMintIterator struct {
	Event *ILynexV3PoolMint // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolMint)
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
		it.Event = new(ILynexV3PoolMint)
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
func (it *ILynexV3PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolMint represents a Mint event raised by the ILynexV3Pool contract.
type ILynexV3PoolMint struct {
	Sender          common.Address
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*ILynexV3PoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolMintIterator{contract: _ILynexV3Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolMint, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolMint)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseMint(log types.Log) (*ILynexV3PoolMint, error) {
	event := new(ILynexV3PoolMint)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ILynexV3Pool contract.
type ILynexV3PoolSwapIterator struct {
	Event *ILynexV3PoolSwap // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolSwap)
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
		it.Event = new(ILynexV3PoolSwap)
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
func (it *ILynexV3PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolSwap represents a Swap event raised by the ILynexV3Pool contract.
type ILynexV3PoolSwap struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Price     *big.Int
	Liquidity *big.Int
	Tick      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*ILynexV3PoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolSwapIterator{contract: _ILynexV3Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolSwap)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseSwap(log types.Log) (*ILynexV3PoolSwap, error) {
	event := new(ILynexV3PoolSwap)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3PoolTickSpacingIterator is returned from FilterTickSpacing and is used to iterate over the raw logs and unpacked data for TickSpacing events raised by the ILynexV3Pool contract.
type ILynexV3PoolTickSpacingIterator struct {
	Event *ILynexV3PoolTickSpacing // Event containing the contract specifics and raw log

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
func (it *ILynexV3PoolTickSpacingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3PoolTickSpacing)
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
		it.Event = new(ILynexV3PoolTickSpacing)
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
func (it *ILynexV3PoolTickSpacingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3PoolTickSpacingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3PoolTickSpacing represents a TickSpacing event raised by the ILynexV3Pool contract.
type ILynexV3PoolTickSpacing struct {
	NewTickSpacing *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTickSpacing is a free log retrieval operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_ILynexV3Pool *ILynexV3PoolFilterer) FilterTickSpacing(opts *bind.FilterOpts) (*ILynexV3PoolTickSpacingIterator, error) {

	logs, sub, err := _ILynexV3Pool.contract.FilterLogs(opts, "TickSpacing")
	if err != nil {
		return nil, err
	}
	return &ILynexV3PoolTickSpacingIterator{contract: _ILynexV3Pool.contract, event: "TickSpacing", logs: logs, sub: sub}, nil
}

// WatchTickSpacing is a free log subscription operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_ILynexV3Pool *ILynexV3PoolFilterer) WatchTickSpacing(opts *bind.WatchOpts, sink chan<- *ILynexV3PoolTickSpacing) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Pool.contract.WatchLogs(opts, "TickSpacing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3PoolTickSpacing)
				if err := _ILynexV3Pool.contract.UnpackLog(event, "TickSpacing", log); err != nil {
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

// ParseTickSpacing is a log parse operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_ILynexV3Pool *ILynexV3PoolFilterer) ParseTickSpacing(log types.Log) (*ILynexV3PoolTickSpacing, error) {
	event := new(ILynexV3PoolTickSpacing)
	if err := _ILynexV3Pool.contract.UnpackLog(event, "TickSpacing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
