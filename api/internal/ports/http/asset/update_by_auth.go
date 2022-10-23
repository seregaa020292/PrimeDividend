package asset

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/asset/command"
)

func (h HandlerAsset) UpdateUserAsset(w http.ResponseWriter, r *http.Request, assetId openapi.AssetId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	var asset openapi.AssetUpdate

	if err := respond.DecodeValidate(&asset); err != nil {
		respond.Err(err)
		return
	}

	if err := h.commandEdit.Exec(command.PayloadUpdate{
		UserID:     user.ID,
		AssetID:    assetId,
		Quantity:   asset.Quantity,
		Amount:     asset.Amount,
		NotationAt: asset.NotationAt,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusOK)
}
