// Package quotes implements multiple price feed adapters.
package quotes

import (
	"fmt"

	"github.com/layer-3/clearsync/pkg/quotes/binance"
	"github.com/layer-3/clearsync/pkg/quotes/bitfaker"
	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/kraken"
	"github.com/layer-3/clearsync/pkg/quotes/opendax"
	"github.com/layer-3/clearsync/pkg/quotes/uniswap"
)

type Driver interface {
	Start() error
	Stop() error
	Subscribe(market common.Market) error
	Unsubscribe(market common.Market) error
}

func NewDriver(config common.Config, outbox chan<- common.TradeEvent) (Driver, error) {
	allDrivers := map[common.DriverType]Driver{
		common.DriverBinance:   binance.New(config, outbox),
		common.DriverKraken:    kraken.New(config, outbox),
		common.DriverOpendax:   opendax.New(config, outbox),
		common.DriverBitfaker:  bitfaker.New(config, outbox),
		common.DriverUniswapV3: uniswap.New(config, outbox),
	}

	driver, ok := allDrivers[config.Driver]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", config.Driver.String())
	}
	return driver, nil
}
