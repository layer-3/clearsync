package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/layer-3/clearsync/pkg/abi/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/abi/kernel_ecdsa_validator_v2_2"
	"github.com/layer-3/clearsync/pkg/abi/kernel_factory_v2_2"
	"github.com/layer-3/clearsync/pkg/abi/kernel_v2_2"
	"github.com/layer-3/clearsync/pkg/userop"
)

func TestSimulatedRPC(t *testing.T) {
	ctx := context.Background()

	// 1. Start a local Ethereum node
	rpcURL, _ /*accountBalance*/ := startEthNode(ctx, t)
	ethClient, err := ethclient.Dial(rpcURL.String())
	require.NoError(t, err)

	// 2. Deploy the required contracts
	addresses := deployContracts(ctx, t, ethClient)

	// 3. Start the bundler
	signer, err := ecrypto.GenerateKey()
	require.NoError(t, err, "failed to generate SIGNER private key for bundler")
	utility, err := ecrypto.GenerateKey()
	require.NoError(t, err, "failed to generate UTILITY private key for bundler")
	bundlerURL := startBundler(ctx, t, rpcURL, addresses.entryPoint, signer, utility)

	// 4. Start the user operation client
	config := buildConfig(t, rpcURL, *bundlerURL, addresses)
	_ /*client*/, err = userop.NewClient(config)
	require.NoError(t, err)

	// 5. Run transactions
	// TODO: set up tests

	<-time.After(60 * time.Second)
}

func startEthNode(ctx context.Context, t *testing.T) (url.URL, *AccountBalance) {
	rpcURL, err := url.Parse("http://localhost:8545")
	require.NoError(t, err, "failed to parse local RPC URL")

	gethContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "ethereum/client-go:stable",
			// 8545 TCP, used by the HTTP based JSON RPC API
			// 8546 TCP, used by the WebSocket based JSON RPC API
			// 8547 TCP, used by the GraphQL API
			// 30303 TCP and UDP, used by the P2P protocol running the network
			ExposedPorts: []string{"8545/tcp", "8546/tcp", "8547/tcp", "30303/tcp", "30303/udp"},
			Entrypoint:   []string{"geth", "--dev", "--http", "--http.api=eth,web3,net"},
			WaitingFor:   wait.ForLog("server started"),
		},
		Started: true,
	})
	require.NoError(t, err, "failed to start Go-Ethereum container")

	t.Cleanup(func() {
		err := gethContainer.Terminate(ctx)
		require.NoError(t, err, "failed to terminate Go-Ethereum container")
	})

	return *rpcURL, NewAccountBalance(gethContainer, *rpcURL)
}

func startBundler(
	ctx context.Context,
	t *testing.T,
	rpcURL url.URL,
	entryPoint common.Address,
	signer *ecdsa.PrivateKey,
	utility *ecdsa.PrivateKey,
) *url.URL {
	bundlerURL, err := url.Parse("http://localhost:3000")
	require.NoError(t, err, "failed to parse local bundler URL")

	altoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "ghcr.io/pimlicolabs/alto:v1.0.1",
			Entrypoint: []string{"pnpm", "start",
				"--networkName", "mainnet", // check Go-Ethereum container logs to find out configured network
				"--entryPoint", entryPoint.Hex(), // the contract should already be deployed on Go-Ethereum node
				"--signerPrivateKeys", privateKeyToString(signer),
				"--utilityPrivateKey", privateKeyToString(utility),
				"--minBalance", "0",
				"--rpcUrl", rpcURL.String(),
			},
			WaitingFor: wait.ForLog("server started"),
		},
		Started: true,
	})
	require.NoError(t, err, "failed to start Alto bundler container")

	t.Cleanup(func() {
		if err := altoContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate Alto container: %s", err)
		}
	})

	return bundlerURL
}

type contracts struct {
	entryPoint common.Address
	validator  common.Address
	factory    common.Address
	logic      common.Address
	paymaster  common.Address
}

func deployContracts(ctx context.Context, t *testing.T, ethClient *ethclient.Client) contracts {
	chainID, err := ethClient.ChainID(ctx)
	require.NoError(t, err)

	entryPoint, _, _, err := entry_point_v0_6_0.DeployEntryPoint(nil, ethClient)
	require.NoError(t, err)

	validator, _, _, err := kernel_ecdsa_validator_v2_2.DeployKernelECDSAValidator(nil, ethClient)
	require.NoError(t, err)

	factoryOwner, err := NewAccount(chainID)
	require.NoError(t, err)

	factory, _, _, err := kernel_factory_v2_2.DeployKernelFactory(nil, ethClient, factoryOwner.Address, entryPoint)
	require.NoError(t, err)

	logic, _, _, err := kernel_v2_2.DeployKernel(nil, ethClient, entryPoint)
	require.NoError(t, err)

	var paymaster common.Address // TODO: deploy Paymaster contract

	return contracts{
		entryPoint: entryPoint,
		validator:  validator,
		factory:    factory,
		logic:      logic,
		paymaster:  paymaster,
	}
}

func privateKeyToString(privateKey *ecdsa.PrivateKey) string {
	// Serialize the private key to PEM format
	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Failed to marshal ECDSA private key: %v", err)
	}
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: x509Encoded})

	// Convert the PEM to a string (PEM is already a string format)
	return string(pemEncoded)
}

func buildConfig(t *testing.T, rpcURL, bundlerURL url.URL, addresses contracts) userop.ClientConfig {
	config, err := userop.NewClientConfigFromEnv()
	require.NoError(t, err)

	config.ProviderURL = rpcURL
	config.BundlerURL = bundlerURL
	config.EntryPoint = addresses.entryPoint
	config.SmartWallet = userop.SmartWalletConfig{
		Type:           &userop.SmartWalletKernel,
		ECDSAValidator: addresses.validator,
		Logic:          addresses.factory,
		Factory:        addresses.logic,
	}
	config.Paymaster.Type = &userop.PaymasterPimlicoERC20
	config.Paymaster.Address = addresses.paymaster

	return config
}
