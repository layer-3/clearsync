package smart_wallet

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/layer-3/clearsync/pkg/abi/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/abi/simple_account"
)

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func mustTrue[T any](x T, b bool) T {
	if !b {
		panic(fmt.Errorf("unexpected false boolean value"))
	}
	return x
}

var entryPointABI = must(abi.JSON(strings.NewReader(entry_point_v0_6_0.EntryPointMetaData.ABI)))
var simpleAccountABI = must(abi.JSON(strings.NewReader(simple_account.SimpleAccountMetaData.ABI)))

// kernelExecuteABI is used to execute a transaction on Zerodev Kernel smart account.
var KernelExecuteABI = must(abi.JSON(strings.NewReader(`[
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      },
      {
        "internalType": "enum Operation",
        "name": "",
        "type": "uint8"
      }
    ],
    "name": "execute",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [{
      "components": [
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        },
        {
          "internalType": "bytes",
          "name": "data",
          "type": "bytes"
        }
      ],
      "internalType": "struct Call[]",
      "name": "calls",
      "type": "tuple[]"
    }],
    "name": "executeBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  }
]`)))

// kernelDeployWalletABI is used to deploy a new smart account on Zerodev Kernel.
var kernelDeployWalletABI = must(abi.JSON(strings.NewReader(`[{
  "inputs":[
    {
      "internalType":"address",
      "name":"_implementation",
      "type":"address"
    },
    {
      "internalType":"bytes",
      "name":"_data",
      "type":"bytes"
    },
    {
      "internalType":"uint256",
      "name":"_index",
      "type":"uint256"
    }
  ],
  "name":"createAccount",
  "outputs":[
    {
      "internalType":"address",
      "name":"proxy",
      "type":"address"
    }
  ],
  "stateMutability":"payable",
  "type":"function"
}]`)))

// kernelInitABI is the init ABI, used to initialise Zerodev Kernel smart account.
var kernelInitABI = must(abi.JSON(strings.NewReader(`[{
  "inputs": [
    {
      "internalType": "contract IKernelValidator",
      "name": "_defaultValidator",
      "type": "address"
    },
    {
      "internalType": "bytes",
      "name": "_data",
      "type": "bytes"
    }
  ],
  "name": "initialize",
  "outputs": [],
  "stateMutability": "payable",
  "type": "function"
}]`)))

// biconomyDeployWalletABI is used to deploy a new smart account on Biconomy.
var biconomyDeployWalletABI = must(abi.JSON(strings.NewReader(`[{
  "inputs": [
    {
      "internalType": "address",
      "name": "moduleSetupContract",
      "type": "address"
    },
    {
      "internalType": "bytes",
      "name": "moduleSetupData",
      "type": "bytes"
    },
    {
      "internalType": "uint256",
      "name": "index",
      "type": "uint256"
    }
  ],
  "name": "deployCounterFactualAccount",
  "outputs": [{
    "internalType": "address",
    "name": "proxy",
    "type": "address"
  }],
  "stateMutability": "nonpayable",
  "type": "function"
}]`)))

// biconomyInitABI is the init ABI, used to initialise Biconomy smart account.
var biconomyInitABI = must(abi.JSON(strings.NewReader(`[
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "handler",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "moduleSetupContract",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "name": "moduleSetupData",
        "type": "bytes"
      }
    ],
    "name": "init",
    "outputs": [{
      "internalType": "address",
      "name": "",
      "type": "address"
    }],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{
      "internalType": "address",
      "name": "eoaOwner",
      "type": "address"
    }],
    "name": "initForSmartAccount",
    "outputs": [{
      "internalType": "address",
      "name": "",
      "type": "address"
    }],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`)))
