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
	Bin: "0x608060405234801561001057600080fd5b50610e6c806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80639936d81214610030575b600080fd5b61004361003e366004610758565b61005a565b60405161005192919061081f565b60405180910390f35b60006060816100698780610877565b915050600260ff8216146100ab575050604080518082019091526012815271372830b93a34b1b4b830b73a3990109e901960711b6020820152600091506103bc565b600085900361028d578060ff166100c585602001356103c5565b60ff16146101125760405162461bcd60e51b8152602060048201526015602482015274021756e616e696d6f75733b207c70726f6f667c3d3605c1b60448201526064015b60405180910390fd5b61011c84806108c8565b61012d906060810190604001610903565b65ffffffffffff16600003610156575050604080516020810190915260008152600191506103bc565b61016084806108c8565b610171906060810190604001610903565b65ffffffffffff1660010361019a575050604080516020810190915260008152600191506103bc565b6101a484806108c8565b6101b5906060810190604001610903565b65ffffffffffff16600303610245576101ce84806108c8565b6101df906080810190606001610935565b61022b5760405162461bcd60e51b815260206004820152601f60248201527f2166696e616c3b207475726e4e756d3e3d33202626207c70726f6f667c3d30006044820152606401610109565b5050604080516020810190915260008152600191506103bc565b60405162461bcd60e51b815260206004820181905260248201527f6261642063616e646964617465207475726e4e756d3b207c70726f6f667c3d306044820152606401610109565b600185900361038e576102d0868660008181106102ac576102ac610950565b90506020028101906102be9190610966565b6102c790610cec565b8260ff166103f6565b61031a868660008181106102e6576102e6610950565b90506020028101906102f89190610966565b61030290806108c8565b61030c9080610877565b61031591610dc8565b6104a6565b61032761030285806108c8565b61022b8686600081811061033d5761033d610950565b905060200281019061034f9190610966565b61035990806108c8565b6103639080610877565b61036c91610dc8565b61037686806108c8565b6103809080610877565b61038991610dc8565b61052d565b505060408051808201909152601081526f0c4c2c840e0e4dedecc40d8cadccee8d60831b6020820152600091505b94509492505050565b6000805b82156103f0576103da600184610deb565b90921691806103e881610dfe565b9150506103c9565b92915050565b81600001516040015165ffffffffffff1660011461044e5760405162461bcd60e51b8152602060048201526015602482015274706f737466756e642e7475726e4e756d20213d203160581b6044820152606401610109565b8061045c83602001516103c5565b60ff16146104a25760405162461bcd60e51b8152602060048201526013602482015272706f737466756e642021756e616e696d6f757360681b6044820152606401610109565b5050565b60005b81518110156104a25760008282815181106104c6576104c6610950565b6020026020010151905080604001515160011461051a5760405162461bcd60e51b81526020600482015260126024820152717c616c6c6f636174696f6e737c20213d203160701b6044820152606401610109565b508061052581610e1d565b9150506104a9565b805182511461057e5760405162461bcd60e51b815260206004820152601860248201527f7c6f7574636f6d65417c20213d207c6f7574636f6d65427c00000000000000006044820152606401610109565b60005b825181101561073b5781818151811061059c5761059c610950565b6020026020010151600001516001600160a01b03168382815181106105c3576105c3610950565b6020026020010151600001516001600160a01b0316146106165760405162461bcd60e51b815260206004820152600e60248201526d0c2e6e6cae840dad2e6dac2e8c6d60931b6044820152606401610109565b600083828151811061062a5761062a610950565b60200260200101516040015160008151811061064857610648610950565b60200260200101519050600083838151811061066657610666610950565b60200260200101516040015160008151811061068457610684610950565b6020026020010151905080600001518260000151036106dd5760405162461bcd60e51b8152602060048201526015602482015274064657374696e6174696f6e206d757374207377617605c1b6044820152606401610109565b80602001518260200151146107265760405162461bcd60e51b815260206004820152600f60248201526e0c2dadeeadce840dad2e6dac2e8c6d608b1b6044820152606401610109565b5050808061073390610e1d565b915050610581565b505050565b60006040828403121561075257600080fd5b50919050565b6000806000806060858703121561076e57600080fd5b843567ffffffffffffffff8082111561078657600080fd5b908601906080828903121561079a57600080fd5b909450602086013590808211156107b057600080fd5b818701915087601f8301126107c457600080fd5b8135818111156107d357600080fd5b8860208260051b85010111156107e857600080fd5b60208301955080945050604087013591508082111561080657600080fd5b5061081387828801610740565b91505092959194509250565b821515815260006020604081840152835180604085015260005b8181101561085557858101830151858201606001528201610839565b506000606082860101526060601f19601f830116850101925050509392505050565b6000808335601e1984360301811261088e57600080fd5b83018035915067ffffffffffffffff8211156108a957600080fd5b6020019150600581901b36038213156108c157600080fd5b9250929050565b60008235607e198336030181126108de57600080fd5b9190910192915050565b803565ffffffffffff811681146108fe57600080fd5b919050565b60006020828403121561091557600080fd5b61091e826108e8565b9392505050565b803580151581146108fe57600080fd5b60006020828403121561094757600080fd5b61091e82610925565b634e487b7160e01b600052603260045260246000fd5b60008235603e198336030181126108de57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff811182821017156109b5576109b561097c565b60405290565b6040516060810167ffffffffffffffff811182821017156109b5576109b561097c565b6040805190810167ffffffffffffffff811182821017156109b5576109b561097c565b604051601f8201601f1916810167ffffffffffffffff81118282101715610a2a57610a2a61097c565b604052919050565b600067ffffffffffffffff821115610a4c57610a4c61097c565b5060051b60200190565b600082601f830112610a6757600080fd5b813567ffffffffffffffff811115610a8157610a8161097c565b610a94601f8201601f1916602001610a01565b818152846020838601011115610aa957600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610ad757600080fd5b81356020610aec610ae783610a32565b610a01565b82815260059290921b84018101918181019086841115610b0b57600080fd5b8286015b84811015610bb457803567ffffffffffffffff80821115610b305760008081fd5b908801906080828b03601f1901811315610b4a5760008081fd5b610b52610992565b8784013581526040808501358983015260608086013560ff81168114610b785760008081fd5b83830152928501359284841115610b9157600091508182fd5b610b9f8e8b86890101610a56565b90830152508652505050918301918301610b0f565b509695505050505050565b6000610bcd610ae784610a32565b8381529050602080820190600585901b840186811115610bec57600080fd5b845b81811015610ce157803567ffffffffffffffff80821115610c0f5760008081fd5b908701906060828b031215610c245760008081fd5b610c2c6109bb565b82356001600160a01b0381168114610c445760008081fd5b81528286013582811115610c585760008081fd5b83016040818d03811315610c6c5760008081fd5b610c746109de565b823560048110610c845760008081fd5b81528289013585811115610c985760008081fd5b610ca48f828601610a56565b828b0152508389015284810135915083821115610cc15760008081fd5b610ccd8d838701610ac6565b908301525086525050928201928201610bee565b505050509392505050565b600060408236031215610cfe57600080fd5b610d066109de565b823567ffffffffffffffff80821115610d1e57600080fd5b818501915060808236031215610d3357600080fd5b610d3b610992565b823582811115610d4a57600080fd5b830136601f820112610d5b57600080fd5b610d6a36823560208401610bbf565b825250602083013582811115610d7f57600080fd5b610d8b36828601610a56565b602083015250610d9d604084016108e8565b6040820152610dae60608401610925565b606082015283525050602092830135928101929092525090565b600061091e368484610bbf565b634e487b7160e01b600052601160045260246000fd5b818103818111156103f0576103f0610dd5565b600060ff821660ff8103610e1457610e14610dd5565b60010192915050565b600060018201610e2f57610e2f610dd5565b506001019056fea2646970667358221220f12d40194b8784569097dc7b732cec536e689c4655e8c96aab90b9638d15ae8e64736f6c63430008140033",
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
