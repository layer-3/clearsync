// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idummy_executor

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

// DummyExecutorMetaData contains all meta data concerning the DummyExecutor contract.
var DummyExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPayload\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b506040516102bc3803806102bc83398101604081905261002f91610059565b6001600081905580546001600160a01b0319166001600160a01b0392909216919091179055610089565b60006020828403121561006b57600080fd5b81516001600160a01b038116811461008257600080fd5b9392505050565b610224806100986000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631cff79cd1461003b578063f887ea4014610050575b600080fd5b61004e61004936600461012e565b61007f565b005b600154610063906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b610087610104565b6001546001600160a01b031633146100b257604051630693060f60e01b815260040160405180910390fd5b826001600160a01b03167f0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd1183836040516100ed9291906101bf565b60405180910390a26100ff6001600055565b505050565b60026000540361012757604051633ee5aeb560e01b815260040160405180910390fd5b6002600055565b60008060006040848603121561014357600080fd5b83356001600160a01b038116811461015a57600080fd5b9250602084013567ffffffffffffffff8082111561017757600080fd5b818601915086601f83011261018b57600080fd5b81358181111561019a57600080fd5b8760208285010111156101ac57600080fd5b6020830194508093505050509250925092565b60208152816020820152818360408301376000818301604090810191909152601f909201601f1916010191905056fea2646970667358221220c0e36b32f30e225abfae9b6f4302d6a52342f3a113d5dfa27e48dc9c07c3b4dd64736f6c63430008170033",
}

// DummyExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyExecutorMetaData.ABI instead.
var DummyExecutorABI = DummyExecutorMetaData.ABI

// DummyExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DummyExecutorMetaData.Bin instead.
var DummyExecutorBin = DummyExecutorMetaData.Bin

// DeployDummyExecutor deploys a new Ethereum contract, binding an instance of DummyExecutor to it.
func DeployDummyExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *DummyExecutor, error) {
	parsed, err := DummyExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DummyExecutorBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DummyExecutor{DummyExecutorCaller: DummyExecutorCaller{contract: contract}, DummyExecutorTransactor: DummyExecutorTransactor{contract: contract}, DummyExecutorFilterer: DummyExecutorFilterer{contract: contract}}, nil
}

// DummyExecutor is an auto generated Go binding around an Ethereum contract.
type DummyExecutor struct {
	DummyExecutorCaller     // Read-only binding to the contract
	DummyExecutorTransactor // Write-only binding to the contract
	DummyExecutorFilterer   // Log filterer for contract events
}

// DummyExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyExecutorSession struct {
	Contract     *DummyExecutor    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyExecutorCallerSession struct {
	Contract *DummyExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DummyExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyExecutorTransactorSession struct {
	Contract     *DummyExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DummyExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyExecutorRaw struct {
	Contract *DummyExecutor // Generic contract binding to access the raw methods on
}

// DummyExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyExecutorCallerRaw struct {
	Contract *DummyExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// DummyExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyExecutorTransactorRaw struct {
	Contract *DummyExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyExecutor creates a new instance of DummyExecutor, bound to a specific deployed contract.
func NewDummyExecutor(address common.Address, backend bind.ContractBackend) (*DummyExecutor, error) {
	contract, err := bindDummyExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyExecutor{DummyExecutorCaller: DummyExecutorCaller{contract: contract}, DummyExecutorTransactor: DummyExecutorTransactor{contract: contract}, DummyExecutorFilterer: DummyExecutorFilterer{contract: contract}}, nil
}

// NewDummyExecutorCaller creates a new read-only instance of DummyExecutor, bound to a specific deployed contract.
func NewDummyExecutorCaller(address common.Address, caller bind.ContractCaller) (*DummyExecutorCaller, error) {
	contract, err := bindDummyExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyExecutorCaller{contract: contract}, nil
}

// NewDummyExecutorTransactor creates a new write-only instance of DummyExecutor, bound to a specific deployed contract.
func NewDummyExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyExecutorTransactor, error) {
	contract, err := bindDummyExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyExecutorTransactor{contract: contract}, nil
}

// NewDummyExecutorFilterer creates a new log filterer instance of DummyExecutor, bound to a specific deployed contract.
func NewDummyExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyExecutorFilterer, error) {
	contract, err := bindDummyExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyExecutorFilterer{contract: contract}, nil
}

// bindDummyExecutor binds a generic wrapper to an already deployed contract.
func bindDummyExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyExecutor *DummyExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyExecutor.Contract.DummyExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyExecutor *DummyExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyExecutor.Contract.DummyExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyExecutor *DummyExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyExecutor.Contract.DummyExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyExecutor *DummyExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyExecutor *DummyExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyExecutor *DummyExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyExecutor.Contract.contract.Transact(opts, method, params...)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_DummyExecutor *DummyExecutorCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DummyExecutor.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_DummyExecutor *DummyExecutorSession) Router() (common.Address, error) {
	return _DummyExecutor.Contract.Router(&_DummyExecutor.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_DummyExecutor *DummyExecutorCallerSession) Router() (common.Address, error) {
	return _DummyExecutor.Contract.Router(&_DummyExecutor.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address beneficiary, bytes data) returns()
func (_DummyExecutor *DummyExecutorTransactor) Execute(opts *bind.TransactOpts, beneficiary common.Address, data []byte) (*types.Transaction, error) {
	return _DummyExecutor.contract.Transact(opts, "execute", beneficiary, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address beneficiary, bytes data) returns()
func (_DummyExecutor *DummyExecutorSession) Execute(beneficiary common.Address, data []byte) (*types.Transaction, error) {
	return _DummyExecutor.Contract.Execute(&_DummyExecutor.TransactOpts, beneficiary, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address beneficiary, bytes data) returns()
func (_DummyExecutor *DummyExecutorTransactorSession) Execute(beneficiary common.Address, data []byte) (*types.Transaction, error) {
	return _DummyExecutor.Contract.Execute(&_DummyExecutor.TransactOpts, beneficiary, data)
}

// DummyExecutorExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the DummyExecutor contract.
type DummyExecutorExecutedIterator struct {
	Event *DummyExecutorExecuted // Event containing the contract specifics and raw log

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
func (it *DummyExecutorExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyExecutorExecuted)
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
		it.Event = new(DummyExecutorExecuted)
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
func (it *DummyExecutorExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyExecutorExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyExecutorExecuted represents a Executed event raised by the DummyExecutor contract.
type DummyExecutorExecuted struct {
	Beneficiary common.Address
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd11.
//
// Solidity: event Executed(address indexed beneficiary, bytes data)
func (_DummyExecutor *DummyExecutorFilterer) FilterExecuted(opts *bind.FilterOpts, beneficiary []common.Address) (*DummyExecutorExecutedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _DummyExecutor.contract.FilterLogs(opts, "Executed", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &DummyExecutorExecutedIterator{contract: _DummyExecutor.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd11.
//
// Solidity: event Executed(address indexed beneficiary, bytes data)
func (_DummyExecutor *DummyExecutorFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *DummyExecutorExecuted, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _DummyExecutor.contract.WatchLogs(opts, "Executed", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyExecutorExecuted)
				if err := _DummyExecutor.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd11.
//
// Solidity: event Executed(address indexed beneficiary, bytes data)
func (_DummyExecutor *DummyExecutorFilterer) ParseExecuted(log types.Log) (*DummyExecutorExecuted, error) {
	event := new(DummyExecutorExecuted)
	if err := _DummyExecutor.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
