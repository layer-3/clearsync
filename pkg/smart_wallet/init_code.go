package smart_wallet

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

func IsAccountDeployed(provider *ethclient.Client, swAddress common.Address) (bool, error) {
	byteCode, err := provider.CodeAt(context.Background(), swAddress, nil)
	if err != nil {
		return false, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	}

	// assume that the smart account is deployed if it has non-zero byte code
	return len(byteCode) != 0, nil
}

func GetInitCode(provider *ethclient.Client, smartWalletConfig Config, smartWalletAddress, ownerAddress common.Address, index decimal.Decimal) ([]byte, error) {
	// check if smart account is already deployed
	isDeployed, err := IsAccountDeployed(provider, smartWalletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	}
	var initCode []byte

	// if sender == zeroAddress OR smart account is not deployed
	// then we need to calculate the init code
	if smartWalletAddress == (common.Address{}) || !isDeployed {

		switch typ := *smartWalletConfig.Type; typ {
		case SimpleAccountType:
			return nil, fmt.Errorf("%w: %s", ErrSmartWalletNotSupported, typ)
		case BiconomyType: // not tested
			initCode, err = GetBiconomyInitCode(ownerAddress, index, smartWalletConfig.Factory, smartWalletConfig.ECDSAValidator)
		case KernelType:
			initCode, err = GetKernelInitCode(ownerAddress, index, smartWalletConfig.Factory, smartWalletConfig.Logic, smartWalletConfig.ECDSAValidator)
		default:
			return nil, fmt.Errorf("unknown smart wallet type: %s", typ)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to get init code: %w", err)
		}
	}

	return initCode, nil
}

// getKernelInitCode returns a middleware that sets the init code
// for a Zerodev Kernel smart account. The init code deploys
// a smart account if it is not already deployed.
func GetKernelInitCode(owner common.Address, index decimal.Decimal, factory, accountLogic, ecdsaValidator common.Address) ([]byte, error) {
	// Initialize Kernel Smart Account with default validation module and its calldata
	// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/abstract/KernelStorage.sol#L35
	initData, err := kernelInitABI.Pack("initialize", ecdsaValidator, owner.Bytes())
	if err != nil {
		panic(fmt.Errorf("failed to pack init data: %w", err))
	}

	// Deploy Kernel Smart Account by calling `factory.createAccount`
	// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/factory/KernelFactory.sol#L25
	callData, err := kernelDeployWalletABI.Pack("createAccount", accountLogic, initData, index.BigInt())
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

// getBiconomyInitCode returns a middleware that sets the init code for a Biconomy smart account.
// The init code deploys a smart account if it is not already deployed.
// NOTE: this was NOT tested. Use at your own risk or wait for the package to be updated.
func GetBiconomyInitCode(owner common.Address, index decimal.Decimal, factory, ecdsaValidator common.Address) ([]byte, error) {
	// Initialize SCW validation module with owner address
	// see https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/modules/EcdsaOwnershipRegistryModule.sol#L43
	ecdsaOwnershipInitData, err := biconomyInitABI.Pack("initForSmartAccount", owner.Bytes())
	if err != nil {
		panic(fmt.Errorf("failed to pack init data: %w", err))
	}

	// Deploy Biconomy SCW by calling `factory.createAccount`
	// see https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/factory/SmartAccountFactory.sol#L112
	callData, err := biconomyDeployWalletABI.Pack("createAccount", ecdsaValidator, ecdsaOwnershipInitData, index.BigInt())
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
