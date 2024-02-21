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

const maxBlockGas = 30_000_000

type ctxKey int

const (
	ctxKeySigner ctxKey = iota
	ctxKeyOwner
	ctxKeyIndex
)

// middleware is a function that modifies a user operation.
// It is used to create a pipeline of operations to be executed
// to fill in the user operation with all the necessary data.
type middleware func(ctx context.Context, op *UserOperation) error

func getNonce(entryPoint *entry_point.EntryPoint) middleware {
	return func(_ context.Context, op *UserOperation) error {
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

type smartWalletInitCodeGenerator func(op UserOperation, owner common.Address, index decimal.Decimal) ([]byte, error)

func getInitCode(providerRPC *ethclient.Client, smartWalletConfig SmartWallet) (middleware, error) {
	var getInitCode smartWalletInitCodeGenerator
	switch typ := smartWalletConfig.Type; typ {
	case SmartWalletSimpleAccount:
		return nil, fmt.Errorf("%w: %s", ErrSmartWalletNotSupported, typ)
	case SmartWalletBiconomy: // not tested
		getInitCode = getBiconomyInitCode(smartWalletConfig.Factory, smartWalletConfig.ECDSAValidator)
	case SmartWalletKernel:
		getInitCode = getKernelInitCode(smartWalletConfig.Factory, smartWalletConfig.Logic, smartWalletConfig.ECDSAValidator)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", typ)
	}

	return func(ctx context.Context, op *UserOperation) error {
		var initCode []byte
		var err error

		// check if smart account is already deployed
		isDeployed, err := isAccountDeployed(providerRPC, op.Sender)
		if err != nil {
			return fmt.Errorf("failed to check if smart account is already deployed: %w", err)
		}

		// if sender == zeroAddress OR smart account is not deployed
		// then we need to calculate the init code
		if op.Sender == (common.Address{}) || !isDeployed {
			owner, ok := ctx.Value(ctxKeyOwner).(common.Address)
			if !ok {
				return fmt.Errorf("`owner` not found, but required in context to get init code")
			}

			index, ok := ctx.Value(ctxKeyIndex).(decimal.Decimal)
			if !ok {
				return fmt.Errorf("`index` not found, but required in context to get init code")
			}

			initCode, err = getInitCode(*op, owner, index)
			if err != nil {
				return fmt.Errorf("failed to get init code: %w", err)
			}
		}

		op.InitCode = initCode
		return nil
	}, nil
}

// getKernelInitCode returns a middleware that sets the init code
// for a Zerodev Kernel smart account. The init code deploys
// a smart account if it is not already deployed.
func getKernelInitCode(factory common.Address, accountLogic common.Address, ecdsaValidator common.Address) smartWalletInitCodeGenerator {
	initABI, err := abi.JSON(strings.NewReader(kernelInitABI))
	if err != nil {
		panic(err)
	}

	createAccountABI, err := abi.JSON(strings.NewReader(kernelDeployWalletABI))
	if err != nil {
		panic(err)
	}

	return func(op UserOperation, owner common.Address, index decimal.Decimal) ([]byte, error) {
		// Initialize Kernel Smart Account with default validation module and its calldata
		// see https://github.com/zerodevapp/kernel/blob/v2.3/src/abstract/KernelStorage.sol#L35
		initData, err := initABI.Pack("initialize", ecdsaValidator, owner.Bytes())
		if err != nil {
			panic(fmt.Errorf("failed to pack init data: %w", err))
		}

		// Deploy Kernel Smart Account by calling `factory.createAccount`
		// see https://github.com/zerodevapp/kernel/blob/v2.3/src/factory/KernelFactory.sol#L25
		callData, err := createAccountABI.Pack("createAccount", accountLogic, initData, index.BigInt())
		if err != nil {
			panic(fmt.Errorf("failed to pack createAccount data: %w", err))
		}

		// Pack factory address and deployment data for `CreateSender` in EntryPoint
		// see https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/SenderCreator.sol#L15
		initCode := make([]byte, len(factory)+len(callData))
		copy(initCode, factory.Bytes())
		copy(initCode[len(factory):], callData)

		slog.Debug("built initCode", "initCode", hexutil.Encode(initCode))

		return initCode, nil
	}
}

// getBiconomyInitCode returns a middleware that sets the init code for a Biconomy smart account.
// The init code deploys a smart account if it is not already deployed.
// !!! NOT TESTED extensively since we settled with the Zerodev Kernel smart account !!!
func getBiconomyInitCode(factory, ecdsaValidator common.Address) smartWalletInitCodeGenerator {
	initABI, err := abi.JSON(strings.NewReader(biconomyInitABI))
	if err != nil {
		panic(err)
	}

	createAccountABI, err := abi.JSON(strings.NewReader(biconomyDeployWalletABI))
	if err != nil {
		panic(err)
	}

	return func(op UserOperation, owner common.Address, index decimal.Decimal) ([]byte, error) {
		// Initialize SCW validation module with owner address
		// see https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/modules/EcdsaOwnershipRegistryModule.sol#L43
		ecdsaOwnershipInitData, err := initABI.Pack("initForSmartAccount", owner.Bytes())
		if err != nil {
			panic(fmt.Errorf("failed to pack init data: %w", err))
		}

		// Deploy Biconomy SCW by calling `factory.createAccount`
		// see https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/factory/SmartAccountFactory.sol#L112
		callData, err := createAccountABI.Pack("createAccount", ecdsaValidator, ecdsaOwnershipInitData, index.BigInt())
		if err != nil {
			panic(fmt.Errorf("failed to pack createAccount data: %w", err))
		}

		// Pack factory address and deployment data for `CreateSender` in EntryPoint
		// see https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/SenderCreator.sol#L15
		initCode := make([]byte, len(factory)+len(callData))
		copy(initCode, factory.Bytes())
		copy(initCode[len(factory):], callData)

		slog.Debug("built initCode", "initCode", hexutil.Encode(initCode))

		return initCode, nil
	}
}

func getGasPrice(providerRPC *ethclient.Client, gasConfig GasConfig) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		slog.Info("getting gas price")

		// Get the latest block to read its base fee
		block, err := providerRPC.BlockByNumber(ctx, nil)
		if err != nil {
			return err
		}
		blockBaseFee := block.BaseFee()

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

		// Increase maxPriorityFeePerGas to give user more
		// flexibility in setting the gas price.
		maxPriorityFeePerGas.Mul(
			maxPriorityFeePerGas,
			gasConfig.MaxPriorityFeePerGasMultiplier.BigInt())

		// Calculate maxFeePerGas
		maxFeePerGas := new(big.Int).Mul(blockBaseFee, gasConfig.MaxFeePerGasMultiplier.BigInt())
		maxFeePerGas.Add(maxFeePerGas, maxPriorityFeePerGas)

		op.MaxFeePerGas = decimal.NewFromBigInt(maxFeePerGas, 0)
		op.MaxPriorityFeePerGas = decimal.NewFromBigInt(maxPriorityFeePerGas, 0)

		maxBlockGas := decimal.NewFromInt(maxBlockGas)
		op.CallGasLimit = maxBlockGas
		op.VerificationGasLimit = maxBlockGas
		op.PreVerificationGas = maxBlockGas

		return nil
	}
}

func getBiconomyPaymasterAndData(
	bundlerRPC *rpc.Client,
	paymasterCtx map[string]any,
	entryPoint common.Address,
) middleware {
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

		// Biconomy-standardized gas estimation with paymaster
		// see https://docs.biconomy.io/Paymaster/api/sponsor-useroperation#sponsorship-paymaster
		var est gasEstimate
		if err := bundlerRPC.CallContext(
			ctx,
			&est,
			"pm_sponsorUserOperation",
			opModified,
			entryPoint,
			paymasterCtx,
		); err != nil {
			return fmt.Errorf("failed to call pm_sponsorUserOperation: %w", err)
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

func getPimlicoERC20PaymasterData(
	bundlerRPC *rpc.Client,
	entryPoint common.Address,
	paymaster common.Address,
	verificationGasOverhead decimal.Decimal,
) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		estimate := estimateUserOperationGas(bundlerRPC, entryPoint)
		if err := estimate(ctx, op); err != nil {
			return err
		}

		op.VerificationGasLimit = op.VerificationGasLimit.Add(verificationGasOverhead)
		op.PaymasterAndData = paymaster.Bytes()

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

func (est gasEstimate) convert() (
	preVerificationGas *big.Int,
	verificationGasLimit *big.Int,
	callGasLimit *big.Int,
	err error,
) {
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
	// The gas estimate is increased by 10_000 to account for
	// a rare yet hard to debug bundler gas estimation inaccuracy.
	// It is NOT a random value, see how it works in the original Alto code:
	// https://github.com/pimlicolabs/alto/blob/a0a9a4906af809d97611c7f0e0f032e50c4c45cb/src/entrypoint-0.6/rpc/gasEstimation.ts#L277-L279
	gasEstimationOverhead := decimal.NewFromInt(10_000)

	return func(ctx context.Context, op *UserOperation) error {
		slog.Info("estimating gas")

		// ERC4337-standardized gas estimation
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

		op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0).Add(gasEstimationOverhead)
		op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
		op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)

		return nil
	}
}

func sign(entryPoint common.Address, chainID *big.Int) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		signer, ok := ctx.Value(ctxKeySigner).(Signer)
		if !ok {
			return fmt.Errorf("signer not found in context")
		}

		signature, err := signer(*op, entryPoint, chainID)
		if err != nil {
			return fmt.Errorf("failed to sign user operation: %w", err)
		}

		op.Signature = signature

		b, err := op.MarshalJSON()
		if err != nil {
			return fmt.Errorf("failed to marshal user operation: %w", err)
		}
		slog.Debug("userop signed",
			"hash", op.UserOpHash(entryPoint, chainID).String(),
			"json", string(b))

		return nil
	}
}
