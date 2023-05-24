package op

import (
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	testingx "github.com/octohelm/x/testing"
)

func TestTap(t *testing.T) {
	src := xiter.Seq(func(yield func(int) bool) {
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
