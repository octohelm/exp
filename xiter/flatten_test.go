package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestFlatten(t *testing.T) {
	src1 := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	src2 := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	flattened := Flatten(Of(src1, src2))

	testingx.Expect(t, slices.Collect(flattened), testingx.Equal([]int{
		0, 1, 2,
		0, 1, 2,
	}))
}
