package provider

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/provider/query"
)

func (h HandlerProvider) GetProvider(w http.ResponseWriter, r *http.Request, providerId openapi.ProviderId) {
	respond := h.responder.Http(w, r)

	provider, err := h.queryGetById.Fetch(query.ID(providerId))
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetOne(provider))
}
