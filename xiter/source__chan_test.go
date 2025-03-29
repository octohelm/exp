package xiter

import (
	"context"
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestChan(t *testing.T) {
	c := make(chan int)
	defer close(c)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for i, v := range []int{1, 2, 3, 4, 5} {
			c <- v

			if i == 2 {
				cancel()
				return
			}
		}
	}()

	values := slices.Collect(RecvContext(ctx, c))
	testingx.Expect(t, values, testingx.Equal([]int{
		1, 2, 3,
	}))
}
