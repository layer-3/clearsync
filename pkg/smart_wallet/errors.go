package smart_wallet

import "errors"

var (
	// ErrSmartWalletNotSupported is returned on attempt to build client with an unsupported smart wallet type.
	// Make sure that the smart wallet type you are trying to use
	// has no `unsupported` or `untested` tags in `wallet_type.go` source file.
	ErrSmartWalletNotSupported = errors.New("smart wallet not supported")
)
