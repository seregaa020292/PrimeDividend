//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"primedividend/quotes/internal/app"
	"primedividend/quotes/internal/config"
	"primedividend/quotes/pkg/db/redis"
)

func Initialize(ctx context.Context, cfg config.Config) app.App {
	wire.Build(
		redis.NewRedis,

		app.NewApp,
	)

	return app.App{}
}
