package asset

import (
	"net/http"

	"primedivident/internal/infrastructure/server/middlewares/helper"
	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/internal/modules/asset/command"
)

func (h HandlerAsset) CreateUserAsset(w http.ResponseWriter, r *http.Request, portfolioId openapi.PortfolioId) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	var asset openapi.AssetAdd

	if err := respond.DecodeValidate(&asset); err != nil {
		respond.Err(err)
		return
	}

	if err := h.commandCreate.Exec(command.PayloadCreate{
		UserID:      user.ID,
		PortfolioID: portfolioId,
		Amount:      asset.Amount,
		Quantity:    asset.Quantity,
		MarketID:    asset.MarketId,
		NotationAt:  asset.NotationAt,
	}); err != nil {
		respond.Err(err)
		return
	}

	respond.WriteHeader(http.StatusNoContent)
}
