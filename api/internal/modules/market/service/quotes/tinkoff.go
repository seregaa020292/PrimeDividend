package quotes

import (
	"log"
	"os"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"

	"primedivident/internal/config"
	"primedivident/pkg/utils"
	"primedivident/pkg/utils/errlog"
)

type Tinkoff struct {
	message chan Message
	client  *sdk.StreamingClient
}

func NewTinkoff(config config.Tinkoff) Tinkoff {
	logger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)
	client, err := sdk.NewStreamingClient(logger, config.AuthToken)
	if err != nil {
		log.Printf("Error tinkoff %s\n", err)
	} else {
		log.Println("Start Tinkoff stream")
	}

	return Tinkoff{
		message: make(chan Message),
		client:  client,
	}
}

func (t Tinkoff) Read() {
	defer t.Close()

	if err := t.client.RunReadLoop(t.readLoop); err != nil {
		log.Fatalln(err)
	}
}

func (t Tinkoff) Message() <-chan Message {
	return t.message
}

func (t Tinkoff) Subscribe(identity string) error {
	return t.client.SubscribeCandle(identity, sdk.CandleInterval5Min, utils.RandomString(12))
}

func (t Tinkoff) Unsubscribe(identity string) error {
	return t.client.UnsubscribeCandle(identity, sdk.CandleInterval5Min, utils.RandomString(12))
}

func (t Tinkoff) Close() {
	errlog.Println(t.client.Close())
}

func (t Tinkoff) readLoop(event any) error {
	if sdkEvent, ok := event.(sdk.CandleEvent); ok {
		t.handleMessage(sdkEvent)
	}
	return nil
}

func (t Tinkoff) handleMessage(event sdk.CandleEvent) {
	t.message <- Message{
		Provider: TinkoffProvider,
		Identity: event.Candle.FIGI,
		Price:    event.Candle.ClosePrice,
	}
}
