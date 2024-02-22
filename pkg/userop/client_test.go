package userop

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestClientNewUserOp(t *testing.T) {
	t.Parallel()

	t.Run("Error when no wallet deployed and no wallet deployment opts", func(t *testing.T) {
		client := bundlerMockedClient(t, providerURL)
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
		client := bundlerMockedClient(t, providerURL)
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
