package auth

import (
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/pkg/errorn"
)

func (h HandlerAuth) JoinNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	strategyNetwork := h.strategy.Network().Get(auth.Name(network))

	if strategyNetwork == nil {
		respond.Err(errorn.ErrNotFound)
		return
	}

	state := strategy.GenCookieOauthState(w, r)

	respond.Redirect(strategyNetwork.Callback(state))
}
