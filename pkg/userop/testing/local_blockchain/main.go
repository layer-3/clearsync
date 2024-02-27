package local_backend

import (
	"fmt"
	"os/exec"
)

func main() {
	// Start blockchain RPC
	_, err := NewSimulatedBackend()
	if err != nil {
		panic(fmt.Errorf("failed to build simulated backend: %v", err))
	}

	if err := exec.Command("geth", "--dev", "--http", "--http.api=eth,web3,net").Run(); err != nil {
		return
	}

	// Start bundler RPC
	if err := exec.Command("forge").Run(); err != nil {
		return
	}

	// Start paymaster RPC
	// if err := exec.Command("forge").Run(); err != nil {
	// 	return
	// }
}
