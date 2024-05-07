// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivoucher

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
	Target          common.Address
	Action          uint8
	Beneficiary     common.Address
	Expire          uint64
	ChainId         uint32
	VoucherCodeHash [32]byte
	EncodedParams   []byte
}

// IVoucherMetaData contains all meta data concerning the IVoucher contract.
var IVoucherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"IncorrectSigner\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expire\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"voucherCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedParams\",\"type\":\"bytes\"}],\"internalType\":\"structIVoucher.Voucher\",\"name\":\"voucher\",\"type\":\"tuple\"}],\"name\":\"InvalidVoucher\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"voucherCodeHash\",\"type\":\"bytes32\"}],\"name\":\"VoucherAlreadyUsed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"voucherCodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"VoucherUsed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expire\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"voucherCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedParams\",\"type\":\"bytes\"}],\"internalType\":\"structIVoucher.Voucher\",\"name\":\"voucher\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"useVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expire\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"voucherCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedParams\",\"type\":\"bytes\"}],\"internalType\":\"structIVoucher.Voucher[]\",\"name\":\"vouchers\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"useVouchers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// UseVoucher is a paid mutator transaction binding the contract method 0xb91d544e.
//
// Solidity: function useVoucher((address,uint8,address,uint64,uint32,bytes32,bytes) voucher, bytes signature) returns()
func (_IVoucher *IVoucherTransactor) UseVoucher(opts *bind.TransactOpts, voucher IVoucherVoucher, signature []byte) (*types.Transaction, error) {
	return _IVoucher.contract.Transact(opts, "useVoucher", voucher, signature)
}

// UseVoucher is a paid mutator transaction binding the contract method 0xb91d544e.
//
// Solidity: function useVoucher((address,uint8,address,uint64,uint32,bytes32,bytes) voucher, bytes signature) returns()
func (_IVoucher *IVoucherSession) UseVoucher(voucher IVoucherVoucher, signature []byte) (*types.Transaction, error) {
	return _IVoucher.Contract.UseVoucher(&_IVoucher.TransactOpts, voucher, signature)
}

// UseVoucher is a paid mutator transaction binding the contract method 0xb91d544e.
//
// Solidity: function useVoucher((address,uint8,address,uint64,uint32,bytes32,bytes) voucher, bytes signature) returns()
func (_IVoucher *IVoucherTransactorSession) UseVoucher(voucher IVoucherVoucher, signature []byte) (*types.Transaction, error) {
	return _IVoucher.Contract.UseVoucher(&_IVoucher.TransactOpts, voucher, signature)
}

// UseVouchers is a paid mutator transaction binding the contract method 0x268f3b25.
//
// Solidity: function useVouchers((address,uint8,address,uint64,uint32,bytes32,bytes)[] vouchers, bytes signature) returns()
func (_IVoucher *IVoucherTransactor) UseVouchers(opts *bind.TransactOpts, vouchers []IVoucherVoucher, signature []byte) (*types.Transaction, error) {
	return _IVoucher.contract.Transact(opts, "useVouchers", vouchers, signature)
}

// UseVouchers is a paid mutator transaction binding the contract method 0x268f3b25.
//
// Solidity: function useVouchers((address,uint8,address,uint64,uint32,bytes32,bytes)[] vouchers, bytes signature) returns()
func (_IVoucher *IVoucherSession) UseVouchers(vouchers []IVoucherVoucher, signature []byte) (*types.Transaction, error) {
	return _IVoucher.Contract.UseVouchers(&_IVoucher.TransactOpts, vouchers, signature)
}

// UseVouchers is a paid mutator transaction binding the contract method 0x268f3b25.
//
// Solidity: function useVouchers((address,uint8,address,uint64,uint32,bytes32,bytes)[] vouchers, bytes signature) returns()
func (_IVoucher *IVoucherTransactorSession) UseVouchers(vouchers []IVoucherVoucher, signature []byte) (*types.Transaction, error) {
	return _IVoucher.Contract.UseVouchers(&_IVoucher.TransactOpts, vouchers, signature)
}

// IVoucherVoucherUsedIterator is returned from FilterVoucherUsed and is used to iterate over the raw logs and unpacked data for VoucherUsed events raised by the IVoucher contract.
type IVoucherVoucherUsedIterator struct {
	Event *IVoucherVoucherUsed // Event containing the contract specifics and raw log

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
func (it *IVoucherVoucherUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVoucherVoucherUsed)
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
		it.Event = new(IVoucherVoucherUsed)
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
func (it *IVoucherVoucherUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVoucherVoucherUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVoucherVoucherUsed represents a VoucherUsed event raised by the IVoucher contract.
type IVoucherVoucherUsed struct {
	Wallet          common.Address
	Action          uint8
	VoucherCodeHash [32]byte
	ChainId         uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVoucherUsed is a free log retrieval operation binding the contract event 0x7697a9d8fa2cf7dd6cb14a34b76d2e6577843156db8afa185d9181e28f607ba9.
//
// Solidity: event VoucherUsed(address wallet, uint8 action, bytes32 voucherCodeHash, uint32 chainId)
func (_IVoucher *IVoucherFilterer) FilterVoucherUsed(opts *bind.FilterOpts) (*IVoucherVoucherUsedIterator, error) {

	logs, sub, err := _IVoucher.contract.FilterLogs(opts, "VoucherUsed")
	if err != nil {
		return nil, err
	}
	return &IVoucherVoucherUsedIterator{contract: _IVoucher.contract, event: "VoucherUsed", logs: logs, sub: sub}, nil
}

// WatchVoucherUsed is a free log subscription operation binding the contract event 0x7697a9d8fa2cf7dd6cb14a34b76d2e6577843156db8afa185d9181e28f607ba9.
//
// Solidity: event VoucherUsed(address wallet, uint8 action, bytes32 voucherCodeHash, uint32 chainId)
func (_IVoucher *IVoucherFilterer) WatchVoucherUsed(opts *bind.WatchOpts, sink chan<- *IVoucherVoucherUsed) (event.Subscription, error) {

	logs, sub, err := _IVoucher.contract.WatchLogs(opts, "VoucherUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVoucherVoucherUsed)
				if err := _IVoucher.contract.UnpackLog(event, "VoucherUsed", log); err != nil {
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

// ParseVoucherUsed is a log parse operation binding the contract event 0x7697a9d8fa2cf7dd6cb14a34b76d2e6577843156db8afa185d9181e28f607ba9.
//
// Solidity: event VoucherUsed(address wallet, uint8 action, bytes32 voucherCodeHash, uint32 chainId)
func (_IVoucher *IVoucherFilterer) ParseVoucherUsed(log types.Log) (*IVoucherVoucherUsed, error) {
	event := new(IVoucherVoucherUsed)
	if err := _IVoucher.contract.UnpackLog(event, "VoucherUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
