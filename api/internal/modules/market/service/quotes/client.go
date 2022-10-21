package quotes

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"

	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/utils/errlog"
)

const (
	pongWait   = time.Second * 60
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	User    entity.JwtPayload
	Conn    *websocket.Conn
	Quotes  *Quotes
	Message chan []byte
}

func NewClient(user entity.JwtPayload, conn *websocket.Conn, quotes *Quotes) Client {
	//utils.Println(quotes.tinkoff.Subscribe("BBG002GHV6L9"))

	return Client{
		User:    user,
		Conn:    conn,
		Quotes:  quotes,
		Message: make(chan []byte),
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
		var message Message
		if err := c.Conn.ReadJSON(&message); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v\n", err)
			}
			break
		}

		c.handleMessage(message)
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
		case message, ok := <-c.Message:
			if !ok {
				errlog.Println(c.Conn.WriteMessage(websocket.CloseMessage, []byte{}))
				return
			}

			errlog.Println(c.Conn.WriteMessage(websocket.TextMessage, message))
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c Client) Send(message Message) {
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
	c.Quotes.leave <- c
}

func (c Client) handleMessage(message Message) {
}

func (c Client) setReadDeadline() {
	errlog.Println(c.Conn.SetReadDeadline(time.Now().Add(pongWait)))
}
