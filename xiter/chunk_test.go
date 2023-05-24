package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestChunk(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
		}
	})

	chunked := Chunk(src, 3)

	testingx.Expect(t, slices.Collect(chunked), testingx.Equal([][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}))
}
