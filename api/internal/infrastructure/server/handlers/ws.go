package handlers

import "primedivident/internal/ports/ws/market"

type WsHandlers struct {
	Market market.HandlerMarket
}

func NewWsHandlers(market market.HandlerMarket) WsHandlers {
	return WsHandlers{
		Market: market,
	}
}
