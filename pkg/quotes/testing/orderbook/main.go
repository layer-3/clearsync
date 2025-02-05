package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/layer-3/clearsync/pkg/quotes"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	market := quotes.NewMarket("btc", "usdt")
	const topLevels = 1 // Configurable number of top levels to track
	outbox := make(chan quotes.BinanceOrderBookOutboxEvent, 128)

	_, err := quotes.NewBinanceOrderBook(ctx, market, topLevels, outbox)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for update := range outbox {
			log.Println("Top levels updated:", update.Asks[0], update.Bids[0])
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
}
