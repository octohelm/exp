package xiter_test

import (
	"slices"
	"testing"

	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/exp/xiter"
)

func TestMerge(t *testing.T) {
	seq0 := xiter.Of(0, 2, 4)
	seq1 := xiter.Of(1, 3, 5)

	merged := xiter.Merge(seq1, seq0)

	t.Run("基础合并结果校验", func(t *testing.T) {
		values := slices.Sorted(merged)

		Then(t, "合并后的有序集合应包含所有元素",
			Expect(values, Equal([]int{0, 1, 2, 3, 4, 5})),
		)
	})

	t.Run("合并流的原始顺序校验", func(t *testing.T) {
		values := slices.Collect(merged)
		slices.Sort(values)

		Then(t, "并发合并时应保证元素完整且不重复",
			Expect(values, Equal([]int{0, 1, 2, 3, 4, 5})),
		)
	})
}
