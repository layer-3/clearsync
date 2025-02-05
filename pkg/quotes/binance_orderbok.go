package quotes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

// gapThreshold is the maximum number of missing updates before a full reset is
// triggered. This is used to prevent the order book from getting out of sync.
// The value is random, but should be large enough to prevent unnecessary resets
// yet small enough to prevent the order book from getting too far out of sync.
const gapThreshold = 25

var ErrInvalidSnapshot = errors.New("invalid snapshot")

// BinanceDepthSnapshot represents a snapshot of the order book
// that is obtained from the REST API, not the WebSocket.
type BinanceDepthSnapshot struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// BinanceDepthEvent represents an incremental update to
// the order book that is obtained from the WebSocket.
type BinanceDepthEvent struct {
	FirstUpdateID int64      `json:"U"`
	LastUpdateID  int64      `json:"u"`
	Bids          [][]string `json:"b"`
	Asks          [][]string `json:"a"`
}

type BinanceOrderBook struct {
	mu           sync.Mutex
	LastUpdateID int64
	Bids         map[string]string
	Asks         map[string]string
	Updates      chan<- BinanceOrderBookOutboxEvent
	TopLevels    uint
}

type BinanceOrderBookOutboxEvent struct {
	Bids []BinanceOrderBookLevel
	Asks []BinanceOrderBookLevel
}

type BinanceOrderBookLevel struct {
	Price  decimal.Decimal
	Amount decimal.Decimal
}

func NewBinanceOrderBook(ctx context.Context, market Market, topLevels uint, outbox chan<- BinanceOrderBookOutboxEvent) (*BinanceOrderBook, error) {
	if topLevels == 0 {
		return nil, errors.New("top levels must be greater than 0")
	}
	if outbox == nil {
		return nil, errors.New("outbox is required")
	}

	ob := &BinanceOrderBook{
		Bids:      make(map[string]string),
		Asks:      make(map[string]string),
		Updates:   outbox,
		TopLevels: topLevels,
	}

	go func() {
		defer close(ob.Updates)
		for {
			err := ob.connectWebSocket(ctx, market)
			select {
			case <-ctx.Done():
				return
			default:
				loggerBinance.Errorw("reconnecting after connection failure", "error", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	return ob, nil
}

func (ob *BinanceOrderBook) applySnapshot(snapshot BinanceDepthSnapshot) {
	ob.mu.Lock()
	defer ob.mu.Unlock()
	ob.LastUpdateID = snapshot.LastUpdateID
	ob.Bids = make(map[string]string)
	ob.Asks = make(map[string]string)
	for _, bid := range snapshot.Bids {
		ob.Bids[bid[0]] = bid[1]
	}
	for _, ask := range snapshot.Asks {
		ob.Asks[ask[0]] = ask[1]
	}
}

func (ob *BinanceOrderBook) applyUpdate(ctx context.Context, event BinanceDepthEvent) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	if event.LastUpdateID < ob.LastUpdateID {
		return // ignore outdated events
	}
	if event.FirstUpdateID > ob.LastUpdateID {
		loggerBinance.Warn("Gap detected, resetting order book")
		return // indicates a gap, requiring a full reset
	}

	var wg sync.WaitGroup
	processOrders := func(orders [][]string, orderMap map[string]string, isBid bool) {
		defer wg.Done()
		for _, order := range orders {
			select {
			case <-ctx.Done():
				return // stop processing on cancellation
			default:
				price, quantity := order[0], order[1]
				if quantity == "0" {
					delete(orderMap, price)
				} else {
					orderMap[price] = quantity
				}
			}
		}
	}

	wg.Add(2)
	go processOrders(event.Bids, ob.Bids, true)
	go processOrders(event.Asks, ob.Asks, false)
	wg.Wait()

	ob.LastUpdateID = event.LastUpdateID
}

func (ob *BinanceOrderBook) notifyUpdate(ctx context.Context) {
	// Parse top bids and asks concurrently
	var wg sync.WaitGroup
	var topBids, topAsks []BinanceOrderBookLevel
	var bidErr, askErr error

	wg.Add(2)
	go func() {
		defer wg.Done()
		topBids, bidErr = parseOrderBookLevel(ctx, ob.TopLevels, ob.Bids)
	}()
	go func() {
		defer wg.Done()
		topAsks, askErr = parseOrderBookLevel(ctx, ob.TopLevels, ob.Asks)
	}()
	wg.Wait()

	// Stop processing if context was canceled
	if ctx.Err() != nil {
		return
	}
	if bidErr != nil {
		loggerBinance.Errorw("Failed to parse top bids:", "error", bidErr)
		return
	}
	if askErr != nil {
		loggerBinance.Errorw("Failed to parse top asks:", "error", askErr)
		return
	}

	// Push update while respecting context
	select {
	case ob.Updates <- BinanceOrderBookOutboxEvent{Bids: topBids, Asks: topAsks}:
	case <-ctx.Done():
		// Context was canceled, so we drop the update
	}
}

func parseOrderBookLevel(ctx context.Context, topLevels uint, orderbookSide map[string]string) ([]BinanceOrderBookLevel, error) {
	levels := make([]BinanceOrderBookLevel, topLevels)
	i := 0
	for price, qty := range orderbookSide {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if uint(i) >= topLevels {
			break
		}

		var err error
		levels[i].Price, err = decimal.NewFromString(price)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price (`%s`): %s", price, err)
		}
		levels[i].Amount, err = decimal.NewFromString(qty)
		if err != nil {
			return nil, fmt.Errorf("failed to parse amount (`%s`): %s", qty, err)
		}
		i++
	}

	return levels, nil
}

func (*BinanceOrderBook) fetchSnapshot(ctx context.Context, market Market) (BinanceDepthSnapshot, error) {
	symbol := strings.ReplaceAll(strings.ToUpper(market.String()), "/", "")
	url := fmt.Sprintf("https://api.binance.com/api/v3/depth?symbol=%s&limit=5000", symbol)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return BinanceDepthSnapshot{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return BinanceDepthSnapshot{}, fmt.Errorf("failed to fetch snapshot: %w", err)
	}
	defer resp.Body.Close()

	var snapshot BinanceDepthSnapshot
	if err := json.NewDecoder(resp.Body).Decode(&snapshot); err != nil {
		return BinanceDepthSnapshot{}, fmt.Errorf("failed to decode snapshot: %w", err)
	}
	return snapshot, nil
}

func (*BinanceOrderBook) keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}

func (ob *BinanceOrderBook) connectWebSocket(ctx context.Context, market Market) error {
	// Establish a WebSocket connection
	symbol := strings.ReplaceAll(strings.ToLower(market.String()), "/", "")
	loggerBinance.Debugw("Connecting to WebSocket", "symbol", symbol)
	url := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth@100ms", symbol)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	defer conn.Close()
	ob.keepAlive(conn, time.Minute)
	loggerBinance.Debug("Connected to WebSocket")

	// Fetch & apply orderbook snapshot
	snapshot, err := ob.fetchSnapshot(ctx, market)
	if err != nil {
		return fmt.Errorf("failed to fetch snapshot: %w", err)
	}
	if snapshot.LastUpdateID == 0 {
		return ErrInvalidSnapshot
	}
	loggerBinance.Info("Snapshot fetched")

	ob.applySnapshot(snapshot)
	loggerBinance.Infow("Snapshot applied", "LastUpdateID", snapshot.LastUpdateID)

	// Process orderbook updates
	nextExpectedID := snapshot.LastUpdateID + 1
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		loggerBinance.Debug("Waiting for message")

		_, message, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("failed to read message: %w", err)
		}
		loggerBinance.Debug("Message received")

		var event BinanceDepthEvent
		if err := json.Unmarshal(message, &event); err != nil {
			return fmt.Errorf("failed to unmarshal message: %w", err)
		}
		loggerBinance.Debugw("Event parsed", "LastUpdateID", event.LastUpdateID)

		// Ignore already-applied events
		if event.LastUpdateID < nextExpectedID {
			loggerBinance.Debugw("Ignoring outdated event",
				"EventLastUpdateID", event.LastUpdateID,
				"ExpectedLastUpdateID", nextExpectedID)
			continue
		}

		// Detect and handle gaps
		if event.FirstUpdateID > nextExpectedID {
			loggerBinance.Warnw("Gap detected, checking if a reset is needed",
				"ExpectedFirstUpdateID", nextExpectedID,
				"ReceivedFirstUpdateID", event.FirstUpdateID)

			// If gap is too large, fetch a new snapshot
			if event.FirstUpdateID-nextExpectedID > gapThreshold {
				loggerBinance.Warnw("Large gap detected, refetching snapshot",
					"GapSize", event.FirstUpdateID-nextExpectedID,
					"Threshold", gapThreshold)
				snapshot, err := ob.fetchSnapshot(ctx, market)
				if err != nil {
					return fmt.Errorf("failed to fetch snapshot: %w", err)
				}
				ob.applySnapshot(snapshot)
				nextExpectedID = snapshot.LastUpdateID + 1
				continue
			}

			continue // otherwise, just buffer missing events
		}

		// Apply update immediately
		loggerBinance.Debug("Applying update")
		ob.applyUpdate(ctx, event)
		loggerBinance.Infow("Order book updated", "LastUpdateID", event.LastUpdateID)
		ob.notifyUpdate(ctx)
		loggerBinance.Debug("Update notified")

		// Update next expected ID
		nextExpectedID = event.LastUpdateID + 1
	}
}
