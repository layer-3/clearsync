package quotes

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
)

type wsTransport interface {
	IsConnected() bool
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

type wsDialer interface {
	Dial(url string, header http.Header) (wsTransport, *http.Response, error)
}

type wsDialWrapper struct {
	rps int
}

func (w *wsDialWrapper) WithRateLimit(rps int) {
	w.rps = rps
}

func (w *wsDialWrapper) Dial(url string, header http.Header) (wsTransport, *http.Response, error) {
	conn, resp, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, resp, err
	}

	connWrapper := &wsConn{
		conn:    conn,
		limiter: rate.NewLimiter(rate.Limit(w.rps), w.rps),
	}
	return connWrapper, resp, err
}

type wsConn struct {
	conn    *websocket.Conn
	mu      sync.RWMutex
	limiter *rate.Limiter
}

func (w *wsConn) IsConnected() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.conn != nil
}

func (w *wsConn) ReadMessage() (messageType int, p []byte, err error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.conn.ReadMessage()
}

func (w *wsConn) WriteMessage(messageType int, data []byte) error {
	if w.limiter != nil {
		// Wait until the limiter allows sending a message
		ctx := context.Background()
		if err := w.limiter.Wait(ctx); err != nil {
			return err
		}
	}

	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.conn.WriteMessage(messageType, data)
}

func (w *wsConn) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.Close()
}
