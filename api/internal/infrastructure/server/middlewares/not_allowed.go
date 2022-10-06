package middlewares

import (
	"net/http"

	"primedivident/internal/infrastructure/server/response"
	"primedivident/pkg/errs"
)

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	respond := response.NewRespondBuilder(w, r)

	err := errs.NotFound.New("метод не поддерживается")

	respond.Err(err)
}
