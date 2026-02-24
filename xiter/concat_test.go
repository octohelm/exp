package xiter_test

import (
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestConcat(t *testing.T) {
	t.Run("合并多个序列", func(t *testing.T) {
		seq0 := xiter.Of(0, 2, 4)
		seq1 := xiter.Of(1, 3, 5)

		seq := xiter.Concat(seq0, seq1)

		Then(t, "序列应按顺序连接",
			Expect(slices.Collect(seq), Equal([]int{
				0, 2, 4, 1, 3, 5,
			})),
		)
	})

	t.Run("包含空序列的合并", func(t *testing.T) {
		seq0 := xiter.Of[int]()
		seq1 := xiter.Of(1, 2)
		seq2 := xiter.Of[int]()

		seq := xiter.Concat(seq0, seq1, seq2)

		Then(t, "空序列不应影响合并结果",
			Expect(slices.Collect(seq), Equal([]int{1, 2})),
		)
	})

	t.Run("边界条件", func(t *testing.T) {
		t.Run("全部为空", func(t *testing.T) {
			seq := xiter.Concat[int]()

			Then(t, "不提供输入时应返回空迭代器",
				Expect(slices.Collect(seq), Be(cmp.Len[[]int](0))),
			)
		})
	})
}
