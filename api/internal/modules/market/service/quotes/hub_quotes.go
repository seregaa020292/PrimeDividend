package quotes

import (
	"log"

	"github.com/gorilla/websocket"

	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/market/repository"
	"primedivident/internal/modules/market/service/quotes/providers"
)

type HubQuotes struct {
	assetRepo repository.AssetRepository
	tinkoff   providers.Tinkoff
	clients   Clients
	join      chan Client
	leave     chan Client
}

func NewHubQuotes(assetRepo repository.AssetRepository, tinkoff providers.Tinkoff) *HubQuotes {
	return &HubQuotes{
		assetRepo: assetRepo,
		tinkoff:   tinkoff,
		clients:   NewClients(),
		join:      make(chan Client),
		leave:     make(chan Client),
	}
}

func (q HubQuotes) Join(user entity.JwtPayload, conn *websocket.Conn) {
	client := NewClient(user, conn, &q)

	q.join <- client

	go q.tinkoff.Read()
	go client.Read()
	go client.Write()
}

func (q HubQuotes) Run() {
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

func (q HubQuotes) Close() {
	log.Println("Stop HubQuotes")

	q.tinkoff.Close()
	q.disconnects()
}

func (q HubQuotes) add(client Client) {
	q.clients.Add(client)
}

func (q HubQuotes) disconnects() {
	for _, client := range q.clients.connects {
		q.disconnect(client)
	}
}

func (q HubQuotes) disconnect(client Client) {
	if q.clients.Exist(client) {
		client.Close()
		q.clients.Remove(client.User.ID)
	}
}
