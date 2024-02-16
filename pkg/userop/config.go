package userop

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shopspring/decimal"
)

// ClientConfig represents the configuration for the user operation client.
type ClientConfig struct {
	SmartWalletType SmartWalletType `yaml:"smart_wallet"`
	ProviderURL     string          `yaml:"provider_url"`
	BundlerURL      string          `yaml:"bundler_url"`
	ChainID         *big.Int        `yaml:"chain_id"`
	EntryPoint      common.Address  `yaml:"entry_point"`
	Paymaster       PaymasterConfig `yaml:"paymaster"`
	Signer          Signer
}

// PaymasterConfig represents the configuration for the paymaster.
type PaymasterConfig struct {
	Type    PaymasterType  `yaml:"type"`
	URL     string         `yaml:"url"`
	Address common.Address `yaml:"address"`

	PimlicoERC20       PimlicoERC20Config       `yaml:"pimlico_erc20"`
	PimlicoVerifying   PimlicoVerifyingConfig   `yaml:"pimlico_verifying"`
	BiconomyERC20      BiconomyERC20Config      `yaml:"biconomy_erc20"`
	BiconomySponsoring BiconomySponsoringConfig `yaml:"biconomy_sponsoring"`
}

func (c *PaymasterConfig) init() {
	switch c.Type {
	case PaymasterPimlicoERC20:
		c.PimlicoERC20.init()
	case PaymasterPimlicoVerifying:
		c.PimlicoVerifying.init()
	case PaymasterBiconomyERC20:
		c.BiconomyERC20.init()
	case PaymasterBiconomySponsoring:
		c.BiconomySponsoring.init()
	default:
		panic(fmt.Errorf("unknown paymaster type: %s", c.Type))
	}
}

type PimlicoERC20Config struct {
	// MaxTokenCost specifies the limit for tokens to spend.
	// Operations requiring user to pay more
	// than specified amount of tokens for gas will fail.
	MaxTokenCost decimal.Decimal `json:"maxTokenCost"`
}

func (config *PimlicoERC20Config) init() {
	*config = PimlicoERC20Config{
		MaxTokenCost: decimal.NewFromInt(1_000_000),
	}
}

// PimlicoVerifyingConfig represents the configuration for the Pimlico paymaster.
// See its RPC endpoint docs at https://docs.pimlico.io/paymaster/verifying-paymaster/reference/endpoints#pm_sponsoruseroperation-v2
type PimlicoVerifyingConfig struct {
	SponsorshipPolicyID string `yaml:"sponsorship_policy_id"`
}

func (config *PimlicoVerifyingConfig) init() {
	// no default values
}

type BiconomyERC20Config struct {
	Mode               string            `yaml:"mode"`
	CalculateGasLimits bool              `yaml:"calculate_gas_limits" env-default:"true"`
	TokenInfo          BiconomyTokenInfo `yaml:"token_info"`
}

func (config *BiconomyERC20Config) init() {
	*config = BiconomyERC20Config{
		Mode:               "ERC20",
		CalculateGasLimits: true,
		TokenInfo:          BiconomyTokenInfo{},
	}
}

type BiconomyTokenInfo struct {
	FeeTokenAddress common.Address `yaml:"fee_token_address"`
}

type BiconomySponsoringConfig struct {
	Mode               string                        `yaml:"mode"`
	CalculateGasLimits bool                          `yaml:"calculate_gas_limits"`
	ExpiryDuration     int                           `yaml:"expiry_duration"`
	SponsorshipInfo    BiconomySponsorshipInfoConfig `yaml:"sponsorship_info"`
}

// TODO: implement
func (config *BiconomySponsoringConfig) init() {
	*config = BiconomySponsoringConfig{
		Mode:               "SPONSORED",
		CalculateGasLimits: true,
		ExpiryDuration:     300, // 5 minutes
		SponsorshipInfo: BiconomySponsorshipInfoConfig{
			WebhookData: make(map[string]any),
			SmartAccountInfo: BiconomySmartAccountInfo{
				Name:    "BICONOMY",
				Version: "2.0.0",
			},
		},
	}
}

type BiconomySponsorshipInfoConfig struct {
	WebhookData      map[string]any           `yaml:"webhook_data"`
	SmartAccountInfo BiconomySmartAccountInfo `yaml:"smart_account_info"`
}

type BiconomySmartAccountInfo struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// NewClientConfigFromFile reads the
// client configuration from a file.
func NewClientConfigFromFile(path string) (ClientConfig, error) {
	var config ClientConfig

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return config, err
	}

	config.Paymaster.init()
	return config, nil
}

// NewClientConfigFromEnv reads the client
// configuration from environment variables.
func NewClientConfigFromEnv() (ClientConfig, error) {
	var config ClientConfig

	if err := cleanenv.ReadEnv(&config); err != nil {
		return config, err
	}

	config.Paymaster.init()
	return config, nil
}

// Signer represents a function that signs a user operation.
type Signer func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error)
