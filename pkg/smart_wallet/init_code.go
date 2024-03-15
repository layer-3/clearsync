package smart_wallet

import (
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

func GetInitCode(smartWalletConfig Config, ownerAddress common.Address, index decimal.Decimal) ([]byte, error) {
	var initCode []byte
	var err error

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

	return initCode, nil
}

func GetInitCodeFromFactoryCallData(smartWalletConfig Config, factoryCallData []byte) ([]byte, error) {
	var initCode []byte
	var err error

	switch typ := *smartWalletConfig.Type; typ {
	case SimpleAccountType:
		return nil, fmt.Errorf("%w: %s", ErrSmartWalletNotSupported, typ)
	case BiconomyType: // not tested
		initCode, err = GetBiconomyInitCodeFromFactoryCallData(smartWalletConfig.Factory, factoryCallData)
	case KernelType:
		initCode, err = GetKernelInitCodeFromFactoryCallData(smartWalletConfig.Factory, factoryCallData)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", typ)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get init code: %w", err)
	}

	return initCode, nil
}

// GetKernelInitCode the init code for a Zerodev Kernel smart account.
// The init code deploys a smart account if it is not already deployed.
func GetKernelInitCode(owner common.Address, index decimal.Decimal, factory, accountLogic, ecdsaValidator common.Address) ([]byte, error) {
	callData, err := GetKernelFactoryCallData(owner, index, accountLogic, ecdsaValidator)
	if err != nil {
		return nil, fmt.Errorf("failed to get init data: %w", err)
	}

	return GetKernelInitCodeFromFactoryCallData(factory, callData)
}

// GetKernelInitCodeFromFactoryCallData returns the init code for a Zerodev Kernel smart account,
// given the Kernel Factory calldata. The init code deploys a smart account if it is not already deployed.
func GetKernelInitCodeFromFactoryCallData(factory common.Address, factoryCallData []byte) ([]byte, error) {
	// Pack factory address and deployment data for `CreateSender` in EntryPoint
	// see https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/SenderCreator.sol#L15
	initCode := make([]byte, len(factory)+len(factoryCallData))
	copy(initCode, factory.Bytes())
	copy(initCode[len(factory):], factoryCallData)

	slog.Debug("built initCode", "initCode", hexutil.Encode(initCode))

	return initCode, nil
}

// GetBiconomyInitCode returns the init code for a Biconomy smart account.
// The init code deploys a smart account if it is not already deployed.
// NOTE: this was NOT tested. Use at your own risk or wait for the package to be updated.
func GetBiconomyInitCode(owner common.Address, index decimal.Decimal, factory, ecdsaValidator common.Address) ([]byte, error) {
	callData, err := GetBiconomyFactoryCallData(owner, index, ecdsaValidator)
	if err != nil {
		return nil, fmt.Errorf("failed to get init data: %w", err)
	}

	return GetBiconomyInitCodeFromFactoryCallData(factory, callData)
}

// GetBiconomyInitCodeFromFactoryCallData returns the init code for a Biconomy smart account,
// given the Biconomy Factory calldata. The init code deploys a smart account if it is not already deployed.
func GetBiconomyInitCodeFromFactoryCallData(factory common.Address, factoryCallData []byte) ([]byte, error) {
	// Pack factory address and deployment data for `CreateSender` in EntryPoint
	// see https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/SenderCreator.sol#L15
	initCode := make([]byte, len(factory)+len(factoryCallData))
	copy(initCode, factory.Bytes())
	copy(initCode[len(factory):], factoryCallData)

	slog.Debug("built initCode", "initCode", hexutil.Encode(initCode))

	return initCode, nil
}
