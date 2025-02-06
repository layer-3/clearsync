// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imulti_asset_holder

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

// IMultiAssetHolderReclaimArgs is an auto generated low-level Go binding around an user-defined struct.
type IMultiAssetHolderReclaimArgs struct {
	SourceChannelId       [32]byte
	SourceStateHash       [32]byte
	SourceOutcomeBytes    []byte
	SourceAssetIndex      *big.Int
	IndexOfTargetInSource *big.Int
	TargetStateHash       [32]byte
	TargetOutcomeBytes    []byte
	TargetAssetIndex      *big.Int
}

// ImultiAssetHolderMetaData contains all meta data concerning the ImultiAssetHolder contract.
var ImultiAssetHolderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalHoldings\",\"type\":\"uint256\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationHoldings\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeld\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceAssetIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"targetStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"targetOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"targetAssetIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiAssetHolder.ReclaimArgs\",\"name\":\"reclaimArgs\",\"type\":\"tuple\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"fromChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"outcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ImultiAssetHolderABI is the input ABI used to generate the binding from.
// Deprecated: Use ImultiAssetHolderMetaData.ABI instead.
var ImultiAssetHolderABI = ImultiAssetHolderMetaData.ABI

// ImultiAssetHolder is an auto generated Go binding around an Ethereum contract.
type ImultiAssetHolder struct {
	ImultiAssetHolderCaller     // Read-only binding to the contract
	ImultiAssetHolderTransactor // Write-only binding to the contract
	ImultiAssetHolderFilterer   // Log filterer for contract events
}

// ImultiAssetHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImultiAssetHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImultiAssetHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImultiAssetHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImultiAssetHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImultiAssetHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImultiAssetHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImultiAssetHolderSession struct {
	Contract     *ImultiAssetHolder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ImultiAssetHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImultiAssetHolderCallerSession struct {
	Contract *ImultiAssetHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ImultiAssetHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImultiAssetHolderTransactorSession struct {
	Contract     *ImultiAssetHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ImultiAssetHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImultiAssetHolderRaw struct {
	Contract *ImultiAssetHolder // Generic contract binding to access the raw methods on
}

// ImultiAssetHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImultiAssetHolderCallerRaw struct {
	Contract *ImultiAssetHolderCaller // Generic read-only contract binding to access the raw methods on
}

// ImultiAssetHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImultiAssetHolderTransactorRaw struct {
	Contract *ImultiAssetHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImultiAssetHolder creates a new instance of ImultiAssetHolder, bound to a specific deployed contract.
func NewImultiAssetHolder(address common.Address, backend bind.ContractBackend) (*ImultiAssetHolder, error) {
	contract, err := bindImultiAssetHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolder{ImultiAssetHolderCaller: ImultiAssetHolderCaller{contract: contract}, ImultiAssetHolderTransactor: ImultiAssetHolderTransactor{contract: contract}, ImultiAssetHolderFilterer: ImultiAssetHolderFilterer{contract: contract}}, nil
}

// NewImultiAssetHolderCaller creates a new read-only instance of ImultiAssetHolder, bound to a specific deployed contract.
func NewImultiAssetHolderCaller(address common.Address, caller bind.ContractCaller) (*ImultiAssetHolderCaller, error) {
	contract, err := bindImultiAssetHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolderCaller{contract: contract}, nil
}

// NewImultiAssetHolderTransactor creates a new write-only instance of ImultiAssetHolder, bound to a specific deployed contract.
func NewImultiAssetHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ImultiAssetHolderTransactor, error) {
	contract, err := bindImultiAssetHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolderTransactor{contract: contract}, nil
}

// NewImultiAssetHolderFilterer creates a new log filterer instance of ImultiAssetHolder, bound to a specific deployed contract.
func NewImultiAssetHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ImultiAssetHolderFilterer, error) {
	contract, err := bindImultiAssetHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolderFilterer{contract: contract}, nil
}

// bindImultiAssetHolder binds a generic wrapper to an already deployed contract.
func bindImultiAssetHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImultiAssetHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImultiAssetHolder *ImultiAssetHolderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImultiAssetHolder.Contract.ImultiAssetHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImultiAssetHolder *ImultiAssetHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.ImultiAssetHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImultiAssetHolder *ImultiAssetHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.ImultiAssetHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImultiAssetHolder *ImultiAssetHolderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImultiAssetHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImultiAssetHolder *ImultiAssetHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImultiAssetHolder *ImultiAssetHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_ImultiAssetHolder *ImultiAssetHolderTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ImultiAssetHolder.contract.Transact(opts, "deposit", asset, destination, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_ImultiAssetHolder *ImultiAssetHolderSession) Deposit(asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.Deposit(&_ImultiAssetHolder.TransactOpts, asset, destination, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_ImultiAssetHolder *ImultiAssetHolderTransactorSession) Deposit(asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.Deposit(&_ImultiAssetHolder.TransactOpts, asset, destination, expectedHeld, amount)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_ImultiAssetHolder *ImultiAssetHolderTransactor) Reclaim(opts *bind.TransactOpts, reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _ImultiAssetHolder.contract.Transact(opts, "reclaim", reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_ImultiAssetHolder *ImultiAssetHolderSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.Reclaim(&_ImultiAssetHolder.TransactOpts, reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_ImultiAssetHolder *ImultiAssetHolderTransactorSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.Reclaim(&_ImultiAssetHolder.TransactOpts, reclaimArgs)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_ImultiAssetHolder *ImultiAssetHolderTransactor) Transfer(opts *bind.TransactOpts, assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _ImultiAssetHolder.contract.Transact(opts, "transfer", assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_ImultiAssetHolder *ImultiAssetHolderSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.Transfer(&_ImultiAssetHolder.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_ImultiAssetHolder *ImultiAssetHolderTransactorSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _ImultiAssetHolder.Contract.Transfer(&_ImultiAssetHolder.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// ImultiAssetHolderAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the ImultiAssetHolder contract.
type ImultiAssetHolderAllocationUpdatedIterator struct {
	Event *ImultiAssetHolderAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *ImultiAssetHolderAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImultiAssetHolderAllocationUpdated)
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
		it.Event = new(ImultiAssetHolderAllocationUpdated)
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
func (it *ImultiAssetHolderAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImultiAssetHolderAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImultiAssetHolderAllocationUpdated represents a AllocationUpdated event raised by the ImultiAssetHolder contract.
type ImultiAssetHolderAllocationUpdated struct {
	ChannelId       [32]byte
	AssetIndex      *big.Int
	InitialHoldings *big.Int
	FinalHoldings   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) FilterAllocationUpdated(opts *bind.FilterOpts, channelId [][32]byte) (*ImultiAssetHolderAllocationUpdatedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _ImultiAssetHolder.contract.FilterLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolderAllocationUpdatedIterator{contract: _ImultiAssetHolder.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *ImultiAssetHolderAllocationUpdated, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _ImultiAssetHolder.contract.WatchLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImultiAssetHolderAllocationUpdated)
				if err := _ImultiAssetHolder.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
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

// ParseAllocationUpdated is a log parse operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) ParseAllocationUpdated(log types.Log) (*ImultiAssetHolderAllocationUpdated, error) {
	event := new(ImultiAssetHolderAllocationUpdated)
	if err := _ImultiAssetHolder.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImultiAssetHolderDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ImultiAssetHolder contract.
type ImultiAssetHolderDepositedIterator struct {
	Event *ImultiAssetHolderDeposited // Event containing the contract specifics and raw log

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
func (it *ImultiAssetHolderDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImultiAssetHolderDeposited)
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
		it.Event = new(ImultiAssetHolderDeposited)
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
func (it *ImultiAssetHolderDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImultiAssetHolderDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImultiAssetHolderDeposited represents a Deposited event raised by the ImultiAssetHolder contract.
type ImultiAssetHolderDeposited struct {
	Destination         [32]byte
	Asset               common.Address
	DestinationHoldings *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) FilterDeposited(opts *bind.FilterOpts, destination [][32]byte) (*ImultiAssetHolderDepositedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ImultiAssetHolder.contract.FilterLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolderDepositedIterator{contract: _ImultiAssetHolder.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ImultiAssetHolderDeposited, destination [][32]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ImultiAssetHolder.contract.WatchLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImultiAssetHolderDeposited)
				if err := _ImultiAssetHolder.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) ParseDeposited(log types.Log) (*ImultiAssetHolderDeposited, error) {
	event := new(ImultiAssetHolderDeposited)
	if err := _ImultiAssetHolder.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImultiAssetHolderReclaimedIterator is returned from FilterReclaimed and is used to iterate over the raw logs and unpacked data for Reclaimed events raised by the ImultiAssetHolder contract.
type ImultiAssetHolderReclaimedIterator struct {
	Event *ImultiAssetHolderReclaimed // Event containing the contract specifics and raw log

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
func (it *ImultiAssetHolderReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImultiAssetHolderReclaimed)
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
		it.Event = new(ImultiAssetHolderReclaimed)
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
func (it *ImultiAssetHolderReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImultiAssetHolderReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImultiAssetHolderReclaimed represents a Reclaimed event raised by the ImultiAssetHolder contract.
type ImultiAssetHolderReclaimed struct {
	ChannelId  [32]byte
	AssetIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReclaimed is a free log retrieval operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) FilterReclaimed(opts *bind.FilterOpts, channelId [][32]byte) (*ImultiAssetHolderReclaimedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _ImultiAssetHolder.contract.FilterLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &ImultiAssetHolderReclaimedIterator{contract: _ImultiAssetHolder.contract, event: "Reclaimed", logs: logs, sub: sub}, nil
}

// WatchReclaimed is a free log subscription operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) WatchReclaimed(opts *bind.WatchOpts, sink chan<- *ImultiAssetHolderReclaimed, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _ImultiAssetHolder.contract.WatchLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImultiAssetHolderReclaimed)
				if err := _ImultiAssetHolder.contract.UnpackLog(event, "Reclaimed", log); err != nil {
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

// ParseReclaimed is a log parse operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_ImultiAssetHolder *ImultiAssetHolderFilterer) ParseReclaimed(log types.Log) (*ImultiAssetHolderReclaimed, error) {
	event := new(ImultiAssetHolderReclaimed)
	if err := _ImultiAssetHolder.contract.UnpackLog(event, "Reclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
