package auth

import (
	"fmt"
	"net/http"

	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/pkg/errorn"
)

func (h HandlerAuth) ConfirmNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	if err := auth.ValidateOauthState(r); err != nil {
		respond.Err(errorn.ErrForbidden.Wrap(err))
		return
	}

	strategy := h.authService.NetworkStrategy(auth.Key(network))
	if strategy == nil {
		err := fmt.Errorf("strategy %s not found", network)
		respond.Err(errorn.ErrNotFound.Wrap(err))
		return
	}

	tokens, err := strategy.Login(r.FormValue("code"))
	if err != nil {
		respond.Err(errorn.ErrUnauthorized.Wrap(err))
		return
	}

	auth.SetCookieRefreshToken(tokens.RefreshToken, w, r)

	respond.Json(http.StatusAccepted, tokens.AccessToken)
}
