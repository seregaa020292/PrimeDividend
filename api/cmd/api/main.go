package main

import (
	"primedivident/internal/config"
	"primedivident/internal/infrastructures/server/http/handler"
	"primedivident/internal/infrastructures/wire"
	"primedivident/pkg/graceful"
)

func main() {
	server := wire.InitializeServer()

	//config.GetConfig()

	g := graceful.NewGraceful(config.TimeoutShutdown)
	g.Shutdown(graceful.Operations{
		server.Stop,
	})

	server.Run(handler.Handlers)
}
