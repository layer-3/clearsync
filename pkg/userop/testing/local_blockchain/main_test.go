package local_blockchain

import (
	"context"
	"log/slog"
	"math/big"
	"os"
	"testing"

	"github.com/layer-3/clearsync/pkg/artifacts/test_erc20"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setLogLevel(level slog.Level) {
	lvl := new(slog.LevelVar)
	lvl.Set(level)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	slog.SetDefault(logger)
}

func TestSimulatedRPC(t *testing.T) {
	setLogLevel(slog.LevelDebug)
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
		bundler := NewBundler(ctx, t, node, addresses.EntryPoint)
		slog.Info("connecting to bundler", "bundlerURL", bundler.LocalURL.String())
	}
	bundler := NewBundler(ctx, t, node, addresses.EntryPoint)

	// 4. Build client
	client := BuildClient(t, node.LocalURL, bundler.LocalURL, addresses, userop.PaymasterConfig{
		Type: &userop.PaymasterDisabled,
	})

	// 5. Create and fund smart account
	eoa, receiver, swAddress := setupAccounts(ctx, t, client, node)

	// 6. Submit user operation
	signer := userop.SignerForKernel(signer.NewLocalSigner(eoa.PrivateKey))
	transferAmount := decimal.NewFromInt(1 /* 1 wei */).BigInt()
	calls := smart_wallet.Calls{{To: receiver.Address, Value: transferAmount}}
	params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: eoa.Address}
	op, err := client.NewUserOp(ctx, swAddress, signer, calls, params, nil)
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

func TestSimulatedPaymaster(t *testing.T) {
	setLogLevel(slog.LevelDebug)
	ctx := context.Background()

	node := NewEthNode(ctx, t)
	slog.Info("connecting to Ethereum node", "rpcURL", node.LocalURL.String())

	// Deploy the required contracts
	addresses := SetupContracts(ctx, t, node)

	bundler := NewBundler(ctx, t, node, addresses.EntryPoint)
	balance := decimal.NewFromFloat(50e18 /* 50 ETH */).BigInt()
	owner, err := NewAccountWithBalance(ctx, balance, node)
	require.NoError(t, err, "failed to create owner account")

	// Deploy ERC20 token
	token, _, _, err := test_erc20.DeployTestERC20(owner.TransactOpts, node.Client, "TestToken", "TT", 6, big.NewInt(1000*1e6)) // 1000 tokens
	require.NoError(t, err, "failed to deploy ERC20 token")

	// Deploy paymaster
	paymasterURL := SetupPaymaster(ctx, t, node, bundler, token)

	// Build client
	client := BuildClient(t, node.LocalURL, bundler.LocalURL, addresses, userop.PaymasterConfig{
		Type:    &userop.PaymasterPimlicoVerifying,
		URL:     paymasterURL.String(),
		Address: addresses.Paymaster,
	})

	// Create smart account
	eoa, err := NewAccount(ctx, node) // EOA without funds
	require.NoError(t, err, "failed to create EOA")
	slog.Info("eoa", "address", eoa.Address)

	swAddress, err := client.GetAccountAddress(ctx, eoa.Address, decimal.Zero)
	sw := Account{Address: swAddress}
	require.NoError(t, err, "failed to compute sender account address")
	slog.Info("sender", "address", sw.Address)

	// Send userop
	signer := userop.SignerForKernel(signer.NewLocalSigner(eoa.PrivateKey))
	transferAmount := decimal.NewFromInt(0 /* 0 wei */).BigInt()
	calls := smart_wallet.Calls{{To: sw.Address, Value: transferAmount}}
	params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: eoa.Address}
	op, err := client.NewUserOp(ctx, sw.Address, signer, calls, params, nil)
	assert.NoError(t, err, "failed to create new user operation")
	slog.Info("ready to send", "userop", op)

	done, err := client.SendUserOp(ctx, op)
	assert.NoError(t, err, "failed to send user operation")

	receipt := <-done
	slog.Info("transaction mined", "receipt", receipt)
	assert.True(t, receipt.Success)
}
