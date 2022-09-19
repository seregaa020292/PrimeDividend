package config

type App struct {
	Name       string `env:"APP_NAME" env-default:""`
	Env        string `env:"APP_ENV" env-default:"development"`
	SiteOrigin string `env:"SITE_ORIGIN" env-required:"true"`
	LogLevel   string `env:"LOG_LEVEL" env-default:"debug"`
}

func (a App) IsDevelopment() bool {
	return a.Env == "development"
}
