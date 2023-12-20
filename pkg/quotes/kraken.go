package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

type kraken struct {
	once        *once
	conn        wsTransport
	dialer      wsDialer
	url         string
	retryPeriod time.Duration

	availablePairs sync.Map
	streams        sync.Map
	tradeSampler   tradeSampler
	outbox         chan<- TradeEvent
}

func newKraken(config Config, outbox chan<- TradeEvent) *kraken {
	url := "wss://ws.kraken.com/v2"
	if config.URL != "" {
		url = config.URL
	}

	return &kraken{
		once:         newOnce(),
		url:          url,
		dialer:       wsDialWrapper{},
		retryPeriod:  config.ReconnectPeriod,
		tradeSampler: *newTradeSampler(config.TradeSampler),
		outbox:       outbox,
	}
}

func (k *kraken) Start() error {
	var startErr error
	k.once.Start(func() {
		if err := k.getKrakenPairs(); err != nil {
			startErr = err
			return
		}

		if err := k.connect(); err != nil {
			startErr = err
			return
		}

		go k.listen()
	})
	return startErr
}

func (k *kraken) Stop() error {
	var stopErr error
	k.once.Stop(func() {
		conn := k.conn
		k.conn = nil

		if conn == nil {
			return
		}
		stopErr = conn.Close()
	})
	return stopErr
}

type subscribeMessage struct {
	Method string             `json:"method"`
	Params subscriptionParams `json:"params"`
}

type subscriptionParams struct {
	Channel  string   `json:"channel"`
	Snapshot bool     `json:"snapshot"`
	Symbol   []string `json:"symbol"`
}

func (k *kraken) Name() DriverType {
	return DriverKraken
}

func (k *kraken) Subscribe(market Market) error {
	if _, ok := k.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	symbol := fmt.Sprintf("%s%s", strings.ToUpper(market.BaseUnit), strings.ToUpper(market.QuoteUnit))
	if _, ok := k.availablePairs.Load(symbol); !ok {
		return fmt.Errorf("market %s doesn't exist in Kraken", symbol)
	}

	pair := fmt.Sprintf("%s/%s", strings.ToUpper(market.BaseUnit), strings.ToUpper(market.QuoteUnit))
	subMsg := subscribeMessage{
		Method: "subscribe",
		Params: subscriptionParams{
			Channel:  "trade",
			Snapshot: true,
			Symbol:   []string{pair},
		},
	}

	if err := k.writeConn(subMsg); err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedSub, err)
	}

	k.streams.Store(market, struct{}{})
	return nil
}

func (k *kraken) Unsubscribe(market Market) error {
	if _, ok := k.streams.Load(market); !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	pair := fmt.Sprintf("%s/%s", strings.ToUpper(market.BaseUnit), strings.ToUpper(market.QuoteUnit))
	unsubMsg := subscribeMessage{
		Method: "unsubscribe",
		Params: subscriptionParams{
			Channel:  "trade",
			Snapshot: true,
			Symbol:   []string{pair},
		},
	}

	if err := k.writeConn(unsubMsg); err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedUnsub, err)
	}
	k.streams.Delete(market)
	return nil
}

func (k *kraken) writeConn(msg subscribeMessage) error {
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
	for {
		var err error
		k.conn, _, err = k.dialer.Dial(k.url, nil)
		if err != nil {
			logger.Error(err)
			time.Sleep(k.retryPeriod)
			continue
		}

		return nil
	}
}

type krakenEvent[T krakenStatus | krakenTrade] struct {
	Channel string `json:"channel"`
	Type    string `json:"type"`
	Data    []T    `json:"data"`
}

type krakenStatus struct {
	ApiVersion   string `json:"api_version"`
	ConnectionId uint64 `json:"connection_id"`
	System       string `json:"system"`
	Version      string `json:"version"`
}

type krakenTrade struct {
	OrdType   string    `json:"ord_type"`
	Price     float64   `json:"price"`
	Qty       float64   `json:"qty"`
	Side      string    `json:"side"` // "buy" | "sell"
	Symbol    string    `json:"symbol"`
	Timestamp time.Time `json:"timestamp"`
	TradeId   int       `json:"trade_id"`
}

type krakenResult struct {
	Method string `json:"method"`
	Result struct {
		Channel  string `json:"channel"`
		Snapshot bool   `json:"snapshot"`
		Symbol   string `json:"symbol"`
	} `json:"result"`
	Success bool      `json:"success"`
	TimeIn  time.Time `json:"time_in"`
	TimeOut time.Time `json:"time_out"`
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
			logger.Errorf("error reading Kraken message: %v", err)

			k.connect()
			k.streams.Range(func(m, value any) bool {
				market := m.(Market)
				if err := k.Subscribe(market); err != nil {
					logger.Warnf("Error subscribing to market %s: %s", market, err)
					return false
				}
				return true
			})

			continue
		}

		tradeEvents, err := k.parseMessage(rawMsg)
		if err != nil {
			logger.Errorf("error parsing Kraken message: %v", err)
			continue
		}
		if tradeEvents == nil {
			continue
		}

		for _, tr := range tradeEvents {
			if !k.tradeSampler.Allow(tr) {
				continue
			}

			k.outbox <- tr
		}
	}
}

func (k *kraken) parseMessage(rawMsg []byte) ([]TradeEvent, error) {
	var ticker krakenEvent[krakenTrade]
	if err := json.Unmarshal(rawMsg, &ticker); err == nil && ticker.Channel != "heartbeat" {
		return buildKrakenEvents(ticker.Data), nil
	}

	var status krakenEvent[krakenStatus]
	if err := json.Unmarshal(rawMsg, &status); err == nil && ticker.Channel != "heartbeat" {
		// TODO: Handle KrakenEvent[KrakenStatus]
		return nil, nil
	}

	var result krakenResult
	if err := json.Unmarshal(rawMsg, &result); err != nil {
		return nil, err
	}
	if ticker.Channel == "heartbeat" {
		return nil, nil
	}

	if !result.Success {
		return nil, fmt.Errorf("failed to subscribe to Kraken market %s: ", result.Result.Symbol)
	}
	return nil, fmt.Errorf("unknown Kraken message: %s", string(rawMsg))
}

type krakenAssetPairs struct {
	Error  []interface{}         `json:"error"`
	Result map[string]krakenPair `json:"result"`
}

type krakenPair struct {
	Altname           string        `json:"altname"`
	Wsname            string        `json:"wsname"`
	AclassBase        string        `json:"aclass_base"`
	Base              string        `json:"base"`
	AclassQuote       string        `json:"aclass_quote"`
	Quote             string        `json:"quote"`
	Lot               string        `json:"lot"`
	CostDecimals      int           `json:"cost_decimals"`
	PairDecimals      int           `json:"pair_decimals"`
	LotDecimals       int           `json:"lot_decimals"`
	LotMultiplier     int           `json:"lot_multiplier"`
	LeverageBuy       []interface{} `json:"leverage_buy"`
	LeverageSell      []interface{} `json:"leverage_sell"`
	Fees              [][]float64   `json:"fees"`
	FeesMaker         [][]float64   `json:"fees_maker"`
	FeeVolumeCurrency string        `json:"fee_volume_currency"`
	MarginCall        int           `json:"margin_call"`
	MarginStop        int           `json:"margin_stop"`
	Ordermin          string        `json:"ordermin"`
	Costmin           string        `json:"costmin"`
	TickSize          string        `json:"tick_size"`
	Status            string        `json:"status"`
}

func (k *kraken) getKrakenPairs() error {
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
			logger.Errorf("error closing HTTP response body: %v", err)
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
		return fmt.Errorf("failed to unmarshal Kraken pairs response: %v", err)
	}

	// Convert pairs to map

	for _, pair := range pairs.Result {
		if pair.Status != "online" {
			symbol := fmt.Sprintf("%s%s", strings.ToUpper(pair.Base), strings.ToUpper(pair.Quote))
			logger.Warnf("market %s doesn't exist in Kraken", symbol)
			continue
		}
		k.availablePairs.Store(pair.Altname, pair)
	}

	return nil
}

func buildKrakenEvents(trades []krakenTrade) []TradeEvent {
	var events []TradeEvent
	for _, tr := range trades {
		price := decimal.NewFromFloat(tr.Price)
		amount := decimal.NewFromFloat(tr.Qty)

		takerType := TakerTypeBuy
		if tr.Side == "sell" {
			takerType = TakerTypeSell
		}

		events = append(events, TradeEvent{
			Source:    DriverKraken,
			Market:    strings.ToLower(tr.Symbol),
			Price:     price,
			Amount:    amount,
			Total:     price.Mul(amount),
			TakerType: takerType,
			CreatedAt: tr.Timestamp,
		})
	}

	return events
}
