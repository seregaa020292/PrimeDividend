package quotes

import "github.com/google/uuid"

type Clients map[uuid.UUID]Client

func (c Clients) Add(client Client) {
	if _, ok := c[client.User.ID]; !ok {
		c[client.User.ID] = client
	}
}

func (c Clients) Exist(client Client) bool {
	_, ok := c[client.User.ID]
	return ok
}
