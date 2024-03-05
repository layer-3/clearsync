package local_blockchain

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"
)

func TestSimulatedRPC(t *testing.T) {
	ctx := context.Background()

	// 1. Start a local Ethereum node
	for i := 0; i < 3; i++ { // starting multiple nodes to test reusing existing nodes
		ethNode := NewEthNode(ctx, t)
		slog.Info("connecting to Ethereum node", "rpcURL", ethNode.LocalURL.String())
	}
	ethNode := NewEthNode(ctx, t)
	slog.Info("connecting to Ethereum node", "rpcURL", ethNode.LocalURL.String())

	// 2. Deploy the required contracts
	addresses := SetupContracts(ctx, t, ethNode)

	// 3. Start the bundler
	for i := 0; i < 3; i++ { // starting multiple bundlers to test reusing existing bundlers
		bundlerURL := NewBundler(ctx, t, ethNode.ContainerURL, addresses.entryPoint)
		slog.Info("connecting to bundler", "bundlerURL", bundlerURL.String())
	}
	bundlerURL := NewBundler(ctx, t, ethNode.ContainerURL, addresses.entryPoint)

	// 4. Run transactions
	// privateKey, err := crypto.HexToECDSA("26b556ff5c77f622504ed5e474919db6e4533fdc62b2f5965a26a6b22eb86f3f")
	// require.NoError(t, err, "failed to parse private key")
	// signer := userop.SignerForKernel(newExampleECDSASigner(privateKey))
	// sender := common.HexToAddress("")
	// calls := []userop.Call{}
	// params := &userop.WalletDeploymentOpts{Index: decimal.Zero, Owner: common.Address{}}

	_ = buildClient(t, ethNode.LocalURL, *bundlerURL, addresses)
	// op, err := client.NewUserOp(ctx, sender, signer, calls, params)
	// require.NoError(t, err, "failed to create new user operation")
	// done, err := client.SendUserOp(ctx, op)
	// require.NoError(t, err, "failed to send user operation")
	//
	// receipt := <-done
	// slog.Info("transaction mined", "receipt", receipt)
	fmt.Println("waiting for 60 seconds")
	<-time.After(60 * time.Second)
}
