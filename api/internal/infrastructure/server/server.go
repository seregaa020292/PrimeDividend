package server

import (
	"context"
	"log"
	"net/http"

	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructure/server/routes"
)

type Server struct {
	server *http.Server
}

func NewServer(routes routes.Routes) Server {
	return Server{
		server: &http.Server{
			ReadHeaderTimeout: consts.ServerReadHeaderTimeout,
			Addr:              consts.ServerAddr,
			Handler:           routes.Handle(),
		},
	}
}

func (s *Server) Run() {
	log.Println("Starting HTTP server")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("unable to start HTTP server: %s", err)
	}
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Stop HTTP server")

	return s.server.Shutdown(ctx)
}
