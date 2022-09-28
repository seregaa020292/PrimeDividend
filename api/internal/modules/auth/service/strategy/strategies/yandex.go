package strategies

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/internal/modules/auth/service/strategy/repository"
)

type yandexStrategy struct {
	oauth      *oauth2.Config
	jwtTokens  auth.JwtTokens
	repository repository.Repository
}

func NewYandexStrategy(
	cfg config.YandexOAuth2,
	jwtTokens auth.JwtTokens,
	repository repository.Repository,
) categorize.NetworkStrategy {
	return yandexStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     yandex.Endpoint,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
		},
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (y yandexStrategy) Callback(state string) string {
	return y.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (y yandexStrategy) Login(code string, session entity.FingerprintSession) (auth.Tokens, error) {
	panic("implement me")
}
