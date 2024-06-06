// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilynex_v2_pair

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

// PairObservation is an auto generated low-level Go binding around an user-defined struct.
type PairObservation struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}

// ILynexPairMetaData contains all meta data concerning the ILynexPair contract.
var ILynexPairMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockTimestampLast\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimFees\",\"inputs\":[],\"outputs\":[{\"name\":\"claimed0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable0\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimable1\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"current\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCumulativePrices\",\"inputs\":[],\"outputs\":[{\"name\":\"reserve0Cumulative\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve1Cumulative\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountOut\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReserves\",\"inputs\":[],\"outputs\":[{\"name\":\"_reserve0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_reserve1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_blockTimestampLast\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"index0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"index1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isStable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastObservation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPair.Observation\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve0Cumulative\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve1Cumulative\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metadata\",\"inputs\":[],\"outputs\":[{\"name\":\"dec0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dec1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"r0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"r1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"st\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"t0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"t1\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"liquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"observationLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"observations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve0Cumulative\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve1Cumulative\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prices\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"points\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quote\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"granularity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reserve0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reserve0CumulativeLast\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reserve1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reserve1CumulativeLast\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sample\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"points\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"window\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"skim\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supplyIndex0\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supplyIndex1\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"amount0Out\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1Out\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sync\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fees\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0In\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1In\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount0Out\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1Out\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Sync\",\"inputs\":[{\"name\":\"reserve0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserve1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ILynexPairABI is the input ABI used to generate the binding from.
// Deprecated: Use ILynexPairMetaData.ABI instead.
var ILynexPairABI = ILynexPairMetaData.ABI

// ILynexPair is an auto generated Go binding around an Ethereum contract.
type ILynexPair struct {
	ILynexPairCaller     // Read-only binding to the contract
	ILynexPairTransactor // Write-only binding to the contract
	ILynexPairFilterer   // Log filterer for contract events
}

// ILynexPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILynexPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILynexPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILynexPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILynexPairSession struct {
	Contract     *ILynexPair       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILynexPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILynexPairCallerSession struct {
	Contract *ILynexPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ILynexPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILynexPairTransactorSession struct {
	Contract     *ILynexPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ILynexPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILynexPairRaw struct {
	Contract *ILynexPair // Generic contract binding to access the raw methods on
}

// ILynexPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILynexPairCallerRaw struct {
	Contract *ILynexPairCaller // Generic read-only contract binding to access the raw methods on
}

// ILynexPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILynexPairTransactorRaw struct {
	Contract *ILynexPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILynexPair creates a new instance of ILynexPair, bound to a specific deployed contract.
func NewILynexPair(address common.Address, backend bind.ContractBackend) (*ILynexPair, error) {
	contract, err := bindILynexPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILynexPair{ILynexPairCaller: ILynexPairCaller{contract: contract}, ILynexPairTransactor: ILynexPairTransactor{contract: contract}, ILynexPairFilterer: ILynexPairFilterer{contract: contract}}, nil
}

// NewILynexPairCaller creates a new read-only instance of ILynexPair, bound to a specific deployed contract.
func NewILynexPairCaller(address common.Address, caller bind.ContractCaller) (*ILynexPairCaller, error) {
	contract, err := bindILynexPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexPairCaller{contract: contract}, nil
}

// NewILynexPairTransactor creates a new write-only instance of ILynexPair, bound to a specific deployed contract.
func NewILynexPairTransactor(address common.Address, transactor bind.ContractTransactor) (*ILynexPairTransactor, error) {
	contract, err := bindILynexPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexPairTransactor{contract: contract}, nil
}

// NewILynexPairFilterer creates a new log filterer instance of ILynexPair, bound to a specific deployed contract.
func NewILynexPairFilterer(address common.Address, filterer bind.ContractFilterer) (*ILynexPairFilterer, error) {
	contract, err := bindILynexPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILynexPairFilterer{contract: contract}, nil
}

// bindILynexPair binds a generic wrapper to an already deployed contract.
func bindILynexPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILynexPairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexPair *ILynexPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexPair.Contract.ILynexPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexPair *ILynexPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexPair.Contract.ILynexPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexPair *ILynexPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexPair.Contract.ILynexPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexPair *ILynexPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexPair *ILynexPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexPair *ILynexPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexPair.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Allowance(&_ILynexPair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Allowance(&_ILynexPair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.BalanceOf(&_ILynexPair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.BalanceOf(&_ILynexPair.CallOpts, arg0)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) BlockTimestampLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "blockTimestampLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint256)
func (_ILynexPair *ILynexPairSession) BlockTimestampLast() (*big.Int, error) {
	return _ILynexPair.Contract.BlockTimestampLast(&_ILynexPair.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) BlockTimestampLast() (*big.Int, error) {
	return _ILynexPair.Contract.BlockTimestampLast(&_ILynexPair.CallOpts)
}

// Claimable0 is a free data retrieval call binding the contract method 0x4d5a9f8a.
//
// Solidity: function claimable0(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Claimable0(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "claimable0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable0 is a free data retrieval call binding the contract method 0x4d5a9f8a.
//
// Solidity: function claimable0(address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) Claimable0(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Claimable0(&_ILynexPair.CallOpts, arg0)
}

// Claimable0 is a free data retrieval call binding the contract method 0x4d5a9f8a.
//
// Solidity: function claimable0(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Claimable0(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Claimable0(&_ILynexPair.CallOpts, arg0)
}

// Claimable1 is a free data retrieval call binding the contract method 0xa1ac4d13.
//
// Solidity: function claimable1(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Claimable1(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "claimable1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable1 is a free data retrieval call binding the contract method 0xa1ac4d13.
//
// Solidity: function claimable1(address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) Claimable1(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Claimable1(&_ILynexPair.CallOpts, arg0)
}

// Claimable1 is a free data retrieval call binding the contract method 0xa1ac4d13.
//
// Solidity: function claimable1(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Claimable1(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Claimable1(&_ILynexPair.CallOpts, arg0)
}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_ILynexPair *ILynexPairCaller) Current(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "current", tokenIn, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_ILynexPair *ILynexPairSession) Current(tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _ILynexPair.Contract.Current(&_ILynexPair.CallOpts, tokenIn, amountIn)
}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_ILynexPair *ILynexPairCallerSession) Current(tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _ILynexPair.Contract.Current(&_ILynexPair.CallOpts, tokenIn, amountIn)
}

// CurrentCumulativePrices is a free data retrieval call binding the contract method 0x1df8c717.
//
// Solidity: function currentCumulativePrices() view returns(uint256 reserve0Cumulative, uint256 reserve1Cumulative, uint256 blockTimestamp)
func (_ILynexPair *ILynexPairCaller) CurrentCumulativePrices(opts *bind.CallOpts) (struct {
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
	BlockTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "currentCumulativePrices")

	outstruct := new(struct {
		Reserve0Cumulative *big.Int
		Reserve1Cumulative *big.Int
		BlockTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0Cumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1Cumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentCumulativePrices is a free data retrieval call binding the contract method 0x1df8c717.
//
// Solidity: function currentCumulativePrices() view returns(uint256 reserve0Cumulative, uint256 reserve1Cumulative, uint256 blockTimestamp)
func (_ILynexPair *ILynexPairSession) CurrentCumulativePrices() (struct {
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
	BlockTimestamp     *big.Int
}, error) {
	return _ILynexPair.Contract.CurrentCumulativePrices(&_ILynexPair.CallOpts)
}

// CurrentCumulativePrices is a free data retrieval call binding the contract method 0x1df8c717.
//
// Solidity: function currentCumulativePrices() view returns(uint256 reserve0Cumulative, uint256 reserve1Cumulative, uint256 blockTimestamp)
func (_ILynexPair *ILynexPairCallerSession) CurrentCumulativePrices() (struct {
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
	BlockTimestamp     *big.Int
}, error) {
	return _ILynexPair.Contract.CurrentCumulativePrices(&_ILynexPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ILynexPair *ILynexPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ILynexPair *ILynexPairSession) Decimals() (uint8, error) {
	return _ILynexPair.Contract.Decimals(&_ILynexPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ILynexPair *ILynexPairCallerSession) Decimals() (uint8, error) {
	return _ILynexPair.Contract.Decimals(&_ILynexPair.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(address)
func (_ILynexPair *ILynexPairCaller) Fees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(address)
func (_ILynexPair *ILynexPairSession) Fees() (common.Address, error) {
	return _ILynexPair.Contract.Fees(&_ILynexPair.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(address)
func (_ILynexPair *ILynexPairCallerSession) Fees() (common.Address, error) {
	return _ILynexPair.Contract.Fees(&_ILynexPair.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_ILynexPair *ILynexPairSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.GetAmountOut(&_ILynexPair.CallOpts, amountIn, tokenIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.GetAmountOut(&_ILynexPair.CallOpts, amountIn, tokenIn)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_ILynexPair *ILynexPairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_ILynexPair *ILynexPairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	return _ILynexPair.Contract.GetReserves(&_ILynexPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1, uint256 _blockTimestampLast)
func (_ILynexPair *ILynexPairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast *big.Int
}, error) {
	return _ILynexPair.Contract.GetReserves(&_ILynexPair.CallOpts)
}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Index0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "index0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_ILynexPair *ILynexPairSession) Index0() (*big.Int, error) {
	return _ILynexPair.Contract.Index0(&_ILynexPair.CallOpts)
}

// Index0 is a free data retrieval call binding the contract method 0x32c0defd.
//
// Solidity: function index0() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Index0() (*big.Int, error) {
	return _ILynexPair.Contract.Index0(&_ILynexPair.CallOpts)
}

// Index1 is a free data retrieval call binding the contract method 0xbda39cad.
//
// Solidity: function index1() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Index1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "index1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index1 is a free data retrieval call binding the contract method 0xbda39cad.
//
// Solidity: function index1() view returns(uint256)
func (_ILynexPair *ILynexPairSession) Index1() (*big.Int, error) {
	return _ILynexPair.Contract.Index1(&_ILynexPair.CallOpts)
}

// Index1 is a free data retrieval call binding the contract method 0xbda39cad.
//
// Solidity: function index1() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Index1() (*big.Int, error) {
	return _ILynexPair.Contract.Index1(&_ILynexPair.CallOpts)
}

// IsStable is a free data retrieval call binding the contract method 0x09047bdd.
//
// Solidity: function isStable() view returns(bool)
func (_ILynexPair *ILynexPairCaller) IsStable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "isStable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStable is a free data retrieval call binding the contract method 0x09047bdd.
//
// Solidity: function isStable() view returns(bool)
func (_ILynexPair *ILynexPairSession) IsStable() (bool, error) {
	return _ILynexPair.Contract.IsStable(&_ILynexPair.CallOpts)
}

// IsStable is a free data retrieval call binding the contract method 0x09047bdd.
//
// Solidity: function isStable() view returns(bool)
func (_ILynexPair *ILynexPairCallerSession) IsStable() (bool, error) {
	return _ILynexPair.Contract.IsStable(&_ILynexPair.CallOpts)
}

// LastObservation is a free data retrieval call binding the contract method 0x8a7b8cf2.
//
// Solidity: function lastObservation() view returns((uint256,uint256,uint256))
func (_ILynexPair *ILynexPairCaller) LastObservation(opts *bind.CallOpts) (PairObservation, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "lastObservation")

	if err != nil {
		return *new(PairObservation), err
	}

	out0 := *abi.ConvertType(out[0], new(PairObservation)).(*PairObservation)

	return out0, err

}

// LastObservation is a free data retrieval call binding the contract method 0x8a7b8cf2.
//
// Solidity: function lastObservation() view returns((uint256,uint256,uint256))
func (_ILynexPair *ILynexPairSession) LastObservation() (PairObservation, error) {
	return _ILynexPair.Contract.LastObservation(&_ILynexPair.CallOpts)
}

// LastObservation is a free data retrieval call binding the contract method 0x8a7b8cf2.
//
// Solidity: function lastObservation() view returns((uint256,uint256,uint256))
func (_ILynexPair *ILynexPairCallerSession) LastObservation() (PairObservation, error) {
	return _ILynexPair.Contract.LastObservation(&_ILynexPair.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_ILynexPair *ILynexPairCaller) Metadata(opts *bind.CallOpts) (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "metadata")

	outstruct := new(struct {
		Dec0 *big.Int
		Dec1 *big.Int
		R0   *big.Int
		R1   *big.Int
		St   bool
		T0   common.Address
		T1   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dec0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Dec1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.R0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.R1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.St = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.T0 = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.T1 = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_ILynexPair *ILynexPairSession) Metadata() (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	return _ILynexPair.Contract.Metadata(&_ILynexPair.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(uint256 dec0, uint256 dec1, uint256 r0, uint256 r1, bool st, address t0, address t1)
func (_ILynexPair *ILynexPairCallerSession) Metadata() (struct {
	Dec0 *big.Int
	Dec1 *big.Int
	R0   *big.Int
	R1   *big.Int
	St   bool
	T0   common.Address
	T1   common.Address
}, error) {
	return _ILynexPair.Contract.Metadata(&_ILynexPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ILynexPair *ILynexPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ILynexPair *ILynexPairSession) Name() (string, error) {
	return _ILynexPair.Contract.Name(&_ILynexPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ILynexPair *ILynexPairCallerSession) Name() (string, error) {
	return _ILynexPair.Contract.Name(&_ILynexPair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Nonces(&_ILynexPair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.Nonces(&_ILynexPair.CallOpts, arg0)
}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) ObservationLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "observationLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_ILynexPair *ILynexPairSession) ObservationLength() (*big.Int, error) {
	return _ILynexPair.Contract.ObservationLength(&_ILynexPair.CallOpts)
}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) ObservationLength() (*big.Int, error) {
	return _ILynexPair.Contract.ObservationLength(&_ILynexPair.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint256 timestamp, uint256 reserve0Cumulative, uint256 reserve1Cumulative)
func (_ILynexPair *ILynexPairCaller) Observations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "observations", arg0)

	outstruct := new(struct {
		Timestamp          *big.Int
		Reserve0Cumulative *big.Int
		Reserve1Cumulative *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve0Cumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Reserve1Cumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint256 timestamp, uint256 reserve0Cumulative, uint256 reserve1Cumulative)
func (_ILynexPair *ILynexPairSession) Observations(arg0 *big.Int) (struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}, error) {
	return _ILynexPair.Contract.Observations(&_ILynexPair.CallOpts, arg0)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint256 timestamp, uint256 reserve0Cumulative, uint256 reserve1Cumulative)
func (_ILynexPair *ILynexPairCallerSession) Observations(arg0 *big.Int) (struct {
	Timestamp          *big.Int
	Reserve0Cumulative *big.Int
	Reserve1Cumulative *big.Int
}, error) {
	return _ILynexPair.Contract.Observations(&_ILynexPair.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x5881c475.
//
// Solidity: function prices(address tokenIn, uint256 amountIn, uint256 points) view returns(uint256[])
func (_ILynexPair *ILynexPairCaller) Prices(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, points *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "prices", tokenIn, amountIn, points)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0x5881c475.
//
// Solidity: function prices(address tokenIn, uint256 amountIn, uint256 points) view returns(uint256[])
func (_ILynexPair *ILynexPairSession) Prices(tokenIn common.Address, amountIn *big.Int, points *big.Int) ([]*big.Int, error) {
	return _ILynexPair.Contract.Prices(&_ILynexPair.CallOpts, tokenIn, amountIn, points)
}

// Prices is a free data retrieval call binding the contract method 0x5881c475.
//
// Solidity: function prices(address tokenIn, uint256 amountIn, uint256 points) view returns(uint256[])
func (_ILynexPair *ILynexPairCallerSession) Prices(tokenIn common.Address, amountIn *big.Int, points *big.Int) ([]*big.Int, error) {
	return _ILynexPair.Contract.Prices(&_ILynexPair.CallOpts, tokenIn, amountIn, points)
}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256 amountOut)
func (_ILynexPair *ILynexPairCaller) Quote(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "quote", tokenIn, amountIn, granularity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256 amountOut)
func (_ILynexPair *ILynexPairSession) Quote(tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	return _ILynexPair.Contract.Quote(&_ILynexPair.CallOpts, tokenIn, amountIn, granularity)
}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256 amountOut)
func (_ILynexPair *ILynexPairCallerSession) Quote(tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	return _ILynexPair.Contract.Quote(&_ILynexPair.CallOpts, tokenIn, amountIn, granularity)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ILynexPair *ILynexPairSession) Reserve0() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve0(&_ILynexPair.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Reserve0() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve0(&_ILynexPair.CallOpts)
}

// Reserve0CumulativeLast is a free data retrieval call binding the contract method 0xbf944dbc.
//
// Solidity: function reserve0CumulativeLast() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Reserve0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "reserve0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0CumulativeLast is a free data retrieval call binding the contract method 0xbf944dbc.
//
// Solidity: function reserve0CumulativeLast() view returns(uint256)
func (_ILynexPair *ILynexPairSession) Reserve0CumulativeLast() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve0CumulativeLast(&_ILynexPair.CallOpts)
}

// Reserve0CumulativeLast is a free data retrieval call binding the contract method 0xbf944dbc.
//
// Solidity: function reserve0CumulativeLast() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Reserve0CumulativeLast() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve0CumulativeLast(&_ILynexPair.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ILynexPair *ILynexPairSession) Reserve1() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve1(&_ILynexPair.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Reserve1() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve1(&_ILynexPair.CallOpts)
}

// Reserve1CumulativeLast is a free data retrieval call binding the contract method 0xc245febc.
//
// Solidity: function reserve1CumulativeLast() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) Reserve1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "reserve1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1CumulativeLast is a free data retrieval call binding the contract method 0xc245febc.
//
// Solidity: function reserve1CumulativeLast() view returns(uint256)
func (_ILynexPair *ILynexPairSession) Reserve1CumulativeLast() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve1CumulativeLast(&_ILynexPair.CallOpts)
}

// Reserve1CumulativeLast is a free data retrieval call binding the contract method 0xc245febc.
//
// Solidity: function reserve1CumulativeLast() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) Reserve1CumulativeLast() (*big.Int, error) {
	return _ILynexPair.Contract.Reserve1CumulativeLast(&_ILynexPair.CallOpts)
}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_ILynexPair *ILynexPairCaller) Sample(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "sample", tokenIn, amountIn, points, window)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_ILynexPair *ILynexPairSession) Sample(tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _ILynexPair.Contract.Sample(&_ILynexPair.CallOpts, tokenIn, amountIn, points, window)
}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_ILynexPair *ILynexPairCallerSession) Sample(tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _ILynexPair.Contract.Sample(&_ILynexPair.CallOpts, tokenIn, amountIn, points, window)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_ILynexPair *ILynexPairCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_ILynexPair *ILynexPairSession) Stable() (bool, error) {
	return _ILynexPair.Contract.Stable(&_ILynexPair.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_ILynexPair *ILynexPairCallerSession) Stable() (bool, error) {
	return _ILynexPair.Contract.Stable(&_ILynexPair.CallOpts)
}

// SupplyIndex0 is a free data retrieval call binding the contract method 0x9f767c88.
//
// Solidity: function supplyIndex0(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) SupplyIndex0(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "supplyIndex0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyIndex0 is a free data retrieval call binding the contract method 0x9f767c88.
//
// Solidity: function supplyIndex0(address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) SupplyIndex0(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.SupplyIndex0(&_ILynexPair.CallOpts, arg0)
}

// SupplyIndex0 is a free data retrieval call binding the contract method 0x9f767c88.
//
// Solidity: function supplyIndex0(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) SupplyIndex0(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.SupplyIndex0(&_ILynexPair.CallOpts, arg0)
}

// SupplyIndex1 is a free data retrieval call binding the contract method 0x205aabf1.
//
// Solidity: function supplyIndex1(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCaller) SupplyIndex1(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "supplyIndex1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyIndex1 is a free data retrieval call binding the contract method 0x205aabf1.
//
// Solidity: function supplyIndex1(address ) view returns(uint256)
func (_ILynexPair *ILynexPairSession) SupplyIndex1(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.SupplyIndex1(&_ILynexPair.CallOpts, arg0)
}

// SupplyIndex1 is a free data retrieval call binding the contract method 0x205aabf1.
//
// Solidity: function supplyIndex1(address ) view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) SupplyIndex1(arg0 common.Address) (*big.Int, error) {
	return _ILynexPair.Contract.SupplyIndex1(&_ILynexPair.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ILynexPair *ILynexPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ILynexPair *ILynexPairSession) Symbol() (string, error) {
	return _ILynexPair.Contract.Symbol(&_ILynexPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ILynexPair *ILynexPairCallerSession) Symbol() (string, error) {
	return _ILynexPair.Contract.Symbol(&_ILynexPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ILynexPair *ILynexPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ILynexPair *ILynexPairSession) Token0() (common.Address, error) {
	return _ILynexPair.Contract.Token0(&_ILynexPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ILynexPair *ILynexPairCallerSession) Token0() (common.Address, error) {
	return _ILynexPair.Contract.Token0(&_ILynexPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ILynexPair *ILynexPairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ILynexPair *ILynexPairSession) Token1() (common.Address, error) {
	return _ILynexPair.Contract.Token1(&_ILynexPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ILynexPair *ILynexPairCallerSession) Token1() (common.Address, error) {
	return _ILynexPair.Contract.Token1(&_ILynexPair.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_ILynexPair *ILynexPairCaller) Tokens(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "tokens")

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_ILynexPair *ILynexPairSession) Tokens() (common.Address, common.Address, error) {
	return _ILynexPair.Contract.Tokens(&_ILynexPair.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address, address)
func (_ILynexPair *ILynexPairCallerSession) Tokens() (common.Address, common.Address, error) {
	return _ILynexPair.Contract.Tokens(&_ILynexPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ILynexPair *ILynexPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexPair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ILynexPair *ILynexPairSession) TotalSupply() (*big.Int, error) {
	return _ILynexPair.Contract.TotalSupply(&_ILynexPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ILynexPair *ILynexPairCallerSession) TotalSupply() (*big.Int, error) {
	return _ILynexPair.Contract.TotalSupply(&_ILynexPair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.Contract.Approve(&_ILynexPair.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.Contract.Approve(&_ILynexPair.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _ILynexPair.Contract.Burn(&_ILynexPair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _ILynexPair.Contract.Burn(&_ILynexPair.TransactOpts, to)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_ILynexPair *ILynexPairTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_ILynexPair *ILynexPairSession) ClaimFees() (*types.Transaction, error) {
	return _ILynexPair.Contract.ClaimFees(&_ILynexPair.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(uint256 claimed0, uint256 claimed1)
func (_ILynexPair *ILynexPairTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _ILynexPair.Contract.ClaimFees(&_ILynexPair.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ILynexPair *ILynexPairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ILynexPair *ILynexPairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _ILynexPair.Contract.Mint(&_ILynexPair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ILynexPair *ILynexPairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _ILynexPair.Contract.Mint(&_ILynexPair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ILynexPair *ILynexPairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ILynexPair *ILynexPairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ILynexPair.Contract.Permit(&_ILynexPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ILynexPair *ILynexPairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ILynexPair.Contract.Permit(&_ILynexPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_ILynexPair *ILynexPairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_ILynexPair *ILynexPairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _ILynexPair.Contract.Skim(&_ILynexPair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_ILynexPair *ILynexPairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _ILynexPair.Contract.Skim(&_ILynexPair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ILynexPair *ILynexPairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ILynexPair *ILynexPairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ILynexPair.Contract.Swap(&_ILynexPair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ILynexPair *ILynexPairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ILynexPair.Contract.Swap(&_ILynexPair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ILynexPair *ILynexPairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ILynexPair *ILynexPairSession) Sync() (*types.Transaction, error) {
	return _ILynexPair.Contract.Sync(&_ILynexPair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ILynexPair *ILynexPairTransactorSession) Sync() (*types.Transaction, error) {
	return _ILynexPair.Contract.Sync(&_ILynexPair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.Contract.Transfer(&_ILynexPair.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.Contract.Transfer(&_ILynexPair.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.Contract.TransferFrom(&_ILynexPair.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_ILynexPair *ILynexPairTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ILynexPair.Contract.TransferFrom(&_ILynexPair.TransactOpts, src, dst, amount)
}

// ILynexPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ILynexPair contract.
type ILynexPairApprovalIterator struct {
	Event *ILynexPairApproval // Event containing the contract specifics and raw log

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
func (it *ILynexPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairApproval)
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
		it.Event = new(ILynexPairApproval)
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
func (it *ILynexPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairApproval represents a Approval event raised by the ILynexPair contract.
type ILynexPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ILynexPair *ILynexPairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ILynexPairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairApprovalIterator{contract: _ILynexPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ILynexPair *ILynexPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ILynexPairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairApproval)
				if err := _ILynexPair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ILynexPair *ILynexPairFilterer) ParseApproval(log types.Log) (*ILynexPairApproval, error) {
	event := new(ILynexPairApproval)
	if err := _ILynexPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ILynexPair contract.
type ILynexPairBurnIterator struct {
	Event *ILynexPairBurn // Event containing the contract specifics and raw log

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
func (it *ILynexPairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairBurn)
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
		it.Event = new(ILynexPairBurn)
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
func (it *ILynexPairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairBurn represents a Burn event raised by the ILynexPair contract.
type ILynexPairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_ILynexPair *ILynexPairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ILynexPairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairBurnIterator{contract: _ILynexPair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_ILynexPair *ILynexPairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ILynexPairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairBurn)
				if err := _ILynexPair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_ILynexPair *ILynexPairFilterer) ParseBurn(log types.Log) (*ILynexPairBurn, error) {
	event := new(ILynexPairBurn)
	if err := _ILynexPair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the ILynexPair contract.
type ILynexPairClaimIterator struct {
	Event *ILynexPairClaim // Event containing the contract specifics and raw log

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
func (it *ILynexPairClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairClaim)
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
		it.Event = new(ILynexPairClaim)
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
func (it *ILynexPairClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairClaim represents a Claim event raised by the ILynexPair contract.
type ILynexPairClaim struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) FilterClaim(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*ILynexPairClaimIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Claim", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairClaimIterator{contract: _ILynexPair.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *ILynexPairClaim, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Claim", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairClaim)
				if err := _ILynexPair.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) ParseClaim(log types.Log) (*ILynexPairClaim, error) {
	event := new(ILynexPairClaim)
	if err := _ILynexPair.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairFeesIterator is returned from FilterFees and is used to iterate over the raw logs and unpacked data for Fees events raised by the ILynexPair contract.
type ILynexPairFeesIterator struct {
	Event *ILynexPairFees // Event containing the contract specifics and raw log

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
func (it *ILynexPairFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairFees)
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
		it.Event = new(ILynexPairFees)
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
func (it *ILynexPairFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairFees represents a Fees event raised by the ILynexPair contract.
type ILynexPairFees struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFees is a free log retrieval operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) FilterFees(opts *bind.FilterOpts, sender []common.Address) (*ILynexPairFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Fees", senderRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairFeesIterator{contract: _ILynexPair.contract, event: "Fees", logs: logs, sub: sub}, nil
}

// WatchFees is a free log subscription operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) WatchFees(opts *bind.WatchOpts, sink chan<- *ILynexPairFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Fees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairFees)
				if err := _ILynexPair.contract.UnpackLog(event, "Fees", log); err != nil {
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

// ParseFees is a log parse operation binding the contract event 0x112c256902bf554b6ed882d2936687aaeb4225e8cd5b51303c90ca6cf43a8602.
//
// Solidity: event Fees(address indexed sender, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) ParseFees(log types.Log) (*ILynexPairFees, error) {
	event := new(ILynexPairFees)
	if err := _ILynexPair.contract.UnpackLog(event, "Fees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ILynexPair contract.
type ILynexPairMintIterator struct {
	Event *ILynexPairMint // Event containing the contract specifics and raw log

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
func (it *ILynexPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairMint)
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
		it.Event = new(ILynexPairMint)
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
func (it *ILynexPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairMint represents a Mint event raised by the ILynexPair contract.
type ILynexPairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*ILynexPairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairMintIterator{contract: _ILynexPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ILynexPairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairMint)
				if err := _ILynexPair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_ILynexPair *ILynexPairFilterer) ParseMint(log types.Log) (*ILynexPairMint, error) {
	event := new(ILynexPairMint)
	if err := _ILynexPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ILynexPair contract.
type ILynexPairSwapIterator struct {
	Event *ILynexPairSwap // Event containing the contract specifics and raw log

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
func (it *ILynexPairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairSwap)
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
		it.Event = new(ILynexPairSwap)
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
func (it *ILynexPairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairSwap represents a Swap event raised by the ILynexPair contract.
type ILynexPairSwap struct {
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
func (_ILynexPair *ILynexPairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ILynexPairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairSwapIterator{contract: _ILynexPair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ILynexPair *ILynexPairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ILynexPairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairSwap)
				if err := _ILynexPair.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_ILynexPair *ILynexPairFilterer) ParseSwap(log types.Log) (*ILynexPairSwap, error) {
	event := new(ILynexPairSwap)
	if err := _ILynexPair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the ILynexPair contract.
type ILynexPairSyncIterator struct {
	Event *ILynexPairSync // Event containing the contract specifics and raw log

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
func (it *ILynexPairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairSync)
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
		it.Event = new(ILynexPairSync)
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
func (it *ILynexPairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairSync represents a Sync event raised by the ILynexPair contract.
type ILynexPairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ILynexPair *ILynexPairFilterer) FilterSync(opts *bind.FilterOpts) (*ILynexPairSyncIterator, error) {

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &ILynexPairSyncIterator{contract: _ILynexPair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ILynexPair *ILynexPairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *ILynexPairSync) (event.Subscription, error) {

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairSync)
				if err := _ILynexPair.contract.UnpackLog(event, "Sync", log); err != nil {
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
func (_ILynexPair *ILynexPairFilterer) ParseSync(log types.Log) (*ILynexPairSync, error) {
	event := new(ILynexPairSync)
	if err := _ILynexPair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ILynexPair contract.
type ILynexPairTransferIterator struct {
	Event *ILynexPairTransfer // Event containing the contract specifics and raw log

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
func (it *ILynexPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexPairTransfer)
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
		it.Event = new(ILynexPairTransfer)
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
func (it *ILynexPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexPairTransfer represents a Transfer event raised by the ILynexPair contract.
type ILynexPairTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ILynexPair *ILynexPairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ILynexPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILynexPair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILynexPairTransferIterator{contract: _ILynexPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ILynexPair *ILynexPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ILynexPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILynexPair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexPairTransfer)
				if err := _ILynexPair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ILynexPair *ILynexPairFilterer) ParseTransfer(log types.Log) (*ILynexPairTransfer, error) {
	event := new(ILynexPairTransfer)
	if err := _ILynexPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
