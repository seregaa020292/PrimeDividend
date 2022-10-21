package currency

import "primedivident/internal/infrastructure/server/response"

type HandlerCurrency struct {
	responder response.Responder
}

func NewHandler(
	responder response.Responder,
) HandlerCurrency {
	return HandlerCurrency{
		responder: responder,
	}
}
