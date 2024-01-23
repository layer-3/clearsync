package quotes

import "errors"

var (
	errNotStarted     = errors.New("driver is not started")
	errAlreadyStarted = errors.New("driver is already started")
	errAlreadyStopped = errors.New("driver is already stopped")
	errNotSubbed      = errors.New("market not subscribed")
	errAlreadySubbed  = errors.New("market already subscribed")
	errFailedSub      = errors.New("failed to subscribe to market")
	errFailedUnsub    = errors.New("failed to unsubscribe from market")
)
