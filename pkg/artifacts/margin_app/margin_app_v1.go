// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package margin_app

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

// ExitFormatAllocation is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatAllocation struct {
	Destination    [32]byte
	Amount         *big.Int
	AllocationType uint8
	Metadata       []byte
}

// ExitFormatAssetMetadata is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatAssetMetadata struct {
	AssetType uint8
	Metadata  []byte
}

// ExitFormatSingleAssetExit is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatSingleAssetExit struct {
	Asset         common.Address
	AssetMetadata ExitFormatAssetMetadata
	Allocations   []ExitFormatAllocation
}

// INitroTypesFixedPart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesFixedPart struct {
	Participants      []common.Address
	ChannelNonce      uint64
	AppDefinition     common.Address
	ChallengeDuration *big.Int
}

// INitroTypesRecoveredVariablePart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesRecoveredVariablePart struct {
	VariablePart INitroTypesVariablePart
	SignedBy     *big.Int
}

// INitroTypesVariablePart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesVariablePart struct {
	Outcome []ExitFormatSingleAssetExit
	AppData []byte
	TurnNum *big.Int
	IsFinal bool
}

// MarginAppV1MetaData contains all meta data concerning the MarginAppV1 contract.
var MarginAppV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e7d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80639936d81214610030575b600080fd5b61004361003e366004610854565b61005a565b60405161005192919061091b565b60405180910390f35b60006060816100698780610973565b915050600260ff8216146100c45760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792032207061727469636970616e747320737570706f7274656400000060448201526064015b60405180910390fd5b8060ff166100d585602001356104c0565b60ff16146101125760405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b60448201526064016100bb565b60008590036102965761012584806109c4565b6101369060608101906040016109e4565b65ffffffffffff1660000361015f575050604080516020810190915260008152600191506104b7565b61016984806109c4565b61017a9060608101906040016109e4565b65ffffffffffff166001036101a3575050604080516020810190915260008152600191506104b7565b60036101af85806109c4565b6101c09060608101906040016109e4565b65ffffffffffff161061024e576101d784806109c4565b6101e8906080810190606001610a13565b6102345760405162461bcd60e51b815260206004820152601f60248201527f2166696e616c3b207475726e4e756d3e3d33202626207c70726f6f667c3d300060448201526064016100bb565b5050604080516020810190915260008152600191506104b7565b60405162461bcd60e51b815260206004820181905260248201527f6261642063616e646964617465207475726e4e756d3b207c70726f6f667c3d3060448201526064016100bb565b600185900361047c5760026102ab85806109c4565b6102bc9060608101906040016109e4565b65ffffffffffff16101561030b5760405162461bcd60e51b81526020600482015260166024820152757475726e4e756d3c32202626207c70726f6f667c3d3160501b60448201526064016100bb565b8585600081811061031e5761031e610a35565b90506020028101906103309190610a4b565b61033a90806109c4565b61034b9060608101906040016109e4565b65ffffffffffff1660011461039a5760405162461bcd60e51b8152602060048201526015602482015274706f737466756e642e7475726e4e756d20213d203160581b60448201526064016100bb565b8060ff166103cf878760008181106103b4576103b4610a35565b90506020028101906103c69190610a4b565b602001356104c0565b60ff16146104155760405162461bcd60e51b8152602060048201526013602482015272706f737466756e642021756e616e696d6f757360681b60448201526064016100bb565b6102348686600081811061042b5761042b610a35565b905060200281019061043d9190610a4b565b61044790806109c4565b6104519080610973565b61045a91610ca4565b61046486806109c4565b61046e9080610973565b61047791610ca4565b6104f1565b60405162461bcd60e51b815260206004820152601060248201526f0c4c2c840e0e4dedecc40d8cadccee8d60831b60448201526064016100bb565b94509492505050565b6000805b82156104eb576104d5600184610de9565b90921691806104e381610dfc565b9150506104c4565b92915050565b81516001148015610503575080516001145b61054f5760405162461bcd60e51b815260206004820152601a60248201527f696e636f7272656374206e756d626572206f662061737365747300000000000060448201526064016100bb565b8160008151811061056257610562610a35565b60200260200101516040015151600214801561059d57508060008151811061058c5761058c610a35565b602002602001015160400151516002145b6105e95760405162461bcd60e51b815260206004820152601f60248201527f696e636f7272656374206e756d626572206f6620616c6c6f636174696f6e730060448201526064016100bb565b806000815181106105fc576105fc610a35565b60200260200101516040015160008151811061061a5761061a610a35565b6020026020010151600001518260008151811061063957610639610a35565b60200260200101516040015160008151811061065757610657610a35565b6020026020010151600001511480156106e757508060008151811061067e5761067e610a35565b60200260200101516040015160018151811061069c5761069c610a35565b602002602001015160000151826000815181106106bb576106bb610a35565b6020026020010151604001516001815181106106d9576106d9610a35565b602002602001015160000151145b6107335760405162461bcd60e51b815260206004820152601a60248201527f64657374696e6174696f6e732063616e6e6f74206368616e676500000000000060448201526064016100bb565b60008060005b60028110156107e6578460008151811061075557610755610a35565b602002602001015160400151818151811061077257610772610a35565b602002602001015160200151836107899190610e1b565b92508360008151811061079e5761079e610a35565b60200260200101516040015181815181106107bb576107bb610a35565b602002602001015160200151826107d29190610e1b565b9150806107de81610e2e565b915050610739565b508082146108365760405162461bcd60e51b815260206004820152601d60248201527f746f74616c20616c6c6f63617465642063616e6e6f74206368616e676500000060448201526064016100bb565b50505050565b60006040828403121561084e57600080fd5b50919050565b6000806000806060858703121561086a57600080fd5b843567ffffffffffffffff8082111561088257600080fd5b908601906080828903121561089657600080fd5b909450602086013590808211156108ac57600080fd5b818701915087601f8301126108c057600080fd5b8135818111156108cf57600080fd5b8860208260051b85010111156108e457600080fd5b60208301955080945050604087013591508082111561090257600080fd5b5061090f8782880161083c565b91505092959194509250565b821515815260006020604081840152835180604085015260005b8181101561095157858101830151858201606001528201610935565b506000606082860101526060601f19601f830116850101925050509392505050565b6000808335601e1984360301811261098a57600080fd5b83018035915067ffffffffffffffff8211156109a557600080fd5b6020019150600581901b36038213156109bd57600080fd5b9250929050565b60008235607e198336030181126109da57600080fd5b9190910192915050565b6000602082840312156109f657600080fd5b813565ffffffffffff81168114610a0c57600080fd5b9392505050565b600060208284031215610a2557600080fd5b81358015158114610a0c57600080fd5b634e487b7160e01b600052603260045260246000fd5b60008235603e198336030181126109da57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715610a9a57610a9a610a61565b60405290565b6040516060810167ffffffffffffffff81118282101715610a9a57610a9a610a61565b6040805190810167ffffffffffffffff81118282101715610a9a57610a9a610a61565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b0f57610b0f610a61565b604052919050565b600067ffffffffffffffff821115610b3157610b31610a61565b5060051b60200190565b600082601f830112610b4c57600080fd5b813567ffffffffffffffff811115610b6657610b66610a61565b610b79601f8201601f1916602001610ae6565b818152846020838601011115610b8e57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610bbc57600080fd5b81356020610bd1610bcc83610b17565b610ae6565b82815260059290921b84018101918181019086841115610bf057600080fd5b8286015b84811015610c9957803567ffffffffffffffff80821115610c155760008081fd5b908801906080828b03601f1901811315610c2f5760008081fd5b610c37610a77565b8784013581526040808501358983015260608086013560ff81168114610c5d5760008081fd5b83830152928501359284841115610c7657600091508182fd5b610c848e8b86890101610b3b565b90830152508652505050918301918301610bf4565b509695505050505050565b6000610cb2610bcc84610b17565b80848252602080830192508560051b850136811115610cd057600080fd5b855b81811015610dc757803567ffffffffffffffff80821115610cf35760008081fd5b818901915060608236031215610d095760008081fd5b610d11610aa0565b82356001600160a01b0381168114610d295760008081fd5b81528286013582811115610d3d5760008081fd5b8301604036829003811315610d525760008081fd5b610d5a610ac3565b823560048110610d6a5760008081fd5b81528289013585811115610d7e5760008081fd5b610d8a36828601610b3b565b828b0152508389015284810135915083821115610da75760008081fd5b610db336838701610bab565b908301525087525050938201938201610cd2565b50919695505050505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156104eb576104eb610dd3565b600060ff821660ff8103610e1257610e12610dd3565b60010192915050565b808201808211156104eb576104eb610dd3565b600060018201610e4057610e40610dd3565b506001019056fea2646970667358221220a8cf425b7eb2696fe3d1513563d557ccf33ac60c6fdeb87178edc4b487bec93564736f6c63430008140033",
}

// MarginAppV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use MarginAppV1MetaData.ABI instead.
var MarginAppV1ABI = MarginAppV1MetaData.ABI

// MarginAppV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarginAppV1MetaData.Bin instead.
var MarginAppV1Bin = MarginAppV1MetaData.Bin

// DeployMarginAppV1 deploys a new Ethereum contract, binding an instance of MarginAppV1 to it.
func DeployMarginAppV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarginAppV1, error) {
	parsed, err := MarginAppV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarginAppV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarginAppV1{MarginAppV1Caller: MarginAppV1Caller{contract: contract}, MarginAppV1Transactor: MarginAppV1Transactor{contract: contract}, MarginAppV1Filterer: MarginAppV1Filterer{contract: contract}}, nil
}

// MarginAppV1 is an auto generated Go binding around an Ethereum contract.
type MarginAppV1 struct {
	MarginAppV1Caller     // Read-only binding to the contract
	MarginAppV1Transactor // Write-only binding to the contract
	MarginAppV1Filterer   // Log filterer for contract events
}

// MarginAppV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MarginAppV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginAppV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MarginAppV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginAppV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarginAppV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarginAppV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarginAppV1Session struct {
	Contract     *MarginAppV1      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarginAppV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarginAppV1CallerSession struct {
	Contract *MarginAppV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarginAppV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarginAppV1TransactorSession struct {
	Contract     *MarginAppV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarginAppV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MarginAppV1Raw struct {
	Contract *MarginAppV1 // Generic contract binding to access the raw methods on
}

// MarginAppV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarginAppV1CallerRaw struct {
	Contract *MarginAppV1Caller // Generic read-only contract binding to access the raw methods on
}

// MarginAppV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarginAppV1TransactorRaw struct {
	Contract *MarginAppV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMarginAppV1 creates a new instance of MarginAppV1, bound to a specific deployed contract.
func NewMarginAppV1(address common.Address, backend bind.ContractBackend) (*MarginAppV1, error) {
	contract, err := bindMarginAppV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarginAppV1{MarginAppV1Caller: MarginAppV1Caller{contract: contract}, MarginAppV1Transactor: MarginAppV1Transactor{contract: contract}, MarginAppV1Filterer: MarginAppV1Filterer{contract: contract}}, nil
}

// NewMarginAppV1Caller creates a new read-only instance of MarginAppV1, bound to a specific deployed contract.
func NewMarginAppV1Caller(address common.Address, caller bind.ContractCaller) (*MarginAppV1Caller, error) {
	contract, err := bindMarginAppV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarginAppV1Caller{contract: contract}, nil
}

// NewMarginAppV1Transactor creates a new write-only instance of MarginAppV1, bound to a specific deployed contract.
func NewMarginAppV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MarginAppV1Transactor, error) {
	contract, err := bindMarginAppV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarginAppV1Transactor{contract: contract}, nil
}

// NewMarginAppV1Filterer creates a new log filterer instance of MarginAppV1, bound to a specific deployed contract.
func NewMarginAppV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MarginAppV1Filterer, error) {
	contract, err := bindMarginAppV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarginAppV1Filterer{contract: contract}, nil
}

// bindMarginAppV1 binds a generic wrapper to an already deployed contract.
func bindMarginAppV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarginAppV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarginAppV1 *MarginAppV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarginAppV1.Contract.MarginAppV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarginAppV1 *MarginAppV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarginAppV1.Contract.MarginAppV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarginAppV1 *MarginAppV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarginAppV1.Contract.MarginAppV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarginAppV1 *MarginAppV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarginAppV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarginAppV1 *MarginAppV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarginAppV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarginAppV1 *MarginAppV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarginAppV1.Contract.contract.Transact(opts, method, params...)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_MarginAppV1 *MarginAppV1Caller) StateIsSupported(opts *bind.CallOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	var out []interface{}
	err := _MarginAppV1.contract.Call(opts, &out, "stateIsSupported", fixedPart, proof, candidate)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_MarginAppV1 *MarginAppV1Session) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _MarginAppV1.Contract.StateIsSupported(&_MarginAppV1.CallOpts, fixedPart, proof, candidate)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_MarginAppV1 *MarginAppV1CallerSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _MarginAppV1.Contract.StateIsSupported(&_MarginAppV1.CallOpts, fixedPart, proof, candidate)
}
