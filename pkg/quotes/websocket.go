package quotes

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type wsTransport interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

type wsDialer interface {
	Dial(url string, header http.Header) (wsTransport, *http.Response, error)
}

type wsDialWrapper struct{}

func (wsDialWrapper) Dial(url string, header http.Header) (wsTransport, *http.Response, error) {
	return websocket.DefaultDialer.Dial(url, header)
}
