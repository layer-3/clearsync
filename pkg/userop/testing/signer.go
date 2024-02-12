package main

import (
	"fmt"
	"log"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/layer-3/clearsync/pkg/userop"
)

// GetUserOperationHash computes the hash of a UserOperation.
func getUserOperationHash(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) common.Hash {
	packedOp := op.Pack()
	hashedOp := crypto.Keccak256(packedOp)

	arguments := abi.Arguments{
		{Type: abi.Type{T: abi.BytesTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
	}

	packed, err := arguments.Pack(
		hashedOp,
		entryPoint,
		chainID,
	)
	if err != nil {
		panic(err) // Handle error appropriately in real code
	}

	return crypto.Keccak256Hash(packed)
}

func signUserOp(userOperation userop.UserOperation, entryPoint common.Address, chainID *big.Int) common.Hash {
	slog.Info("signing user operation")

	hash := getUserOperationHash(userOperation, entryPoint, chainID)
	privateKey, err := crypto.HexToECDSA("9be4e88f3d84ff8b58ec9f11f047245cad2f8e2c6cf43ca4087331ffc0fea6c8")
	if err != nil {
		log.Fatal(err)
	}

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return common.Hash{}
	}

	signatureStr := fmt.Sprintf("0x%s", common.Bytes2Hex(signature))
	result := common.Hash([]byte(signatureStr))
	slog.Info("user operation signed", "hash", result.Hex())
	return result
}
