package heartbeat_test

import (
	"slices"
	"testing"
	"time"

	"github.com/nandaarta/concurrency-go/heartbeat"
)

func TestHeartbeat(t *testing.T) {
	t.Parallel() // parallel execution with other test case

	var (
		pulseInput    []string      = []string{"piip", "puup", "pap", "piup", "pwip", "paap", "pouep", "peep", "peop"}
		timeout       time.Duration = 2 * time.Second
		pulseInterval time.Duration = 20 * time.Millisecond
		done          chan any
	)

	done = make(chan any)

	time.AfterFunc(timeout, func() { close(done) })

	res, pulses := heartbeat.DoSomething(done, pulseInput, pulseInterval)

	for _, pulse := range pulses {
		if pulse != heartbeat.SomeBeat {
			t.Fatalf("expected: %q, but got: %q", heartbeat.SomeBeat, pulse)
		}
	}

	if len(res) != len(pulseInput) {
		t.Fatalf("expected : %d, but got: %d", len(pulseInput), len(res))
	}

	for _, pulseString := range pulseInput {
		if !slices.Contains(res, pulseString) {
			t.Errorf("expected: %q, but doesn't exist", pulseString)
		}
	}

}
