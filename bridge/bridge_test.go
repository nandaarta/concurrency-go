package bridge_test

import (
	"slices"
	"testing"

	"github.com/nandaarta/concurrency-go/bridge"
)

func TestBridge(t *testing.T) {
	t.Parallel()

	var (
		values = []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		exp    = []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	done := make(chan any)
	defer close(done)

	bridgeRes := bridge.Bridge(done, bridge.Generate(values))

	res := make([]any, 0)
	for value := range bridgeRes {
		res = append(res, value)
	}

	if !slices.Equal(exp, res) {
		t.Errorf("want %+v, but actual result is %+v", exp, res)
	}
}
