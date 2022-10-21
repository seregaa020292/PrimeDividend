package provider

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerProvider) GetProvider(w http.ResponseWriter, r *http.Request, providerId openapi.ProviderId) {
	respond := h.responder.Http(w, r)

	respond.Json(http.StatusOK, nil)
}
