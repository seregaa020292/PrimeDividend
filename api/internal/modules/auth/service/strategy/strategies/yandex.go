package strategies

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/pkg/errorn"
)

const OauthUrlYandex = ""

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

func (s yandexStrategy) Callback(state string) string {
	return s.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (s yandexStrategy) Login(code string, accountability entity.Accountability) (auth.Tokens, error) {
	var response responseYandex

	token, err := s.ClientNetwork(&response, s.oauth, code, func(token *oauth2.Token) string {
		return OauthUrlYandex
	})
	if err != nil {
		return auth.Tokens{}, err
	}

	log.Printf("%+v", token)

	network := entity.Network{
		Identity: "",
		Email:    "",
		Name:     "",
	}

	user, err := s.UserAttachNetwork(network, auth.Yandex)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnauthorized.Wrap(err)
	}

	return s.CreateSessionTokens(auth.Yandex, user, accountability)
}
