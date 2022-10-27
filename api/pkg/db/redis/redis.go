package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"

	"primedivident/internal/config"
	"primedivident/pkg/utils/errlog"
)

const (
	MinIdleConns = 200
	PoolSize     = 12000
	PoolTimeout  = 240 * time.Second
)

type Redis struct {
	*redis.Client
}

func NewRedis(config config.Redis) *Redis {
	log.Println("Start Redis")

	connect := redis.NewClient(&redis.Options{
		Addr:         config.Dsn(),
		MinIdleConns: MinIdleConns,
		PoolSize:     PoolSize,
		PoolTimeout:  PoolTimeout,
	})

	return &Redis{connect}
}

func (r Redis) Close() {
	log.Println("Stop Redis")

	errlog.Println(r.Client.Close())
}
