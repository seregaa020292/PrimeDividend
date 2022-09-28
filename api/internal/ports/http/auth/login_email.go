package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
)

func (h HandlerAuth) LoginEmail(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	var user openapi.LoginUser

	if err := respond.DecodeValidate(&user); err != nil {
		respond.Err(err)
		return
	}

	tokens, err := h.strategy.Password().Get(auth.Email).Login(
		user.Email,
		user.Password,
		entity.FingerprintSession{
			IP:        r.RemoteAddr,
			UserAgent: r.UserAgent(),
		})
	if err != nil {
		respond.Err(err)
		return
	}

	strategy.SetCookieRefreshToken(tokens.RefreshToken, w, r)

	respond.Json(http.StatusOK, tokens.AccessToken)
}
