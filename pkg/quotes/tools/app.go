// App for testing quotes drivers.
// Set your driver here or as a console argument: `go run . index`
package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver"
)

// Usage example: `go run . binance syncswap`
func main() {
	go func() {
		// Start pprof server
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			panic(err)
		}
	}()

	if err := log.SetLogLevel("*", "info"); err != nil {
		panic(err)
	}

	var drivers []common.DriverType
	if len(os.Args) >= 2 {
		drivers = make([]common.DriverType, 0, len(os.Args[1:]))
		for _, arg := range os.Args[1:] {
			parsedDriver, err := common.ToDriverType(arg)
			if err != nil {
				panic(err)
			}
			drivers = append(drivers, parsedDriver)
		}
	}

	config, err := driver.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}
	if len(drivers) > 0 {
		// Override default values only if drivers are provided
		config.Drivers = drivers
	}

	outbox := make(chan common.TradeEvent, 128)
	outboxStop := make(chan struct{}, 1)
	go func() {
		// You may add a lot of markets to subscribe
		// and considering imposed rate limits of your RPC provider
		// it may take a while to get the first trade
		// if you run outbox processing AFTER subscriptions.
		// That's why we start processing in an async manner beforehand.
		for e := range outbox {
			slog.Info("new trade",
				"source", e.Source,
				"market", e.Market,
				"side", e.TakerType.String(),
				"price", e.Price.String(),
				"amount", e.Amount.String())
		}
		outboxStop <- struct{}{}
	}()

	driver, err := driver.NewDriver(config, outbox, nil, nil)
	if err != nil {
		panic(err)
	}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	slog.Info("starting", "config", jsonConfig)

	if err := driver.Start(); err != nil {
		panic(err)
	}

	markets := []common.Market{
		// Add your markets here
		common.NewMarket("eth", "usd"),
		common.NewMarket("btc", "usd"),
		common.NewMarket("lube", "usdc"),
		common.NewMarket("linda", "usdc"),
	}

	atLeastOne := false
	for _, market := range markets {
		if err = driver.Subscribe(market); err != nil {
			slog.Warn("failed to subscribe", "market", market, "err", err)
			continue
		}
		atLeastOne = true
		slog.Info("subscribed", "market", market.String())
	}

	if !atLeastOne {
		panic("failed to subscribe to at least one market")
	}

	slog.Info("waiting for trades")
	<-outboxStop
}
