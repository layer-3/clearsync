// Package driver implements adapters for trade event streaming.
package driver

import (
	"fmt"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

// New creates an instance of trades streaming driver.
//
// If no drivers appear in the `config.Drivers` array,
// the constructor assumes no drives are configured and returns an error.
//
// Params:
//   - config: contains the configuration for the driver(s) to be created
//   - outbox: a channel where the driver sends aggregated trades
//   - external: an optional adapter to fetch historical data from instead of querying RPC provider,
//     If you don't { have / need } a historical data adapter, pass `nil` here.
func New(config Config, outbox chan<- common.TradeEvent, external base.HistoricalDataDriver) (base.Driver, error) {
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
	}

	switch config.Drivers[0] {
	case common.DriverBinance:
		return newBinance(config.Binance, outbox, external)
	case common.DriverKraken:
		return newKraken(config.Kraken, outbox, external)
	case common.DriverMexc:
		return newMexc(config.Mexc, outbox, external), nil
	case common.DriverOpendax:
		return newOpendax(config.Opendax, outbox, external)
	case common.DriverBitfaker:
		return newBitfaker(config.Bitfaker, outbox)
	case common.DriverUniswapV3:
		return newUniswapV3(config.Rpc.Ethereum, config.UniswapV3, outbox, external)
	case common.DriverSyncswap:
		return newSyncswap(config.Rpc.Linea, config.Syncswap, outbox, external)
	case common.DriverQuickswap:
		return newQuickswap(config.Rpc.Polygon, config.Quickswap, outbox, external)
	case common.DriverSectaV2:
		return newSectaV2(config.Rpc.Linea, config.SectaV2, outbox, external)
	case common.DriverSectaV3:
		return newSectaV3(config.Rpc.Linea, config.SectaV3, outbox, external)
	case common.DriverLynexV2:
		return newLynexV2(config.Rpc.Linea, config.LynexV2, outbox, external)
	case common.DriverLynexV3:
		return newLynexV3(config.Rpc.Linea, config.LynexV3, outbox, external)
	default:
		return nil, fmt.Errorf("unknown driver: %s", config.Drivers)
	}
}
