package quotes

import (
	"github.com/gorilla/websocket"

	"primedivident/internal/modules/auth/entity"
)

type Quotes struct {
	tinkoff Tinkoff
	clients Clients
	join    chan Client
	leave   chan Client
}

func NewQuotes(tinkoff Tinkoff) *Quotes {
	return &Quotes{
		tinkoff: tinkoff,
		clients: NewClients(),
		join:    make(chan Client),
		leave:   make(chan Client),
	}
}

func (q Quotes) Join(user entity.JwtPayload, conn *websocket.Conn) {
	client := NewClient(user, conn, &q)

	q.join <- client

	go q.tinkoff.Read()
	go client.Read()
	go client.Write()
}

func (q Quotes) Run() {
	for {
		select {
		case client := <-q.join:
			q.add(client)
		case client := <-q.leave:
			q.disconnect(client)
		case message := <-q.tinkoff.Message():
			q.clients.Broadcast(message)
		}
	}
}

func (q Quotes) add(client Client) {
	q.clients.Add(client)
}

func (q Quotes) disconnect(client Client) {
	if q.clients.Exist(client) {
		defer client.Close()
		defer q.tinkoff.Close()

		q.clients.Remove(client.User.ID)
	}
}
