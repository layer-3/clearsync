package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/layer-3/clearsync/pkg/abi/itoken"
	"github.com/layer-3/clearsync/pkg/session_key"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/shopspring/decimal"
)

var (
	userOpConfig           = exampleUserOpConfig
	signer                 = exampleSigner
	userOpSigner           = exampleUserOpSigner
	sessionKeyConfig       = exampleSessionKeyConfig
	sessionKeySigner       = exampleSessionKeySigner
	sessionKeyUserOpSigner = exampleSessionKeyUserOpSigner
)

func main() {
	setLogLevel(slog.LevelInfo)

	var (
		chainId     = big.NewInt(137) // Matic
		owner       = common.HexToAddress("0x2185da3337cad307fd48dFDabA6D4C66A9fD2c71")
		smartWallet = common.HexToAddress("0x69b36b0Cb89b1666d85Ed4fF48243730E9c53405")
		receiver    = common.HexToAddress("0x2185da3337cad307fd48dFDabA6D4C66A9fD2c71")
		token       = common.HexToAddress("0x18e73A5333984549484348A94f4D219f4faB7b81") // Duckies
		amount      = decimal.RequireFromString("1000")                                 // wei
	)

	ctx := context.Background()

	// create smartWallet userOpClient (with specific Wallet and Paymaster types)
	userOpClient, err := userop.NewClient(userOpConfig)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	// calculate smart wallet address
	walletAddress, err := userOpClient.GetAccountAddress(ctx, owner, decimal.Zero)
	if err != nil {
		panic(fmt.Errorf("failed to get wallet address: %w", err))
	}
	slog.Debug("wallet address", "address", walletAddress)

	// You can send native tokens to any address.
	transferNative := userop.Call{
		To:    receiver,
		Value: amount.BigInt(),
	}
	if err := createAndSendUserop(userOpClient, userOpSigner, smartWallet, []userop.Call{transferNative}); err != nil {
		panic(err)
	}

	// smart wallet should be deployed now

	sessionKeyClient, err := session_key.NewClient(sessionKeyConfig)
	if err != nil {
		panic(fmt.Errorf("failed to create session key client: %w", err))
	}

	incompleteEnablingSKSigner, err := sessionKeyClient.GetIncompleteEnablingUserOpSigner(sessionKeySigner)
	if err != nil {
		panic(fmt.Errorf("failed to get incomplete enabling user op signer: %w", err))
	}

	// skSigner := sessionKeyClient.GetUserOpSigner(exampleSessionKeySigner)

	// enable and use session key
	transferERC20, err := newTransferERC20Call(token, receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build transfer erc20 call: %w", err))
	}

	op, err := userOpClient.NewUserOp(ctx, smartWallet, incompleteEnablingSKSigner, userop.Calls{transferERC20}, walletDeploymentOpts)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	incompleteSig := op.Signature

	sessionData, err := session_key.UnpackEnableData(incompleteSig)
	if err != nil {
		panic(fmt.Errorf("failed to unpack enable data: %w", err))
	}

	enableDigest := sessionKeyClient.GetEnableDataDigest(smartWallet, sessionData, session_key.KernelExecuteBatchSig, chainId)

	enableSig, err := signer.Sign(enableDigest)
	if err != nil {
		panic(fmt.Errorf("failed to sign enable data digest: %w", err))
	}

	offset := session_key.KernelEnableSigOffset
	copy(incompleteSig[offset:offset+65], enableSig.Raw())
	op.Signature = incompleteSig

	sendUserop(userOpClient, op)

	// use session key
	approveCall, err := newApproveCall(token, receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build approve call: %w", err))
	}

	createAndSendUserop(userOpClient, sessionKeyUserOpSigner, smartWallet, []userop.Call{approveCall})
}

func setLogLevel(level slog.Level) {
	lvl := new(slog.LevelVar)
	lvl.Set(level)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	slog.SetDefault(logger)
}

// Creates and sends the user operation.
// NOTE: when sending the first userOp from a Smart Wallet,
// `config.example.go/walletDeploymentOpts` must contain Smart Wallet owner EOA address and SW index (0 by default).
func createAndSendUserop(client userop.Client, signer userop.Signer, smartWallet common.Address, calls []userop.Call) error {
	ctx := context.Background()

	op, err := client.NewUserOp(ctx, smartWallet, signer, calls, walletDeploymentOpts)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	return sendUserop(client, op)
}

func sendUserop(client userop.Client, op userop.UserOperation) error {
	ctx := context.Background()

	b, err := op.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal userop: %w", err)
	}
	slog.Debug("sending user operation", "op", string(b))

	waitForUserOp, err := client.SendUserOp(ctx, op)
	if err != nil {
		return fmt.Errorf("failed to send userop: %w", err)
	}

	userOpReceipt := <-waitForUserOp

	slog.Info("user operation verified", "userOpReceipt", userOpReceipt)

	return nil
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
		Value:    big.NewInt(0),
		CallData: callData,
	}, nil
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
		Value:    big.NewInt(0),
		CallData: callData,
	}, nil
}
