package market

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/infrastructure/socket"
	"primedivident/internal/modules/market/service/quotes"
)

type HandlerMarket struct {
	responder response.Responder
	socket    socket.Upgrader
	quotes    *quotes.Quotes
}

func NewHandlerMarket(
	responder response.Responder,
	socket socket.Upgrader,
	quotes *quotes.Quotes,
) HandlerMarket {
	return HandlerMarket{
		responder: responder,
		socket:    socket,
		quotes:    quotes,
	}
}
