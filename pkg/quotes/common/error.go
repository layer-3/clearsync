package common

import "errors"

var (
	ErrNotSubbed     = errors.New("market not subscribed")
	ErrAlreadySubbed = errors.New("market already subscribed")
	ErrFailedSub     = errors.New("failed to subscribe to market")
	ErrFailedUnsub   = errors.New("failed to unsubscribe from market")
)

