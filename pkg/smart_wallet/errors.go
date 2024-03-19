package smart_wallet

import "errors"

var (
	// ErrNoSmartWalletType is returned when no smart wallet type is specified.
	// Make sure to specify the smart wallet type in the configuration.
	ErrNoSmartWalletType = errors.New("no smart wallet type specified (nil)")
	// ErrSmartWalletNotSupported is returned on attempt to perform some actions with an unsupported smart wallet type.
	// Make sure that the smart wallet type you are trying to use
	// has no `unsupported` or `untested` tags in the source files.
	ErrSmartWalletNotSupported = errors.New("smart wallet not supported")
)
