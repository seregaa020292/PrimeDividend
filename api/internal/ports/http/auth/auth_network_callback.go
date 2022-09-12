package auth

import (
	"fmt"
	"net/http"

	"primedivident/internal/infrastructure/auth"
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/pkg/logger"
	"primedivident/pkg/response"
)

func (h HandlerAuth) AuthNetworkCallback(w http.ResponseWriter, r *http.Request, network openapi.Network) {
	respond := response.New(w, r)

	oauthState, _ := r.Cookie(auth.OauthState)
	state := r.FormValue("state")
	code := r.FormValue("code")

	if state != oauthState.Value {
		logger.GetLogger().Errorf("%s", "Invalid oauth google state")
		respond.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	data, err := auth.GetUser(code, auth.VkOAuth2Config)
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
