// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilynex_factory

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

// ILynexFactoryMetaData contains all meta data concerning the ILynexFactory contract.
var ILynexFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REFERRAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFERRAL_FEE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dibs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingFeeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dibs\",\"type\":\"address\"}],\"name\":\"setDibs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_refFee\",\"type\":\"uint256\"}],\"name\":\"setReferralFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatileFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ILynexFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ILynexFactoryMetaData.ABI instead.
var ILynexFactoryABI = ILynexFactoryMetaData.ABI

// ILynexFactory is an auto generated Go binding around an Ethereum contract.
type ILynexFactory struct {
	ILynexFactoryCaller     // Read-only binding to the contract
	ILynexFactoryTransactor // Write-only binding to the contract
	ILynexFactoryFilterer   // Log filterer for contract events
}

// ILynexFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILynexFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILynexFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILynexFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILynexFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILynexFactorySession struct {
	Contract     *ILynexFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILynexFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILynexFactoryCallerSession struct {
	Contract *ILynexFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ILynexFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILynexFactoryTransactorSession struct {
	Contract     *ILynexFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ILynexFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILynexFactoryRaw struct {
	Contract *ILynexFactory // Generic contract binding to access the raw methods on
}

// ILynexFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILynexFactoryCallerRaw struct {
	Contract *ILynexFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ILynexFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILynexFactoryTransactorRaw struct {
	Contract *ILynexFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILynexFactory creates a new instance of ILynexFactory, bound to a specific deployed contract.
func NewILynexFactory(address common.Address, backend bind.ContractBackend) (*ILynexFactory, error) {
	contract, err := bindILynexFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILynexFactory{ILynexFactoryCaller: ILynexFactoryCaller{contract: contract}, ILynexFactoryTransactor: ILynexFactoryTransactor{contract: contract}, ILynexFactoryFilterer: ILynexFactoryFilterer{contract: contract}}, nil
}

// NewILynexFactoryCaller creates a new read-only instance of ILynexFactory, bound to a specific deployed contract.
func NewILynexFactoryCaller(address common.Address, caller bind.ContractCaller) (*ILynexFactoryCaller, error) {
	contract, err := bindILynexFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexFactoryCaller{contract: contract}, nil
}

// NewILynexFactoryTransactor creates a new write-only instance of ILynexFactory, bound to a specific deployed contract.
func NewILynexFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ILynexFactoryTransactor, error) {
	contract, err := bindILynexFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILynexFactoryTransactor{contract: contract}, nil
}

// NewILynexFactoryFilterer creates a new log filterer instance of ILynexFactory, bound to a specific deployed contract.
func NewILynexFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ILynexFactoryFilterer, error) {
	contract, err := bindILynexFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILynexFactoryFilterer{contract: contract}, nil
}

// bindILynexFactory binds a generic wrapper to an already deployed contract.
func bindILynexFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILynexFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexFactory *ILynexFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexFactory.Contract.ILynexFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexFactory *ILynexFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexFactory.Contract.ILynexFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexFactory *ILynexFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexFactory.Contract.ILynexFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILynexFactory *ILynexFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILynexFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILynexFactory *ILynexFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILynexFactory *ILynexFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILynexFactory.Contract.contract.Transact(opts, method, params...)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) MAXFEE() (*big.Int, error) {
	return _ILynexFactory.Contract.MAXFEE(&_ILynexFactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) MAXFEE() (*big.Int, error) {
	return _ILynexFactory.Contract.MAXFEE(&_ILynexFactory.CallOpts)
}

// MAXREFERRALFEE is a free data retrieval call binding the contract method 0x1e61079c.
//
// Solidity: function MAX_REFERRAL_FEE() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) MAXREFERRALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "MAX_REFERRAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREFERRALFEE is a free data retrieval call binding the contract method 0x1e61079c.
//
// Solidity: function MAX_REFERRAL_FEE() view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) MAXREFERRALFEE() (*big.Int, error) {
	return _ILynexFactory.Contract.MAXREFERRALFEE(&_ILynexFactory.CallOpts)
}

// MAXREFERRALFEE is a free data retrieval call binding the contract method 0x1e61079c.
//
// Solidity: function MAX_REFERRAL_FEE() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) MAXREFERRALFEE() (*big.Int, error) {
	return _ILynexFactory.Contract.MAXREFERRALFEE(&_ILynexFactory.CallOpts)
}

// REFERRALFEELIMIT is a free data retrieval call binding the contract method 0x63257389.
//
// Solidity: function REFERRAL_FEE_LIMIT() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) REFERRALFEELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "REFERRAL_FEE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REFERRALFEELIMIT is a free data retrieval call binding the contract method 0x63257389.
//
// Solidity: function REFERRAL_FEE_LIMIT() view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) REFERRALFEELIMIT() (*big.Int, error) {
	return _ILynexFactory.Contract.REFERRALFEELIMIT(&_ILynexFactory.CallOpts)
}

// REFERRALFEELIMIT is a free data retrieval call binding the contract method 0x63257389.
//
// Solidity: function REFERRAL_FEE_LIMIT() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) REFERRALFEELIMIT() (*big.Int, error) {
	return _ILynexFactory.Contract.REFERRALFEELIMIT(&_ILynexFactory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_ILynexFactory *ILynexFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_ILynexFactory *ILynexFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _ILynexFactory.Contract.AllPairs(&_ILynexFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_ILynexFactory *ILynexFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _ILynexFactory.Contract.AllPairs(&_ILynexFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) AllPairsLength() (*big.Int, error) {
	return _ILynexFactory.Contract.AllPairsLength(&_ILynexFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _ILynexFactory.Contract.AllPairsLength(&_ILynexFactory.CallOpts)
}

// Dibs is a free data retrieval call binding the contract method 0x7be1623e.
//
// Solidity: function dibs() view returns(address)
func (_ILynexFactory *ILynexFactoryCaller) Dibs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "dibs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dibs is a free data retrieval call binding the contract method 0x7be1623e.
//
// Solidity: function dibs() view returns(address)
func (_ILynexFactory *ILynexFactorySession) Dibs() (common.Address, error) {
	return _ILynexFactory.Contract.Dibs(&_ILynexFactory.CallOpts)
}

// Dibs is a free data retrieval call binding the contract method 0x7be1623e.
//
// Solidity: function dibs() view returns(address)
func (_ILynexFactory *ILynexFactoryCallerSession) Dibs() (common.Address, error) {
	return _ILynexFactory.Contract.Dibs(&_ILynexFactory.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_ILynexFactory *ILynexFactoryCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_ILynexFactory *ILynexFactorySession) FeeManager() (common.Address, error) {
	return _ILynexFactory.Contract.FeeManager(&_ILynexFactory.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_ILynexFactory *ILynexFactoryCallerSession) FeeManager() (common.Address, error) {
	return _ILynexFactory.Contract.FeeManager(&_ILynexFactory.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0x512b45ea.
//
// Solidity: function getFee(bool _stable) view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) GetFee(opts *bind.CallOpts, _stable bool) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "getFee", _stable)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0x512b45ea.
//
// Solidity: function getFee(bool _stable) view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) GetFee(_stable bool) (*big.Int, error) {
	return _ILynexFactory.Contract.GetFee(&_ILynexFactory.CallOpts, _stable)
}

// GetFee is a free data retrieval call binding the contract method 0x512b45ea.
//
// Solidity: function getFee(bool _stable) view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) GetFee(_stable bool) (*big.Int, error) {
	return _ILynexFactory.Contract.GetFee(&_ILynexFactory.CallOpts, _stable)
}

// GetInitializable is a free data retrieval call binding the contract method 0xeb13c4cf.
//
// Solidity: function getInitializable() view returns(address, address, bool)
func (_ILynexFactory *ILynexFactoryCaller) GetInitializable(opts *bind.CallOpts) (common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "getInitializable")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetInitializable is a free data retrieval call binding the contract method 0xeb13c4cf.
//
// Solidity: function getInitializable() view returns(address, address, bool)
func (_ILynexFactory *ILynexFactorySession) GetInitializable() (common.Address, common.Address, bool, error) {
	return _ILynexFactory.Contract.GetInitializable(&_ILynexFactory.CallOpts)
}

// GetInitializable is a free data retrieval call binding the contract method 0xeb13c4cf.
//
// Solidity: function getInitializable() view returns(address, address, bool)
func (_ILynexFactory *ILynexFactoryCallerSession) GetInitializable() (common.Address, common.Address, bool, error) {
	return _ILynexFactory.Contract.GetInitializable(&_ILynexFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address , address , bool ) view returns(address)
func (_ILynexFactory *ILynexFactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 bool) (common.Address, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "getPair", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address , address , bool ) view returns(address)
func (_ILynexFactory *ILynexFactorySession) GetPair(arg0 common.Address, arg1 common.Address, arg2 bool) (common.Address, error) {
	return _ILynexFactory.Contract.GetPair(&_ILynexFactory.CallOpts, arg0, arg1, arg2)
}

// GetPair is a free data retrieval call binding the contract method 0x6801cc30.
//
// Solidity: function getPair(address , address , bool ) view returns(address)
func (_ILynexFactory *ILynexFactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address, arg2 bool) (common.Address, error) {
	return _ILynexFactory.Contract.GetPair(&_ILynexFactory.CallOpts, arg0, arg1, arg2)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address ) view returns(bool)
func (_ILynexFactory *ILynexFactoryCaller) IsPair(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "isPair", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address ) view returns(bool)
func (_ILynexFactory *ILynexFactorySession) IsPair(arg0 common.Address) (bool, error) {
	return _ILynexFactory.Contract.IsPair(&_ILynexFactory.CallOpts, arg0)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address ) view returns(bool)
func (_ILynexFactory *ILynexFactoryCallerSession) IsPair(arg0 common.Address) (bool, error) {
	return _ILynexFactory.Contract.IsPair(&_ILynexFactory.CallOpts, arg0)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_ILynexFactory *ILynexFactoryCaller) PairCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "pairCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_ILynexFactory *ILynexFactorySession) PairCodeHash() ([32]byte, error) {
	return _ILynexFactory.Contract.PairCodeHash(&_ILynexFactory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_ILynexFactory *ILynexFactoryCallerSession) PairCodeHash() ([32]byte, error) {
	return _ILynexFactory.Contract.PairCodeHash(&_ILynexFactory.CallOpts)
}

// Pairs is a free data retrieval call binding the contract method 0xffb0a4a0.
//
// Solidity: function pairs() view returns(address[])
func (_ILynexFactory *ILynexFactoryCaller) Pairs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "pairs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Pairs is a free data retrieval call binding the contract method 0xffb0a4a0.
//
// Solidity: function pairs() view returns(address[])
func (_ILynexFactory *ILynexFactorySession) Pairs() ([]common.Address, error) {
	return _ILynexFactory.Contract.Pairs(&_ILynexFactory.CallOpts)
}

// Pairs is a free data retrieval call binding the contract method 0xffb0a4a0.
//
// Solidity: function pairs() view returns(address[])
func (_ILynexFactory *ILynexFactoryCallerSession) Pairs() ([]common.Address, error) {
	return _ILynexFactory.Contract.Pairs(&_ILynexFactory.CallOpts)
}

// PendingFeeManager is a free data retrieval call binding the contract method 0x8a4fa0d2.
//
// Solidity: function pendingFeeManager() view returns(address)
func (_ILynexFactory *ILynexFactoryCaller) PendingFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "pendingFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingFeeManager is a free data retrieval call binding the contract method 0x8a4fa0d2.
//
// Solidity: function pendingFeeManager() view returns(address)
func (_ILynexFactory *ILynexFactorySession) PendingFeeManager() (common.Address, error) {
	return _ILynexFactory.Contract.PendingFeeManager(&_ILynexFactory.CallOpts)
}

// PendingFeeManager is a free data retrieval call binding the contract method 0x8a4fa0d2.
//
// Solidity: function pendingFeeManager() view returns(address)
func (_ILynexFactory *ILynexFactoryCallerSession) PendingFeeManager() (common.Address, error) {
	return _ILynexFactory.Contract.PendingFeeManager(&_ILynexFactory.CallOpts)
}

// StableFee is a free data retrieval call binding the contract method 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) StableFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "stableFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFee is a free data retrieval call binding the contract method 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) StableFee() (*big.Int, error) {
	return _ILynexFactory.Contract.StableFee(&_ILynexFactory.CallOpts)
}

// StableFee is a free data retrieval call binding the contract method 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) StableFee() (*big.Int, error) {
	return _ILynexFactory.Contract.StableFee(&_ILynexFactory.CallOpts)
}

// VolatileFee is a free data retrieval call binding the contract method 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCaller) VolatileFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILynexFactory.contract.Call(opts, &out, "volatileFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatileFee is a free data retrieval call binding the contract method 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (_ILynexFactory *ILynexFactorySession) VolatileFee() (*big.Int, error) {
	return _ILynexFactory.Contract.VolatileFee(&_ILynexFactory.CallOpts)
}

// VolatileFee is a free data retrieval call binding the contract method 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (_ILynexFactory *ILynexFactoryCallerSession) VolatileFee() (*big.Int, error) {
	return _ILynexFactory.Contract.VolatileFee(&_ILynexFactory.CallOpts)
}

// AcceptFeeManager is a paid mutator transaction binding the contract method 0xf94c53c7.
//
// Solidity: function acceptFeeManager() returns()
func (_ILynexFactory *ILynexFactoryTransactor) AcceptFeeManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILynexFactory.contract.Transact(opts, "acceptFeeManager")
}

// AcceptFeeManager is a paid mutator transaction binding the contract method 0xf94c53c7.
//
// Solidity: function acceptFeeManager() returns()
func (_ILynexFactory *ILynexFactorySession) AcceptFeeManager() (*types.Transaction, error) {
	return _ILynexFactory.Contract.AcceptFeeManager(&_ILynexFactory.TransactOpts)
}

// AcceptFeeManager is a paid mutator transaction binding the contract method 0xf94c53c7.
//
// Solidity: function acceptFeeManager() returns()
func (_ILynexFactory *ILynexFactoryTransactorSession) AcceptFeeManager() (*types.Transaction, error) {
	return _ILynexFactory.Contract.AcceptFeeManager(&_ILynexFactory.TransactOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pair)
func (_ILynexFactory *ILynexFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _ILynexFactory.contract.Transact(opts, "createPair", tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pair)
func (_ILynexFactory *ILynexFactorySession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _ILynexFactory.Contract.CreatePair(&_ILynexFactory.TransactOpts, tokenA, tokenB, stable)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address tokenA, address tokenB, bool stable) returns(address pair)
func (_ILynexFactory *ILynexFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address, stable bool) (*types.Transaction, error) {
	return _ILynexFactory.Contract.CreatePair(&_ILynexFactory.TransactOpts, tokenA, tokenB, stable)
}

// SetDibs is a paid mutator transaction binding the contract method 0x0c74db12.
//
// Solidity: function setDibs(address _dibs) returns()
func (_ILynexFactory *ILynexFactoryTransactor) SetDibs(opts *bind.TransactOpts, _dibs common.Address) (*types.Transaction, error) {
	return _ILynexFactory.contract.Transact(opts, "setDibs", _dibs)
}

// SetDibs is a paid mutator transaction binding the contract method 0x0c74db12.
//
// Solidity: function setDibs(address _dibs) returns()
func (_ILynexFactory *ILynexFactorySession) SetDibs(_dibs common.Address) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetDibs(&_ILynexFactory.TransactOpts, _dibs)
}

// SetDibs is a paid mutator transaction binding the contract method 0x0c74db12.
//
// Solidity: function setDibs(address _dibs) returns()
func (_ILynexFactory *ILynexFactoryTransactorSession) SetDibs(_dibs common.Address) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetDibs(&_ILynexFactory.TransactOpts, _dibs)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_ILynexFactory *ILynexFactoryTransactor) SetFee(opts *bind.TransactOpts, _stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _ILynexFactory.contract.Transact(opts, "setFee", _stable, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_ILynexFactory *ILynexFactorySession) SetFee(_stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetFee(&_ILynexFactory.TransactOpts, _stable, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xe1f76b44.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (_ILynexFactory *ILynexFactoryTransactorSession) SetFee(_stable bool, _fee *big.Int) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetFee(&_ILynexFactory.TransactOpts, _stable, _fee)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_ILynexFactory *ILynexFactoryTransactor) SetFeeManager(opts *bind.TransactOpts, _feeManager common.Address) (*types.Transaction, error) {
	return _ILynexFactory.contract.Transact(opts, "setFeeManager", _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_ILynexFactory *ILynexFactorySession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetFeeManager(&_ILynexFactory.TransactOpts, _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_ILynexFactory *ILynexFactoryTransactorSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetFeeManager(&_ILynexFactory.TransactOpts, _feeManager)
}

// SetReferralFee is a paid mutator transaction binding the contract method 0x713494d7.
//
// Solidity: function setReferralFee(uint256 _refFee) returns()
func (_ILynexFactory *ILynexFactoryTransactor) SetReferralFee(opts *bind.TransactOpts, _refFee *big.Int) (*types.Transaction, error) {
	return _ILynexFactory.contract.Transact(opts, "setReferralFee", _refFee)
}

// SetReferralFee is a paid mutator transaction binding the contract method 0x713494d7.
//
// Solidity: function setReferralFee(uint256 _refFee) returns()
func (_ILynexFactory *ILynexFactorySession) SetReferralFee(_refFee *big.Int) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetReferralFee(&_ILynexFactory.TransactOpts, _refFee)
}

// SetReferralFee is a paid mutator transaction binding the contract method 0x713494d7.
//
// Solidity: function setReferralFee(uint256 _refFee) returns()
func (_ILynexFactory *ILynexFactoryTransactorSession) SetReferralFee(_refFee *big.Int) (*types.Transaction, error) {
	return _ILynexFactory.Contract.SetReferralFee(&_ILynexFactory.TransactOpts, _refFee)
}

// ILynexFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the ILynexFactory contract.
type ILynexFactoryPairCreatedIterator struct {
	Event *ILynexFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *ILynexFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILynexFactoryPairCreated)
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
		it.Event = new(ILynexFactoryPairCreated)
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
func (it *ILynexFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILynexFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILynexFactoryPairCreated represents a PairCreated event raised by the ILynexFactory contract.
type ILynexFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Stable bool
	Pair   common.Address
	Arg4   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xc4805696c66d7cf352fc1d6bb633ad5ee82f6cb577c453024b6e0eb8306c6fc9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, bool stable, address pair, uint256 arg4)
func (_ILynexFactory *ILynexFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*ILynexFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ILynexFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &ILynexFactoryPairCreatedIterator{contract: _ILynexFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xc4805696c66d7cf352fc1d6bb633ad5ee82f6cb577c453024b6e0eb8306c6fc9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, bool stable, address pair, uint256 arg4)
func (_ILynexFactory *ILynexFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *ILynexFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ILynexFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILynexFactoryPairCreated)
				if err := _ILynexFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xc4805696c66d7cf352fc1d6bb633ad5ee82f6cb577c453024b6e0eb8306c6fc9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, bool stable, address pair, uint256 arg4)
func (_ILynexFactory *ILynexFactoryFilterer) ParsePairCreated(log types.Log) (*ILynexFactoryPairCreated, error) {
	event := new(ILynexFactoryPairCreated)
	if err := _ILynexFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
