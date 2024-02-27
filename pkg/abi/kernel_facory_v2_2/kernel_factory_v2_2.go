// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kernel_factory_v2_2

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

// KernelFactoryMetaData contains all meta data concerning the KernelFactory contract.
var KernelFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DeploymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaltDoesNotStartWithCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Deployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllowedImplementation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"predictDeterministicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"predicted\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"name\":\"setEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346100b257601f610f5538819003918201601f19168301916001600160401b038311848410176100b75780849260409485528339810103126100b25780516001600160a01b0391828216918290036100b257602001519182168092036100b25780638b78c6d8195560007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a3600080546001600160a01b031916919091179055604051610e8790816100ce8239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6040608081526004908136101561001557600080fd5b600091823560e01c9081630396cb6014610b4a5781632569296214610ae1578163296601cd1461089c5781634d6cb700146107f75781635414dff01461079c57816354d1f13d14610738578163584465f2146106b65781636544c8281461064f578163715018a6146105d05781638da5cb5b1461055e578163b0d691fe1461050d578163bb30a9741461046b57838263bb9fe6bf146103cf578263c23a5cea1461030457508163db4c545e146102bc578163f04e283e146101f3578163f2fde38b1461013f575063fee81cf4146100eb57600080fd5b3461013b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b57602091610125610c00565b9063389a75e1600c525281600c20549051908152f35b5080fd5b839060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b57610173610c00565b9061017c610c56565b8160601b156101e8575073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff748739278181547f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08580a35580f35b637448fbae8352601cfd5b8360207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b957610226610c00565b61022e610c56565b63389a75e1600c528082526020600c2092835442116102ae57508173ffffffffffffffffffffffffffffffffffffffff929355167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff748739278181547f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08580a35580f35b636f5e88188352601cfd5b80fd5b50503461013b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b57602090608960136102fb610cbb565b01209051908152f35b809184346103cb5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103cb5781359173ffffffffffffffffffffffffffffffffffffffff8084168094036103c757610360610c56565b84541692833b156103c7576024859283855196879485937fc23a5cea0000000000000000000000000000000000000000000000000000000085528401525af19081156103be57506103ae5750f35b6103b790610df6565b6102b95780f35b513d84823e3d90fd5b8480fd5b5050fd5b809184346103cb57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103cb57610408610c56565b73ffffffffffffffffffffffffffffffffffffffff83541691823b1561046657839283918351809581937fbb9fe6bf0000000000000000000000000000000000000000000000000000000083525af19081156103be57506103ae5750f35b505050fd5b50503461013b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b576104a3610c00565b90602435918215158093036105095773ffffffffffffffffffffffffffffffffffffffff906104d0610c56565b168352600160205282209060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00835416911617905580f35b8380fd5b50503461013b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b5773ffffffffffffffffffffffffffffffffffffffff60209254169051908152f35b50503461013b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b5760209073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff7487392754915191168152f35b83807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b957610602610c56565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffff748739278181547f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a35580f35b50503461013b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b5760ff8160209373ffffffffffffffffffffffffffffffffffffffff6106a3610c00565b1681526001855220541690519015158152f35b83903461013b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261013b573573ffffffffffffffffffffffffffffffffffffffff811680910361013b5761070e610c56565b7fffffffffffffffffffffffff000000000000000000000000000000000000000082541617815580f35b83807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b95763389a75e1600c52338152806020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c928280a280f35b8284346102b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b9575073ffffffffffffffffffffffffffffffffffffffff6107ef60209335610c8e565b915191168152f35b8284346102b957817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b95782359067ffffffffffffffff82116102b957506107ef6bffffffffffffffffffffffff61086e60209573ffffffffffffffffffffffffffffffffffffffff94369101610c28565b6108938680518381948b830196873781016024358b8201520389810184520182610e39565b51902016610c8e565b828460607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b9576108d0610c00565b9160243567ffffffffffffffff8111610add576108f09036908601610c28565b9373ffffffffffffffffffffffffffffffffffffffff9384821681526020966001885260ff858320541615610a5b576bffffffffffffffffffffffff8551898101908987833761095488828d8d82019060443590820152038d810184520182610e39565b5190201696331560011715610a4f5761096b610cbb565b90601382019860898a2060ff86536035523060601b600152806015526055852099856035528a3b156109a3575b8b8b8b8b5191168152f35b856089929394959697989b50f5978815610a4357918185939284938884527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8d85015289840137870190348a5af115610a2957507f09e48df7857bd0c1e0d31bb8a85d42cf1874817895f171c917f6ee2cea73ec20818692a38480808080808080610998565b3d15610a3857503d81803e3d90fd5b63301164258252601cfd5b8363301164258652601cfd5b50632f6348368252601cfd5b608490888651917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602960248201527f4b65726e656c466163746f72793a20696d706c656d656e746174696f6e206e6f60448201527f7420616c6c6f77656400000000000000000000000000000000000000000000006064820152fd5b8280fd5b83807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b95763389a75e1600c523381526202a30042016020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d8280a280f35b91905060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610add5782823563ffffffff811680910361013b57610b90610c56565b73ffffffffffffffffffffffffffffffffffffffff82541693843b15610add57602490845195869384927f0396cb6000000000000000000000000000000000000000000000000000000000845283015234905af19081156103be5750610bf4575080f35b610bfd90610df6565b80f35b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610c2357565b600080fd5b9181601f84011215610c235782359167ffffffffffffffff8311610c235760208381860195010111610c2357565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff74873927543303610c8057565b6382b429006000526004601cfd5b60896013610c9a610cbb565b012060ff6000536035523060601b6001526015526055600020906000603552565b604051903060701c15610d5f57666052573d6000fd607b8301527f3d356020355560408036111560525736038060403d373d3d355af43d6000803e60748301527f3735a920a3ca505d382bbc545af43d6000803e6052573d6000fd5b3d6000f35b60548301527f14605757363d3d37363d7f360894a13ba1a3210667c828492db98dca3e2076cc60348301523060148301526c607f3d8160093d39f33d3d33738252565b66604c573d6000fd60758301527f3d3560203555604080361115604c5736038060403d373d3d355af43d6000803e606e8301527f3735a920a3ca505d382bbc545af43d6000803e604c573d6000fd5b3d6000f35b604e8301527f14605157363d3d37363d7f360894a13ba1a3210667c828492db98dca3e2076cc602e83015230600e8301526c60793d8160093d39f33d3d336d8252565b67ffffffffffffffff8111610e0a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610e0a5760405256fea164736f6c6343000812000a",
}

// KernelFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use KernelFactoryMetaData.ABI instead.
var KernelFactoryABI = KernelFactoryMetaData.ABI

// KernelFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KernelFactoryMetaData.Bin instead.
var KernelFactoryBin = KernelFactoryMetaData.Bin

// DeployKernelFactory deploys a new Ethereum contract, binding an instance of KernelFactory to it.
func DeployKernelFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _entryPoint common.Address) (common.Address, *types.Transaction, *KernelFactory, error) {
	parsed, err := KernelFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KernelFactoryBin), backend, _owner, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KernelFactory{KernelFactoryCaller: KernelFactoryCaller{contract: contract}, KernelFactoryTransactor: KernelFactoryTransactor{contract: contract}, KernelFactoryFilterer: KernelFactoryFilterer{contract: contract}}, nil
}

// KernelFactory is an auto generated Go binding around an Ethereum contract.
type KernelFactory struct {
	KernelFactoryCaller     // Read-only binding to the contract
	KernelFactoryTransactor // Write-only binding to the contract
	KernelFactoryFilterer   // Log filterer for contract events
}

// KernelFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KernelFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KernelFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KernelFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KernelFactorySession struct {
	Contract     *KernelFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KernelFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KernelFactoryCallerSession struct {
	Contract *KernelFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KernelFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KernelFactoryTransactorSession struct {
	Contract     *KernelFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KernelFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KernelFactoryRaw struct {
	Contract *KernelFactory // Generic contract binding to access the raw methods on
}

// KernelFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KernelFactoryCallerRaw struct {
	Contract *KernelFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// KernelFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KernelFactoryTransactorRaw struct {
	Contract *KernelFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKernelFactory creates a new instance of KernelFactory, bound to a specific deployed contract.
func NewKernelFactory(address common.Address, backend bind.ContractBackend) (*KernelFactory, error) {
	contract, err := bindKernelFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KernelFactory{KernelFactoryCaller: KernelFactoryCaller{contract: contract}, KernelFactoryTransactor: KernelFactoryTransactor{contract: contract}, KernelFactoryFilterer: KernelFactoryFilterer{contract: contract}}, nil
}

// NewKernelFactoryCaller creates a new read-only instance of KernelFactory, bound to a specific deployed contract.
func NewKernelFactoryCaller(address common.Address, caller bind.ContractCaller) (*KernelFactoryCaller, error) {
	contract, err := bindKernelFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryCaller{contract: contract}, nil
}

// NewKernelFactoryTransactor creates a new write-only instance of KernelFactory, bound to a specific deployed contract.
func NewKernelFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*KernelFactoryTransactor, error) {
	contract, err := bindKernelFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryTransactor{contract: contract}, nil
}

// NewKernelFactoryFilterer creates a new log filterer instance of KernelFactory, bound to a specific deployed contract.
func NewKernelFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*KernelFactoryFilterer, error) {
	contract, err := bindKernelFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryFilterer{contract: contract}, nil
}

// bindKernelFactory binds a generic wrapper to an already deployed contract.
func bindKernelFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KernelFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KernelFactory *KernelFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KernelFactory.Contract.KernelFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KernelFactory *KernelFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelFactory.Contract.KernelFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KernelFactory *KernelFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KernelFactory.Contract.KernelFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KernelFactory *KernelFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KernelFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KernelFactory *KernelFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KernelFactory *KernelFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KernelFactory.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_KernelFactory *KernelFactoryCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_KernelFactory *KernelFactorySession) EntryPoint() (common.Address, error) {
	return _KernelFactory.Contract.EntryPoint(&_KernelFactory.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_KernelFactory *KernelFactoryCallerSession) EntryPoint() (common.Address, error) {
	return _KernelFactory.Contract.EntryPoint(&_KernelFactory.CallOpts)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x4d6cb700.
//
// Solidity: function getAccountAddress(bytes _data, uint256 _index) view returns(address)
func (_KernelFactory *KernelFactoryCaller) GetAccountAddress(opts *bind.CallOpts, _data []byte, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "getAccountAddress", _data, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountAddress is a free data retrieval call binding the contract method 0x4d6cb700.
//
// Solidity: function getAccountAddress(bytes _data, uint256 _index) view returns(address)
func (_KernelFactory *KernelFactorySession) GetAccountAddress(_data []byte, _index *big.Int) (common.Address, error) {
	return _KernelFactory.Contract.GetAccountAddress(&_KernelFactory.CallOpts, _data, _index)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x4d6cb700.
//
// Solidity: function getAccountAddress(bytes _data, uint256 _index) view returns(address)
func (_KernelFactory *KernelFactoryCallerSession) GetAccountAddress(_data []byte, _index *big.Int) (common.Address, error) {
	return _KernelFactory.Contract.GetAccountAddress(&_KernelFactory.CallOpts, _data, _index)
}

// InitCodeHash is a free data retrieval call binding the contract method 0xdb4c545e.
//
// Solidity: function initCodeHash() view returns(bytes32 result)
func (_KernelFactory *KernelFactoryCaller) InitCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "initCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InitCodeHash is a free data retrieval call binding the contract method 0xdb4c545e.
//
// Solidity: function initCodeHash() view returns(bytes32 result)
func (_KernelFactory *KernelFactorySession) InitCodeHash() ([32]byte, error) {
	return _KernelFactory.Contract.InitCodeHash(&_KernelFactory.CallOpts)
}

// InitCodeHash is a free data retrieval call binding the contract method 0xdb4c545e.
//
// Solidity: function initCodeHash() view returns(bytes32 result)
func (_KernelFactory *KernelFactoryCallerSession) InitCodeHash() ([32]byte, error) {
	return _KernelFactory.Contract.InitCodeHash(&_KernelFactory.CallOpts)
}

// IsAllowedImplementation is a free data retrieval call binding the contract method 0x6544c828.
//
// Solidity: function isAllowedImplementation(address ) view returns(bool)
func (_KernelFactory *KernelFactoryCaller) IsAllowedImplementation(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "isAllowedImplementation", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedImplementation is a free data retrieval call binding the contract method 0x6544c828.
//
// Solidity: function isAllowedImplementation(address ) view returns(bool)
func (_KernelFactory *KernelFactorySession) IsAllowedImplementation(arg0 common.Address) (bool, error) {
	return _KernelFactory.Contract.IsAllowedImplementation(&_KernelFactory.CallOpts, arg0)
}

// IsAllowedImplementation is a free data retrieval call binding the contract method 0x6544c828.
//
// Solidity: function isAllowedImplementation(address ) view returns(bool)
func (_KernelFactory *KernelFactoryCallerSession) IsAllowedImplementation(arg0 common.Address) (bool, error) {
	return _KernelFactory.Contract.IsAllowedImplementation(&_KernelFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_KernelFactory *KernelFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_KernelFactory *KernelFactorySession) Owner() (common.Address, error) {
	return _KernelFactory.Contract.Owner(&_KernelFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_KernelFactory *KernelFactoryCallerSession) Owner() (common.Address, error) {
	return _KernelFactory.Contract.Owner(&_KernelFactory.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_KernelFactory *KernelFactoryCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_KernelFactory *KernelFactorySession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _KernelFactory.Contract.OwnershipHandoverExpiresAt(&_KernelFactory.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_KernelFactory *KernelFactoryCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _KernelFactory.Contract.OwnershipHandoverExpiresAt(&_KernelFactory.CallOpts, pendingOwner)
}

// PredictDeterministicAddress is a free data retrieval call binding the contract method 0x5414dff0.
//
// Solidity: function predictDeterministicAddress(bytes32 salt) view returns(address predicted)
func (_KernelFactory *KernelFactoryCaller) PredictDeterministicAddress(opts *bind.CallOpts, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _KernelFactory.contract.Call(opts, &out, "predictDeterministicAddress", salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictDeterministicAddress is a free data retrieval call binding the contract method 0x5414dff0.
//
// Solidity: function predictDeterministicAddress(bytes32 salt) view returns(address predicted)
func (_KernelFactory *KernelFactorySession) PredictDeterministicAddress(salt [32]byte) (common.Address, error) {
	return _KernelFactory.Contract.PredictDeterministicAddress(&_KernelFactory.CallOpts, salt)
}

// PredictDeterministicAddress is a free data retrieval call binding the contract method 0x5414dff0.
//
// Solidity: function predictDeterministicAddress(bytes32 salt) view returns(address predicted)
func (_KernelFactory *KernelFactoryCallerSession) PredictDeterministicAddress(salt [32]byte) (common.Address, error) {
	return _KernelFactory.Contract.PredictDeterministicAddress(&_KernelFactory.CallOpts, salt)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_KernelFactory *KernelFactoryTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_KernelFactory *KernelFactorySession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _KernelFactory.Contract.AddStake(&_KernelFactory.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_KernelFactory *KernelFactoryTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _KernelFactory.Contract.AddStake(&_KernelFactory.TransactOpts, unstakeDelaySec)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_KernelFactory *KernelFactoryTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_KernelFactory *KernelFactorySession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _KernelFactory.Contract.CancelOwnershipHandover(&_KernelFactory.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_KernelFactory *KernelFactoryTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _KernelFactory.Contract.CancelOwnershipHandover(&_KernelFactory.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_KernelFactory *KernelFactoryTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_KernelFactory *KernelFactorySession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.CompleteOwnershipHandover(&_KernelFactory.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_KernelFactory *KernelFactoryTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.CompleteOwnershipHandover(&_KernelFactory.TransactOpts, pendingOwner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _data, uint256 _index) payable returns(address proxy)
func (_KernelFactory *KernelFactoryTransactor) CreateAccount(opts *bind.TransactOpts, _implementation common.Address, _data []byte, _index *big.Int) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "createAccount", _implementation, _data, _index)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _data, uint256 _index) payable returns(address proxy)
func (_KernelFactory *KernelFactorySession) CreateAccount(_implementation common.Address, _data []byte, _index *big.Int) (*types.Transaction, error) {
	return _KernelFactory.Contract.CreateAccount(&_KernelFactory.TransactOpts, _implementation, _data, _index)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x296601cd.
//
// Solidity: function createAccount(address _implementation, bytes _data, uint256 _index) payable returns(address proxy)
func (_KernelFactory *KernelFactoryTransactorSession) CreateAccount(_implementation common.Address, _data []byte, _index *big.Int) (*types.Transaction, error) {
	return _KernelFactory.Contract.CreateAccount(&_KernelFactory.TransactOpts, _implementation, _data, _index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_KernelFactory *KernelFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_KernelFactory *KernelFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _KernelFactory.Contract.RenounceOwnership(&_KernelFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_KernelFactory *KernelFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KernelFactory.Contract.RenounceOwnership(&_KernelFactory.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_KernelFactory *KernelFactoryTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_KernelFactory *KernelFactorySession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _KernelFactory.Contract.RequestOwnershipHandover(&_KernelFactory.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_KernelFactory *KernelFactoryTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _KernelFactory.Contract.RequestOwnershipHandover(&_KernelFactory.TransactOpts)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_KernelFactory *KernelFactoryTransactor) SetEntryPoint(opts *bind.TransactOpts, _entryPoint common.Address) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "setEntryPoint", _entryPoint)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_KernelFactory *KernelFactorySession) SetEntryPoint(_entryPoint common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.SetEntryPoint(&_KernelFactory.TransactOpts, _entryPoint)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_KernelFactory *KernelFactoryTransactorSession) SetEntryPoint(_entryPoint common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.SetEntryPoint(&_KernelFactory.TransactOpts, _entryPoint)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xbb30a974.
//
// Solidity: function setImplementation(address _implementation, bool _allow) returns()
func (_KernelFactory *KernelFactoryTransactor) SetImplementation(opts *bind.TransactOpts, _implementation common.Address, _allow bool) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "setImplementation", _implementation, _allow)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xbb30a974.
//
// Solidity: function setImplementation(address _implementation, bool _allow) returns()
func (_KernelFactory *KernelFactorySession) SetImplementation(_implementation common.Address, _allow bool) (*types.Transaction, error) {
	return _KernelFactory.Contract.SetImplementation(&_KernelFactory.TransactOpts, _implementation, _allow)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xbb30a974.
//
// Solidity: function setImplementation(address _implementation, bool _allow) returns()
func (_KernelFactory *KernelFactoryTransactorSession) SetImplementation(_implementation common.Address, _allow bool) (*types.Transaction, error) {
	return _KernelFactory.Contract.SetImplementation(&_KernelFactory.TransactOpts, _implementation, _allow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_KernelFactory *KernelFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_KernelFactory *KernelFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.TransferOwnership(&_KernelFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_KernelFactory *KernelFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.TransferOwnership(&_KernelFactory.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_KernelFactory *KernelFactoryTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_KernelFactory *KernelFactorySession) UnlockStake() (*types.Transaction, error) {
	return _KernelFactory.Contract.UnlockStake(&_KernelFactory.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_KernelFactory *KernelFactoryTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _KernelFactory.Contract.UnlockStake(&_KernelFactory.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_KernelFactory *KernelFactoryTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _KernelFactory.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_KernelFactory *KernelFactorySession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.WithdrawStake(&_KernelFactory.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_KernelFactory *KernelFactoryTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _KernelFactory.Contract.WithdrawStake(&_KernelFactory.TransactOpts, withdrawAddress)
}

// KernelFactoryDeployedIterator is returned from FilterDeployed and is used to iterate over the raw logs and unpacked data for Deployed events raised by the KernelFactory contract.
type KernelFactoryDeployedIterator struct {
	Event *KernelFactoryDeployed // Event containing the contract specifics and raw log

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
func (it *KernelFactoryDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelFactoryDeployed)
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
		it.Event = new(KernelFactoryDeployed)
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
func (it *KernelFactoryDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelFactoryDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelFactoryDeployed represents a Deployed event raised by the KernelFactory contract.
type KernelFactoryDeployed struct {
	Proxy          common.Address
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeployed is a free log retrieval operation binding the contract event 0x09e48df7857bd0c1e0d31bb8a85d42cf1874817895f171c917f6ee2cea73ec20.
//
// Solidity: event Deployed(address indexed proxy, address indexed implementation)
func (_KernelFactory *KernelFactoryFilterer) FilterDeployed(opts *bind.FilterOpts, proxy []common.Address, implementation []common.Address) (*KernelFactoryDeployedIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _KernelFactory.contract.FilterLogs(opts, "Deployed", proxyRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryDeployedIterator{contract: _KernelFactory.contract, event: "Deployed", logs: logs, sub: sub}, nil
}

// WatchDeployed is a free log subscription operation binding the contract event 0x09e48df7857bd0c1e0d31bb8a85d42cf1874817895f171c917f6ee2cea73ec20.
//
// Solidity: event Deployed(address indexed proxy, address indexed implementation)
func (_KernelFactory *KernelFactoryFilterer) WatchDeployed(opts *bind.WatchOpts, sink chan<- *KernelFactoryDeployed, proxy []common.Address, implementation []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _KernelFactory.contract.WatchLogs(opts, "Deployed", proxyRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelFactoryDeployed)
				if err := _KernelFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
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

// ParseDeployed is a log parse operation binding the contract event 0x09e48df7857bd0c1e0d31bb8a85d42cf1874817895f171c917f6ee2cea73ec20.
//
// Solidity: event Deployed(address indexed proxy, address indexed implementation)
func (_KernelFactory *KernelFactoryFilterer) ParseDeployed(log types.Log) (*KernelFactoryDeployed, error) {
	event := new(KernelFactoryDeployed)
	if err := _KernelFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KernelFactoryOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the KernelFactory contract.
type KernelFactoryOwnershipHandoverCanceledIterator struct {
	Event *KernelFactoryOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *KernelFactoryOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelFactoryOwnershipHandoverCanceled)
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
		it.Event = new(KernelFactoryOwnershipHandoverCanceled)
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
func (it *KernelFactoryOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelFactoryOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelFactoryOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the KernelFactory contract.
type KernelFactoryOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_KernelFactory *KernelFactoryFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*KernelFactoryOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _KernelFactory.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryOwnershipHandoverCanceledIterator{contract: _KernelFactory.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_KernelFactory *KernelFactoryFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *KernelFactoryOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _KernelFactory.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelFactoryOwnershipHandoverCanceled)
				if err := _KernelFactory.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_KernelFactory *KernelFactoryFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*KernelFactoryOwnershipHandoverCanceled, error) {
	event := new(KernelFactoryOwnershipHandoverCanceled)
	if err := _KernelFactory.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KernelFactoryOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the KernelFactory contract.
type KernelFactoryOwnershipHandoverRequestedIterator struct {
	Event *KernelFactoryOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *KernelFactoryOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelFactoryOwnershipHandoverRequested)
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
		it.Event = new(KernelFactoryOwnershipHandoverRequested)
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
func (it *KernelFactoryOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelFactoryOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelFactoryOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the KernelFactory contract.
type KernelFactoryOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_KernelFactory *KernelFactoryFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*KernelFactoryOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _KernelFactory.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryOwnershipHandoverRequestedIterator{contract: _KernelFactory.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_KernelFactory *KernelFactoryFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *KernelFactoryOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _KernelFactory.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelFactoryOwnershipHandoverRequested)
				if err := _KernelFactory.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_KernelFactory *KernelFactoryFilterer) ParseOwnershipHandoverRequested(log types.Log) (*KernelFactoryOwnershipHandoverRequested, error) {
	event := new(KernelFactoryOwnershipHandoverRequested)
	if err := _KernelFactory.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KernelFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KernelFactory contract.
type KernelFactoryOwnershipTransferredIterator struct {
	Event *KernelFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KernelFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelFactoryOwnershipTransferred)
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
		it.Event = new(KernelFactoryOwnershipTransferred)
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
func (it *KernelFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the KernelFactory contract.
type KernelFactoryOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_KernelFactory *KernelFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*KernelFactoryOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KernelFactory.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KernelFactoryOwnershipTransferredIterator{contract: _KernelFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_KernelFactory *KernelFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KernelFactoryOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KernelFactory.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelFactoryOwnershipTransferred)
				if err := _KernelFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_KernelFactory *KernelFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*KernelFactoryOwnershipTransferred, error) {
	event := new(KernelFactoryOwnershipTransferred)
	if err := _KernelFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
