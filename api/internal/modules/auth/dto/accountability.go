package dto

import (
	"net/http"

	"primedivident/internal/modules/auth/entity"
)

func AccountabilityByRequest(r *http.Request) entity.Accountability {
	return entity.Accountability{
		IP:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
		Origin:    r.Host,
	}
}
