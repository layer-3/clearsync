package smart_wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func GetFactoryCallData(smartWalletConfig Config, ownerAddress common.Address, index decimal.Decimal) ([]byte, error) {
	var initCode []byte
	var err error

	if smartWalletConfig.Type == nil {
		return nil, ErrNoSmartWalletType
	}

	switch typ := *smartWalletConfig.Type; typ {
	case SimpleAccountType:
		return nil, fmt.Errorf("%w: %s", ErrSmartWalletNotSupported, typ)
	case BiconomyType: // not tested
		initCode, err = GetBiconomyFactoryCallData(ownerAddress, index, smartWalletConfig.ECDSAValidator)
	case KernelType:
		initCode, err = GetKernelFactoryCallData(ownerAddress, index, smartWalletConfig.Logic, smartWalletConfig.ECDSAValidator)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", typ)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get init code: %w", err)
	}

	return initCode, nil
}

// getKernelInitData returns the calldata needed call the factory
// to deploy a Zerodev Kernel smart account.
func GetKernelFactoryCallData(owner common.Address, index decimal.Decimal, accountLogic, ecdsaValidator common.Address) ([]byte, error) {
	// Initialize Kernel Smart Account with default validation module and its calldata
	// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/abstract/KernelStorage.sol#L35
	initData, err := kernelInitABI.Pack("initialize", ecdsaValidator, owner.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to pack init data: %w", err)
	}

	// Deploy Kernel Smart Account by calling `factory.createAccount`
	// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/factory/KernelFactory.sol#L25
	callData, err := kernelDeployWalletABI.Pack("createAccount", accountLogic, initData, index.BigInt())
	if err != nil {
		return nil, fmt.Errorf("failed to pack createAccount data: %w", err)
	}

	return callData, nil
}

// getKernelInitData returns the calldata needed call the factory
// to deploy a Biconomy smart account.
func GetBiconomyFactoryCallData(owner common.Address, index decimal.Decimal, ecdsaValidator common.Address) ([]byte, error) {
	// Initialize SCW validation module with owner address
	// see https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/modules/EcdsaOwnershipRegistryModule.sol#L43
	ecdsaOwnershipInitData, err := biconomyInitABI.Pack("initForSmartAccount", owner.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to pack init data: %w", err)
	}

	// Deploy Biconomy SCW by calling `factory.createAccount`
	// see https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/factory/SmartAccountFactory.sol#L112
	callData, err := biconomyDeployWalletABI.Pack("createAccount", ecdsaValidator, ecdsaOwnershipInitData, index.BigInt())
	if err != nil {
		return nil, fmt.Errorf("failed to pack createAccount data: %w", err)
	}

	return callData, nil
}
