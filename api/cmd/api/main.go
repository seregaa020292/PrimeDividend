package main

import (
	"primedivident/internal/config"
	"primedivident/internal/infrastructure/wire"
	"primedivident/pkg/graceful"
)

func main() {
	cfg := config.GetConfig()
	ctx := graceful.SignContext()

	server := wire.Initialize(ctx, cfg)
	server.Run()
}
