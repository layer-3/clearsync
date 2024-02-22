package userop

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestClientNewUserOp(t *testing.T) {
	t.Parallel()

	t.Run("Error when no wallet deployed and no wallet deployment opts", func(t *testing.T) {
		client := bundlerMockedClient(t, defaultProviderURL)
		ctx := context.Background()

		// create random owner so that no wallet is deployed
		var randomOwner common.Address
		for {
			randomOwner = randomAddress()
			isDeployed, err := client.IsAccountDeployed(ctx, randomOwner, decimal.Zero)
			require.NoError(t, err)

			if !isDeployed {
				break
			}
		}

		smartWallet, err := client.GetAccountAddress(ctx, randomOwner, decimal.Zero)
		require.NoError(t, err)

		// create userop without wallet deployment opts
		_, err = client.NewUserOp(ctx, smartWallet, nil, nil, nil)

		// assert error
		require.EqualError(t, err, ErrNoWalletDeploymentOpts.Error())
	})

	t.Run("Error when no wallet deployed and owner is zero in wallet deployment opts", func(t *testing.T) {
		client := bundlerMockedClient(t, defaultProviderURL)
		ctx := context.Background()

		// create random owner so that no wallet is deployed
		var randomOwner common.Address
		for {
			randomOwner = randomAddress()
			isDeployed, err := client.IsAccountDeployed(ctx, randomOwner, decimal.Zero)
			require.NoError(t, err)

			if !isDeployed {
				break
			}
		}

		wdo := &WalletDeploymentOpts{}

		// create userop with wallet deployment opts with zero owner
		_, err := client.NewUserOp(ctx, common.Address{}, nil, nil, wdo)

		// assert error
		require.EqualError(t, err, ErrNoWalletOwner.Error())
	})
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("Error when entrypoint address is empty", func(t *testing.T) {
		conf := mockedConfig()
		conf.EntryPoint = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidEntryPointAddress.Error())
	})

	t.Run("Error when factory address is empty", func(t *testing.T) {
		conf := mockedConfig()
		conf.SmartWallet.Factory = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidFactoryAddress.Error())
	})

	t.Run("Error when logic address is empty", func(t *testing.T) {
		conf := mockedConfig()
		conf.SmartWallet.Logic = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidLogicAddress.Error())
	})

	t.Run("Error when ECDSA validator address is empty", func(t *testing.T) {
		conf := mockedConfig()
		conf.SmartWallet.ECDSAValidator = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidECDSAValidatorAddress.Error())
	})

	t.Run("Error when paymaster address is empty", func(t *testing.T) {
		conf := mockedConfig()
		conf.Paymaster.Type = &PaymasterPimlicoERC20
		conf.Paymaster.Address = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidPaymasterAddress.Error())
	})

	t.Run("Paymaster address can be empty if no paymaster config", func(t *testing.T) {
		conf := mockedConfig()
		conf.ProviderURL = *must(url.Parse(defaultProviderURL))
		conf.Paymaster = PaymasterConfig{}

		_, err := NewClient(conf)

		require.NoError(t, err)
	})

	t.Run("Paymaster address can be empty if paymaster is disabled", func(t *testing.T) {
		conf := mockedConfig()
		conf.ProviderURL = *must(url.Parse(defaultProviderURL))
		conf.Paymaster.Type = &PaymasterDisabled

		_, err := NewClient(conf)

		require.NoError(t, err)
	})
}
