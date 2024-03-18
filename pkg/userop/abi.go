package userop

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

var (
	address = must(abi.NewType("address", "", nil))
	uint256 = must(abi.NewType("uint256", "", nil))
	bytes32 = must(abi.NewType("bytes32", "", nil))
	bytes   = must(abi.NewType("bytes", "", nil))

	// keccak256("UserOperationEvent(bytes32,address,address,uint256,bool,uint256, uint256)")
	userOpEventID = common.HexToHash("0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f")
	// keccak256("UserOperationRevertReason(bytes32,address,uint256,bytes)")
	userOpRevertReasonID = common.HexToHash("0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201")
)

// ABI with UserOp events from EntryPoint contract.
var entryPointUserOpEventsABI = must(abi.JSON(strings.NewReader(`[
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
]`)))
