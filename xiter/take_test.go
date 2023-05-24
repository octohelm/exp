package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestTake(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	values := Take(src, 2)

	testingx.Expect(t, slices.Collect(values), testingx.Equal([]int{
		0, 1,
	}))
}
