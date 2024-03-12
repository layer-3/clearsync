package local_blockchain

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/layer-3/clearsync/pkg/abi/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/abi/kernel_ecdsa_validator_v2_2"
	"github.com/layer-3/clearsync/pkg/abi/kernel_factory_v2_2"
	"github.com/layer-3/clearsync/pkg/abi/kernel_v2_2"
	"github.com/layer-3/clearsync/pkg/artifacts/session_key_validator"
	"github.com/layer-3/clearsync/pkg/userop"
)

type EthNode struct {
	Container    testcontainers.Container
	Client       *ethclient.Client
	LocalURL     url.URL
	ContainerURL url.URL
}

var (
	ethActiveNode  *EthNode
	ethActiveUsers int64
	ethMutex       sync.Mutex
)

func ethNodeClenaupFactory(ctx context.Context, t *testing.T) func() {
	return func() {
		ethMutex.Lock()
		defer ethMutex.Unlock()

		ethActiveUsers--
		if ethActiveUsers <= 0 {
			err := ethActiveNode.Container.Terminate(ctx)
			require.NoError(t, err, "failed to terminate Go-Ethereum container")

			ethActiveNode = nil
		}
	}
}

func NewEthNode(ctx context.Context, t *testing.T) *EthNode {
	ethMutex.Lock()
	defer ethMutex.Unlock()
	slog.Info("starting Go-Ethereum node...")

	if ethActiveNode != nil {
		ethActiveUsers++
		t.Cleanup(ethNodeClenaupFactory(ctx, t))

		slog.Info("reusing existing Ethereum node")
		return ethActiveNode
	}

	var (
		err           error
		gethContainer testcontainers.Container
		ethClient     *ethclient.Client
		rpcURL        *url.URL
		containerURL  *url.URL
	)

	if u := os.Getenv("GETH_NODE_RPC_URL"); u != "" {
		rpcURL, err = url.Parse(u)
		require.NoError(t, err, "failed to parse local RPC URL")
		containerURL = rpcURL
	} else {
		// TODO: use in-memory database instead of container volumes
		gethContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "ethereum/client-go:v1.13.14",
				// 8545 TCP, used by the HTTP based JSON RPC API
				// 8546 TCP, used by the WebSocket based JSON RPC API
				// 8547 TCP, used by the GraphQL API
				// 30303 TCP and UDP, used by the P2P protocol running the network
				ExposedPorts: []string{"8545:8545/tcp", "8546:8546/tcp", "8547:8547/tcp", "30303:30303/tcp", "30303:30303/udp"},
				Cmd:          []string{"--dev", "--http", "--http.api=eth,web3,net", "--http.addr=0.0.0.0", "--http.corsdomain='*'", "--http.vhosts='*'"},
				WaitingFor:   wait.ForLog("server started"),
			},
			Started: true,
		})
		require.NoError(t, err, "failed to start Go-Ethereum container")

		containerIP, err := gethContainer.ContainerIP(ctx)
		require.NoError(t, err, "failed to get Go-Ethereum container IP")
		containerPort, err := gethContainer.MappedPort(ctx, "8545")
		require.NoError(t, err, "failed to get Go-Ethereum container port")

		rpcURL, err = url.Parse(fmt.Sprintf("http://0.0.0.0:%s", containerPort.Port()))
		require.NoError(t, err, "failed to parse local RPC URL")
		containerURL, err = url.Parse(fmt.Sprintf("http://%s:%s", containerIP, containerPort.Port()))
		require.NoError(t, err, "failed to parse container RPC URL")
	}

	ethClient, err = ethclient.Dial(rpcURL.String())
	require.NoError(t, err)

	slog.Info("Go-Ethereum container started", "rpcURL", rpcURL.String())
	ethActiveNode = &EthNode{
		Container:    gethContainer,
		Client:       ethClient,
		LocalURL:     *rpcURL,
		ContainerURL: *containerURL,
	}

	t.Cleanup(ethNodeClenaupFactory(ctx, t))
	ethActiveUsers++

	return ethActiveNode
}

type BundlerNode struct {
	Container    testcontainers.Container
	ContainerURL *url.URL
}

var (
	bundlerActiveNode  *BundlerNode
	bundlerActiveUsers int64
	bundlerMutex       sync.Mutex
)

func bundlerNodeClenaupFactory(ctx context.Context, t *testing.T) func() {
	return func() {
		bundlerMutex.Lock()
		defer bundlerMutex.Unlock()

		bundlerActiveUsers--
		if bundlerActiveUsers <= 0 {
			if err := bundlerActiveNode.Container.Terminate(ctx); err != nil {
				t.Fatalf("failed to terminate Alto container: %s", err)
			}

			bundlerActiveNode = nil
		}
	}
}

func NewBundler(ctx context.Context, t *testing.T, node *EthNode, entryPoint common.Address) *url.URL {
	bundlerMutex.Lock()
	defer bundlerMutex.Unlock()
	slog.Info("starting bundler...")

	if bundlerActiveNode != nil {
		bundlerActiveUsers++
		t.Cleanup(bundlerNodeClenaupFactory(ctx, t))

		slog.Info("reusing existing Alto bundler")
		return bundlerActiveNode.ContainerURL
	}

	const port = "3000"
	balance := decimal.NewFromFloat(100e18 /* 100 ETH */).BigInt()
	bundlerAccount, err := NewAccountWithBalance(ctx, balance, node)
	require.NoError(t, err, "failed to fund bundler account")
	privateKey := hexutil.Encode(crypto.FromECDSA(bundlerAccount.PrivateKey))

	altoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "ghcr.io/pimlicolabs/alto:v1.0.1",
			Entrypoint: []string{
				"pnpm", "start",
				"--port", port,
				"--networkName", "mainnet",
				"--entryPoint", entryPoint.Hex(), // the contract should already be deployed on Go-Ethereum node
				"--signerPrivateKeys", privateKey,
				"--utilityPrivateKey", privateKey,
				"--minBalance", "1000000000000000000", // 1 ETH
				"--rpcUrl", node.ContainerURL.String(),
				"--logLevel", "info",
				"--noEthCallOverrideSupport", "true",
				"--useUserOperationGasLimitsForSubmission", "true",
			},
			ImagePlatform: "linux/amd64",
			ExposedPorts:  []string{fmt.Sprintf("%s:%s/tcp", port, port)},
			WaitingFor:    wait.ForLog("Server listening at"),
		},
		Started: true,
	})
	require.NoError(t, err, "failed to start Alto bundler container")

	containerPort, err := altoContainer.MappedPort(ctx, port)
	require.NoError(t, err, "failed to get Alto bundler container port")
	bundlerURL, err := url.Parse(fmt.Sprintf("http://0.0.0.0:%s", containerPort.Port()))

	bundlerActiveNode = &BundlerNode{
		Container:    altoContainer,
		ContainerURL: bundlerURL,
	}

	t.Cleanup(bundlerNodeClenaupFactory(ctx, t))
	bundlerActiveUsers++

	require.NoError(t, err, "failed to parse local bundler URL")

	return bundlerActiveNode.ContainerURL
}

type Contracts struct {
	EntryPoint          common.Address
	Validator           common.Address
	Factory             common.Address
	Logic               common.Address
	Paymaster           common.Address
	SessionKeyValidator common.Address
}

var (
	cachedContracts *Contracts
	contractsMutex  sync.Mutex
)

func SetupContracts(ctx context.Context, t *testing.T, node *EthNode) Contracts {
	contractsMutex.Lock()
	defer contractsMutex.Unlock()

	if cachedContracts != nil {
		return *cachedContracts
	}
	chainID, err := node.Client.ChainID(ctx)
	require.NoError(t, err)
	slog.Info("chainID", "chainID", chainID)

	balance := decimal.NewFromFloat(50e18 /* 50 ETH */).BigInt()
	owner, err := NewAccountWithBalance(ctx, balance, node)
	require.NoError(t, err, "failed to create owner account")

	entryPoint, _, _, err := entry_point_v0_6_0.DeployEntryPoint(owner.TransactOpts, node.Client)
	require.NoError(t, err)
	slog.Info("deployed EntryPoint contract", "address", entryPoint)

	validator, _, _, err := kernel_ecdsa_validator_v2_2.DeployKernelECDSAValidator(owner.TransactOpts, node.Client)
	require.NoError(t, err)
	slog.Info("deployed KernelECDSAValidator contract", "address", validator)

	logic, _, _, err := kernel_v2_2.DeployKernel(owner.TransactOpts, node.Client, entryPoint)
	require.NoError(t, err)
	slog.Info("deployed Kernel contract", "address", logic)

	factory, _, FactoryContract, err := kernel_factory_v2_2.DeployKernelFactory(owner.TransactOpts, node.Client, owner.Address, entryPoint)
	require.NoError(t, err)
	slog.Info("deployed KernelFactory contract", "address", factory)

	tx, err := FactoryContract.SetImplementation(owner.TransactOpts, logic, true)
	require.NoError(t, err)
	slog.Info("set Kernel implementation", "tx", tx.Hash().Hex())

	var paymaster common.Address // TODO: deploy Paymaster contract

	sessionKeyValidator, _, _, err := session_key_validator.DeploySessionKeyValidator(owner.TransactOpts, node.Client)
	require.NoError(t, err)
	slog.Info("deployed SessionKeyValidator contract", "address", sessionKeyValidator)

	slog.Info("done deploying contracts")
	contracts := Contracts{
		EntryPoint:          entryPoint,
		Validator:           validator,
		Factory:             factory,
		Logic:               logic,
		Paymaster:           paymaster,
		SessionKeyValidator: sessionKeyValidator,
	}

	cachedContracts = &contracts

	return *cachedContracts
}

func buildClient(t *testing.T, rpcURL, bundlerURL url.URL, addresses Contracts) userop.Client {
	config, err := userop.NewClientConfigFromEnv()
	require.NoError(t, err)

	config.ProviderURL = rpcURL
	config.BundlerURL = bundlerURL
	config.EntryPoint = addresses.EntryPoint
	config.SmartWallet = userop.SmartWalletConfig{
		Type:           &userop.SmartWalletKernel,
		ECDSAValidator: addresses.Validator,
		Logic:          addresses.Logic,
		Factory:        addresses.Factory,
	}
	config.Paymaster.Type = &userop.PaymasterDisabled
	config.Paymaster.Address = addresses.Paymaster

	client, err := userop.NewClient(config)
	require.NoError(t, err)
	return client
}
