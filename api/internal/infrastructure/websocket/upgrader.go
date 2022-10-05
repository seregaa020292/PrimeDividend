package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

const bufferSize = 1024

type Upgrader struct {
	*websocket.Upgrader
}

func NewUpgrader() Upgrader {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  bufferSize,
		WriteBufferSize: bufferSize,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return r.Method == http.MethodGet
	}

	return Upgrader{Upgrader: upgrader}
}

func (ws Upgrader) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return ws.Upgrader.Upgrade(w, r, nil)
}
