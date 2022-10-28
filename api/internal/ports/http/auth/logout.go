package auth

import (
	"net/http"

	"primedividend/api/internal/modules/auth/service/strategy"
)

func (h HandlerAuth) Logout(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	refreshToken, err := strategy.GetCookieRefreshToken(r)
	if err != nil {
		respond.Err(err)
		return
	}

	if err := h.strategy.Logout(refreshToken); err != nil {
		respond.Err(err)
		return
	}

	strategy.RemoveCookieRefreshToken(w, r)

	respond.WriteHeader(http.StatusOK)
}
