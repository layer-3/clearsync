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
	"github.com/layer-3/clearsync/pkg/quotes"
)

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

	driverType := quotes.DriverIndex
	if len(os.Args) == 2 {
		parsedDriver, err := quotes.ToDriverType(os.Args[1])
		if err != nil {
			panic(err)
		}
		driverType = parsedDriver
	}

	config, err := quotes.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}
	config.Driver = driverType

	syncswap, err := quotes.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}
	syncswap.Driver = quotes.DriverSyncswap

	outbox := make(chan quotes.TradeEvent, 128)
	outboxStop := make(chan struct{}, 1)
	go func() {
		// You may add a lot of markets to subscribe
		// and considering imposed rate limits
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

	driver, err := quotes.NewDriver(config, outbox)
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

	markets := []quotes.Market{
		// Add your markets here
		quotes.NewMarket("weth", "usdc"),
		quotes.NewMarket("lube", "usdc"),
		quotes.NewMarket("linda", "usdc"),
	}

	for _, market := range markets {
		if err = driver.Subscribe(market); err != nil {
			slog.Warn("failed to subscribe", "market", market, "err", err)
			continue
		}
		slog.Info("subscribed", "market", market.String())
	}

	slog.Info("waiting for trades")
	<-outboxStop
}
