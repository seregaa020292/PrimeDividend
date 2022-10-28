package middlewares

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/pkg/errs"
)

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	respond := response.NewRespondBuilder(w, r)

	err := errs.NotFound.New("Метод не поддерживается")

	respond.Err(err)
}
