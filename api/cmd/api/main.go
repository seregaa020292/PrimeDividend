package main

import (
	"primedivident/internal/infrastructures/server/http/handler"
	"primedivident/pkg/graceful"
	"time"
)

func main() {
	server := InitializeServer()

	g := graceful.NewGraceful(5 * time.Second)
	g.Shutdown(graceful.Operations{
		server.Stop,
	})

	server.Run(handler.Handlers)
}
