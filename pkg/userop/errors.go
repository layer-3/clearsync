package userop

import "errors"

var (
	// ErrNoWalletDeploymentOpts is returned when the wallet deployment opts
	// are required to build and submit userop, but they are not specified.
	ErrNoWalletDeploymentOpts = errors.New("wallet deployment opts not specified")
	// ErrNoWalletOwnerInWDO is returned when the wallet owner is not specified.
	ErrNoWalletOwnerInWDO = errors.New("wallet deployment opts: wallet owner not specified")
	// ErrInvalidPollDuration is returned when blockchain poll duration is non-present or its format is invalid.
	ErrInvalidPollDuration = errors.New("invalid blockchain poll duration")
	// ErrInvalidEntryPointAddress is returned when the entrypoint address is invalid.
	ErrInvalidEntryPointAddress = errors.New("invalid entry point address")
	// ErrInvalidFactoryAddress is returned when the factory address is invalid.
	ErrInvalidFactoryAddress = errors.New("invalid factory address")
	// ErrInvalidLogicAddress is returned when the logic address is invalid.
	ErrInvalidLogicAddress = errors.New("invalid logic address")
	// ErrInvalidECDSAValidatorAddress is returned when the ECDSA validator address is invalid.
	ErrInvalidECDSAValidatorAddress = errors.New("invalid ECDSA validator address")
	// ErrInvalidPaymasterURL is returned when the paymaster URL is invalid.
	ErrInvalidPaymasterURL = errors.New("invalid paymaster URL")
	// ErrInvalidPaymasterAddress is returned when the paymaster address is invalid.
	ErrInvalidPaymasterAddress = errors.New("invalid paymaster address")
	// ErrNoSigner is returned when the signer is not specified.
	ErrNoSigner = errors.New("signer not specified")
	// ErrNoCalls is returned when the calls are not specified.
	ErrNoCalls = errors.New("calls not specified")
	// ErrPaymasterNotSupported is returned on attempt to build client with an unsupported paymaster type.
	// Make sure that the paymaster type you are trying to use
	// has no `unsupported` or `untested` tags in `paymaster_type.go` source file.
	ErrPaymasterNotSupported = errors.New("paymaster type not supported")
)
