package market

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/infrastructure/socket"
	"primedividend/api/internal/modules/market/service/quotes"
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
