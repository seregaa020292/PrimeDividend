package middlewares

import (
	"net/http"

	"primedivident/internal/infrastructure/server/response"
	"primedivident/pkg/errs"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	respond := response.NewRespondBuilder(w, r)

	err := errs.NotFound.New("Не найдено")

	respond.Err(err)
}
