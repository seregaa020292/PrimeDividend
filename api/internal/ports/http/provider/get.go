package provider

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
)

func (h HandlerProvider) GetProvider(w http.ResponseWriter, r *http.Request, providerId openapi.ProviderId) {
	respond := h.responder.Http(w, r)

	provider, err := h.queryGetById.Fetch(providerId)
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(provider))
}
