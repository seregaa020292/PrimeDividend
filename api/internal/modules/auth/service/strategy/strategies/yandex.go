package strategies

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
)

type yandexStrategy struct {
	oauth *oauth2.Config
	strategy.Service
}

func NewYandexStrategy(cfg config.YandexOAuth2, service strategy.Service) categorize.NetworkStrategy {
	return yandexStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     yandex.Endpoint,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
		},
		Service: service,
	}
}

func (y yandexStrategy) Callback(state string) string {
	return y.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (y yandexStrategy) Login(code string, accountability entity.Accountability) (auth.Tokens, error) {
	panic("implement me")
}
