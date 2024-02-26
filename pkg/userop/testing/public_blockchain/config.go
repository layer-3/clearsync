package public_blockchain

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/userop"
)

var (
	config = userop.ClientConfig{
		ProviderURL: *must(url.Parse("https://polygon-mainnet.infura.io/v3/16a479138e474a8cb10bd9a26a02fbae")),
		BundlerURL:  *must(url.Parse("https://api.pimlico.io/v1/polygon/rpc?apikey=d1b58599-575e-4eee-afa9-eb057a6be7d6")),
		EntryPoint:  common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"),
		Gas: userop.GasConfig{
			MaxPriorityFeePerGasMultiplier: decimal.RequireFromString("1.13"),
			MaxFeePerGasMultiplier:         decimal.RequireFromString("2"),
		},
		SmartWallet: userop.SmartWalletConfig{
			Type: &userop.SmartWalletKernel,
			// Zerodev Kernel factory address:
			Factory: common.HexToAddress("0x5de4839a76cf55d0c90e2061ef4386d962E15ae3"),
			// Zerodev Kernel implementation (logic) address:
			Logic:          common.HexToAddress("0x0DA6a956B9488eD4dd761E59f52FDc6c8068E6B5"),
			ECDSAValidator: common.HexToAddress("0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"),
		},
		Paymaster: userop.PaymasterConfig{
			Type: &userop.PaymasterPimlicoERC20,
			// URL:     "https://api.pimlico.io/v2/polygon/rpc?apikey=d1b58599-575e-4eee-afa9-eb057a6be7d6",
			Address: common.HexToAddress("0x00000000003011eEF3f79892ba3d521E5Ba5C5c0"),
			PimlicoERC20: userop.PimlicoERC20Config{
				VerificationGasOverhead: decimal.NewFromInt(10_000),
			},
		},
	}

	signer = userop.SignerForKernel(must(crypto.HexToECDSA(
		"26b556ff5c77f622504ed5e474919db6e4533fdc62b2f5965a26a6b22eb86f3f")))
)
