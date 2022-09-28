package auth

import (
	"net/http"

	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/pkg/errorn"
)

func (h HandlerAuth) Logout(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	refreshToken, err := strategy.GetCookieRefreshToken(r)
	if err != nil {
		respond.Err(errorn.ErrForbidden.Wrap(err))
		return
	}

	if err := h.strategy.Logout(refreshToken); err != nil {
		respond.Err(errorn.ErrUnknown.Wrap(err))
		return
	}

	strategy.RemoveCookieRefreshToken(w, r)

	respond.WriteHeader(http.StatusOK)
}