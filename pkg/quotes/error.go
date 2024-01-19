package quotes

import "errors"

var (
	errNotSubbed     = errors.New("market not subscribed")
	errAlreadySubbed = errors.New("market already subscribed")
	errFailedSub     = errors.New("failed to subscribe to market")
	errFailedUnsub   = errors.New("failed to unsubscribe from market")
)
