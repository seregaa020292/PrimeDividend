package main

import (
	"primedividend/api/internal/config"
	"primedividend/api/internal/infrastructure/wire"
	"primedividend/api/pkg/graceful"
)

func main() {
	cfg := config.GetConfig()
	ctx := graceful.SignContext()

	server := wire.Initialize(ctx, cfg)
	server.Run()
}
