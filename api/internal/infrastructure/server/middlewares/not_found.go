package middlewares

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/pkg/errs"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	respond := response.NewRespondBuilder(w, r)

	err := errs.NotFound.New("Не найдено")

	respond.Err(err)
}
