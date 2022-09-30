package dto

import (
	"net/http"

	"primedivident/internal/modules/auth/service/strategy/auth"
)

func AccountabilityByRequest(r *http.Request) auth.Accountability {
	return auth.Accountability{
		IP:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
		Origin:    r.Host,
	}
}
