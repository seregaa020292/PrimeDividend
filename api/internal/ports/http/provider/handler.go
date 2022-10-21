package provider

import "primedivident/internal/infrastructure/server/response"

type HandlerProvider struct {
	responder response.Responder
}

func NewHandler(
	responder response.Responder,
) HandlerProvider {
	return HandlerProvider{
		responder: responder,
	}
}
