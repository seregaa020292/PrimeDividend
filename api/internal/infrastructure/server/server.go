package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"primedivident/internal/config/consts"
)

type (
	Middlewares interface {
		Setup(router chi.Router)
	}
	Handlers interface {
		Setup(router chi.Router)
	}
)

type Server struct {
	server      *http.Server
	middlewares Middlewares
	handlers    Handlers
}

func NewServer(middlewares Middlewares, handlers Handlers) Server {
	return Server{
		middlewares: middlewares,
		handlers:    handlers,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()

	s.middlewares.Setup(router)
	s.handlers.Setup(router)

	s.server = &http.Server{
		ReadHeaderTimeout: consts.ServerReadHeaderTimeout,
		Addr:              consts.ServerAddr,
		Handler:           router,
	}

	log.Println("Starting HTTP server")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("unable to start HTTP server: %s", err)
	}
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Stop HTTP server")

	return s.server.Shutdown(ctx)
}
