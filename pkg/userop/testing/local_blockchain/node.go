package local_blockchain

import (
	"bufio"
	"context"
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
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/layer-3/clearsync/pkg/abi/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/abi/kernel_ecdsa_validator_v2_2"
	"github.com/layer-3/clearsync/pkg/abi/kernel_factory_v2_2"
	"github.com/layer-3/clearsync/pkg/abi/kernel_v2_2"
	"github.com/layer-3/clearsync/pkg/artifacts/session_key_validator_v2_4"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"

	_ "embed"
)

type EthNode struct {
	Container    testcontainers.Container
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
		// TODO: add test cleanups with container termination
		gethContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "ethereum/client-go:v1.13.14",
				// 8545 TCP, used by the HTTP based JSON RPC API
				// 8546 TCP, used by the WebSocket based JSON RPC API
				// 8547 TCP, used by the GraphQL API
				// 30303 TCP and UDP, used by the P2P protocol running the network
				ExposedPorts: []string{"8545:8545/tcp", "8546:8546/tcp", "8547:8547/tcp", "30303:30303/tcp", "30303:30303/udp"},
				Cmd:          []string{"--dev", "--http", "--ws", "--http.api=eth,web3,net", "--http.addr=0.0.0.0", "--http.corsdomain='*'", "--http.vhosts='*'", "--ws.addr=0.0.0.0", "--ws.origins='*'"},
				WaitingFor:   wait.ForLog("server started"),
			},
			Started: true,
		})
		require.NoError(t, err, "failed to start Go-Ethereum container")

		containerIP, err := gethContainer.ContainerIP(ctx)
		require.NoError(t, err, "failed to get Go-Ethereum container IP")
		// As a rpc port we are using ws port for subscription
		// As a container port we are using http port for bundler
		rpcPort, err := gethContainer.MappedPort(ctx, "8546")
		require.NoError(t, err, "failed to get Go-Ethereum rpc port")
		containerPort, err := gethContainer.MappedPort(ctx, "8545")
		require.NoError(t, err, "failed to get Go-Ethereum container port")

		rpcURL, err = url.Parse(fmt.Sprintf("ws://0.0.0.0:%s", rpcPort.Port()))
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

	return ethActiveNode
}

type BundlerNode struct {
	Container    testcontainers.Container
	ContainerURL *url.URL
}

var (
	bundlerActiveNode *BundlerNode
	bundlerMutex      sync.Mutex
)

func NewBundler(ctx context.Context, t *testing.T, node *EthNode, entryPoint common.Address) *url.URL {
	bundlerMutex.Lock()
	defer bundlerMutex.Unlock()
	slog.Info("starting bundler...")

	if bundlerActiveNode != nil {
		slog.Info("reusing existing Alto bundler")
		return bundlerActiveNode.ContainerURL
	}

	var (
		altoContainer testcontainers.Container
		bundlerURL    *url.URL
	)

	if uEnv := os.Getenv("BUNDLER_RPC_URL"); uEnv != "" {
		u, err := url.Parse(uEnv)
		if err != nil {
			// configuration error, so we shut down the app
			panic(err)
		}

		bundlerURL = u
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

		containerPort, err := altoContainer.MappedPort(ctx, port)
		require.NoError(t, err, "failed to get Alto bundler container port")

		bundlerURL, err = url.Parse(fmt.Sprintf("http://0.0.0.0:%s", containerPort.Port()))
		require.NoError(t, err, "failed to parse local bundler URL")
	}

	bundlerActiveNode = &BundlerNode{
		Container:    altoContainer,
		ContainerURL: bundlerURL,
	}

	return bundlerActiveNode.ContainerURL
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

func buildClient(t *testing.T, rpcURL, bundlerURL url.URL, addresses Contracts) userop.Client {
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
	config.Paymaster.Type = &userop.PaymasterDisabled
	config.Paymaster.Address = addresses.Paymaster

	client, err := userop.NewClient(config)
	require.NoError(t, err)
	return client
}
