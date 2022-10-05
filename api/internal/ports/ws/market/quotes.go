package market

import (
	"log"
	"net/http"
)

func (h HandlerMarket) Quotes(w http.ResponseWriter, r *http.Request) {
	respond := h.responder.Http(w, r)

	conn, err := h.websocket.Upgrade(w, r)
	if err != nil {
		respond.Err(err)
		return
	}

	log.Println(conn)
}
