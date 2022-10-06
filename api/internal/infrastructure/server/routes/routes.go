package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructure/server/handlers"
	"primedivident/internal/infrastructure/server/middlewares"
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/auth/service/strategy"
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

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(httprate.LimitByIP(consts.RequestLimit, consts.WindowLength))
	router.Use(middlewares.NewStructLogger())
	router.Use(middleware.Recoverer)
	router.Use(middlewares.CorsHandler())
	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)

	r.http.Handle(router, []openapi.MiddlewareFunc{
		middlewares.AuthSwagger(swagger.Router, r.strategy.VerifyAccess),
	})

	router.Route("/ws", func(router chi.Router) {
		router.Group(func(router chi.Router) {
			//router.Use(middlewares.Auth)
			router.Get("/quotes", r.ws.Market.Quotes)
		})
	})

	router.NotFound(middlewares.NotFound)
	router.MethodNotAllowed(middlewares.NotAllowed)

	return router
}
