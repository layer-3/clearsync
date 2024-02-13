package userop

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point"
	"github.com/layer-3/clearsync/pkg/abi/simple_account_factory"
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

func getInitCode(factoryAddress, ownerAddress common.Address) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		factory, err := abi.JSON(strings.NewReader(simple_account_factory.SimpleAccountFactoryMetaData.ABI))
		if err != nil {
			return fmt.Errorf("failed to parse SimpleAccountFactory ABI: %w", err)
		}

		initCode, err := factory.Pack("deployCounterFactualAccount", ownerAddress, new(big.Int))
		if err != nil {
			return fmt.Errorf("failed to pack initCode: %w", err)
		}

		op.InitCode = initCode
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
	CallGasLimit         string `json:"callGasLimit"`
	VerificationGasLimit string `json:"verificationGasLimit"`
	PreVerificationGas   string `json:"preVerificationGas"`
}

func estimateUserOperationGas(bundlerRPC *rpc.Client, entryPoint common.Address) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		slog.Info("estimating gas")

		var est gasEstimate
		if err := bundlerRPC.CallContext(
			ctx,
			&est,
			"eth_estimateUserOperationGas",
			op,
			entryPoint,
		); err != nil {
			return fmt.Errorf("error estimating gas: %w", err)
		}

		preVerificationGas, err := hexutil.DecodeBig(est.PreVerificationGas)
		if err != nil {
			return fmt.Errorf("failed to parse preVerificationGas: %w (got '%s')", err, est.PreVerificationGas)
		}
		verificationGasLimit, err := hexutil.DecodeBig(est.VerificationGasLimit)
		if err != nil {
			return fmt.Errorf("failed to parse verificationGasLimit: %w (got '%s')", err, est.VerificationGasLimit)
		}
		callGasLimit, err := hexutil.DecodeBig(est.CallGasLimit)
		if err != nil {
			return fmt.Errorf("failed to parse callGasLimit: %w (got '%s')", err, est.CallGasLimit)
		}

		op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0)
		op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
		op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)

		return nil
	}
}

func sign(signer Signer, entryPoint common.Address, chainID *big.Int) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		signature, err := signer(*op, entryPoint, chainID)
		if err != nil {
			return fmt.Errorf("failed to sign user operation: %w", err)
		}

		op.Signature = signature
		return nil
	}
}
