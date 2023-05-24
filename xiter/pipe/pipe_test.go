package pipe_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	xiterop "github.com/octohelm/exp/xiter/op"
	testingx "github.com/octohelm/x/testing"

	. "github.com/octohelm/exp/xiter/pipe"
)

func TestCompose(t *testing.T) {
	valuesSeq := xiter.Seq(func(yield func(i int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
		}
	})

	t.Run("one", func(t *testing.T) {
		values := Pipe2(
			valuesSeq,
			Do(xiterop.Filter, func(x int) bool { return x%2 == 0 }),
			slices.Collect,
		)

		testingx.Expect(t, values, testingx.Equal([]int{0, 2, 4, 6, 8}))
	})

	t.Run("two", func(t *testing.T) {
		values := Pipe3(
			valuesSeq,
			Do(xiterop.Chunk[int], 3),
			Do(xiterop.Map, func(values []int) int {
				return Pipe(
					slices.Values(values),
					Do2(xiter.Reduce, 0, func(ret int, v int) int { return ret + v }),
				)
			}),
			slices.Collect,
		)

		testingx.Expect(t, values, testingx.Equal([]int{
			0 + 1 + 2,
			3 + 4 + 5,
			6 + 7 + 8,
			9,
		}))
	})

	t.Run("three", func(t *testing.T) {
		values := Pipe4(
			valuesSeq,
			Do(xiterop.Filter, func(x int) bool { return x%2 == 0 }),
			Do(xiterop.Map, func(x int) string { return fmt.Sprintf("%d", x*x) }),
			Do(xiterop.Filter, func(x string) bool { return len(x) == 2 }),
			slices.Collect,
		)

		testingx.Expect(t, values, testingx.Equal([]string{
			"16", "36", "64",
		}))
	})
}
