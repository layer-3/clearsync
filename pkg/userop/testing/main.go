package main

import (
	"context"
	"fmt"
	"log/slog"
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

	approve, err := newApproveCall()
	if err != nil {
		panic(fmt.Errorf("failed to build call: %w", err))
	}
	if err := send(client, approve); err != nil {
		panic(err)
	}

	transfer, err := newTransferCall()
	if err != nil {
		panic(fmt.Errorf("failed to build call: %w", err))
	}
	if err := send(client, transfer); err != nil {
		panic(err)
	}
}

func newApproveCall() (userop.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ERC20 ABI: %w", err))
	}

	callData, err := erc20.Pack("approve", config.Paymaster.Address, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to pack transfer data: %w", err))
	}

	return userop.Call{
		To:       token,
		Value:    decimal.Zero,
		CallData: callData,
	}, nil
}

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

func send(client userop.UserOperationClient, call userop.Call) error {
	ctx := context.Background()

	op, err := client.NewUserOp(ctx, sender, []userop.Call{call})
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	b, err := op.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal userop: %w", err)
	}
	slog.Info("sending user operation", "op", string(b))

	callback := func() {}
	if err := client.SendUserOp(ctx, op, callback); err != nil {
		return fmt.Errorf("failed to send userop: %w", err)
	}
	return nil
}
