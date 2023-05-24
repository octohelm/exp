package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestSkip(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	values := Skip(src, 1)

	testingx.Expect(t, slices.Collect(values), testingx.Equal([]int{
		1, 2,
	}))
}
