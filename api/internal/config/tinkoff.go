package config

type Tinkoff struct {
	AuthToken string `env:"TINKOFF_AUTH_TOKEN" env-required:"true"`
}
