package userop

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestClientNewUserOp(t *testing.T) {
	// Skipping the test since it uses public RPC endpoint
	// which often fails with 429 Too Many Requests.
	t.Skip()

	t.Parallel()

	t.Run("Error when signer is not specified", func(t *testing.T) {
		t.Parallel()

		client := bundlerMock(t, defaultProviderURL())
		ctx := context.Background()

		_, err := client.NewUserOp(ctx, common.Address{}, nil, nil, nil, nil)
		require.EqualError(t, err, ErrNoSigner.Error())
	})

	t.Run("Error when no calls specified", func(t *testing.T) {
		t.Parallel()

		client := bundlerMock(t, defaultProviderURL())
		ctx := context.Background()

		// create random owner so that no wallet is deployed
		randomOwner := randomOwnerWithoutAccount(client, t)
		smartWallet, err := client.GetAccountAddress(ctx, randomOwner, decimal.Zero)
		require.NoError(t, err)

		// create userop without wallet deployment opts
		_, err = client.NewUserOp(ctx, smartWallet, SignerForKernel(nil), nil, nil, nil)

		// assert error
		require.EqualError(t, err, ErrNoCalls.Error())
	})

	t.Run("Error when no wallet deployed and no wallet deployment opts", func(t *testing.T) {
		t.Parallel()

		client := bundlerMock(t, defaultProviderURL())
		ctx := context.Background()

		// create random owner so that no wallet is deployed
		randomOwner := randomOwnerWithoutAccount(client, t)
		smartWallet, err := client.GetAccountAddress(ctx, randomOwner, decimal.Zero)
		require.NoError(t, err)

		calls := smart_wallet.Calls{
			{
				To: common.Address{},
			},
		}

		// create userop without wallet deployment opts
		_, err = client.NewUserOp(ctx, smartWallet, SignerForKernel(nil), calls, nil, nil)

		// assert error
		require.EqualError(t, err, ErrNoWalletDeploymentOpts.Error())
	})

	t.Run("Error when no wallet deployed and owner is zero in wallet deployment opts", func(t *testing.T) {
		t.Parallel()

		client := bundlerMock(t, defaultProviderURL())
		ctx := context.Background()

		calls := smart_wallet.Calls{
			{
				To: common.Address{},
			},
		}

		wdo := &WalletDeploymentOpts{}
		// create userop with wallet deployment opts with zero owner
		_, err := client.NewUserOp(ctx, common.Address{}, SignerForKernel(nil), calls, wdo, nil)

		// assert error
		require.EqualError(t, err, ErrNoWalletOwnerInWDO.Error())
	})

	t.Run("Error when no calls specified", func(t *testing.T) {
		t.Parallel()

		client := bundlerMock(t, defaultProviderURL())
		ctx := context.Background()

		// create random owner so that no wallet is deployed
		randomOwner := randomOwnerWithoutAccount(client, t)
		wdo := &WalletDeploymentOpts{
			Owner: randomOwner,
		}

		// create userop with wallet deployment opts with zero owner
		_, err := client.NewUserOp(ctx, common.Address{}, SignerForKernel(nil), nil, wdo, nil)

		// assert error
		require.EqualError(t, err, ErrNoCalls.Error())
	})
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("Error when entrypoint address is empty", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.EntryPoint = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidEntryPointAddress.Error())
	})

	t.Run("Error when factory address is empty", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.SmartWallet.Factory = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidFactoryAddress.Error())
	})

	t.Run("Error when logic address is empty", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.SmartWallet.Logic = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidLogicAddress.Error())
	})

	t.Run("Error when ECDSA validator address is empty", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.SmartWallet.ECDSAValidator = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidECDSAValidatorAddress.Error())
	})

	t.Run("Error when paymaster address is empty", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.Paymaster.Type = &PaymasterPimlicoERC20
		conf.Paymaster.Address = common.Address{}

		_, err := NewClient(conf)

		require.EqualError(t, errors.Unwrap(err), ErrInvalidPaymasterAddress.Error())
	})

	t.Run("Paymaster address can be empty if no paymaster config", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.ProviderURL = defaultProviderURL()
		conf.Paymaster = PaymasterConfig{}

		_, err := NewClient(conf)

		require.NoError(t, err)
	})

	t.Run("Paymaster address can be empty if paymaster is disabled", func(t *testing.T) {
		t.Parallel()

		conf := mockConfig()
		conf.ProviderURL = defaultProviderURL()
		conf.Paymaster.Type = &PaymasterDisabled

		_, err := NewClient(conf)

		require.NoError(t, err)
	})

	t.Run("Logger level is info by default", func(t *testing.T) {
		// t.Parallel() // can't be run in parallel due to global logger

		conf := mockConfig()
		conf.ProviderURL = defaultProviderURL()
		conf.LoggerLevel = ""

		_, err := NewClient(conf)
		require.NoError(t, err)

		ctx := context.Background()

		require.False(t, slog.Default().Enabled(ctx, slog.LevelDebug))
		require.True(t, slog.Default().Enabled(ctx, slog.LevelInfo))
	})

	t.Run("Logger level is parsed correctly", func(t *testing.T) {
		// t.Parallel() // can't be run in parallel due to global logger

		conf := mockConfig()
		conf.ProviderURL = defaultProviderURL()
		conf.LoggerLevel = "debug"

		_, err := NewClient(conf)
		require.NoError(t, err)

		ctx := context.Background()

		require.True(t, slog.Default().Enabled(ctx, slog.LevelDebug))
	})

	t.Run("Error on incorrect logger level", func(t *testing.T) {
		// t.Parallel() // can't be run in parallel due to global logger

		conf := mockConfig()
		conf.ProviderURL = defaultProviderURL()
		conf.LoggerLevel = "deafbeef"

		_, err := NewClient(conf)
		require.ErrorContains(t, err, "failed to set logger level")
	})
}
