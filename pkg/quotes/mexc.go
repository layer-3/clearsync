package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerMexc = log.Logger("mexc")

type mexc struct {
	once               *once
	usdcToUSDT         bool
	assetsUpdatePeriod time.Duration
	idlePeriod         time.Duration
	exchangeInfo       *mexcExchangeInfoService
	filter             Filter
	history            HistoricalData
	batcherInbox       chan<- TradeEvent
	outbox             chan<- TradeEvent
	streams            safe.Map[Market, chan struct{}]
	symbolToMarket     safe.Map[string, Market]
	assets             safe.Map[Market, mexcSymbol]
	requestID          atomic.Int64
}

type mexcSymbol struct {
	Symbol     string `json:"symbol"`
	Status     string `json:"status"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

// Define the mexcDeal struct
type mexcDeal struct {
	Price    string `json:"p"`
	Quantity string `json:"v"`
	Side     int    `json:"S"`
	Time     int64  `json:"t"`
}

// Define the mexcTradeMessage struct
type mexcTradeMessage struct {
	C      string    `json:"c"`
	D      mexcDeals `json:"d"`
	Symbol string    `json:"s"`
	Time   int64     `json:"t"`
}

type mexcDeals struct {
	Deals []mexcDeal `json:"deals"`
	Event string     `json:"e"`
}

type mexcExchangeInfoService struct {
	client *mexcClient
}

func (s *mexcExchangeInfoService) Do(ctx context.Context) (*mexcExchangeInfo, error) {
	url := "https://api.mexc.com/api/v3/exchangeInfo"
	resp, err := s.client.get(ctx, url)
	if err != nil {
		return nil, err
	}
	var exchangeInfo mexcExchangeInfo
	if err := json.Unmarshal(resp, &exchangeInfo); err != nil {
		return nil, err
	}
	return &exchangeInfo, nil
}

type mexcExchangeInfo struct {
	Symbols []mexcSymbol `json:"symbols"`
}

type mexcClient struct {
	baseURL    string
	httpClient *http.Client
}

func newMexcClient(baseURL string) *mexcClient {
	return &mexcClient{baseURL: baseURL, httpClient: &http.Client{}}
}

func (c *mexcClient) get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func newMexc(config MexcConfig, outbox chan<- TradeEvent, history HistoricalData) Driver {
	batcherInbox := make(chan TradeEvent, 1024)
	go batchMexc(config.BatchPeriod, batcherInbox, outbox)

	driver := &mexc{
		once:               newOnce(),
		usdcToUSDT:         config.USDCtoUSDT,
		assetsUpdatePeriod: config.AssetsUpdatePeriod,
		exchangeInfo:       &mexcExchangeInfoService{client: newMexcClient("https://api.mexc.com")},
		filter:             NewFilter(config.Filter),
		history:            history,
		batcherInbox:       batcherInbox,
		outbox:             outbox,
		streams:            safe.NewMap[Market, chan struct{}](),
		symbolToMarket:     safe.NewMap[string, Market](),
		assets:             safe.NewMap[Market, mexcSymbol](),
		requestID:          atomic.Int64{},
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

	return driver
}

func (b *mexc) ActiveDrivers() []DriverType {
	return []DriverType{DriverMexc}
}

func (b *mexc) ExchangeType() ExchangeType {
	return ExchangeTypeCEX
}

func (b *mexc) Start() error {
	if started := b.once.Start(func() {}); !started {
		return ErrAlreadyStarted
	}
	return nil
}

func (b *mexc) Stop() error {
	stopped := b.once.Stop(func() {
		b.streams.Range(func(market Market, _ chan struct{}) bool {
			err := b.Unsubscribe(market)
			return err == nil
		})

		b.streams = safe.NewMap[Market, chan struct{}]()
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return nil
}

func (b *mexc) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return ErrNotStarted
	}

	if b.usdcToUSDT && market.Quote() == "usd" {
		if err := b.Subscribe(NewMarket(market.Base(), "usdt")); err != nil {
			loggerMexc.Warnw("failed to subscribe to USDT", "market", market, "error", err)
		}

		if err := b.Subscribe(NewMarket(market.Base(), "usdc")); err != nil {
			loggerMexc.Warnw("failed to subscribe to USDC", "market", market, "error", err)
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

	stopCh := make(chan struct{})
	b.streams.Store(market, stopCh)

	go b.watchTrades(symbol, stopCh)

	recordSubscribed(DriverMexc, market)
	loggerMexc.Infow("subscribed", "market", market)
	return nil
}

func (b *mexc) watchTrades(symbol string, stopCh chan struct{}) {
	idle := time.NewTimer(b.idlePeriod)
	doneCh := make(chan struct{})
	defer idle.Stop()
	defer close(doneCh)

	url := "wss://wbs.mexc.com/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		loggerMexc.Errorw("failed to connect to websocket", "error", err)
		return
	}
	defer conn.Close()

	subMsg := map[string]interface{}{
		"id":     b.requestID.Load(),
		"method": "SUBSCRIPTION",
		"params": []string{"spot@public.deals.v3.api@" + strings.ToUpper(symbol)},
	}
	b.requestID.Add(1)
	if err := conn.WriteJSON(subMsg); err != nil {
		loggerMexc.Errorw("failed to subscribe", "error", err)
		return
	}

	for {
		select {
		case <-stopCh:
			unsubMsg := map[string]interface{}{
				"id":     b.requestID.Load(),
				"method": "UNSUBSCRIPTION",
				"params": []string{"spot@public.deals.v3.api@" + strings.ToUpper(symbol)},
			}
			b.requestID.Add(1)
			conn.WriteJSON(unsubMsg)
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				loggerMexc.Errorw("read error", "error", err)
				// Reconnect logic
				const maxRetries = 5
				for i := 0; i < maxRetries; i++ {
					loggerMexc.Infow("attempting to reconnect", "attempt", i+1)
					conn, _, err = websocket.DefaultDialer.Dial("wss://wbs.mexc.com/ws", nil)
					if err == nil {
						subMsg := map[string]interface{}{
							"id":     b.requestID.Load(),
							"method": "SUBSCRIPTION",
							"params": []string{"spot@public.deals.v3.api@" + strings.ToUpper(symbol)},
						}
						b.requestID.Add(1)
						if err := conn.WriteJSON(subMsg); err != nil {
							loggerMexc.Errorw("failed to resubscribe", "error", err)
						} else {
							loggerMexc.Infow("resubscribed successfully", "symbol", symbol)
							break
						}
					}
					time.Sleep(time.Second * time.Duration(1<<i)) // Exponential backoff
				}
				if err != nil {
					loggerMexc.Errorw("failed to reconnect after max retries", "error", err)
					return
				}
			} else {
				// {"c":"spot@public.deals.v3.api@ETHUSDT","d":{"deals":[{"p":"3709.00","v":"0.00172","S":1,"t":1716379423968}],"e":"spot@public.deals.v3.api"},"s":"ETHUSDT","t":1716379423970}
				// Unmarshal the JSON message into the struct
				var tradeMsg mexcTradeMessage
				err = json.Unmarshal([]byte(message), &tradeMsg)
				if err != nil {
					fmt.Println("Error unmarshalling JSON:", err)
					return
				}

				b.handleTrade(idle, tradeMsg)
			}
		}
	}
}

func (b *mexc) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return ErrNotStarted
	}

	stopCh, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	stopCh <- struct{}{}
	close(stopCh)

	b.streams.Delete(market)
	recordUnsubscribed(DriverMexc, market)
	return nil
}

type mexcAggregatedTrades struct {
	AggregateTradeID any    `json:"a"`
	FirstTradeID     any    `json:"f"`
	LastTradeID      any    `json:"l"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	Timestamp        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"` // Was the buyer the maker?
	IsBestPriceMatch bool   `json:"M"` // Was the trade the best price match?
}

type mexcAggregatedTradesError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (b *mexc) HistoricalData(ctx context.Context, market Market, window time.Duration, limit uint64) ([]TradeEvent, error) {
	trades, err := fetchHistoryDataFromExternalSource(ctx, b.history, market, window, limit, loggerMexc)
	if err == nil && len(trades) > 0 {
		return trades, nil
	}

	// Build request
	const baseURL = "https://api.mexc.com/api/v3/aggTrades"

	symbol := strings.ToUpper(market.Base() + market.Quote())
	if strings.ToLower(market.Quote()) == "usd" {
		symbol = symbol + "T" // USD -> USDT
	}
	endTime := time.Now().UnixMilli()
	startTime := time.Now().Add(-window).UnixMilli()

	url := fmt.Sprintf("%s?symbol=%s&startTime=%d&endTime=%d&limit=%d",
		baseURL, symbol, startTime, endTime, limit)

	// Fetch historical data
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch historical data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse historical data
	var aggTrades []mexcAggregatedTrades
	if err := json.Unmarshal(body, &aggTrades); err != nil {
		var aggTradesErr mexcAggregatedTradesError
		if err := json.Unmarshal(body, &aggTradesErr); err == nil {
			return nil, fmt.Errorf("failed to fetch historical data: %s (code %d)", aggTradesErr.Msg, aggTradesErr.Code)
		}
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Convert aggregated trades to trade events
	trades = make([]TradeEvent, 0, len(aggTrades))
	for _, trade := range aggTrades {
		price, err := decimal.NewFromString(trade.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price: %+v", trade.Price)
		}

		amount, err := decimal.NewFromString(trade.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to parse quantity: %+v", trade.Quantity)
		}

		takerType := TakerTypeBuy
		if trade.IsBuyerMaker {
			takerType = TakerTypeSell
		}

		trades = append(trades, TradeEvent{
			Source:    DriverMexc,
			Market:    market,
			Price:     price,
			Amount:    amount,
			Total:     price.Mul(amount),
			TakerType: takerType,
			CreatedAt: time.UnixMilli(trade.Timestamp),
		})
	}

	sortTradeEventsInPlace(trades)
	return trades, nil
}

func (b *mexc) updateAssets() {
	var exchangeInfo *mexcExchangeInfo
	var err error
	for {
		exchangeInfo, err = b.exchangeInfo.Do(context.Background())
		if err == nil {
			break
		}
		loggerMexc.Errorw("failed to fetch exchange info", "error", err)
		<-time.After(5 * time.Second)
		continue
	}

	for _, symbol := range exchangeInfo.Symbols {
		if symbol.Status != "ENABLED" { // only interested in active pairs
			continue
		}

		market := NewMarket(symbol.BaseAsset, symbol.QuoteAsset)
		b.assets.Store(market, symbol)
	}
}

func (b *mexc) handleTrade(idle *time.Timer, trade mexcTradeMessage) {
	idle.Reset(b.idlePeriod)

	for _, deal := range trade.D.Deals {
		tradeEvent, err := b.buildEvent(deal, trade.Symbol)
		if err != nil {
			loggerMexc.Errorw("failed to build trade event", "event", trade, "error", err)
			return
		}

		if !b.filter.Allow(tradeEvent) {
			return
		}
		b.batcherInbox <- tradeEvent
	}
}

func (b *mexc) buildEvent(tr mexcDeal, symbol string) (TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		return TradeEvent{}, fmt.Errorf("failed to parse price: %+v", tr.Price)
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		return TradeEvent{}, fmt.Errorf("failed to parse quantity: %+v", tr.Quantity)
	}

	market, ok := b.symbolToMarket.Load(strings.ToLower(symbol))
	if !ok {
		return TradeEvent{}, fmt.Errorf("failed to load market: %+v", symbol)
	}

	if b.usdcToUSDT && (market.quoteUnit == "usdt" || market.quoteUnit == "usdc") {
		market.quoteUnit = "usd"
	}

	takerType := TakerTypeBuy
	if tr.Side == 2 {
		takerType = TakerTypeSell
	}

	return TradeEvent{
		Source:    DriverMexc,
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.UnixMilli(tr.Time),
	}, nil
}

func batchMexc(batchPeriod time.Duration, inbox <-chan TradeEvent, outbox chan<- TradeEvent) {
	marketTrades := make(map[Market][]TradeEvent)
	timer := time.NewTimer(batchPeriod)
	defer timer.Stop()

	for {
		select {
		case trade := <-inbox:
			marketTrades[trade.Market] = append(marketTrades[trade.Market], trade)
		case <-timer.C:
			for market, trades := range marketTrades {
				if event := combineTradesMexc(trades); event != nil {
					marketTrades[market] = nil
					outbox <- *event
				}
			}
			timer.Reset(batchPeriod)
		}
	}
}

func combineTradesMexc(trades []TradeEvent) *TradeEvent {
	if len(trades) == 0 {
		return nil
	}

	totalAmount := decimal.Zero
	totalValue := decimal.Zero
	netAmount := decimal.Zero

	for _, trade := range trades {
		totalAmount = totalAmount.Add(trade.Amount)
		totalValue = totalValue.Add(trade.Amount.Mul(trade.Price))

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
	var side TakerType
	if netAmount.GreaterThanOrEqual(decimal.Zero) {
		side = TakerTypeSell
	} else {
		side = TakerTypeBuy
		netAmount = netAmount.Abs()
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
