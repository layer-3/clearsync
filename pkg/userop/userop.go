// Package userop provides an ERC-4337 pseudo-transaction object called a UserOperation
// which is used to execute actions through a smart contract account.
// This isn't to be mistaken for a regular transaction type.
package userop

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// UserOperation represents an EIP-4337 style transaction for a smart contract account.
type UserOperation struct {
	Sender               common.Address  `json:"sender"`
	Nonce                decimal.Decimal `json:"nonce"`
	InitCode             []byte          `json:"initCode"`
	CallData             []byte          `json:"callData"`
	CallGasLimit         decimal.Decimal `json:"callGasLimit"`
	VerificationGasLimit decimal.Decimal `json:"verificationGasLimit"`
	PreVerificationGas   decimal.Decimal `json:"preVerificationGas"`
	MaxFeePerGas         decimal.Decimal `json:"maxFeePerGas"`
	MaxPriorityFeePerGas decimal.Decimal `json:"maxPriorityFeePerGas"`
	PaymasterAndData     []byte          `json:"paymasterAndData"`
	Signature            []byte          `json:"signature,omitempty"`
}

// GetFactory returns the address portion of InitCode if applicable.
// Otherwise, it returns the zero address.
func (op *UserOperation) GetFactory() common.Address {
	if len(op.InitCode) < common.AddressLength {
		return common.HexToAddress("0x")
	}

	return common.BytesToAddress(op.InitCode[:common.AddressLength])
}

// GetFactoryData returns the data portion of InitCode if applicable.
// Otherwise, it returns an empty byte array.
func (op *UserOperation) GetFactoryData() []byte {
	if len(op.InitCode) < common.AddressLength {
		return []byte{}
	}

	return op.InitCode[common.AddressLength:]
}

// UserOpHash returns the hash of the userOp + entryPoint address + chainID.
func (op *UserOperation) UserOpHash(entryPoint common.Address, chainID *big.Int) common.Hash {
	args := abi.Arguments{
		{Name: "sender", Type: address},
		{Name: "nonce", Type: uint256},
		{Name: "hashInitCode", Type: bytes32},
		{Name: "hashCallData", Type: bytes32},
		{Name: "callGasLimit", Type: uint256},
		{Name: "verificationGasLimit", Type: uint256},
		{Name: "preVerificationGas", Type: uint256},
		{Name: "maxFeePerGas", Type: uint256},
		{Name: "maxPriorityFeePerGas", Type: uint256},
		{Name: "hashPaymasterAndData", Type: bytes32},
	}
	packed, err := args.Pack(
		op.Sender,
		op.Nonce.BigInt(),
		crypto.Keccak256Hash(op.InitCode),
		crypto.Keccak256Hash(op.CallData),
		op.CallGasLimit.BigInt(),
		op.VerificationGasLimit.BigInt(),
		op.PreVerificationGas.BigInt(),
		op.MaxFeePerGas.BigInt(),
		op.MaxPriorityFeePerGas.BigInt(),
		crypto.Keccak256Hash(op.PaymasterAndData),
	)
	if err != nil { // This should never happen
		panic(err)
	}

	return crypto.Keccak256Hash(
		crypto.Keccak256(packed),
		common.LeftPadBytes(entryPoint.Bytes(), 32),
		common.LeftPadBytes(chainID.Bytes(), 32),
	)
}

// MarshalJSON returns a JSON encoding of the UserOperation.
func (op UserOperation) MarshalJSON() ([]byte, error) {
	ic := "0x"
	if fa := op.GetFactory(); fa != common.HexToAddress("0x") {
		ic = fmt.Sprintf("%s%s", fa, common.Bytes2Hex(op.GetFactoryData()))
	}

	return json.Marshal(&struct {
		Sender               string `json:"sender"`
		Nonce                string `json:"nonce"`
		InitCode             string `json:"initCode"`
		CallData             string `json:"callData"`
		CallGasLimit         string `json:"callGasLimit"`
		VerificationGasLimit string `json:"verificationGasLimit"`
		PreVerificationGas   string `json:"preVerificationGas"`
		MaxFeePerGas         string `json:"maxFeePerGas"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
		PaymasterAndData     string `json:"paymasterAndData"`
		Signature            string `json:"signature"`
	}{
		Sender:               op.Sender.String(),
		Nonce:                hexutil.EncodeBig(op.Nonce.BigInt()),
		InitCode:             ic,
		CallData:             hexutil.Encode(op.CallData),
		CallGasLimit:         hexutil.EncodeBig(op.CallGasLimit.BigInt()),
		VerificationGasLimit: hexutil.EncodeBig(op.VerificationGasLimit.BigInt()),
		PreVerificationGas:   hexutil.EncodeBig(op.PreVerificationGas.BigInt()),
		MaxFeePerGas:         hexutil.EncodeBig(op.MaxFeePerGas.BigInt()),
		MaxPriorityFeePerGas: hexutil.EncodeBig(op.MaxPriorityFeePerGas.BigInt()),
		PaymasterAndData:     hexutil.Encode(op.PaymasterAndData),
		Signature:            hexutil.Encode(op.Signature),
	})
}

// SignWithECDSA signs the hash with the given private key using the ECDSA algorithm.
func (op UserOperation) SignWithECDSA(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	ethMessageHash := computeEthSignedMessageHash(hash)

	signature, err := crypto.Sign(ethMessageHash, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign user operation: %w", err)
	}

	// To make the signature compatible with
	// Ethereum's ECDSA recovery, ensure V is 27 or 28.
	if signature[64] < 27 {
		signature[64] += 27
	}

	slog.Debug("user operation signed:", "hash", common.Bytes2Hex(hash), "signature", hexutil.Encode(signature))
	return signature, nil
}

// computeEthSignedMessageHash accepts an arbitrary message, prepends a known message,
// and hashes the result using keccak256. The known message added to the input before hashing is
// "\x19Ethereum Signed Message:\n" + len(message).
func computeEthSignedMessageHash(message []byte) []byte {
	return crypto.Keccak256([]byte(
		fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), string(message)),
	))
}
