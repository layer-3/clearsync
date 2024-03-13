package smart_wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// SmartWalletConfig represents the configuration
// for the smart wallet to be used with the client.
type Config struct {
	Type           *Type          `yaml:"type" env:"TYPE"`
	ECDSAValidator common.Address `yaml:"ecdsa_validator" env:"ECDSA_VALIDATOR"`
	Logic          common.Address `yaml:"logic" env:"LOGIC"`
	Factory        common.Address `yaml:"factory" env:"FACTORY"`
}

func (sw *Config) Init() {
	sw.Type = &Type{}
}

// Type represents an enum for supported ERC-4337 smart wallets
// that can be used with the client to send user operations from.
type Type struct {
	slug string
}

var (
	// SimpleAccountType represents a smart wallet type for a Simple Account wallet.
	SimpleAccountType = Type{"simple_account"}
	// BiconomyType represents a type for BiconomyType Smart Account wallet.
	BiconomyType = Type{"biconomy"}
	// KernelType represents a type for Zerodev KernelType wallet.
	KernelType = Type{"kernel"}
)

// String returns the string representation of a SmartWalletType.
func (t Type) String() string {
	return t.slug
}

// UnmarshalYAML unmarshals the YAML representation of a SmartWalletType.
func (t *Type) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawValue string
	err := unmarshal(&rawValue)
	if err != nil {
		return err
	}

	switch rawValue {
	case SimpleAccountType.String():
		*t = SimpleAccountType
	case BiconomyType.String():
		*t = BiconomyType
	case KernelType.String():
		*t = KernelType
	default:
		return fmt.Errorf("unknown smart wallet type: %s", rawValue)
	}

	return nil
}

// UnmarshalJSON unmarshals the JSON representation of a SmartWalletType.
func (t *Type) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case SimpleAccountType.String():
		*t = SimpleAccountType
	case BiconomyType.String():
		*t = BiconomyType
	case KernelType.String():
		*t = KernelType
	default:
		return fmt.Errorf("unknown smart wallet type: %s", string(b))
	}

	return nil
}

// SetValue implements the cleanenv.Setter interface.
func (t *Type) SetValue(s string) error {
	switch s {
	case SimpleAccountType.String():
		*t = SimpleAccountType
	case BiconomyType.String():
		*t = BiconomyType
	case KernelType.String():
		*t = KernelType
	default:
		return fmt.Errorf("unknown smart wallet type: %s", s)
	}

	return nil
}
