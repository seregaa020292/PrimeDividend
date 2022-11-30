package graceful

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func SignContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case <-sign:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	return ctx
}
