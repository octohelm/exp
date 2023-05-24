package xiter

import (
	testingx "github.com/octohelm/x/testing"
	"iter"
	"slices"
	"testing"
	"time"
)

func TestSwitchMap(t *testing.T) {
	src := Of(0, 1, 2, 3, 4)

	mapped := SwitchMap(src, func(x int) iter.Seq[int] {
		if x%2 == 0 {
			return Of(x * 2)
		}
		return Of(x * x)
	})

	values := slices.Collect(mapped)
	testingx.Expect(t, values, testingx.Equal([]int{
		0 * 2,
		1 * 1,
		2 * 2,
		3 * 3,
		4 * 2,
	}))

	time.Sleep(1 * time.Second)
}
