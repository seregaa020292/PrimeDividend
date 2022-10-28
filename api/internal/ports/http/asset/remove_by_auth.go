package asset

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
	"primedividend/api/internal/infrastructure/server/openapi"
	"primedividend/api/internal/modules/asset/command"
)

func (h HandlerAsset) RemoveUserAsset(w http.ResponseWriter, r *http.Request, assetId openapi.AssetId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	if err := h.commandRemove.Exec(command.PayloadRemove{
		UserID:  user.ID,
		AssetID: assetId,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}
