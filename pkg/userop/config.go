package userop

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shopspring/decimal"
)

// ClientConfig represents the configuration
// for the user operation client.
type ClientConfig struct {
	ProviderURL string          `yaml:"provider_url"`
	BundlerURL  string          `yaml:"bundler_url"`
	EntryPoint  common.Address  `yaml:"entry_point"`
	Gas         GasConfig       `yaml:"gas"`
	SmartWallet SmartWallet     `yaml:"smart_wallet"`
	Paymaster   PaymasterConfig `yaml:"paymaster"`
}

// GasConfig represents the configuration for the userop transaction gas fees.
type GasConfig struct {
	MaxPriorityFeePerGasMultiplier decimal.Decimal `yaml:"max_priority_fee_per_gas_multiplier"` // percentage
	MaxFeePerGasMultiplier         decimal.Decimal `yaml:"max_fee_per_gas_multiplier"`          // percentage
}

// Init initializes the GasConfig with default values.
func (c *GasConfig) Init() {
	*c = GasConfig{
		MaxPriorityFeePerGasMultiplier: decimal.RequireFromString("1.13"),
		MaxFeePerGasMultiplier:         decimal.RequireFromString("2"),
	}
}

// SmartWallet represents the configuration
// for the smart wallet to be used with the client.
type SmartWallet struct {
	Type           SmartWalletType `yaml:"type"`
	ECDSAValidator common.Address  `yaml:"ecdsa_validator"`
	Logic          common.Address  `yaml:"logic"`
	Factory        common.Address  `yaml:"factory"`
}

// PaymasterConfig represents the configuration
// for the paymaster to be used with the client.
type PaymasterConfig struct {
	Type    PaymasterType  `yaml:"type"`
	URL     string         `yaml:"url"`
	Address common.Address `yaml:"address"`

	PimlicoERC20       PimlicoERC20Config       `yaml:"pimlico_erc20"`
	PimlicoVerifying   PimlicoVerifyingConfig   `yaml:"pimlico_verifying"`
	BiconomyERC20      BiconomyERC20Config      `yaml:"biconomy_erc20"`
	BiconomySponsoring BiconomySponsoringConfig `yaml:"biconomy_sponsoring"`
}

// Init initializes the PaymasterConfig with default values.
func (c *PaymasterConfig) Init() {
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

// PimlicoERC20Config represents the configuration for the Pimlico ERC20 paymaster.
type PimlicoERC20Config struct {
	// MaxTokenCost specifies the limit for tokens to spend.
	// Operations requiring user to pay more
	// than specified amount of tokens for gas will fail.
	MaxTokenCost            decimal.Decimal `json:"maxTokenCost"` // unused for now
	VerificationGasOverhead decimal.Decimal `yaml:"verification_gas_overhead"`
}

func (config *PimlicoERC20Config) init() {
	*config = PimlicoERC20Config{
		MaxTokenCost:            decimal.NewFromInt(1_000_000),
		VerificationGasOverhead: decimal.NewFromInt(10_000),
	}
}

// PimlicoVerifyingConfig represents the configuration for the Pimlico Verifying paymaster.
// See the RPC endpoint docs at https://docs.pimlico.io/paymaster/verifying-paymaster/reference/endpoints#pm_sponsoruseroperation-v2
type PimlicoVerifyingConfig struct {
	SponsorshipPolicyID string `yaml:"sponsorship_policy_id"`
}

func (config *PimlicoVerifyingConfig) init() {
	// no default values
}

// BiconomyERC20Config represents the configuration for the Biconomy ERC20 paymaster.
// See the RPC endpoint docs at https://docs.biconomy.io/Paymaster/api/sponsor-useroperation#2-mode-is-erc20-
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

// BiconomyTokenInfo represents the token
// used to pay for fees for the Biconomy paymaster.
type BiconomyTokenInfo struct {
	FeeTokenAddress common.Address `yaml:"fee_token_address"`
}

// BiconomySponsoringConfig represents the configuration for the Biconomy Sponsoring paymaster.
// See the RPC endpoint docs at https://docs.biconomy.io/Paymaster/api/sponsor-useroperation#1-mode-is-sponsored-
type BiconomySponsoringConfig struct {
	Mode               string                        `yaml:"mode"`
	CalculateGasLimits bool                          `yaml:"calculate_gas_limits"`
	ExpiryDuration     int                           `yaml:"expiry_duration"`
	SponsorshipInfo    BiconomySponsorshipInfoConfig `yaml:"sponsorship_info"`
}

func (config *BiconomySponsoringConfig) init() {
	*config = BiconomySponsoringConfig{
		Mode:               "SPONSORED",
		CalculateGasLimits: true,
		ExpiryDuration:     300, // 5 minutes
		SponsorshipInfo: BiconomySponsorshipInfoConfig{
			WebhookData: make(map[string]any),
			SmartAccountInfo: BiconomySmartAccountInfo{
				Name:    "BICONOMY",
				Version: "2.0.0", // NOTE: the version of a smart account affects the bundler's behavior
			},
		},
	}
}

// BiconomySponsorshipInfoConfig represents the configuration
// for transaction sponsoring for the Biconomy Sponsoring paymaster.
// More about webhooks: https://docs.biconomy.io/Paymaster/api/webhookapi
type BiconomySponsorshipInfoConfig struct {
	WebhookData      map[string]any           `yaml:"webhook_data"`
	SmartAccountInfo BiconomySmartAccountInfo `yaml:"smart_account_info"`
}

// BiconomySmartAccountInfo represents the configuration
// for the Biconomy smart contract that sponsors transactions.
type BiconomySmartAccountInfo struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// NewClientConfigFromFile reads the
// client configuration from a file.
func NewClientConfigFromFile(path string) (ClientConfig, error) {
	var config ClientConfig
	config.Gas.Init()
	config.Paymaster.Init()

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return config, err
	}

	return config, nil
}

// NewClientConfigFromEnv reads the client
// configuration from environment variables.
func NewClientConfigFromEnv() (ClientConfig, error) {
	var config ClientConfig
	config.Gas.Init()
	config.Paymaster.Init()

	if err := cleanenv.ReadEnv(&config); err != nil {
		return config, err
	}

	return config, nil
}

// Signer represents a handler that signs a user operation.
// The handler DOES NOT modify the operation itself,
// but rather builds and returns the signature.
type Signer func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error)
