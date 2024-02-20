package userop

import "fmt"

// SmartWalletType represents a type for supported ERC-4337 smart wallets
// that can be used with the client to send user operations from.
type SmartWalletType struct {
	slug string
}

var (
	// SmartWalletSimpleAccount represents a smart wallet type for a Simple Account wallet.
	SmartWalletSimpleAccount = SmartWalletType{"simple_account"}
	// SmartWalletBiconomy represents a type for Biconomy Smart Account wallet.
	SmartWalletBiconomy = SmartWalletType{"biconomy"}
	// SmartWalletKernel represents a type for Zerodev Kernel wallet.
	SmartWalletKernel = SmartWalletType{"kernel"}
)

// String returns the string representation of a SmartWalletType.
func (t SmartWalletType) String() string {
	return t.slug
}

// UnmarshalYAML unmarshals the YAML representation of a SmartWalletType.
func (t *SmartWalletType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawValue string
	err := unmarshal(&rawValue)
	if err != nil {
		return err
	}

	switch rawValue {
	case SmartWalletSimpleAccount.String():
		*t = SmartWalletSimpleAccount
	case SmartWalletBiconomy.String():
		*t = SmartWalletBiconomy
	case SmartWalletKernel.String():
		*t = SmartWalletKernel
	default:
		return fmt.Errorf("unknown smart wallet type: %s", rawValue)
	}

	return nil
}

// UnmarshalJSON unmarshals the JSON representation of a SmartWalletType.
func (t *SmartWalletType) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case SmartWalletSimpleAccount.String():
		*t = SmartWalletSimpleAccount
	case SmartWalletBiconomy.String():
		*t = SmartWalletBiconomy
	case SmartWalletKernel.String():
		*t = SmartWalletKernel
	default:
		return fmt.Errorf("unknown smart wallet type: %s", string(b))
	}

	return nil
}
