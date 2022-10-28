package pubsub

import (
	"context"

	"primedivident/pkg/db/redis"
)

type Publisher interface {
	Publish(channel string, message any) error
}

type publisher struct {
	ctx context.Context
	db  *redis.Redis
}

func NewPublisher(db *redis.Redis) Publisher {
	return publisher{
		ctx: context.Background(),
		db:  db,
	}
}

func (p publisher) Publish(channel string, message any) error {
	return p.db.Publish(p.ctx, channel, message).Err()
}
