package auth

import "net/http"

func (h HandlerAuth) AuthEmailConfirm(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
