package auth

import (
	"net/http"

	"primedivident/internal/config/consts"
	"primedivident/pkg/token"
	"primedivident/pkg/utils"
)

func GetCookieRefreshToken(r *http.Request) (string, error) {
	refreshToken, err := r.Cookie(consts.RefreshToken)
	if err != nil {
		return "", err
	}

	return refreshToken.Value, nil
}

func SetCookieRefreshToken(refreshToken token.Token, w http.ResponseWriter, r *http.Request) {
	cookie := utils.GenCookie(consts.RefreshToken, refreshToken.Value, &http.Cookie{
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Domain:   r.URL.Hostname(),
		Expires:  refreshToken.ExpiresAt,
	})

	http.SetCookie(w, cookie)
}

func RemoveCookieRefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie := utils.GenCookie(consts.RefreshToken, "", &http.Cookie{
		SameSite: http.SameSiteStrictMode,
		Domain:   r.URL.Hostname(),
		MaxAge:   -1,
	})

	http.SetCookie(w, cookie)
}
