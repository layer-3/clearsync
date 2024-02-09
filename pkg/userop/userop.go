// Package userop provides a ERC-4337 pseudo-transaction object called a UserOperation
// which is used to execute actions through a smart contract account.
// This isn't to be mistaken for a regular transaction type.
package userop

import (
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
