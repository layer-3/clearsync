package driver

import (
	"context"
	"fmt"
	"strings"
	"time"

	gobinance "github.com/adshao/go-binance/v2"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
	"github.com/layer-3/clearsync/pkg/quotes/filter"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerBinance = log.Logger("binance")

type binance struct {
	once               *common.Once
	usdcToUSDT         bool
	assetsUpdatePeriod time.Duration
	idlePeriod         time.Duration
	binanceClient      *gobinance.Client
	filter             filter.Filter
	history            common.HistoricalDataDriver
	batcherInbox       chan<- common.TradeEvent
	outbox             chan<- common.TradeEvent
	streams            safe.Map[common.Market, chan struct{}]
	symbolToMarket     safe.Map[string, common.Market]
	assets             safe.Map[common.Market, gobinance.Symbol]
}

func newBinance(config BinanceConfig, outbox chan<- common.TradeEvent, history common.HistoricalDataDriver) (common.Driver, error) {
	traderFilter, err := filter.New(config.Filter, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create filter: %w", err)
	}

	batcherInbox := make(chan common.TradeEvent, 1024)
	go batch(config.BatchPeriod, batcherInbox, outbox)

	driver := &binance{
		once:               common.NewOnce(),
		usdcToUSDT:         config.USDCtoUSDT,
		assetsUpdatePeriod: config.AssetsUpdatePeriod,
		idlePeriod:         config.IdlePeriod,
		binanceClient:      gobinance.NewClient("", ""),
		filter:             traderFilter,
		history:            history,
		batcherInbox:       batcherInbox,
		outbox:             outbox,
		streams:            safe.NewMap[common.Market, chan struct{}](),
		symbolToMarket:     safe.NewMap[string, common.Market](),
		assets:             safe.NewMap[common.Market, gobinance.Symbol](),
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

func (b *binance) Type() (common.DriverType, common.ExchangeType) {
	return common.DriverBinance, common.ExchangeTypeCEX
}

func (b *binance) Start() error {
	return b.once.Start(func() error {
		base.CexConfigured.CompareAndSwap(false, true)
		return nil
	})
}

func (b *binance) Stop() error {
	return b.once.Stop(func() error {
		b.streams.Range(func(market common.Market, _ chan struct{}) bool {
			err := b.Unsubscribe(market)
			return err == nil
		})

		b.streams = safe.NewMap[common.Market, chan struct{}]()
		base.CexConfigured.CompareAndSwap(true, false)
		return nil
	})
}

func (b *binance) Subscribe(market common.Market) error {
	if !b.once.IsStarted() {
		return common.ErrNotStarted
	}

	if b.usdcToUSDT && market.Quote() == "usd" {
		if err := b.Subscribe(common.NewMarket(market.Base(), "usdt")); err != nil {
			loggerBinance.Warnw("failed to subscribe to USDT", "market", market, "error", err)
		}

		if err := b.Subscribe(common.NewMarket(market.Base(), "usdc")); err != nil {
			loggerBinance.Warnw("failed to subscribe to USDC", "market", market, "error", err)
		}
		return nil
	}

	symbol := strings.ToLower(market.Base() + market.Quote())
	b.symbolToMarket.Store(symbol, market)

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
	}

	if _, ok := b.assets.Load(market); !ok {
		return fmt.Errorf("market does not exist: %s", market)
	}

	idle := time.NewTimer(b.idlePeriod)
	doneCh, stopCh, err := gobinance.WsTradeServe(symbol, b.handleTrade(market, idle), b.handleErr(market))
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, common.ErrFailedSub, err)
	}
	b.streams.Store(market, stopCh)

	go func() {
		defer idle.Stop()

		select {
		case <-doneCh:
			loggerBinance.Warnw("market stopped", "market", market)
		case <-idle.C:
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

	base.RecordSubscribed(common.DriverBinance, market)
	loggerBinance.Infow("subscribed", "market", market)
	return nil
}

func (b *binance) Unsubscribe(market common.Market) error {
	if !b.once.IsStarted() {
		return common.ErrNotStarted
	}

	stopCh, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
	}

	stopCh <- struct{}{}
	close(stopCh)

	b.streams.Delete(market)
	base.RecordUnsubscribed(common.DriverBinance, market)
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

		market := common.NewMarket(symbol.BaseAsset, symbol.QuoteAsset)
		b.assets.Store(market, symbol)
	}
}

func (b *binance) handleTrade(
	market common.Market,
	idle *time.Timer,
) func(*gobinance.WsTradeEvent) {
	return func(event *gobinance.WsTradeEvent) {
		idle.Reset(b.idlePeriod)

		tradeEvent, err := b.buildEvent(event, market)
		if err != nil {
			loggerBinance.Errorw("failed to build trade event", "event", event, "error", err)
			return
		}

		if !b.filter.Allow(tradeEvent) {
			return
		}
		b.batcherInbox <- tradeEvent
	}
}

func (b *binance) handleErr(market common.Market) func(error) {
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

func (b *binance) buildEvent(tr *gobinance.WsTradeEvent, market common.Market) (common.TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		return common.TradeEvent{}, fmt.Errorf("failed to parse price: %+v", tr.Price)
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		return common.TradeEvent{}, fmt.Errorf("failed to parse quantity: %+v", tr.Quantity)
	}

	if b.usdcToUSDT && (market.QuoteUnit == "usdt" || market.QuoteUnit == "usdc") {
		market.QuoteUnit = "usd"
	}

	// IsBuyerMaker: true => the trade was initiated by the sell-side; the buy-side was the order book already.
	// IsBuyerMaker: false => the trade was initiated by the buy-side; the sell-side was the order book already.
	takerType := common.TakerTypeBuy
	if tr.IsBuyerMaker {
		takerType = common.TakerTypeSell
	}

	return common.TradeEvent{
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.UnixMilli(tr.TradeTime),
	}, nil
}

func (b *binance) HistoricalData(ctx context.Context, market common.Market, window time.Duration, limit uint64) ([]common.TradeEvent, error) {
	trades, err := base.FetchHistoryDataFromExternalSource(ctx, b.history, market, window, limit, loggerBinance)
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
	trades = make([]common.TradeEvent, 0, limit)
	for _, aggTrade := range aggTrades {
		trade, err := b.buildEvent(&gobinance.WsTradeEvent{
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

	common.SortTradeEventsInPlace(trades)
	return trades, nil
}

func batch(batchPeriod time.Duration, inbox <-chan common.TradeEvent, outbox chan<- common.TradeEvent) {
	marketTrades := make(map[common.Market][]common.TradeEvent)
	timer := time.NewTimer(batchPeriod)
	defer timer.Stop()

	for {
		select {
		case trade := <-inbox:
			marketTrades[trade.Market] = append(marketTrades[trade.Market], trade)
		case <-timer.C:
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

func combineTrades(trades []common.TradeEvent) *common.TradeEvent {
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
		if trade.TakerType == common.TakerTypeBuy {
			netAmount = netAmount.Add(trade.Amount)
		} else if trade.TakerType == common.TakerTypeSell {
			netAmount = netAmount.Sub(trade.Amount)
		}
	}

	if totalAmount.IsZero() {
		return nil
	}

	avgPrice := totalValue.Div(totalAmount)
	// Determine net side (buy or sell)
	var side common.TakerType
	if netAmount.GreaterThanOrEqual(decimal.Zero) {
		side = common.TakerTypeSell // "buy" (yes, it looks inverted)
	} else {
		side = common.TakerTypeBuy // "sell"
	}

	return &common.TradeEvent{
		Market:    trades[0].Market,
		Price:     avgPrice,
		Amount:    totalAmount,
		Total:     avgPrice.Mul(totalAmount),
		TakerType: side,
		CreatedAt: time.Now(),
	}
}
