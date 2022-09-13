package decorator

import "context"

type CommandHandler[C any] interface {
	Exec(cmd C) error
}

type CommandCtxHandler[C any] interface {
	Exec(ctx context.Context, cmd C) error
}

type QueryHandler[Q any, R any] interface {
	Fetch(q Q) (R, error)
}

type QueryCtxHandler[Q any, R any] interface {
	Fetch(ctx context.Context, q Q) (R, error)
}
