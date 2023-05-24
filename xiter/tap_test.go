package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestTap(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	count := 0

	_ = slices.Collect(
		Tap(src, func(e int) {
			count += e
		}),
	)

	testingx.Expect(t, count, testingx.Be(0+1+2))
}
