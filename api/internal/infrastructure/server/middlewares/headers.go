package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

var Headers = []func(http.Handler) http.Handler{
	middleware.SetHeader("X-Content-Type-Options", "nosniff"),
	middleware.SetHeader("X-Frame-Options", "deny"),
}
