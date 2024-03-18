package userop

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/shopspring/decimal"
)

// ClientConfig represents the configuration
// for the user operation client.
type ClientConfig struct {
	ProviderURL string              `yaml:"provider_url" env:"USEROP_CLIENT_PROVIDER_URL"`
	BundlerURL  string              `yaml:"bundler_url" env:"USEROP_CLIENT_BUNDLER_URL"`
	PollPeriod  time.Duration       `yaml:"poll_period" env:"USEROP_CLIENT_POLL_PERIOD" env-default:"100ms"`
	EntryPoint  common.Address      `yaml:"entry_point" env:"USEROP_CLIENT_ENTRY_POINT"`
	Gas         GasConfig           `yaml:"gas"`
	SmartWallet smart_wallet.Config `yaml:"smart_wallet"`
	Paymaster   PaymasterConfig     `yaml:"paymaster"`
}

func (conf *ClientConfig) Init() {
	conf.PollPeriod = 100 * time.Millisecond
	conf.Gas.Init()
	conf.Paymaster.Init()
	conf.SmartWallet.Init()
}

func (conf ClientConfig) validate() error {
	if conf.PollPeriod.String() == "" {
		return ErrInvalidPollDuration
	}

	if conf.EntryPoint == (common.Address{}) {
		return ErrInvalidEntryPointAddress
	}

	if conf.SmartWallet.Factory == (common.Address{}) {
		return ErrInvalidFactoryAddress
	}

	if conf.SmartWallet.Logic == (common.Address{}) {
		return ErrInvalidLogicAddress
	}

	if conf.SmartWallet.ECDSAValidator == (common.Address{}) {
		return ErrInvalidECDSAValidatorAddress
	}

	if conf.Paymaster.Type != nil && *conf.Paymaster.Type != PaymasterDisabled {
		if conf.Paymaster.Address == (common.Address{}) {
			return ErrInvalidPaymasterAddress
		}
	}

	return nil
}

// GasConfig represents the configuration for the userop transaction gas fees.
type GasConfig struct {
	MaxPriorityFeePerGasMultiplier decimal.Decimal `yaml:"max_priority_fee_per_gas_multiplier" env:"GAS_CONFIG_MAX_PRIORITY_FEE_PER_GAS_MULTIPLIER"` // percentage, 2.42 means 242% increase
	MaxFeePerGasMultiplier         decimal.Decimal `yaml:"max_fee_per_gas_multiplier" env:"GAS_CONFIG_MAX_FEE_PER_GAS_MULTIPLIER"`                   // percentage
}

// Init initializes the GasConfig with default values.
func (c *GasConfig) Init() {
	*c = GasConfig{
		MaxPriorityFeePerGasMultiplier: decimal.RequireFromString("1.13"),
		MaxFeePerGasMultiplier:         decimal.RequireFromString("2"),
	}
}

// PaymasterConfig represents the configuration
// for the paymaster to be used with the client.
type PaymasterConfig struct {
	Type    *PaymasterType `yaml:"type" env:"TYPE"` // nil is equivalent to PaymasterDisabled
	URL     string         `yaml:"url" env:"URL"`
	Address common.Address `yaml:"address" env:"ADDRESS"`

	PimlicoERC20       PimlicoERC20Config       `yaml:"pimlico_erc20" env-prefix:"PAYMASTER_CONFIG_PIMLICO_ERC20_"`
	PimlicoVerifying   PimlicoVerifyingConfig   `yaml:"pimlico_verifying" env-prefix:"PAYMASTER_CONFIG_PIMLICO_VERIFYING_"`
	BiconomyERC20      BiconomyERC20Config      `yaml:"biconomy_erc20" env-prefix:"PAYMASTER_CONFIG_BICONOMY_ERC20_"`
	BiconomySponsoring BiconomySponsoringConfig `yaml:"biconomy_sponsoring" env-prefix:"PAYMASTER_CONFIG_BICONOMY_SPONSORING_"`
}

// Init initializes the PaymasterConfig with default values.
func (c *PaymasterConfig) Init() {
	c.Type = &PaymasterType{}

	c.PimlicoERC20.Init()
	c.PimlicoVerifying.Init()
	c.BiconomyERC20.Init()
	c.BiconomySponsoring.Init()
}

// PimlicoERC20Config represents the configuration for the Pimlico ERC20 paymaster.
type PimlicoERC20Config struct {
	// MaxTokenCost specifies the limit for tokens to spend.
	// Operations requiring user to pay more
	// than specified amount of tokens for gas will fail.
	MaxTokenCost            decimal.Decimal `json:"maxTokenCost" env:"MAX_TOKEN_COST"` // unused for now
	VerificationGasOverhead decimal.Decimal `yaml:"verification_gas_overhead" env:"VERIFICATION_GAS_OVERHEAD"`
}

func (config *PimlicoERC20Config) Init() {
	*config = PimlicoERC20Config{
		MaxTokenCost:            decimal.NewFromInt(1_000_000),
		VerificationGasOverhead: decimal.NewFromInt(10_000),
	}
}

// PimlicoVerifyingConfig represents the configuration for the Pimlico Verifying paymaster.
// See the RPC endpoint docs at https://docs.pimlico.io/paymaster/verifying-paymaster/reference/endpoints#pm_sponsoruseroperation-v2
type PimlicoVerifyingConfig struct {
	SponsorshipPolicyID string `yaml:"sponsorship_policy_id" env:"SPONSORSHIP_POLICY_ID"`
}

func (config *PimlicoVerifyingConfig) Init() {
	// no default values
}

// BiconomyERC20Config represents the configuration for the Biconomy ERC20 paymaster.
// See the RPC endpoint docs at https://docs.biconomy.io/Paymaster/api/sponsor-useroperation#2-mode-is-erc20-
type BiconomyERC20Config struct {
	Mode               string            `yaml:"mode" env:"MODE"`
	CalculateGasLimits bool              `yaml:"calculate_gas_limits" env:"CALCULATE_GAS_LIMITS"`
	TokenInfo          BiconomyTokenInfo `yaml:"token_info" env-prefix:"TOKEN_INFO_"`
}

func (config *BiconomyERC20Config) Init() {
	*config = BiconomyERC20Config{
		Mode:               "ERC20",
		CalculateGasLimits: true,
		TokenInfo:          BiconomyTokenInfo{},
	}
}

// BiconomyTokenInfo represents the token
// used to pay for fees for the Biconomy paymaster.
type BiconomyTokenInfo struct {
	FeeTokenAddress common.Address `yaml:"fee_token_address" env:"FEE_TOKEN"`
}

// BiconomySponsoringConfig represents the configuration for the Biconomy Sponsoring paymaster.
// See the RPC endpoint docs at https://docs.biconomy.io/Paymaster/api/sponsor-useroperation#1-mode-is-sponsored-
type BiconomySponsoringConfig struct {
	Mode               string                        `yaml:"mode" env:"MODE"`
	CalculateGasLimits bool                          `yaml:"calculate_gas_limits" env:"CALCULATE_GAS_LIMITS"`
	ExpiryDuration     int                           `yaml:"expiry_duration" env:"EXPIRY_DURATION"`
	SponsorshipInfo    BiconomySponsorshipInfoConfig `yaml:"sponsorship_info" env-prefix:"SPONSORSHIP_INFO_"`
}

func (config *BiconomySponsoringConfig) Init() {
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
	WebhookData      map[string]any           `yaml:"webhook_data" env:"WEBHOOK_DATA"`
	SmartAccountInfo BiconomySmartAccountInfo `yaml:"smart_account_info" env-prefix:"SMART_ACCOUNT_INFO_"`
}

// BiconomySmartAccountInfo represents the configuration
// for the Biconomy smart contract that sponsors transactions.
type BiconomySmartAccountInfo struct {
	Name    string `yaml:"name" env:"NAME"`
	Version string `yaml:"version" env:"VERSION"`
}

// NewClientConfigFromFile reads the
// client configuration from a file.
func NewClientConfigFromFile(path string) (config ClientConfig, err error) {
	config.Init()
	return config, cleanenv.ReadConfig(path, &config)
}

// NewClientConfigFromEnv reads the client
// configuration from environment variables.
func NewClientConfigFromEnv() (config ClientConfig, err error) {
	config.Init()
	return config, cleanenv.ReadEnv(&config)
}

// ECDSASigner represents a handler that signs a message using ecdsa private key.
type ECDSASigner interface {
	Sign(msg []byte) ([]byte, error)
}

// Signer represents a handler that signs a user operation.
// The handler DOES NOT modify the operation itself,
// but rather builds and returns the signature.
type Signer func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error)
