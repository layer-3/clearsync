// Package quotes implements multiple price feed adapters.
package quotes

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Driver is an interface that represents trades adapter.
// It is used to stream trades from different exchanges
// and aggregate them into a single outbox channel.
type Driver interface {
	// ActiveDrivers returns all configured data providers.
	ActiveDrivers() []DriverType
	// ExchangeType returns the type of the exchange.
	ExchangeType() ExchangeType

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
	HistoricalData
}

type HistoricalData interface {
	// HistoricalData returns historical trade data for the given market.
	// The returned data is ordered from oldest to newest.
	// The window parameter defines the time range to fetch data for starting from now.
	HistoricalData(ctx context.Context, market Market, window time.Duration, limit uint64) ([]TradeEvent, error)
}

// NewDriver creates a new instance of the driver.
//
// If no drivers appear in the `config.Drivers` array,
// the constructor assumes no drives are configured and returns an error.
//
// Index driver is configured automatically
// if at least one of the following conditions is met:
//  1. `config.Drivers` contains multiple drivers;
//  2. a valid non-nil `inbox` channel is provided;
//
// Params:
//   - config: contains the configuration for the driver(s) to be created
//   - outbox: a channel where the driver sends aggregated trades
//   - inbox:  an optional channel where the package user can send trades from his own source.
//     If you don't { have / need } your own source, pass `nil` here.
//   - trades: an optional adapter to fetch historical data from instead of querying RPC provider,
//     If you don't { have / need } a historical data adapter, pass `nil` here.
func NewDriver(config Config, outbox chan<- TradeEvent, inbox <-chan TradeEvent, history HistoricalData) (Driver, error) {
	// Remove duplicate drivers
	seen := make(map[DriverType]struct{})
	var unique []DriverType
	for _, driver := range config.Drivers {
		if _, ok := seen[driver]; !ok {
			seen[driver] = struct{}{}
			unique = append(unique, driver)
		}
	}
	config.Drivers = unique

	if len(config.Drivers) == 0 {
		return nil, fmt.Errorf("no drivers are configured")
	} else if len(config.Drivers) > 1 || inbox != nil {
		return newIndex(config, outbox, inbox, history)
	}

	switch config.Drivers[0] {
	case DriverBinance:
		return newBinance(config.Binance, outbox, history)
	case DriverKraken:
		return newKraken(config.Kraken, outbox, history)
	case DriverMexc:
		return newMexc(config.Mexc, outbox, history), nil
	case DriverOpendax:
		return newOpendax(config.Opendax, outbox, history)
	case DriverBitfaker:
		return newBitfaker(config.Bitfaker, outbox)
	case DriverUniswapV3:
		return newUniswapV3(config.UniswapV3, outbox, history)
	case DriverSyncswap:
		return newSyncswap(config.Syncswap, outbox, history)
	case DriverQuickswap:
		return newQuickswap(config.Quickswap, outbox, history)
	case DriverSectaV2:
		return newSectaV2(config.SectaV2, outbox, history)
	case DriverSectaV3:
		return newSectaV3(config.SectaV3, outbox, history)
	case DriverLynexV2:
		return newLynexV2(config.LynexV2, outbox), nil
	default:
		return nil, fmt.Errorf("unknown driver: %s", config.Drivers)
	}
}

var MarketSubscriptions = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "price_feed_value",
		Help: "Current trades subscriptions by provider and market.",
	},
	[]string{"provider", "market"}, // labels
)

func recordSubscribed(provider DriverType, market Market) {
	MarketSubscriptions.WithLabelValues(provider.String(), market.String()).Inc()
}

func recordUnsubscribed(provider DriverType, market Market) {
	MarketSubscriptions.WithLabelValues(provider.String(), market.String()).Dec()
}
