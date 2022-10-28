package auth

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/auth/dto"
	"primedividend/api/internal/modules/auth/service/strategy"
)

func (h HandlerAuth) RefreshToken(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	refreshToken, err := strategy.GetCookieRefreshToken(r)
	if err != nil {
		respond.Err(err)
		return
	}

	if err := h.strategy.VerifyRefresh(refreshToken); err != nil {
		strategy.RemoveCookieRefreshToken(w, r)
		respond.Err(err)
		return
	}

	tokens, err := h.strategy.Refresh(refreshToken, dto.AccountabilityByRequest(r))
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
