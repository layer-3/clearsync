// App for testing quotes drivers.
// Set your driver here or as a console argument: `go run . binance`
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/layer-3/clearsync/pkg/quotes"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

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

	err = driver.Start()
	if err != nil {
		panic(err)
	}

	market := quotes.Market{BaseUnit: "usdc", QuoteUnit: "weth"}
	err = driver.Subscribe(market)
	if err != nil {
		panic(err)
	}

	log.Printf("Subscribed for market %s", market)
	for e := range outbox {
		log.Printf("market: %s, price: %.5f, amount: %s\n", e.Market, e.Price.InexactFloat64(), e.Amount.String())
	}
}
