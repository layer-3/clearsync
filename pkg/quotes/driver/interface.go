// Package quotes implements multiple price feed adapters.
package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

// Driver is an interface that represents trades adapter.
// It is used to stream trades from different exchanges
// and aggregate them into a single outbox channel.
type Driver interface {
	// ActiveDrivers returns all configured data providers.
	ActiveDrivers() []common.DriverType
	// ExchangeType returns the type of the exchange.
	ExchangeType() common.ExchangeType

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
	Subscribe(market common.Market) error
	// Unsubscribe closes the streaming connection for the given market.
	// After calling this method, the driver won't send any more trades for the given market.
	// If the market is not subscribed yet, this method returns an error.
	Unsubscribe(market common.Market) error
	HistoricalData
}

type HistoricalData interface {
	// HistoricalData returns historical trade data for the given market.
	// The returned data is ordered from oldest to newest.
	// The window parameter defines the time range to fetch data for starting from now.
	HistoricalData(ctx context.Context, market common.Market, window time.Duration, limit uint64) ([]common.TradeEvent, error)
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
func NewDriver(config Config, outbox chan<- common.TradeEvent, inbox <-chan common.TradeEvent, history HistoricalData) (Driver, error) {
	// Remove duplicate drivers
	seen := make(map[common.DriverType]struct{})
	var unique []common.DriverType
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
	case common.DriverBinance:
		return newBinance(config.Binance, outbox, history)
	case common.DriverKraken:
		return newKraken(config.Kraken, outbox, history)
	case common.DriverMexc:
		return newMexc(config.Mexc, outbox, history), nil
	case common.DriverOpendax:
		return newOpendax(config.Opendax, outbox, history)
	case common.DriverBitfaker:
		return newBitfaker(config.Bitfaker, outbox)
	case common.DriverUniswapV3:
		return newUniswapV3(config.Rpc.Ethereum, config.UniswapV3, outbox, history)
	case common.DriverSyncswap:
		return newSyncswap(config.Rpc.Linea, config.Syncswap, outbox, history)
	case common.DriverQuickswap:
		return newQuickswap(config.Rpc.Polygon, config.Quickswap, outbox, history)
	case common.DriverSectaV2:
		return newSectaV2(config.Rpc.Linea, config.SectaV2, outbox, history)
	case common.DriverSectaV3:
		return newSectaV3(config.Rpc.Linea, config.SectaV3, outbox, history)
	case common.DriverLynexV2:
		return newLynexV2(config.Rpc.Linea, config.LynexV2, outbox, history)
	case common.DriverLynexV3:
		return newLynexV3(config.Rpc.Linea, config.LynexV3, outbox, history)
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

func recordSubscribed(provider common.DriverType, market common.Market) {
	MarketSubscriptions.WithLabelValues(provider.String(), market.String()).Inc()
}

func recordUnsubscribed(provider common.DriverType, market common.Market) {
	MarketSubscriptions.WithLabelValues(provider.String(), market.String()).Dec()
}
