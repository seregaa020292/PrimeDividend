// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"primedividend/quotes/internal/app"
	"primedividend/quotes/internal/config"
	"primedividend/quotes/pkg/db/redis"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, cfg config.Config) app.App {
	redisRedis := redis.NewRedis(cfg)
	appApp := app.NewApp(ctx, redisRedis)
	return appApp
}