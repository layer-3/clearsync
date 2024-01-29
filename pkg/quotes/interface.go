// Package quotes implements multiple price feed adapters.
package quotes

import "fmt"

type Driver interface {
	Start() error
	Stop() error
	Subscribe(market Market) error
	Unsubscribe(market Market) error
}

func NewDriver(config Config, outbox chan<- TradeEvent) (Driver, error) {
	switch config.DriverType() {
	case DriverBinance:
		return newBinance(config.(BinanceConfig), outbox), nil
	case DriverKraken:
		return newKraken(config.(KrakenConfig), outbox), nil
	case DriverOpendax:
		return newOpendax(config.(OpendaxConfig), outbox), nil
	case DriverBitfaker:
		return newBitfaker(config.(BitfakerConfig), outbox), nil
	case DriverUniswapV3Api:
		return newUniswapV3Api(config.(UniswapV3ApiConfig), outbox), nil
	case DriverUniswapV3Geth:
		return newUniswapV3Geth(config.(UniswapV3GethConfig), outbox), nil
	case DriverSyncswap:
		return newSyncswap(config.(SyncswapConfig), outbox), nil
	default:
		return nil, fmt.Errorf("driver is not supported: %s", config.DriverType())
	}
}
