package auth

import (
	"net/http"

	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/pkg/errorn"
)

func (h HandlerAuth) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	refreshToken, err := strategy.GetCookieRefreshToken(r)
	if err != nil {
		respond.Err(errorn.ErrForbidden.Wrap(err))
		return
	}

	tokens, err := h.strategy.Refresh(refreshToken, dto.AccountabilityByRequest(r))
	if err != nil {
		respond.Err(err)
		return
	}

	strategy.SetCookieRefreshToken(tokens.RefreshToken, w, r)

	respond.Json(http.StatusOK, tokens.AccessToken)
}
