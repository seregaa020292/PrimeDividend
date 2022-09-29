package strategies

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/odnoklassniki"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/pkg/errorn"
	"primedivident/pkg/secure"
)

const (
	oauthUrlOK    = "https://api.ok.ru/fb.do?method=%s&application_key=%s&sig=%s&access_token=%s"
	oauthMethodOK = "users.getCurrentUser"
	oauthSigOK    = "application_key=%smethod=%s%s"
)

type okStrategy struct {
	clientKey string
	oauth     *oauth2.Config
	strategy.Service
}

func NewOkStrategy(cfg config.OkOAuth2, service strategy.Service) categorize.NetworkStrategy {
	return okStrategy{
		clientKey: cfg.ClientKey,
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     odnoklassniki.Endpoint,
		},
		Service: service,
	}
}

func (s okStrategy) Callback(state string) string {
	return s.oauth.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s okStrategy) Login(code string, accountability entity.Accountability) (auth.Tokens, error) {
	var response responseOK

	_, err := s.ClientNetwork(&response, s.oauth, code, s.urlApi)
	if err != nil {
		return auth.Tokens{}, err
	}

	network := entity.Network{
		Identity: response.UID,
		Email:    response.Email,
		Name:     fmt.Sprintf("%s %s", response.LastName, response.FirstName),
	}

	user, err := s.UserAttachNetwork(network, auth.Ok)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnauthorized.Wrap(err)
	}

	return s.CreateSessionTokens(auth.Ok, user, accountability)
}

func (s okStrategy) urlApi(token *oauth2.Token) string {
	secretKey := secure.GetMD5Hash(fmt.Sprintf("%s%s", token.AccessToken, s.oauth.ClientSecret))

	sign := secure.GetMD5Hash(fmt.Sprintf(oauthSigOK, s.clientKey, oauthMethodOK, secretKey))

	return fmt.Sprintf(oauthUrlOK, oauthMethodOK, s.clientKey, sign, token.AccessToken)
}
