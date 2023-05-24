package op

import (
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	testingx "github.com/octohelm/x/testing"
)

func TestFlatten(t *testing.T) {
	src1 := xiter.Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	src2 := xiter.Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	flattened := Flatten(xiter.Of(src1, src2))

	testingx.Expect(t, slices.Collect(flattened), testingx.Equal([]int{
		0, 1, 2,
		0, 1, 2,
	}))
}
