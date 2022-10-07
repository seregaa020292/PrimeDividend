package market

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"primedivident/internal/infrastructure/server/middlewares/helper"
)

func (h HandlerMarket) Quotes(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	user, err := helper.UserFromCtx(r.Context())
	if err != nil {
		respond.Err(err)
		return
	}

	conn, err := h.websocket.Upgrade(w, r)
	if err != nil {
		respond.Err(err)
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ID: %s, Role: %s", user.ID, user.Role)))
}
