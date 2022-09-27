package config

import "time"

type Jwt struct {
	AccessSecretKey  string        `env:"JWT_ACCESS_SECRET_KEY" env-required:"true"`
	AccessExpiresIn  time.Duration `env:"JWT_ACCESS_EXPIRES_IN" env-required:"true"`
	RefreshSecretKey string        `env:"JWT_REFRESH_SECRET_KEY" env-required:"true"`
	RefreshExpiresIn time.Duration `env:"JWT_REFRESH_EXPIRES_IN" env-required:"true"`
}
