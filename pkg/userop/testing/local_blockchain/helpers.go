package local_blockchain

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func SendNative(ctx context.Context, node *EthNode, from, to Account, fundAmount *big.Int) error {
	chainID, err := node.Client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	nonce, err := node.Client.PendingNonceAt(ctx, from.Address)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	gasLimit := uint64(21000)
	suggestedGasTipCap, err := node.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &to.Address,
		Value:     fundAmount,
		GasFeeCap: suggestedGasTipCap,
		GasTipCap: suggestedGasTipCap,
		Gas:       gasLimit,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), from.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	err = node.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	_, err = WaitMined(ctx, node, signedTx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}

	return nil
}

func WaitMined(ctx context.Context, node *EthNode, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(50 * time.Millisecond)
	defer queryTicker.Stop()

	for {
		receipt, err := node.Client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func setupAccounts(ctx context.Context, t *testing.T, userOpClient userop.Client, node *EthNode) (Account, Account, common.Address) {
	eoaBalance := decimal.NewFromFloat(2e18 /* 2 ETH */).BigInt()
	eoa, err := NewAccountWithBalance(ctx, eoaBalance, node) // EOA = Externally Owned Account
	require.NoError(t, err, "failed to create EOA")
	slog.Info("eoa", "address", eoa.Address)

	swAddress, err := userOpClient.GetAccountAddress(ctx, eoa.Address, decimal.Zero)
	sender := Account{Address: swAddress}
	require.NoError(t, err, "failed to compute sender account address")
	slog.Info("sender", "address", sender.Address)
	err = SendNative(ctx, node, eoa, sender, big.NewInt(1e18 /* 1 ETH */))
	require.NoError(t, err, "failed to fund sender account")

	receiver, err := NewAccount(ctx, node)
	require.NoError(t, err, "failed to create receiver account")
	slog.Info("receiver", "address", receiver.Address)

	return eoa, receiver, swAddress
}
