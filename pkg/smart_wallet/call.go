package smart_wallet

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Call represents sufficient data to build a single transaction,
// which is a part of a user operation to be executed in a batch with other ones.
type Call struct {
	To       common.Address
	Value    *big.Int // wei
	CallData []byte
}

type Calls []Call

func BuildCallData(swt Type, calls Calls) ([]byte, error) {
	switch swt {
	case SimpleAccountType, BiconomyType:
		return calls.PackForSimpleAccount()
	case KernelType:
		return calls.PackForKernel()
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", swt)
	}
}

// PackForSimpleAccount packs CallData for SimpleAccount smart wallet.
func (calls Calls) PackForSimpleAccount() ([]byte, error) {
	addresses := make([]common.Address, len(calls))
	calldatas := make([][]byte, len(calls))
	for i, call := range calls {
		addresses[i] = call.To
		calldatas[i] = call.CallData
	}

	// Pack the data for the `executeBatch` smart account function.
	// Biconomy v2.0: https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol#L128
	// NOTE: you can NOT send native token with SimpleAccount v0.6.0 because of `executeBatch` signature
	data, err := simpleAccountABI.Pack("executeBatch", addresses, calldatas)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for SimpleAccount: %w", err)
	}
	return data, nil
}

// UnpackCallsForSimpleAccount unpacks CallData for SimpleAccount smart wallet.
func UnpackCallsForSimpleAccount(data []byte) (Calls, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("invalid data length")
	}

	values, err := simpleAccountABI.Methods["executeBatch"].Inputs.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack executeBatch data for SimpleAccount: %w", err)
	}

	addresses, ok := values[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("failed to unpack addresses for SimpleAccount")
	}

	calldatas, ok := values[1].([][]byte)
	if !ok {
		return nil, fmt.Errorf("failed to unpack calldatas for SimpleAccount")
	}

	if len(addresses) != len(calldatas) {
		return nil, fmt.Errorf("addresses and calldatas length mismatch")
	}

	calls := make(Calls, len(addresses))
	for i := range addresses {
		calls[i] = Call{
			To:       addresses[i],
			Value:    big.NewInt(0),
			CallData: calldatas[i],
		}
	}

	return calls, nil
}

// callStructKernel represents a call to the Zerodev Kernel smart wallet.
// The idea is the same as in Call type,
// but tailed specifically to the Zerodev Kernel ABI.
type callStructKernel struct {
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`
	Data  []byte         `json:"data"`
}

// handleCallKernel packs calldata for Zerodev Kernel smart wallet.
func (calls Calls) PackForKernel() ([]byte, error) {
	params := make([]callStructKernel, len(calls))
	for i, call := range calls {
		value := big.NewInt(0)
		if call.Value != nil {
			value.Set(call.Value)
		}

		params[i] = callStructKernel{
			To:    call.To,
			Value: value,
			Data:  call.CallData,
		}
	}

	// pack the data for the `executeBatch` smart account function
	// Zerodev Kernel v2.2: https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L82
	data, err := KernelExecuteABI.Pack("executeBatch", params)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for Kernel: %w", err)
	}
	return data, nil
}

// UnpackCallsForKernel unpacks CallData for Zerodev Kernel smart wallet.
func UnpackCallsForKernel(data []byte) (Calls, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("invalid data length")
	}

	values, err := KernelExecuteABI.Methods["executeBatch"].Inputs.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack executeBatch data for Kernel: %w", err)
	}

	params, ok := values[0].([]struct {
		To    common.Address `json:"to"`
		Value *big.Int       `json:"value"`
		Data  []byte         `json:"data"`
	})
	if !ok {
		return nil, fmt.Errorf("failed to unpack params for Kernel")
	}

	calls := make(Calls, len(params))
	for i, param := range params {
		calls[i] = Call{
			To:       param.To,
			Value:    param.Value,
			CallData: param.Data,
		}
	}

	return calls, nil
}
