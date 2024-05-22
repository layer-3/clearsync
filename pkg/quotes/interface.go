// Package quotes implements multiple price feed adapters.
package quotes

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Driver interface {
	ActiveDrivers() []DriverType
	ExchangeType() ExchangeType
	Start() error
	Stop() error
	Subscribe(market Market) error
	Unsubscribe(market Market) error
	SetInbox(inbox <-chan TradeEvent)
}

func NewDriver(config Config, outbox chan<- TradeEvent) (Driver, error) {
	if len(config.Drivers) == 0 {
		return nil, fmt.Errorf("no drivers are configured")
	} else if len(config.Drivers) > 1 {
		return newIndex(config, outbox), nil
	}

	switch config.Drivers[0] {
	case DriverBinance:
		return newBinance(config.Binance, outbox), nil
	case DriverKraken:
		return newKraken(config.Kraken, outbox), nil
	case DriverMexc:
		return newMexc(config.Mexc, outbox), nil
	case DriverOpendax:
		return newOpendax(config.Opendax, outbox), nil
	case DriverBitfaker:
		return newBitfaker(config.Bitfaker, outbox), nil
	case DriverUniswapV3:
		return newUniswapV3(config.UniswapV3, outbox), nil
	case DriverSyncswap:
		return newSyncswap(config.Syncswap, outbox), nil
	case DriverQuickswap:
		return newQuickswap(config.Quickswap, outbox), nil
	case DriverSectaV2:
		return newSectaV2(config.SectaV2, outbox), nil
	case DriverSectaV3:
		return newSectaV3(config.SectaV3, outbox), nil
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
