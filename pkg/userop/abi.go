package userop

import "github.com/ethereum/go-ethereum/common"

var (
	// keccak256("UserOperationEvent(bytes32,address,address,uint256,bool,uint256, uint256)")
	userOpEventID = common.HexToHash("0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f")
	// keccak256("UserOperationRevertReason(bytes32,address,uint256,bytes)")
	userOpRevertReasonID = common.HexToHash("0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201")
)

// ABI with UserOp events from EntryPoint contract.
const entrypointUserOpEventsABI = `[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "bytes32",
        "name": "userOpHash",
        "type": "bytes32"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "paymaster",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "nonce",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "success",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "actualGasCost",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "actualGasUsed",
        "type": "uint256"
      }
    ],
    "name": "UserOperationEvent",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "bytes32",
        "name": "userOpHash",
        "type": "bytes32"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "factory",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "paymaster",
        "type": "address"
      }
    ],
    "name": "AccountDeployed",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "bytes32",
        "name": "userOpHash",
        "type": "bytes32"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "nonce",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "revertReason",
        "type": "bytes"
      }
    ],
    "name": "UserOperationRevertReason",
    "type": "event"
  }
]`

// kernelExecuteABI is used to execute a transaction on Zerodev Kernel smart account.
const kernelExecuteABI = `[
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
]`

// kernelDeployWalletABI is used to deploy a new smart account on Zerodev Kernel.
const kernelDeployWalletABI = `[{
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
}]`

// kernelInitABI is the init ABI, used to initialise Zerodev Kernel smart account.
const kernelInitABI = `[{
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
}]`

// biconomyDeployWalletABI is used to deploy a new smart account on Biconomy.
const biconomyDeployWalletABI = `[{
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
}]`

// biconomyInitABI is the init ABI, used to initialise Biconomy smart account.
const biconomyInitABI = `[
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
  },
]`
