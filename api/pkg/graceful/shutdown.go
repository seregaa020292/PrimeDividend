package graceful

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type (
	Operation  func(ctx context.Context) error
	Operations []Operation
)

type Graceful struct {
	ctx        context.Context
	timeout    time.Duration
	operations *operations
}

func NewGraceful(timeout time.Duration) Graceful {
	return Graceful{
		ctx:     context.Background(),
		timeout: timeout,
	}
}

type operations struct {
	len  int
	list Operations

	ctx context.Context
	wg  sync.WaitGroup
}

func newOperations(ctx context.Context, list Operations) *operations {
	return &operations{
		len:  len(list),
		list: list,
		ctx:  ctx,
		wg:   sync.WaitGroup{},
	}
}

func (g *Graceful) Shutdown(operations Operations) {
	g.operations = newOperations(g.ctx, operations)

	if g.operations.isEmpty() {
		return
	}

	go func() {
		<-g.stop()

		timeAfterExecuted := g.forceShutdown()
		defer timeAfterExecuted.Stop()

		g.operations.Calls()
	}()
}

func (g *Graceful) stop() chan os.Signal {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return stopChan
}

func (g *Graceful) forceShutdown() *time.Timer {
	return time.AfterFunc(g.timeout, func() {
		log.Println("Force shutdown server")

		os.Exit(0)
	})
}

func (o *operations) isEmpty() bool {
	return o.len == 0
}

func (o *operations) Calls() {
	o.wg.Add(o.len)

	for key, op := range o.list {
		go o.call(op, key+1)
	}

	o.wg.Wait()
}

func (o *operations) call(op Operation, key int) {
	defer o.wg.Done()

	log.Printf("Shutdown %d/%d\n", key, o.len)

	if err := op(o.ctx); err != nil {
		log.Fatalf("Error when stop server: %s\n", err)
	}
}
