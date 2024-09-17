// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivoucher_v2

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

// IVoucherVoucher is an auto generated low-level Go binding around an user-defined struct.
type IVoucherVoucher struct {
	ChainId     uint32
	Router      common.Address
	Executor    common.Address
	Beneficiary common.Address
	ExpireAt    uint64
	Nonce       *big.Int
	Data        []byte
	Signature   []byte
}

// IVoucherMetaData contains all meta data concerning the IVoucher contract.
var IVoucherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIssuer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoucherAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoucherExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIVoucher.Voucher\",\"name\":\"voucher\",\"type\":\"tuple\"}],\"name\":\"Used\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expireAt\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIVoucher.Voucher[]\",\"name\":\"vouchers\",\"type\":\"tuple[]\"}],\"name\":\"use\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IVoucherABI is the input ABI used to generate the binding from.
// Deprecated: Use IVoucherMetaData.ABI instead.
var IVoucherABI = IVoucherMetaData.ABI

// IVoucher is an auto generated Go binding around an Ethereum contract.
type IVoucher struct {
	IVoucherCaller     // Read-only binding to the contract
	IVoucherTransactor // Write-only binding to the contract
	IVoucherFilterer   // Log filterer for contract events
}

// IVoucherCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVoucherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVoucherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVoucherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVoucherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVoucherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVoucherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVoucherSession struct {
	Contract     *IVoucher         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVoucherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVoucherCallerSession struct {
	Contract *IVoucherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IVoucherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVoucherTransactorSession struct {
	Contract     *IVoucherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IVoucherRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVoucherRaw struct {
	Contract *IVoucher // Generic contract binding to access the raw methods on
}

// IVoucherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVoucherCallerRaw struct {
	Contract *IVoucherCaller // Generic read-only contract binding to access the raw methods on
}

// IVoucherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVoucherTransactorRaw struct {
	Contract *IVoucherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVoucher creates a new instance of IVoucher, bound to a specific deployed contract.
func NewIVoucher(address common.Address, backend bind.ContractBackend) (*IVoucher, error) {
	contract, err := bindIVoucher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVoucher{IVoucherCaller: IVoucherCaller{contract: contract}, IVoucherTransactor: IVoucherTransactor{contract: contract}, IVoucherFilterer: IVoucherFilterer{contract: contract}}, nil
}

// NewIVoucherCaller creates a new read-only instance of IVoucher, bound to a specific deployed contract.
func NewIVoucherCaller(address common.Address, caller bind.ContractCaller) (*IVoucherCaller, error) {
	contract, err := bindIVoucher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVoucherCaller{contract: contract}, nil
}

// NewIVoucherTransactor creates a new write-only instance of IVoucher, bound to a specific deployed contract.
func NewIVoucherTransactor(address common.Address, transactor bind.ContractTransactor) (*IVoucherTransactor, error) {
	contract, err := bindIVoucher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVoucherTransactor{contract: contract}, nil
}

// NewIVoucherFilterer creates a new log filterer instance of IVoucher, bound to a specific deployed contract.
func NewIVoucherFilterer(address common.Address, filterer bind.ContractFilterer) (*IVoucherFilterer, error) {
	contract, err := bindIVoucher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVoucherFilterer{contract: contract}, nil
}

// bindIVoucher binds a generic wrapper to an already deployed contract.
func bindIVoucher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVoucherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVoucher *IVoucherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVoucher.Contract.IVoucherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVoucher *IVoucherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVoucher.Contract.IVoucherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVoucher *IVoucherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVoucher.Contract.IVoucherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVoucher *IVoucherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVoucher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVoucher *IVoucherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVoucher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVoucher *IVoucherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVoucher.Contract.contract.Transact(opts, method, params...)
}

// Use is a paid mutator transaction binding the contract method 0x142cfda8.
//
// Solidity: function use((uint32,address,address,address,uint64,uint128,bytes,bytes)[] vouchers) returns()
func (_IVoucher *IVoucherTransactor) Use(opts *bind.TransactOpts, vouchers []IVoucherVoucher) (*types.Transaction, error) {
	return _IVoucher.contract.Transact(opts, "use", vouchers)
}

// Use is a paid mutator transaction binding the contract method 0x142cfda8.
//
// Solidity: function use((uint32,address,address,address,uint64,uint128,bytes,bytes)[] vouchers) returns()
func (_IVoucher *IVoucherSession) Use(vouchers []IVoucherVoucher) (*types.Transaction, error) {
	return _IVoucher.Contract.Use(&_IVoucher.TransactOpts, vouchers)
}

// Use is a paid mutator transaction binding the contract method 0x142cfda8.
//
// Solidity: function use((uint32,address,address,address,uint64,uint128,bytes,bytes)[] vouchers) returns()
func (_IVoucher *IVoucherTransactorSession) Use(vouchers []IVoucherVoucher) (*types.Transaction, error) {
	return _IVoucher.Contract.Use(&_IVoucher.TransactOpts, vouchers)
}

// IVoucherUsedIterator is returned from FilterUsed and is used to iterate over the raw logs and unpacked data for Used events raised by the IVoucher contract.
type IVoucherUsedIterator struct {
	Event *IVoucherUsed // Event containing the contract specifics and raw log

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
func (it *IVoucherUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVoucherUsed)
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
		it.Event = new(IVoucherUsed)
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
func (it *IVoucherUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVoucherUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVoucherUsed represents a Used event raised by the IVoucher contract.
type IVoucherUsed struct {
	Voucher IVoucherVoucher
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUsed is a free log retrieval operation binding the contract event 0xe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb0.
//
// Solidity: event Used((uint32,address,address,address,uint64,uint128,bytes,bytes) voucher)
func (_IVoucher *IVoucherFilterer) FilterUsed(opts *bind.FilterOpts) (*IVoucherUsedIterator, error) {

	logs, sub, err := _IVoucher.contract.FilterLogs(opts, "Used")
	if err != nil {
		return nil, err
	}
	return &IVoucherUsedIterator{contract: _IVoucher.contract, event: "Used", logs: logs, sub: sub}, nil
}

// WatchUsed is a free log subscription operation binding the contract event 0xe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb0.
//
// Solidity: event Used((uint32,address,address,address,uint64,uint128,bytes,bytes) voucher)
func (_IVoucher *IVoucherFilterer) WatchUsed(opts *bind.WatchOpts, sink chan<- *IVoucherUsed) (event.Subscription, error) {

	logs, sub, err := _IVoucher.contract.WatchLogs(opts, "Used")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVoucherUsed)
				if err := _IVoucher.contract.UnpackLog(event, "Used", log); err != nil {
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

// ParseUsed is a log parse operation binding the contract event 0xe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb0.
//
// Solidity: event Used((uint32,address,address,address,uint64,uint128,bytes,bytes) voucher)
func (_IVoucher *IVoucherFilterer) ParseUsed(log types.Log) (*IVoucherUsed, error) {
	event := new(IVoucherUsed)
	if err := _IVoucher.contract.UnpackLog(event, "Used", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
