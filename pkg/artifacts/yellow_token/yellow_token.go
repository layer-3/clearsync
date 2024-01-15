// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yellow_token

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

// YellowTokenMetaData contains all meta data concerning the YellowToken contract.
var YellowTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b50604051620015a4380380620015a48339810160408190526200003591620003fe565b6040805180820190915260018152603160f81b6020820152839081908185600362000061838262000502565b50600462000070828262000502565b50620000829150839050600562000140565b610120526200009381600662000140565b61014052815160208084019190912060e052815190820120610100524660a0526200012160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525062000137338262000179565b5050506200064a565b600060208351101562000160576200015883620001bb565b905062000173565b816200016d848262000502565b5060ff90505b92915050565b6001600160a01b038216620001a95760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b620001b760008383620001fe565b5050565b600080829050601f81511115620001e9578260405163305a27a960e01b8152600401620001a09190620005ce565b8051620001f68262000603565b179392505050565b6001600160a01b0383166200022d57806002600082825462000221919062000628565b90915550620002a19050565b6001600160a01b03831660009081526020819052604090205481811015620002825760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620001a0565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216620002bf57600280548290039055620002de565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200032491815260200190565b60405180910390a3505050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620003645781810151838201526020016200034a565b50506000910152565b600082601f8301126200037f57600080fd5b81516001600160401b03808211156200039c576200039c62000331565b604051601f8301601f19908116603f01168101908282118183101715620003c757620003c762000331565b81604052838152866020858801011115620003e157600080fd5b620003f484602083016020890162000347565b9695505050505050565b6000806000606084860312156200041457600080fd5b83516001600160401b03808211156200042c57600080fd5b6200043a878388016200036d565b945060208601519150808211156200045157600080fd5b5062000460868287016200036d565b925050604084015190509250925092565b600181811c908216806200048657607f821691505b602082108103620004a757634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620004fd576000816000526020600020601f850160051c81016020861015620004d85750805b601f850160051c820191505b81811015620004f957828155600101620004e4565b5050505b505050565b81516001600160401b038111156200051e576200051e62000331565b62000536816200052f845462000471565b84620004ad565b602080601f8311600181146200056e5760008415620005555750858301515b600019600386901b1c1916600185901b178555620004f9565b600085815260208120601f198616915b828110156200059f578886015182559484019460019091019084016200057e565b5085821015620005be5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020815260008251806020840152620005ef81604085016020870162000347565b601f01601f19169190910160400192915050565b80516020808301519190811015620004a75760001960209190910360031b1b16919050565b808201808211156200017357634e487b7160e01b600052601160045260246000fd5b60805160a05160c05160e051610100516101205161014051610eff620006a560003960006106fe015260006106d10152600061067901526000610651015260006105ac015260006105d6015260006106000152610eff6000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806370a082311161008c57806395d89b411161006657806395d89b41146101a8578063a9059cbb146101b0578063d505accf146101c3578063dd62ed3e146101d857600080fd5b806370a08231146101515780637ecebe001461017a57806384b0196e1461018d57600080fd5b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461011557806323b872dd14610127578063313ce5671461013a5780633644e51514610149575b600080fd5b6100dc610211565b6040516100e99190610c62565b60405180910390f35b610105610100366004610c98565b6102a3565b60405190151581526020016100e9565b6002545b6040519081526020016100e9565b610105610135366004610cc2565b6102bd565b604051600881526020016100e9565b6101196102e1565b61011961015f366004610cfe565b6001600160a01b031660009081526020819052604090205490565b610119610188366004610cfe565b6102f0565b61019561030e565b6040516100e99796959493929190610d19565b6100dc610354565b6101056101be366004610c98565b610363565b6101d66101d1366004610db2565b610371565b005b6101196101e6366004610e25565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60606003805461022090610e58565b80601f016020809104026020016040519081016040528092919081815260200182805461024c90610e58565b80156102995780601f1061026e57610100808354040283529160200191610299565b820191906000526020600020905b81548152906001019060200180831161027c57829003601f168201915b5050505050905090565b6000336102b18185856104b0565b60019150505b92915050565b6000336102cb8582856104c2565b6102d6858585610540565b506001949350505050565b60006102eb61059f565b905090565b6001600160a01b0381166000908152600760205260408120546102b7565b6000606080600080600060606103226106ca565b61032a6106f7565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606004805461022090610e58565b6000336102b1818585610540565b8342111561039a5760405163313c898160e11b8152600481018590526024015b60405180910390fd5b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886103e78c6001600160a01b0316600090815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e001604051602081830303815290604052805190602001209050600061044282610724565b9050600061045282878787610751565b9050896001600160a01b0316816001600160a01b031614610499576040516325c0072360e11b81526001600160a01b0380831660048301528b166024820152604401610391565b6104a48a8a8a6104b0565b50505050505050505050565b6104bd838383600161077f565b505050565b6001600160a01b03838116600090815260016020908152604080832093861683529290522054600019811461053a578181101561052b57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610391565b61053a8484848403600061077f565b50505050565b6001600160a01b03831661056a57604051634b637e8f60e11b815260006004820152602401610391565b6001600160a01b0382166105945760405163ec442f0560e01b815260006004820152602401610391565b6104bd838383610854565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156105f857507f000000000000000000000000000000000000000000000000000000000000000046145b1561062257507f000000000000000000000000000000000000000000000000000000000000000090565b6102eb604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60606102eb7f0000000000000000000000000000000000000000000000000000000000000000600561097e565b60606102eb7f0000000000000000000000000000000000000000000000000000000000000000600661097e565b60006102b761073161059f565b8360405161190160f01b8152600281019290925260228201526042902090565b60008060008061076388888888610a29565b9250925092506107738282610af8565b50909695505050505050565b6001600160a01b0384166107a95760405163e602df0560e01b815260006004820152602401610391565b6001600160a01b0383166107d357604051634a1406b160e11b815260006004820152602401610391565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561053a57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161084691815260200190565b60405180910390a350505050565b6001600160a01b03831661087f5780600260008282546108749190610e92565b909155506108f19050565b6001600160a01b038316600090815260208190526040902054818110156108d25760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610391565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b03821661090d5760028054829003905561092c565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161097191815260200190565b60405180910390a3505050565b606060ff83146109985761099183610bb5565b90506102b7565b8180546109a490610e58565b80601f01602080910402602001604051908101604052809291908181526020018280546109d090610e58565b8015610a1d5780601f106109f257610100808354040283529160200191610a1d565b820191906000526020600020905b815481529060010190602001808311610a0057829003601f168201915b505050505090506102b7565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610a645750600091506003905082610aee565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610ab8573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610ae457506000925060019150829050610aee565b9250600091508190505b9450945094915050565b6000826003811115610b0c57610b0c610eb3565b03610b15575050565b6001826003811115610b2957610b29610eb3565b03610b475760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610b5b57610b5b610eb3565b03610b7c5760405163fce698f760e01b815260048101829052602401610391565b6003826003811115610b9057610b90610eb3565b03610bb1576040516335e2f38360e21b815260048101829052602401610391565b5050565b60606000610bc283610bf4565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b600060ff8216601f8111156102b757604051632cd44ac360e21b815260040160405180910390fd5b6000815180845260005b81811015610c4257602081850181015186830182015201610c26565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610c756020830184610c1c565b9392505050565b80356001600160a01b0381168114610c9357600080fd5b919050565b60008060408385031215610cab57600080fd5b610cb483610c7c565b946020939093013593505050565b600080600060608486031215610cd757600080fd5b610ce084610c7c565b9250610cee60208501610c7c565b9150604084013590509250925092565b600060208284031215610d1057600080fd5b610c7582610c7c565b60ff60f81b881681526000602060e06020840152610d3a60e084018a610c1c565b8381036040850152610d4c818a610c1c565b606085018990526001600160a01b038816608086015260a0850187905284810360c08601528551808252602080880193509091019060005b81811015610da057835183529284019291840191600101610d84565b50909c9b505050505050505050505050565b600080600080600080600060e0888a031215610dcd57600080fd5b610dd688610c7c565b9650610de460208901610c7c565b95506040880135945060608801359350608088013560ff81168114610e0857600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215610e3857600080fd5b610e4183610c7c565b9150610e4f60208401610c7c565b90509250929050565b600181811c90821680610e6c57607f821691505b602082108103610e8c57634e487b7160e01b600052602260045260246000fd5b50919050565b808201808211156102b757634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fdfea26469706673582212203b21f1a4595c370666ac5aa5b70cc03a17018d72f3860c39361c097671ff189b64736f6c63430008160033",
}

// YellowTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use YellowTokenMetaData.ABI instead.
var YellowTokenABI = YellowTokenMetaData.ABI

// YellowTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YellowTokenMetaData.Bin instead.
var YellowTokenBin = YellowTokenMetaData.Bin

// DeployYellowToken deploys a new Ethereum contract, binding an instance of YellowToken to it.
func DeployYellowToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, supply *big.Int) (common.Address, *types.Transaction, *YellowToken, error) {
	parsed, err := YellowTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YellowTokenBin), backend, name, symbol, supply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &YellowToken{YellowTokenCaller: YellowTokenCaller{contract: contract}, YellowTokenTransactor: YellowTokenTransactor{contract: contract}, YellowTokenFilterer: YellowTokenFilterer{contract: contract}}, nil
}

// YellowToken is an auto generated Go binding around an Ethereum contract.
type YellowToken struct {
	YellowTokenCaller     // Read-only binding to the contract
	YellowTokenTransactor // Write-only binding to the contract
	YellowTokenFilterer   // Log filterer for contract events
}

// YellowTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type YellowTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YellowTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YellowTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YellowTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YellowTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YellowTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YellowTokenSession struct {
	Contract     *YellowToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YellowTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YellowTokenCallerSession struct {
	Contract *YellowTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YellowTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YellowTokenTransactorSession struct {
	Contract     *YellowTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YellowTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type YellowTokenRaw struct {
	Contract *YellowToken // Generic contract binding to access the raw methods on
}

// YellowTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YellowTokenCallerRaw struct {
	Contract *YellowTokenCaller // Generic read-only contract binding to access the raw methods on
}

// YellowTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YellowTokenTransactorRaw struct {
	Contract *YellowTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYellowToken creates a new instance of YellowToken, bound to a specific deployed contract.
func NewYellowToken(address common.Address, backend bind.ContractBackend) (*YellowToken, error) {
	contract, err := bindYellowToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YellowToken{YellowTokenCaller: YellowTokenCaller{contract: contract}, YellowTokenTransactor: YellowTokenTransactor{contract: contract}, YellowTokenFilterer: YellowTokenFilterer{contract: contract}}, nil
}

// NewYellowTokenCaller creates a new read-only instance of YellowToken, bound to a specific deployed contract.
func NewYellowTokenCaller(address common.Address, caller bind.ContractCaller) (*YellowTokenCaller, error) {
	contract, err := bindYellowToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YellowTokenCaller{contract: contract}, nil
}

// NewYellowTokenTransactor creates a new write-only instance of YellowToken, bound to a specific deployed contract.
func NewYellowTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*YellowTokenTransactor, error) {
	contract, err := bindYellowToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YellowTokenTransactor{contract: contract}, nil
}

// NewYellowTokenFilterer creates a new log filterer instance of YellowToken, bound to a specific deployed contract.
func NewYellowTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*YellowTokenFilterer, error) {
	contract, err := bindYellowToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YellowTokenFilterer{contract: contract}, nil
}

// bindYellowToken binds a generic wrapper to an already deployed contract.
func bindYellowToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YellowTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YellowToken *YellowTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YellowToken.Contract.YellowTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YellowToken *YellowTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YellowToken.Contract.YellowTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YellowToken *YellowTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YellowToken.Contract.YellowTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YellowToken *YellowTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YellowToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YellowToken *YellowTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YellowToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YellowToken *YellowTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YellowToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YellowToken *YellowTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YellowToken *YellowTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YellowToken.Contract.DOMAINSEPARATOR(&_YellowToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YellowToken *YellowTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YellowToken.Contract.DOMAINSEPARATOR(&_YellowToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YellowToken *YellowTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YellowToken *YellowTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YellowToken.Contract.Allowance(&_YellowToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YellowToken.Contract.Allowance(&_YellowToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YellowToken *YellowTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YellowToken *YellowTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YellowToken.Contract.BalanceOf(&_YellowToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YellowToken.Contract.BalanceOf(&_YellowToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_YellowToken *YellowTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_YellowToken *YellowTokenSession) Decimals() (uint8, error) {
	return _YellowToken.Contract.Decimals(&_YellowToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_YellowToken *YellowTokenCallerSession) Decimals() (uint8, error) {
	return _YellowToken.Contract.Decimals(&_YellowToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_YellowToken *YellowTokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_YellowToken *YellowTokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _YellowToken.Contract.Eip712Domain(&_YellowToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_YellowToken *YellowTokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _YellowToken.Contract.Eip712Domain(&_YellowToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YellowToken *YellowTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YellowToken *YellowTokenSession) Name() (string, error) {
	return _YellowToken.Contract.Name(&_YellowToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YellowToken *YellowTokenCallerSession) Name() (string, error) {
	return _YellowToken.Contract.Name(&_YellowToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_YellowToken *YellowTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_YellowToken *YellowTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _YellowToken.Contract.Nonces(&_YellowToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _YellowToken.Contract.Nonces(&_YellowToken.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YellowToken *YellowTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YellowToken *YellowTokenSession) Symbol() (string, error) {
	return _YellowToken.Contract.Symbol(&_YellowToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YellowToken *YellowTokenCallerSession) Symbol() (string, error) {
	return _YellowToken.Contract.Symbol(&_YellowToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YellowToken *YellowTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YellowToken *YellowTokenSession) TotalSupply() (*big.Int, error) {
	return _YellowToken.Contract.TotalSupply(&_YellowToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _YellowToken.Contract.TotalSupply(&_YellowToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_YellowToken *YellowTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_YellowToken *YellowTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Approve(&_YellowToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Approve(&_YellowToken.TransactOpts, spender, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_YellowToken *YellowTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_YellowToken *YellowTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YellowToken.Contract.Permit(&_YellowToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_YellowToken *YellowTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YellowToken.Contract.Permit(&_YellowToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_YellowToken *YellowTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_YellowToken *YellowTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Transfer(&_YellowToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Transfer(&_YellowToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_YellowToken *YellowTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_YellowToken *YellowTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.TransferFrom(&_YellowToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.TransferFrom(&_YellowToken.TransactOpts, from, to, value)
}

// YellowTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YellowToken contract.
type YellowTokenApprovalIterator struct {
	Event *YellowTokenApproval // Event containing the contract specifics and raw log

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
func (it *YellowTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenApproval)
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
		it.Event = new(YellowTokenApproval)
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
func (it *YellowTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenApproval represents a Approval event raised by the YellowToken contract.
type YellowTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YellowToken *YellowTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YellowTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenApprovalIterator{contract: _YellowToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YellowToken *YellowTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YellowTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenApproval)
				if err := _YellowToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YellowToken *YellowTokenFilterer) ParseApproval(log types.Log) (*YellowTokenApproval, error) {
	event := new(YellowTokenApproval)
	if err := _YellowToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the YellowToken contract.
type YellowTokenEIP712DomainChangedIterator struct {
	Event *YellowTokenEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *YellowTokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenEIP712DomainChanged)
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
		it.Event = new(YellowTokenEIP712DomainChanged)
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
func (it *YellowTokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the YellowToken contract.
type YellowTokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_YellowToken *YellowTokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*YellowTokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &YellowTokenEIP712DomainChangedIterator{contract: _YellowToken.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_YellowToken *YellowTokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *YellowTokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenEIP712DomainChanged)
				if err := _YellowToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_YellowToken *YellowTokenFilterer) ParseEIP712DomainChanged(log types.Log) (*YellowTokenEIP712DomainChanged, error) {
	event := new(YellowTokenEIP712DomainChanged)
	if err := _YellowToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YellowToken contract.
type YellowTokenTransferIterator struct {
	Event *YellowTokenTransfer // Event containing the contract specifics and raw log

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
func (it *YellowTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenTransfer)
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
		it.Event = new(YellowTokenTransfer)
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
func (it *YellowTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenTransfer represents a Transfer event raised by the YellowToken contract.
type YellowTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YellowToken *YellowTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YellowTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenTransferIterator{contract: _YellowToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YellowToken *YellowTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YellowTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenTransfer)
				if err := _YellowToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YellowToken *YellowTokenFilterer) ParseTransfer(log types.Log) (*YellowTokenTransfer, error) {
	event := new(YellowTokenTransfer)
	if err := _YellowToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
