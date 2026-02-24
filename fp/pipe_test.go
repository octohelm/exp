package fp_test

import (
	"fmt"
	"slices"
	"testing"

	. "github.com/octohelm/exp/fp"
	"github.com/octohelm/exp/xiter"
	. "github.com/octohelm/x/testing/v2"
)

func TestCompose(t *testing.T) {
	// 准备基础数据源
	valuesSeq := xiter.Seq(func(yield func(i int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
		}
	})

	t.Run("Pipe2: 基础过滤", func(t *testing.T) {
		Then(t, "经过偶数过滤后，应只保留偶数序列",
			Expect(
				Pipe2(
					valuesSeq,
					Do(xiter.Filter, func(x int) bool { return x%2 == 0 }),
					slices.Collect,
				),
				Equal([]int{0, 2, 4, 6, 8}),
			),
		)
	})

	t.Run("Pipe3: 分块聚合", func(t *testing.T) {
		Then(t, "按3分块并求和后，应得到每组累加值",
			Expect(
				Pipe3(
					valuesSeq,
					Do(xiter.Chunk[int], 3),
					Do(xiter.Map, func(values []int) int {
						return Pipe(
							slices.Values(values),
							Do2(xiter.Reduce, 0, func(ret int, v int) int { return ret + v }),
						)
					}),
					slices.Collect,
				),
				Equal([]int{
					0 + 1 + 2, // 3
					3 + 4 + 5, // 12
					6 + 7 + 8, // 21
					9,         // 9
				}),
			),
		)
	})

	t.Run("Pipe4: 链式变换与过滤", func(t *testing.T) {
		Then(t, "经过平方变换并过滤长度为2的字符串后，应符合预期",
			Expect(
				Pipe4(
					valuesSeq,
					Do(xiter.Filter, func(x int) bool { return x%2 == 0 }),
					Do(xiter.Map, func(x int) string { return fmt.Sprintf("%d", x*x) }),
					Do(xiter.Filter, func(x string) bool { return len(x) == 2 }),
					slices.Collect,
				),
				Equal([]string{"16", "36", "64"}),
			),
		)
	})
}
