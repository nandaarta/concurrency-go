package fanout_test

import (
	"context"
	"testing"
	"time"

	"github.com/nandaarta/concurrency-go/fanout"
)

var timeout = 10 * time.Second

func TestFanOutWithoutSemaphore(t *testing.T) {
	t.Parallel()

	// Given
	var (
		numberOfResults         = 20
		expectedNumberOfResults = 20
		expectedValuePerResult  = "any-value"
	)

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	res, err := fanout.FanOut(ctx, numberOfResults)
	if err != nil {
		t.Fatalf("error : %v", err)
	}

	if len(res) != expectedNumberOfResults {
		t.Errorf("expected %d, but actual result is %d", expectedNumberOfResults, len(res))
	}

	for _, v := range res {
		if v != expectedValuePerResult {
			t.Errorf("expected should be always %s, but got actual %s", expectedValuePerResult, v)
		}
	}
}

func TestFanOutWithSemaphore(t *testing.T) {
	t.Parallel()
	var (
		numberOfResults         = 20
		expectedNumberOfResults = 20
		expectedValuePerResult  = "any-value"
	)

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	res, err := fanout.FanOutSemaphore(ctx, numberOfResults)
	if err != nil {
		t.Fatalf("error : %v", err)
	}

	if err != nil {
		t.Fatalf("error : %v", err)
	}

	if len(res) != expectedNumberOfResults {
		t.Errorf("expected %d, but actual result is %d", expectedNumberOfResults, len(res))
	}

	for _, v := range res {
		if v != expectedValuePerResult {
			t.Errorf("expected should be always %s, but got actual %s", expectedValuePerResult, v)
		}
	}
}
