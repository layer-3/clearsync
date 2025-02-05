package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/layer-3/clearsync/pkg/quotes"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	market := quotes.NewMarket("bnb", "btc")
	const topLevels = 2 // Configurable number of top levels to track
	outbox := make(chan quotes.BinanceOrderBookOutboxEvent, 128)

	_, err := quotes.NewBinanceOrderBook(ctx, market, topLevels, outbox)
	if err != nil {
		log.Fatal(err)
	}

	for update := range outbox {
		log.Println("Top levels updated:", update)
	}
}
