package main

import (
	"fmt"
	"os"

	"github.com/layer-3/clearsync/pkg/quotes"
)

type driver struct {
	driverInterface quotes.Driver
}

// App for testing quotes drivers
func main() {
	// Set your driver here or as a console argument: go run . binance
	// Available drivers: binance, kraken, opendax, bitfaker, uniswap_v3, wip:index
	driverName := quotes.DriverBinance

	if len(os.Args) == 2 {
		parsedDriver, err := quotes.ToDriverType(os.Args[1])
		if err != nil {
			panic(err)
		}
		driverName = parsedDriver
	}

	outbox := make(chan quotes.TradeEvent, 128)
	driver, err := quotes.NewDriver(quotes.Config{Driver: driverName}, outbox)
	if err != nil {
		panic(err)
	}

	err = driver.Start()
	if err != nil {
		panic(err)
	}

	err = driver.Subscribe(quotes.Market{BaseUnit: "btc", QuoteUnit: "usdt"})
	if err != nil {
		panic(err)
	}

	for e := range outbox {
		fmt.Printf("market: %s, price: %.5f, amount: %s\n", e.Market, e.Price.InexactFloat64(), e.Amount.String())
	}
}
