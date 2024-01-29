// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc1271

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

// IERC1271MetaData contains all meta data concerning the IERC1271 contract.
var IERC1271MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC1271ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1271MetaData.ABI instead.
var IERC1271ABI = IERC1271MetaData.ABI

// IERC1271 is an auto generated Go binding around an Ethereum contract.
type IERC1271 struct {
	IERC1271Caller     // Read-only binding to the contract
	IERC1271Transactor // Write-only binding to the contract
	IERC1271Filterer   // Log filterer for contract events
}

// IERC1271Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1271Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1271Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1271Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1271Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1271Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1271Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1271Session struct {
	Contract     *IERC1271         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1271CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1271CallerSession struct {
	Contract *IERC1271Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC1271TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1271TransactorSession struct {
	Contract     *IERC1271Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC1271Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1271Raw struct {
	Contract *IERC1271 // Generic contract binding to access the raw methods on
}

// IERC1271CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1271CallerRaw struct {
	Contract *IERC1271Caller // Generic read-only contract binding to access the raw methods on
}

// IERC1271TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1271TransactorRaw struct {
	Contract *IERC1271Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1271 creates a new instance of IERC1271, bound to a specific deployed contract.
func NewIERC1271(address common.Address, backend bind.ContractBackend) (*IERC1271, error) {
	contract, err := bindIERC1271(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1271{IERC1271Caller: IERC1271Caller{contract: contract}, IERC1271Transactor: IERC1271Transactor{contract: contract}, IERC1271Filterer: IERC1271Filterer{contract: contract}}, nil
}

// NewIERC1271Caller creates a new read-only instance of IERC1271, bound to a specific deployed contract.
func NewIERC1271Caller(address common.Address, caller bind.ContractCaller) (*IERC1271Caller, error) {
	contract, err := bindIERC1271(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1271Caller{contract: contract}, nil
}

// NewIERC1271Transactor creates a new write-only instance of IERC1271, bound to a specific deployed contract.
func NewIERC1271Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1271Transactor, error) {
	contract, err := bindIERC1271(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1271Transactor{contract: contract}, nil
}

// NewIERC1271Filterer creates a new log filterer instance of IERC1271, bound to a specific deployed contract.
func NewIERC1271Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC1271Filterer, error) {
	contract, err := bindIERC1271(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1271Filterer{contract: contract}, nil
}

// bindIERC1271 binds a generic wrapper to an already deployed contract.
func bindIERC1271(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1271MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1271 *IERC1271Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1271.Contract.IERC1271Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1271 *IERC1271Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1271.Contract.IERC1271Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1271 *IERC1271Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1271.Contract.IERC1271Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1271 *IERC1271CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1271.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1271 *IERC1271TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1271.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1271 *IERC1271TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1271.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_IERC1271 *IERC1271Caller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _IERC1271.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_IERC1271 *IERC1271Session) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _IERC1271.Contract.IsValidSignature(&_IERC1271.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_IERC1271 *IERC1271CallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _IERC1271.Contract.IsValidSignature(&_IERC1271.CallOpts, hash, signature)
}
