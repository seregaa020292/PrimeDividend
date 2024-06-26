package quotes

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"

	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/internal/modules/market/service/quotes/message"
	"primedividend/api/pkg/utils/errlog"
)

const (
	pongWait   = time.Second * 60
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	User      entity.JwtPayload
	Conn      *websocket.Conn
	HubQuotes *HubQuotes
	Message   chan []byte
}

func NewClient(user entity.JwtPayload, conn *websocket.Conn, hubQuotes *HubQuotes) Client {
	//utils.Println(quotes.tinkoff.Subscribe("BBG002GHV6L9"))

	return Client{
		User:      user,
		Conn:      conn,
		HubQuotes: hubQuotes,
		Message:   make(chan []byte),
	}
}

func (c Client) Read() {
	defer c.leave()

	c.setReadDeadline()

	c.Conn.SetPongHandler(func(string) error {
		c.setReadDeadline()
		return nil
	})

	for {
		var msg message.Message
		if err := c.Conn.ReadJSON(&msg); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v\n", err)
			}
			break
		}

		c.handleMessage(msg)
	}
}

func (c Client) Write() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		errlog.FnPrintln(c.Conn.Close)
	}()

	for {
		select {
		case msg, ok := <-c.Message:
			if !ok {
				errlog.Println(c.Conn.WriteMessage(websocket.CloseMessage, []byte{}))
				return
			}

			errlog.Println(c.Conn.WriteMessage(websocket.TextMessage, msg))
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c Client) Send(message message.Message) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	c.Message <- msg
}

func (c Client) Close() {
	close(c.Message)
	errlog.Println(c.Conn.Close())
}

func (c Client) leave() {
	c.HubQuotes.leave <- c
}

func (c Client) handleMessage(msg message.Message) {
}

func (c Client) setReadDeadline() {
	errlog.Println(c.Conn.SetReadDeadline(time.Now().Add(pongWait)))
}
