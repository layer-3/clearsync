package main

import (
	"crypto/ecdsa"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/userop"
)

type exampleECDSASigner struct {
	privateKey *ecdsa.PrivateKey
}

func newExampleECDSASigner(privateKey *ecdsa.PrivateKey) exampleECDSASigner {
	return exampleECDSASigner{privateKey: privateKey}
}

func (s exampleECDSASigner) Sign(msg []byte) ([]byte, error) {
	return crypto.Sign(msg, s.privateKey)
}

var (
	exampleConfig = userop.ClientConfig{
		ProviderURL: *must(url.Parse("https://NETWORK.infura.io/v3/YOUR_INFURA_API_KEY")),
		BundlerURL:  *must(url.Parse("https://api.pimlico.io/v1/NETWORK/rpc?apikey=YOUR_PIMLICO_API_KEY")),
		EntryPoint:  common.HexToAddress("ENTRY_POINT_ADDRESS"),
		SmartWallet: userop.SmartWalletConfig{
			// Example of a Kernel Smart Wallet config with Kernel v2.2.
			Type: &userop.SmartWalletKernel,
			// Zerodev Kernel factory address:
			Factory: common.HexToAddress("0x5de4839a76cf55d0c90e2061ef4386d962E15ae3"),
			// Zerodev Kernel implementation (logic) address:
			Logic:          common.HexToAddress("0x0DA6a956B9488eD4dd761E59f52FDc6c8068E6B5"),
			ECDSAValidator: common.HexToAddress("0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"),
		},
		Paymaster: userop.PaymasterConfig{
			// Example of a Pimlico ERC20 Paymaster config.
			Type:    &userop.PaymasterPimlicoERC20,
			URL:     url.URL{},
			Address: common.HexToAddress("0xa683b47e447De6c8A007d9e294e87B6Db333Eb18"),
			PimlicoERC20: userop.PimlicoERC20Config{
				VerificationGasOverhead: decimal.RequireFromString("10000"),
			},
		},
		Gas: userop.GasConfig{
			// These are values by default.
			MaxPriorityFeePerGasMultiplier: decimal.RequireFromString("1.13"),
			MaxFeePerGasMultiplier:         decimal.RequireFromString("2"),
		},
	}

	walletDeploymentOpts = &userop.WalletDeploymentOpts{}

	exampleSigner = userop.SignerForKernel(
		newExampleECDSASigner(
			must(crypto.HexToECDSA(
				"YOUR_PRIVATE_KEY",
			)),
		),
	)
)

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}
