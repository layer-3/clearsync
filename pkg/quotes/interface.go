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
	allDrivers := map[DriverType]Driver{
		DriverBinance:       newBinance(config, outbox),
		DriverKraken:        newKraken(config, outbox),
		DriverOpendax:       newOpendax(config, outbox),
		DriverBitfaker:      newBitfaker(config, outbox),
		DriverUniswapV3Api:  newUniswapV3Api(config, outbox),
		DriverUniswapV3Geth: newUniswapV3Geth(config, outbox),
	}

	driver, ok := allDrivers[config.Driver]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", config.Driver.String())
	}
	return driver, nil
}
