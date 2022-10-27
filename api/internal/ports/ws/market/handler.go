package market

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/infrastructure/socket"
	"primedivident/internal/modules/market/service/quotes"
	"primedivident/pkg/db/redis"
)

type HandlerMarket struct {
	responder response.Responder
	socket    socket.Upgrader
	redis     *redis.Redis
	quotes    *quotes.HubQuotes
}

func NewHandlerMarket(
	responder response.Responder,
	socket socket.Upgrader,
	redis *redis.Redis,
	quotes *quotes.HubQuotes,
) HandlerMarket {
	return HandlerMarket{
		responder: responder,
		socket:    socket,
		quotes:    quotes,
	}
}
