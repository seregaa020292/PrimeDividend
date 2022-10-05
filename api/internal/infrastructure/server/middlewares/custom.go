package middlewares

import (
	"net/http"

	"github.com/getkin/kin-openapi/routers"
)

func Custom(router routers.Router) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			route, _, err := router.FindRoute(r)
			if err != nil {
				panic(err)
			}

			// TODO:
			_ = route.Operation.OperationID

			next(w, r)
		}
	}
}
