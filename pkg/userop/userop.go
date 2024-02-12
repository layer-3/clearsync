// Package userop provides a ERC-4337 pseudo-transaction object called a UserOperation
// which is used to execute actions through a smart contract account.
// This isn't to be mistaken for a regular transaction type.
package userop

import (
	"encoding/hex"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// TODO: verify userop validity (https://www.erc4337.io/docs/bundlers/running-a-bundler)

// UserOperation represents an EIP-4337 style transaction for a smart contract account.
type UserOperation struct {
	Sender               common.Address  `json:"sender,omitempty"`
	Nonce                decimal.Decimal `json:"nonce,omitempty"`
	InitCode             []byte          `json:"initCode"`
	CallData             []byte          `json:"callData"`
	CallGasLimit         decimal.Decimal `json:"callGasLimit,omitempty"`
	VerificationGasLimit decimal.Decimal `json:"verificationGasLimit,omitempty"`
	PreVerificationGas   decimal.Decimal `json:"preVerificationGas,omitempty"`
	MaxFeePerGas         decimal.Decimal `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas decimal.Decimal `json:"maxPriorityFeePerGas,omitempty"`
	PaymasterAndData     []byte          `json:"paymasterAndData"`
	Signature            common.Hash     `json:"signature,omitempty"`
}

// Marshal method implements custom serialization for user operation.
// Namely, it converts []byte fields to hex strings and decimal.Decimal fields to strings.
func (op *UserOperation) Marshal() map[string]interface{} {
	// Prepare the object for serialization
	result := map[string]interface{}{
		"sender":               op.Sender,
		"nonce":                "0x" + op.Nonce.BigInt().Text(16),
		"initCode":             "0x" + hex.EncodeToString(op.InitCode),
		"callData":             "0x" + hex.EncodeToString(op.CallData),
		"callGasLimit":         "0x" + op.CallGasLimit.BigInt().Text(16),
		"verificationGasLimit": "0x" + op.VerificationGasLimit.BigInt().Text(16),
		"preVerificationGas":   "0x" + op.PreVerificationGas.BigInt().Text(16),
		"maxFeePerGas":         "0x" + op.MaxFeePerGas.BigInt().Text(16),
		"maxPriorityFeePerGas": "0x" + op.MaxPriorityFeePerGas.BigInt().Text(16),
		"paymasterAndData":     "0x" + hex.EncodeToString(op.PaymasterAndData),
		"signature":            op.Signature,
	}

	slog.Info("marshalling", "userop", result)
	return result
}
