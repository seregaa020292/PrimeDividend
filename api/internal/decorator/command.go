package decorator

import "context"

type CommandCtxHandler[C any] interface {
	Exec(ctx context.Context, cmd C) error
}

type CommandHandler[C any] interface {
	Exec(cmd C) error
}
