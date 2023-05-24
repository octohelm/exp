package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestMap(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	mapped := Map(src, func(x int) int {
		return x * x
	})

	testingx.Expect(t, slices.Collect(mapped), testingx.Equal([]int{
		0, 1, 4,
	}))
}
