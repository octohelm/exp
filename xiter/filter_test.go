package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestFilter(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
		}
	})

	filtered := Filter(src, func(x int) bool {
		return x%2 == 0
	})

	testingx.Expect(t, slices.Collect(filtered), testingx.Equal([]int{
		0, 2, 4, 6, 8,
	}))
}
