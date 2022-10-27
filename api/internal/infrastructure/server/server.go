package server

import (
	"context"
	"log"
	"net/http"

	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructure/server/routes"
	"primedivident/pkg/graceful"
	"primedivident/pkg/utils/errlog"
)

type Server struct {
	ctx     context.Context
	downApp graceful.ShutdownApp
	server  *http.Server
}

func NewServer(ctx context.Context, downApp graceful.ShutdownApp, routes routes.Routes) Server {
	return Server{
		ctx:     ctx,
		downApp: downApp,
		server: &http.Server{
			ReadHeaderTimeout: consts.ServerReadHeaderTimeout,
			Addr:              consts.ServerAddr,
			Handler:           routes.Handle(),
		},
	}
}

func (s *Server) Run() {
	go s.Stop()

	log.Println("Starting HTTP server")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("unable to start HTTP server: %s", err)
	}
}

func (s *Server) Stop() {
	<-s.ctx.Done()

	s.downApp.Run()

	log.Println("Stop HTTP server")

	errlog.Println(s.server.Shutdown(s.ctx))
}
