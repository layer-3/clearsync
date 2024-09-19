// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivoucher_v2

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

// IVoucherVoucher is an auto generated low-level Go binding around an user-defined struct.
type IVoucherVoucher struct {
	ChainId     uint32
	Router      common.Address
	Executor    common.Address
	Beneficiary common.Address
	ExpireAt    uint64
	Nonce       *big.Int
	Data        []byte
	Signature   []byte
}

// VoucherRouterMetaData contains all meta data concerning the VoucherRouter contract.
var VoucherRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultIssuer_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultIssuer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executorIssuers\",\"inputs\":[{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultIssuer\",\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecutorIssuer\",\"inputs\":[{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"use\",\"inputs\":[{\"name\":\"vouchers\",\"type\":\"tuple[]\",\"internalType\":\"structIVoucher.Voucher[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expireAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usedVouchers\",\"inputs\":[{\"name\":\"uid\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"isUsed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Used\",\"inputs\":[{\"name\":\"voucher\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIVoucher.Voucher\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expireAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExecutor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIssuer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVouchersLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VoucherAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VoucherExpired\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620012453803806200124583398101604081905262000034916200014f565b816001600160a01b0381166200006457604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6200006f81620000c4565b5060016002556001600160a01b0381166200009d57604051635edff10b60e11b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b03929092169190911790555062000187565b600180546001600160a01b0319169055620000df81620000e2565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146200014a57600080fd5b919050565b600080604083850312156200016357600080fd5b6200016e8362000132565b91506200017e6020840162000132565b90509250929050565b6110ae80620001976000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806394765d631161007157806394765d63146101105780639a8d0ad314610143578063c092db0e1461016c578063e30c39781461017f578063ec31b60314610190578063f2fde38b146101a357600080fd5b8063142cfda8146100ae578063672bfdc6146100c3578063715018a6146100d657806379ba5097146100de5780638da5cb5b146100e6575b600080fd5b6100c16100bc366004610a83565b6101b6565b005b6100c16100d1366004610b14565b6102e0565b6100c1610331565b6100c1610345565b6000546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b61013361011e366004610b4d565b60056020526000908152604090205460ff1681565b6040519015158152602001610107565b6100f3610151366004610b14565b6004602052600090815260409020546001600160a01b031681565b6003546100f3906001600160a01b031681565b6001546001600160a01b03166100f3565b6100c161019e366004610b68565b61038e565b6100c16101b1366004610b14565b6103c4565b6101be610435565b60008190036101e0576040516378b87a6f60e11b815260040160405180910390fd5b60005b818110156102d15761021783838381811061020057610200610b9b565b90506020028101906102129190610bb1565b61045d565b61024383838381811061022c5761022c610b9b565b905060200281019061023e9190610bb1565b610589565b61026f83838381811061025857610258610b9b565b905060200281019061026a9190610bb1565b6106b5565b7fe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb08383838181106102a2576102a2610b9b565b90506020028101906102b49190610bb1565b6040516102c19190610c73565b60405180910390a16001016101e3565b506102dc6001600255565b5050565b6102e8610784565b6001600160a01b03811661030f57604051635edff10b60e11b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b610339610784565b61034360006107b1565b565b60015433906001600160a01b031681146103825760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b61038b816107b1565b50565b610396610784565b6001600160a01b03918216600090815260046020526040902080546001600160a01b03191691909216179055565b6103cc610784565b600180546001600160a01b0383166001600160a01b031990911681179091556103fd6000546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b600280540361045757604051633ee5aeb560e01b815260040160405180910390fd5b60028055565b60006004816104726060850160408601610b14565b6001600160a01b039081168252602082019290925260400160002054169050600019810161049e575050565b6001600160a01b0381166104ba57506003546001600160a01b03165b60006105506104cc60e0850185610d71565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061054a92506105179150610512905087610e85565b6107ca565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b906107e3565b9050816001600160a01b0316816001600160a01b03161461058457604051638baa579f60e01b815260040160405180910390fd5b505050565b466105976020830183610f59565b63ffffffff16146105bb57604051633d23e4d160e11b815260040160405180910390fd5b306105cc6040830160208401610b14565b6001600160a01b0316146105f35760405163466d7fef60e01b815260040160405180910390fd5b60006106056060830160408401610b14565b6001600160a01b03160361062c5760405163710c949760e01b815260040160405180910390fd5b61063c60a0820160808301610f74565b67ffffffffffffffff1642111561066657604051630abfec3f60e11b815260040160405180910390fd5b6005600061067a60c0840160a08501610b4d565b6001600160801b0316815260208101919091526040016000205460ff161561038b5760405163e58f39a760e01b815260040160405180910390fd5b6001600560006106cb60c0850160a08601610b4d565b6001600160801b0316815260208101919091526040908101600020805460ff1916921515929092179091556107069060608301908301610b14565b6001600160a01b0316631cff79cd6107246080840160608501610b14565b61073160c0850185610d71565b6040518463ffffffff1660e01b815260040161074f93929190610f8f565b600060405180830381600087803b15801561076957600080fd5b505af115801561077d573d6000803e3d6000fd5b5050505050565b6000546001600160a01b031633146103435760405163118cdaa760e01b8152336004820152602401610379565b600180546001600160a01b031916905561038b8161080d565b60006107d58261085d565b805190602001209050919050565b6000806000806107f386866108ae565b92509250925061080382826108fb565b5090949350505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6060816000015182602001518360400151846060015185608001518660a001518760c001516040516020016108989796959493929190610fbd565b6040516020818303038152906040529050919050565b600080600083516041036108e85760208401516040850151606086015160001a6108da888285856109b4565b9550955095505050506108f4565b50508151600091506002905b9250925092565b600082600381111561090f5761090f611062565b03610918575050565b600182600381111561092c5761092c611062565b0361094a5760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561095e5761095e611062565b0361097f5760405163fce698f760e01b815260048101829052602401610379565b600382600381111561099357610993611062565b036102dc576040516335e2f38360e21b815260048101829052602401610379565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156109ef5750600091506003905082610a79565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610a43573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610a6f57506000925060019150829050610a79565b9250600091508190505b9450945094915050565b60008060208385031215610a9657600080fd5b823567ffffffffffffffff80821115610aae57600080fd5b818501915085601f830112610ac257600080fd5b813581811115610ad157600080fd5b8660208260051b8501011115610ae657600080fd5b60209290920196919550909350505050565b80356001600160a01b0381168114610b0f57600080fd5b919050565b600060208284031215610b2657600080fd5b610b2f82610af8565b9392505050565b80356001600160801b0381168114610b0f57600080fd5b600060208284031215610b5f57600080fd5b610b2f82610b36565b60008060408385031215610b7b57600080fd5b610b8483610af8565b9150610b9260208401610af8565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b6000823560fe19833603018112610bc757600080fd5b9190910192915050565b803563ffffffff81168114610b0f57600080fd5b803567ffffffffffffffff81168114610b0f57600080fd5b6000808335601e19843603018112610c1457600080fd5b830160208101925035905067ffffffffffffffff811115610c3457600080fd5b803603821315610c4357600080fd5b9250929050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815263ffffffff610c8583610bd1565b1660208201526000610c9960208401610af8565b6001600160a01b038116604084015250610cb560408401610af8565b6001600160a01b038116606084015250610cd160608401610af8565b6001600160a01b038116608084015250610ced60808401610be5565b67ffffffffffffffff811660a084015250610d0a60a08401610b36565b6001600160801b03811660c084015250610d2760c0840184610bfd565b6101008060e0860152610d3f61012086018385610c4a565b9250610d4e60e0870187610bfd565b868503601f1901838801529250610d66848483610c4a565b979650505050505050565b6000808335601e19843603018112610d8857600080fd5b83018035915067ffffffffffffffff821115610da357600080fd5b602001915036819003821315610c4357600080fd5b634e487b7160e01b600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715610df257610df2610db8565b60405290565b600082601f830112610e0957600080fd5b813567ffffffffffffffff80821115610e2457610e24610db8565b604051601f8301601f19908116603f01168101908282118183101715610e4c57610e4c610db8565b81604052838152866020858801011115610e6557600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006101008236031215610e9857600080fd5b610ea0610dce565b610ea983610bd1565b8152610eb760208401610af8565b6020820152610ec860408401610af8565b6040820152610ed960608401610af8565b6060820152610eea60808401610be5565b6080820152610efb60a08401610b36565b60a082015260c083013567ffffffffffffffff80821115610f1b57600080fd5b610f2736838701610df8565b60c084015260e0850135915080821115610f4057600080fd5b50610f4d36828601610df8565b60e08301525092915050565b600060208284031215610f6b57600080fd5b610b2f82610bd1565b600060208284031215610f8657600080fd5b610b2f82610be5565b6001600160a01b0384168152604060208201819052600090610fb49083018486610c4a565b95945050505050565b63ffffffff881681526000602060018060a01b03808a166020850152808916604085015280881660608501525067ffffffffffffffff861660808401526001600160801b03851660a084015260e060c084015283518060e085015260005b81811015611038578581018301518582016101000152820161101b565b506101009150600082828601015281601f19601f8301168501019250505098975050505050505050565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220f0be4b440300366a1f44df918c6441df4be696cf62d77b8dc98db7dab7aeb80164736f6c63430008170033",
}

// VoucherRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use VoucherRouterMetaData.ABI instead.
var VoucherRouterABI = VoucherRouterMetaData.ABI

// VoucherRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoucherRouterMetaData.Bin instead.
var VoucherRouterBin = VoucherRouterMetaData.Bin

// DeployVoucherRouter deploys a new Ethereum contract, binding an instance of VoucherRouter to it.
func DeployVoucherRouter(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, defaultIssuer_ common.Address) (common.Address, *types.Transaction, *VoucherRouter, error) {
	parsed, err := VoucherRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoucherRouterBin), backend, owner, defaultIssuer_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VoucherRouter{VoucherRouterCaller: VoucherRouterCaller{contract: contract}, VoucherRouterTransactor: VoucherRouterTransactor{contract: contract}, VoucherRouterFilterer: VoucherRouterFilterer{contract: contract}}, nil
}

// VoucherRouter is an auto generated Go binding around an Ethereum contract.
type VoucherRouter struct {
	VoucherRouterCaller     // Read-only binding to the contract
	VoucherRouterTransactor // Write-only binding to the contract
	VoucherRouterFilterer   // Log filterer for contract events
}

// VoucherRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoucherRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoucherRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoucherRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoucherRouterSession struct {
	Contract     *VoucherRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoucherRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoucherRouterCallerSession struct {
	Contract *VoucherRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VoucherRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoucherRouterTransactorSession struct {
	Contract     *VoucherRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VoucherRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoucherRouterRaw struct {
	Contract *VoucherRouter // Generic contract binding to access the raw methods on
}

// VoucherRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoucherRouterCallerRaw struct {
	Contract *VoucherRouterCaller // Generic read-only contract binding to access the raw methods on
}

// VoucherRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoucherRouterTransactorRaw struct {
	Contract *VoucherRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoucherRouter creates a new instance of VoucherRouter, bound to a specific deployed contract.
func NewVoucherRouter(address common.Address, backend bind.ContractBackend) (*VoucherRouter, error) {
	contract, err := bindVoucherRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoucherRouter{VoucherRouterCaller: VoucherRouterCaller{contract: contract}, VoucherRouterTransactor: VoucherRouterTransactor{contract: contract}, VoucherRouterFilterer: VoucherRouterFilterer{contract: contract}}, nil
}

// NewVoucherRouterCaller creates a new read-only instance of VoucherRouter, bound to a specific deployed contract.
func NewVoucherRouterCaller(address common.Address, caller bind.ContractCaller) (*VoucherRouterCaller, error) {
	contract, err := bindVoucherRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoucherRouterCaller{contract: contract}, nil
}

// NewVoucherRouterTransactor creates a new write-only instance of VoucherRouter, bound to a specific deployed contract.
func NewVoucherRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*VoucherRouterTransactor, error) {
	contract, err := bindVoucherRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoucherRouterTransactor{contract: contract}, nil
}

// NewVoucherRouterFilterer creates a new log filterer instance of VoucherRouter, bound to a specific deployed contract.
func NewVoucherRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*VoucherRouterFilterer, error) {
	contract, err := bindVoucherRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoucherRouterFilterer{contract: contract}, nil
}

// bindVoucherRouter binds a generic wrapper to an already deployed contract.
func bindVoucherRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoucherRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoucherRouter *VoucherRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoucherRouter.Contract.VoucherRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoucherRouter *VoucherRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoucherRouter.Contract.VoucherRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoucherRouter *VoucherRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoucherRouter.Contract.VoucherRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoucherRouter *VoucherRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoucherRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoucherRouter *VoucherRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoucherRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoucherRouter *VoucherRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoucherRouter.Contract.contract.Transact(opts, method, params...)
}

// DefaultIssuer is a free data retrieval call binding the contract method 0xc092db0e.
//
// Solidity: function defaultIssuer() view returns(address)
func (_VoucherRouter *VoucherRouterCaller) DefaultIssuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoucherRouter.contract.Call(opts, &out, "defaultIssuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultIssuer is a free data retrieval call binding the contract method 0xc092db0e.
//
// Solidity: function defaultIssuer() view returns(address)
func (_VoucherRouter *VoucherRouterSession) DefaultIssuer() (common.Address, error) {
	return _VoucherRouter.Contract.DefaultIssuer(&_VoucherRouter.CallOpts)
}

// DefaultIssuer is a free data retrieval call binding the contract method 0xc092db0e.
//
// Solidity: function defaultIssuer() view returns(address)
func (_VoucherRouter *VoucherRouterCallerSession) DefaultIssuer() (common.Address, error) {
	return _VoucherRouter.Contract.DefaultIssuer(&_VoucherRouter.CallOpts)
}

// ExecutorIssuers is a free data retrieval call binding the contract method 0x9a8d0ad3.
//
// Solidity: function executorIssuers(address executor) view returns(address issuer)
func (_VoucherRouter *VoucherRouterCaller) ExecutorIssuers(opts *bind.CallOpts, executor common.Address) (common.Address, error) {
	var out []interface{}
	err := _VoucherRouter.contract.Call(opts, &out, "executorIssuers", executor)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorIssuers is a free data retrieval call binding the contract method 0x9a8d0ad3.
//
// Solidity: function executorIssuers(address executor) view returns(address issuer)
func (_VoucherRouter *VoucherRouterSession) ExecutorIssuers(executor common.Address) (common.Address, error) {
	return _VoucherRouter.Contract.ExecutorIssuers(&_VoucherRouter.CallOpts, executor)
}

// ExecutorIssuers is a free data retrieval call binding the contract method 0x9a8d0ad3.
//
// Solidity: function executorIssuers(address executor) view returns(address issuer)
func (_VoucherRouter *VoucherRouterCallerSession) ExecutorIssuers(executor common.Address) (common.Address, error) {
	return _VoucherRouter.Contract.ExecutorIssuers(&_VoucherRouter.CallOpts, executor)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VoucherRouter *VoucherRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoucherRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VoucherRouter *VoucherRouterSession) Owner() (common.Address, error) {
	return _VoucherRouter.Contract.Owner(&_VoucherRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VoucherRouter *VoucherRouterCallerSession) Owner() (common.Address, error) {
	return _VoucherRouter.Contract.Owner(&_VoucherRouter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VoucherRouter *VoucherRouterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoucherRouter.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VoucherRouter *VoucherRouterSession) PendingOwner() (common.Address, error) {
	return _VoucherRouter.Contract.PendingOwner(&_VoucherRouter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VoucherRouter *VoucherRouterCallerSession) PendingOwner() (common.Address, error) {
	return _VoucherRouter.Contract.PendingOwner(&_VoucherRouter.CallOpts)
}

// UsedVouchers is a free data retrieval call binding the contract method 0x94765d63.
//
// Solidity: function usedVouchers(uint128 uid) view returns(bool isUsed)
func (_VoucherRouter *VoucherRouterCaller) UsedVouchers(opts *bind.CallOpts, uid *big.Int) (bool, error) {
	var out []interface{}
	err := _VoucherRouter.contract.Call(opts, &out, "usedVouchers", uid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedVouchers is a free data retrieval call binding the contract method 0x94765d63.
//
// Solidity: function usedVouchers(uint128 uid) view returns(bool isUsed)
func (_VoucherRouter *VoucherRouterSession) UsedVouchers(uid *big.Int) (bool, error) {
	return _VoucherRouter.Contract.UsedVouchers(&_VoucherRouter.CallOpts, uid)
}

// UsedVouchers is a free data retrieval call binding the contract method 0x94765d63.
//
// Solidity: function usedVouchers(uint128 uid) view returns(bool isUsed)
func (_VoucherRouter *VoucherRouterCallerSession) UsedVouchers(uid *big.Int) (bool, error) {
	return _VoucherRouter.Contract.UsedVouchers(&_VoucherRouter.CallOpts, uid)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VoucherRouter *VoucherRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoucherRouter.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VoucherRouter *VoucherRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _VoucherRouter.Contract.AcceptOwnership(&_VoucherRouter.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VoucherRouter *VoucherRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VoucherRouter.Contract.AcceptOwnership(&_VoucherRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VoucherRouter *VoucherRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoucherRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VoucherRouter *VoucherRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _VoucherRouter.Contract.RenounceOwnership(&_VoucherRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VoucherRouter *VoucherRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VoucherRouter.Contract.RenounceOwnership(&_VoucherRouter.TransactOpts)
}

// SetDefaultIssuer is a paid mutator transaction binding the contract method 0x672bfdc6.
//
// Solidity: function setDefaultIssuer(address issuer) returns()
func (_VoucherRouter *VoucherRouterTransactor) SetDefaultIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _VoucherRouter.contract.Transact(opts, "setDefaultIssuer", issuer)
}

// SetDefaultIssuer is a paid mutator transaction binding the contract method 0x672bfdc6.
//
// Solidity: function setDefaultIssuer(address issuer) returns()
func (_VoucherRouter *VoucherRouterSession) SetDefaultIssuer(issuer common.Address) (*types.Transaction, error) {
	return _VoucherRouter.Contract.SetDefaultIssuer(&_VoucherRouter.TransactOpts, issuer)
}

// SetDefaultIssuer is a paid mutator transaction binding the contract method 0x672bfdc6.
//
// Solidity: function setDefaultIssuer(address issuer) returns()
func (_VoucherRouter *VoucherRouterTransactorSession) SetDefaultIssuer(issuer common.Address) (*types.Transaction, error) {
	return _VoucherRouter.Contract.SetDefaultIssuer(&_VoucherRouter.TransactOpts, issuer)
}

// SetExecutorIssuer is a paid mutator transaction binding the contract method 0xec31b603.
//
// Solidity: function setExecutorIssuer(address executor, address issuer) returns()
func (_VoucherRouter *VoucherRouterTransactor) SetExecutorIssuer(opts *bind.TransactOpts, executor common.Address, issuer common.Address) (*types.Transaction, error) {
	return _VoucherRouter.contract.Transact(opts, "setExecutorIssuer", executor, issuer)
}

// SetExecutorIssuer is a paid mutator transaction binding the contract method 0xec31b603.
//
// Solidity: function setExecutorIssuer(address executor, address issuer) returns()
func (_VoucherRouter *VoucherRouterSession) SetExecutorIssuer(executor common.Address, issuer common.Address) (*types.Transaction, error) {
	return _VoucherRouter.Contract.SetExecutorIssuer(&_VoucherRouter.TransactOpts, executor, issuer)
}

// SetExecutorIssuer is a paid mutator transaction binding the contract method 0xec31b603.
//
// Solidity: function setExecutorIssuer(address executor, address issuer) returns()
func (_VoucherRouter *VoucherRouterTransactorSession) SetExecutorIssuer(executor common.Address, issuer common.Address) (*types.Transaction, error) {
	return _VoucherRouter.Contract.SetExecutorIssuer(&_VoucherRouter.TransactOpts, executor, issuer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VoucherRouter *VoucherRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VoucherRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VoucherRouter *VoucherRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VoucherRouter.Contract.TransferOwnership(&_VoucherRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VoucherRouter *VoucherRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VoucherRouter.Contract.TransferOwnership(&_VoucherRouter.TransactOpts, newOwner)
}

// Use is a paid mutator transaction binding the contract method 0x142cfda8.
//
// Solidity: function use((uint32,address,address,address,uint64,uint128,bytes,bytes)[] vouchers) returns()
func (_VoucherRouter *VoucherRouterTransactor) Use(opts *bind.TransactOpts, vouchers []IVoucherVoucher) (*types.Transaction, error) {
	return _VoucherRouter.contract.Transact(opts, "use", vouchers)
}

// Use is a paid mutator transaction binding the contract method 0x142cfda8.
//
// Solidity: function use((uint32,address,address,address,uint64,uint128,bytes,bytes)[] vouchers) returns()
func (_VoucherRouter *VoucherRouterSession) Use(vouchers []IVoucherVoucher) (*types.Transaction, error) {
	return _VoucherRouter.Contract.Use(&_VoucherRouter.TransactOpts, vouchers)
}

// Use is a paid mutator transaction binding the contract method 0x142cfda8.
//
// Solidity: function use((uint32,address,address,address,uint64,uint128,bytes,bytes)[] vouchers) returns()
func (_VoucherRouter *VoucherRouterTransactorSession) Use(vouchers []IVoucherVoucher) (*types.Transaction, error) {
	return _VoucherRouter.Contract.Use(&_VoucherRouter.TransactOpts, vouchers)
}

// VoucherRouterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the VoucherRouter contract.
type VoucherRouterOwnershipTransferStartedIterator struct {
	Event *VoucherRouterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *VoucherRouterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherRouterOwnershipTransferStarted)
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
		it.Event = new(VoucherRouterOwnershipTransferStarted)
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
func (it *VoucherRouterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherRouterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherRouterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the VoucherRouter contract.
type VoucherRouterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VoucherRouter *VoucherRouterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VoucherRouterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VoucherRouter.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoucherRouterOwnershipTransferStartedIterator{contract: _VoucherRouter.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VoucherRouter *VoucherRouterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *VoucherRouterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VoucherRouter.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherRouterOwnershipTransferStarted)
				if err := _VoucherRouter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VoucherRouter *VoucherRouterFilterer) ParseOwnershipTransferStarted(log types.Log) (*VoucherRouterOwnershipTransferStarted, error) {
	event := new(VoucherRouterOwnershipTransferStarted)
	if err := _VoucherRouter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VoucherRouter contract.
type VoucherRouterOwnershipTransferredIterator struct {
	Event *VoucherRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VoucherRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherRouterOwnershipTransferred)
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
		it.Event = new(VoucherRouterOwnershipTransferred)
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
func (it *VoucherRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherRouterOwnershipTransferred represents a OwnershipTransferred event raised by the VoucherRouter contract.
type VoucherRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VoucherRouter *VoucherRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VoucherRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VoucherRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoucherRouterOwnershipTransferredIterator{contract: _VoucherRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VoucherRouter *VoucherRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VoucherRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VoucherRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherRouterOwnershipTransferred)
				if err := _VoucherRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VoucherRouter *VoucherRouterFilterer) ParseOwnershipTransferred(log types.Log) (*VoucherRouterOwnershipTransferred, error) {
	event := new(VoucherRouterOwnershipTransferred)
	if err := _VoucherRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherRouterUsedIterator is returned from FilterUsed and is used to iterate over the raw logs and unpacked data for Used events raised by the VoucherRouter contract.
type VoucherRouterUsedIterator struct {
	Event *VoucherRouterUsed // Event containing the contract specifics and raw log

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
func (it *VoucherRouterUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherRouterUsed)
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
		it.Event = new(VoucherRouterUsed)
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
func (it *VoucherRouterUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherRouterUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherRouterUsed represents a Used event raised by the VoucherRouter contract.
type VoucherRouterUsed struct {
	Voucher IVoucherVoucher
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUsed is a free log retrieval operation binding the contract event 0xe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb0.
//
// Solidity: event Used((uint32,address,address,address,uint64,uint128,bytes,bytes) voucher)
func (_VoucherRouter *VoucherRouterFilterer) FilterUsed(opts *bind.FilterOpts) (*VoucherRouterUsedIterator, error) {

	logs, sub, err := _VoucherRouter.contract.FilterLogs(opts, "Used")
	if err != nil {
		return nil, err
	}
	return &VoucherRouterUsedIterator{contract: _VoucherRouter.contract, event: "Used", logs: logs, sub: sub}, nil
}

// WatchUsed is a free log subscription operation binding the contract event 0xe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb0.
//
// Solidity: event Used((uint32,address,address,address,uint64,uint128,bytes,bytes) voucher)
func (_VoucherRouter *VoucherRouterFilterer) WatchUsed(opts *bind.WatchOpts, sink chan<- *VoucherRouterUsed) (event.Subscription, error) {

	logs, sub, err := _VoucherRouter.contract.WatchLogs(opts, "Used")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherRouterUsed)
				if err := _VoucherRouter.contract.UnpackLog(event, "Used", log); err != nil {
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

// ParseUsed is a log parse operation binding the contract event 0xe119867e6fc31f0cd6fded9dd3fdf7841204668080a573db5e5bd791a78cbbb0.
//
// Solidity: event Used((uint32,address,address,address,uint64,uint128,bytes,bytes) voucher)
func (_VoucherRouter *VoucherRouterFilterer) ParseUsed(log types.Log) (*VoucherRouterUsed, error) {
	event := new(VoucherRouterUsed)
	if err := _VoucherRouter.contract.UnpackLog(event, "Used", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
