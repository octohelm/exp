package op

import (
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	testingx "github.com/octohelm/x/testing"
)

func TestChunk(t *testing.T) {
	src := xiter.Seq(func(yield func(int) bool) {
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
