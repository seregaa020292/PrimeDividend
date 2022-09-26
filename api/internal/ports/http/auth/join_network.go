package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/pkg/errorn"
)

func (h HandlerAuth) JoinNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	strategy := h.authService.NetworkStrategy(auth.Key(network))

	if strategy == nil {
		respond.Err(errorn.ErrNotFound)
		return
	}

	state := auth.GenCookieOauthState(w, r)

	respond.Redirect(strategy.Callback(state))
}
