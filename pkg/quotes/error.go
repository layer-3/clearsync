package quotes

import "errors"

var (
	errNotStarted     = errors.New("driver is not started; call `Start()` first or wait for it to finish")
	errAlreadyStarted = errors.New("driver is already started")
	errAlreadyStopped = errors.New("driver is already stopped")
	errInvalidWsUrl   = errors.New("websocket URL must start with ws:// or wss://")
	errNotSubbed      = errors.New("market not subscribed")
	errAlreadySubbed  = errors.New("market already subscribed")
	errFailedSub      = errors.New("failed to subscribe to market")
	errFailedUnsub    = errors.New("failed to unsubscribe from market")
)
