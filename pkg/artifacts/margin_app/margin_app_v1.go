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
	Bin: "0x60806040523461001b575b60405161105f610029823961105f90f35b610023600080fd5b61000a56fe60806040526004361015610018575b610016600080fd5b005b60003560e01c639936d8120361000e576100306101c8565b61000e565b600080fd5b908160809103126100485790565b610050600080fd5b90565b909182601f8301121561009b575b60208235926001600160401b03841161008e575b01926020830284011161008457565b61008c600080fd5b565b610096600080fd5b610075565b6100a3600080fd5b610061565b908160409103126100485790565b90606082820312610146575b6100de8183356001600160401b038111610139575b840161003a565b926100506101018360208601356001600160401b03811161012c575b8601610053565b9390946040810135906001600160401b03821161011f575b016100a8565b610127600080fd5b610119565b610134600080fd5b6100fa565b610141600080fd5b6100d7565b61014e600080fd5b6100c2565b60005b8381106101665750506000910152565b8181015183820152602001610156565b601f01601f191690565b6101a16101766020936101aa93610195815190565b80835293849260200190565b95869101610153565b0190565b901515815260406020820181905261005092910190610180565b5034610200575b6101e66101dd3660046100b6565b92919091610919565b906101fc6101f360405190565b928392836101ae565b0390f35b610208600080fd5b6101cf565b903590601e193682900301821215610252575b0160208135916001600160401b038311610245575b0191602082023603831361008457565b61024d600080fd5b610235565b61025a600080fd5b610220565b61026c6100506100509290565b60ff1690565b1561027957565b60405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792032207061727469636970616e747320737570706f727465640000006044820152606490fd5b0390fd5b805b0361003557565b35610050816102c2565b156102dc57565b60405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b6044820152606490fd5b6100506100506100509290565b903590607e193682900301821215610331570190565b6101aa600080fd5b65ffffffffffff81166102c4565b3561005081610339565b61035e6100506100509290565b65ffffffffffff1690565b50634e487b7160e01b600052604160045260246000fd5b90601f01601f191681019081106001600160401b038211176103a157604052565b6103a9610369565b604052565b9061008c6103bb60405190565b9283610380565b6101aa6020916001600160401b0381116103e057601f01601f191690565b610176610369565b906103fa6103f5836103c2565b6103ae565b918252565b61005060006103e8565b6100506103ff565b8015156102c4565b3561005081610411565b1561042a57565b60405162461bcd60e51b815260206004820152601f60248201527f2166696e616c3b207475726e4e756d3e3d33202626207c70726f6f667c3d30006044820152606490fd5b1561047657565b60405162461bcd60e51b81526020600482015260166024820152757475726e4e756d3c32202626207c70726f6f667c3d3160501b6044820152606490fd5b50634e487b7160e01b600052603260045260246000fd5b903590603e193682900301821215610331570190565b90610050926020918110156104fa575b028101906104cb565b6105026104b4565b6104f1565b1561050e57565b60405162461bcd60e51b8152602060048201526015602482015274706f737466756e642e7475726e4e756d20213d203160581b6044820152606490fd5b1561055257565b60405162461bcd60e51b8152602060048201526013602482015272706f737466756e642021756e616e696d6f757360681b6044820152606490fd5b602080916001600160401b0381116105a457020190565b6105ac610369565b020190565b6001600160a01b0381166102c4565b9050359061008c826105b1565b6004111561003557565b9050359061008c826105cd565b90826000939282370152565b92919061008c916106036103f5836103c2565b94828652602086019183820111156105e45761061d600080fd5b6105e4565b906100509181601f8201121561063e575b6020813591016105f0565b610646600080fd5b610633565b9190610692906040848203126106a6575b61066660406103ae565b93600061067383836105d7565b908601526020810135906001600160401b038211610699575b01610622565b6020830152565b6106a1600080fd5b61068c565b6106ae600080fd5b61065c565b9050359061008c826102c2565b60ff81166102c4565b9050359061008c826106c0565b919061074090608084820312610747575b6106f160806103ae565b9360006106fe83836106b3565b908601526020610710838284016106b3565b908601526040610722838284016106c9565b908601526060810135906001600160401b0382116106995701610622565b6060830152565b61074f600080fd5b6106e7565b909291926107646103f58261058d565b93818552602080860192028301928184116107c3575b80925b84841061078b575050505050565b602080916107ab8587356001600160401b0381116107b6575b86016106d6565b81520193019261077d565b6107be600080fd5b6107a4565b6107cb600080fd5b61077a565b906100509181601f820112156107ec575b602081359101610754565b6107f4600080fd5b6107e1565b919061086490606084820312610885575b61081460606103ae565b93600061082183836105c0565b908601526108448260208301356001600160401b038111610878575b830161064b565b60208601526040810135906001600160401b03821161086b575b016107d0565b6040830152565b610873600080fd5b61085e565b610880600080fd5b61083d565b61088d600080fd5b61080a565b909291926108a26103f58261058d565b9381855260208086019202830192818411610901575b80925b8484106108c9575050505050565b602080916108e98587356001600160401b0381116108f4575b86016107f9565b8152019301926108bb565b6108fc600080fd5b6108e2565b610909600080fd5b6108b8565b610050913691610892565b61093b610936610940929593949561092f600090565b508061020d565b905090565b61025f565b60029261095961094f8561025f565b60ff841614610272565b61097d61097061096b602086016102cb565b610c49565b60ff8481169116146102d5565b8460009261098a8461030e565b8714610b02576001966109a361099f8961030e565b9190565b146109e05760405162461bcd60e51b815260206004820152601060248201526f0c4c2c840e0e4dedecc40d8cadccee8d60831b6044820152606490fd5b838501956109ee878761031b565b6040016109fa90610347565b90610a0490610351565b610a179165ffffffffffff16101561046f565b610a208461030e565b610a2b9083856104e1565b848101610a379161031b565b604001610a4390610347565b610a4c88610351565b610a5e9165ffffffffffff1614610507565b610a678461030e565b610a729083856104e1565b602001610a7e906102cb565b610a8790610c49565b9060ff169060ff1614610a999061054b565b610aa28361030e565b90610aac926104e1565b818101610ab89161031b565b818101610ac49161020d565b93610acf919361031b565b908101610adb9161020d565b92610ae6919261090e565b91610af09161090e565b610af991610e0b565b90610050610409565b5050509150915080820190610b3a610b2e610b286040610b22868861031b565b01610347565b92610351565b9165ffffffffffff1690565b14610bf657610b4e6040610b22838561031b565b91600192610b5e610b2e85610351565b14610beb57610b726040610b22848461031b565b610b7f610b2e6003610351565b1015610bce5760405162461bcd60e51b8152806102be600482016020808252818101527f6261642063616e646964617465207475726e4e756d3b207c70726f6f667c3d30604082015260600190565b6060610be0610af993610be69361031b565b01610419565b610423565b505090610050610409565b5050600190610050610409565b50634e487b7160e01b600052601160045260246000fd5b60019060ff1660ff8114610c2c570190565b6101aa610c03565b91908203918211610c4157565b61008c610c03565b600090610c558261025f565b905b610c608361030e565b811115610c9057610c88610c6091610c81610c7b600161030e565b82610c34565b1692610c1a565b919050610c57565b50905090565b15610c9d57565b60405162461bcd60e51b815260206004820152601a60248201527f696e636f7272656374206e756d626572206f66206173736574730000000000006044820152606490fd5b9060208091610cef845190565b811015610cfd575b02010190565b610d056104b4565b610cf7565b15610d1157565b60405162461bcd60e51b815260206004820152601f60248201527f696e636f7272656374206e756d626572206f6620616c6c6f636174696f6e73006044820152606490fd5b15610d5d57565b60405162461bcd60e51b815260206004820152601a60248201527f64657374696e6174696f6e732063616e6e6f74206368616e67650000000000006044820152606490fd5b6001906000198114610c2c570190565b91908201809211610c4157565b15610dc657565b60405162461bcd60e51b815260206004820152601d60248201527f746f74616c20616c6c6f63617465642063616e6e6f74206368616e67650000006044820152606490fd5b90610e14825190565b90600191610e2461099f8461030e565b1480611009575b610e3490610c96565b600091610e556040610e4e610e488661030e565b87610ce2565b5101515190565b92600293610e6561099f8661030e565b1480610fe0575b610e7590610d0a565b610ea881610ea26040610e90610e8a8461030e565b8a610ce2565b510151610e9c8361030e565b90610ce2565b51015190565b610ec961099f61005084610ea26040610e90610ec38461030e565b8b610ce2565b1480610f8c575b90610ee083949392969596610d56565b6000948596610eee8361030e565b955b610f0e575b505050505050610f0861099f61008c9390565b14610dbf565b909192939495610f1d8261030e565b871015610f8657610f72610f59610f7892610f536020610ea28c6040610f4b610f458d61030e565b8c610ce2565b510151610ce2565b90610db2565b98610f536020610ea28b6040610f4b8d610e9c8d61030e565b96610da2565b949392919083979697610ef0565b95610ef5565b50610ee0610fb182610ea26040610fa5610ec38461030e565b510151610e9c8761030e565b610fd861099f61005085610ea26040610fcc610f458461030e565b510151610e9c8a61030e565b149050610ed0565b50610e75610ff56040610e4e610e488561030e565b61100161099f8761030e565b149050610e6c565b50610e34611015825190565b61102161099f8561030e565b149050610e2b56fea264697066735822122098276d2d433473a093da843b44cdd25179039405a4e8e0b31585747878b8dd7464736f6c63430008110033",
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
