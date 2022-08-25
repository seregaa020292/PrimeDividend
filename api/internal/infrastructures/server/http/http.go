package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"primedivident/internal/infrastructures/server/http/middlewares"
)

type Handler func(router chi.Router) http.Handler

type Handlers []Handler

func RunHTTPServer(createHandlers Handlers) {
	apiRouter := chi.NewRouter()
	middlewares.Setup(apiRouter)

	rootRouter := chi.NewRouter()

	for _, createHandler := range createHandlers {
		rootRouter.Mount("/", createHandler(apiRouter))
	}

	//logrus.Info("Starting HTTP server")

	err := http.ListenAndServe(":3000", rootRouter)
	if err != nil {
		//logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
