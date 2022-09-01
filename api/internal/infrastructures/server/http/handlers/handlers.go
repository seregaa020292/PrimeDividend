package handlers

import (
	"github.com/go-chi/chi/v5"

	serverHttp "primedivident/internal/infrastructures/server/http"
	"primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/portfolio"
)

type handlers struct {
}

func NewHandlers() serverHttp.Handlers {
	return handlers{}
}

func (h handlers) Setup(router chi.Router) {
	portfolio.HandlerFromMux(portfolio.NewHandler(), router)
	instrument.HandlerFromMux(instrument.NewHandler(), router)
}
