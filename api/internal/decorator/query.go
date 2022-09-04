package decorator

import "context"

type QueryCtxHandler[Q any, R any] interface {
	Fetch(ctx context.Context, q Q) (R, error)
}

type QueryHandler[Q any, R any] interface {
	Fetch(q Q) (R, error)
}
