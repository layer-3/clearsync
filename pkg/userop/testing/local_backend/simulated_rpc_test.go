package local_backend

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
)

func TestRPC(t *testing.T) {
	// Arrange
	n := newSimulatedRPC(t)
	cl, err := rpc.DialOptions(context.Background(), n.HTTPEndpoint())
	require.NoError(t, err)

	// Act
	for i := 0; i < 5; i++ {
		var x uint64
		err = cl.CallContext(context.Background(), &x, "eth_blockNumber")
		require.NoError(t, err)

		// Assert
		fmt.Println(x)
		<-time.After(100 * time.Millisecond)
	}
}

// node.Node -> rpc.API -> SimulatedBackend (ethclient.Client)
func newSimulatedRPC(t *testing.T) *node.Node {
	var secret [32]byte
	if _, err := rand.Read(secret[:]); err != nil {
		t.Fatalf("failed to create jwt secret: %v", err)
	}

	// Geth must read it from a file, and does not support in-memory JWT secrets, so we create a temporary file.
	jwtPath := path.Join(t.TempDir(), "jwt_secret")
	if err := os.WriteFile(jwtPath, []byte(hexutil.Encode(secret[:])), 0600); err != nil {
		t.Fatalf("failed to prepare jwt secret file: %v", err)
	}

	// We get ports assigned by the node automatically
	conf := &node.Config{
		HTTPHost:  "127.0.0.1",
		HTTPPort:  0,
		WSHost:    "127.0.0.1",
		WSPort:    0,
		AuthAddr:  "127.0.0.1",
		AuthPort:  0,
		JWTSecret: jwtPath,

		WSModules:   []string{"eth"},
		HTTPModules: []string{"eth"},
	}
	simulatedNode, err := node.New(conf)
	if err != nil {
		t.Fatalf("could not create a new node: %v", err)
	}

	backend, err := NewSimulatedBackend()
	require.NoError(t, err)

	// register dummy apis, so we can test the modules are available and reachable with authentication
	simulatedNode.RegisterAPIs([]rpc.API{{
		Namespace:     "eth",
		Version:       "1.0",
		Service:       blockNumberRPC{backend},
		Authenticated: false, // no authentication required for a public handler
	}})
	if err := simulatedNode.Start(); err != nil {
		t.Fatalf("failed to start test node: %v", err)
	}

	return simulatedNode
}

type blockNumberRPC struct {
	backend *SimulatedBackend
}

// BlockNumber implements the `eth_blockNumber` method.
func (rpc blockNumberRPC) BlockNumber() (uint64, error) {
	// advance one block
	_ = rpc.backend.Commit()

	number, err := rpc.backend.Client().BlockNumber(context.Background())
	if err != nil {
		return 0, fmt.Errorf("failed to get block number: %w", err)
	}

	return number, nil
}
