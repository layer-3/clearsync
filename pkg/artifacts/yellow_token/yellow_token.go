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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b50604051620016eb380380620016eb833981016040819052620000359162000355565b6040805180820190915260018152603160f81b6020820152839081908185600362000061838262000456565b50600462000070828262000456565b50620000829150839050600562000140565b610120526200009381600662000140565b61014052815160208084019190912060e052815190820120610100524660a0526200012160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525062000137338262000179565b5050506200059e565b60006020835110156200016057620001588362000240565b905062000173565b816200016d848262000456565b5060ff90505b92915050565b6001600160a01b038216620001d55760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064015b60405180910390fd5b8060026000828254620001e9919062000522565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600080829050601f815111156200026e578260405163305a27a960e01b8152600401620001cc919062000544565b80516200027b8262000579565b179392505050565b505050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620002bb578181015183820152602001620002a1565b50506000910152565b600082601f830112620002d657600080fd5b81516001600160401b0380821115620002f357620002f362000288565b604051601f8301601f19908116603f011681019082821181831017156200031e576200031e62000288565b816040528381528660208588010111156200033857600080fd5b6200034b8460208301602089016200029e565b9695505050505050565b6000806000606084860312156200036b57600080fd5b83516001600160401b03808211156200038357600080fd5b6200039187838801620002c4565b94506020860151915080821115620003a857600080fd5b50620003b786828701620002c4565b925050604084015190509250925092565b600181811c90821680620003dd57607f821691505b602082108103620003fe57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028357600081815260208120601f850160051c810160208610156200042d5750805b601f850160051c820191505b818110156200044e5782815560010162000439565b505050505050565b81516001600160401b0381111562000472576200047262000288565b6200048a81620004838454620003c8565b8462000404565b602080601f831160018114620004c25760008415620004a95750858301515b600019600386901b1c1916600185901b1785556200044e565b600085815260208120601f198616915b82811015620004f357888601518255948401946001909101908401620004d2565b5085821015620005125787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200017357634e487b7160e01b600052601160045260246000fd5b6020815260008251806020840152620005658160408501602087016200029e565b601f01601f19169190910160400192915050565b80516020808301519190811015620003fe5760001960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516110f2620005f9600039600061038f0152600061036401526000610a27015260006109ff0152600061095a01526000610984015260006109ae01526110f26000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d7146101e9578063a9059cbb146101fc578063d505accf1461020f578063dd62ed3e1461022457600080fd5b806370a082311461018a5780637ecebe00146101b357806384b0196e146101c657806395d89b41146101e157600080fd5b806323b872dd116100d357806323b872dd1461014d578063313ce567146101605780633644e5151461016f578063395093511461017757600080fd5b806306fdde03146100fa578063095ea7b31461011857806318160ddd1461013b575b600080fd5b610102610237565b60405161010f9190610e5e565b60405180910390f35b61012b610126366004610e94565b6102c9565b604051901515815260200161010f565b6002545b60405190815260200161010f565b61012b61015b366004610ebe565b6102e3565b6040516008815260200161010f565b61013f610307565b61012b610185366004610e94565b610316565b61013f610198366004610efa565b6001600160a01b031660009081526020819052604090205490565b61013f6101c1366004610efa565b610338565b6101ce610356565b60405161010f9796959493929190610f15565b6101026103df565b61012b6101f7366004610e94565b6103ee565b61012b61020a366004610e94565b61046e565b61022261021d366004610fab565b61047c565b005b61013f61023236600461101e565b6105e0565b60606003805461024690611051565b80601f016020809104026020016040519081016040528092919081815260200182805461027290611051565b80156102bf5780601f10610294576101008083540402835291602001916102bf565b820191906000526020600020905b8154815290600101906020018083116102a257829003601f168201915b5050505050905090565b6000336102d781858561060b565b60019150505b92915050565b6000336102f185828561072f565b6102fc8585856107a9565b506001949350505050565b600061031161094d565b905090565b6000336102d781858561032983836105e0565b6103339190611085565b61060b565b6001600160a01b0381166000908152600760205260408120546102dd565b60006060808280808361038a7f00000000000000000000000000000000000000000000000000000000000000006005610a78565b6103b57f00000000000000000000000000000000000000000000000000000000000000006006610a78565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606004805461024690611051565b600033816103fc82866105e0565b9050838110156104615760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102fc828686840361060b565b6000336102d78185856107a9565b834211156104cc5760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e650000006044820152606401610458565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886104fb8c610b23565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e001604051602081830303815290604052805190602001209050600061055682610b4b565b9050600061056682878787610b78565b9050896001600160a01b0316816001600160a01b0316146105c95760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e617475726500006044820152606401610458565b6105d48a8a8a61060b565b50505050505050505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b03831661066d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610458565b6001600160a01b0382166106ce5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610458565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b600061073b84846105e0565b905060001981146107a357818110156107965760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610458565b6107a3848484840361060b565b50505050565b6001600160a01b03831661080d5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610458565b6001600160a01b03821661086f5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610458565b6001600160a01b038316600090815260208190526040902054818110156108e75760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610458565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36107a3565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156109a657507f000000000000000000000000000000000000000000000000000000000000000046145b156109d057507f000000000000000000000000000000000000000000000000000000000000000090565b610311604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b606060ff8314610a9257610a8b83610ba0565b90506102dd565b818054610a9e90611051565b80601f0160208091040260200160405190810160405280929190818152602001828054610aca90611051565b8015610b175780601f10610aec57610100808354040283529160200191610b17565b820191906000526020600020905b815481529060010190602001808311610afa57829003601f168201915b505050505090506102dd565b6001600160a01b03811660009081526007602052604090208054600181018255905b50919050565b60006102dd610b5861094d565b8360405161190160f01b8152600281019290925260228201526042902090565b6000806000610b8987878787610bdf565b91509150610b9681610ca3565b5095945050505050565b60606000610bad83610df0565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610c165750600090506003610c9a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610c6a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610c9357600060019250925050610c9a565b9150600090505b94509492505050565b6000816004811115610cb757610cb76110a6565b03610cbf5750565b6001816004811115610cd357610cd36110a6565b03610d205760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610458565b6002816004811115610d3457610d346110a6565b03610d815760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610458565b6003816004811115610d9557610d956110a6565b03610ded5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610458565b50565b600060ff8216601f8111156102dd57604051632cd44ac360e21b815260040160405180910390fd5b6000815180845260005b81811015610e3e57602081850181015186830182015201610e22565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610e716020830184610e18565b9392505050565b80356001600160a01b0381168114610e8f57600080fd5b919050565b60008060408385031215610ea757600080fd5b610eb083610e78565b946020939093013593505050565b600080600060608486031215610ed357600080fd5b610edc84610e78565b9250610eea60208501610e78565b9150604084013590509250925092565b600060208284031215610f0c57600080fd5b610e7182610e78565b60ff60f81b881681526000602060e081840152610f3560e084018a610e18565b8381036040850152610f47818a610e18565b606085018990526001600160a01b038816608086015260a0850187905284810360c0860152855180825283870192509083019060005b81811015610f9957835183529284019291840191600101610f7d565b50909c9b505050505050505050505050565b600080600080600080600060e0888a031215610fc657600080fd5b610fcf88610e78565b9650610fdd60208901610e78565b95506040880135945060608801359350608088013560ff8116811461100157600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561103157600080fd5b61103a83610e78565b915061104860208401610e78565b90509250929050565b600181811c9082168061106557607f821691505b602082108103610b4557634e487b7160e01b600052602260045260246000fd5b808201808211156102dd57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fdfea264697066735822122069ca53e4b2f1f37d26273d4eaffc0b440f73fbb5590ac4c825bafe6debf4408d64736f6c63430008140033",
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Approve(&_YellowToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Approve(&_YellowToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YellowToken *YellowTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YellowToken *YellowTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.DecreaseAllowance(&_YellowToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.DecreaseAllowance(&_YellowToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YellowToken *YellowTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YellowToken *YellowTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.IncreaseAllowance(&_YellowToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.IncreaseAllowance(&_YellowToken.TransactOpts, spender, addedValue)
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
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Transfer(&_YellowToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Transfer(&_YellowToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.TransferFrom(&_YellowToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YellowToken *YellowTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.TransferFrom(&_YellowToken.TransactOpts, from, to, amount)
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
