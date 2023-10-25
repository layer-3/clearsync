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
	Bin: "0x60806040523461001b575b60405161111c610029823961111c90f35b610023600080fd5b61000a56fe60806040526004361015610018575b610016600080fd5b005b60003560e01c639936d8120361000e576100306101c8565b61000e565b600080fd5b908160809103126100485790565b610050600080fd5b90565b909182601f8301121561009b575b60208235926001600160401b03841161008e575b01926020830284011161008457565b61008c600080fd5b565b610096600080fd5b610075565b6100a3600080fd5b610061565b908160409103126100485790565b90606082820312610146575b6100de8183356001600160401b038111610139575b840161003a565b926100506101018360208601356001600160401b03811161012c575b8601610053565b9390946040810135906001600160401b03821161011f575b016100a8565b610127600080fd5b610119565b610134600080fd5b6100fa565b610141600080fd5b6100d7565b61014e600080fd5b6100c2565b60005b8381106101665750506000910152565b8181015183820152602001610156565b601f01601f191690565b6101a16101766020936101aa93610195815190565b80835293849260200190565b95869101610153565b0190565b901515815260406020820181905261005092910190610180565b5034610200575b6101e66101dd3660046100b6565b929190916109ce565b906101fc6101f360405190565b928392836101ae565b0390f35b610208600080fd5b6101cf565b903590601e193682900301821215610252575b0160208135916001600160401b038311610245575b0191602082023603831361008457565b61024d600080fd5b610235565b61025a600080fd5b610220565b61026c6100506100509290565b60ff1690565b50634e487b7160e01b600052604160045260246000fd5b90601f01601f191681019081106001600160401b038211176102aa57604052565b6102b2610272565b604052565b9061008c6102c460405190565b9283610289565b6101aa6020916001600160401b0381116102e957601f01601f191690565b610176610272565b906103036102fe836102cb565b6102b7565b918252565b61031260126102f1565b71372830b93a34b1b4b830b73a3990109e901960711b602082015290565b610050610308565b6100506100506100509290565b805b0361003557565b3561005081610345565b1561035f57565b60405162461bcd60e51b8152602060048201526015602482015274021756e616e696d6f75733b207c70726f6f667c3d3605c1b6044820152606490fd5b0390fd5b903590607e1936829003018212156103b6570190565b6101aa600080fd5b65ffffffffffff8116610347565b35610050816103be565b6103e36100506100509290565b65ffffffffffff1690565b61005060006102f1565b6100506103ee565b801515610347565b3561005081610400565b1561041957565b60405162461bcd60e51b815260206004820152601f60248201527f2166696e616c3b207475726e4e756d3e3d33202626207c70726f6f667c3d30006044820152606490fd5b50634e487b7160e01b600052603260045260246000fd5b903590603e1936829003018212156103b6570190565b90610050926020918110156104a4575b02810190610475565b6104ac61045e565b61049b565b602080916001600160401b0381116104c857020190565b6104d0610272565b020190565b6001600160a01b038116610347565b9050359061008c826104d5565b6004111561003557565b9050359061008c826104f1565b90826000939282370152565b92919061008c916105276102fe836102cb565b948286526020860191838201111561050857610541600080fd5b610508565b906100509181601f82011215610562575b602081359101610514565b61056a600080fd5b610557565b91906105b6906040848203126105ca575b61058a60406102b7565b93600061059783836104fb565b908601526020810135906001600160401b0382116105bd575b01610546565b6020830152565b6105c5600080fd5b6105b0565b6105d2600080fd5b610580565b9050359061008c82610345565b60ff8116610347565b9050359061008c826105e4565b91906106649060808482031261066b575b61061560806102b7565b93600061062283836105d7565b908601526020610634838284016105d7565b908601526040610646838284016105ed565b908601526060810135906001600160401b0382116105bd5701610546565b6060830152565b610673600080fd5b61060b565b909291926106886102fe826104b1565b93818552602080860192028301928184116106e7575b80925b8484106106af575050505050565b602080916106cf8587356001600160401b0381116106da575b86016105fa565b8152019301926106a1565b6106e2600080fd5b6106c8565b6106ef600080fd5b61069e565b906100509181601f82011215610710575b602081359101610678565b610718600080fd5b610705565b9190610788906060848203126107a9575b61073860606102b7565b93600061074583836104e4565b908601526107688260208301356001600160401b03811161079c575b830161056f565b60208601526040810135906001600160401b03821161078f575b016106f4565b6040830152565b610797600080fd5b610782565b6107a4600080fd5b610761565b6107b1600080fd5b61072e565b909291926107c66102fe826104b1565b9381855260208086019202830192818411610825575b80925b8484106107ed575050505050565b6020809161080d8587356001600160401b038111610818575b860161071d565b8152019301926107df565b610820600080fd5b610806565b61082d600080fd5b6107dc565b906100509181601f8201121561084e575b6020813591016107b6565b610856600080fd5b610843565b9050359061008c826103be565b9050359061008c82610400565b91909160808184031261090f575b6108ef61089060806102b7565b936108ad8184356001600160401b038111610902575b8501610832565b85526108ce8160208501356001600160401b0381116108f5575b8501610546565b602086015260406108e18282860161085b565b908601526060809301610868565b90830152565b6108fd600080fd5b6108c7565b61090a600080fd5b6108a6565b610917600080fd5b610883565b91909160408184031261096d575b6108ef61093760406102b7565b936109548184356001600160401b038111610960575b8501610875565b855260209283016105d7565b610968600080fd5b61094d565b610975600080fd5b61092a565b61005090369061091c565b6100506100506100509260ff1690565b6100509136916107b6565b6109aa60106102f1565b6f0c4c2c840e0e4dedecc40d8cadccee8d60831b602082015290565b6100506109a0565b6109ef6109ea6109f4929593956109e3600090565b508061020d565b905090565b61025f565b926109ff600261025f565b60ff851603610c615782600094610a1586610338565b8514610b3057600194610a2e610a2a87610338565b9190565b14610a41575050505050906100506109c6565b610b1386610b0a610b02610af9610af08888610ae7610ad1610b279e9f610ade8f9d91610b219f8f90610b1b9f610a99610a93610a8e610a9f94610acb99610a8889610338565b9161048b565b61097a565b91610985565b90610d92565b610ad6610ad1610acb610ac2610ab98c8c610a8889610338565b868101906103a0565b8581019061020d565b90610995565b610e56565b01809e6103a0565b8a81019061020d565b610a8887610338565b848101906103a0565b8381019061020d565b9590976103a0565b9081019061020d565b929094610995565b92610995565b90610fbd565b906100506103f8565b91505081939250610b62610b5b610b54610b4f6020610b68960161034e565b610cbf565b9260ff1690565b9160ff1690565b14610358565b80820190610b99610b8d610b876040610b8186886103a0565b016103cc565b926103d6565b9165ffffffffffff1690565b14610c5457610bad6040610b8183856103a0565b91600192610bbd610b8d856103d6565b14610c4957610bd16040610b8184846103a0565b610bde610b8d60036103d6565b14610c2c5760405162461bcd60e51b81528061039c600482016020808252818101527f6261642063616e646964617465207475726e4e756d3b207c70726f6f667c3d30604082015260600190565b6060610c3e610b2793610c44936103a0565b01610408565b610412565b5050906100506103f8565b50506001906100506103f8565b50505050600090610050610330565b50634e487b7160e01b600052601160045260246000fd5b60019060ff1660ff8114610c99570190565b6101aa610c70565b9190610cac565b9290565b8203918211610cb757565b61008c610c70565b600090610ccb8261025f565b905b610cd683610338565b811115610d0657610cfe610cd691610cf7610cf16001610338565b82610ca1565b1692610c87565b919050610ccd565b50905090565b15610d1357565b60405162461bcd60e51b8152602060048201526015602482015274706f737466756e642e7475726e4e756d20213d203160581b6044820152606490fd5b15610d5757565b60405162461bcd60e51b8152602060048201526013602482015272706f737466756e642021756e616e696d6f757360681b6044820152606490fd5b90610dd7610a93610ca8610b4f602061008c96610dd2610dbf60406000840151015165ffffffffffff1690565b610dcc610b8d60016103d6565b14610d0c565b015190565b14610d50565b6001906000198114610c99570190565b9060208091610dfa845190565b811015610e08575b02010190565b610e1061045e565b610e02565b15610e1c57565b60405162461bcd60e51b81526020600482015260126024820152717c616c6c6f636174696f6e737c20213d203160701b6044820152606490fd5b90610e616000610338565b610e6c610050845190565b811015610ead5780610ea3610e906040610e89610ea89588610ded565b5101515190565b610e9d610a2a6001610338565b14610e15565b610ddd565b610e61565b509050565b15610eb957565b60405162461bcd60e51b815260206004820152601860248201527f7c6f7574636f6d65417c20213d207c6f7574636f6d65427c00000000000000006044820152606490fd5b15610f0557565b60405162461bcd60e51b815260206004820152600e60248201526d0c2e6e6cae840dad2e6dac2e8c6d60931b6044820152606490fd5b15610f4257565b60405162461bcd60e51b8152602060048201526015602482015274064657374696e6174696f6e206d757374207377617605c1b6044820152606490fd5b15610f8657565b60405162461bcd60e51b815260206004820152600f60248201526e0c2dadeeadce840dad2e6dac2e8c6d608b1b6044820152606490fd5b9190610fde610fca845190565b610fd8610a2a610050855190565b14610eb2565b600092610fea84610338565b610ff5610050835190565b8110156110df5780610ea3866110d4610a2a61005060206110ce818b610dd26110ab604061109f8f9e6110da9f8d6110928561108985856110848d61107e6110716110658a6110508861105f6110989f611050838692610ded565b5101516001600160a01b031690565b96610ded565b6001600160a01b031690565b916001600160a01b031690565b14610efe565b610ded565b51015191610338565b90610ded565b5195610ded565b5101516110928b610338565b51986110c7610a2a6100506110c08487015190565b938d015190565b1415610f3b565b94015190565b14610f7f565b610fea565b505050905056fea264697066735822122086abe5eaace65421c26487f0970d8d88dac8299ce4975c9c832b71c1b4aa9f7b64736f6c63430008110033",
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
