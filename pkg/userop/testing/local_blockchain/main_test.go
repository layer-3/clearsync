package local_blockchain

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestSimulatedRPC(t *testing.T) {
	ctx := context.Background()

	// 1. Start a local Ethereum node
	for i := 0; i < 3; i++ { // starting multiple nodes to test reusing existing nodes
		ethNode := NewEthNode(ctx, t)
		slog.Info("connecting to Ethereum node", "rpcURL", ethNode.LocalURL.String())
	}
	node := NewEthNode(ctx, t)
	slog.Info("connecting to Ethereum node", "rpcURL", node.LocalURL.String())

	// 2. Deploy the required contracts
	addresses := SetupContracts(ctx, t, node)

	// 3. Start the bundler
	for i := 0; i < 3; i++ { // starting multiple bundlers to test reusing existing bundlers
		bundlerURL := NewBundler(ctx, t, node, addresses.EntryPoint)
		slog.Info("connecting to bundler", "bundlerURL", bundlerURL.String())
	}
	bundlerURL := *NewBundler(ctx, t, node, addresses.EntryPoint)

	// 4. Build client
	client := buildClient(t, node.LocalURL, bundlerURL, addresses)

	// 5. Create and fund smart account
	eoaBalance := decimal.NewFromFloat(2e18 /* 2 ETH */).BigInt()
	eoa, err := NewAccountWithBalance(ctx, eoaBalance, node) // EOA = Externally Owned Account
	require.NoError(t, err, "failed to create EOA")
	slog.Info("eoa", "address", eoa.Address)

	senderAddress, err := client.GetAccountAddress(ctx, eoa.Address, decimal.Zero)
	sender := Account{Address: senderAddress}
	require.NoError(t, err, "failed to compute sender account address")
	slog.Info("sender", "address", sender.Address)
	sendNative(ctx, t, node, eoa, sender, decimal.NewFromFloat(1e18 /* 1 ETH */))

	receiver, err := NewAccount(ctx, node)
	require.NoError(t, err, "failed to create receiver account")
	slog.Info("receiver", "address", receiver.Address)

	// 6. Submit user operation
	signer := userop.SignerForKernel(signer.NewLocalSigner(eoa.PrivateKey))
	transferAmount := decimal.NewFromInt(1 /* 1 wei */).BigInt()
	calls := []userop.Call{{To: receiver.Address, Value: transferAmount}}
	params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: eoa.Address}
	op, err := client.NewUserOp(ctx, sender.Address, signer, calls, params, nil)
	require.NoError(t, err, "failed to create new user operation")
	slog.Info("ready to send", "userop", op)

	done, err := client.SendUserOp(ctx, op)
	require.NoError(t, err, "failed to send user operation")

	receipt := <-done
	slog.Info("transaction mined", "receipt", receipt)
	require.True(t, receipt.Success)

	receiverBalance, err := node.Client.BalanceAt(ctx, receiver.Address, nil)
	require.NoError(t, err, "failed to fetch receiver new balance")
	require.Equal(t, transferAmount, receiverBalance, "new balance should equal the transfer amount")
}

func sendNative(ctx context.Context, t *testing.T, node *EthNode, from, to Account, fundAmount decimal.Decimal) {
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
