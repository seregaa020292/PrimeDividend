package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	serverHttp "primedivident/internal/infrastructures/server/http"
	"primedivident/internal/ports/http/portfolio"
)

func main() {
	serverHttp.RunHTTPServer(serverHttp.Handlers{
		func(router chi.Router) http.Handler {
			return portfolio.HandlerFromMux(portfolio.NewHttpServer(), router)
		},
	})
}
