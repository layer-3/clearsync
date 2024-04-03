package public_blockchain

import (
	"context"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/itoken"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
)

var (
	config               = exampleConfig
	signer               = exampleSigner
	walletDeploymentOpts = exampleWalletDeploymentOpts
	gasLimitOverrides    = exampleGasLimitOverrides

	swartWalletIndex = decimal.Zero
	owner            = common.HexToAddress("0x2185da3337cad307fd48dFDabA6D4C66A9fD2c71")
	smartWallet      = common.HexToAddress("0x69b36b0Cb89b1666d85Ed4fF48243730E9c53405")
	receiver         = common.HexToAddress("0x2185da3337cad307fd48dFDabA6D4C66A9fD2c71")
	token            = common.HexToAddress("0x18e73A5333984549484348A94f4D219f4faB7b81") // Duckies
	amount           = decimal.RequireFromString("1000000000000000000")                  // wei

	ducklingsGame    = common.HexToAddress("0xb66bf78cad7cbab51988ddc792652cbabdff7675") // Duckies
	ducklingsGameABI = `[{
		  "inputs": [{
	        "internalType": "uint8",
	        "name": "size",
	        "type": "uint8"
	      }],
		  "name": "mintPack",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		}]`
)

func TestPublicRPC(t *testing.T) {
	t.Skip("this test is for manual execution only")
	setLogLevel(slog.LevelDebug)

	// create smartWallet client (with specific Wallet and Paymaster types)
	client, err := userop.NewClient(config)
	require.NoError(t, err, "failed to create userop client")

	// calculate smart wallet address
	walletAddress, err := client.GetAccountAddress(context.Background(), owner, swartWalletIndex)
	require.NoError(t, err, "failed to calculate smart wallet address")
	slog.Debug("wallet address", "address", walletAddress)

	// You can send native tokens to any address.
	transferNative, err := newTransferNativeCall(receiver, amount.BigInt())
	require.NoError(t, err, "failed to build transfer native call")
	err = send(t, client, smartWallet, smart_wallet.Calls{transferNative})
	require.NoError(t, err, "failed to send transfer native call")

	// NOTE: prior to using PimlicoERC20Paymaster, make sure to approve the
	// paymaster contract to spend your fee token.
	approve, err := newApproveCall(t, token, receiver, amount)
	require.NoError(t, err, "failed to build approve call")
	err = send(t, client, smartWallet, smart_wallet.Calls{approve})
	require.NoError(t, err, "failed to send approve call")

	// Now this call can be paid with ERC20 tokens using PimlicoERC20Paymaster.
	transferERC20, err := newTransferERC20Call(t, token, receiver, amount)
	require.NoError(t, err, "failed to build transfer ERC20 call")
	err = send(t, client, smartWallet, smart_wallet.Calls{transferERC20})
	require.NoError(t, err, "failed to send transfer ERC20 call")

	// You can also submit several calls in a single userOp.
	mintPrice := decimal.RequireFromString("5000000000") // 50 duckies for 1 Duckling
	approveToGame, err := newApproveCall(t, token, receiver, mintPrice)
	require.NoError(t, err, "failed to build approve to game call")

	mintPack, err := newCallFromABI(t, ducklingsGame, ducklingsGameABI, big.NewInt(0), "mintPack", uint8(1))
	require.NoError(t, err, "failed to build mint pack call")
	err = send(t, client, smartWallet, smart_wallet.Calls{approveToGame, mintPack})
	require.NoError(t, err, "failed to send mint pack call")
}

func setLogLevel(level slog.Level) {
	lvl := new(slog.LevelVar)
	lvl.Set(level)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	slog.SetDefault(logger)
}

// Encodes an `approve` call to the `token` contract, approving `amount` to be spent by `spender`.
func newApproveCall(t *testing.T, token, spender common.Address, amount decimal.Decimal) (smart_wallet.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	require.NoError(t, err, "failed to parse ERC20 ABI")

	callData, err := erc20.Pack("approve", spender, amount.BigInt())
	require.NoError(t, err, "failed to pack approve data")

	return smart_wallet.Call{
		To:       token,
		Value:    big.NewInt(0),
		CallData: callData,
	}, nil
}

// Encodes a `transfer` call of a native token, transferring `amount` to `receiver`.
func newTransferNativeCall(receiver common.Address, amount *big.Int) (smart_wallet.Call, error) {
	return smart_wallet.Call{
		To:    receiver,
		Value: amount,
	}, nil
}

// Encodes a `transfer` call to the `token` contract, transferring `amount` to `receiver`.
func newTransferERC20Call(t *testing.T, token, receiver common.Address, amount decimal.Decimal) (smart_wallet.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	require.NoError(t, err, "failed to parse ERC20 ABI")

	callData, err := erc20.Pack("transfer", receiver, amount.BigInt())
	require.NoError(t, err, "failed to pack transfer data")

	return smart_wallet.Call{
		To:       token,
		Value:    big.NewInt(0),
		CallData: callData,
	}, nil
}

// Encodes a call to the `contract` with the given `value`, `method` and `args`.
func newCallFromABI(t *testing.T, contract common.Address, stringABI string, value *big.Int, method string, args ...interface{}) (smart_wallet.Call, error) {
	ABI, err := abi.JSON(strings.NewReader(stringABI))
	require.NoError(t, err, "failed to parse ABI")

	callData, err := ABI.Pack(method, args...)
	require.NoError(t, err, "failed to pack call data")

	return smart_wallet.Call{
		To:       contract,
		Value:    value,
		CallData: callData,
	}, nil
}

// Creates and sends the user operation.
// NOTE: when sending the first userOp from a Smart Wallet,
// `config.example.go/walletDeploymentOpts` must contain Smart Wallet owner EOA address and SW index (0 by default).
func send(t *testing.T, client userop.Client, smartWallet common.Address, calls smart_wallet.Calls) error {
	ctx := context.Background()

	overrides := &userop.Overrides{
		GasLimits: gasLimitOverrides,
	}
	op, err := client.NewUserOp(ctx, smartWallet, signer, calls, walletDeploymentOpts, overrides)
	require.NoError(t, err, "failed to create user operation")

	b, err := op.MarshalJSON()
	require.NoError(t, err, "failed to marshal user operation")
	slog.Debug("sending user operation", "op", string(b))

	waitForUserOp, err := client.SendUserOp(ctx, op)
	require.NoError(t, err, "failed to send user operation")

	userOpReceipt := <-waitForUserOp

	slog.Info("user operation verified", "userOpReceipt", userOpReceipt)
	require.True(t, userOpReceipt.Success)

	return nil
}
