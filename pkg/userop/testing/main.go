package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/itoken"
	"github.com/layer-3/clearsync/pkg/userop"
)

func main() {

	var (
		owner       = common.HexToAddress("0x2185da3337cad307fd48dFDabA6D4C66A9fD2c71")
		smartWallet = common.HexToAddress("0x69b36b0Cb89b1666d85Ed4fF48243730E9c53405")
		receiver    = common.HexToAddress("0x2185da3337cad307fd48dFDabA6D4C66A9fD2c71")
		duckies     = common.HexToAddress("0x18e73A5333984549484348A94f4D219f4faB7b81") // Duckies
		amount      = decimal.RequireFromString("1000")                                 // wei

		ducklingsGame    = common.HexToAddress("0xb66bf78cad7cbab51988ddc792652cbabdff7675") // Duckies
		ducklingsGameABI = `[{
      "inputs": [
        {
          "internalType": "uint8",
          "name": "size",
          "type": "uint8"
        }
      ],
      "name": "mintPack",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }]`
	)

	// create smartWallet client (with specific Wallet and Paymaster types)
	client, err := userop.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	// calculate smart wallet address
	walletAddress, err := client.GetAccountAddress(context.Background(), owner, decimal.Zero)
	if err != nil {
		panic(fmt.Errorf("failed to get wallet address: %w", err))
	}
	slog.Info("wallet address", "address", walletAddress)

	// You can send native tokens to any address.
	transferNative, err := newTransferNativeCall(receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build transfer native call: %w", err))
	}
	if err := send(client, smartWallet, []userop.Call{transferNative}); err != nil {
		panic(err)
	}

	// NOTE: prior to using PimlicoERC20Paymaster, make sure to approve the
	// paymaster contract to spend your fee token.
	approve, err := newApproveCall(duckies, receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build approve call: %w", err))
	}
	if err := send(client, smartWallet, []userop.Call{approve}); err != nil {
		panic(err)
	}

	// Now this call can be paid with ERC20 tokens using PimlicoERC20Paymaster.
	transferERC20, err := newTransferERC20Call(duckies, receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build transfer erc20 call: %w", err))
	}
	if err := send(client, smartWallet, []userop.Call{transferERC20}); err != nil {
		panic(err)
	}

	// You can also submit several calls in a single userOp.
	mintPrice := decimal.RequireFromString("5000000000") // 50 duckies for 1 Duckling
	approveToGame, err := newApproveCall(duckies, ducklingsGame, mintPrice)
	if err != nil {
		panic(fmt.Errorf("failed to build approve to game call: %w", err))
	}

	mintPack, err := newCallFromABI(ducklingsGame, ducklingsGameABI, decimal.NewFromInt(0), "mintPack", 1)
	if err != nil {
		panic(fmt.Errorf("failed to build mint pack call: %w", err))
	}

	if err := send(client, smartWallet, []userop.Call{approveToGame, mintPack}); err != nil {
		panic(err)
	}
}

// Encodes an `approve` call to the `token` contract, approving `amount` to be spent by `spender`.
func newApproveCall(token, spender common.Address, amount decimal.Decimal) (userop.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ERC20 ABI: %w", err))
	}

	callData, err := erc20.Pack("approve", spender, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to pack transfer data: %w", err))
	}

	return userop.Call{
		To:       token,
		Value:    decimal.Zero,
		CallData: callData,
	}, nil
}

// Encodes a `transfer` call of a native token, transferring `amount` to `receiver`.
func newTransferNativeCall(receiver common.Address, amount decimal.Decimal) (userop.Call, error) {
	return userop.Call{
		To:    receiver,
		Value: amount,
	}, nil
}

// Encodes a `transfer` call to the `token` contract, transferring `amount` to `receiver`.
func newTransferERC20Call(token, receiver common.Address, amount decimal.Decimal) (userop.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ERC20 ABI: %w", err))
	}

	callData, err := erc20.Pack("transfer", receiver, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to pack transfer data: %w", err))
	}

	return userop.Call{
		To:       token,
		Value:    decimal.Zero,
		CallData: callData,
	}, nil
}

// Encodes a call to the `contract` with the given `value`, `method` and `args`.
func newCallFromABI(contract common.Address, stringABI string, value decimal.Decimal, method string, args ...interface{}) (userop.Call, error) {
	ABI, err := abi.JSON(strings.NewReader(stringABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ABI: %w", err))
	}

	callData, err := ABI.Pack(method, args...)
	if err != nil {
		panic(fmt.Errorf("failed to pack call data: %w", err))
	}

	return userop.Call{
		To:       contract,
		Value:    value,
		CallData: callData,
	}, nil
}

// Creates and sends the user operation.
// NOTE: when sending the first userOp from a Smart Wallet,
// `config.example.go/walletDeploymentOpts` must contain Smart Wallet owner EOA address and SW index (0 by default).
func send(client userop.UserOperationClient, smartWallet common.Address, calls []userop.Call) error {
	ctx := context.Background()

	op, err := client.NewUserOp(ctx, smartWallet, signer, calls, walletDeploymentOpts)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	b, err := op.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal userop: %w", err)
	}
	slog.Info("sending user operation", "op", string(b))

	if _, err := client.SendUserOp(ctx, op); err != nil {
		return fmt.Errorf("failed to send userop: %w", err)
	}
	return nil
}
