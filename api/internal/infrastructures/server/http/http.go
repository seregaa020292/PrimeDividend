package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"primedivident/internal/infrastructures/server/http/middlewares"
)

const (
	startPath = "/"
	addr      = ":3000"
)

type (
	Handler  func(router chi.Router) http.Handler
	Handlers []Handler
)

type Server struct {
	server *http.Server
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Run(createHandlers Handlers) {
	apiRouter := chi.NewRouter()
	middlewares.Setup(apiRouter)

	router := chi.NewRouter()

	for _, createHandler := range createHandlers {
		router.Mount(startPath, createHandler(apiRouter))
	}

	log.Println("Starting HTTP server")

	s.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("Unable to start HTTP server: %s", err)
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
