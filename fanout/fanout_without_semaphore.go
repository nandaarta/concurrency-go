package fanout

import (
	"context"
	"fmt"
)

func FanOut(ctx context.Context, numOfWorkers int) (result []string, err error) {
	workers := make([]<-chan string, numOfWorkers)

	for i := 0; i < numOfWorkers; i++ {
		workers[i] = generate(ctx)
	}

	result = make([]string, 0, numOfWorkers)

	for i := range workers {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("worker data %s", ctx.Err())
		case value := <-workers[i]: // gets All data from all channel here
			result = append(result, value)
		}
	}

	return
}

func generate(ctx context.Context) chan string {
	dataStream := make(chan string)

	go func(ctx context.Context) {
		defer close(dataStream)

		select {
		case <-ctx.Done():
			return
		case dataStream <- "any-value":
			// do something or do nothing at all
		}
	}(ctx)

	return dataStream
}
