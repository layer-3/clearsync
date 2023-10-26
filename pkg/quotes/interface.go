// Package quotes implements multiple price feed adapters.
package quotes

import (
	"fmt"

	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
	"github.com/layer-3/neodax/finex/pkg/config"
	"github.com/layer-3/neodax/finex/pkg/event"
	"github.com/layer-3/neodax/finex/pkg/websocket/client"
	"github.com/layer-3/neodax/finex/pkg/websocket/protocol"
)

var logger = log.Logger("quotes")

const (
	DriverBinance  = "binance"
	DriverKraken   = "kraken"
	DriverOpendax  = "opendax"
	DriverBitfaker = "bitfaker"
)

// TODO: remove Finex models dependency
type Driver interface {
	Init(markets cache.Market, outbox chan trade.Event, output chan<- event.Event, config config.QuoteFeed, dialer client.WSDialer) error
	Start() error
	Subscribe(base, quote string) error
	Close() error
}

func DriverFromName(name string) (Driver, error) {
	switch name {
	case DriverBinance:
		return &Binance{}, nil
	case DriverKraken:
		return &Kraken{}, nil
	case DriverOpendax:
		return &Opendax{}, nil
	case DriverBitfaker:
		return &Bitfaker{}, nil
	default:
		return nil, fmt.Errorf("unknown driver %s", name)
	}
}

func GetRoutingEvent(trade trade.Event) (*event.Event, error) {
	tradeMsg := protocol.NewEvent(
		protocol.EventPublic,
		protocol.EventTrade,
		[]interface{}{
			trade.Market,
			trade.ID,
			trade.Price,
			trade.Amount,
			trade.Total,
			trade.CreatedAt,
			trade.TakerType,
			trade.Source,
		},
	)

	encodedMsg, err := tradeMsg.Encode()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	event, err := event.EventFromRoutingKey("public."+trade.Market+".trades", encodedMsg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return event, nil
}
