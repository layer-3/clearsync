package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/layer-3/clearsync/pkg/userop"
)

// packUserOp packs a UserOperation into a byte slice according to the ABI encoding.
func packUserOp(userOperation userop.UserOperation) []byte {
	hashedInitCode := crypto.Keccak256(userOperation.InitCode)
	hashedCallData := crypto.Keccak256(userOperation.CallData)
	hashedPaymasterAndData := crypto.Keccak256(userOperation.PaymasterAndData)

	arguments := abi.Arguments{
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.BytesTy}},
		{Type: abi.Type{T: abi.BytesTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.BytesTy}},
	}

	packed, err := arguments.Pack(
		userOperation.Sender,
		userOperation.Nonce,
		hashedInitCode,
		hashedCallData,
		userOperation.CallGasLimit,
		userOperation.VerificationGasLimit,
		userOperation.PreVerificationGas,
		userOperation.MaxFeePerGas,
		userOperation.MaxPriorityFeePerGas,
		hashedPaymasterAndData,
	)
	if err != nil {
		panic(err) // Handle error appropriately in real code
	}
	return packed
}

// GetUserOperationHash computes the hash of a UserOperation.
func getUserOperationHash(userOperation userop.UserOperation, entryPoint common.Address, chainID *big.Int) common.Hash {
	packedOp := packUserOp(userOperation)
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

func main() {
	client, err := userop.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	op, err := client.NewUserOp(context.Background(), sender, receiver, token, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	callback := func() {}
	if err := client.SendUserOp(context.Background(), op, callback); err != nil {
		panic(fmt.Errorf("failed to send userop: %w", err))
	}
}
