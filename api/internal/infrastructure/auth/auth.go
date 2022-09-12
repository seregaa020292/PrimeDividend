package auth

import (
	"context"
	"io"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"

	"primedivident/pkg/datetime"
	"primedivident/pkg/utils"
)

const OauthState = "oauth-state"

func GenStateOauthCookie(w http.ResponseWriter, r *http.Request) string {
	newUUID, _ := uuid.NewUUID()

	state := newUUID.String()

	cookie := utils.GenCookie(OauthState, state, &http.Cookie{
		Secure:   true,
		HttpOnly: true,
		Domain:   r.URL.Hostname(),
		Expires:  datetime.GetNow().AddDate(1, 0, 0),
	})

	http.SetCookie(w, cookie)

	return state
}

func GetUser(code string, cfg *oauth2.Config) ([]byte, error) {
	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := VkOAuth2Config.Client(context.Background(), token)

	response, err := client.Get(oauthVkUrlAPI)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
