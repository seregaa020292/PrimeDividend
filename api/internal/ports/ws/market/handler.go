package market

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/infrastructure/socket"
	"primedivident/internal/modules/market/service/quotes"
)

type HandlerMarket struct {
	responder response.Responder
	socket    socket.Upgrader
	quotes    *quotes.HubQuotes
}

func NewHandlerMarket(
	responder response.Responder,
	socket socket.Upgrader,
	quotes *quotes.HubQuotes,
) HandlerMarket {
	return HandlerMarket{
		responder: responder,
		socket:    socket,
		quotes:    quotes,
	}
}
