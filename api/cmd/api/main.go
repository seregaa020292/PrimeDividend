package main

import (
	"primedivident/internal/config"
	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructures/wire"
	"primedivident/pkg/graceful"
)

func main() {
	cfg := config.GetConfig()

	server := wire.Initialize(cfg)

	g := graceful.NewGraceful(consts.TimeoutShutdown)
	g.Shutdown(graceful.Operations{
		server.Stop,
	})

	server.Run()
}
