package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	serverHttp "primedivident/internal/infrastructures/server/http"
	"primedivident/internal/ports/http/portfolio"
)

var Handlers = serverHttp.Handlers{
	func(router chi.Router) http.Handler {
		return portfolio.HandlerFromMux(portfolio.NewHttpServer(), router)
	},
}
