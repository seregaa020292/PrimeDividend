package quotes

import (
	"github.com/google/uuid"

	"primedividend/api/internal/modules/market/service/quotes/message"
)

type Clients struct {
	connects map[uuid.UUID]Client
}

func NewClients() Clients {
	return Clients{
		connects: make(map[uuid.UUID]Client),
	}
}

func (c Clients) Add(client Client) {
	if _, ok := c.connects[client.User.ID]; !ok {
		c.connects[client.User.ID] = client
	}
}

func (c Clients) Remove(clientID uuid.UUID) {
	delete(c.connects, clientID)
}

func (c Clients) Exist(client Client) bool {
	_, ok := c.connects[client.User.ID]
	return ok
}

func (c Clients) Broadcast(message message.Message) {
	for _, client := range c.connects {
		client.Send(message)
	}
}

func (c Clients) Length() int {
	return len(c.connects)
}
