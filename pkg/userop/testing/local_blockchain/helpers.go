package local_blockchain

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func SendNative(ctx context.Context, t *testing.T, node *EthNode, from, to Account, fundAmount decimal.Decimal) {
	chainID, err := node.Client.ChainID(ctx)
	require.NoError(t, err, "Error getting chain ID")

	nonce, err := node.Client.PendingNonceAt(ctx, from.Address)
	require.NoError(t, err, "Error getting nonce")

	gasLimit := uint64(21000)
	gasPrice, err := node.Client.SuggestGasPrice(ctx)
	require.NoError(t, err, "Error suggesting gas price")

	tx := types.NewTransaction(nonce, to.Address, fundAmount.BigInt(), gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), from.PrivateKey)
	require.NoError(t, err, "Error signing transaction")

	err = node.Client.SendTransaction(ctx, signedTx)
	require.NoError(t, err, "Error sending transaction")

	_, err = waitMined(ctx, node, signedTx)
	require.NoError(t, err, "Error waiting for transaction to be mined")
}

// waitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func waitMined(ctx context.Context, node *EthNode, tx *types.Transaction) (*types.Receipt, error) {
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
