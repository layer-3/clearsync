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

// IMultiAssetHolderMetaData contains all meta data concerning the IMultiAssetHolder contract.
var IMultiAssetHolderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalHoldings\",\"type\":\"uint256\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationHoldings\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeld\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceAssetIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"targetStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"targetOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"targetAssetIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiAssetHolder.ReclaimArgs\",\"name\":\"reclaimArgs\",\"type\":\"tuple\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"fromChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"outcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMultiAssetHolderABI is the input ABI used to generate the binding from.
// Deprecated: Use IMultiAssetHolderMetaData.ABI instead.
var IMultiAssetHolderABI = IMultiAssetHolderMetaData.ABI

// IMultiAssetHolder is an auto generated Go binding around an Ethereum contract.
type IMultiAssetHolder struct {
	IMultiAssetHolderCaller     // Read-only binding to the contract
	IMultiAssetHolderTransactor // Write-only binding to the contract
	IMultiAssetHolderFilterer   // Log filterer for contract events
}

// IMultiAssetHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMultiAssetHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMultiAssetHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMultiAssetHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMultiAssetHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMultiAssetHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMultiAssetHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMultiAssetHolderSession struct {
	Contract     *IMultiAssetHolder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IMultiAssetHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMultiAssetHolderCallerSession struct {
	Contract *IMultiAssetHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IMultiAssetHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMultiAssetHolderTransactorSession struct {
	Contract     *IMultiAssetHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IMultiAssetHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMultiAssetHolderRaw struct {
	Contract *IMultiAssetHolder // Generic contract binding to access the raw methods on
}

// IMultiAssetHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMultiAssetHolderCallerRaw struct {
	Contract *IMultiAssetHolderCaller // Generic read-only contract binding to access the raw methods on
}

// IMultiAssetHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMultiAssetHolderTransactorRaw struct {
	Contract *IMultiAssetHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMultiAssetHolder creates a new instance of IMultiAssetHolder, bound to a specific deployed contract.
func NewIMultiAssetHolder(address common.Address, backend bind.ContractBackend) (*IMultiAssetHolder, error) {
	contract, err := bindIMultiAssetHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolder{IMultiAssetHolderCaller: IMultiAssetHolderCaller{contract: contract}, IMultiAssetHolderTransactor: IMultiAssetHolderTransactor{contract: contract}, IMultiAssetHolderFilterer: IMultiAssetHolderFilterer{contract: contract}}, nil
}

// NewIMultiAssetHolderCaller creates a new read-only instance of IMultiAssetHolder, bound to a specific deployed contract.
func NewIMultiAssetHolderCaller(address common.Address, caller bind.ContractCaller) (*IMultiAssetHolderCaller, error) {
	contract, err := bindIMultiAssetHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolderCaller{contract: contract}, nil
}

// NewIMultiAssetHolderTransactor creates a new write-only instance of IMultiAssetHolder, bound to a specific deployed contract.
func NewIMultiAssetHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*IMultiAssetHolderTransactor, error) {
	contract, err := bindIMultiAssetHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolderTransactor{contract: contract}, nil
}

// NewIMultiAssetHolderFilterer creates a new log filterer instance of IMultiAssetHolder, bound to a specific deployed contract.
func NewIMultiAssetHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*IMultiAssetHolderFilterer, error) {
	contract, err := bindIMultiAssetHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolderFilterer{contract: contract}, nil
}

// bindIMultiAssetHolder binds a generic wrapper to an already deployed contract.
func bindIMultiAssetHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMultiAssetHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMultiAssetHolder *IMultiAssetHolderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMultiAssetHolder.Contract.IMultiAssetHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMultiAssetHolder *IMultiAssetHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.IMultiAssetHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMultiAssetHolder *IMultiAssetHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.IMultiAssetHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMultiAssetHolder *IMultiAssetHolderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMultiAssetHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMultiAssetHolder *IMultiAssetHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMultiAssetHolder *IMultiAssetHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_IMultiAssetHolder *IMultiAssetHolderTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IMultiAssetHolder.contract.Transact(opts, "deposit", asset, destination, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_IMultiAssetHolder *IMultiAssetHolderSession) Deposit(asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.Deposit(&_IMultiAssetHolder.TransactOpts, asset, destination, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_IMultiAssetHolder *IMultiAssetHolderTransactorSession) Deposit(asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.Deposit(&_IMultiAssetHolder.TransactOpts, asset, destination, expectedHeld, amount)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_IMultiAssetHolder *IMultiAssetHolderTransactor) Reclaim(opts *bind.TransactOpts, reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _IMultiAssetHolder.contract.Transact(opts, "reclaim", reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_IMultiAssetHolder *IMultiAssetHolderSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.Reclaim(&_IMultiAssetHolder.TransactOpts, reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_IMultiAssetHolder *IMultiAssetHolderTransactorSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.Reclaim(&_IMultiAssetHolder.TransactOpts, reclaimArgs)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_IMultiAssetHolder *IMultiAssetHolderTransactor) Transfer(opts *bind.TransactOpts, assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _IMultiAssetHolder.contract.Transact(opts, "transfer", assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_IMultiAssetHolder *IMultiAssetHolderSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.Transfer(&_IMultiAssetHolder.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_IMultiAssetHolder *IMultiAssetHolderTransactorSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _IMultiAssetHolder.Contract.Transfer(&_IMultiAssetHolder.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// IMultiAssetHolderAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the IMultiAssetHolder contract.
type IMultiAssetHolderAllocationUpdatedIterator struct {
	Event *IMultiAssetHolderAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *IMultiAssetHolderAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiAssetHolderAllocationUpdated)
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
		it.Event = new(IMultiAssetHolderAllocationUpdated)
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
func (it *IMultiAssetHolderAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiAssetHolderAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiAssetHolderAllocationUpdated represents a AllocationUpdated event raised by the IMultiAssetHolder contract.
type IMultiAssetHolderAllocationUpdated struct {
	ChannelId       [32]byte
	AssetIndex      *big.Int
	InitialHoldings *big.Int
	FinalHoldings   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) FilterAllocationUpdated(opts *bind.FilterOpts, channelId [][32]byte) (*IMultiAssetHolderAllocationUpdatedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IMultiAssetHolder.contract.FilterLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolderAllocationUpdatedIterator{contract: _IMultiAssetHolder.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *IMultiAssetHolderAllocationUpdated, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IMultiAssetHolder.contract.WatchLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiAssetHolderAllocationUpdated)
				if err := _IMultiAssetHolder.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
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
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) ParseAllocationUpdated(log types.Log) (*IMultiAssetHolderAllocationUpdated, error) {
	event := new(IMultiAssetHolderAllocationUpdated)
	if err := _IMultiAssetHolder.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMultiAssetHolderDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IMultiAssetHolder contract.
type IMultiAssetHolderDepositedIterator struct {
	Event *IMultiAssetHolderDeposited // Event containing the contract specifics and raw log

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
func (it *IMultiAssetHolderDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiAssetHolderDeposited)
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
		it.Event = new(IMultiAssetHolderDeposited)
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
func (it *IMultiAssetHolderDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiAssetHolderDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiAssetHolderDeposited represents a Deposited event raised by the IMultiAssetHolder contract.
type IMultiAssetHolderDeposited struct {
	Destination         [32]byte
	Asset               common.Address
	DestinationHoldings *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) FilterDeposited(opts *bind.FilterOpts, destination [][32]byte) (*IMultiAssetHolderDepositedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IMultiAssetHolder.contract.FilterLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolderDepositedIterator{contract: _IMultiAssetHolder.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IMultiAssetHolderDeposited, destination [][32]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IMultiAssetHolder.contract.WatchLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiAssetHolderDeposited)
				if err := _IMultiAssetHolder.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) ParseDeposited(log types.Log) (*IMultiAssetHolderDeposited, error) {
	event := new(IMultiAssetHolderDeposited)
	if err := _IMultiAssetHolder.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMultiAssetHolderReclaimedIterator is returned from FilterReclaimed and is used to iterate over the raw logs and unpacked data for Reclaimed events raised by the IMultiAssetHolder contract.
type IMultiAssetHolderReclaimedIterator struct {
	Event *IMultiAssetHolderReclaimed // Event containing the contract specifics and raw log

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
func (it *IMultiAssetHolderReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMultiAssetHolderReclaimed)
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
		it.Event = new(IMultiAssetHolderReclaimed)
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
func (it *IMultiAssetHolderReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMultiAssetHolderReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMultiAssetHolderReclaimed represents a Reclaimed event raised by the IMultiAssetHolder contract.
type IMultiAssetHolderReclaimed struct {
	ChannelId  [32]byte
	AssetIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReclaimed is a free log retrieval operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) FilterReclaimed(opts *bind.FilterOpts, channelId [][32]byte) (*IMultiAssetHolderReclaimedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IMultiAssetHolder.contract.FilterLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IMultiAssetHolderReclaimedIterator{contract: _IMultiAssetHolder.contract, event: "Reclaimed", logs: logs, sub: sub}, nil
}

// WatchReclaimed is a free log subscription operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) WatchReclaimed(opts *bind.WatchOpts, sink chan<- *IMultiAssetHolderReclaimed, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IMultiAssetHolder.contract.WatchLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMultiAssetHolderReclaimed)
				if err := _IMultiAssetHolder.contract.UnpackLog(event, "Reclaimed", log); err != nil {
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
func (_IMultiAssetHolder *IMultiAssetHolderFilterer) ParseReclaimed(log types.Log) (*IMultiAssetHolderReclaimed, error) {
	event := new(IMultiAssetHolderReclaimed)
	if err := _IMultiAssetHolder.contract.UnpackLog(event, "Reclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
