// App for testing quotes drivers.
// Set your driver here or as a console argument: `go run . binance`
package main

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/layer-3/clearsync/pkg/quotes"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

	driverName := quotes.DriverBinance
	if len(os.Args) == 2 {
		parsedDriver, err := quotes.ToDriverType(os.Args[1])
		if err != nil {
			panic(err)
		}
		driverName = parsedDriver
	}

	config, err := quotes.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}
	config.Driver = driverName

	outbox := make(chan quotes.TradeEvent, 128)
	driver, err := quotes.NewDriver(config, outbox)
	if err != nil {
		panic(err)
	}

	if err := driver.Start(); err != nil {
		panic(err)
	}

	market := quotes.Market{BaseUnit: "usdc", QuoteUnit: "weth"}
	if err = driver.Subscribe(market); err != nil {
		panic(err)
	}

	slog.Info("Subscribed", "market", market)
	for e := range outbox {
		slog.Info("new trade",
			"market", e.Market,
			"side", e.TakerType.String(),
			"price", e.Price.String(),
			"amount", e.Amount.String())
	}
}
