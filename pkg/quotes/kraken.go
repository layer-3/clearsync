package quotes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

type kraken struct {
	conn        wsTransport
	connMu      sync.Mutex
	dialer      wsDialer
	url         string
	retryPeriod time.Duration
	isConnected atomic.Bool

	streams      sync.Map
	tradeSampler tradeSampler
	outbox       chan<- TradeEvent
}

func newKraken(config Config, outbox chan<- TradeEvent) *kraken {
	url := "wss://ws.kraken.com/v2"
	if config.URL != "" {
		url = config.URL
	}

	return &kraken{
		url:          url,
		dialer:       wsDialWrapper{},
		retryPeriod:  config.ReconnectPeriod,
		tradeSampler: *newTradeSampler(config.TradeSampler),
		outbox:       outbox,
	}
}

func (k *kraken) Start(markets []Market) error {
	if len(markets) == 0 {
		return errors.New("no markets specified")
	}

	if err := k.connect(); err != nil {
		return err
	}

	if err := k.subscribe(markets); err != nil {
		return err
	}

	k.listen()
	return nil
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

func (k *kraken) Subscribe(market Market) error {
	if _, ok := k.streams.Load(market); ok {
		return fmt.Errorf("market %s already subscribed", market)
	}
	k.streams.Store(market, struct{}{})

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
		return fmt.Errorf("market %s: failed to subscribe: %v", market, err)
	}
	return nil
}

func (k *kraken) Unsubscribe(market Market) error {
	if _, ok := k.streams.Load(market); !ok {
		return fmt.Errorf("market %s not subscribed", market)
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
		return fmt.Errorf("market %s: failed to unsubscribe: %v", market, err)
	}
	k.streams.Delete(market)
	return nil
}

func (k *kraken) Stop() error {
	k.connMu.Lock()
	defer k.connMu.Unlock()

	conn := k.conn
	k.conn = nil

	if conn == nil {
		return nil
	}
	return conn.Close()
}

func (k *kraken) writeConn(msg subscribeMessage) error {
	k.connMu.Lock()
	defer k.connMu.Unlock()

	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshalling subscription message: %v", err)
	}

	for !k.isConnected.Load() {
	}

	if err := k.conn.WriteMessage(websocket.TextMessage, payload); err != nil {
		return fmt.Errorf("error writing subscription message: %v", err)
	}
	return nil
}

func (k *kraken) connect() error {
	k.connMu.Lock()
	defer k.connMu.Unlock()

	if k.conn != nil {
		return errors.New("already connected")
	}

	for {
		var err error
		k.conn, _, err = k.dialer.Dial(k.url, nil)
		if err != nil {
			logger.Error(err)
			time.Sleep(k.retryPeriod)
			continue
		}

		k.isConnected.Store(true)
		return nil
	}
}

func (k *kraken) subscribe(markets []Market) error {
	availablePairs, err := getKrakenPairs()
	if err != nil {
		return err
	}

	for _, m := range markets {
		symbol := fmt.Sprintf("%s%s", strings.ToUpper(m.BaseUnit), strings.ToUpper(m.QuoteUnit))
		if pair, ok := availablePairs[symbol]; !ok || pair.Status != "online" {
			logger.Warnf("market %s doesn't exist in Kraken", symbol)
			continue
		}

		if err := k.Subscribe(m); err != nil {
			logger.Warnf("failed to subscribe to Kraken market %s: %v", symbol, err)
			continue
		}

		logger.Infof("quotes service connected to Kraken %s market", symbol)
		<-time.After(25 * time.Millisecond) // to cope with rate limits
	}

	return nil
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
		if !k.isConnected.Load() {
			<-time.After(k.retryPeriod)
			continue
		}

		k.connMu.Lock()
		conn := k.conn
		k.connMu.Unlock()
		if conn == nil {
			continue
		}

		_, rawMsg, err := conn.ReadMessage()
		if err != nil {
			logger.Errorf("error reading Kraken message: %v", err)
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

func getKrakenPairs() (map[string]krakenPair, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.kraken.com/0/public/AssetPairs", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pairs krakenAssetPairs
	if err := json.Unmarshal(body, &pairs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Kraken pairs response: %v", err)
	}

	return pairs.Result, nil
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
