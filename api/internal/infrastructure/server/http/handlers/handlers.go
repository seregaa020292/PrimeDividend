package handlers

import (
	"github.com/go-chi/chi/v5"

	serverHttp "primedivident/internal/infrastructure/server/http"
	"primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/portfolio"
)

type handlers struct {
	portfolio  portfolio.ServerInterface
	instrument instrument.ServerInterface
}

func NewHandlers(
	portfolio portfolio.ServerInterface,
	instrument instrument.ServerInterface,
) serverHttp.Handlers {
	return handlers{
		portfolio:  portfolio,
		instrument: instrument,
	}
}

func (h handlers) Setup(router chi.Router) {
	portfolio.HandlerFromMux(h.portfolio, router)
	instrument.HandlerFromMux(h.instrument, router)
}
