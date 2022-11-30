package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"

	"primedividend/quotes/internal/config"
	"primedividend/quotes/pkg/utils/errlog"
)

const (
	MinIdleConns = 200
	PoolSize     = 12000
	PoolTimeout  = 240 * time.Second
)

type Redis struct {
	*redis.Client
}

func NewRedis(config config.Config) *Redis {
	log.Println("Start Redis")

	connect := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		MinIdleConns: MinIdleConns,
		PoolSize:     PoolSize,
		PoolTimeout:  PoolTimeout,
	})

	if err := connect.Ping(context.Background()).Err(); err != nil {
		log.Fatalln(err)
	}

	return &Redis{Client: connect}
}

func (r Redis) Close() {
	log.Println("Stop Redis")

	errlog.Println(r.Client.Close())
}
