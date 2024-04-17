// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancake_v3_factory

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

// PancakeV3FactoryMetaData contains all meta data concerning the PancakeV3Factory contract.
var PancakeV3FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"FeeAmountExtraInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"SetLmPoolDeployer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"WhiteListAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacingExtraInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lmPoolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setFeeAmountExtraInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol0\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol1\",\"type\":\"uint32\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lmPool\",\"type\":\"address\"}],\"name\":\"setLmPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"setLmPoolDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"setWhiteListAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516118753803806118758339818101604052602081101561003357600080fd5b50516001600160601b0319606082901b16608052600080546001600160a01b0319163390811782556040519091907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c908290a37f1bd07f61ef326b4de236f5b68f225f46ff76ee2c375ae31a06da201c49c70c12805462ffffff19166001908117909155604080518082018252600080825260208281018581526064808452600390925292517f6b16ef514f22b74729cbea5cc7babfecbdc2309e530ca716643d11fe929eed2e8054945115156101000261ff001992151560ff199096169590951791909116939093179092559151909160008051602061183583398151915291a36040805160008152600160208201528151606492600080516020611855833981519152928290030190a27f344a86d038cc67650617710ee5afca4f5d1ed60d199ecd86852cac7a55b2d3e5805462ffffff1916600a9081179091556040805180820182526000808252600160208381019182526101f4808452600390915292517f5ed261ce397475c8f8ccd7526f550ae383248415591df3d1b32ee25c9ab0af2e8054925115156101000261ff001992151560ff1990941693909317919091169190911790559151909160008051602061183583398151915291a360408051600081526001602082015281516101f492600080516020611855833981519152928290030190a27f18ea07d45b61092cf379823b7e255753fc01638d9bcaaef647c0748469d0c8cb805462ffffff191660329081179091556040805180820182526000808252600160208381019182526109c4808452600390915292517f2cb06da9fad5bc9043c9933b28e89aaba34d84764c67113fa1d4256f6b23f7558054925115156101000261ff001992151560ff1990941693909317919091169190911790559151909160008051602061183583398151915291a360408051600081526001602082015281516109c492600080516020611855833981519152928290030190a27f1ca239af1d44623dfaa87ee0cbbbe4bbeb2112df36e66deedafd694350d045cd805462ffffff191660c8908117909155604080518082018252600080825260016020838101918252612710808452600390915292517fbed90d45c8c5fb2e8fcae0027c6e57da3d943cdb82d794c1080bce28e166f2118054925115156101000261ff001992151560ff1990941693909317919091169190911790559151909160008051602061183583398151915291a3604080516000815260016020820152815161271092600080516020611855833981519152928290030190a25060805160601c61141f610416600039806106f352806110db525061141f6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637e8435e6116100975780638da5cb5b116100665780638da5cb5b146103a55780638ff38e80146103ad578063a1671295146103df578063e4a86a9914610428576100f5565b80637e8435e6146102c357806380d6a7921461030a57806388e8006d1461033d5780638a7c195f1461037a576100f5565b806322afcccb116100d357806322afcccb146101dc5780633119049a1461021557806343db87da1461021d5780635e492ac8146102bb576100f5565b806311ff5e8d146100fa57806313af4035146101375780631698ee821461016a575b600080fd5b6101356004803603604081101561011057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610463565b005b6101356004803603602081101561014d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610590565b6101b36004803603606081101561018057600080fd5b50803573ffffffffffffffffffffffffffffffffffffffff908116916020810135909116906040013562ffffff166106a3565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101fe600480360360208110156101f257600080fd5b503562ffffff166106dc565b6040805160029290920b8252519081900360200190f35b6101b36106f1565b61027a6004803603608081101561023357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135909116906fffffffffffffffffffffffffffffffff60408201358116916060013516610715565b60405180836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b6101b361086b565b610135600480360360608110156102d957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169063ffffffff60208201358116916040013516610887565b6101356004803603602081101561032057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166109a5565b61035f6004803603602081101561035357600080fd5b503562ffffff16610a9a565b60408051921515835290151560208301528051918290030190f35b6101356004803603604081101561039057600080fd5b5062ffffff813516906020013560020b610ab8565b6101b3610cc3565b610135600480360360608110156103c357600080fd5b5062ffffff813516906020810135151590604001351515610cdf565b6101b3600480360360608110156103f557600080fd5b50803573ffffffffffffffffffffffffffffffffffffffff908116916020810135909116906040013562ffffff16610e50565b6101356004803603604081101561043e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001351515611234565b60005473ffffffffffffffffffffffffffffffffffffffff163314806104a0575060055473ffffffffffffffffffffffffffffffffffffffff1633145b61050b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4e6f74206f776e6572206f72204c4d20706f6f6c206465706c6f796572000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff1663cc7e7fa2826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561057457600080fd5b505af1158015610588573d6000803e3d6000fd5b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461061657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff808516939216917fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c91a3600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600260209081526000938452604080852082529284528284209052825290205473ffffffffffffffffffffffffffffffffffffffff1681565b60016020526000908152604090205460020b81565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461079e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b604080517f85b6672900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526fffffffffffffffffffffffffffffffff8088166024840152861660448301528251908916926385b6672992606480820193918290030181600087803b15801561082b57600080fd5b505af115801561083f573d6000803e3d6000fd5b505050506040513d604081101561085557600080fd5b5080516020909101519097909650945050505050565b60055473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff16331461090d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b604080517fb0d0d21100000000000000000000000000000000000000000000000000000000815263ffffffff808516600483015283166024820152905173ffffffffffffffffffffffffffffffffffffffff85169163b0d0d21191604480830192600092919082900301818387803b15801561098857600080fd5b505af115801561099c573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a2b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca40790600090a250565b60036020526000908152604090205460ff8082169161010090041682565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b3e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b620f42408262ffffff1610610b5257600080fd5b60008160020b138015610b6957506140008160020b125b610b7257600080fd5b62ffffff8216600090815260016020526040902054600290810b900b15610b9857600080fd5b62ffffff828116600081815260016020818152604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000016600289900b9788161790558051808201825284815280830193845285855260039092528084209151825493517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909416901515177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010093151593909302929092179055517fc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc9190a3604080516000815260016020820152815162ffffff8516927fed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88928290030190a25050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff163314610d6557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b62ffffff8316600090815260016020526040902054600290810b900b610d8a57600080fd5b604080518082018252831515808252831515602080840182815262ffffff89166000818152600384528790209551865492511515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009094169390931716919091179094558451928352820152825191927fed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d8892918290030190a2505050565b60008273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610e8b57600080fd5b6000808473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1610610ec8578486610ecb565b85855b909250905073ffffffffffffffffffffffffffffffffffffffff8216610ef057600080fd5b62ffffff8416600090815260016020908152604080832054600383529281902081518083019092525460ff8082161515835261010090910416151591810191909152600291820b9182900b15801590610f4a575080602001515b610fb557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f666565206973206e6f7420617661696c61626c65207965740000000000000000604482015290519081900360640190fd5b805115611024573360009081526004602052604090205460ff16611024576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806113e16032913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84811660009081526002602090815260408083208785168452825280832062ffffff8b168452909152902054161561107057600080fd5b604080517ffad5359f00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8681166024830152858116604483015262ffffff89166064830152600285900b608483015291517f00000000000000000000000000000000000000000000000000000000000000009092169163fad5359f9160a4808201926020929091908290030181600087803b15801561112657600080fd5b505af115801561113a573d6000803e3d6000fd5b505050506040513d602081101561115057600080fd5b505173ffffffffffffffffffffffffffffffffffffffff80861660008181526002602081815260408084208a871680865290835281852062ffffff8f168087529084528286208054988a167fffffffffffffffffffffffff0000000000000000000000000000000000000000998a1681179091558287528585528387208888528552838720828852855295839020805490981686179097558151938a900b8452918301939093528251959a5093947f783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118929181900390910190a4505050509392505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112ba57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604090205460ff161515811515141561135557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f7374617465206e6f74206368616e676500000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526004602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016851515908117909155825190815291517faec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a29281900390910190a2505056fe757365722073686f756c6420626520696e20746865207768697465206c69737420666f722074686973206665652074696572a164736f6c6343000706000ac66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecced85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88",
}

// PancakeV3FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeV3FactoryMetaData.ABI instead.
var PancakeV3FactoryABI = PancakeV3FactoryMetaData.ABI

// PancakeV3FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PancakeV3FactoryMetaData.Bin instead.
var PancakeV3FactoryBin = PancakeV3FactoryMetaData.Bin

// DeployPancakeV3Factory deploys a new Ethereum contract, binding an instance of PancakeV3Factory to it.
func DeployPancakeV3Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _poolDeployer common.Address) (common.Address, *types.Transaction, *PancakeV3Factory, error) {
	parsed, err := PancakeV3FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PancakeV3FactoryBin), backend, _poolDeployer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PancakeV3Factory{PancakeV3FactoryCaller: PancakeV3FactoryCaller{contract: contract}, PancakeV3FactoryTransactor: PancakeV3FactoryTransactor{contract: contract}, PancakeV3FactoryFilterer: PancakeV3FactoryFilterer{contract: contract}}, nil
}

// PancakeV3Factory is an auto generated Go binding around an Ethereum contract.
type PancakeV3Factory struct {
	PancakeV3FactoryCaller     // Read-only binding to the contract
	PancakeV3FactoryTransactor // Write-only binding to the contract
	PancakeV3FactoryFilterer   // Log filterer for contract events
}

// PancakeV3FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeV3FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeV3FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeV3FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeV3FactorySession struct {
	Contract     *PancakeV3Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeV3FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeV3FactoryCallerSession struct {
	Contract *PancakeV3FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PancakeV3FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeV3FactoryTransactorSession struct {
	Contract     *PancakeV3FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PancakeV3FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeV3FactoryRaw struct {
	Contract *PancakeV3Factory // Generic contract binding to access the raw methods on
}

// PancakeV3FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeV3FactoryCallerRaw struct {
	Contract *PancakeV3FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeV3FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeV3FactoryTransactorRaw struct {
	Contract *PancakeV3FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeV3Factory creates a new instance of PancakeV3Factory, bound to a specific deployed contract.
func NewPancakeV3Factory(address common.Address, backend bind.ContractBackend) (*PancakeV3Factory, error) {
	contract, err := bindPancakeV3Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeV3Factory{PancakeV3FactoryCaller: PancakeV3FactoryCaller{contract: contract}, PancakeV3FactoryTransactor: PancakeV3FactoryTransactor{contract: contract}, PancakeV3FactoryFilterer: PancakeV3FactoryFilterer{contract: contract}}, nil
}

// NewPancakeV3FactoryCaller creates a new read-only instance of PancakeV3Factory, bound to a specific deployed contract.
func NewPancakeV3FactoryCaller(address common.Address, caller bind.ContractCaller) (*PancakeV3FactoryCaller, error) {
	contract, err := bindPancakeV3Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryCaller{contract: contract}, nil
}

// NewPancakeV3FactoryTransactor creates a new write-only instance of PancakeV3Factory, bound to a specific deployed contract.
func NewPancakeV3FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeV3FactoryTransactor, error) {
	contract, err := bindPancakeV3Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryTransactor{contract: contract}, nil
}

// NewPancakeV3FactoryFilterer creates a new log filterer instance of PancakeV3Factory, bound to a specific deployed contract.
func NewPancakeV3FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeV3FactoryFilterer, error) {
	contract, err := bindPancakeV3Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryFilterer{contract: contract}, nil
}

// bindPancakeV3Factory binds a generic wrapper to an already deployed contract.
func bindPancakeV3Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeV3FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV3Factory *PancakeV3FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV3Factory.Contract.PancakeV3FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV3Factory *PancakeV3FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.PancakeV3FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV3Factory *PancakeV3FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.PancakeV3FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV3Factory *PancakeV3FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV3Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV3Factory *PancakeV3FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV3Factory *PancakeV3FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3Factory *PancakeV3FactoryCaller) FeeAmountTickSpacing(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeV3Factory.contract.Call(opts, &out, "feeAmountTickSpacing", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3Factory *PancakeV3FactorySession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _PancakeV3Factory.Contract.FeeAmountTickSpacing(&_PancakeV3Factory.CallOpts, arg0)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3Factory *PancakeV3FactoryCallerSession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _PancakeV3Factory.Contract.FeeAmountTickSpacing(&_PancakeV3Factory.CallOpts, arg0)
}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3Factory *PancakeV3FactoryCaller) FeeAmountTickSpacingExtraInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	var out []interface{}
	err := _PancakeV3Factory.contract.Call(opts, &out, "feeAmountTickSpacingExtraInfo", arg0)

	outstruct := new(struct {
		WhitelistRequested bool
		Enabled            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WhitelistRequested = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3Factory *PancakeV3FactorySession) FeeAmountTickSpacingExtraInfo(arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	return _PancakeV3Factory.Contract.FeeAmountTickSpacingExtraInfo(&_PancakeV3Factory.CallOpts, arg0)
}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3Factory *PancakeV3FactoryCallerSession) FeeAmountTickSpacingExtraInfo(arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	return _PancakeV3Factory.Contract.FeeAmountTickSpacingExtraInfo(&_PancakeV3Factory.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3Factory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3Factory *PancakeV3FactorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _PancakeV3Factory.Contract.GetPool(&_PancakeV3Factory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _PancakeV3Factory.Contract.GetPool(&_PancakeV3Factory.CallOpts, arg0, arg1, arg2)
}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCaller) LmPoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3Factory.contract.Call(opts, &out, "lmPoolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3Factory *PancakeV3FactorySession) LmPoolDeployer() (common.Address, error) {
	return _PancakeV3Factory.Contract.LmPoolDeployer(&_PancakeV3Factory.CallOpts)
}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCallerSession) LmPoolDeployer() (common.Address, error) {
	return _PancakeV3Factory.Contract.LmPoolDeployer(&_PancakeV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3Factory *PancakeV3FactorySession) Owner() (common.Address, error) {
	return _PancakeV3Factory.Contract.Owner(&_PancakeV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCallerSession) Owner() (common.Address, error) {
	return _PancakeV3Factory.Contract.Owner(&_PancakeV3Factory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3Factory.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3Factory *PancakeV3FactorySession) PoolDeployer() (common.Address, error) {
	return _PancakeV3Factory.Contract.PoolDeployer(&_PancakeV3Factory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3Factory *PancakeV3FactoryCallerSession) PoolDeployer() (common.Address, error) {
	return _PancakeV3Factory.Contract.PoolDeployer(&_PancakeV3Factory.CallOpts)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3Factory *PancakeV3FactoryTransactor) CollectProtocol(opts *bind.TransactOpts, pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "collectProtocol", pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3Factory *PancakeV3FactorySession) CollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.CollectProtocol(&_PancakeV3Factory.TransactOpts, pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) CollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.CollectProtocol(&_PancakeV3Factory.TransactOpts, pool, recipient, amount0Requested, amount1Requested)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3Factory *PancakeV3FactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3Factory *PancakeV3FactorySession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.CreatePool(&_PancakeV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.CreatePool(&_PancakeV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.EnableFeeAmount(&_PancakeV3Factory.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.EnableFeeAmount(&_PancakeV3Factory.TransactOpts, fee, tickSpacing)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) SetFeeAmountExtraInfo(opts *bind.TransactOpts, fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "setFeeAmountExtraInfo", fee, whitelistRequested, enabled)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) SetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetFeeAmountExtraInfo(&_PancakeV3Factory.TransactOpts, fee, whitelistRequested, enabled)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) SetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetFeeAmountExtraInfo(&_PancakeV3Factory.TransactOpts, fee, whitelistRequested, enabled)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) SetFeeProtocol(opts *bind.TransactOpts, pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "setFeeProtocol", pool, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) SetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetFeeProtocol(&_PancakeV3Factory.TransactOpts, pool, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) SetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetFeeProtocol(&_PancakeV3Factory.TransactOpts, pool, feeProtocol0, feeProtocol1)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) SetLmPool(opts *bind.TransactOpts, pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "setLmPool", pool, lmPool)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) SetLmPool(pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetLmPool(&_PancakeV3Factory.TransactOpts, pool, lmPool)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) SetLmPool(pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetLmPool(&_PancakeV3Factory.TransactOpts, pool, lmPool)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) SetLmPoolDeployer(opts *bind.TransactOpts, _lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "setLmPoolDeployer", _lmPoolDeployer)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) SetLmPoolDeployer(_lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetLmPoolDeployer(&_PancakeV3Factory.TransactOpts, _lmPoolDeployer)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) SetLmPoolDeployer(_lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetLmPoolDeployer(&_PancakeV3Factory.TransactOpts, _lmPoolDeployer)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetOwner(&_PancakeV3Factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetOwner(&_PancakeV3Factory.TransactOpts, _owner)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactor) SetWhiteListAddress(opts *bind.TransactOpts, user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3Factory.contract.Transact(opts, "setWhiteListAddress", user, verified)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3Factory *PancakeV3FactorySession) SetWhiteListAddress(user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetWhiteListAddress(&_PancakeV3Factory.TransactOpts, user, verified)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3Factory *PancakeV3FactoryTransactorSession) SetWhiteListAddress(user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3Factory.Contract.SetWhiteListAddress(&_PancakeV3Factory.TransactOpts, user, verified)
}

// PancakeV3FactoryFeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the PancakeV3Factory contract.
type PancakeV3FactoryFeeAmountEnabledIterator struct {
	Event *PancakeV3FactoryFeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *PancakeV3FactoryFeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FactoryFeeAmountEnabled)
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
		it.Event = new(PancakeV3FactoryFeeAmountEnabled)
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
func (it *PancakeV3FactoryFeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FactoryFeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FactoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the PancakeV3Factory contract.
type PancakeV3FactoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*PancakeV3FactoryFeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryFeeAmountEnabledIterator{contract: _PancakeV3Factory.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *PancakeV3FactoryFeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FactoryFeeAmountEnabled)
				if err := _PancakeV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) ParseFeeAmountEnabled(log types.Log) (*PancakeV3FactoryFeeAmountEnabled, error) {
	event := new(PancakeV3FactoryFeeAmountEnabled)
	if err := _PancakeV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator is returned from FilterFeeAmountExtraInfoUpdated and is used to iterate over the raw logs and unpacked data for FeeAmountExtraInfoUpdated events raised by the PancakeV3Factory contract.
type PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator struct {
	Event *PancakeV3FactoryFeeAmountExtraInfoUpdated // Event containing the contract specifics and raw log

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
func (it *PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FactoryFeeAmountExtraInfoUpdated)
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
		it.Event = new(PancakeV3FactoryFeeAmountExtraInfoUpdated)
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
func (it *PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FactoryFeeAmountExtraInfoUpdated represents a FeeAmountExtraInfoUpdated event raised by the PancakeV3Factory contract.
type PancakeV3FactoryFeeAmountExtraInfoUpdated struct {
	Fee                *big.Int
	WhitelistRequested bool
	Enabled            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountExtraInfoUpdated is a free log retrieval operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) FilterFeeAmountExtraInfoUpdated(opts *bind.FilterOpts, fee []*big.Int) (*PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.FilterLogs(opts, "FeeAmountExtraInfoUpdated", feeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryFeeAmountExtraInfoUpdatedIterator{contract: _PancakeV3Factory.contract, event: "FeeAmountExtraInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeAmountExtraInfoUpdated is a free log subscription operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) WatchFeeAmountExtraInfoUpdated(opts *bind.WatchOpts, sink chan<- *PancakeV3FactoryFeeAmountExtraInfoUpdated, fee []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.WatchLogs(opts, "FeeAmountExtraInfoUpdated", feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FactoryFeeAmountExtraInfoUpdated)
				if err := _PancakeV3Factory.contract.UnpackLog(event, "FeeAmountExtraInfoUpdated", log); err != nil {
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

// ParseFeeAmountExtraInfoUpdated is a log parse operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) ParseFeeAmountExtraInfoUpdated(log types.Log) (*PancakeV3FactoryFeeAmountExtraInfoUpdated, error) {
	event := new(PancakeV3FactoryFeeAmountExtraInfoUpdated)
	if err := _PancakeV3Factory.contract.UnpackLog(event, "FeeAmountExtraInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3FactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PancakeV3Factory contract.
type PancakeV3FactoryOwnerChangedIterator struct {
	Event *PancakeV3FactoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *PancakeV3FactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FactoryOwnerChanged)
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
		it.Event = new(PancakeV3FactoryOwnerChanged)
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
func (it *PancakeV3FactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FactoryOwnerChanged represents a OwnerChanged event raised by the PancakeV3Factory contract.
type PancakeV3FactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PancakeV3FactoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryOwnerChangedIterator{contract: _PancakeV3Factory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PancakeV3FactoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FactoryOwnerChanged)
				if err := _PancakeV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) ParseOwnerChanged(log types.Log) (*PancakeV3FactoryOwnerChanged, error) {
	event := new(PancakeV3FactoryOwnerChanged)
	if err := _PancakeV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3FactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PancakeV3Factory contract.
type PancakeV3FactoryPoolCreatedIterator struct {
	Event *PancakeV3FactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *PancakeV3FactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FactoryPoolCreated)
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
		it.Event = new(PancakeV3FactoryPoolCreated)
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
func (it *PancakeV3FactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FactoryPoolCreated represents a PoolCreated event raised by the PancakeV3Factory contract.
type PancakeV3FactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*PancakeV3FactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryPoolCreatedIterator{contract: _PancakeV3Factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PancakeV3FactoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FactoryPoolCreated)
				if err := _PancakeV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) ParsePoolCreated(log types.Log) (*PancakeV3FactoryPoolCreated, error) {
	event := new(PancakeV3FactoryPoolCreated)
	if err := _PancakeV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3FactorySetLmPoolDeployerIterator is returned from FilterSetLmPoolDeployer and is used to iterate over the raw logs and unpacked data for SetLmPoolDeployer events raised by the PancakeV3Factory contract.
type PancakeV3FactorySetLmPoolDeployerIterator struct {
	Event *PancakeV3FactorySetLmPoolDeployer // Event containing the contract specifics and raw log

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
func (it *PancakeV3FactorySetLmPoolDeployerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FactorySetLmPoolDeployer)
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
		it.Event = new(PancakeV3FactorySetLmPoolDeployer)
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
func (it *PancakeV3FactorySetLmPoolDeployerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FactorySetLmPoolDeployerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FactorySetLmPoolDeployer represents a SetLmPoolDeployer event raised by the PancakeV3Factory contract.
type PancakeV3FactorySetLmPoolDeployer struct {
	LmPoolDeployer common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetLmPoolDeployer is a free log retrieval operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) FilterSetLmPoolDeployer(opts *bind.FilterOpts, lmPoolDeployer []common.Address) (*PancakeV3FactorySetLmPoolDeployerIterator, error) {

	var lmPoolDeployerRule []interface{}
	for _, lmPoolDeployerItem := range lmPoolDeployer {
		lmPoolDeployerRule = append(lmPoolDeployerRule, lmPoolDeployerItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.FilterLogs(opts, "SetLmPoolDeployer", lmPoolDeployerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactorySetLmPoolDeployerIterator{contract: _PancakeV3Factory.contract, event: "SetLmPoolDeployer", logs: logs, sub: sub}, nil
}

// WatchSetLmPoolDeployer is a free log subscription operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) WatchSetLmPoolDeployer(opts *bind.WatchOpts, sink chan<- *PancakeV3FactorySetLmPoolDeployer, lmPoolDeployer []common.Address) (event.Subscription, error) {

	var lmPoolDeployerRule []interface{}
	for _, lmPoolDeployerItem := range lmPoolDeployer {
		lmPoolDeployerRule = append(lmPoolDeployerRule, lmPoolDeployerItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.WatchLogs(opts, "SetLmPoolDeployer", lmPoolDeployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FactorySetLmPoolDeployer)
				if err := _PancakeV3Factory.contract.UnpackLog(event, "SetLmPoolDeployer", log); err != nil {
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

// ParseSetLmPoolDeployer is a log parse operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) ParseSetLmPoolDeployer(log types.Log) (*PancakeV3FactorySetLmPoolDeployer, error) {
	event := new(PancakeV3FactorySetLmPoolDeployer)
	if err := _PancakeV3Factory.contract.UnpackLog(event, "SetLmPoolDeployer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3FactoryWhiteListAddedIterator is returned from FilterWhiteListAdded and is used to iterate over the raw logs and unpacked data for WhiteListAdded events raised by the PancakeV3Factory contract.
type PancakeV3FactoryWhiteListAddedIterator struct {
	Event *PancakeV3FactoryWhiteListAdded // Event containing the contract specifics and raw log

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
func (it *PancakeV3FactoryWhiteListAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FactoryWhiteListAdded)
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
		it.Event = new(PancakeV3FactoryWhiteListAdded)
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
func (it *PancakeV3FactoryWhiteListAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FactoryWhiteListAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FactoryWhiteListAdded represents a WhiteListAdded event raised by the PancakeV3Factory contract.
type PancakeV3FactoryWhiteListAdded struct {
	User     common.Address
	Verified bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhiteListAdded is a free log retrieval operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) FilterWhiteListAdded(opts *bind.FilterOpts, user []common.Address) (*PancakeV3FactoryWhiteListAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.FilterLogs(opts, "WhiteListAdded", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FactoryWhiteListAddedIterator{contract: _PancakeV3Factory.contract, event: "WhiteListAdded", logs: logs, sub: sub}, nil
}

// WatchWhiteListAdded is a free log subscription operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) WatchWhiteListAdded(opts *bind.WatchOpts, sink chan<- *PancakeV3FactoryWhiteListAdded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeV3Factory.contract.WatchLogs(opts, "WhiteListAdded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FactoryWhiteListAdded)
				if err := _PancakeV3Factory.contract.UnpackLog(event, "WhiteListAdded", log); err != nil {
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

// ParseWhiteListAdded is a log parse operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3Factory *PancakeV3FactoryFilterer) ParseWhiteListAdded(log types.Log) (*PancakeV3FactoryWhiteListAdded, error) {
	event := new(PancakeV3FactoryWhiteListAdded)
	if err := _PancakeV3Factory.contract.UnpackLog(event, "WhiteListAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
