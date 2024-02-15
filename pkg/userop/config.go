package userop

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
)

// ClientConfig represents the configuration for the user operation client.
type ClientConfig struct {
	ProviderURL string
	BundlerURL  string
	ChainID     *big.Int
	EntryPoint  common.Address
	Paymaster   PaymasterConfig
	Signer      Signer
}

// PaymasterConfig represents the configuration for the paymaster.
type PaymasterConfig struct {
	URL     string
	Address common.Address
	Ctx     map[string]any
}

// Signer represents a function that signs a user operation.
type Signer func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error)

// NewClientConfigFromFile reads the
// client configuration from a file.
func NewClientConfigFromFile(path string) (ClientConfig, error) {
	var config ClientConfig
	return config, cleanenv.ReadConfig(path, &config)
}

// NewClientConfigFromEnv reads the client
// configuration from environment variables.
func NewClientConfigFromEnv() (ClientConfig, error) {
	var config ClientConfig
	return config, cleanenv.ReadEnv(&config)
}
