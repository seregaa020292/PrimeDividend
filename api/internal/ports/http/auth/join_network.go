package auth

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/auth/service/strategy"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

func (h HandlerAuth) JoinNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	strategyNetwork := h.strategy.Network().Get(auth.Name(network))

	if strategyNetwork == nil {
		respond.Err(errs.NotFound.New(errmsg.CouldNotBeFound))
		return
	}

	state := strategy.GenCookieOauthState(w, r)

	respond.Redirect(strategyNetwork.Callback(state))
}
