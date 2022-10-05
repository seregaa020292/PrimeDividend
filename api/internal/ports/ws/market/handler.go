package market

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/infrastructure/websocket"
)

type HandlerMarket struct {
	responder response.Responder
	websocket websocket.Upgrader
}

func NewHandlerMarket(responder response.Responder, websocket websocket.Upgrader) HandlerMarket {
	return HandlerMarket{
		responder: responder,
		websocket: websocket,
	}
}
