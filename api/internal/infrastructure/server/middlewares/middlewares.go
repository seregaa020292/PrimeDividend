package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"primedivident/internal/config/consts"
)

type middlewareFunc = func(next http.Handler) http.Handler

type middlewares struct {
}

func NewMiddlewares() Handlers {
	return middlewares{}
}

func Setup(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(httprate.LimitByIP(consts.RequestLimit, consts.WindowLength))
	router.Use(newStructuredLogger())
	router.Use(middleware.Recoverer)

	router.Use(corsHandler())

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)

	router.Use(middleware.NoCache)
}
