package common

import (
	"context"
	"time"
)

// Driver is an interface that represents trades adapter.
// It is used to stream trades from different exchanges
// and aggregate them into a single outbox channel.
type Driver interface {
	// Type returns the type of the configured data provider
	// and the type of the exchange it represents.
	Type() (DriverType, ExchangeType)
	// Start handles the initialization of the driver.
	// It should be called before any other method.
	Start() error
	// Stop handles the cleanup of the driver.
	// It unsubscribes from all markets and closes all open connections.
	// After calling Stop, the driver can't be used anymore
	// and needs to be again with Start method.
	Stop() error
	// Subscribe establishes a streaming connection to fetch trades for the given market.
	// The driver sends trades to the outbox channel configured in the constructor function.
	// If the market is already subscribed, this method returns an error.
	Subscribe(market Market) error
	// Unsubscribe closes the streaming connection for the given market.
	// After calling this method, the driver won't send any more trades for the given market.
	// If the market is not subscribed yet, this method returns an error.
	Unsubscribe(market Market) error
	HistoricalDataDriver
}

// HistoricalDataDriver is an interface that represents trades adapter
// that can fetch historical data on demand.
type HistoricalDataDriver interface {
	// HistoricalData returns historical trade data for the given market.
	// The returned data is ordered from oldest to newest.
	// The window parameter defines the time range to fetch data for starting from now.
	HistoricalData(ctx context.Context, market Market, window time.Duration, limit uint64) ([]TradeEvent, error)
}
