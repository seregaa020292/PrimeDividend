package main

import (
	"primedividend/quotes/internal/config"
	"primedividend/quotes/internal/wire"
	"primedividend/quotes/pkg/graceful"
)

func main() {
	cfg := config.NewConfig()
	ctx := graceful.SignContext()

	app := wire.Initialize(ctx, cfg)
	app.Run()
}
