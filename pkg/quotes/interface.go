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
	allDrivers := map[DriverType]func(Config, chan<- TradeEvent) Driver{
		DriverIndex:         newIndex,
		DriverBinance:       newBinance,
		DriverKraken:        newKraken,
		DriverOpendax:       newOpendax,
		DriverBitfaker:      newBitfaker,
		DriverUniswapV3Api:  newUniswapV3Api,
		DriverUniswapV3Geth: newUniswapV3Geth,
	}

	driver, ok := allDrivers[config.Driver]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", config.Driver.String())
	}
	return driver(config, outbox), nil
}
