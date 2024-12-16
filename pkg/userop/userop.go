// Package userop provides an ERC-4337 pseudo-transaction object called a UserOperation
// which is used to execute actions through a smart contract account.
// This isn't to be mistaken for a regular transaction type.
package userop

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// TODO: replace `decimal.Decimal` with `*big.Int` as corresponding fields are always integers.
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
func (op *UserOperation) UserOpHash(entryPoint common.Address, chainID *big.Int) (common.Hash, error) {
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
		return common.Hash{}, fmt.Errorf("failed to pack UserOperation: %w", err)
	}

	return crypto.Keccak256Hash(
		crypto.Keccak256(packed),
		common.LeftPadBytes(entryPoint.Bytes(), 32),
		common.LeftPadBytes(chainID.Bytes(), 32),
	), nil
}

// DeepCopy creates and returns a deep copy of the UserOperation.
func (op *UserOperation) DeepCopy() *UserOperation {
	if op == nil {
		return nil
	}

	copyOp := &UserOperation{
		Sender:               op.Sender,
		Nonce:                op.Nonce,
		CallGasLimit:         op.CallGasLimit,
		VerificationGasLimit: op.VerificationGasLimit,
		PreVerificationGas:   op.PreVerificationGas,
		MaxFeePerGas:         op.MaxFeePerGas,
		MaxPriorityFeePerGas: op.MaxPriorityFeePerGas,
	}

	copyOp.InitCode = make([]byte, len(op.InitCode))
	copy(copyOp.InitCode, op.InitCode)

	copyOp.CallData = make([]byte, len(op.CallData))
	copy(copyOp.CallData, op.CallData)

	copyOp.PaymasterAndData = make([]byte, len(op.PaymasterAndData))
	copy(copyOp.PaymasterAndData, op.PaymasterAndData)

	copyOp.Signature = make([]byte, len(op.Signature))
	copy(copyOp.Signature, op.Signature)

	return copyOp
}

type UserOperationDTO struct {
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
}

// MarshalJSON returns a JSON encoding of the UserOperation.
func (op UserOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&UserOperationDTO{
		Sender:               op.Sender.String(),
		Nonce:                hexutil.EncodeBig(op.Nonce.BigInt()),
		InitCode:             hexutil.Encode(op.InitCode),
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

// UnmarshalJSON decodes a JSON encoding into a UserOperation.
func (op *UserOperation) UnmarshalJSON(data []byte) error {
	var uoDTO UserOperationDTO
	if err := json.Unmarshal(data, &uoDTO); err != nil {
		return err
	}

	var err error
	op.Sender = common.HexToAddress(uoDTO.Sender)

	if uoDTO.Nonce == "" {
		op.Nonce = decimal.NewFromInt(0)
	} else {
		if nonceBI, ok := big.NewInt(0).SetString(uoDTO.Nonce, 0); !ok {
			return fmt.Errorf("invalid nonce: %w", err)
		} else {
			op.Nonce = decimal.NewFromBigInt(nonceBI, 0)
		}
	}

	if uoDTO.InitCode == "" {
		op.InitCode = []byte{}
	} else {
		if op.InitCode, err = hexutil.Decode(uoDTO.InitCode); err != nil {
			return fmt.Errorf("invalid initCode: %w", err)
		}
	}

	if uoDTO.CallData == "" {
		op.CallData = []byte{}
	} else {
		if op.CallData, err = hexutil.Decode(uoDTO.CallData); err != nil {
			return fmt.Errorf("invalid callData: %w", err)
		}
	}

	if uoDTO.CallGasLimit == "" {
		op.CallGasLimit = decimal.NewFromInt(0)
	} else {
		if callGasLimitBI, ok := new(big.Int).SetString(uoDTO.CallGasLimit, 0); !ok {
			return fmt.Errorf("invalid callGasLimit: %w", err)
		} else {
			op.CallGasLimit = decimal.NewFromBigInt(callGasLimitBI, 0)
		}
	}

	if uoDTO.VerificationGasLimit == "" {
		op.VerificationGasLimit = decimal.NewFromInt(0)
	} else {
		if verificationGasLimitBI, ok := new(big.Int).SetString(uoDTO.VerificationGasLimit, 0); !ok {
			return fmt.Errorf("invalid verificationGasLimit: %w", err)
		} else {
			op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimitBI, 0)
		}
	}

	if uoDTO.PreVerificationGas == "" {
		op.PreVerificationGas = decimal.NewFromInt(0)
	} else {
		if preVerificationGasBI, ok := new(big.Int).SetString(uoDTO.PreVerificationGas, 0); !ok {
			return fmt.Errorf("invalid preVerificationGas: %w", err)
		} else {
			op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGasBI, 0)
		}
	}

	if uoDTO.MaxFeePerGas == "" {
		op.MaxFeePerGas = decimal.NewFromInt(0)
	} else {
		if maxFeePerGasBI, ok := new(big.Int).SetString(uoDTO.MaxFeePerGas, 0); !ok {
			return fmt.Errorf("invalid maxFeePerGas: %w", err)
		} else {
			op.MaxFeePerGas = decimal.NewFromBigInt(maxFeePerGasBI, 0)
		}
	}

	if uoDTO.MaxPriorityFeePerGas == "" {
		op.MaxPriorityFeePerGas = decimal.NewFromInt(0)
	} else {
		if maxPriorityFeePerGasBI, ok := new(big.Int).SetString(uoDTO.MaxPriorityFeePerGas, 0); !ok {
			return fmt.Errorf("invalid maxPriorityFeePerGas: %w", err)
		} else {
			op.MaxPriorityFeePerGas = decimal.NewFromBigInt(maxPriorityFeePerGasBI, 0)
		}
	}

	if uoDTO.PaymasterAndData == "" {
		op.PaymasterAndData = []byte{}
	} else {
		if op.PaymasterAndData, err = hexutil.Decode(uoDTO.PaymasterAndData); err != nil {
			return fmt.Errorf("invalid paymasterAndData: %w", err)
		}
	}

	if uoDTO.Signature == "" {
		op.Signature = []byte{}
	} else {
		if op.Signature, err = hexutil.Decode(uoDTO.Signature); err != nil {
			return fmt.Errorf("invalid signature: %w", err)
		}
	}

	return nil
}
