package userop

import "fmt"

type PaymasterType struct {
	slug string
}

var (
	PaymasterPimlicoERC20       = PaymasterType{"pimlico_erc20"}
	PaymasterPimlicoVerifying   = PaymasterType{"pimlico_verifying"}
	PaymasterBiconomyERC20      = PaymasterType{"biconomy_erc20"}
	PaymasterBiconomySponsoring = PaymasterType{"biconomy_sponsoring"}
	// PaymasterEthInfinitismERC20     = PaymasterType{"eth_infinitism_erc20"}     // unsupported
	// PaymasterEthInfinitismVerifying = PaymasterType{"eth_infinitism_verifying"} // unsupported
)

func (t PaymasterType) String() string {
	return t.slug
}

// UnmarshalYAML unmarshalls the YAML representation of a PaymasterType.
func (t *PaymasterType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawValue string
	err := unmarshal(&rawValue)
	if err != nil {
		return err
	}

	switch rawValue {
	case PaymasterPimlicoERC20.String():
		*t = PaymasterPimlicoERC20
	case PaymasterPimlicoVerifying.String():
		*t = PaymasterPimlicoVerifying
	case PaymasterBiconomyERC20.String():
		*t = PaymasterBiconomyERC20
	case PaymasterBiconomySponsoring.String():
		*t = PaymasterBiconomySponsoring
	default:
		return fmt.Errorf("unknown paymaster type: %s", rawValue)
	}

	return nil
}

// UnmarshalJSON unmarshalls the JSON representation of a PaymasterType.
func (t *PaymasterType) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case PaymasterPimlicoERC20.String():
		*t = PaymasterPimlicoERC20
	case PaymasterPimlicoVerifying.String():
		*t = PaymasterPimlicoVerifying
	case PaymasterBiconomyERC20.String():
		*t = PaymasterBiconomyERC20
	case PaymasterBiconomySponsoring.String():
		*t = PaymasterBiconomySponsoring
	default:
		return fmt.Errorf("unknown paymaster type: %s", string(b))
	}

	return nil
}
