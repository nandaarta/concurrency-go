package context

import "time"

func DoSomething(doneCh <-chan any) <-chan string {
	result := make(chan string)

	go func() {
		defer close(result)

		select {
		case <-doneCh:
			result <- "process not completed !!"
		case <-time.After(5 * time.Second):
			result <- "process completed !!"
		}
	}()

	return result
}
