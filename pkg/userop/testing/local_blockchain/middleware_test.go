package local_blockchain

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGasLimitOverrides(t *testing.T) {
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
		bundlerURL := NewBundler(ctx, t, node, addresses.EntryPoint)
		slog.Info("connecting to bundler", "bundlerURL", bundlerURL.String())
	}
	bundlerURL := *NewBundler(ctx, t, node, addresses.EntryPoint)

	overrides := &userop.Overrides{
		GasLimits: &userop.GasLimitOverrides{
			CallGasLimit:         big.NewInt(420e3),
			VerificationGasLimit: big.NewInt(420e3),
			PreVerificationGas:   big.NewInt(420e3),
		},
	}

	t.Run("overrides persist if no paymaster", func(t *testing.T) {
		// 4. Build client
		client := buildClient(t, node.LocalURL, bundlerURL, addresses)

		// 5. Create and fund smart account
		eoa, receiver, swAddress := setupAccounts(ctx, t, client, node)

		// 6. Create user operation
		signer := userop.SignerForKernel(signer.NewLocalSigner(eoa.PrivateKey))
		transferAmount := decimal.NewFromInt(1 /* 1 wei */).BigInt()
		calls := smart_wallet.Calls{{To: receiver.Address, Value: transferAmount}}
		params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: eoa.Address}
		op, err := client.NewUserOp(ctx, swAddress, signer, calls, params, overrides)
		require.NoError(t, err, "failed to create new user operation")

		// 7. Overrides must be persisted
		require.Equal(t, overrides.GasLimits.CallGasLimit.Cmp(op.CallGasLimit.BigInt()), 0, "call gas limit override must persist")
		require.Equal(t, overrides.GasLimits.VerificationGasLimit.Cmp(op.VerificationGasLimit.BigInt()), 0, "verification gas limit override must persist")
		require.Equal(t, overrides.GasLimits.PreVerificationGas.Cmp(op.PreVerificationGas.BigInt()), 0, "pre-verification gas override must persist")
	})

	t.Run("overrides persist if ERC20 paymaster", func(t *testing.T) {
		// 4. Build client
		config := defaultClientConfig(t, node.LocalURL, bundlerURL, addresses)
		config.Paymaster.Type = &userop.PaymasterPimlicoERC20
		config.Paymaster.Address = addresses.EntryPoint // any address with code
		config.Paymaster.PimlicoERC20 = userop.PimlicoERC20Config{
			VerificationGasOverhead: decimal.Zero,
		}

		client, err := userop.NewClient(config)
		require.NoError(t, err)

		// 5. Create and fund smart account
		eoa, receiver, swAddress := setupAccounts(ctx, t, client, node)

		// 6. Create user operation
		signer := userop.SignerForKernel(signer.NewLocalSigner(eoa.PrivateKey))
		transferAmount := decimal.NewFromInt(1 /* 1 wei */).BigInt()
		calls := smart_wallet.Calls{{To: receiver.Address, Value: transferAmount}}
		params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: eoa.Address}
		op, err := client.NewUserOp(ctx, swAddress, signer, calls, params, overrides)
		require.NoError(t, err, "failed to create new user operation")

		// 7. Overrides must be persisted
		require.Equal(t, overrides.GasLimits.CallGasLimit.Cmp(op.CallGasLimit.BigInt()), 0, "call gas limit must persist")
		require.Equal(t, overrides.GasLimits.VerificationGasLimit.Cmp(op.VerificationGasLimit.BigInt()), 0, "verification gas limit must persist")
		require.Equal(t, overrides.GasLimits.PreVerificationGas.Cmp(op.PreVerificationGas.BigInt()), 0, "pre-verification gas must persist")
	})

	t.Run("overrides are overwritten with estimation if verifying paymaster", func(t *testing.T) {

		// 4. Mock server
		responseLimit := int64(69000)

		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				response := userop.GasEstimate{
					CallGasLimit:         responseLimit,
					VerificationGasLimit: responseLimit,
					PreVerificationGas:   responseLimit,
					PaymasterAndData:     "0xdeadbeef",
				}
				jsonResponse := map[string]interface{}{
					"jsonrpc": "2.0",
					"result":  response,
					"id":      1,
				}
				jsonResponseBytes, err := json.Marshal(jsonResponse)
				require.NoError(t, err)
				w.Write(jsonResponseBytes)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}))
		defer testServer.Close()

		// 5. Build client
		config := defaultClientConfig(t, node.LocalURL, bundlerURL, addresses)
		config.Paymaster.Type = &userop.PaymasterPimlicoVerifying
		config.Paymaster.URL = testServer.URL

		client, err := userop.NewClient(config)
		require.NoError(t, err)

		// 6. Create and fund smart account
		eoa, receiver, swAddress := setupAccounts(ctx, t, client, node)

		// 7. Create user operation
		signer := userop.SignerForKernel(signer.NewLocalSigner(eoa.PrivateKey))
		transferAmount := decimal.NewFromInt(1 /* 1 wei */).BigInt()
		calls := smart_wallet.Calls{{To: receiver.Address, Value: transferAmount}}
		params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: eoa.Address}
		op, err := client.NewUserOp(ctx, swAddress, signer, calls, params, overrides)
		require.NoError(t, err, "failed to create new user operation")

		// 8. Overrides must be overwritten
		require.Equal(t, op.CallGasLimit.Cmp(decimal.NewFromInt(responseLimit)), 0, "call gas limit override must be overwritten")
		require.Equal(t, op.VerificationGasLimit.Cmp(decimal.NewFromInt(responseLimit)), 0, "verification gas limit override must be overwritten")
		require.Equal(t, op.PreVerificationGas.Cmp(decimal.NewFromInt(responseLimit)), 0, "pre-verification gas override must be overwritten")
	})
}
