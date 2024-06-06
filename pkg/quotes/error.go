package quotes

import "errors"

var (
	ErrNotStarted     = errors.New("driver is not started; call `Start()` first or wait for it to finish")
	ErrAlreadyStarted = errors.New("driver is already started")
	ErrAlreadyStopped = errors.New("driver is already stopped")
	ErrInvalidWsUrl   = errors.New("websocket URL must start with ws:// or wss://")
	ErrNotSubbed      = errors.New("market not subscribed")
	ErrAlreadySubbed  = errors.New("market already subscribed")
	ErrFailedSub      = errors.New("failed to subscribe to market")
	ErrFailedUnsub    = errors.New("failed to unsubscribe from market")
	ErrSwapParsing    = errors.New("recovered in from panic during swap parsing")
	ErrMarketDisabled = errors.New("market is disabled")
)
