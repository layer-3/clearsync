package userop

import "errors"

// TODO: look through the code and find errors that can be exported
var (
	ErrNoWalletDeploymentOpts = errors.New("wallet deployment opts not specified")
)
