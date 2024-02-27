package main

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bw6-633/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestRPC(t *testing.T) {
	ctx := context.Background()

	rpcURL, err := startEthNode(ctx, t)
	if err != nil {
		panic(err)
	}

	var entryPoint common.Address // TODO: deploy EntryPoint contract
	var signer ecdsa.PrivateKey   // TODO: generate a private key
	var utility ecdsa.PrivateKey  // TODO: generate a private key

	if err := startBundler(ctx, t, *rpcURL, entryPoint, signer, utility); err != nil {
		panic(err)
	}

	<-time.After(60 * time.Second)
}

func startEthNode(ctx context.Context, t *testing.T) (*url.URL, error) {
	rpcURL, err := url.Parse("http://localhost:8545")
	if err != nil {
		return nil, fmt.Errorf("failed to parse local RPC URL: %w", err)
	}

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
	if err != nil {
		return nil, fmt.Errorf("failed to start Go-Ethereum container: %w", err)
	}

	t.Cleanup(func() {
		if err := gethContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate Go-Ethereum container: %s", err)
		}
	})

	return rpcURL, nil
}

func startBundler(
	ctx context.Context,
	t *testing.T,
	rpcURL url.URL,
	entryPoint common.Address,
	signer ecdsa.PrivateKey,
	utility ecdsa.PrivateKey,
) error {
	altoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "ghcr.io/pimlicolabs/alto:v1.0.1",
			Entrypoint: []string{"pnpm", "start",
				"--networkName", "mainnet", // check geth container logs to find out configured network
				"--entryPoint", entryPoint.Hex(), // the contract should be already deployed on geth node
				"--signerPrivateKeys", hexutils.BytesToHex(signer.Bytes()),
				"--utilityPrivateKey", hexutils.BytesToHex(utility.Bytes()),
				"--minBalance", "0",
				"--rpcUrl", rpcURL.String(),
			},
			WaitingFor: wait.ForLog("server started"),
		},
		Started: true,
	})
	if err != nil {
		return fmt.Errorf("failed to start Alto bundler container: %w", err)
	}

	t.Cleanup(func() {
		if err := altoContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate Alto container: %s", err)
		}
	})

	return nil
}
