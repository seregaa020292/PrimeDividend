package config

import "fmt"

type Redis struct {
	Host string `env:"REDIS_HOST" env-required:"true"`
	Port int    `env:"REDIS_PORT" env-required:"true"`
}

func (r Redis) Dsn() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
