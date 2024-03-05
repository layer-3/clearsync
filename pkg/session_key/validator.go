package session_key

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	mt "github.com/layer-3/go-merkletree"

	"github.com/layer-3/clearsync/pkg/signer"
)

const (
	multiPermissionProofABI = `
		[
			{
				"components":[
					{"internalType":"uint32","name":"index","type":"uint32"},
					{"internalType":"address","name":"target","type":"address"},
					{"internalType":"bytes4","name":"sig","type":"bytes4"},
					{"internalType":"uint256","name":"valueLimit","type":"uint256"},
					{
						"components":[
							{"internalType":"uint256","name":"offset","type":"uint256"},
							{"internalType":"enum ParamCondition","name":"condition","type":"uint8"},
							{"internalType":"bytes32","name":"param","type":"bytes32"}
						],
						"internalType":"struct ParamRule[]",
						"name":"rules",
						"type":"tuple[]"
					},{
						"components":[
							{"internalType":"ValidAfter","name":"validAfter","type":"uint48"},
							{"internalType":"uint48","name":"interval","type":"uint48"},
							{"internalType":"uint48","name":"runs","type":"uint48"}
						],
						"internalType":"struct ExecutionRule",
						"name":"executionRule",
						"type":"tuple"
					}
				],
				"internalType":"struct Permission[]",
				"name":"permission",
				"type":"tuple[]"
			},{
				"internalType":"bytes32[][]",
				"name":"merkleProof",
				"type":"bytes32[][]"
			}
		]
	`
)

func PackEnableValidatorSignature(enableData []byte, validator, executor common.Address, digestSig signer.Signature) []byte {
	signature := make([]byte, 0, 4+6+6+20+20+32+len(enableData)+32+32+32+1)
	// "enable validator" (0x00000002) mode
	// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
	signature = append(signature, []byte{0x00, 0x00, 0x00, 0x02}...)
	signature = append(signature, enableData[52:58]...)
	signature = append(signature, enableData[58:64]...)
	signature = append(signature, validator.Bytes()...)
	signature = append(signature, executor.Bytes()...)
	signature = append(signature, big.NewInt(int64(len(enableData))).FillBytes(make([]byte, 32))...)
	signature = append(signature, enableData...)
	signature = append(signature, big.NewInt(65).FillBytes(make([]byte, 32))...)
	signature = append(signature, digestSig.R...)
	signature = append(signature, digestSig.S...)
	signature = append(signature, digestSig.V)

	return signature
}

func PackUseSessionKeySignature(sessionKey common.Address, sessionKeySig signer.Signature, permissions []Permission, proofs []mt.Proof) ([]byte, error) {
	var args abi.Arguments
	dec := json.NewDecoder(strings.NewReader(multiPermissionProofABI))
	if err := dec.Decode(&args); err != nil {
		return nil, err
	}

	kernelPermissions := make([]kernelPermission, len(permissions))
	for i, permission := range permissions {
		kernelPermissions[i] = permission.toKernelPermission(uint32(i))
	}

	normProofs := make([][][32]byte, len(proofs))
	for i, proof := range proofs {
		normProofs[i] = make([][32]byte, len(proof.Siblings))
		for j, sibling := range proof.Siblings {
			copy(normProofs[i][j][:], sibling[:])
		}
	}

	permissionProof, err := args.Pack(kernelPermissions, normProofs)
	if err != nil {
		return nil, err
	}

	// session key (20) + sessionKeySig (65) + abi.encode(permissions, merkleProof)
	signature := make([]byte, 0, 20+32+len(permissionProof))
	signature = append(signature, sessionKey.Bytes()...)
	signature = append(signature, sessionKeySig.R...)
	signature = append(signature, sessionKeySig.S...)
	signature = append(signature, sessionKeySig.V)
	signature = append(signature, permissionProof...)

	return signature, nil
}
