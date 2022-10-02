package main

import (
	"primedivident/internal/config"
	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructure/wire"
	"primedivident/pkg/datetime"
	"primedivident/pkg/graceful"
)

func main() {
	datetime.InitLocation()

	cfg := config.GetConfig()

	server := wire.Initialize(cfg)

	g := graceful.NewGraceful(consts.TimeoutShutdown)
	g.Shutdown(graceful.Operations{
		server.Stop,
	})

	server.Run()
}
