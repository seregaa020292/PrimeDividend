package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"primedividend/api/internal/config/consts"
	"primedividend/api/internal/handlers"
	"primedividend/api/internal/infrastructure/server/middlewares"
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/auth/service/strategy"
)

type Routes struct {
	strategy strategy.Strategy
	http     handlers.HttpHandlers
	ws       handlers.WsHandlers
}

func NewRoutes(
	strategy strategy.Strategy,
	http handlers.HttpHandlers,
	ws handlers.WsHandlers,
) Routes {
	return Routes{
		strategy: strategy,
		http:     http,
		ws:       ws,
	}
}

func (r Routes) Handle() chi.Router {
	router := chi.NewRouter()
	swagger := openapi.NewSwagger()

	authValidate := middlewares.NewAuthValidate(swagger.Router, r.strategy)
	auth := middlewares.NewAuth(r.strategy)

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(httprate.LimitByIP(consts.RequestLimit, consts.WindowLength))
	router.Use(middlewares.NewStructLogger())
	router.Use(middleware.Recoverer)
	router.Use(middlewares.CorsHandler())
	router.Use(middlewares.Headers...)
	router.Use(middleware.NoCache)

	r.http.Handle(router, []openapi.MiddlewareFunc{
		authValidate.Middleware,
	})

	router.Route("/ws", func(router chi.Router) {
		router.Group(func(router chi.Router) {
			router.Use(auth.Middleware)
			router.Get("/quotes", r.ws.Market.Quotes)
		})
	})

	router.NotFound(middlewares.NotFound)
	router.MethodNotAllowed(middlewares.NotAllowed)

	return router
}
