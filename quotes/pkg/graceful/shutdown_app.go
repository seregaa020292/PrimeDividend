package graceful

import "sync"

type ShutdownApp []func()

func (s ShutdownApp) Run() {
	var wg sync.WaitGroup

	wg.Add(len(s))

	for _, fn := range s {
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(fn)
	}

	wg.Wait()
}
