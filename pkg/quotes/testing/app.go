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

	var drivers []quotes.DriverType
	if len(os.Args) >= 2 {
		drivers = make([]quotes.DriverType, 0, len(os.Args[1:]))
		for _, arg := range os.Args[1:] {
			parsedDriver, err := quotes.ToDriverType(arg)
			if err != nil {
				panic(err)
			}
			drivers = append(drivers, parsedDriver)
		}
	}

	config, err := quotes.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}

	if len(drivers) > 0 {
		// Override default values only if drivers are provided
		config.Drivers = drivers
	}

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
		quotes.NewMarket("weth", "usd"),
		// quotes.NewMarket("lube", "usdc"),
		// quotes.NewMarket("linda", "usdc"),
	}

	// atLeastOne := false
	for _, market := range markets {
		if err = driver.Subscribe(market); err != nil {
			slog.Warn("failed to subscribe", "market", market, "err", err)
			continue
		}
		// atLeastOne = true
		slog.Info("subscribed", "market", market.String())
	}

	// if !atLeastOne {
	// 	panic("failed to subscribe to at least one market")
	// }

	slog.Info("waiting for trades")
	<-outboxStop
}

// in memory [{3045.2119233911228102 4.4068 0 binance 2024-04-17 15:53:39.982798 +0200 CEST m=+5.017420459}]
// panic: decimal division by 0

// goroutine 38 [running]:
// github.com/shopspring/decimal.Decimal.QuoRem({0x14000456b60, 0xfffffff0}, {0x14000456ba0, 0x0}, 0x10)
//         /Users/dimast/go/pkg/mod/github.com/shopspring/decimal@v1.3.1/decimal.go:565 +0x244
// github.com/shopspring/decimal.Decimal.DivRound({0x14000456b60?, 0xfffffff0?}, {0x14000456ba0?, 0x0?}, 0x10)
//         /Users/dimast/go/pkg/mod/github.com/shopspring/decimal@v1.3.1/decimal.go:607 +0x3c
// github.com/shopspring/decimal.Decimal.Div(...)
