package quotes

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	gobinance "github.com/adshao/go-binance/v2"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/safe"
)

var (
	loggerBinance = log.Logger("binance")
	cexConfigured = atomic.Bool{}
)

type binanceStream struct {
	aggTradeStop chan struct{}
	midPriceStop context.CancelFunc
}

type binance struct {
	once                 *once
	usdcToUSDT           bool
	assetsUpdatePeriod   time.Duration
	idlePeriod           time.Duration
	binanceClient        *gobinance.Client
	filter               Filter
	history              HistoricalData
	tradesBatcherInbox   chan<- TradeEvent
	midPriceBatcherInbox chan<- TradeEvent
	outbox               chan<- TradeEvent
	streams              safe.Map[Market, binanceStream]
	symbolToMarket       safe.Map[string, Market]
	assets               safe.Map[Market, gobinance.Symbol]
}

func newBinance(config BinanceConfig, outbox chan<- TradeEvent, history HistoricalData) (Driver, error) {
	tradesBatcherInbox := make(chan TradeEvent, 1024)
	midPriceBatcherInbox := make(chan TradeEvent, 1024)
	go batch(config.BatchPeriod, tradesBatcherInbox, midPriceBatcherInbox, outbox)

	driver := &binance{
		once:                 newOnce(),
		usdcToUSDT:           config.USDCtoUSDT,
		assetsUpdatePeriod:   config.AssetsUpdatePeriod,
		idlePeriod:           config.IdlePeriod,
		binanceClient:        gobinance.NewClient("", ""),
		filter:               NewFilter(config.Filter),
		history:              history,
		tradesBatcherInbox:   tradesBatcherInbox,
		midPriceBatcherInbox: midPriceBatcherInbox,
		outbox:               outbox,
		streams:              safe.NewMap[Market, binanceStream](),
		symbolToMarket:       safe.NewMap[string, Market](),
		assets:               safe.NewMap[Market, gobinance.Symbol](),
	}

	driver.updateAssets()
	go func() {
		ticker := time.NewTicker(driver.assetsUpdatePeriod)
		defer ticker.Stop()
		for {
			<-ticker.C
			driver.updateAssets()
		}
	}()

	return driver, nil
}

func (b *binance) ActiveDrivers() []DriverType {
	return []DriverType{DriverBinance}
}

func (b *binance) ExchangeType() ExchangeType {
	return ExchangeTypeCEX
}

func (b *binance) Start() error {
	started := b.once.Start(func() {
		cexConfigured.CompareAndSwap(false, true)
	})
	if !started {
		return ErrAlreadyStarted
	}
	return nil
}

func (b *binance) Stop() error {
	stopped := b.once.Stop(func() {
		b.streams.Range(func(market Market, _ binanceStream) bool {
			err := b.Unsubscribe(market)
			return err == nil
		})

		b.streams = safe.NewMap[Market, binanceStream]()
		cexConfigured.CompareAndSwap(true, false)
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return nil
}

func (b *binance) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return ErrNotStarted
	}

	if b.usdcToUSDT && market.Quote() == "usd" {
		if err := b.Subscribe(NewMarket(market.Base(), "usdt")); err != nil {
			loggerBinance.Warnw("failed to subscribe to USDT", "market", market, "error", err)
		}

		if err := b.Subscribe(NewMarket(market.Base(), "usdc")); err != nil {
			loggerBinance.Warnw("failed to subscribe to USDC", "market", market, "error", err)
		}
		return nil
	}

	symbol := strings.ToLower(market.Base() + market.Quote())
	b.symbolToMarket.Store(symbol, market)

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	if _, ok := b.assets.Load(market); !ok {
		return fmt.Errorf("market does not exist: %s", market)
	}

	// Subscribe to trades
	idleTrades := time.NewTimer(b.idlePeriod)
	aggTradeDoneCh, aggTradeStopCh, err := gobinance.WsAggTradeServe(symbol, b.handleTrade(market, idleTrades), b.handleErr(market))
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedSub, err)
	}

	// Subscribe to orderbook updates
	idleOrderbook := time.NewTimer(b.idlePeriod)
	ctx, midPriceStop := context.WithCancel(context.Background())
	orderbookUpdatesCh := make(chan BinanceOrderBookOutboxEvent, 128)
	_, err = NewBinanceOrderBook(ctx, market, 1, orderbookUpdatesCh)
	if err != nil {
		midPriceStop()
		return fmt.Errorf("%s: %w: %w", market, ErrFailedSub, err)
	}
	go func() {
		h := b.handleOrderbookUpdates(market, idleOrderbook)
		for update := range orderbookUpdatesCh {
			h(update)
		}
	}()

	// Store handles to stop the streams
	b.streams.Store(market, binanceStream{
		aggTradeStop: aggTradeStopCh,
		midPriceStop: midPriceStop,
	})

	go func() {
		defer idleTrades.Stop()
		defer idleOrderbook.Stop()

		select {
		case <-aggTradeDoneCh:
			midPriceStop() // stop the sibling stream
			loggerBinance.Warnw("market stopped", "market", market)
		case <-ctx.Done():
			aggTradeStopCh <- struct{}{} // stop the sibling stream
			loggerBinance.Warnw("market stopped", "market", market)
		case <-idleTrades.C:
			loggerBinance.Warnw("market inactivity detected", "market", market)
		}

		if _, ok := b.streams.Load(market); !ok {
			return // market was unsubscribed earlier
		}

		loggerBinance.Infow("resubscribing", "market", market)
		if err := b.Unsubscribe(market); err != nil {
			loggerBinance.Errorw("failed to unsubscribe", "market", market, "error", err)
		}

		for {
			err := b.Subscribe(market)
			if err == nil {
				loggerBinance.Infow("resubscribed", "market", market)
				break
			}
			loggerBinance.Errorw("failed to resubscribe", "market", market, "error", err)
			<-time.After(5 * time.Second)
		}
	}()

	recordSubscribed(DriverBinance, market)
	loggerBinance.Infow("subscribed", "market", market)
	return nil
}

func (b *binance) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return ErrNotStarted
	}

	stream, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	stream.aggTradeStop <- struct{}{}
	close(stream.aggTradeStop)
	stream.midPriceStop()

	b.streams.Delete(market)
	recordUnsubscribed(DriverBinance, market)
	return nil
}

func (b *binance) updateAssets() {
	exchangeInfoService := b.binanceClient.NewExchangeInfoService()

	var exchangeInfo *gobinance.ExchangeInfo
	var err error
	for {
		exchangeInfo, err = exchangeInfoService.Do(context.Background())
		if err == nil {
			break
		}
		loggerBinance.Errorw("failed to fetch exchange info", "error", err)
		<-time.After(5 * time.Second)
		continue
	}

	for _, symbol := range exchangeInfo.Symbols {
		if symbol.Status != "TRADING" { // only interested in active pairs
			continue
		}

		market := NewMarket(symbol.BaseAsset, symbol.QuoteAsset)
		b.assets.Store(market, symbol)
	}
}

func (b *binance) handleTrade(
	market Market,
	idle *time.Timer,
) func(*gobinance.WsAggTradeEvent) {
	return func(event *gobinance.WsAggTradeEvent) {
		idle.Reset(b.idlePeriod)

		tradeEvent, err := b.buildEvent(event, market)
		if err != nil {
			loggerBinance.Errorw("failed to build trade event", "event", event, "error", err)
			return
		}

		if !b.filter.Allow(tradeEvent) {
			return
		}
		b.tradesBatcherInbox <- tradeEvent
	}
}

func (b *binance) handleOrderbookUpdates(
	market Market,
	idle *time.Timer,
) func(event BinanceOrderBookOutboxEvent) {
	return func(event BinanceOrderBookOutboxEvent) {
		idle.Reset(b.idlePeriod)

		// Compute mid price
		bidPrice := event.Bids[0].Price
		askPrice := event.Asks[0].Price
		midPrice := bidPrice.Add(askPrice).Div(two)

		bidAmount := event.Bids[0].Amount
		askAmount := event.Asks[0].Amount
		midAmount := bidAmount.Add(askAmount).Div(two)

		b.outbox <- TradeEvent{
			Source:    DriverBinance,
			Market:    market,
			Price:     midPrice,
			Amount:    midAmount,
			Total:     midPrice.Mul(midAmount),
			TakerType: TakerTypeUnknown,
			CreatedAt: time.Now(),
		}
	}
}

func (b *binance) handleErr(market Market) func(error) {
	return func(err error) {
		loggerBinance.Errorw("received error", "market", market, "error", err)
		if err.Error() == "websocket: close 1001 (going away)" || err.Error() == "websocket: close 1006 (abnormal closure): unexpected EOF" {
			// Reconnect logic
			const maxRetries = 5
			b.Unsubscribe(market)
			for i := 0; i < maxRetries; i++ {
				loggerBinance.Infow("attempting to reconnect to market", "market", market, "attempt", i+1)
				if err = b.Subscribe(market); err == nil {
					loggerBinance.Infow("resubscribed successfully", "market", market)
					break
				}
				loggerBinance.Errorw("failed to resubscribe", "market", market, "error", err)
				time.Sleep(time.Second * time.Duration(1<<i)) // Exponential backoff
			}
			if err != nil {
				loggerBinance.Errorw("failed to reconnect after max retries", "market", market, "error", err)
			}
		}
	}
}

func (b *binance) buildEvent(tr *gobinance.WsAggTradeEvent, market Market) (TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		return TradeEvent{}, fmt.Errorf("failed to parse price: %+v", tr.Price)
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		return TradeEvent{}, fmt.Errorf("failed to parse quantity: %+v", tr.Quantity)
	}

	if b.usdcToUSDT && (market.quoteUnit == "usdt" || market.quoteUnit == "usdc") {
		market.quoteUnit = "usd"
	}

	// IsBuyerMaker: true => the trade was initiated by the sell-side; the buy-side was the order book already.
	// IsBuyerMaker: false => the trade was initiated by the buy-side; the sell-side was the order book already.
	takerType := TakerTypeBuy
	if tr.IsBuyerMaker {
		takerType = TakerTypeSell
	}

	return TradeEvent{
		Source:    DriverBinance,
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.UnixMilli(tr.TradeTime),
	}, nil
}

func (b *binance) HistoricalData(ctx context.Context, market Market, window time.Duration, limit uint64) ([]TradeEvent, error) {
	trades, err := fetchHistoryDataFromExternalSource(ctx, b.history, market, window, limit, loggerBinance)
	if err == nil && len(trades) > 0 {
		return trades, nil
	}

	// Fetch data
	aggTradesService := b.binanceClient.NewAggTradesService()

	aggTradesService.StartTime(time.Now().Add(-window).Unix() * 1000)
	aggTradesService.EndTime(time.Now().Unix() * 1000)
	aggTradesService.Limit(int(limit))

	base := strings.ToLower(market.Base())
	quote := strings.ToLower(market.Quote())

	tokenFixtures := map[string]string{
		"usd":  "usdt", // as of spring 2024 Binance does not provide USD spot markets
		"weth": "eth",  // as of spring 2024 Binance does not provide WETH spot markets
	}
	if newBase, ok := tokenFixtures[base]; ok {
		base = newBase
	}
	if newQuote, ok := tokenFixtures[quote]; ok {
		quote = newQuote
	}

	symbol := strings.ToUpper(base + quote)
	aggTradesService.Symbol(symbol)

	aggTrades, err := aggTradesService.Do(ctx)
	if err != nil {
		loggerBinance.Errorw("failed to fetch historical data", "market", market, "error", err)
		return nil, fmt.Errorf("failed to fetch historical data: %w", err)
	}

	// Convert aggregated trades to a trade events
	trades = make([]TradeEvent, 0, limit)
	for _, aggTrade := range aggTrades {
		trade, err := b.buildEvent(&gobinance.WsAggTradeEvent{
			Price:        aggTrade.Price,
			Quantity:     aggTrade.Quantity,
			TradeTime:    aggTrade.Timestamp,
			IsBuyerMaker: aggTrade.IsBuyerMaker,
		}, market)
		if err != nil {
			loggerBinance.Errorw("failed to build trade event", "market", market, "error", err)
			continue
		}
		if trade.Price.IsZero() {
			loggerBinance.Warnw("skipping trade with zero price",
				"market", market,
				"trade", trade,
				"aggregated_trade", aggTrade)
			continue
		}
		trades = append(trades, trade)
	}

	sortTradeEventsInPlace(trades)
	return trades, nil
}

func batch(batchPeriod time.Duration, tradesInbox, midPriceIndox <-chan TradeEvent, outbox chan<- TradeEvent) {
	marketTrades := make(map[Market][]TradeEvent)
	midprices := make(map[Market][]TradeEvent)
	timer := time.NewTimer(batchPeriod)
	defer timer.Stop()

	for {
		select {
		case trade := <-tradesInbox:
			marketTrades[trade.Market] = append(marketTrades[trade.Market], trade)
		case trade := <-midPriceIndox:
			midprices[trade.Market] = append(midprices[trade.Market], trade)
		case <-timer.C:
			// TODO: combine midprices
			for market, trades := range marketTrades {
				if event := combineTrades(trades); event != nil {
					marketTrades[market] = nil
					outbox <- *event
				}
			}
			timer.Reset(batchPeriod)
		}
	}
}

func combineTrades(trades []TradeEvent) *TradeEvent {
	if len(trades) == 0 {
		return nil
	}

	totalAmount := decimal.Zero
	totalValue := decimal.Zero
	netAmount := decimal.Zero

	for _, trade := range trades {
		totalAmount = totalAmount.Add(trade.Amount)
		totalValue = totalValue.Add(trade.Amount.Mul(trade.Price))

		// Update net amount to determine net side (buy or sell)
		if trade.TakerType == TakerTypeBuy {
			netAmount = netAmount.Add(trade.Amount)
		} else if trade.TakerType == TakerTypeSell {
			netAmount = netAmount.Sub(trade.Amount)
		}
	}

	if totalAmount.IsZero() {
		return nil
	}

	avgPrice := totalValue.Div(totalAmount)
	// Determine net side (buy or sell)
	var side TakerType
	if netAmount.GreaterThanOrEqual(decimal.Zero) {
		side = TakerTypeSell // "buy" (yes, it looks inverted)
	} else {
		side = TakerTypeBuy // "sell"
	}

	return &TradeEvent{
		Source:    trades[0].Source,
		Market:    trades[0].Market,
		Price:     avgPrice,
		Amount:    totalAmount,
		Total:     avgPrice.Mul(totalAmount),
		TakerType: side,
		CreatedAt: time.Now(),
	}
}
