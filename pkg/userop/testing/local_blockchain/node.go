package local_blockchain

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"math/big"
	"net/url"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/layer-3/clearsync/pkg/artifacts/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/artifacts/kernel_ecdsa_validator_v2_2"
	"github.com/layer-3/clearsync/pkg/artifacts/kernel_factory_v2_2"
	"github.com/layer-3/clearsync/pkg/artifacts/kernel_v2_2"
	"github.com/layer-3/clearsync/pkg/artifacts/session_key_validator_v2_4"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
)

// These addresses are hardcoded in the Paymaster contract and determined by CREATE2
const ENTRY_POINT_ADDRESS_V06 = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"

type EthNode struct {
	Container    testcontainers.Container
	RPCClient    *rpc.Client
	Client       *ethclient.Client
	LocalURL     url.URL
	ContainerURL url.URL

	mu sync.Mutex
}

type MockedClient struct {
	ethclient.Client
	mu sync.Mutex
}

func (m *MockedClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Client.SendTransaction(ctx, tx)
}

func (n *EthNode) FundAccount(ctx context.Context, to Account, amount *big.Int) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if pkStr := os.Getenv("DEPLOYER_PK"); pkStr != "" {
		privateKey, err := crypto.HexToECDSA(pkStr)
		if err != nil {
			return fmt.Errorf("failed to parse deployer private key: %w", err)
		}

		deployerAccount := Account{
			PrivateKey: privateKey,
			Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		}

		err = SendNative(ctx, n, deployerAccount, to, amount)
		if err != nil {
			return fmt.Errorf("failed to send native: %w", err)
		}

		return nil
	}

	gethCmd := fmt.Sprintf(
		"eth.sendTransaction({from: eth.coinbase, to: '%s', value: web3.toWei(%d, 'wei')})",
		to.Address, amount.Uint64(),
	)

	exitCode, result, err := n.Container.Exec(ctx, []string{"geth", "attach", "--exec", gethCmd, n.LocalURL.String()})
	if err != nil || exitCode != 0 {
		return fmt.Errorf("failed to exec increment balance: %w (exit code %d)", err, exitCode)
	}

	scanner := bufio.NewScanner(result)
	var txHash string
	for scanner.Scan() && txHash == "" {
		txHash = regexp.MustCompile("0x[0-9a-fA-F]{64}").FindString(scanner.Text())
	}

	if txHash == "" {
		return fmt.Errorf("failed to find transaction hash in geth output")
	}

	tx, _, err := n.Client.TransactionByHash(context.Background(), common.Hash(hexutil.MustDecode(txHash)))
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}

	return nil
}

var (
	ethActiveNode *EthNode
	ethMutex      sync.Mutex
)

func NewEthNode(ctx context.Context, t *testing.T) *EthNode {
	ethMutex.Lock()
	defer ethMutex.Unlock()
	slog.Info("starting Go-Ethereum node...")

	if ethActiveNode != nil {
		slog.Info("reusing existing Ethereum node")
		return ethActiveNode
	}

	var (
		err            error
		anvilContainer testcontainers.Container
		ethClient      *ethclient.Client
		rpcURL         *url.URL
		containerURL   *url.URL
	)

	if u := os.Getenv("GETH_NODE_RPC_URL"); u != "" {
		rpcURL, err = url.Parse(u)
		require.NoError(t, err, "failed to parse local RPC URL")
		containerURL = rpcURL
	} else {
		// TODO: use in-memory database instead of container volumes
		// TODO: add test cleanups with container termination
		anvilContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Name:  "anvil",
				Image: "quay.io/openware/geth:v0.1.7-amd64",
				// 8545 TCP, used by the HTTP based JSON RPC API
				// 8546 TCP, used by the WebSocket based JSON RPC API
				// 8547 TCP, used by the GraphQL API
				// 30303 TCP and UDP, used by the P2P protocol running the network
				ExposedPorts: []string{"8545:8545/tcp"},
				Cmd:          []string{"anvil"},
				Env: map[string]string{
					"ANVIL_IP_ADDR": "0.0.0.0",
				},
				WaitingFor: wait.ForLog("Listening on"),
			},
			Started: true,
		})
		require.NoError(t, err, "failed to start Go-Ethereum container")

		// Hardcoded deployer private key for testing purposes
		os.Setenv("DEPLOYER_PK", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")

		containerIP, err := anvilContainer.ContainerIP(ctx)
		require.NoError(t, err, "failed to get Go-Ethereum container IP")
		// As a rpc port we are using ws port for subscription
		// As a container port we are using http port for bundler
		rpcPort, err := anvilContainer.MappedPort(ctx, "8545")
		require.NoError(t, err, "failed to get Go-Ethereum rpc port")

		rpcURL, err = url.Parse(fmt.Sprintf("ws://0.0.0.0:%s", rpcPort.Port()))
		require.NoError(t, err, "failed to parse local RPC URL")
		containerURL, err = url.Parse(fmt.Sprintf("http://%s:%s", containerIP, rpcPort.Port()))
		require.NoError(t, err, "failed to parse container RPC URL")
	}

	rpcClient, err := rpc.DialContext(context.Background(), rpcURL.String())
	require.NoError(t, err)

	ethClient = ethclient.NewClient(rpcClient)

	slog.Info("Go-Ethereum container started", "rpcURL", rpcURL.String())
	ethActiveNode = &EthNode{
		Container:    anvilContainer,
		Client:       ethClient,
		RPCClient:    rpcClient,
		LocalURL:     *rpcURL,
		ContainerURL: *containerURL,
	}

	return ethActiveNode
}

type BundlerNode struct {
	Container    testcontainers.Container
	LocalURL     url.URL
	ContainerURL url.URL
}

var (
	bundlerActiveNode *BundlerNode
	bundlerMutex      sync.Mutex
)

func NewBundler(ctx context.Context, t *testing.T, node *EthNode, entryPoint common.Address) *BundlerNode {
	bundlerMutex.Lock()
	defer bundlerMutex.Unlock()
	slog.Info("starting bundler...")

	if bundlerActiveNode != nil {
		slog.Info("reusing existing Alto bundler")
		return bundlerActiveNode
	}

	var (
		err           error
		altoContainer testcontainers.Container
		rpcURL        *url.URL
		containerURL  *url.URL
	)

	if uEnv := os.Getenv("BUNDLER_RPC_URL"); uEnv != "" {
		rpcURL, err = url.Parse(uEnv)
		if err != nil {
			// configuration error, so we shut down the app
			panic(err)
		}

		containerURL = rpcURL
	} else {
		const port = "3000"
		balance := decimal.NewFromFloat(100e18 /* 100 ETH */).BigInt()
		bundlerAccount, err := NewAccountWithBalance(ctx, balance, node)
		require.NoError(t, err, "failed to fund bundler account")
		privateKey := hexutil.Encode(crypto.FromECDSA(bundlerAccount.PrivateKey))

		altoContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "quay.io/openware/bundler:c7dd933",
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

		containerIP, err := altoContainer.ContainerIP(ctx)
		require.NoError(t, err, "failed to get Go-Ethereum container IP")
		// As a rpc port we are using ws port for subscription
		// As a container port we are using http port for other services
		rpcPort, err := altoContainer.MappedPort(ctx, port)
		require.NoError(t, err, "failed to get Go-Ethereum rpc port")

		rpcURL, err = url.Parse(fmt.Sprintf("http://0.0.0.0:%s", rpcPort.Port()))
		require.NoError(t, err, "failed to parse local RPC URL")

		containerURL, err = url.Parse(fmt.Sprintf("http://%s:%s", containerIP, rpcPort.Port()))
		require.NoError(t, err, "failed to parse container RPC URL")
	}

	bundlerActiveNode = &BundlerNode{
		Container:    altoContainer,
		ContainerURL: *containerURL,
		LocalURL:     *rpcURL,
	}

	return bundlerActiveNode
}

type Contracts struct {
	EntryPoint          common.Address
	ECDSAValidator      common.Address
	SessionKeyValidator common.Address
	Logic               common.Address
	Factory             common.Address
	Paymaster           common.Address
}

var (
	cachedContracts *Contracts
	contractsMutex  sync.Mutex
)

func SetupContracts(ctx context.Context, t *testing.T, node *EthNode) Contracts {
	contractsMutex.Lock()
	defer contractsMutex.Unlock()

	if v := os.Getenv("BUNDLER_USE_HARDCODED_CONTRACTS"); v == "true" {
		return getContractAddressesFromEnv()
	}

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
	// NOTE: it will only work with local anvil node
	// Setting code for the hardcoded entry point address
	entryPointCode, err := node.Client.CodeAt(ctx, entryPoint, nil)
	require.NoError(t, err)

	err = node.RPCClient.CallContext(ctx, nil, "anvil_setCode", ENTRY_POINT_ADDRESS_V06, hexutil.Encode(entryPointCode))
	require.NoError(t, err)

	entryPoint = common.HexToAddress(ENTRY_POINT_ADDRESS_V06)
	slog.Info("deployed EntryPoint contract", "address", entryPoint)

	ecdsaValidator, _, _, err := kernel_ecdsa_validator_v2_2.DeployKernelECDSAValidator(owner.TransactOpts, node.Client)
	require.NoError(t, err)
	slog.Info("deployed KernelECDSAValidator contract", "address", ecdsaValidator)

	sessionKeyValidator, _, _, err := session_key_validator_v2_4.DeploySessionKeyValidator(owner.TransactOpts, node.Client)
	require.NoError(t, err)
	slog.Info("deployed SessionKeyValidator contract", "address", sessionKeyValidator)

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

	slog.Info("done deploying contracts")
	contracts := Contracts{
		EntryPoint:          entryPoint,
		ECDSAValidator:      ecdsaValidator,
		SessionKeyValidator: sessionKeyValidator,
		Logic:               logic,
		Factory:             factory,
		Paymaster:           paymaster,
	}

	cachedContracts = &contracts

	return *cachedContracts
}

func getContractAddressesFromEnv() Contracts {
	return Contracts{
		EntryPoint:          common.HexToAddress(os.Getenv("ENTRY_POINT_ADDRESS")),
		ECDSAValidator:      common.HexToAddress(os.Getenv("KERNEL_ECDSA_VALIDATOR_ADDRESS")),
		Factory:             common.HexToAddress(os.Getenv("KERNEL_FACTORY_ADDRESS")),
		Logic:               common.HexToAddress(os.Getenv("KERNEL_ADDRESS")),
		Paymaster:           common.HexToAddress(os.Getenv("PAYMASTER_ADDRESS")),
		SessionKeyValidator: common.HexToAddress(os.Getenv("SESSION_KEY_VALIDATOR_ADDRESS")),
	}
}

func defaultClientConfig(t *testing.T, rpcURL, bundlerURL url.URL, addresses Contracts, paymasterConfig userop.PaymasterConfig) userop.ClientConfig {
	config, err := userop.NewClientConfigFromEnv()
	require.NoError(t, err)

	config.ProviderURL = rpcURL.String()
	config.BundlerURL = bundlerURL.String()
	config.EntryPoint = addresses.EntryPoint
	config.SmartWallet = smart_wallet.Config{
		Type:           &smart_wallet.KernelType,
		ECDSAValidator: addresses.ECDSAValidator,
		Logic:          addresses.Logic,
		Factory:        addresses.Factory,
	}
	config.Paymaster = paymasterConfig

	return config
}

func BuildClient(t *testing.T, rpcURL, bundlerURL url.URL, addresses Contracts, paymasterConfig userop.PaymasterConfig) userop.Client {
	config := defaultClientConfig(t, rpcURL, bundlerURL, addresses, paymasterConfig)
	client, err := userop.NewClient(config)
	require.NoError(t, err)
	return client
}

func SetupPaymaster(ctx context.Context, t *testing.T, node *EthNode, bundler *BundlerNode) *url.URL {
	var (
		err    error
		rpcURL *url.URL
	)

	if u := os.Getenv("PAYMASTER_RPC_URL"); u != "" {
		rpcURL, err = url.Parse(u)
		require.NoError(t, err, "failed to parse local RPC URL")
	} else {
		const port = "3001"
		paymasterContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "quay.io/openware/paymaster:v0.1.0-mock-amd64",
				Env: map[string]string{
					"ALTO_RPC":  bundler.ContainerURL.String(),
					"ANVIL_RPC": node.ContainerURL.String(),
					"PORT":      port,
				},
				ImagePlatform: "linux/amd64",
				ExposedPorts:  []string{fmt.Sprintf("%s:%s/tcp", port, port)},
				WaitingFor:    wait.ForLog("Listening on"),
			},
			Started: true,
		})
		require.NoError(t, err, "failed to start Paymaster container")

		rpcPort, err := paymasterContainer.MappedPort(ctx, port)
		require.NoError(t, err, "failed to get Paymaster container port")

		rpcURL, err = url.Parse(fmt.Sprintf("http://0.0.0.0:%s", rpcPort.Port()))
		require.NoError(t, err, "failed to parse local Paymaster URL")
	}

	return rpcURL
}
