package main

import (
	"fmt"

	"github.com/layer-3/clearsync/pkg/quotes"
)

// Usage example of quotes package
func main() {
	driverConfigs := []quotes.Config{
		{
			Driver: quotes.DriverBinance,
		},
		// {
		// 	Driver: quotes.DriverUniswapV3,
		// },
	}
	outbox := make(chan quotes.TradeEvent, 128)

	indexAggregator, err := quotes.NewIndexAggregator(driverConfigs, quotes.DefaultWeightsMap, outbox)
	if err != nil {
		panic(err)
	}

	err = indexAggregator.Start([]quotes.Market{{BaseUnit: "btc", QuoteUnit: "usdt"}})
	if err != nil {
		panic(err)
	}

	err = indexAggregator.Subscribe(quotes.Market{BaseUnit: "btc", QuoteUnit: "usdt"})
	if err != nil {
		panic(err)
	}

	for e := range outbox {
		fmt.Printf("market: %s, price: %.5f, amount: %s\n", e.Market, e.Price.InexactFloat64(), e.Amount.String())
	}
}
