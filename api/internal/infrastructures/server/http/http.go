package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructures/server/http/middlewares"
)

type Handlers interface {
	Setup(router chi.Router)
}

type Server struct {
	server   *http.Server
	handlers Handlers
}

func NewServer(handlers Handlers) Server {
	return Server{
		handlers: handlers,
	}
}

func (s *Server) Run() {
	log.Println("Starting HTTP server")

	router := chi.NewRouter()

	middlewares.Setup(router)
	s.handlers.Setup(router)

	s.server = &http.Server{
		Addr:    consts.ServerAddr,
		Handler: router,
	}

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("Unable to start HTTP server: %s", err)
	}
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Stop HTTP server")

	return s.server.Shutdown(ctx)
}
