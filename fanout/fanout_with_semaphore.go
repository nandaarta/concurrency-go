package fanout

import (
	"context"
	"fmt"
	"runtime"
)

func FanOutSemaphore(ctx context.Context, numOfWorkers int) (result []string, err error) {
	workers := make([]<-chan string, numOfWorkers)
	sem := newSemaphore()

	for i := 0; i < numOfWorkers; i++ {
		workers[i] = generateWithSemaphore(ctx, sem)
	}

	result = make([]string, 0, numOfWorkers)

	for i := range workers {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("worker data %s", ctx.Err())
		case value := <-workers[i]:
			result = append(result, value)
		}
	}

	return
}

func newSemaphore() chan bool {
	workerProcess := runtime.GOMAXPROCS(0) // you can set number of CPUs depending on your usecase

	return make(chan bool, workerProcess)
}

func generateWithSemaphore(ctx context.Context, semaphore chan bool) chan string {
	dataStream := make(chan string, 1) // set data stream as buffered channel

	go func(ctx context.Context) {
		defer close(dataStream)

		semaphore <- true
		{
			select {
			case <-ctx.Done():
				return
			case dataStream <- "any-value":
				// do something or do nothing at all
			}
		}
		<-semaphore
	}(ctx)

	return dataStream

}
