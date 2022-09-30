package strategy

import (
	"fmt"
	"net/http"

	"primedivident/pkg/token"
	"primedivident/pkg/utils"
)

const RefreshToken = "refresh-token"

func GetCookieRefreshToken(r *http.Request) (string, error) {
	refreshToken, err := r.Cookie(RefreshToken)
	if err != nil {
		return "", err
	}
	if refreshToken.Value == "" {
		return "", fmt.Errorf("%s is empty", RefreshToken)
	}

	return refreshToken.Value, nil
}

func SetCookieRefreshToken(refreshToken token.Token, w http.ResponseWriter, r *http.Request) {
	cookie := utils.GenCookie(RefreshToken, refreshToken.Value, &http.Cookie{
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Domain:   r.URL.Hostname(),
		Expires:  refreshToken.ExpiresAt,
	})

	http.SetCookie(w, cookie)
}

func RemoveCookieRefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie := utils.GenCookie(RefreshToken, "", &http.Cookie{
		SameSite: http.SameSiteStrictMode,
		Domain:   r.URL.Hostname(),
		MaxAge:   -1,
	})

	http.SetCookie(w, cookie)
}
