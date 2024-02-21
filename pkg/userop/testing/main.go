package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/itoken"
	"github.com/layer-3/clearsync/pkg/userop"
)

func main() {
	client, err := userop.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	walletAddress, err := client.GetAccountAddress(context.Background(), owner, decimal.Zero)
	if err != nil {
		panic(fmt.Errorf("failed to get wallet address: %w", err))
	}

	slog.Info("wallet address", "address", walletAddress)

	approve, err := newApproveCall()
	if err != nil {
		panic(fmt.Errorf("failed to build call: %w", err))
	}
	if err := send(client, []userop.Call{approve}); err != nil {
		panic(err)
	}

	transfer, err := newTransferCall()
	if err != nil {
		panic(fmt.Errorf("failed to build call: %w", err))
	}
	if err := send(client, []userop.Call{transfer}); err != nil {
		panic(err)
	}
}

// Encodes an `approve` call to the `token` contract, approving `amount` to be spent by `receiver`.
func newApproveCall() (userop.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ERC20 ABI: %w", err))
	}

	callData, err := erc20.Pack("approve", receiver, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to pack transfer data: %w", err))
	}

	return userop.Call{
		To:       token,
		Value:    decimal.Zero,
		CallData: callData,
	}, nil
}

// Encodes a `transfer` call to the `token` contract, transferring `amount` to `receiver`.
func newTransferCall() (userop.Call, error) {
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

// TODO: add a `newCallFromABI` function that takes an ABI, method name and args and returns a `Call` object.
func newCallFromABI(abi string, value big.Int, method string, args ...interface{}) (userop.Call, error) {
	// ...
	return userop.Call{}, nil
}

// Creates and sends the user operation.
// NOTE: when sending the first userOp from a Smart Wallet,
// `config.example.go/walletDeploymentOpts` must contain Smart Wallet owner EOA address and SW index (0 by default).
func send(client userop.UserOperationClient, calls []userop.Call) error {
	ctx := context.Background()

	op, err := client.NewUserOp(ctx, sender, signer, calls, walletDeploymentOpts)
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
