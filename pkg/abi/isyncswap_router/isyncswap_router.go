// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isyncswap_router

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

// ISyncSwapRouterMetaData contains all meta data concerning the ISyncSwapRouter contract.
var ISyncSwapRouterMetaData = &bind.MetaData{
	ABI: "[]",
}

// ISyncSwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISyncSwapRouterMetaData.ABI instead.
var ISyncSwapRouterABI = ISyncSwapRouterMetaData.ABI

// ISyncSwapRouter is an auto generated Go binding around an Ethereum contract.
type ISyncSwapRouter struct {
	ISyncSwapRouterCaller     // Read-only binding to the contract
	ISyncSwapRouterTransactor // Write-only binding to the contract
	ISyncSwapRouterFilterer   // Log filterer for contract events
}

// ISyncSwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISyncSwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISyncSwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISyncSwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISyncSwapRouterSession struct {
	Contract     *ISyncSwapRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISyncSwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISyncSwapRouterCallerSession struct {
	Contract *ISyncSwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISyncSwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISyncSwapRouterTransactorSession struct {
	Contract     *ISyncSwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISyncSwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISyncSwapRouterRaw struct {
	Contract *ISyncSwapRouter // Generic contract binding to access the raw methods on
}

// ISyncSwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISyncSwapRouterCallerRaw struct {
	Contract *ISyncSwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ISyncSwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISyncSwapRouterTransactorRaw struct {
	Contract *ISyncSwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISyncSwapRouter creates a new instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouter(address common.Address, backend bind.ContractBackend) (*ISyncSwapRouter, error) {
	contract, err := bindISyncSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouter{ISyncSwapRouterCaller: ISyncSwapRouterCaller{contract: contract}, ISyncSwapRouterTransactor: ISyncSwapRouterTransactor{contract: contract}, ISyncSwapRouterFilterer: ISyncSwapRouterFilterer{contract: contract}}, nil
}

// NewISyncSwapRouterCaller creates a new read-only instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*ISyncSwapRouterCaller, error) {
	contract, err := bindISyncSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouterCaller{contract: contract}, nil
}

// NewISyncSwapRouterTransactor creates a new write-only instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISyncSwapRouterTransactor, error) {
	contract, err := bindISyncSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouterTransactor{contract: contract}, nil
}

// NewISyncSwapRouterFilterer creates a new log filterer instance of ISyncSwapRouter, bound to a specific deployed contract.
func NewISyncSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISyncSwapRouterFilterer, error) {
	contract, err := bindISyncSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapRouterFilterer{contract: contract}, nil
}

// bindISyncSwapRouter binds a generic wrapper to an already deployed contract.
func bindISyncSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISyncSwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapRouter *ISyncSwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapRouter.Contract.ISyncSwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapRouter *ISyncSwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.ISyncSwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapRouter *ISyncSwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.ISyncSwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapRouter *ISyncSwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapRouter *ISyncSwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapRouter *ISyncSwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapRouter.Contract.contract.Transact(opts, method, params...)
}
