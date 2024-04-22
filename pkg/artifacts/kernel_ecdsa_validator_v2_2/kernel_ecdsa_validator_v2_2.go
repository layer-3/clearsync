// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kernel_ecdsa_validator_v2_2

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// KernelECDSAValidatorMetaData contains all meta data concerning the KernelECDSAValidator contract.
var KernelECDSAValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kernel\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ecdsaValidatorStorage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"ValidationData\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"_userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"ValidationData\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657610639908161001c8239f35b600080fdfe60406080815260048036101561001457600080fd5b600091823560e01c80630c9595561461028357806320709efc1461021f578063333daf92146101c15780633a871cdd14610155578381638fc925aa146100f25750639ea9bd591461006457600080fd5b346100ee57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ee5761009a610382565b9160243567ffffffffffffffff81116100ea57936100be839260209636910161034f565b505033815280855273ffffffffffffffffffffffffffffffffffffffff91829120541691519216148152f35b8480fd5b8280fd5b92905060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101515780359067ffffffffffffffff821161014c5761013d9136910161034f565b50503382528160205281205580f35b505050fd5b5050fd5b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc836060368301126101ba5783359167ffffffffffffffff83116101bd576101609083360301126101ba57506020926101b3916024359101610496565b9051908152f35b80fd5b5080fd5b5082346101ba57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101ba576024359067ffffffffffffffff82116101ba57506020926102186101b3923690830161034f565b9135610598565b8382346101bd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bd576020918173ffffffffffffffffffffffffffffffffffffffff9182610273610382565b1681528085522054169051908152f35b509060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ee5780359067ffffffffffffffff821161034b576102cd9136910161034f565b6014116100ee5773ffffffffffffffffffffffffffffffffffffffff903560601c91338452836020528320805490837fffffffffffffffffffffffff0000000000000000000000000000000000000000831617905516337f381c0d11398486654573703c51ee8210ce9461764d133f9f0e53b6a5397053318480a480f35b8380fd5b9181601f8401121561037d5782359167ffffffffffffffff831161037d576020838186019501011161037d57565b600080fd5b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361037d57565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561037d570180359067ffffffffffffffff821161037d5760200191813603831361037d57565b92919267ffffffffffffffff9182811161046757604051927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f81601f8501160116840190848210908211176104675760405282948184528183011161037d578281602093846000960137010152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000929173ffffffffffffffffffffffffffffffffffffffff90848335838116908190036101bd578152806020528260408220541693826020527b19457468657265756d205369676e6564204d6573736167653a0a33328252603c600420918461051961014084019461051361050c87876103a5565b36916103f6565b90610549565b168614610540575061050c6105329392610513926103a5565b160361053a57565b60019150565b96505050505050565b6001608060006041602094969596604080519880519285526060810151851a88528781015182520151606052145afa51913d1561058a576000606052604052565b638baa579f6000526004601cfd5b90916000923384528360205273ffffffffffffffffffffffffffffffffffffffff918260408620541693836105d76105d13685876103f6565b83610549565b1685146106245761061592610513916020527b19457468657265756d205369676e6564204d6573736167653a0a33328752603c6004209236916103f6565b160361061e5790565b50600190565b50505050509056fea164736f6c6343000812000a",
}

// KernelECDSAValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use KernelECDSAValidatorMetaData.ABI instead.
var KernelECDSAValidatorABI = KernelECDSAValidatorMetaData.ABI

// KernelECDSAValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KernelECDSAValidatorMetaData.Bin instead.
var KernelECDSAValidatorBin = KernelECDSAValidatorMetaData.Bin

// DeployKernelECDSAValidator deploys a new Ethereum contract, binding an instance of KernelECDSAValidator to it.
func DeployKernelECDSAValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KernelECDSAValidator, error) {
	parsed, err := KernelECDSAValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KernelECDSAValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KernelECDSAValidator{KernelECDSAValidatorCaller: KernelECDSAValidatorCaller{contract: contract}, KernelECDSAValidatorTransactor: KernelECDSAValidatorTransactor{contract: contract}, KernelECDSAValidatorFilterer: KernelECDSAValidatorFilterer{contract: contract}}, nil
}

// KernelECDSAValidator is an auto generated Go binding around an Ethereum contract.
type KernelECDSAValidator struct {
	KernelECDSAValidatorCaller     // Read-only binding to the contract
	KernelECDSAValidatorTransactor // Write-only binding to the contract
	KernelECDSAValidatorFilterer   // Log filterer for contract events
}

// KernelECDSAValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type KernelECDSAValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelECDSAValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KernelECDSAValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelECDSAValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KernelECDSAValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelECDSAValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KernelECDSAValidatorSession struct {
	Contract     *KernelECDSAValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KernelECDSAValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KernelECDSAValidatorCallerSession struct {
	Contract *KernelECDSAValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// KernelECDSAValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KernelECDSAValidatorTransactorSession struct {
	Contract     *KernelECDSAValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// KernelECDSAValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type KernelECDSAValidatorRaw struct {
	Contract *KernelECDSAValidator // Generic contract binding to access the raw methods on
}

// KernelECDSAValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KernelECDSAValidatorCallerRaw struct {
	Contract *KernelECDSAValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// KernelECDSAValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KernelECDSAValidatorTransactorRaw struct {
	Contract *KernelECDSAValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKernelECDSAValidator creates a new instance of KernelECDSAValidator, bound to a specific deployed contract.
func NewKernelECDSAValidator(address common.Address, backend bind.ContractBackend) (*KernelECDSAValidator, error) {
	contract, err := bindKernelECDSAValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KernelECDSAValidator{KernelECDSAValidatorCaller: KernelECDSAValidatorCaller{contract: contract}, KernelECDSAValidatorTransactor: KernelECDSAValidatorTransactor{contract: contract}, KernelECDSAValidatorFilterer: KernelECDSAValidatorFilterer{contract: contract}}, nil
}

// NewKernelECDSAValidatorCaller creates a new read-only instance of KernelECDSAValidator, bound to a specific deployed contract.
func NewKernelECDSAValidatorCaller(address common.Address, caller bind.ContractCaller) (*KernelECDSAValidatorCaller, error) {
	contract, err := bindKernelECDSAValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KernelECDSAValidatorCaller{contract: contract}, nil
}

// NewKernelECDSAValidatorTransactor creates a new write-only instance of KernelECDSAValidator, bound to a specific deployed contract.
func NewKernelECDSAValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*KernelECDSAValidatorTransactor, error) {
	contract, err := bindKernelECDSAValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KernelECDSAValidatorTransactor{contract: contract}, nil
}

// NewKernelECDSAValidatorFilterer creates a new log filterer instance of KernelECDSAValidator, bound to a specific deployed contract.
func NewKernelECDSAValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*KernelECDSAValidatorFilterer, error) {
	contract, err := bindKernelECDSAValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KernelECDSAValidatorFilterer{contract: contract}, nil
}

// bindKernelECDSAValidator binds a generic wrapper to an already deployed contract.
func bindKernelECDSAValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KernelECDSAValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KernelECDSAValidator *KernelECDSAValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KernelECDSAValidator.Contract.KernelECDSAValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KernelECDSAValidator *KernelECDSAValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.KernelECDSAValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KernelECDSAValidator *KernelECDSAValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.KernelECDSAValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KernelECDSAValidator *KernelECDSAValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KernelECDSAValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KernelECDSAValidator *KernelECDSAValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KernelECDSAValidator *KernelECDSAValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.contract.Transact(opts, method, params...)
}

// EcdsaValidatorStorage is a free data retrieval call binding the contract method 0x20709efc.
//
// Solidity: function ecdsaValidatorStorage(address ) view returns(address owner)
func (_KernelECDSAValidator *KernelECDSAValidatorCaller) EcdsaValidatorStorage(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _KernelECDSAValidator.contract.Call(opts, &out, "ecdsaValidatorStorage", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaValidatorStorage is a free data retrieval call binding the contract method 0x20709efc.
//
// Solidity: function ecdsaValidatorStorage(address ) view returns(address owner)
func (_KernelECDSAValidator *KernelECDSAValidatorSession) EcdsaValidatorStorage(arg0 common.Address) (common.Address, error) {
	return _KernelECDSAValidator.Contract.EcdsaValidatorStorage(&_KernelECDSAValidator.CallOpts, arg0)
}

// EcdsaValidatorStorage is a free data retrieval call binding the contract method 0x20709efc.
//
// Solidity: function ecdsaValidatorStorage(address ) view returns(address owner)
func (_KernelECDSAValidator *KernelECDSAValidatorCallerSession) EcdsaValidatorStorage(arg0 common.Address) (common.Address, error) {
	return _KernelECDSAValidator.Contract.EcdsaValidatorStorage(&_KernelECDSAValidator.CallOpts, arg0)
}

// ValidCaller is a free data retrieval call binding the contract method 0x9ea9bd59.
//
// Solidity: function validCaller(address _caller, bytes ) view returns(bool)
func (_KernelECDSAValidator *KernelECDSAValidatorCaller) ValidCaller(opts *bind.CallOpts, _caller common.Address, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _KernelECDSAValidator.contract.Call(opts, &out, "validCaller", _caller, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidCaller is a free data retrieval call binding the contract method 0x9ea9bd59.
//
// Solidity: function validCaller(address _caller, bytes ) view returns(bool)
func (_KernelECDSAValidator *KernelECDSAValidatorSession) ValidCaller(_caller common.Address, arg1 []byte) (bool, error) {
	return _KernelECDSAValidator.Contract.ValidCaller(&_KernelECDSAValidator.CallOpts, _caller, arg1)
}

// ValidCaller is a free data retrieval call binding the contract method 0x9ea9bd59.
//
// Solidity: function validCaller(address _caller, bytes ) view returns(bool)
func (_KernelECDSAValidator *KernelECDSAValidatorCallerSession) ValidCaller(_caller common.Address, arg1 []byte) (bool, error) {
	return _KernelECDSAValidator.Contract.ValidCaller(&_KernelECDSAValidator.CallOpts, _caller, arg1)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 hash, bytes signature) view returns(uint256)
func (_KernelECDSAValidator *KernelECDSAValidatorCaller) ValidateSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) (*big.Int, error) {
	var out []interface{}
	err := _KernelECDSAValidator.contract.Call(opts, &out, "validateSignature", hash, signature)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 hash, bytes signature) view returns(uint256)
func (_KernelECDSAValidator *KernelECDSAValidatorSession) ValidateSignature(hash [32]byte, signature []byte) (*big.Int, error) {
	return _KernelECDSAValidator.Contract.ValidateSignature(&_KernelECDSAValidator.CallOpts, hash, signature)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 hash, bytes signature) view returns(uint256)
func (_KernelECDSAValidator *KernelECDSAValidatorCallerSession) ValidateSignature(hash [32]byte, signature []byte) (*big.Int, error) {
	return _KernelECDSAValidator.Contract.ValidateSignature(&_KernelECDSAValidator.CallOpts, hash, signature)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes ) payable returns()
func (_KernelECDSAValidator *KernelECDSAValidatorTransactor) Disable(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _KernelECDSAValidator.contract.Transact(opts, "disable", arg0)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes ) payable returns()
func (_KernelECDSAValidator *KernelECDSAValidatorSession) Disable(arg0 []byte) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.Disable(&_KernelECDSAValidator.TransactOpts, arg0)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes ) payable returns()
func (_KernelECDSAValidator *KernelECDSAValidatorTransactorSession) Disable(arg0 []byte) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.Disable(&_KernelECDSAValidator.TransactOpts, arg0)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) payable returns()
func (_KernelECDSAValidator *KernelECDSAValidatorTransactor) Enable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _KernelECDSAValidator.contract.Transact(opts, "enable", _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) payable returns()
func (_KernelECDSAValidator *KernelECDSAValidatorSession) Enable(_data []byte) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.Enable(&_KernelECDSAValidator.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) payable returns()
func (_KernelECDSAValidator *KernelECDSAValidatorTransactorSession) Enable(_data []byte) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.Enable(&_KernelECDSAValidator.TransactOpts, _data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) _userOp, bytes32 _userOpHash, uint256 ) payable returns(uint256 validationData)
func (_KernelECDSAValidator *KernelECDSAValidatorTransactor) ValidateUserOp(opts *bind.TransactOpts, _userOp UserOperation, _userOpHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _KernelECDSAValidator.contract.Transact(opts, "validateUserOp", _userOp, _userOpHash, arg2)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) _userOp, bytes32 _userOpHash, uint256 ) payable returns(uint256 validationData)
func (_KernelECDSAValidator *KernelECDSAValidatorSession) ValidateUserOp(_userOp UserOperation, _userOpHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.ValidateUserOp(&_KernelECDSAValidator.TransactOpts, _userOp, _userOpHash, arg2)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) _userOp, bytes32 _userOpHash, uint256 ) payable returns(uint256 validationData)
func (_KernelECDSAValidator *KernelECDSAValidatorTransactorSession) ValidateUserOp(_userOp UserOperation, _userOpHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _KernelECDSAValidator.Contract.ValidateUserOp(&_KernelECDSAValidator.TransactOpts, _userOp, _userOpHash, arg2)
}

// KernelECDSAValidatorOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the KernelECDSAValidator contract.
type KernelECDSAValidatorOwnerChangedIterator struct {
	Event *KernelECDSAValidatorOwnerChanged // Event containing the contract specifics and raw log

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
func (it *KernelECDSAValidatorOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelECDSAValidatorOwnerChanged)
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
		it.Event = new(KernelECDSAValidatorOwnerChanged)
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
func (it *KernelECDSAValidatorOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelECDSAValidatorOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelECDSAValidatorOwnerChanged represents a OwnerChanged event raised by the KernelECDSAValidator contract.
type KernelECDSAValidatorOwnerChanged struct {
	Kernel   common.Address
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0x381c0d11398486654573703c51ee8210ce9461764d133f9f0e53b6a539705331.
//
// Solidity: event OwnerChanged(address indexed kernel, address indexed oldOwner, address indexed newOwner)
func (_KernelECDSAValidator *KernelECDSAValidatorFilterer) FilterOwnerChanged(opts *bind.FilterOpts, kernel []common.Address, oldOwner []common.Address, newOwner []common.Address) (*KernelECDSAValidatorOwnerChangedIterator, error) {

	var kernelRule []interface{}
	for _, kernelItem := range kernel {
		kernelRule = append(kernelRule, kernelItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KernelECDSAValidator.contract.FilterLogs(opts, "OwnerChanged", kernelRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KernelECDSAValidatorOwnerChangedIterator{contract: _KernelECDSAValidator.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0x381c0d11398486654573703c51ee8210ce9461764d133f9f0e53b6a539705331.
//
// Solidity: event OwnerChanged(address indexed kernel, address indexed oldOwner, address indexed newOwner)
func (_KernelECDSAValidator *KernelECDSAValidatorFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *KernelECDSAValidatorOwnerChanged, kernel []common.Address, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var kernelRule []interface{}
	for _, kernelItem := range kernel {
		kernelRule = append(kernelRule, kernelItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KernelECDSAValidator.contract.WatchLogs(opts, "OwnerChanged", kernelRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelECDSAValidatorOwnerChanged)
				if err := _KernelECDSAValidator.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0x381c0d11398486654573703c51ee8210ce9461764d133f9f0e53b6a539705331.
//
// Solidity: event OwnerChanged(address indexed kernel, address indexed oldOwner, address indexed newOwner)
func (_KernelECDSAValidator *KernelECDSAValidatorFilterer) ParseOwnerChanged(log types.Log) (*KernelECDSAValidatorOwnerChanged, error) {
	event := new(KernelECDSAValidatorOwnerChanged)
	if err := _KernelECDSAValidator.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
