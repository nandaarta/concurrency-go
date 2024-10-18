package fanin

import (
	"context"
	"sync"
)

func FanIn(ctx context.Context, workers ...<-chan int) <-chan int {
	var (
		wg              sync.WaitGroup
		multiplexedData chan int
	)

	multiplexedData = make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()

		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedData <- i * 10: // let's say we want to multiply the value by 10
				// do something or nothing here
			}
		}
	}

	wg.Add(len(workers))
	for _, w := range workers {
		go multiplex(w)
	}

	// wait all the reads to complete without blocking
	go func() {
		wg.Wait()
		close(multiplexedData)
	}()

	return multiplexedData
}

func newWorker(ctx context.Context, value int) <-chan int {
	valueCh := make(chan int)

	go func() {
		defer close(valueCh) // close this to notify receiver that channel is closed

		select {
		case <-ctx.Done():
			return
		case valueCh <- value:
			// do something or nothing at all
		}
	}()

	return valueCh
}

func WorkerGenerator(ctx context.Context, values ...int) []<-chan int {
	res := make([]<-chan int, 0, len(values))

	for _, value := range values {
		res = append(res, newWorker(ctx, value))
	}
	return res
}
