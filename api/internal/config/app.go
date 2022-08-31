package config

type App struct {
	Env        string `env:"APP_ENV" env-default:"development"`
	SiteOrigin string `env:"SITE_ORIGIN" env-default:"http://localhost"`
	LogLevel   string `env:"LOG_LEVEL" env-default:"debug"`
}

func (a App) IsDevelopment() bool {
	return a.Env == "development"
}
