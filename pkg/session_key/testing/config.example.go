package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/session_key"
	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
)

var (
	exampleUserOpConfig = userop.ClientConfig{
		ProviderURL: "https://NETWORK.infura.io/v3/YOUR_INFURA_API_KEY",
		BundlerURL:  "https://api.pimlico.io/v1/NETWORK/rpc?apikey=YOUR_PIMLICO_API_KEY",
		EntryPoint:  common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"),
		SmartWallet: smart_wallet.Config{
			// Example of a Kernel Smart Wallet config with Kernel v2.2.
			Type: &smart_wallet.KernelType,
			// Zerodev Kernel factory address:
			Factory: common.HexToAddress("0x5de4839a76cf55d0c90e2061ef4386d962E15ae3"),
			// Zerodev Kernel implementation (logic) address:
			Logic:          common.HexToAddress("0x0DA6a956B9488eD4dd761E59f52FDc6c8068E6B5"),
			ECDSAValidator: common.HexToAddress("0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"),
		},
		Paymaster: userop.PaymasterConfig{
			// Example of a Pimlico ERC20 Paymaster config.
			Type:    &userop.PaymasterPimlicoERC20,
			URL:     "",
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

	exampleWalletDeploymentOpts = &userop.WalletDeploymentOpts{}

	exampleSigner = signer_pkg.NewLocalSigner(
		must(crypto.HexToECDSA(
			"deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
		)),
	)

	exampleUserOpSigner = userop.SignerForKernel(exampleSigner)

	exampleSessionKeySigner = signer_pkg.NewLocalSigner(
		must(crypto.HexToECDSA(
			"deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
		)),
	)

	exampleSessionKeyConfig = session_key.Config{
		ProviderURL:                "https://NETWORK.infura.io/v3/YOUR_INFURA_API_KEY",
		ExecuteInBatch:             true,
		SessionKeyValidAfter:       0,
		SessionKeyValidUntil:       0,
		SessionKeyValidatorAddress: common.HexToAddress("0x5C06CE2b673fD5E6e56076e40DD46aB67f5a72A5"),
		ExecutorAddress:            common.HexToAddress("0x0000000000000000000000000000000000000000"),
		PaymasterAddress:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Permissions:                []session_key.Permission{
			// list of your permissions to be applied to every session key
		},
	}
)

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}
