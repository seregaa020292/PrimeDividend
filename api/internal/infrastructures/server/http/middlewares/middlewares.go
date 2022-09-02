package middlewares

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"primedivident/internal/config/consts"
)

func Setup(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(httprate.LimitByIP(consts.RequestLimit, consts.WindowLength))
	router.Use(NewStructuredLogger())
	router.Use(middleware.Recoverer)

	router.Use(corsHandler())

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)

	router.Use(middleware.NoCache)
}
