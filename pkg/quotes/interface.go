// Package quotes implements multiple price feed adapters.
package quotes

import "fmt"

type Driver interface {
	Name() DriverType
	Type() Type
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
	case DriverIndex:
		return newIndex(config, outbox), nil
	case DriverBinance:
		return newBinance(config.Binance, outbox), nil
	case DriverKraken:
		return newKraken(config.Kraken, outbox), nil
	case DriverOpendax:
		return newOpendax(config.Opendax, outbox), nil
	case DriverBitfaker:
		return newBitfaker(config.Bitfaker, outbox), nil
	case DriverUniswapV3Api:
		return newUniswapV3Api(config.UniswapV3Api, outbox), nil
	case DriverUniswapV3Geth:
		return newUniswapV3Geth(config.UniswapV3Geth, outbox), nil
	case DriverSyncswap:
		return newSyncswap(config.Syncswap, outbox), nil
	case DriverQuickswap:
		return newQuickswap(config.Quickswap, outbox), nil
	default:
		return nil, fmt.Errorf("driver is not supported: %s", config.Drivers)
	}
}
