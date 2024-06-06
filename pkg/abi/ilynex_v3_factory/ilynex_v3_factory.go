// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilynex_v3_factory

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

// ILynexV3FactoryMetaData contains all meta data concerning the ILynexV3Factory contract.
var ILynexV3FactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createPool\",\"inputs\":[{\"name\":\"tokenA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultCommunityFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"farmingAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolByPair\",\"inputs\":[{\"name\":\"tokenA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolDeployer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBaseFeeConfiguration\",\"inputs\":[{\"name\":\"alpha1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"alpha2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"beta1\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"beta2\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gamma1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gamma2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"volumeBeta\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"volumeGamma\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultCommunityFee\",\"inputs\":[{\"name\":\"newDefaultCommunityFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFarmingAddress\",\"inputs\":[{\"name\":\"_farmingAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVaultAddress\",\"inputs\":[{\"name\":\"_vaultAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vaultAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DefaultCommunityFee\",\"inputs\":[{\"name\":\"newDefaultCommunityFee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FarmingAddress\",\"inputs\":[{\"name\":\"newFarmingAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeConfiguration\",\"inputs\":[{\"name\":\"alpha1\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"alpha2\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"beta1\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"beta2\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"gamma1\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"gamma2\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"volumeBeta\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"volumeGamma\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"baseFee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Owner\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Pool\",\"inputs\":[{\"name\":\"token0\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token1\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultAddress\",\"inputs\":[{\"name\":\"newVaultAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ILynexV3FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ILynexV3FactoryMetaData.ABI instead.
var ILynexV3FactoryABI = ILynexV3FactoryMetaData.ABI

// ILynexV3Factory is an auto generated Go binding around an Ethereum contract.
type ILynexV3Factory struct {
	ILynexV3FactoryCaller     // Read-only binding to the contract
	ILynexV3FactoryTransactor // Write-only binding to the contract
	ILynexV3FactoryFilterer   // Log filterer for contract events
}

// ILynexV3FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILynexV3FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexV3FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILynexV3FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexV3FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILynexV3FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexV3FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILynexV3FactorySession struct {
	Contract     *ILynexV3Factory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILynexV3FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILynexV3FactoryCallerSession struct {
	Contract *ILynexV3FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ILynexV3FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILynexV3FactoryTransactorSession struct {
	Contract     *ILynexV3FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ILynexV3FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILynexV3FactoryRaw struct {
	Contract *ILynexV3Factory // Generic contract binding to access the raw methods on
}

// ILynexV3FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILynexV3FactoryCallerRaw struct {
	Contract *ILynexV3FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ILynexV3FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILynexV3FactoryTransactorRaw struct {
	Contract *ILynexV3FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILynexV3Factory creates a new instance of ILynexV3Factory, bound to a specific deployed contract.
func NewILynexV3Factory(address common.Address, backend bind.ContractBackend) (*ILynexV3Factory, error) {
	contract, err := bindILynexV3Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILynexV3Factory{ILynexV3FactoryCaller: ILynexV3FactoryCaller{contract: contract}, ILynexV3FactoryTransactor: ILynexV3FactoryTransactor{contract: contract}, ILynexV3FactoryFilterer: ILynexV3FactoryFilterer{contract: contract}}, nil
}

// NewILynexV3FactoryCaller creates a new read-only instance of ILynexV3Factory, bound to a specific deployed contract.
func NewILynexV3FactoryCaller(address common.Address, caller bind.ContractCaller) (*ILynexV3FactoryCaller, error) {
	contract, err := bindILynexV3Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryCaller{contract: contract}, nil
}

// NewILynexV3FactoryTransactor creates a new write-only instance of ILynexV3Factory, bound to a specific deployed contract.
func NewILynexV3FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ILynexV3FactoryTransactor, error) {
	contract, err := bindILynexV3Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryTransactor{contract: contract}, nil
}

// NewILynexV3FactoryFilterer creates a new log filterer instance of ILynexV3Factory, bound to a specific deployed contract.
func NewILynexV3FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ILynexV3FactoryFilterer, error) {
	contract, err := bindILynexV3Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryFilterer{contract: contract}, nil
}

// bindILynexV3Factory binds a generic wrapper to an already deployed contract.
func bindILynexV3Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILynexV3FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexV3Factory *ILynexV3FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexV3Factory.Contract.ILynexV3FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexV3Factory *ILynexV3FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.ILynexV3FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexV3Factory *ILynexV3FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.ILynexV3FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexV3Factory *ILynexV3FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexV3Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexV3Factory *ILynexV3FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexV3Factory *ILynexV3FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.contract.Transact(opts, method, params...)
}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint16)
func (_ILynexV3Factory *ILynexV3FactoryCaller) DefaultCommunityFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ILynexV3Factory.contract.Call(opts, &out, "defaultCommunityFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint16)
func (_ILynexV3Factory *ILynexV3FactorySession) DefaultCommunityFee() (uint16, error) {
	return _ILynexV3Factory.Contract.DefaultCommunityFee(&_ILynexV3Factory.CallOpts)
}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint16)
func (_ILynexV3Factory *ILynexV3FactoryCallerSession) DefaultCommunityFee() (uint16, error) {
	return _ILynexV3Factory.Contract.DefaultCommunityFee(&_ILynexV3Factory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCaller) FarmingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Factory.contract.Call(opts, &out, "farmingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_ILynexV3Factory *ILynexV3FactorySession) FarmingAddress() (common.Address, error) {
	return _ILynexV3Factory.Contract.FarmingAddress(&_ILynexV3Factory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCallerSession) FarmingAddress() (common.Address, error) {
	return _ILynexV3Factory.Contract.FarmingAddress(&_ILynexV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ILynexV3Factory *ILynexV3FactorySession) Owner() (common.Address, error) {
	return _ILynexV3Factory.Contract.Owner(&_ILynexV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCallerSession) Owner() (common.Address, error) {
	return _ILynexV3Factory.Contract.Owner(&_ILynexV3Factory.CallOpts)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_ILynexV3Factory *ILynexV3FactoryCaller) PoolByPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Factory.contract.Call(opts, &out, "poolByPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_ILynexV3Factory *ILynexV3FactorySession) PoolByPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _ILynexV3Factory.Contract.PoolByPair(&_ILynexV3Factory.CallOpts, tokenA, tokenB)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address tokenA, address tokenB) view returns(address pool)
func (_ILynexV3Factory *ILynexV3FactoryCallerSession) PoolByPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _ILynexV3Factory.Contract.PoolByPair(&_ILynexV3Factory.CallOpts, tokenA, tokenB)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Factory.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_ILynexV3Factory *ILynexV3FactorySession) PoolDeployer() (common.Address, error) {
	return _ILynexV3Factory.Contract.PoolDeployer(&_ILynexV3Factory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCallerSession) PoolDeployer() (common.Address, error) {
	return _ILynexV3Factory.Contract.PoolDeployer(&_ILynexV3Factory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexV3Factory.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_ILynexV3Factory *ILynexV3FactorySession) VaultAddress() (common.Address, error) {
	return _ILynexV3Factory.Contract.VaultAddress(&_ILynexV3Factory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_ILynexV3Factory *ILynexV3FactoryCallerSession) VaultAddress() (common.Address, error) {
	return _ILynexV3Factory.Contract.VaultAddress(&_ILynexV3Factory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_ILynexV3Factory *ILynexV3FactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.contract.Transact(opts, "createPool", tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_ILynexV3Factory *ILynexV3FactorySession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.CreatePool(&_ILynexV3Factory.TransactOpts, tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_ILynexV3Factory *ILynexV3FactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.CreatePool(&_ILynexV3Factory.TransactOpts, tokenA, tokenB)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactor) SetBaseFeeConfiguration(opts *bind.TransactOpts, alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _ILynexV3Factory.contract.Transact(opts, "setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_ILynexV3Factory *ILynexV3FactorySession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetBaseFeeConfiguration(&_ILynexV3Factory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactorSession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetBaseFeeConfiguration(&_ILynexV3Factory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x8d5a8711.
//
// Solidity: function setDefaultCommunityFee(uint16 newDefaultCommunityFee) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactor) SetDefaultCommunityFee(opts *bind.TransactOpts, newDefaultCommunityFee uint16) (*types.Transaction, error) {
	return _ILynexV3Factory.contract.Transact(opts, "setDefaultCommunityFee", newDefaultCommunityFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x8d5a8711.
//
// Solidity: function setDefaultCommunityFee(uint16 newDefaultCommunityFee) returns()
func (_ILynexV3Factory *ILynexV3FactorySession) SetDefaultCommunityFee(newDefaultCommunityFee uint16) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetDefaultCommunityFee(&_ILynexV3Factory.TransactOpts, newDefaultCommunityFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x8d5a8711.
//
// Solidity: function setDefaultCommunityFee(uint16 newDefaultCommunityFee) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactorSession) SetDefaultCommunityFee(newDefaultCommunityFee uint16) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetDefaultCommunityFee(&_ILynexV3Factory.TransactOpts, newDefaultCommunityFee)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactor) SetFarmingAddress(opts *bind.TransactOpts, _farmingAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.contract.Transact(opts, "setFarmingAddress", _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_ILynexV3Factory *ILynexV3FactorySession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetFarmingAddress(&_ILynexV3Factory.TransactOpts, _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactorSession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetFarmingAddress(&_ILynexV3Factory.TransactOpts, _farmingAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_ILynexV3Factory *ILynexV3FactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetOwner(&_ILynexV3Factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetOwner(&_ILynexV3Factory.TransactOpts, _owner)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_ILynexV3Factory *ILynexV3FactorySession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetVaultAddress(&_ILynexV3Factory.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_ILynexV3Factory *ILynexV3FactoryTransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _ILynexV3Factory.Contract.SetVaultAddress(&_ILynexV3Factory.TransactOpts, _vaultAddress)
}

// ILynexV3FactoryDefaultCommunityFeeIterator is returned from FilterDefaultCommunityFee and is used to iterate over the raw logs and unpacked data for DefaultCommunityFee events raised by the ILynexV3Factory contract.
type ILynexV3FactoryDefaultCommunityFeeIterator struct {
	Event *ILynexV3FactoryDefaultCommunityFee // Event containing the contract specifics and raw log

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
func (it *ILynexV3FactoryDefaultCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3FactoryDefaultCommunityFee)
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
		it.Event = new(ILynexV3FactoryDefaultCommunityFee)
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
func (it *ILynexV3FactoryDefaultCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3FactoryDefaultCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3FactoryDefaultCommunityFee represents a DefaultCommunityFee event raised by the ILynexV3Factory contract.
type ILynexV3FactoryDefaultCommunityFee struct {
	NewDefaultCommunityFee uint16
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDefaultCommunityFee is a free log retrieval operation binding the contract event 0x6b5c342391f543846fce47a925e7eba910f7bec232b08633308ca93fdd0fdf0d.
//
// Solidity: event DefaultCommunityFee(uint16 newDefaultCommunityFee)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) FilterDefaultCommunityFee(opts *bind.FilterOpts) (*ILynexV3FactoryDefaultCommunityFeeIterator, error) {

	logs, sub, err := _ILynexV3Factory.contract.FilterLogs(opts, "DefaultCommunityFee")
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryDefaultCommunityFeeIterator{contract: _ILynexV3Factory.contract, event: "DefaultCommunityFee", logs: logs, sub: sub}, nil
}

// WatchDefaultCommunityFee is a free log subscription operation binding the contract event 0x6b5c342391f543846fce47a925e7eba910f7bec232b08633308ca93fdd0fdf0d.
//
// Solidity: event DefaultCommunityFee(uint16 newDefaultCommunityFee)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) WatchDefaultCommunityFee(opts *bind.WatchOpts, sink chan<- *ILynexV3FactoryDefaultCommunityFee) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Factory.contract.WatchLogs(opts, "DefaultCommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3FactoryDefaultCommunityFee)
				if err := _ILynexV3Factory.contract.UnpackLog(event, "DefaultCommunityFee", log); err != nil {
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

// ParseDefaultCommunityFee is a log parse operation binding the contract event 0x6b5c342391f543846fce47a925e7eba910f7bec232b08633308ca93fdd0fdf0d.
//
// Solidity: event DefaultCommunityFee(uint16 newDefaultCommunityFee)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) ParseDefaultCommunityFee(log types.Log) (*ILynexV3FactoryDefaultCommunityFee, error) {
	event := new(ILynexV3FactoryDefaultCommunityFee)
	if err := _ILynexV3Factory.contract.UnpackLog(event, "DefaultCommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3FactoryFarmingAddressIterator is returned from FilterFarmingAddress and is used to iterate over the raw logs and unpacked data for FarmingAddress events raised by the ILynexV3Factory contract.
type ILynexV3FactoryFarmingAddressIterator struct {
	Event *ILynexV3FactoryFarmingAddress // Event containing the contract specifics and raw log

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
func (it *ILynexV3FactoryFarmingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3FactoryFarmingAddress)
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
		it.Event = new(ILynexV3FactoryFarmingAddress)
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
func (it *ILynexV3FactoryFarmingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3FactoryFarmingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3FactoryFarmingAddress represents a FarmingAddress event raised by the ILynexV3Factory contract.
type ILynexV3FactoryFarmingAddress struct {
	NewFarmingAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFarmingAddress is a free log retrieval operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) FilterFarmingAddress(opts *bind.FilterOpts, newFarmingAddress []common.Address) (*ILynexV3FactoryFarmingAddressIterator, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _ILynexV3Factory.contract.FilterLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryFarmingAddressIterator{contract: _ILynexV3Factory.contract, event: "FarmingAddress", logs: logs, sub: sub}, nil
}

// WatchFarmingAddress is a free log subscription operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) WatchFarmingAddress(opts *bind.WatchOpts, sink chan<- *ILynexV3FactoryFarmingAddress, newFarmingAddress []common.Address) (event.Subscription, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _ILynexV3Factory.contract.WatchLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3FactoryFarmingAddress)
				if err := _ILynexV3Factory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
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

// ParseFarmingAddress is a log parse operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) ParseFarmingAddress(log types.Log) (*ILynexV3FactoryFarmingAddress, error) {
	event := new(ILynexV3FactoryFarmingAddress)
	if err := _ILynexV3Factory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3FactoryFeeConfigurationIterator is returned from FilterFeeConfiguration and is used to iterate over the raw logs and unpacked data for FeeConfiguration events raised by the ILynexV3Factory contract.
type ILynexV3FactoryFeeConfigurationIterator struct {
	Event *ILynexV3FactoryFeeConfiguration // Event containing the contract specifics and raw log

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
func (it *ILynexV3FactoryFeeConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3FactoryFeeConfiguration)
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
		it.Event = new(ILynexV3FactoryFeeConfiguration)
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
func (it *ILynexV3FactoryFeeConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3FactoryFeeConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3FactoryFeeConfiguration represents a FeeConfiguration event raised by the ILynexV3Factory contract.
type ILynexV3FactoryFeeConfiguration struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeConfiguration is a free log retrieval operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) FilterFeeConfiguration(opts *bind.FilterOpts) (*ILynexV3FactoryFeeConfigurationIterator, error) {

	logs, sub, err := _ILynexV3Factory.contract.FilterLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryFeeConfigurationIterator{contract: _ILynexV3Factory.contract, event: "FeeConfiguration", logs: logs, sub: sub}, nil
}

// WatchFeeConfiguration is a free log subscription operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) WatchFeeConfiguration(opts *bind.WatchOpts, sink chan<- *ILynexV3FactoryFeeConfiguration) (event.Subscription, error) {

	logs, sub, err := _ILynexV3Factory.contract.WatchLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3FactoryFeeConfiguration)
				if err := _ILynexV3Factory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
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

// ParseFeeConfiguration is a log parse operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) ParseFeeConfiguration(log types.Log) (*ILynexV3FactoryFeeConfiguration, error) {
	event := new(ILynexV3FactoryFeeConfiguration)
	if err := _ILynexV3Factory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3FactoryOwnerIterator is returned from FilterOwner and is used to iterate over the raw logs and unpacked data for Owner events raised by the ILynexV3Factory contract.
type ILynexV3FactoryOwnerIterator struct {
	Event *ILynexV3FactoryOwner // Event containing the contract specifics and raw log

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
func (it *ILynexV3FactoryOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3FactoryOwner)
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
		it.Event = new(ILynexV3FactoryOwner)
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
func (it *ILynexV3FactoryOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3FactoryOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3FactoryOwner represents a Owner event raised by the ILynexV3Factory contract.
type ILynexV3FactoryOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwner is a free log retrieval operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) FilterOwner(opts *bind.FilterOpts, newOwner []common.Address) (*ILynexV3FactoryOwnerIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ILynexV3Factory.contract.FilterLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryOwnerIterator{contract: _ILynexV3Factory.contract, event: "Owner", logs: logs, sub: sub}, nil
}

// WatchOwner is a free log subscription operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) WatchOwner(opts *bind.WatchOpts, sink chan<- *ILynexV3FactoryOwner, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ILynexV3Factory.contract.WatchLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3FactoryOwner)
				if err := _ILynexV3Factory.contract.UnpackLog(event, "Owner", log); err != nil {
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

// ParseOwner is a log parse operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) ParseOwner(log types.Log) (*ILynexV3FactoryOwner, error) {
	event := new(ILynexV3FactoryOwner)
	if err := _ILynexV3Factory.contract.UnpackLog(event, "Owner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3FactoryPoolIterator is returned from FilterPool and is used to iterate over the raw logs and unpacked data for Pool events raised by the ILynexV3Factory contract.
type ILynexV3FactoryPoolIterator struct {
	Event *ILynexV3FactoryPool // Event containing the contract specifics and raw log

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
func (it *ILynexV3FactoryPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3FactoryPool)
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
		it.Event = new(ILynexV3FactoryPool)
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
func (it *ILynexV3FactoryPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3FactoryPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3FactoryPool represents a Pool event raised by the ILynexV3Factory contract.
type ILynexV3FactoryPool struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPool is a free log retrieval operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) FilterPool(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*ILynexV3FactoryPoolIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ILynexV3Factory.contract.FilterLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryPoolIterator{contract: _ILynexV3Factory.contract, event: "Pool", logs: logs, sub: sub}, nil
}

// WatchPool is a free log subscription operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) WatchPool(opts *bind.WatchOpts, sink chan<- *ILynexV3FactoryPool, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ILynexV3Factory.contract.WatchLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3FactoryPool)
				if err := _ILynexV3Factory.contract.UnpackLog(event, "Pool", log); err != nil {
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

// ParsePool is a log parse operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) ParsePool(log types.Log) (*ILynexV3FactoryPool, error) {
	event := new(ILynexV3FactoryPool)
	if err := _ILynexV3Factory.contract.UnpackLog(event, "Pool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILynexV3FactoryVaultAddressIterator is returned from FilterVaultAddress and is used to iterate over the raw logs and unpacked data for VaultAddress events raised by the ILynexV3Factory contract.
type ILynexV3FactoryVaultAddressIterator struct {
	Event *ILynexV3FactoryVaultAddress // Event containing the contract specifics and raw log

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
func (it *ILynexV3FactoryVaultAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexV3FactoryVaultAddress)
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
		it.Event = new(ILynexV3FactoryVaultAddress)
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
func (it *ILynexV3FactoryVaultAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexV3FactoryVaultAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexV3FactoryVaultAddress represents a VaultAddress event raised by the ILynexV3Factory contract.
type ILynexV3FactoryVaultAddress struct {
	NewVaultAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultAddress is a free log retrieval operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) FilterVaultAddress(opts *bind.FilterOpts, newVaultAddress []common.Address) (*ILynexV3FactoryVaultAddressIterator, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _ILynexV3Factory.contract.FilterLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILynexV3FactoryVaultAddressIterator{contract: _ILynexV3Factory.contract, event: "VaultAddress", logs: logs, sub: sub}, nil
}

// WatchVaultAddress is a free log subscription operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) WatchVaultAddress(opts *bind.WatchOpts, sink chan<- *ILynexV3FactoryVaultAddress, newVaultAddress []common.Address) (event.Subscription, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _ILynexV3Factory.contract.WatchLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexV3FactoryVaultAddress)
				if err := _ILynexV3Factory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
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

// ParseVaultAddress is a log parse operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_ILynexV3Factory *ILynexV3FactoryFilterer) ParseVaultAddress(log types.Log) (*ILynexV3FactoryVaultAddress, error) {
	event := new(ILynexV3FactoryVaultAddress)
	if err := _ILynexV3Factory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
