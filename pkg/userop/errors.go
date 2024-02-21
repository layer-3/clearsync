package userop

import "errors"

var (
	// ErrNoWalletDeploymentOpts is returned when the wallet deployment opts
	// are required to build and submit userop but they are not specified.
	ErrNoWalletDeploymentOpts = errors.New("wallet deployment opts not specified")
	// ErrNoSigner is returned when the signer is not specified.
	ErrNoSigner = errors.New("signer not specified")
	// ErrPaymasterNotSupported is returned on attempt to build client with an unsupported paymaster type.
	// Make sure that the paymaster type you are trying to use
	// has no `unsupported` or `untested` tags in `paymaster_type.go` source file.
	ErrPaymasterNotSupported = errors.New("paymaster type not supported")
	// ErrSmartWalletNotSupported is returned on attempt to build client with an unsupported smart wallet type.
	// Make sure that the smart wallet type you are trying to use
	// has no `unsupported` or `untested` tags in `wallet_type.go` source file.
	ErrSmartWalletNotSupported = errors.New("smart wallet not supported")
)
