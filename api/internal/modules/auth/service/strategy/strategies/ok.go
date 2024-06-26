package strategies

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/odnoklassniki"

	"primedividend/api/internal/config"
	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/internal/modules/auth/service/strategy"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
	"primedividend/api/internal/modules/auth/service/strategy/categorize"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
	"primedividend/api/pkg/secure"
)

const (
	oauthUrlOK    = "https://api.ok.ru/fb.do?method=%s&application_key=%s&sig=%s&access_token=%s"
	oauthSigOK    = "application_key=%smethod=%s%s"
	oauthMethodOK = "users.getCurrentUser"
)

type okStrategy struct {
	oauth  *oauth2.Config
	config config.OkOAuth2
	strategy.Service
}

func NewOkStrategy(cfg config.OkOAuth2, service strategy.Service) categorize.NetworkStrategy {
	return okStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     odnoklassniki.Endpoint,
		},
		config:  cfg,
		Service: service,
	}
}

func (s okStrategy) Callback(state string) string {
	return s.oauth.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s okStrategy) Login(code string, accountability auth.Accountability) (auth.Tokens, error) {
	var response responseOK

	if err := s.ClientNetwork(&response, code, s.oauth, s.urlApi); err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.EncounteredRequestExternal)
	}

	network := entity.Network{
		ClientID: response.UID,
		Email:    response.Email,
		Name:     fmt.Sprintf("%s %s", response.LastName, response.FirstName),
	}

	user, err := s.UserAttachNetwork(network, auth.Ok)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	tokens, err := s.CreateSessionTokens(user, accountability)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return tokens, nil
}

func (s okStrategy) urlApi(token *oauth2.Token) string {
	secretKey := secure.GetMD5Hash(fmt.Sprintf("%s%s", token.AccessToken, s.config.ClientSecret))

	sign := secure.GetMD5Hash(fmt.Sprintf(oauthSigOK, s.config.ClientKey, oauthMethodOK, secretKey))

	return fmt.Sprintf(oauthUrlOK, oauthMethodOK, s.config.ClientKey, sign, token.AccessToken)
}
