package fanin_test

import (
	"context"
	"testing"
	"time"

	"github.com/nandaarta/concurrency-go/fanin"
)

func TestFanIn(t *testing.T) {
	t.Parallel()

	// Given values and expected
	values := []int{1, 2, 3, 4, 5}
	expected := map[int]int{10: 1, 20: 2, 30: 3, 40: 4, 50: 5}
	timeout := 2 * time.Second

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	// When
	got := fanin.FanIn(ctx, fanin.WorkerGenerator(ctx, values...)...)
	// Then
	var count int
	for gotValue := range got {
		count++
		_, ok := expected[gotValue]
		if !ok {
			t.Errorf("unknown result: %d", gotValue)
			continue
		}
	}

	if count != len(expected) {
		t.Errorf("want %d elements, but got %d", len(expected), count)
	}
}
