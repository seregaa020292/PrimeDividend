package pubsub

import (
	"context"

	client "github.com/go-redis/redis/v8"

	"primedivident/pkg/db/redis"
)

type Subscriber interface {
	Subscribe(channels ...string) <-chan *client.Message
}

type subscriber struct {
	ctx context.Context
	db  *redis.Redis
}

func NewSubscriber(db *redis.Redis) Subscriber {
	return subscriber{
		ctx: context.Background(),
		db:  db,
	}
}

func (s subscriber) Subscribe(channels ...string) <-chan *client.Message {
	return s.db.Subscribe(s.ctx, channels...).Channel()
}
