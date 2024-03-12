package public_blockchain

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/userop"
)

var (
	exampleConfig = userop.ClientConfig{
		ProviderURL: "https://NETWORK.infura.io/v3/YOUR_INFURA_API_KEY",
		BundlerURL:  "https://api.pimlico.io/v1/NETWORK/rpc?apikey=YOUR_PIMLICO_API_KEY",
		PollPeriod:  100 * time.Millisecond,
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

	exampleSigner = userop.SignerForKernel(
		signer_pkg.NewLocalSigner(
			must(crypto.HexToECDSA(
				"deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
			)),
		),
	)
)

var exampleWalletDeploymentOpts *userop.WalletDeploymentOpts

// You can set either of gas limits to override the bundler's estimation.
// Or you can set all of them to disable the bundler's estimation.
var exampleGasLimitOverrides *userop.GasLimitOverrides

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}
