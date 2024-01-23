// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isyncswap_pool

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

// IPoolTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type IPoolTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// ISyncSwapPoolMetaData contains all meta data concerning the ISyncSwapPool contract.
var ISyncSwapPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"burnSingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"protocolFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"permit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISyncSwapPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ISyncSwapPoolMetaData.ABI instead.
var ISyncSwapPoolABI = ISyncSwapPoolMetaData.ABI

// ISyncSwapPool is an auto generated Go binding around an Ethereum contract.
type ISyncSwapPool struct {
	ISyncSwapPoolCaller     // Read-only binding to the contract
	ISyncSwapPoolTransactor // Write-only binding to the contract
	ISyncSwapPoolFilterer   // Log filterer for contract events
}

// ISyncSwapPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISyncSwapPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISyncSwapPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISyncSwapPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISyncSwapPoolSession struct {
	Contract     *ISyncSwapPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISyncSwapPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISyncSwapPoolCallerSession struct {
	Contract *ISyncSwapPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISyncSwapPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISyncSwapPoolTransactorSession struct {
	Contract     *ISyncSwapPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISyncSwapPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISyncSwapPoolRaw struct {
	Contract *ISyncSwapPool // Generic contract binding to access the raw methods on
}

// ISyncSwapPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISyncSwapPoolCallerRaw struct {
	Contract *ISyncSwapPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ISyncSwapPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISyncSwapPoolTransactorRaw struct {
	Contract *ISyncSwapPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISyncSwapPool creates a new instance of ISyncSwapPool, bound to a specific deployed contract.
func NewISyncSwapPool(address common.Address, backend bind.ContractBackend) (*ISyncSwapPool, error) {
	contract, err := bindISyncSwapPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPool{ISyncSwapPoolCaller: ISyncSwapPoolCaller{contract: contract}, ISyncSwapPoolTransactor: ISyncSwapPoolTransactor{contract: contract}, ISyncSwapPoolFilterer: ISyncSwapPoolFilterer{contract: contract}}, nil
}

// NewISyncSwapPoolCaller creates a new read-only instance of ISyncSwapPool, bound to a specific deployed contract.
func NewISyncSwapPoolCaller(address common.Address, caller bind.ContractCaller) (*ISyncSwapPoolCaller, error) {
	contract, err := bindISyncSwapPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolCaller{contract: contract}, nil
}

// NewISyncSwapPoolTransactor creates a new write-only instance of ISyncSwapPool, bound to a specific deployed contract.
func NewISyncSwapPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ISyncSwapPoolTransactor, error) {
	contract, err := bindISyncSwapPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolTransactor{contract: contract}, nil
}

// NewISyncSwapPoolFilterer creates a new log filterer instance of ISyncSwapPool, bound to a specific deployed contract.
func NewISyncSwapPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ISyncSwapPoolFilterer, error) {
	contract, err := bindISyncSwapPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolFilterer{contract: contract}, nil
}

// bindISyncSwapPool binds a generic wrapper to an already deployed contract.
func bindISyncSwapPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISyncSwapPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapPool *ISyncSwapPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapPool.Contract.ISyncSwapPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapPool *ISyncSwapPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.ISyncSwapPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapPool *ISyncSwapPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.ISyncSwapPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapPool *ISyncSwapPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapPool *ISyncSwapPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapPool *ISyncSwapPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ISyncSwapPool *ISyncSwapPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ISyncSwapPool *ISyncSwapPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ISyncSwapPool.Contract.DOMAINSEPARATOR(&_ISyncSwapPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ISyncSwapPool.Contract.DOMAINSEPARATOR(&_ISyncSwapPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.Allowance(&_ISyncSwapPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.Allowance(&_ISyncSwapPool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.BalanceOf(&_ISyncSwapPool.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.BalanceOf(&_ISyncSwapPool.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ISyncSwapPool *ISyncSwapPoolSession) Decimals() (uint8, error) {
	return _ISyncSwapPool.Contract.Decimals(&_ISyncSwapPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Decimals() (uint8, error) {
	return _ISyncSwapPool.Contract.Decimals(&_ISyncSwapPool.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address tokenOut, uint256 amountOut, address sender) view returns(uint256 amountIn)
func (_ISyncSwapPool *ISyncSwapPoolCaller) GetAmountIn(opts *bind.CallOpts, tokenOut common.Address, amountOut *big.Int, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "getAmountIn", tokenOut, amountOut, sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address tokenOut, uint256 amountOut, address sender) view returns(uint256 amountIn)
func (_ISyncSwapPool *ISyncSwapPoolSession) GetAmountIn(tokenOut common.Address, amountOut *big.Int, sender common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetAmountIn(&_ISyncSwapPool.CallOpts, tokenOut, amountOut, sender)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address tokenOut, uint256 amountOut, address sender) view returns(uint256 amountIn)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) GetAmountIn(tokenOut common.Address, amountOut *big.Int, sender common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetAmountIn(&_ISyncSwapPool.CallOpts, tokenOut, amountOut, sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address tokenIn, uint256 amountIn, address sender) view returns(uint256 amountOut)
func (_ISyncSwapPool *ISyncSwapPoolCaller) GetAmountOut(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "getAmountOut", tokenIn, amountIn, sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address tokenIn, uint256 amountIn, address sender) view returns(uint256 amountOut)
func (_ISyncSwapPool *ISyncSwapPoolSession) GetAmountOut(tokenIn common.Address, amountIn *big.Int, sender common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetAmountOut(&_ISyncSwapPool.CallOpts, tokenIn, amountIn, sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address tokenIn, uint256 amountIn, address sender) view returns(uint256 amountOut)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) GetAmountOut(tokenIn common.Address, amountIn *big.Int, sender common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetAmountOut(&_ISyncSwapPool.CallOpts, tokenIn, amountIn, sender)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_ISyncSwapPool *ISyncSwapPoolCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_ISyncSwapPool *ISyncSwapPoolSession) GetAssets() ([]common.Address, error) {
	return _ISyncSwapPool.Contract.GetAssets(&_ISyncSwapPool.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) GetAssets() ([]common.Address, error) {
	return _ISyncSwapPool.Contract.GetAssets(&_ISyncSwapPool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 protocolFee)
func (_ISyncSwapPool *ISyncSwapPoolCaller) GetProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "getProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 protocolFee)
func (_ISyncSwapPool *ISyncSwapPoolSession) GetProtocolFee() (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetProtocolFee(&_ISyncSwapPool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 protocolFee)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) GetProtocolFee() (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetProtocolFee(&_ISyncSwapPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) GetReserves(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) GetReserves() (*big.Int, *big.Int, error) {
	return _ISyncSwapPool.Contract.GetReserves(&_ISyncSwapPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) GetReserves() (*big.Int, *big.Int, error) {
	return _ISyncSwapPool.Contract.GetReserves(&_ISyncSwapPool.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ISyncSwapPool *ISyncSwapPoolCaller) GetSwapFee(opts *bind.CallOpts, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "getSwapFee", sender, tokenIn, tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ISyncSwapPool *ISyncSwapPoolSession) GetSwapFee(sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetSwapFee(&_ISyncSwapPool.CallOpts, sender, tokenIn, tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) GetSwapFee(sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ISyncSwapPool.Contract.GetSwapFee(&_ISyncSwapPool.CallOpts, sender, tokenIn, tokenOut, data)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) InvariantLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "invariantLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) InvariantLast() (*big.Int, error) {
	return _ISyncSwapPool.Contract.InvariantLast(&_ISyncSwapPool.CallOpts)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) InvariantLast() (*big.Int, error) {
	return _ISyncSwapPool.Contract.InvariantLast(&_ISyncSwapPool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolSession) Master() (common.Address, error) {
	return _ISyncSwapPool.Contract.Master(&_ISyncSwapPool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Master() (common.Address, error) {
	return _ISyncSwapPool.Contract.Master(&_ISyncSwapPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ISyncSwapPool *ISyncSwapPoolSession) Name() (string, error) {
	return _ISyncSwapPool.Contract.Name(&_ISyncSwapPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Name() (string, error) {
	return _ISyncSwapPool.Contract.Name(&_ISyncSwapPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.Nonces(&_ISyncSwapPool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ISyncSwapPool.Contract.Nonces(&_ISyncSwapPool.CallOpts, owner)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_ISyncSwapPool *ISyncSwapPoolCaller) PoolType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_ISyncSwapPool *ISyncSwapPoolSession) PoolType() (uint16, error) {
	return _ISyncSwapPool.Contract.PoolType(&_ISyncSwapPool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) PoolType() (uint16, error) {
	return _ISyncSwapPool.Contract.PoolType(&_ISyncSwapPool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) Reserve0() (*big.Int, error) {
	return _ISyncSwapPool.Contract.Reserve0(&_ISyncSwapPool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Reserve0() (*big.Int, error) {
	return _ISyncSwapPool.Contract.Reserve0(&_ISyncSwapPool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) Reserve1() (*big.Int, error) {
	return _ISyncSwapPool.Contract.Reserve1(&_ISyncSwapPool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Reserve1() (*big.Int, error) {
	return _ISyncSwapPool.Contract.Reserve1(&_ISyncSwapPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ISyncSwapPool *ISyncSwapPoolSession) Symbol() (string, error) {
	return _ISyncSwapPool.Contract.Symbol(&_ISyncSwapPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Symbol() (string, error) {
	return _ISyncSwapPool.Contract.Symbol(&_ISyncSwapPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolSession) Token0() (common.Address, error) {
	return _ISyncSwapPool.Contract.Token0(&_ISyncSwapPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Token0() (common.Address, error) {
	return _ISyncSwapPool.Contract.Token0(&_ISyncSwapPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolSession) Token1() (common.Address, error) {
	return _ISyncSwapPool.Contract.Token1(&_ISyncSwapPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Token1() (common.Address, error) {
	return _ISyncSwapPool.Contract.Token1(&_ISyncSwapPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolSession) TotalSupply() (*big.Int, error) {
	return _ISyncSwapPool.Contract.TotalSupply(&_ISyncSwapPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _ISyncSwapPool.Contract.TotalSupply(&_ISyncSwapPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapPool.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolSession) Vault() (common.Address, error) {
	return _ISyncSwapPool.Contract.Vault(&_ISyncSwapPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ISyncSwapPool *ISyncSwapPoolCallerSession) Vault() (common.Address, error) {
	return _ISyncSwapPool.Contract.Vault(&_ISyncSwapPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Approve(&_ISyncSwapPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Approve(&_ISyncSwapPool.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256)[] tokenAmounts)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Burn(opts *bind.TransactOpts, data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "burn", data, sender, callback, callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256)[] tokenAmounts)
func (_ISyncSwapPool *ISyncSwapPoolSession) Burn(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Burn(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256)[] tokenAmounts)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Burn(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Burn(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256) tokenAmount)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) BurnSingle(opts *bind.TransactOpts, data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "burnSingle", data, sender, callback, callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256) tokenAmount)
func (_ISyncSwapPool *ISyncSwapPoolSession) BurnSingle(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.BurnSingle(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256) tokenAmount)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) BurnSingle(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.BurnSingle(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes data, address sender, address callback, bytes callbackData) returns(uint256 liquidity)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Mint(opts *bind.TransactOpts, data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "mint", data, sender, callback, callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes data, address sender, address callback, bytes callbackData) returns(uint256 liquidity)
func (_ISyncSwapPool *ISyncSwapPoolSession) Mint(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Mint(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes data, address sender, address callback, bytes callbackData) returns(uint256 liquidity)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Mint(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Mint(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ISyncSwapPool *ISyncSwapPoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Permit(&_ISyncSwapPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Permit(&_ISyncSwapPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address owner, address spender, uint256 amount, uint256 deadline, bytes signature) returns()
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Permit2(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "permit2", owner, spender, amount, deadline, signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address owner, address spender, uint256 amount, uint256 deadline, bytes signature) returns()
func (_ISyncSwapPool *ISyncSwapPoolSession) Permit2(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Permit2(&_ISyncSwapPool.TransactOpts, owner, spender, amount, deadline, signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address owner, address spender, uint256 amount, uint256 deadline, bytes signature) returns()
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Permit2(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Permit2(&_ISyncSwapPool.TransactOpts, owner, spender, amount, deadline, signature)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256) tokenAmount)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Swap(opts *bind.TransactOpts, data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "swap", data, sender, callback, callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256) tokenAmount)
func (_ISyncSwapPool *ISyncSwapPoolSession) Swap(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Swap(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes data, address sender, address callback, bytes callbackData) returns((address,uint256) tokenAmount)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Swap(data []byte, sender common.Address, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Swap(&_ISyncSwapPool.TransactOpts, data, sender, callback, callbackData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Transfer(&_ISyncSwapPool.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.Transfer(&_ISyncSwapPool.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.TransferFrom(&_ISyncSwapPool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ISyncSwapPool *ISyncSwapPoolTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISyncSwapPool.Contract.TransferFrom(&_ISyncSwapPool.TransactOpts, from, to, amount)
}

// ISyncSwapPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ISyncSwapPool contract.
type ISyncSwapPoolApprovalIterator struct {
	Event *ISyncSwapPoolApproval // Event containing the contract specifics and raw log

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
func (it *ISyncSwapPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapPoolApproval)
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
		it.Event = new(ISyncSwapPoolApproval)
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
func (it *ISyncSwapPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapPoolApproval represents a Approval event raised by the ISyncSwapPool contract.
type ISyncSwapPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ISyncSwapPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolApprovalIterator{contract: _ISyncSwapPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ISyncSwapPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapPoolApproval)
				if err := _ISyncSwapPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) ParseApproval(log types.Log) (*ISyncSwapPoolApproval, error) {
	event := new(ISyncSwapPoolApproval)
	if err := _ISyncSwapPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISyncSwapPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ISyncSwapPool contract.
type ISyncSwapPoolBurnIterator struct {
	Event *ISyncSwapPoolBurn // Event containing the contract specifics and raw log

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
func (it *ISyncSwapPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapPoolBurn)
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
		it.Event = new(ISyncSwapPoolBurn)
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
func (it *ISyncSwapPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapPoolBurn represents a Burn event raised by the ISyncSwapPool contract.
type ISyncSwapPoolBurn struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ISyncSwapPoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolBurnIterator{contract: _ISyncSwapPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ISyncSwapPoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapPoolBurn)
				if err := _ISyncSwapPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) ParseBurn(log types.Log) (*ISyncSwapPoolBurn, error) {
	event := new(ISyncSwapPoolBurn)
	if err := _ISyncSwapPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISyncSwapPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ISyncSwapPool contract.
type ISyncSwapPoolMintIterator struct {
	Event *ISyncSwapPoolMint // Event containing the contract specifics and raw log

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
func (it *ISyncSwapPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapPoolMint)
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
		it.Event = new(ISyncSwapPoolMint)
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
func (it *ISyncSwapPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapPoolMint represents a Mint event raised by the ISyncSwapPool contract.
type ISyncSwapPoolMint struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ISyncSwapPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolMintIterator{contract: _ISyncSwapPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ISyncSwapPoolMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapPoolMint)
				if err := _ISyncSwapPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) ParseMint(log types.Log) (*ISyncSwapPoolMint, error) {
	event := new(ISyncSwapPoolMint)
	if err := _ISyncSwapPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISyncSwapPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ISyncSwapPool contract.
type ISyncSwapPoolSwapIterator struct {
	Event *ISyncSwapPoolSwap // Event containing the contract specifics and raw log

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
func (it *ISyncSwapPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapPoolSwap)
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
		it.Event = new(ISyncSwapPoolSwap)
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
func (it *ISyncSwapPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapPoolSwap represents a Swap event raised by the ISyncSwapPool contract.
type ISyncSwapPoolSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ISyncSwapPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolSwapIterator{contract: _ISyncSwapPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ISyncSwapPoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapPoolSwap)
				if err := _ISyncSwapPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) ParseSwap(log types.Log) (*ISyncSwapPoolSwap, error) {
	event := new(ISyncSwapPoolSwap)
	if err := _ISyncSwapPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISyncSwapPoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the ISyncSwapPool contract.
type ISyncSwapPoolSyncIterator struct {
	Event *ISyncSwapPoolSync // Event containing the contract specifics and raw log

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
func (it *ISyncSwapPoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapPoolSync)
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
		it.Event = new(ISyncSwapPoolSync)
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
func (it *ISyncSwapPoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapPoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapPoolSync represents a Sync event raised by the ISyncSwapPool contract.
type ISyncSwapPoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) FilterSync(opts *bind.FilterOpts) (*ISyncSwapPoolSyncIterator, error) {

	logs, sub, err := _ISyncSwapPool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolSyncIterator{contract: _ISyncSwapPool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *ISyncSwapPoolSync) (event.Subscription, error) {

	logs, sub, err := _ISyncSwapPool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapPoolSync)
				if err := _ISyncSwapPool.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) ParseSync(log types.Log) (*ISyncSwapPoolSync, error) {
	event := new(ISyncSwapPoolSync)
	if err := _ISyncSwapPool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISyncSwapPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ISyncSwapPool contract.
type ISyncSwapPoolTransferIterator struct {
	Event *ISyncSwapPoolTransfer // Event containing the contract specifics and raw log

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
func (it *ISyncSwapPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapPoolTransfer)
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
		it.Event = new(ISyncSwapPoolTransfer)
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
func (it *ISyncSwapPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapPoolTransfer represents a Transfer event raised by the ISyncSwapPool contract.
type ISyncSwapPoolTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ISyncSwapPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapPoolTransferIterator{contract: _ISyncSwapPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ISyncSwapPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ISyncSwapPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapPoolTransfer)
				if err := _ISyncSwapPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ISyncSwapPool *ISyncSwapPoolFilterer) ParseTransfer(log types.Log) (*ISyncSwapPoolTransfer, error) {
	event := new(ISyncSwapPoolTransfer)
	if err := _ISyncSwapPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
