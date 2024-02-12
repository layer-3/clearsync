package userop

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point"
)

type middleware func(ctx context.Context, op *UserOperation) error

func getNonce(entryPoint *entry_point.EntryPoint) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		slog.Info("getting nonce")
		key := new(big.Int)
		nonce, err := entryPoint.GetNonce(nil, op.Sender, key)
		if err != nil {
			return err
		}

		op.Nonce = decimal.NewFromBigInt(nonce, 0)
		return nil
	}
}

// TODO: uncomment when smart wallet deployment support is added
// func getInitCode(client *ethclient.Client) middleware {
// 	return func(ctx context.Context, op *UserOperation) error {
// 		slog.Info("Getting init code")
// 		return client.Client().CallContext(ctx, &op.InitCode, "eth_getCode")
// 	}
// }

func getGasPrice(providerRPC *ethclient.Client) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		slog.Info("getting gas price")

		// Get the latest block to read its baseFee
		block, err := providerRPC.BlockByNumber(ctx, nil)
		if err != nil {
			return err
		}
		feePerGas := block.BaseFee()

		var maxPriorityFeePerGasStr string
		if err := providerRPC.Client().CallContext(
			ctx,
			&maxPriorityFeePerGasStr,
			"eth_maxPriorityFeePerGas",
		); err != nil {
			return err
		}

		maxPriorityFeePerGas, ok := new(big.Int).SetString(maxPriorityFeePerGasStr, 0)
		if !ok {
			return fmt.Errorf("failed to parse maxPriorityFeePerGas: %s", maxPriorityFeePerGasStr)
		}

		// Increase maxPriorityFeePerGas by 13%
		tip := new(big.Int).Div(maxPriorityFeePerGas, big.NewInt(100))
		tip.Mul(tip, big.NewInt(13))

		// maxPriorityFeePerGas = maxPriorityFeePerGas + tip
		maxPriorityFeePerGas.Add(maxPriorityFeePerGas, tip)

		// Calculate maxFeePerGas
		maxFeePerGas := new(big.Int).Mul(feePerGas, big.NewInt(2))
		maxFeePerGas.Add(maxFeePerGas, maxPriorityFeePerGas)

		op.MaxFeePerGas = decimal.NewFromBigInt(maxFeePerGas, 0)
		op.MaxPriorityFeePerGas = decimal.NewFromBigInt(maxPriorityFeePerGas, 0)

		return nil
	}
}

// gasEstimate holds gas estimates for a user operation.
type gasEstimate struct {
	PreVerificationGas   *decimal.Decimal `json:"preVerificationGas"`
	VerificationGasLimit *decimal.Decimal `json:"verificationGasLimit"`
	CallGasLimit         *decimal.Decimal `json:"callGasLimit"`
	VerificationGas      *decimal.Decimal `json:"verificationGas"` // TODO: remove this with EntryPoint v0.7
}

func estimateUserOperationGas(bundlerRPC *rpc.Client, entryPoint common.Address) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		slog.Info("estimating gas")

		var est gasEstimate
		if err := bundlerRPC.CallContext(
			ctx,
			&est,
			"eth_estimateUserOperationGas",
			op.Marshal(),
			entryPoint,
		); err != nil {
			return fmt.Errorf("error estimating gas: %w", err)
		}

		op.CallGasLimit = *est.CallGasLimit
		if est.VerificationGasLimit != nil {
			op.VerificationGasLimit = *est.VerificationGasLimit
		} else {
			// Fallback to verificationGas if verificationGasLimit is not available.
			op.VerificationGasLimit = *est.VerificationGas
		}
		op.PreVerificationGas = *est.PreVerificationGas

		return nil
	}
}
