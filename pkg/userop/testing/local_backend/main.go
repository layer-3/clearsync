package local_backend

import "os/exec"

func main() {
	// Start blockchain RPC
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
