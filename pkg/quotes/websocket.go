package quotes

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
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

type wsDialWrapper struct{}

func (wsDialWrapper) Dial(url string, header http.Header) (wsTransport, *http.Response, error) {
	conn, resp, err := websocket.DefaultDialer.Dial(url, header)
	return &wsConn{conn: conn}, resp, err
}

type wsConn struct {
	conn *websocket.Conn
	mu   *sync.Mutex
}

func (w *wsConn) IsConnected() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn != nil
}

func (w *wsConn) ReadMessage() (messageType int, p []byte, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.ReadMessage()
}

func (w *wsConn) WriteMessage(messageType int, data []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.WriteMessage(messageType, data)
}

func (w *wsConn) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.Close()
}
