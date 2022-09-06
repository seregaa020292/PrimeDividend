package mailer

import (
	"fmt"
	"primedivident/pkg/logger"
)

type Observer struct {
	chanMsg chan Message
	sender  Sender
	logger  logger.Logger
}

func NewObserver(sender Sender, poolConn int, logger logger.Logger) Sender {
	observer := &Observer{
		chanMsg: make(chan Message),
		sender:  sender,
		logger:  logger,
	}

	observer.listens(poolConn)

	return observer
}

func (o *Observer) Send(msg Message) error {
	o.chanMsg <- msg
	return nil
}

func (o *Observer) Close() {
	close(o.chanMsg)
}

func (o *Observer) listens(poolConn int) {
	for i := 0; i < poolConn; i++ {
		go o.listen()
	}
}

func (o *Observer) listen() {
	for msg := range o.chanMsg {
		if err := o.sender.Send(msg); err != nil {
			bytes, _ := msg.Bytes()
			o.logger.
				ExtraField("email", string(bytes)).
				ExtraField("fields", fmt.Sprintf("%+v", msg)).
				Errorf("Email send error: %v", err)
		}
	}
}
