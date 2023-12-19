// Package quotes implements multiple price feed adapters.
package quotes

import "fmt"

type Driver interface {
	Name() DriverType
	Start() error
	Stop() error
	Subscribe(market Market) error
	Unsubscribe(market Market) error
}

func NewDriver(config Config, outbox chan<- TradeEvent) (Driver, error) {
	switch config.Driver {
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
	default:
		return nil, fmt.Errorf("driver is not supported: %s", config.Driver)
	}
	return driver(config, outbox), nil
}
