package userop

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
  inputs: [
    {
      internalType: "address",
      name: "moduleSetupContract",
      type: "address"
    },
    {
      internalType: "bytes",
      name: "moduleSetupData",
      type: "bytes"
    },
    {
      internalType: "uint256",
      name: "index",
      type: "uint256"
    }
  ],
  name: "deployCounterFactualAccount",
  outputs: [{
    internalType: "address",
    name: "proxy",
    type: "address"
  }],
  stateMutability: "nonpayable",
  type: "function"
}]`

// kernelInitABI is the init ABI, used to initialise Zerodev Kernel smart account.
const kernelInitABI = `[{
  inputs: [
    {
      internalType: "contract IKernelValidator",
      name: "_defaultValidator",
      type: "address"
    },
    {
      internalType: "bytes",
      name: "_data",
      type: "bytes"
    }
  ],
  name: "initialize",
  outputs: [],
  stateMutability: "payable",
  type: "function"
}]`

// biconomyDeployWalletABI is used to deploy a new smart account on Biconomy.
const biconomyDeployWalletABI = `[{
  inputs: [
    {
      internalType: "address",
      name: "moduleSetupContract",
      type: "address"
    },
    {
      internalType: "bytes",
      name: "moduleSetupData",
      type: "bytes"
    },
    {
      internalType: "uint256",
      name: "index",
      type: "uint256"
    }
  ],
  name: "deployCounterFactualAccount",
  outputs: [{
    internalType: "address",
    name: "proxy",
    type: "address"
  }],
  stateMutability: "nonpayable",
  type: "function"
}]`

// biconomyInitABI is the init ABI, used to initialise Biconomy smart account.
const biconomyInitABI = `[
  {
    inputs: [
      {
        internalType: "address",
        name: "handler",
        type: "address"
      },
      {
        internalType: "address",
        name: "moduleSetupContract",
        type: "address"
      },
      {
        internalType: "bytes",
        name: "moduleSetupData",
        type: "bytes"
      }
    ],
    name: "init",
    outputs: [{
      internalType: "address",
      name: "",
      type: "address"
    }],
    stateMutability: "nonpayable",
    type: "function"
  },
  {
    inputs: [{
      internalType: "address",
      name: "eoaOwner",
      type: "address"
    }],
    name: "initForSmartAccount",
    outputs: [{
      internalType: "address",
      name: "",
      type: "address"
    }],
    stateMutability: "nonpayable",
    type: "function"
  },
]`
