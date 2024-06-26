package provider

import (
	"net/http"

	"primedividend/api/internal/modules/provider/query"
)

func (h HandlerProvider) GetProviders(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	instruments, err := h.queryGetAll.Fetch(query.PayloadAll{})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, h.presenter.GetAll(instruments))
}
