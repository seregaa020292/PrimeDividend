package auth

import (
	"fmt"
	"net/http"

	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/auth/service/auth/strategies"
	"primedivident/pkg/logger"
)

func (h HandlerAuth) ConfirmNetwork(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := h.responder.Http(w, r)

	oauthState, _ := r.Cookie(consts.OauthState)
	state := r.FormValue("state")
	code := r.FormValue("code")

	if state != oauthState.Value {
		logger.GetLogger().Errorf("%s", "Invalid oauth google state")
		respond.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	strategy := h.strategies.Networks.GetStrategy(strategies.Key(network))
	if strategy == nil {
		logger.GetLogger().Errorf("%s", "Invalid oauth google state")
		respond.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	data, err := strategy.Login(code)
	if err != nil {
		logger.GetLogger().Errorf("%s", err)
		respond.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", data)
	//respond.Json(http.StatusAccepted, string(data))
}
