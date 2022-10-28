package market

import (
	"net/http"

	"primedividend/api/internal/infrastructure/server/middlewares/helper"
)

func (h HandlerMarket) Quotes(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	//portfolioID, err := helper.UUID(chi.URLParam(r, "portfolioId"))
	//if err != nil {
	//	respond.Err(err)
	//	return
	//}

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	conn, err := h.socket.Upgrade(w, r)
	if err != nil {
		respond.Err(err)
		return
	}

	go h.quotes.Run()

	h.quotes.Join(user, conn)
}
