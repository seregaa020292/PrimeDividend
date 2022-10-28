package auth

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/auth/dto"
	"primedividend/api/internal/modules/auth/service/strategy"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
)

func (h HandlerAuth) LoginEmail(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	var user openapi.LoginUser

	if err := respond.DecodeValidate(&user); err != nil {
		respond.Err(err)
		return
	}

	tokens, err := h.strategy.Password().Get(auth.Email).
		Login(user.Email, user.Password, dto.AccountabilityByRequest(r))
	if err != nil {
		respond.Err(err)
		return
	}

	strategy.SetCookieRefreshToken(tokens.RefreshToken, w, r)

	respond.Json(http.StatusOK, openapi.AuthToken{
		AccessToken: tokens.AccessToken.Value,
		ExpiresAt:   tokens.AccessToken.ExpiresAt,
	})
}
