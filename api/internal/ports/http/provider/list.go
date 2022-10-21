package provider

import "net/http"

func (h HandlerProvider) GetProviders(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	respond.Json(http.StatusOK, []any{})
}
