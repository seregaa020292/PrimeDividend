package auth

import (
	"net/http"

	"github.com/google/uuid"

	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/pkg/datetime"
	"primedivident/pkg/errorn"
	"primedivident/pkg/utils"
)

func (h HandlerAuth) JoinNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	strategy := h.authService.NetworkStrategy(auth.Key(network))

	if strategy == nil {
		respond.Err(errorn.ErrNotFound)
		return
	}

	state := uuid.New().String()

	cookie := utils.GenCookie(consts.OauthState, state, &http.Cookie{
		Secure:   true,
		HttpOnly: true,
		Domain:   r.URL.Hostname(),
		Expires:  datetime.GetNow().AddDate(1, 0, 0),
	})

	http.SetCookie(w, cookie)

	respond.Redirect(strategy.Callback(state))
}
