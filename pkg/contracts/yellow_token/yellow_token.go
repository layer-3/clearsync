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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premint\",\"type\":\"uint256\"}],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BlacklistedBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BlacklistedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLACKLISTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMPLIANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_SUPPLY_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"premint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"burnBlacklisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBlacklisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001b8338038062001b83833981016040819052620000349162000207565b8282600362000044838262000309565b50600462000053828262000309565b5062000065915060009050336200009d565b620000917f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6336200009d565b60805250620003d59050565b60008281526005602090815260408083206001600160a01b038516845290915290205460ff166200013e5760008281526005602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620000fd3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200016a57600080fd5b81516001600160401b038082111562000187576200018762000142565b604051601f8301601f19908116603f01168101908282118183101715620001b257620001b262000142565b81604052838152602092508683858801011115620001cf57600080fd5b600091505b83821015620001f35785820183015181830184015290820190620001d4565b600093810190920192909252949350505050565b6000806000606084860312156200021d57600080fd5b83516001600160401b03808211156200023557600080fd5b620002438783880162000158565b945060208601519150808211156200025a57600080fd5b50620002698682870162000158565b925050604084015190509250925092565b600181811c908216806200028f57607f821691505b602082108103620002b057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200030457600081815260208120601f850160051c81016020861015620002df5750805b601f850160051c820191505b818110156200030057828155600101620002eb565b5050505b505050565b81516001600160401b0381111562000325576200032562000142565b6200033d816200033684546200027a565b84620002b6565b602080601f8311600181146200037557600084156200035c5750858301515b600019600386901b1c1916600185901b17855562000300565b600085815260208120601f198616915b82811015620003a65788860151825594840194600190910190840162000385565b5085821015620003c55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805161177d62000406600039600081816102d301528181610424015281816106e201526107b9015261177d6000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806362b199c51161010f578063a457c2d7116100a2578063d539139311610071578063d539139314610446578063d547741f1461046d578063dd62ed3e14610480578063f9f92be41461049357600080fd5b8063a457c2d7146103e6578063a9059cbb146103f9578063c6a276c21461040c578063cea4c3aa1461041f57600080fd5b806379cc6790116100de57806379cc6790146103b057806391d14854146103c357806395d89b41146103d6578063a217fddf146103de57600080fd5b806362b199c51461034c578063637c25a11461036157806370a082311461037457806378d06fce1461039d57600080fd5b80632f2ff15d116101875780633950935111610156578063395093511461030a57806340c10f191461031d57806342966c6814610330578063597be6d11461034357600080fd5b80632f2ff15d146102ad578063313ce567146102c2578063355274ea146102d157806336568abe146102f757600080fd5b8063095ea7b3116101c3578063095ea7b31461025c57806318160ddd1461026f57806323b872dd14610277578063248a9ca31461028a57600080fd5b806301ffc9a7146101ea578063062d3bd71461021257806306fdde0314610247575b600080fd5b6101fd6101f8366004611468565b6104a6565b60405190151581526020015b60405180910390f35b6102397f442a94f1a1fac79af32856af2a64f63648cfa2ef3b98610a5bb7cbec4cee698581565b604051908152602001610209565b61024f6104dd565b60405161020991906114b6565b6101fd61026a366004611505565b61056f565b600254610239565b6101fd61028536600461152f565b610587565b61023961029836600461156b565b60009081526005602052604090206001015490565b6102c06102bb366004611584565b6105a5565b005b60405160088152602001610209565b7f0000000000000000000000000000000000000000000000000000000000000000610239565b6102c0610305366004611584565b6105cf565b6101fd610318366004611505565b610652565b6102c061032b366004611505565b610674565b6102c061033e36600461156b565b610760565b61023960065481565b61023960008051602061172883398151915281565b6102c061036f366004611584565b61076d565b6102396103823660046115b0565b6001600160a01b031660009081526020819052604090205490565b6102c06103ab3660046115b0565b6108a7565b6102c06103be366004611505565b610981565b6101fd6103d1366004611584565b610996565b61024f6109c1565b610239600081565b6101fd6103f4366004611505565b6109d0565b6101fd610407366004611505565b610a56565b6102c061041a3660046115b0565b610a72565b6102397f000000000000000000000000000000000000000000000000000000000000000081565b6102397f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6102c061047b366004611584565b610aec565b61023961048e3660046115cb565b610b11565b6102c06104a13660046115b0565b610b3c565b60006001600160e01b03198216637965db0b60e01b14806104d757506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600380546104ec906115f5565b80601f0160208091040260200160405190810160405280929190818152602001828054610518906115f5565b80156105655780601f1061053a57610100808354040283529160200191610565565b820191906000526020600020905b81548152906001019060200180831161054857829003601f168201915b5050505050905090565b60003361057d818585610bb6565b5060019392505050565b600061059284610cda565b61059d848484610d38565b949350505050565b6000828152600560205260409020600101546105c081610d51565b6105ca8383610d5b565b505050565b6001600160a01b03811633146106445760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b61064e8282610de1565b5050565b60003361057d8185856106658383610b11565b61066f9190611645565b610bb6565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661069e81610d51565b6000600654116106e05760405162461bcd60e51b815260206004820152600d60248201526c139bdd081858dd1a5d985d1959609a1b604482015260640161063b565b7f00000000000000000000000000000000000000000000000000000000000000008261070b60025490565b6107159190611645565b11156107565760405162461bcd60e51b815260206004820152601060248201526f04d696e742065786365656473206361760841b604482015260640161063b565b6105ca8383610e48565b61076a3382610f07565b50565b600061077881610d51565b600083116107b75760405162461bcd60e51b815260206004820152600c60248201526b16995c9bc81c1c995b5a5b9d60a21b604482015260640161063b565b7f000000000000000000000000000000000000000000000000000000000000000083111561081d5760405162461bcd60e51b815260206004820152601360248201527205072656d696e7420657863656564732063617606c1b604482015260640161063b565b600654156108615760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e481858dd1a5d985d1959607a1b604482015260640161063b565b4260065561086f8284610e48565b6040518381527f3ec796be1be7d03bff3a62b9fa594a60e947c1809bced06d929f145308ae57ce9060200160405180910390a1505050565b60006108b281610d51565b6108ca60008051602061172883398151915283610996565b6109165760405162461bcd60e51b815260206004820152601a60248201527f4163636f756e74206973206e6f7420626c61636b6c6973746564000000000000604482015260640161063b565b6001600160a01b0382166000908152602081905260409020546109398382610f07565b826001600160a01b03167f4743ca7a69ca9ba3afc645605fdc654fc3cfcf7791e24505160b617a5e5262e88260405161097491815260200190565b60405180910390a2505050565b61098c823383611039565b61064e8282610f07565b60009182526005602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060600480546104ec906115f5565b600033816109de8286610b11565b905083811015610a3e5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161063b565b610a4b8286868403610bb6565b506001949350505050565b6000610a6133610cda565b610a6b83836110b3565b9392505050565b7f442a94f1a1fac79af32856af2a64f63648cfa2ef3b98610a5bb7cbec4cee6985610a9c81610d51565b610ab460008051602061172883398151915283610de1565b6040516001600160a01b038316907ff38e60871ec534937251cd91cad807e15f55f1f6815128faecc256e71994b49790600090a25050565b600082815260056020526040902060010154610b0781610d51565b6105ca8383610de1565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b7f442a94f1a1fac79af32856af2a64f63648cfa2ef3b98610a5bb7cbec4cee6985610b6681610d51565b610b7e60008051602061172883398151915283610d5b565b6040516001600160a01b038316907fffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b85590600090a25050565b6001600160a01b038316610c185760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161063b565b6001600160a01b038216610c795760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161063b565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b610cf260008051602061172883398151915282610996565b1561076a5760405162461bcd60e51b81526020600482015260166024820152751058d8dbdd5b9d081a5cc8189b1858dadb1a5cdd195960521b604482015260640161063b565b600033610d46858285611039565b610a4b8585856110bd565b61076a8133611261565b610d658282610996565b61064e5760008281526005602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610d9d3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610deb8282610996565b1561064e5760008281526005602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6001600160a01b038216610e9e5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161063b565b8060026000828254610eb09190611645565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610f675760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161063b565b6001600160a01b03821660009081526020819052604090205481811015610fdb5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161063b565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b60006110458484610b11565b905060001981146110ad57818110156110a05760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161063b565b6110ad8484848403610bb6565b50505050565b60003361057d8185855b6001600160a01b0383166111215760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161063b565b6001600160a01b0382166111835760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161063b565b6001600160a01b038316600090815260208190526040902054818110156111fb5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161063b565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36110ad565b61126b8282610996565b61064e57611278816112ba565b6112838360206112cc565b604051602001611294929190611658565b60408051601f198184030181529082905262461bcd60e51b825261063b916004016114b6565b60606104d76001600160a01b03831660145b606060006112db8360026116cd565b6112e6906002611645565b67ffffffffffffffff8111156112fe576112fe6116e4565b6040519080825280601f01601f191660200182016040528015611328576020820181803683370190505b509050600360fc1b81600081518110611343576113436116fa565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611372576113726116fa565b60200101906001600160f81b031916908160001a90535060006113968460026116cd565b6113a1906001611645565b90505b6001811115611419576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106113d5576113d56116fa565b1a60f81b8282815181106113eb576113eb6116fa565b60200101906001600160f81b031916908160001a90535060049490941c9361141281611710565b90506113a4565b508315610a6b5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161063b565b60006020828403121561147a57600080fd5b81356001600160e01b031981168114610a6b57600080fd5b60005b838110156114ad578181015183820152602001611495565b50506000910152565b60208152600082518060208401526114d5816040850160208701611492565b601f01601f19169190910160400192915050565b80356001600160a01b038116811461150057600080fd5b919050565b6000806040838503121561151857600080fd5b611521836114e9565b946020939093013593505050565b60008060006060848603121561154457600080fd5b61154d846114e9565b925061155b602085016114e9565b9150604084013590509250925092565b60006020828403121561157d57600080fd5b5035919050565b6000806040838503121561159757600080fd5b823591506115a7602084016114e9565b90509250929050565b6000602082840312156115c257600080fd5b610a6b826114e9565b600080604083850312156115de57600080fd5b6115e7836114e9565b91506115a7602084016114e9565b600181811c9082168061160957607f821691505b60208210810361162957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b808201808211156104d7576104d761162f565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611690816017850160208801611492565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516116c1816028840160208801611492565b01602801949350505050565b80820281158282048414176104d7576104d761162f565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60008161171f5761171f61162f565b50600019019056fe548c7f0307ab2a7ea894e5c7e8c5353cc750bb9385ee2e945f189a9a83daa8eda264697066735822122039848b58acfc3847a0fc9e3477876c0743b5fa9b076baa24f00331bdb239f66364736f6c63430008120033",
}

// YellowTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use YellowTokenMetaData.ABI instead.
var YellowTokenABI = YellowTokenMetaData.ABI

// YellowTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YellowTokenMetaData.Bin instead.
var YellowTokenBin = YellowTokenMetaData.Bin

// DeployYellowToken deploys a new Ethereum contract, binding an instance of YellowToken to it.
func DeployYellowToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, supplyCap *big.Int) (common.Address, *types.Transaction, *YellowToken, error) {
	parsed, err := YellowTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YellowTokenBin), backend, name, symbol, supplyCap)
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

// BLACKLISTEDROLE is a free data retrieval call binding the contract method 0x62b199c5.
//
// Solidity: function BLACKLISTED_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCaller) BLACKLISTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "BLACKLISTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLACKLISTEDROLE is a free data retrieval call binding the contract method 0x62b199c5.
//
// Solidity: function BLACKLISTED_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenSession) BLACKLISTEDROLE() ([32]byte, error) {
	return _YellowToken.Contract.BLACKLISTEDROLE(&_YellowToken.CallOpts)
}

// BLACKLISTEDROLE is a free data retrieval call binding the contract method 0x62b199c5.
//
// Solidity: function BLACKLISTED_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCallerSession) BLACKLISTEDROLE() ([32]byte, error) {
	return _YellowToken.Contract.BLACKLISTEDROLE(&_YellowToken.CallOpts)
}

// COMPLIANCEROLE is a free data retrieval call binding the contract method 0x062d3bd7.
//
// Solidity: function COMPLIANCE_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCaller) COMPLIANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "COMPLIANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// COMPLIANCEROLE is a free data retrieval call binding the contract method 0x062d3bd7.
//
// Solidity: function COMPLIANCE_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenSession) COMPLIANCEROLE() ([32]byte, error) {
	return _YellowToken.Contract.COMPLIANCEROLE(&_YellowToken.CallOpts)
}

// COMPLIANCEROLE is a free data retrieval call binding the contract method 0x062d3bd7.
//
// Solidity: function COMPLIANCE_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCallerSession) COMPLIANCEROLE() ([32]byte, error) {
	return _YellowToken.Contract.COMPLIANCEROLE(&_YellowToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _YellowToken.Contract.DEFAULTADMINROLE(&_YellowToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _YellowToken.Contract.DEFAULTADMINROLE(&_YellowToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenSession) MINTERROLE() ([32]byte, error) {
	return _YellowToken.Contract.MINTERROLE(&_YellowToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_YellowToken *YellowTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _YellowToken.Contract.MINTERROLE(&_YellowToken.CallOpts)
}

// TOKENSUPPLYCAP is a free data retrieval call binding the contract method 0xcea4c3aa.
//
// Solidity: function TOKEN_SUPPLY_CAP() view returns(uint256)
func (_YellowToken *YellowTokenCaller) TOKENSUPPLYCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "TOKEN_SUPPLY_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENSUPPLYCAP is a free data retrieval call binding the contract method 0xcea4c3aa.
//
// Solidity: function TOKEN_SUPPLY_CAP() view returns(uint256)
func (_YellowToken *YellowTokenSession) TOKENSUPPLYCAP() (*big.Int, error) {
	return _YellowToken.Contract.TOKENSUPPLYCAP(&_YellowToken.CallOpts)
}

// TOKENSUPPLYCAP is a free data retrieval call binding the contract method 0xcea4c3aa.
//
// Solidity: function TOKEN_SUPPLY_CAP() view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) TOKENSUPPLYCAP() (*big.Int, error) {
	return _YellowToken.Contract.TOKENSUPPLYCAP(&_YellowToken.CallOpts)
}

// ActivatedAt is a free data retrieval call binding the contract method 0x597be6d1.
//
// Solidity: function activatedAt() view returns(uint256)
func (_YellowToken *YellowTokenCaller) ActivatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "activatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivatedAt is a free data retrieval call binding the contract method 0x597be6d1.
//
// Solidity: function activatedAt() view returns(uint256)
func (_YellowToken *YellowTokenSession) ActivatedAt() (*big.Int, error) {
	return _YellowToken.Contract.ActivatedAt(&_YellowToken.CallOpts)
}

// ActivatedAt is a free data retrieval call binding the contract method 0x597be6d1.
//
// Solidity: function activatedAt() view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) ActivatedAt() (*big.Int, error) {
	return _YellowToken.Contract.ActivatedAt(&_YellowToken.CallOpts)
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

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_YellowToken *YellowTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_YellowToken *YellowTokenSession) Cap() (*big.Int, error) {
	return _YellowToken.Contract.Cap(&_YellowToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_YellowToken *YellowTokenCallerSession) Cap() (*big.Int, error) {
	return _YellowToken.Contract.Cap(&_YellowToken.CallOpts)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_YellowToken *YellowTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_YellowToken *YellowTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _YellowToken.Contract.GetRoleAdmin(&_YellowToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_YellowToken *YellowTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _YellowToken.Contract.GetRoleAdmin(&_YellowToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_YellowToken *YellowTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_YellowToken *YellowTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _YellowToken.Contract.HasRole(&_YellowToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_YellowToken *YellowTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _YellowToken.Contract.HasRole(&_YellowToken.CallOpts, role, account)
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YellowToken *YellowTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _YellowToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YellowToken *YellowTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YellowToken.Contract.SupportsInterface(&_YellowToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YellowToken *YellowTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YellowToken.Contract.SupportsInterface(&_YellowToken.CallOpts, interfaceId)
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

// Activate is a paid mutator transaction binding the contract method 0x637c25a1.
//
// Solidity: function activate(uint256 premint, address account) returns()
func (_YellowToken *YellowTokenTransactor) Activate(opts *bind.TransactOpts, premint *big.Int, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "activate", premint, account)
}

// Activate is a paid mutator transaction binding the contract method 0x637c25a1.
//
// Solidity: function activate(uint256 premint, address account) returns()
func (_YellowToken *YellowTokenSession) Activate(premint *big.Int, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.Activate(&_YellowToken.TransactOpts, premint, account)
}

// Activate is a paid mutator transaction binding the contract method 0x637c25a1.
//
// Solidity: function activate(uint256 premint, address account) returns()
func (_YellowToken *YellowTokenTransactorSession) Activate(premint *big.Int, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.Activate(&_YellowToken.TransactOpts, premint, account)
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

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address account) returns()
func (_YellowToken *YellowTokenTransactor) Blacklist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "blacklist", account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address account) returns()
func (_YellowToken *YellowTokenSession) Blacklist(account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.Blacklist(&_YellowToken.TransactOpts, account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address account) returns()
func (_YellowToken *YellowTokenTransactorSession) Blacklist(account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.Blacklist(&_YellowToken.TransactOpts, account)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_YellowToken *YellowTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_YellowToken *YellowTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Burn(&_YellowToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_YellowToken *YellowTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Burn(&_YellowToken.TransactOpts, amount)
}

// BurnBlacklisted is a paid mutator transaction binding the contract method 0x78d06fce.
//
// Solidity: function burnBlacklisted(address account) returns()
func (_YellowToken *YellowTokenTransactor) BurnBlacklisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "burnBlacklisted", account)
}

// BurnBlacklisted is a paid mutator transaction binding the contract method 0x78d06fce.
//
// Solidity: function burnBlacklisted(address account) returns()
func (_YellowToken *YellowTokenSession) BurnBlacklisted(account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.BurnBlacklisted(&_YellowToken.TransactOpts, account)
}

// BurnBlacklisted is a paid mutator transaction binding the contract method 0x78d06fce.
//
// Solidity: function burnBlacklisted(address account) returns()
func (_YellowToken *YellowTokenTransactorSession) BurnBlacklisted(account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.BurnBlacklisted(&_YellowToken.TransactOpts, account)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_YellowToken *YellowTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_YellowToken *YellowTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.BurnFrom(&_YellowToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_YellowToken *YellowTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.BurnFrom(&_YellowToken.TransactOpts, account, amount)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.GrantRole(&_YellowToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.GrantRole(&_YellowToken.TransactOpts, role, account)
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

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_YellowToken *YellowTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_YellowToken *YellowTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Mint(&_YellowToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_YellowToken *YellowTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YellowToken.Contract.Mint(&_YellowToken.TransactOpts, to, amount)
}

// RemoveBlacklisted is a paid mutator transaction binding the contract method 0xc6a276c2.
//
// Solidity: function removeBlacklisted(address account) returns()
func (_YellowToken *YellowTokenTransactor) RemoveBlacklisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "removeBlacklisted", account)
}

// RemoveBlacklisted is a paid mutator transaction binding the contract method 0xc6a276c2.
//
// Solidity: function removeBlacklisted(address account) returns()
func (_YellowToken *YellowTokenSession) RemoveBlacklisted(account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.RemoveBlacklisted(&_YellowToken.TransactOpts, account)
}

// RemoveBlacklisted is a paid mutator transaction binding the contract method 0xc6a276c2.
//
// Solidity: function removeBlacklisted(address account) returns()
func (_YellowToken *YellowTokenTransactorSession) RemoveBlacklisted(account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.RemoveBlacklisted(&_YellowToken.TransactOpts, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.RenounceRole(&_YellowToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.RenounceRole(&_YellowToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.RevokeRole(&_YellowToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_YellowToken *YellowTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YellowToken.Contract.RevokeRole(&_YellowToken.TransactOpts, role, account)
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

// YellowTokenActivatedIterator is returned from FilterActivated and is used to iterate over the raw logs and unpacked data for Activated events raised by the YellowToken contract.
type YellowTokenActivatedIterator struct {
	Event *YellowTokenActivated // Event containing the contract specifics and raw log

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
func (it *YellowTokenActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenActivated)
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
		it.Event = new(YellowTokenActivated)
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
func (it *YellowTokenActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenActivated represents a Activated event raised by the YellowToken contract.
type YellowTokenActivated struct {
	Premint *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActivated is a free log retrieval operation binding the contract event 0x3ec796be1be7d03bff3a62b9fa594a60e947c1809bced06d929f145308ae57ce.
//
// Solidity: event Activated(uint256 premint)
func (_YellowToken *YellowTokenFilterer) FilterActivated(opts *bind.FilterOpts) (*YellowTokenActivatedIterator, error) {

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "Activated")
	if err != nil {
		return nil, err
	}
	return &YellowTokenActivatedIterator{contract: _YellowToken.contract, event: "Activated", logs: logs, sub: sub}, nil
}

// WatchActivated is a free log subscription operation binding the contract event 0x3ec796be1be7d03bff3a62b9fa594a60e947c1809bced06d929f145308ae57ce.
//
// Solidity: event Activated(uint256 premint)
func (_YellowToken *YellowTokenFilterer) WatchActivated(opts *bind.WatchOpts, sink chan<- *YellowTokenActivated) (event.Subscription, error) {

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "Activated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenActivated)
				if err := _YellowToken.contract.UnpackLog(event, "Activated", log); err != nil {
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

// ParseActivated is a log parse operation binding the contract event 0x3ec796be1be7d03bff3a62b9fa594a60e947c1809bced06d929f145308ae57ce.
//
// Solidity: event Activated(uint256 premint)
func (_YellowToken *YellowTokenFilterer) ParseActivated(log types.Log) (*YellowTokenActivated, error) {
	event := new(YellowTokenActivated)
	if err := _YellowToken.contract.UnpackLog(event, "Activated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// YellowTokenBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the YellowToken contract.
type YellowTokenBlacklistedIterator struct {
	Event *YellowTokenBlacklisted // Event containing the contract specifics and raw log

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
func (it *YellowTokenBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenBlacklisted)
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
		it.Event = new(YellowTokenBlacklisted)
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
func (it *YellowTokenBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenBlacklisted represents a Blacklisted event raised by the YellowToken contract.
type YellowTokenBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed account)
func (_YellowToken *YellowTokenFilterer) FilterBlacklisted(opts *bind.FilterOpts, account []common.Address) (*YellowTokenBlacklistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "Blacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenBlacklistedIterator{contract: _YellowToken.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed account)
func (_YellowToken *YellowTokenFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *YellowTokenBlacklisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "Blacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenBlacklisted)
				if err := _YellowToken.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed account)
func (_YellowToken *YellowTokenFilterer) ParseBlacklisted(log types.Log) (*YellowTokenBlacklisted, error) {
	event := new(YellowTokenBlacklisted)
	if err := _YellowToken.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenBlacklistedBurntIterator is returned from FilterBlacklistedBurnt and is used to iterate over the raw logs and unpacked data for BlacklistedBurnt events raised by the YellowToken contract.
type YellowTokenBlacklistedBurntIterator struct {
	Event *YellowTokenBlacklistedBurnt // Event containing the contract specifics and raw log

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
func (it *YellowTokenBlacklistedBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenBlacklistedBurnt)
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
		it.Event = new(YellowTokenBlacklistedBurnt)
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
func (it *YellowTokenBlacklistedBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenBlacklistedBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenBlacklistedBurnt represents a BlacklistedBurnt event raised by the YellowToken contract.
type YellowTokenBlacklistedBurnt struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklistedBurnt is a free log retrieval operation binding the contract event 0x4743ca7a69ca9ba3afc645605fdc654fc3cfcf7791e24505160b617a5e5262e8.
//
// Solidity: event BlacklistedBurnt(address indexed account, uint256 amount)
func (_YellowToken *YellowTokenFilterer) FilterBlacklistedBurnt(opts *bind.FilterOpts, account []common.Address) (*YellowTokenBlacklistedBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "BlacklistedBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenBlacklistedBurntIterator{contract: _YellowToken.contract, event: "BlacklistedBurnt", logs: logs, sub: sub}, nil
}

// WatchBlacklistedBurnt is a free log subscription operation binding the contract event 0x4743ca7a69ca9ba3afc645605fdc654fc3cfcf7791e24505160b617a5e5262e8.
//
// Solidity: event BlacklistedBurnt(address indexed account, uint256 amount)
func (_YellowToken *YellowTokenFilterer) WatchBlacklistedBurnt(opts *bind.WatchOpts, sink chan<- *YellowTokenBlacklistedBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "BlacklistedBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenBlacklistedBurnt)
				if err := _YellowToken.contract.UnpackLog(event, "BlacklistedBurnt", log); err != nil {
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

// ParseBlacklistedBurnt is a log parse operation binding the contract event 0x4743ca7a69ca9ba3afc645605fdc654fc3cfcf7791e24505160b617a5e5262e8.
//
// Solidity: event BlacklistedBurnt(address indexed account, uint256 amount)
func (_YellowToken *YellowTokenFilterer) ParseBlacklistedBurnt(log types.Log) (*YellowTokenBlacklistedBurnt, error) {
	event := new(YellowTokenBlacklistedBurnt)
	if err := _YellowToken.contract.UnpackLog(event, "BlacklistedBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenBlacklistedRemovedIterator is returned from FilterBlacklistedRemoved and is used to iterate over the raw logs and unpacked data for BlacklistedRemoved events raised by the YellowToken contract.
type YellowTokenBlacklistedRemovedIterator struct {
	Event *YellowTokenBlacklistedRemoved // Event containing the contract specifics and raw log

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
func (it *YellowTokenBlacklistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenBlacklistedRemoved)
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
		it.Event = new(YellowTokenBlacklistedRemoved)
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
func (it *YellowTokenBlacklistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenBlacklistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenBlacklistedRemoved represents a BlacklistedRemoved event raised by the YellowToken contract.
type YellowTokenBlacklistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklistedRemoved is a free log retrieval operation binding the contract event 0xf38e60871ec534937251cd91cad807e15f55f1f6815128faecc256e71994b497.
//
// Solidity: event BlacklistedRemoved(address indexed account)
func (_YellowToken *YellowTokenFilterer) FilterBlacklistedRemoved(opts *bind.FilterOpts, account []common.Address) (*YellowTokenBlacklistedRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "BlacklistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenBlacklistedRemovedIterator{contract: _YellowToken.contract, event: "BlacklistedRemoved", logs: logs, sub: sub}, nil
}

// WatchBlacklistedRemoved is a free log subscription operation binding the contract event 0xf38e60871ec534937251cd91cad807e15f55f1f6815128faecc256e71994b497.
//
// Solidity: event BlacklistedRemoved(address indexed account)
func (_YellowToken *YellowTokenFilterer) WatchBlacklistedRemoved(opts *bind.WatchOpts, sink chan<- *YellowTokenBlacklistedRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "BlacklistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenBlacklistedRemoved)
				if err := _YellowToken.contract.UnpackLog(event, "BlacklistedRemoved", log); err != nil {
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

// ParseBlacklistedRemoved is a log parse operation binding the contract event 0xf38e60871ec534937251cd91cad807e15f55f1f6815128faecc256e71994b497.
//
// Solidity: event BlacklistedRemoved(address indexed account)
func (_YellowToken *YellowTokenFilterer) ParseBlacklistedRemoved(log types.Log) (*YellowTokenBlacklistedRemoved, error) {
	event := new(YellowTokenBlacklistedRemoved)
	if err := _YellowToken.contract.UnpackLog(event, "BlacklistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the YellowToken contract.
type YellowTokenRoleAdminChangedIterator struct {
	Event *YellowTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *YellowTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenRoleAdminChanged)
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
		it.Event = new(YellowTokenRoleAdminChanged)
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
func (it *YellowTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenRoleAdminChanged represents a RoleAdminChanged event raised by the YellowToken contract.
type YellowTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_YellowToken *YellowTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*YellowTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenRoleAdminChangedIterator{contract: _YellowToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_YellowToken *YellowTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *YellowTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenRoleAdminChanged)
				if err := _YellowToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_YellowToken *YellowTokenFilterer) ParseRoleAdminChanged(log types.Log) (*YellowTokenRoleAdminChanged, error) {
	event := new(YellowTokenRoleAdminChanged)
	if err := _YellowToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the YellowToken contract.
type YellowTokenRoleGrantedIterator struct {
	Event *YellowTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *YellowTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenRoleGranted)
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
		it.Event = new(YellowTokenRoleGranted)
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
func (it *YellowTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenRoleGranted represents a RoleGranted event raised by the YellowToken contract.
type YellowTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_YellowToken *YellowTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*YellowTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenRoleGrantedIterator{contract: _YellowToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_YellowToken *YellowTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *YellowTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenRoleGranted)
				if err := _YellowToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_YellowToken *YellowTokenFilterer) ParseRoleGranted(log types.Log) (*YellowTokenRoleGranted, error) {
	event := new(YellowTokenRoleGranted)
	if err := _YellowToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the YellowToken contract.
type YellowTokenRoleRevokedIterator struct {
	Event *YellowTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *YellowTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowTokenRoleRevoked)
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
		it.Event = new(YellowTokenRoleRevoked)
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
func (it *YellowTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowTokenRoleRevoked represents a RoleRevoked event raised by the YellowToken contract.
type YellowTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_YellowToken *YellowTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*YellowTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YellowToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &YellowTokenRoleRevokedIterator{contract: _YellowToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_YellowToken *YellowTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *YellowTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YellowToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowTokenRoleRevoked)
				if err := _YellowToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_YellowToken *YellowTokenFilterer) ParseRoleRevoked(log types.Log) (*YellowTokenRoleRevoked, error) {
	event := new(YellowTokenRoleRevoked)
	if err := _YellowToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
