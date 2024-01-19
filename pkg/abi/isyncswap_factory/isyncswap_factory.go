// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isyncswap_factory

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

// ISyncSwapFactoryMetaData contains all meta data concerning the ISyncSwapFactory contract.
var ISyncSwapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISyncSwapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ISyncSwapFactoryMetaData.ABI instead.
var ISyncSwapFactoryABI = ISyncSwapFactoryMetaData.ABI

// ISyncSwapFactory is an auto generated Go binding around an Ethereum contract.
type ISyncSwapFactory struct {
	ISyncSwapFactoryCaller     // Read-only binding to the contract
	ISyncSwapFactoryTransactor // Write-only binding to the contract
	ISyncSwapFactoryFilterer   // Log filterer for contract events
}

// ISyncSwapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISyncSwapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISyncSwapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISyncSwapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncSwapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISyncSwapFactorySession struct {
	Contract     *ISyncSwapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISyncSwapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISyncSwapFactoryCallerSession struct {
	Contract *ISyncSwapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISyncSwapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISyncSwapFactoryTransactorSession struct {
	Contract     *ISyncSwapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISyncSwapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISyncSwapFactoryRaw struct {
	Contract *ISyncSwapFactory // Generic contract binding to access the raw methods on
}

// ISyncSwapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISyncSwapFactoryCallerRaw struct {
	Contract *ISyncSwapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ISyncSwapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISyncSwapFactoryTransactorRaw struct {
	Contract *ISyncSwapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISyncSwapFactory creates a new instance of ISyncSwapFactory, bound to a specific deployed contract.
func NewISyncSwapFactory(address common.Address, backend bind.ContractBackend) (*ISyncSwapFactory, error) {
	contract, err := bindISyncSwapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapFactory{ISyncSwapFactoryCaller: ISyncSwapFactoryCaller{contract: contract}, ISyncSwapFactoryTransactor: ISyncSwapFactoryTransactor{contract: contract}, ISyncSwapFactoryFilterer: ISyncSwapFactoryFilterer{contract: contract}}, nil
}

// NewISyncSwapFactoryCaller creates a new read-only instance of ISyncSwapFactory, bound to a specific deployed contract.
func NewISyncSwapFactoryCaller(address common.Address, caller bind.ContractCaller) (*ISyncSwapFactoryCaller, error) {
	contract, err := bindISyncSwapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapFactoryCaller{contract: contract}, nil
}

// NewISyncSwapFactoryTransactor creates a new write-only instance of ISyncSwapFactory, bound to a specific deployed contract.
func NewISyncSwapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ISyncSwapFactoryTransactor, error) {
	contract, err := bindISyncSwapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapFactoryTransactor{contract: contract}, nil
}

// NewISyncSwapFactoryFilterer creates a new log filterer instance of ISyncSwapFactory, bound to a specific deployed contract.
func NewISyncSwapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ISyncSwapFactoryFilterer, error) {
	contract, err := bindISyncSwapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapFactoryFilterer{contract: contract}, nil
}

// bindISyncSwapFactory binds a generic wrapper to an already deployed contract.
func bindISyncSwapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISyncSwapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapFactory *ISyncSwapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapFactory.Contract.ISyncSwapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapFactory *ISyncSwapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapFactory.Contract.ISyncSwapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapFactory *ISyncSwapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapFactory.Contract.ISyncSwapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncSwapFactory *ISyncSwapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncSwapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncSwapFactory *ISyncSwapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncSwapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncSwapFactory *ISyncSwapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncSwapFactory.Contract.contract.Transact(opts, method, params...)
}

// GetDeployData is a free data retrieval call binding the contract method 0xd039f622.
//
// Solidity: function getDeployData() view returns(bytes)
func (_ISyncSwapFactory *ISyncSwapFactoryCaller) GetDeployData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ISyncSwapFactory.contract.Call(opts, &out, "getDeployData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDeployData is a free data retrieval call binding the contract method 0xd039f622.
//
// Solidity: function getDeployData() view returns(bytes)
func (_ISyncSwapFactory *ISyncSwapFactorySession) GetDeployData() ([]byte, error) {
	return _ISyncSwapFactory.Contract.GetDeployData(&_ISyncSwapFactory.CallOpts)
}

// GetDeployData is a free data retrieval call binding the contract method 0xd039f622.
//
// Solidity: function getDeployData() view returns(bytes)
func (_ISyncSwapFactory *ISyncSwapFactoryCallerSession) GetDeployData() ([]byte, error) {
	return _ISyncSwapFactory.Contract.GetDeployData(&_ISyncSwapFactory.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x531aa03e.
//
// Solidity: function getPool(address tokenA, address tokenB) view returns(address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryCaller) GetPool(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapFactory.contract.Call(opts, &out, "getPool", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x531aa03e.
//
// Solidity: function getPool(address tokenA, address tokenB) view returns(address pool)
func (_ISyncSwapFactory *ISyncSwapFactorySession) GetPool(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _ISyncSwapFactory.Contract.GetPool(&_ISyncSwapFactory.CallOpts, tokenA, tokenB)
}

// GetPool is a free data retrieval call binding the contract method 0x531aa03e.
//
// Solidity: function getPool(address tokenA, address tokenB) view returns(address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryCallerSession) GetPool(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _ISyncSwapFactory.Contract.GetPool(&_ISyncSwapFactory.CallOpts, tokenA, tokenB)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ISyncSwapFactory *ISyncSwapFactoryCaller) GetSwapFee(opts *bind.CallOpts, pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _ISyncSwapFactory.contract.Call(opts, &out, "getSwapFee", pool, sender, tokenIn, tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ISyncSwapFactory *ISyncSwapFactorySession) GetSwapFee(pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ISyncSwapFactory.Contract.GetSwapFee(&_ISyncSwapFactory.CallOpts, pool, sender, tokenIn, tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ISyncSwapFactory *ISyncSwapFactoryCallerSession) GetSwapFee(pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ISyncSwapFactory.Contract.GetSwapFee(&_ISyncSwapFactory.CallOpts, pool, sender, tokenIn, tokenOut, data)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ISyncSwapFactory *ISyncSwapFactoryCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISyncSwapFactory.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ISyncSwapFactory *ISyncSwapFactorySession) Master() (common.Address, error) {
	return _ISyncSwapFactory.Contract.Master(&_ISyncSwapFactory.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ISyncSwapFactory *ISyncSwapFactoryCallerSession) Master() (common.Address, error) {
	return _ISyncSwapFactory.Contract.Master(&_ISyncSwapFactory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x13b8683f.
//
// Solidity: function createPool(bytes data) returns(address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryTransactor) CreatePool(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ISyncSwapFactory.contract.Transact(opts, "createPool", data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x13b8683f.
//
// Solidity: function createPool(bytes data) returns(address pool)
func (_ISyncSwapFactory *ISyncSwapFactorySession) CreatePool(data []byte) (*types.Transaction, error) {
	return _ISyncSwapFactory.Contract.CreatePool(&_ISyncSwapFactory.TransactOpts, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x13b8683f.
//
// Solidity: function createPool(bytes data) returns(address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryTransactorSession) CreatePool(data []byte) (*types.Transaction, error) {
	return _ISyncSwapFactory.Contract.CreatePool(&_ISyncSwapFactory.TransactOpts, data)
}

// ISyncSwapFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the ISyncSwapFactory contract.
type ISyncSwapFactoryPoolCreatedIterator struct {
	Event *ISyncSwapFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *ISyncSwapFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISyncSwapFactoryPoolCreated)
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
		it.Event = new(ISyncSwapFactoryPoolCreated)
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
func (it *ISyncSwapFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISyncSwapFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISyncSwapFactoryPoolCreated represents a PoolCreated event raised by the ISyncSwapFactory contract.
type ISyncSwapFactoryPoolCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x9c5d829b9b23efc461f9aeef91979ec04bb903feb3bee4f26d22114abfc7335b.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*ISyncSwapFactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ISyncSwapFactory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &ISyncSwapFactoryPoolCreatedIterator{contract: _ISyncSwapFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x9c5d829b9b23efc461f9aeef91979ec04bb903feb3bee4f26d22114abfc7335b.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *ISyncSwapFactoryPoolCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ISyncSwapFactory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISyncSwapFactoryPoolCreated)
				if err := _ISyncSwapFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x9c5d829b9b23efc461f9aeef91979ec04bb903feb3bee4f26d22114abfc7335b.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, address pool)
func (_ISyncSwapFactory *ISyncSwapFactoryFilterer) ParsePoolCreated(log types.Log) (*ISyncSwapFactoryPoolCreated, error) {
	event := new(ISyncSwapFactoryPoolCreated)
	if err := _ISyncSwapFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
