package context_test

import (
	"testing"

	"github.com/nandaarta/concurrency-go/context"
)

func TestDo(t *testing.T) {
	t.Parallel()

	// given done and expected variables
	expected := "process not completed !!"

	doneCh := make(chan any)

	result := context.DoSomething(doneCh)
	close(doneCh)

	resultTest := <-result

	if resultTest != expected {
		t.Errorf("Want : %s, but got result %s", expected, resultTest)
	}
}
