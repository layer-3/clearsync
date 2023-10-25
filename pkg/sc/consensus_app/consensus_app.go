// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus_app

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

// ConsensusAppMetaData contains all meta data concerning the ConsensusApp contract.
var ConsensusAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523461001b575b604051610b5a6100298239610b5a90f35b610023600080fd5b61000a56fe60806040526004361015610018575b610016600080fd5b005b60003560e01c639936d8120361000e576100306101c8565b61000e565b600080fd5b908160809103126100485790565b610050600080fd5b90565b909182601f8301121561009b575b60208235926001600160401b03841161008e575b01926020830284011161008457565b61008c600080fd5b565b610096600080fd5b610075565b6100a3600080fd5b610061565b908160409103126100485790565b90606082820312610146575b6100de8183356001600160401b038111610139575b840161003a565b926100506101018360208601356001600160401b03811161012c575b8601610053565b9390946040810135906001600160401b03821161011f575b016100a8565b610127600080fd5b610119565b610134600080fd5b6100fa565b610141600080fd5b6100d7565b61014e600080fd5b6100c2565b60005b8381106101665750506000910152565b8181015183820152602001610156565b601f01601f191690565b6101a16101766020936101aa93610195815190565b80835293849260200190565b95869101610153565b0190565b901515815260406020820181905261005092910190610180565b5034610200575b6101e66101dd3660046100b6565b9291909161095c565b906101fc6101f360405190565b928392836101ae565b0390f35b610208600080fd5b6101cf565b50634e487b7160e01b600052604160045260246000fd5b90601f01601f191681019081106001600160401b0382111761024557604052565b61024d61020d565b604052565b9061008c61025f60405190565b9283610224565b602080916001600160401b03811161027d57020190565b61028561020d565b020190565b6001600160a01b0381165b0361003557565b9050359061008c8261028a565b909291926102be6102b982610266565b610252565b93818552602080860192028301928184116102fc575b915b8383106102e35750505050565b602080916102f1848661029c565b8152019201916102d6565b610304600080fd5b6102d4565b906100509181601f82011215610325575b6020813591016102a9565b61032d600080fd5b61031a565b6001600160401b038116610295565b9050359061008c82610332565b65ffffffffffff8116610295565b9050359061008c8261034e565b9190916080818403126103e4575b6103d16103846080610252565b936103a18184356001600160401b0381116103d7575b8501610309565b855260206103b182858301610341565b9086015260406103c38282860161029c565b90860152606080930161035c565b90830152565b6103df600080fd5b61039a565b6103ec600080fd5b610377565b610050903690610369565b6004111561003557565b9050359061008c826103fc565b6101aa6020916001600160401b03811161043157601f01601f191690565b61017661020d565b90826000939282370152565b92919061008c916104586102b983610413565b948286526020860191838201111561043957610472600080fd5b610439565b906100509181601f82011215610493575b602081359101610445565b61049b600080fd5b610488565b91906104e7906040848203126104fb575b6104bb6040610252565b9360006104c88383610406565b908601526020810135906001600160401b0382116104ee575b01610477565b6020830152565b6104f6600080fd5b6104e1565b610503600080fd5b6104b1565b80610295565b9050359061008c82610508565b60ff8116610295565b9050359061008c8261051b565b919061059b906080848203126105a2575b61054c6080610252565b936000610559838361050e565b90860152602061056b8382840161050e565b90860152604061057d83828401610524565b908601526060810135906001600160401b0382116104ee5701610477565b6060830152565b6105aa600080fd5b610542565b909291926105bf6102b982610266565b938185526020808601920283019281841161061e575b80925b8484106105e6575050505050565b602080916106068587356001600160401b038111610611575b8601610531565b8152019301926105d8565b610619600080fd5b6105ff565b610626600080fd5b6105d5565b906100509181601f82011215610647575b6020813591016105af565b61064f600080fd5b61063c565b91906106bf906060848203126106e0575b61066f6060610252565b93600061067c838361029c565b9086015261069f8260208301356001600160401b0381116106d3575b83016104a0565b60208601526040810135906001600160401b0382116106c6575b0161062b565b6040830152565b6106ce600080fd5b6106b9565b6106db600080fd5b610698565b6106e8600080fd5b610665565b909291926106fd6102b982610266565b938185526020808601920283019281841161075c575b80925b848410610724575050505050565b602080916107448587356001600160401b03811161074f575b8601610654565b815201930192610716565b610757600080fd5b61073d565b610764600080fd5b610713565b906100509181601f82011215610785575b6020813591016106ed565b61078d600080fd5b61077a565b801515610295565b9050359061008c82610792565b91909160808184031261083b575b6103d16107c26080610252565b936107df8184356001600160401b03811161082e575b8501610769565b85526108008160208501356001600160401b038111610821575b8501610477565b602086015260406108138282860161035c565b90860152606080930161079a565b610829600080fd5b6107f9565b610836600080fd5b6107d8565b610843600080fd5b6107b5565b919091604081840312610899575b6103d16108636040610252565b936108808184356001600160401b03811161088c575b85016107a7565b8552602092830161050e565b610894600080fd5b610879565b6108a1600080fd5b610856565b909291926108b66102b982610266565b9381855260208086019202830192818411610915575b80925b8484106108dd575050505050565b602080916108fd8587356001600160401b038111610908575b8601610848565b8152019301926108cf565b610910600080fd5b6108f6565b61091d600080fd5b6108cc565b6100509136916108a6565b610050903690610848565b906109456102b983610413565b918252565b6100506000610938565b61005061094a565b61098061098c94929361097a61098693610974600090565b506103f1565b94610922565b9161092d565b91610a26565b600190610050610954565b6100506100506100509290565b156109ab57565b60405162461bcd60e51b815260206004820152600a60248201526907c70726f6f667c213d360b41b6044820152606490fd5b6100506100506100509260ff1690565b156109f457565b60405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b6044820152606490fd5b90610a7261005061008c94610a6b610a666020610a44610a78975190565b93610a61600095610a5b610a5788610997565b9190565b146109a4565b015190565b610ad7565b9401515190565b916109dd565b146109ed565b610a8b6100506100509290565b60ff1690565b50634e487b7160e01b600052601160045260246000fd5b60019060ff1660ff8114610aba570190565b6101aa610a91565b91908203918211610acf57565b61008c610a91565b600090610ae382610a7e565b905b610aee83610997565b811115610b1e57610b16610aee91610b0f610b096001610997565b82610ac2565b1692610aa8565b919050610ae5565b5090509056fea2646970667358221220ed7058b6dce4d02d27b75fa2afacd3c6c3e7767058bbbc6641ddd396c8cc19f764736f6c63430008110033",
}

// ConsensusAppABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsensusAppMetaData.ABI instead.
var ConsensusAppABI = ConsensusAppMetaData.ABI

// ConsensusAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsensusAppMetaData.Bin instead.
var ConsensusAppBin = ConsensusAppMetaData.Bin

// DeployConsensusApp deploys a new Ethereum contract, binding an instance of ConsensusApp to it.
func DeployConsensusApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConsensusApp, error) {
	parsed, err := ConsensusAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsensusAppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsensusApp{ConsensusAppCaller: ConsensusAppCaller{contract: contract}, ConsensusAppTransactor: ConsensusAppTransactor{contract: contract}, ConsensusAppFilterer: ConsensusAppFilterer{contract: contract}}, nil
}

// ConsensusApp is an auto generated Go binding around an Ethereum contract.
type ConsensusApp struct {
	ConsensusAppCaller     // Read-only binding to the contract
	ConsensusAppTransactor // Write-only binding to the contract
	ConsensusAppFilterer   // Log filterer for contract events
}

// ConsensusAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsensusAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsensusAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsensusAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsensusAppSession struct {
	Contract     *ConsensusApp     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsensusAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsensusAppCallerSession struct {
	Contract *ConsensusAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ConsensusAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsensusAppTransactorSession struct {
	Contract     *ConsensusAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ConsensusAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsensusAppRaw struct {
	Contract *ConsensusApp // Generic contract binding to access the raw methods on
}

// ConsensusAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsensusAppCallerRaw struct {
	Contract *ConsensusAppCaller // Generic read-only contract binding to access the raw methods on
}

// ConsensusAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsensusAppTransactorRaw struct {
	Contract *ConsensusAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsensusApp creates a new instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusApp(address common.Address, backend bind.ContractBackend) (*ConsensusApp, error) {
	contract, err := bindConsensusApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsensusApp{ConsensusAppCaller: ConsensusAppCaller{contract: contract}, ConsensusAppTransactor: ConsensusAppTransactor{contract: contract}, ConsensusAppFilterer: ConsensusAppFilterer{contract: contract}}, nil
}

// NewConsensusAppCaller creates a new read-only instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusAppCaller(address common.Address, caller bind.ContractCaller) (*ConsensusAppCaller, error) {
	contract, err := bindConsensusApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusAppCaller{contract: contract}, nil
}

// NewConsensusAppTransactor creates a new write-only instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusAppTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsensusAppTransactor, error) {
	contract, err := bindConsensusApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusAppTransactor{contract: contract}, nil
}

// NewConsensusAppFilterer creates a new log filterer instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusAppFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsensusAppFilterer, error) {
	contract, err := bindConsensusApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsensusAppFilterer{contract: contract}, nil
}

// bindConsensusApp binds a generic wrapper to an already deployed contract.
func bindConsensusApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConsensusAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusApp *ConsensusAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsensusApp.Contract.ConsensusAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusApp *ConsensusAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusApp.Contract.ConsensusAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusApp *ConsensusAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusApp.Contract.ConsensusAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusApp *ConsensusAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsensusApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusApp *ConsensusAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusApp *ConsensusAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusApp.Contract.contract.Transact(opts, method, params...)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_ConsensusApp *ConsensusAppCaller) StateIsSupported(opts *bind.CallOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	var out []interface{}
	err := _ConsensusApp.contract.Call(opts, &out, "stateIsSupported", fixedPart, proof, candidate)

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
func (_ConsensusApp *ConsensusAppSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _ConsensusApp.Contract.StateIsSupported(&_ConsensusApp.CallOpts, fixedPart, proof, candidate)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_ConsensusApp *ConsensusAppCallerSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _ConsensusApp.Contract.StateIsSupported(&_ConsensusApp.CallOpts, fixedPart, proof, candidate)
}
