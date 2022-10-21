package decorators

import "context"

type CtxSender[D any] interface {
	Send(ctx context.Context, data D) error
}

type Sender[D any] interface {
	Send(data D) error
}
