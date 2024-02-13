package userop

import (
	"context"
	"encoding/json"
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

func getPaymasterData(paymasterRPC *rpc.Client, paymasterCtx map[string]any, _ common.Address) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		opModified := struct {
			Sender               common.Address  `json:"sender"`
			Nonce                decimal.Decimal `json:"nonce"`
			InitCode             string          `json:"initCode"`
			CallData             string          `json:"callData"`
			CallGasLimit         decimal.Decimal `json:"callGasLimit"`
			VerificationGasLimit decimal.Decimal `json:"verificationGasLimit"`
			PreVerificationGas   decimal.Decimal `json:"preVerificationGas"`
			MaxFeePerGas         decimal.Decimal `json:"maxFeePerGas"`
			MaxPriorityFeePerGas decimal.Decimal `json:"maxPriorityFeePerGas"`
			PaymasterAndData     string          `json:"paymasterAndData"`
			Signature            string          `json:"signature,omitempty"`
		}{
			Sender:               op.Sender,
			Nonce:                op.Nonce,
			InitCode:             hexutil.Encode(op.InitCode),
			CallData:             hexutil.Encode(op.CallData),
			CallGasLimit:         op.CallGasLimit,
			VerificationGasLimit: op.VerificationGasLimit,
			PreVerificationGas:   op.PreVerificationGas,
			MaxFeePerGas:         op.MaxFeePerGas,
			MaxPriorityFeePerGas: op.MaxPriorityFeePerGas,
			PaymasterAndData:     hexutil.Encode(op.PaymasterAndData),
			Signature:            hexutil.Encode(op.Signature),
		}

		var est gasEstimate
		if err := paymasterRPC.CallContext(
			ctx,
			&est,
			"pm_sponsorUserOperation",
			opModified,
			paymasterCtx,
		); err != nil {
			return fmt.Errorf("failed to call pm_sponsorUserOperation: %v", err)
		}

		callGasLimit, verificationGasLimit, preVerificationGas, err := est.convert()
		if err != nil {
			return fmt.Errorf("failed to convert gas estimates: %w", err)
		}

		op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0)
		op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
		op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)

		paymasterAndData, err := hexutil.Decode(est.PaymasterAndData)
		if err != nil {
			return fmt.Errorf("failed to decode paymasterAndData: %w", err)
		}
		op.PaymasterAndData = paymasterAndData

		return nil
	}
}

// gasEstimate holds gas estimates for a user operation.
type gasEstimate struct {
	CallGasLimit         decimal.Decimal `json:"callGasLimit"`
	VerificationGasLimit decimal.Decimal `json:"verificationGasLimit"`
	PreVerificationGas   decimal.Decimal `json:"preVerificationGas"`
	PaymasterAndData     string          `json:"paymasterAndData,omitempty"`
}

func (est *gasEstimate) UnmarshalJSON(data []byte) error {
	type Alias gasEstimate
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(est),
	}

	fmt.Println("data", string(data))
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (est gasEstimate) convert() (*big.Int, *big.Int, *big.Int, error) {
	// preVerificationGas, err := hexutil.DecodeBig(est.PreVerificationGas)
	// if err != nil {
	// 	return nil, nil, nil, fmt.Errorf("failed to parse preVerificationGas: %w (got '%s')", err, est.PreVerificationGas)
	// }
	// verificationGasLimit, err := hexutil.DecodeBig(est.VerificationGasLimit)
	// if err != nil {
	// 	return nil, nil, nil, fmt.Errorf("failed to parse verificationGasLimit: %w (got '%s')", err, est.VerificationGasLimit)
	// }
	// callGasLimit, err := hexutil.DecodeBig(est.CallGasLimit)
	// if err != nil {
	// 	return nil, nil, nil, fmt.Errorf("failed to parse callGasLimit: %w (got '%s')", err, est.CallGasLimit)
	// }

	// return preVerificationGas, verificationGasLimit, callGasLimit, nil
	return est.PreVerificationGas.BigInt(),
		est.VerificationGasLimit.BigInt(),
		est.CallGasLimit.BigInt(),
		nil
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

		callGasLimit, verificationGasLimit, preVerificationGas, err := est.convert()
		if err != nil {
			return fmt.Errorf("failed to convert gas estimates: %w", err)
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

		fmt.Println("to array", op.ToArray(), "hash", op.UserOpHash(entryPoint, chainID).String())
		return nil
	}
}
