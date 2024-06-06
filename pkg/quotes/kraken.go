package quotes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerKraken = log.Logger("kraken")

// kraken implements driver for Kraken WebSocket API v1.9.2.
// See docs here: https://docs.kraken.com/websockets
type kraken struct {
	once        *once
	conn        wsTransport
	dialer      wsDialer
	url         string
	retryPeriod time.Duration

	availablePairs safe.Map[string, krakenPair]
	streams        safe.Map[Market, struct{}]
	filter         Filter
	history        HistoricalData
	outbox         chan<- TradeEvent
}

func newKraken(config KrakenConfig, outbox chan<- TradeEvent, history HistoricalData) (Driver, error) {
	limiter := &wsDialWrapper{}

	// Set rate limit to 1 req/sec
	// as imposed by Kraken API docs here:
	// https://support.kraken.com/hc/en-us/articles/206548367-What-are-the-API-rate-limits-
	limiter.WithRateLimit(1)

	if !(strings.HasPrefix(config.URL, "ws://") || strings.HasPrefix(config.URL, "wss://")) {
		return nil, fmt.Errorf("%s (got '%s')", ErrInvalidWsUrl, config.URL)
	}

	return &kraken{
		once:           newOnce(),
		url:            config.URL,
		dialer:         limiter,
		retryPeriod:    config.ReconnectPeriod,
		availablePairs: safe.NewMap[string, krakenPair](),
		streams:        safe.NewMap[Market, struct{}](),

		filter:  NewFilter(config.Filter),
		history: history,
		outbox:  outbox,
	}, nil
}

func (k *kraken) ActiveDrivers() []DriverType {
	return []DriverType{DriverKraken}
}

func (b *kraken) ExchangeType() ExchangeType {
	return ExchangeTypeCEX
}

func (k *kraken) Start() error {
	var startErr error
	started := k.once.Start(func() {
		if err := k.getPairs(); err != nil {
			startErr = err
			return
		}

		if err := k.connect(); err != nil {
			startErr = err
			return
		}

		go k.listen()
	})

	if !started {
		return ErrAlreadyStarted
	}
	return startErr
}

func (k *kraken) Stop() error {
	var stopErr error
	stopped := k.once.Stop(func() {
		conn := k.conn
		k.conn = nil

		if conn == nil {
			return // connection is already closed
		}

		k.availablePairs = safe.Map[string, krakenPair]{}
		k.streams = safe.Map[Market, struct{}]{} // delete all stopped streams
		stopErr = conn.Close()
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return stopErr
}

type krakenSubscriptionMessage struct {
	Event string `json:"event"` // "subscribe" | "unsubscribe"
	// Pair is a list of currency pairs.
	// Format of each pair is "A/B", where A and B are ISO 4217-A3
	// for standardized assets and popular unique symbol if not standardized.
	Pair         []string                 `json:"pair"`
	Subscription krakenSubscriptionParams `json:"subscription"`
}

type krakenSubscriptionParams struct {
	// Name field sets the subscription target.
	// Available variants: book|ohlc|openOrders|ownTrades|spread|ticker|trade|*
	// * for all available channels depending on the connected environment
	Name string `json:"name"`
}

func (k *kraken) Subscribe(market Market) error {
	if !k.once.Subscribe() {
		return ErrNotStarted
	}

	if _, ok := k.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	if err := k.subscribeUnchecked(market); err != nil {
		return err
	}

	k.streams.Store(market, struct{}{})
	return nil
}

func (k *kraken) subscribeUnchecked(market Market) error {
	pair := fmt.Sprintf("%s/%s", strings.ToUpper(market.Base()), strings.ToUpper(market.Quote()))
	if _, ok := k.availablePairs.Load(pair); !ok {
		return fmt.Errorf("market %s doesn't exist in Kraken", pair)
	}

	subMsg := krakenSubscriptionMessage{
		Event:        "subscribe",
		Pair:         []string{pair},
		Subscription: krakenSubscriptionParams{Name: "trade"},
	}

	if err := k.writeConn(subMsg); err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedSub, err)
	}

	return nil
}

func (k *kraken) Unsubscribe(market Market) error {
	if !k.once.Unsubscribe() {
		return ErrNotStarted
	}

	if _, ok := k.streams.Load(market); !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	if err := k.unsubscribeUnchecked(market); err != nil {
		return err
	}

	k.streams.Delete(market)
	return nil
}

func (*kraken) HistoricalData(_ context.Context, _ Market, _ time.Duration, _ uint64) ([]TradeEvent, error) {
	return nil, errors.New("not implemented")
}

func (k *kraken) unsubscribeUnchecked(market Market) error {
	pair := fmt.Sprintf("%s/%s", strings.ToUpper(market.Base()), strings.ToUpper(market.Quote()))
	if _, ok := k.availablePairs.Load(pair); !ok {
		return fmt.Errorf("market %s doesn't exist in Kraken", pair)
	}

	unsubMsg := krakenSubscriptionMessage{
		Event:        "unsubscribe",
		Pair:         []string{pair},
		Subscription: krakenSubscriptionParams{Name: "trade"},
	}

	if err := k.writeConn(unsubMsg); err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedUnsub, err)
	}

	return nil
}

func (k *kraken) writeConn(msg krakenSubscriptionMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshalling subscription message: %v", err)
	}

	for !k.conn.IsConnected() {
	}

	if err := k.conn.WriteMessage(websocket.TextMessage, payload); err != nil {
		return fmt.Errorf("error writing subscription message: %v", err)
	}
	return nil
}

func (k *kraken) connect() error {
	// Connect to Kraken API

	for {
		var err error
		k.conn, _, err = k.dialer.Dial(k.url, nil)
		if err != nil {
			loggerKraken.Error(err)
			<-time.After(k.retryPeriod)
			continue
		}

		break
	}

	// Read initial message

	_, msg, err := k.conn.ReadMessage()
	if err != nil {
		return err
	}

	// Check if Kraken API is online

	var initResp map[string]interface{}
	if err := json.Unmarshal(msg, &initResp); err != nil {
		return err
	}
	if !(initResp["event"] == "systemStatus" && initResp["status"] == "online") {
		return fmt.Errorf("Kraken API is offline: %v", initResp)
	}
	return nil
}

func (k *kraken) listen() {
	for {
		if k.conn == nil {
			return
		}
		if !k.conn.IsConnected() {
			<-time.After(k.retryPeriod)
			continue
		}

		_, rawMsg, err := k.conn.ReadMessage()
		if err != nil {
			loggerKraken.Errorf("error reading message: %v", err)

			for {
				if err := k.connect(); err == nil {
					break
				}
				<-time.After(5 * time.Second)
			}

			k.streams.Range(func(market Market, _ struct{}) bool {
				if err := k.subscribeUnchecked(market); err != nil {
					loggerKraken.Warnf("failed to subscribe to market %s: %s", market, err)
					// Returning false here would stop iteration over the map,
					// which results in not resubscribing to all markets.
				}
				return true
			})

			continue
		}

		loggerKraken.Infow("raw event", "event", string(rawMsg))

		tradeEvents, err := k.parseMessage(rawMsg)
		if err != nil {
			loggerKraken.Errorf("error parsing message: %v", err)
			continue
		}
		if tradeEvents == nil {
			continue
		}

		for _, tr := range tradeEvents {
			if !k.filter.Allow(tr) {
				continue
			}
			k.outbox <- tr
		}
	}
}

func (k *kraken) parseMessage(rawMsg []byte) ([]TradeEvent, error) {
	// TODO: handle unsubscribe response
	// TODO: handle error response

	var tradeData []interface{}
	var eventData map[string]interface{}
	eventErr := json.Unmarshal(rawMsg, &eventData)
	tradeErr := json.Unmarshal(rawMsg, &tradeData)

	if eventErr != nil && tradeErr != nil {
		return nil, fmt.Errorf("failed to unmarshal message: `%s`", string(rawMsg))
	}
	if eventErr == nil && tradeErr != nil {
		return nil, k.parseEvent(eventData)
	}
	// NOTE: case `updateErr == nil && tradeErr == nil` is considered impossible

	events, err := k.parseTrade(tradeData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse trade `%s`: %w", string(rawMsg), err)
	}
	return k.buildEvents(events)
}

func (k *kraken) parseEvent(eventData map[string]interface{}) error {
	if eventData["event"] == "heartbeat" {
		return nil
	}

	if eventData["event"] == "subscriptionStatus" {
		status := eventData["status"].(string)
		if status == "subscribed" {
			loggerKraken.Infow("subscribed", "pair", eventData["pair"])
			return nil
		}
		if status == "error" {
			return fmt.Errorf("subscription error: %s", eventData["errorMessage"])
		}
	}

	return nil
}

type krakenTrade struct {
	ChannelID   int                 `json:"-"`
	TradeData   []krakenTradeDetail `json:"-"`
	ChannelName string              `json:"-"`
	Pair        string              `json:"-"`
}

type krakenTradeDetail struct {
	Price     string `json:"price"`
	Volume    string `json:"volume"`
	Time      string `json:"time"`
	Side      string `json:"side"`      // "b" for buy | "s" for sell
	OrderType string `json:"orderType"` // "m" for market | "l" for limit
	Misc      string `json:"misc"`
}

func (k *kraken) parseTrade(data []interface{}) (trade krakenTrade, err error) {
	trade.ChannelID = int(data[0].(float64))
	trade.ChannelName = data[2].(string)
	trade.Pair = data[3].(string)

	// Extract trade details

	tradeDetails, ok := data[1].([]interface{})
	if !ok {
		return trade, fmt.Errorf("error in type assertion for trade details")
	}

	for _, item := range tradeDetails {
		itemDetails, ok := item.([]interface{})
		if !ok {
			return trade, fmt.Errorf("error in type assertion for an item in trade details")
		}

		var detail krakenTradeDetail
		detail.Price = itemDetails[0].(string)
		detail.Volume = itemDetails[1].(string)
		detail.Time = itemDetails[2].(string)
		detail.Side = itemDetails[3].(string)
		detail.OrderType = itemDetails[4].(string)
		detail.Misc = itemDetails[5].(string)

		trade.TradeData = append(trade.TradeData, detail)
	}

	return trade, nil
}

func (*kraken) buildEvents(trades krakenTrade) ([]TradeEvent, error) {
	events := make([]TradeEvent, 0, len(trades.TradeData))
	for _, tr := range trades.TradeData {
		price, err := decimal.NewFromString(tr.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price `%s`: %w", tr.Price, err)
		}

		amount, err := decimal.NewFromString(tr.Volume)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price `%s`: %w", tr.Price, err)
		}

		takerType := TakerTypeBuy
		if tr.Side == "s" {
			takerType = TakerTypeSell
		}

		unixTime, err := strconv.ParseFloat(tr.Time, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp `%s`: %w", tr.Time, err)
		}
		sec, dec := math.Modf(unixTime)
		createdAt := time.Unix(int64(sec), int64(dec*(1e9)))

		// According to kraken docs, trade pair should have format: BTC/USDT
		// https://docs.kraken.com/websockets/#message-trade
		market, ok := NewMarketFromString(trades.Pair)
		if !ok {
			return nil, fmt.Errorf("failed to parse trade pair: %s", trades.Pair)
		}

		events = append(events, TradeEvent{
			Source:    DriverKraken,
			Market:    market,
			Price:     price,
			Amount:    amount,
			Total:     price.Mul(amount),
			TakerType: takerType,
			CreatedAt: createdAt,
		})
	}
	return events, nil
}

type krakenAssetPairs struct {
	Error  []any                 `json:"error"`
	Result map[string]krakenPair `json:"result"`
}

type krakenPair struct {
	Altname           string      `json:"altname"`
	Wsname            string      `json:"wsname"`
	AclassBase        string      `json:"aclass_base"`
	Base              string      `json:"base"`
	AclassQuote       string      `json:"aclass_quote"`
	Quote             string      `json:"quote"`
	Lot               string      `json:"lot"`
	CostDecimals      int         `json:"cost_decimals"`
	PairDecimals      int         `json:"pair_decimals"`
	LotDecimals       int         `json:"lot_decimals"`
	LotMultiplier     int         `json:"lot_multiplier"`
	LeverageBuy       []any       `json:"leverage_buy"`
	LeverageSell      []any       `json:"leverage_sell"`
	Fees              [][]float64 `json:"fees"`
	FeesMaker         [][]float64 `json:"fees_maker"`
	FeeVolumeCurrency string      `json:"fee_volume_currency"`
	MarginCall        int         `json:"margin_call"`
	MarginStop        int         `json:"margin_stop"`
	Ordermin          string      `json:"ordermin"`
	Costmin           string      `json:"costmin"`
	TickSize          string      `json:"tick_size"`
	Status            string      `json:"status"`
}

func (k *kraken) getPairs() error {
	// Fetch pairs

	req, err := http.NewRequest(http.MethodGet, "https://api.kraken.com/0/public/AssetPairs", nil)
	if err != nil {
		return fmt.Errorf("HTTP request error: %w", err)
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request error: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			loggerKraken.Errorf("error closing HTTP response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var pairs krakenAssetPairs
	if err := json.Unmarshal(body, &pairs); err != nil {
		return fmt.Errorf("failed to unmarshal pairs response: %v", err)
	}

	// Store active pairs in memory

	for _, pair := range pairs.Result {
		if pair.Status != "online" {
			continue
		}
		k.availablePairs.Store(pair.Wsname, pair)
	}

	return nil
}
