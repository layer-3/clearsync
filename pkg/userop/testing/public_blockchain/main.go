package public_blockchain

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

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
	amount           = decimal.RequireFromString("1000")                                 // wei

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

func main() {
	setLogLevel(slog.LevelInfo)

	// create smartWallet client (with specific Wallet and Paymaster types)
	client, err := userop.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	// calculate smart wallet address
	walletAddress, err := client.GetAccountAddress(context.Background(), owner, swartWalletIndex)
	if err != nil {
		panic(fmt.Errorf("failed to get wallet address: %w", err))
	}
	slog.Debug("wallet address", "address", walletAddress)

	// You can send native tokens to any address.
	transferNative, err := newTransferNativeCall(receiver, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to build transfer native call: %w", err))
	}
	if err := send(client, smartWallet, smart_wallet.Calls{transferNative}); err != nil {
		panic(err)
	}

	// NOTE: prior to using PimlicoERC20Paymaster, make sure to approve the
	// paymaster contract to spend your fee token.
	approve, err := newApproveCall(token, receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build approve call: %w", err))
	}
	if err := send(client, smartWallet, smart_wallet.Calls{approve}); err != nil {
		panic(err)
	}

	// Now this call can be paid with ERC20 tokens using PimlicoERC20Paymaster.
	transferERC20, err := newTransferERC20Call(token, receiver, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build transfer erc20 call: %w", err))
	}
	if err := send(client, smartWallet, smart_wallet.Calls{transferERC20}); err != nil {
		panic(err)
	}

	// You can also submit several calls in a single userOp.
	mintPrice := decimal.RequireFromString("5000000000") // 50 duckies for 1 Duckling
	approveToGame, err := newApproveCall(token, ducklingsGame, mintPrice)
	if err != nil {
		panic(fmt.Errorf("failed to build approve to game call: %w", err))
	}

	mintPack, err := newCallFromABI(ducklingsGame, ducklingsGameABI, big.NewInt(0), "mintPack", uint8(1))
	if err != nil {
		panic(fmt.Errorf("failed to build mint pack call: %w", err))
	}

	if err := send(client, smartWallet, smart_wallet.Calls{approveToGame, mintPack}); err != nil {
		panic(err)
	}
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
func newApproveCall(token, spender common.Address, amount decimal.Decimal) (smart_wallet.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ERC20 ABI: %w", err))
	}

	callData, err := erc20.Pack("approve", spender, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to pack transfer data: %w", err))
	}

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
func newTransferERC20Call(token, receiver common.Address, amount decimal.Decimal) (smart_wallet.Call, error) {
	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ERC20 ABI: %w", err))
	}

	callData, err := erc20.Pack("transfer", receiver, amount.BigInt())
	if err != nil {
		panic(fmt.Errorf("failed to pack transfer data: %w", err))
	}

	return smart_wallet.Call{
		To:       token,
		Value:    big.NewInt(0),
		CallData: callData,
	}, nil
}

// Encodes a call to the `contract` with the given `value`, `method` and `args`.
func newCallFromABI(contract common.Address, stringABI string, value *big.Int, method string, args ...interface{}) (smart_wallet.Call, error) {
	ABI, err := abi.JSON(strings.NewReader(stringABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ABI: %w", err))
	}

	callData, err := ABI.Pack(method, args...)
	if err != nil {
		panic(fmt.Errorf("failed to pack call data: %w", err))
	}

	return smart_wallet.Call{
		To:       contract,
		Value:    value,
		CallData: callData,
	}, nil
}

// Creates and sends the user operation.
// NOTE: when sending the first userOp from a Smart Wallet,
// `config.example.go/walletDeploymentOpts` must contain Smart Wallet owner EOA address and SW index (0 by default).
func send(client userop.Client, smartWallet common.Address, calls smart_wallet.Calls) error {
	ctx := context.Background()

	op, err := client.NewUserOp(ctx, smartWallet, signer, calls, walletDeploymentOpts, gasLimitOverrides)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

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
