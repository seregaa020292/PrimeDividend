package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

func (h HandlerAuth) ConfirmNetwork(
	w http.ResponseWriter,
	r *http.Request,
	network openapi.Network,
	params openapi.ConfirmNetworkParams,
) {
	respond := h.responder.Http(w, r)

	if err := strategy.ValidateOauthState(params.State, r); err != nil {
		respond.Err(err)
		return
	}

	strategyNetwork := h.strategy.Network().Get(auth.Name(network))
	if strategyNetwork == nil {
		respond.Err(errs.NotFound.New(errmsg.CouldNotBeFound))
		return
	}

	tokens, err := strategyNetwork.Login(params.Code, dto.AccountabilityByRequest(r))
	if err != nil {
		respond.Err(err)
		return
	}

	strategy.SetCookieRefreshToken(tokens.RefreshToken, w, r)

	respond.Json(http.StatusAccepted, openapi.AuthToken{
		AccessToken: tokens.AccessToken.Value,
		ExpiresAt:   tokens.AccessToken.ExpiresAt,
	})
}
