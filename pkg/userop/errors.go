package userop

import "errors"

var (
	ErrNoWalletDeploymentOpts = errors.New("wallet deployment opts not specified")
	ErrNoSigner               = errors.New("signer not specified")
	ErrPaymasterNotSupported  = errors.New("paymaster type not supported")
)
