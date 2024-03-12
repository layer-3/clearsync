package userop

import (
	"context"
	"math/big"
	"math/rand"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func defaultProviderURL() string {
	if value := os.Getenv("PROVIDER_URL"); value != "" {
		return value
	}

	// NOTE: this is a public endpoint, which can suffer from rate limiting and deprecation.
	return "https://sepolia.gateway.tenderly.co"
}

func randomAddress() common.Address {
	return common.BigToAddress(big.NewInt(rand.Int63()))
}

func mockConfig() ClientConfig {
	config := ClientConfig{
		ProviderURL: "http://127.0.0.1:42424",
		BundlerURL:  "http://127.0.0.1:42424",
		EntryPoint:  common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"),
		SmartWallet: SmartWalletConfig{
			Type:           &SmartWalletKernel,
			Factory:        common.HexToAddress("0x5de4839a76cf55d0c90e2061ef4386d962E15ae3"),
			Logic:          common.HexToAddress("0x0DA6a956B9488eD4dd761E59f52FDc6c8068E6B5"),
			ECDSAValidator: common.HexToAddress("0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"),
		},
		Paymaster: PaymasterConfig{
			Type: &PaymasterDisabled,
		},
	}

	config.Gas.Init()
	return config
}

func bundlerMock(t *testing.T, providerURL string) Client {
	config := ClientConfig{
		ProviderURL: providerURL,
		BundlerURL:  "http://127.0.0.1:42424",
		EntryPoint:  common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"),
		SmartWallet: SmartWalletConfig{
			Type:           &SmartWalletKernel,
			Factory:        common.HexToAddress("0x5de4839a76cf55d0c90e2061ef4386d962E15ae3"),
			Logic:          common.HexToAddress("0x0DA6a956B9488eD4dd761E59f52FDc6c8068E6B5"),
			ECDSAValidator: common.HexToAddress("0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"),
		},
		Paymaster: PaymasterConfig{
			Type: &PaymasterDisabled,
		},
	}

	config.Gas.Init()

	client, err := NewClient(config)
	require.NoError(t, err)

	return client
}

func randomOwnerWithoutAccount(client Client, t *testing.T) common.Address {
	ctx := context.Background()

	var randomOwner common.Address
	for {
		randomOwner = randomAddress()
		isDeployed, err := client.IsAccountDeployed(ctx, randomOwner, decimal.Zero)
		require.NoError(t, err)

		if !isDeployed {
			break
		}
	}
	return randomOwner
}
