package quotes

import (
	"github.com/gorilla/websocket"

	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/utils"
)

type Quotes struct {
	clients Clients
	join    chan Client
	leave   chan Client
}

func NewQuotes() *Quotes {
	return &Quotes{
		clients: make(Clients),
		join:    make(chan Client),
		leave:   make(chan Client),
	}
}

func (q Quotes) Join(user entity.JwtPayload, conn *websocket.Conn) {
	client := Client{
		User:   user,
		Conn:   conn,
		Quotes: &q,
	}

	q.join <- client

	go client.Read()
}

func (q Quotes) Run() {
	for {
		select {
		case client := <-q.join:
			q.add(client)
		case client := <-q.leave:
			q.disconnect(client)
		}
	}
}

func (q Quotes) broadcast(message Message) {
	for _, client := range q.clients {
		client.Write(message)
	}
}

func (q Quotes) add(client Client) {
	q.clients.Add(client)
}

func (q Quotes) disconnect(client Client) {
	if q.clients.Exist(client) {
		defer utils.Println(client.Conn.Close())

		delete(q.clients, client.User.ID)
	}
}
