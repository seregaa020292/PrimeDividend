package app

import (
	"context"
	"log"
	"time"

	"primedividend/quotes/pkg/db/redis"
)

type App struct {
	ctx context.Context
}

func NewApp(ctx context.Context, redis *redis.Redis) App {
	return App{
		ctx: ctx,
	}
}

func (a App) Run() {
	log.Println("Run")
	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			log.Println("Tick")
		case <-a.ctx.Done():
			log.Println("Shutdown")
			return
		}
	}
}
