package strategies

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"

	"primedividend/api/internal/config"
	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/internal/modules/auth/service/strategy"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
	"primedividend/api/internal/modules/auth/service/strategy/categorize"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

const oauthUrlYandex = "https://login.yandex.ru/info"

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

func (s yandexStrategy) Login(code string, accountability auth.Accountability) (auth.Tokens, error) {
	var response responseYandex

	if err := s.ClientNetwork(&response, code, s.oauth, func(token *oauth2.Token) string {
		return oauthUrlYandex
	}); err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.EncounteredRequestExternal)
	}

	network := entity.Network{
		ClientID: response.ClientID,
		Email:    response.DefaultEmail,
		Name:     fmt.Sprintf("%s %s", response.LastName, response.FirstName),
	}

	user, err := s.UserAttachNetwork(network, auth.Yandex)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	tokens, err := s.CreateSessionTokens(user, accountability)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return tokens, nil
}
