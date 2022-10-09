package quotes

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/utils"
)

type (
	Clients map[uuid.UUID]Client
	Client  struct {
		User   entity.JwtPayload
		Conn   *websocket.Conn
		Quotes *Quotes
	}
)

func (c Client) Read() {
	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}

	c.Quotes.leave <- c
}

func (c Client) Write(message Message) {
	msg, err := json.Marshal(message)

	utils.Println(err)

	utils.Println(c.Conn.WriteMessage(websocket.TextMessage, msg))
}

func (c Clients) Add(client Client) {
	if _, ok := c[client.User.ID]; !ok {
		c[client.User.ID] = client
	}
}

func (c Clients) Exist(client Client) bool {
	_, ok := c[client.User.ID]
	return ok
}
