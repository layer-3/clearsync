// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package escrow_app

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

// EscrowAppMetaData contains all meta data concerning the EscrowApp contract.
var EscrowAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e3f806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80639936d81214610030575b600080fd5b61004361003e366004610743565b61005a565b60405161005192919061080a565b60405180910390f35b60006060816100698780610863565b915050600260ff8216146100ab575050604080518082019091526012815271372830b93a34b1b4b830b73a3990109e901960711b6020820152600091506103bc565b600085900361028d578060ff166100c585602001356103c5565b60ff16146101125760405162461bcd60e51b8152602060048201526015602482015274021756e616e696d6f75733b207c70726f6f667c3d3605c1b60448201526064015b60405180910390fd5b61011c84806108b4565b61012d9060608101906040016108ef565b65ffffffffffff16600003610156575050604080516020810190915260008152600191506103bc565b61016084806108b4565b6101719060608101906040016108ef565b65ffffffffffff1660010361019a575050604080516020810190915260008152600191506103bc565b6101a484806108b4565b6101b59060608101906040016108ef565b65ffffffffffff16600303610245576101ce84806108b4565b6101df906080810190606001610921565b61022b5760405162461bcd60e51b815260206004820152601f60248201527f2166696e616c3b207475726e4e756d3e3d33202626207c70726f6f667c3d30006044820152606401610109565b5050604080516020810190915260008152600191506103bc565b60405162461bcd60e51b815260206004820181905260248201527f6261642063616e646964617465207475726e4e756d3b207c70726f6f667c3d306044820152606401610109565b600185900361038e576102d0868660008181106102ac576102ac61093c565b90506020028101906102be9190610952565b6102c790610cd8565b8260ff166103f6565b61031a868660008181106102e6576102e661093c565b90506020028101906102f89190610952565b61030290806108b4565b61030c9080610863565b61031591610db4565b6104a6565b61032761030285806108b4565b61022b8686600081811061033d5761033d61093c565b905060200281019061034f9190610952565b61035990806108b4565b6103639080610863565b61036c91610db4565b61037686806108b4565b6103809080610863565b61038991610db4565b610523565b505060408051808201909152601081526f0c4c2c840e0e4dedecc40d8cadccee8d60831b6020820152600091505b94509492505050565b6000805b82156103f0576103da600184610dd7565b90921691806103e881610dea565b9150506103c9565b92915050565b81600001516040015165ffffffffffff1660011461044e5760405162461bcd60e51b8152602060048201526015602482015274706f737466756e642e7475726e4e756d20213d203160581b6044820152606401610109565b8061045c83602001516103c5565b60ff16146104a25760405162461bcd60e51b8152602060048201526013602482015272706f737466756e642021756e616e696d6f757360681b6044820152606401610109565b5050565b60005b81518110156104a25760008282815181106104c6576104c661093c565b6020026020010151905080604001515160011461051a5760405162461bcd60e51b81526020600482015260126024820152717c616c6c6f636174696f6e737c20213d203160701b6044820152606401610109565b506001016104a9565b80518251146105745760405162461bcd60e51b815260206004820152601860248201527f7c6f7574636f6d65417c20213d207c6f7574636f6d65427c00000000000000006044820152606401610109565b60005b8251811015610726578181815181106105925761059261093c565b6020026020010151600001516001600160a01b03168382815181106105b9576105b961093c565b6020026020010151600001516001600160a01b03161461060c5760405162461bcd60e51b815260206004820152600e60248201526d0c2e6e6cae840dad2e6dac2e8c6d60931b6044820152606401610109565b60008382815181106106205761062061093c565b60200260200101516040015160008151811061063e5761063e61093c565b60200260200101519050600083838151811061065c5761065c61093c565b60200260200101516040015160008151811061067a5761067a61093c565b6020026020010151905080600001518260000151036106d35760405162461bcd60e51b8152602060048201526015602482015274064657374696e6174696f6e206d757374207377617605c1b6044820152606401610109565b806020015182602001511461071c5760405162461bcd60e51b815260206004820152600f60248201526e0c2dadeeadce840dad2e6dac2e8c6d608b1b6044820152606401610109565b5050600101610577565b505050565b60006040828403121561073d57600080fd5b50919050565b6000806000806060858703121561075957600080fd5b843567ffffffffffffffff8082111561077157600080fd5b908601906080828903121561078557600080fd5b9094506020860135908082111561079b57600080fd5b818701915087601f8301126107af57600080fd5b8135818111156107be57600080fd5b8860208260051b85010111156107d357600080fd5b6020830195508094505060408701359150808211156107f157600080fd5b506107fe8782880161072b565b91505092959194509250565b82151581526000602060406020840152835180604085015260005b8181101561084157858101830151858201606001528201610825565b506000606082860101526060601f19601f830116850101925050509392505050565b6000808335601e1984360301811261087a57600080fd5b83018035915067ffffffffffffffff82111561089557600080fd5b6020019150600581901b36038213156108ad57600080fd5b9250929050565b60008235607e198336030181126108ca57600080fd5b9190910192915050565b803565ffffffffffff811681146108ea57600080fd5b919050565b60006020828403121561090157600080fd5b61090a826108d4565b9392505050565b803580151581146108ea57600080fd5b60006020828403121561093357600080fd5b61090a82610911565b634e487b7160e01b600052603260045260246000fd5b60008235603e198336030181126108ca57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff811182821017156109a1576109a1610968565b60405290565b6040516060810167ffffffffffffffff811182821017156109a1576109a1610968565b6040805190810167ffffffffffffffff811182821017156109a1576109a1610968565b604051601f8201601f1916810167ffffffffffffffff81118282101715610a1657610a16610968565b604052919050565b600067ffffffffffffffff821115610a3857610a38610968565b5060051b60200190565b600082601f830112610a5357600080fd5b813567ffffffffffffffff811115610a6d57610a6d610968565b610a80601f8201601f19166020016109ed565b818152846020838601011115610a9557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610ac357600080fd5b81356020610ad8610ad383610a1e565b6109ed565b82815260059290921b84018101918181019086841115610af757600080fd5b8286015b84811015610ba057803567ffffffffffffffff80821115610b1c5760008081fd5b908801906080828b03601f1901811315610b365760008081fd5b610b3e61097e565b8784013581526040808501358983015260608086013560ff81168114610b645760008081fd5b83830152928501359284841115610b7d57600091508182fd5b610b8b8e8b86890101610a42565b90830152508652505050918301918301610afb565b509695505050505050565b6000610bb9610ad384610a1e565b8381529050602080820190600585901b840186811115610bd857600080fd5b845b81811015610ccd57803567ffffffffffffffff80821115610bfb5760008081fd5b908701906060828b031215610c105760008081fd5b610c186109a7565b82356001600160a01b0381168114610c305760008081fd5b81528286013582811115610c445760008081fd5b83016040818d03811315610c585760008081fd5b610c606109ca565b823560048110610c705760008081fd5b81528289013585811115610c845760008081fd5b610c908f828601610a42565b828b0152508389015284810135915083821115610cad5760008081fd5b610cb98d838701610ab2565b908301525086525050928201928201610bda565b505050509392505050565b600060408236031215610cea57600080fd5b610cf26109ca565b823567ffffffffffffffff80821115610d0a57600080fd5b818501915060808236031215610d1f57600080fd5b610d2761097e565b823582811115610d3657600080fd5b830136601f820112610d4757600080fd5b610d5636823560208401610bab565b825250602083013582811115610d6b57600080fd5b610d7736828601610a42565b602083015250610d89604084016108d4565b6040820152610d9a60608401610911565b606082015283525050602092830135928101929092525090565b600061090a368484610bab565b634e487b7160e01b600052601160045260246000fd5b818103818111156103f0576103f0610dc1565b600060ff821660ff8103610e0057610e00610dc1565b6001019291505056fea2646970667358221220097b92e4405f07ad78058fcb5e9a22d359224a0b3b648278e4e17404fd646d4864736f6c63430008160033",
}

// EscrowAppABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowAppMetaData.ABI instead.
var EscrowAppABI = EscrowAppMetaData.ABI

// EscrowAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EscrowAppMetaData.Bin instead.
var EscrowAppBin = EscrowAppMetaData.Bin

// DeployEscrowApp deploys a new Ethereum contract, binding an instance of EscrowApp to it.
func DeployEscrowApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EscrowApp, error) {
	parsed, err := EscrowAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EscrowAppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EscrowApp{EscrowAppCaller: EscrowAppCaller{contract: contract}, EscrowAppTransactor: EscrowAppTransactor{contract: contract}, EscrowAppFilterer: EscrowAppFilterer{contract: contract}}, nil
}

// EscrowApp is an auto generated Go binding around an Ethereum contract.
type EscrowApp struct {
	EscrowAppCaller     // Read-only binding to the contract
	EscrowAppTransactor // Write-only binding to the contract
	EscrowAppFilterer   // Log filterer for contract events
}

// EscrowAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowAppSession struct {
	Contract     *EscrowApp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowAppCallerSession struct {
	Contract *EscrowAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EscrowAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowAppTransactorSession struct {
	Contract     *EscrowAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EscrowAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowAppRaw struct {
	Contract *EscrowApp // Generic contract binding to access the raw methods on
}

// EscrowAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowAppCallerRaw struct {
	Contract *EscrowAppCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowAppTransactorRaw struct {
	Contract *EscrowAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrowApp creates a new instance of EscrowApp, bound to a specific deployed contract.
func NewEscrowApp(address common.Address, backend bind.ContractBackend) (*EscrowApp, error) {
	contract, err := bindEscrowApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EscrowApp{EscrowAppCaller: EscrowAppCaller{contract: contract}, EscrowAppTransactor: EscrowAppTransactor{contract: contract}, EscrowAppFilterer: EscrowAppFilterer{contract: contract}}, nil
}

// NewEscrowAppCaller creates a new read-only instance of EscrowApp, bound to a specific deployed contract.
func NewEscrowAppCaller(address common.Address, caller bind.ContractCaller) (*EscrowAppCaller, error) {
	contract, err := bindEscrowApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowAppCaller{contract: contract}, nil
}

// NewEscrowAppTransactor creates a new write-only instance of EscrowApp, bound to a specific deployed contract.
func NewEscrowAppTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowAppTransactor, error) {
	contract, err := bindEscrowApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowAppTransactor{contract: contract}, nil
}

// NewEscrowAppFilterer creates a new log filterer instance of EscrowApp, bound to a specific deployed contract.
func NewEscrowAppFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowAppFilterer, error) {
	contract, err := bindEscrowApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowAppFilterer{contract: contract}, nil
}

// bindEscrowApp binds a generic wrapper to an already deployed contract.
func bindEscrowApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EscrowAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EscrowApp *EscrowAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EscrowApp.Contract.EscrowAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EscrowApp *EscrowAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EscrowApp.Contract.EscrowAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EscrowApp *EscrowAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EscrowApp.Contract.EscrowAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EscrowApp *EscrowAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EscrowApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EscrowApp *EscrowAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EscrowApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EscrowApp *EscrowAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EscrowApp.Contract.contract.Transact(opts, method, params...)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_EscrowApp *EscrowAppCaller) StateIsSupported(opts *bind.CallOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	var out []interface{}
	err := _EscrowApp.contract.Call(opts, &out, "stateIsSupported", fixedPart, proof, candidate)

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
func (_EscrowApp *EscrowAppSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _EscrowApp.Contract.StateIsSupported(&_EscrowApp.CallOpts, fixedPart, proof, candidate)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_EscrowApp *EscrowAppCallerSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _EscrowApp.Contract.StateIsSupported(&_EscrowApp.CallOpts, fixedPart, proof, candidate)
}
