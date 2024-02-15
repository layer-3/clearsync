package userop

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

var createKernelAccountABI = `[{
  inputs: [
    {
      internalType: "address",
      name: "moduleSetupContract",
      type: "address"
    },
    {
      internalType: "bytes",
      name: "moduleSetupData",
      type: "bytes"
    },
    {
      internalType: "uint256",
      name: "index",
      type: "uint256"
    }
  ],
  name: "deployCounterFactualAccount",
  outputs: [{
    internalType: "address",
    name: "proxy",
    type: "address"
  }],
  stateMutability: "nonpayable",
  type: "function"
}]`

// kernelInitABI is the init ABI, used to initialise kernel account
var kernelInitABI = `[{
  inputs: [
    {
      internalType: "contract IKernelValidator",
      name: "_defaultValidator",
      type: "address"
    },
    {
      internalType: "bytes",
      name: "_data",
      type: "bytes"
    }
  ],
  name: "initialize",
  outputs: [],
  stateMutability: "payable",
  type: "function"
}]`

func getKernelInitCode(index int64, factory, accountLogic, ecdsaValidator, owner common.Address) middleware {
	initABI, err := abi.JSON(strings.NewReader(kernelInitABI))
	if err != nil {
		panic(err)
	}

	createAccountABI, err := abi.JSON(strings.NewReader(createKernelAccountABI))
	if err != nil {
		panic(err)
	}

	return func(ctx context.Context, op *UserOperation) error {
		initData, err := initABI.Pack("initialize", ecdsaValidator, owner)
		if err != nil {
			return err
		}

		initCode, err := createAccountABI.Pack("createAccount", accountLogic, initData, index)
		if err != nil {
			return err
		}

		op.InitCode = initCode
		return nil
	}
}

var createBiconomyAccountABI = `[{
  inputs: [
    {
      internalType: "address",
      name: "moduleSetupContract",
      type: "address"
    },
    {
      internalType: "bytes",
      name: "moduleSetupData",
      type: "bytes"
    },
    {
      internalType: "uint256",
      name: "index",
      type: "uint256"
    }
  ],
  name: "deployCounterFactualAccount",
  outputs: [{
    internalType: "address",
    name: "proxy",
    type: "address"
  }],
  stateMutability: "nonpayable",
  type: "function"
}]`

var biconomyInitABI = `[
  {
    inputs: [
      {
        internalType: "address",
        name: "handler",
        type: "address"
      },
      {
        internalType: "address",
        name: "moduleSetupContract",
        type: "address"
      },
      {
        internalType: "bytes",
        name: "moduleSetupData",
        type: "bytes"
      }
    ],
    name: "init",
    outputs: [{
      internalType: "address",
      name: "",
      type: "address"
    }],
    stateMutability: "nonpayable",
    type: "function"
  },
  {
    inputs: [{
      internalType: "address",
      name: "eoaOwner",
      type: "address"
    }],
    name: "initForSmartAccount",
    outputs: [{
      internalType: "address",
      name: "",
      type: "address"
    }],
    stateMutability: "nonpayable",
    type: "function"
  },
]`

func getBiconomyInitCode(index int64, factory, accountLogic, ecdsaValidator, owner common.Address) middleware {
	initABI, err := abi.JSON(strings.NewReader(biconomyInitABI))
	if err != nil {
		panic(err)
	}

	createAccountABI, err := abi.JSON(strings.NewReader(createBiconomyAccountABI))
	if err != nil {
		panic(err)
	}

	return func(ctx context.Context, op *UserOperation) error {
		ecdsaOwnershipInitData, err := initABI.Pack("initForSmartAccount", owner)
		if err != nil {
			return err
		}

		initCode, err := createAccountABI.Pack("createAccount", ecdsaValidator, ecdsaOwnershipInitData, index)
		if err != nil {
			return err
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

		maxBlockGas := decimal.NewFromInt(30_000_000)
		op.CallGasLimit = maxBlockGas
		op.VerificationGasLimit = maxBlockGas
		op.PreVerificationGas = maxBlockGas

		return nil
	}
}

func getPaymasterData(paymasterRPC *rpc.Client, paymasterCtx map[string]any, entryPoint common.Address) middleware {
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
			entryPoint,
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
	// depending on provider, any of the following types can be received here: string, int
	CallGasLimit         any `json:"callGasLimit"`
	VerificationGasLimit any `json:"verificationGasLimit"`
	PreVerificationGas   any `json:"preVerificationGas"`

	PaymasterAndData string `json:"paymasterAndData,omitempty"`
}

// func (est *gasEstimate) UnmarshalJSON(data []byte) error {
// 	type Alias gasEstimate
// 	aux := &struct {
// 		*Alias
// 	}{
// 		Alias: (*Alias)(est),
// 	}
//
// 	fmt.Println("data", string(data))
// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (est gasEstimate) convert() (preVerificationGas *big.Int, verificationGasLimit *big.Int, callGasLimit *big.Int, err error) {
	preVerificationGas, err = est.fromAny(est.PreVerificationGas)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("preVerificationGas: %w", err)
	}
	verificationGasLimit, err = est.fromAny(est.VerificationGasLimit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("verificationGasLimit: %w", err)
	}
	callGasLimit, err = est.fromAny(est.CallGasLimit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("callGasLimit: %w", err)
	}

	return preVerificationGas,
		verificationGasLimit,
		callGasLimit,
		nil
}

func (est gasEstimate) fromAny(a any) (*big.Int, error) {
	switch v := a.(type) {
	case string:
		n, err := strconv.ParseInt(v, 16, 64)
		if err != nil { // it IS hexadecimal
			nBig, err := hexutil.DecodeBig(v)
			if err != nil {
				return nil, fmt.Errorf("failed to parse: %w (got '%s')", err, v)
			}
			return nBig, nil
		}

		return new(big.Int).SetInt64(n), nil
	case int64:
		return new(big.Int).SetInt64(v), nil
	case float64:
		return new(big.Int).SetInt64(int64(v)), nil
	default:
		return nil, fmt.Errorf("unexpected type: %T", v)
	}
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

		fmt.Println("hash =", op.UserOpHash(entryPoint, chainID).String())
		fmt.Println("array =", op.ToArray())

		b, err := op.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println("array =", string(b))
		return nil
	}
}
