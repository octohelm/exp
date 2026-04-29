package xiter_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/exp/xiter"
)

func TestSkip(t *testing.T) {
	src := xiter.Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	t.Run("基础跳过逻辑", func(t *testing.T) {
		values := xiter.Skip(src, 1) // 预期 [1, 2]

		Then(t, "应跳过第一个元素",
			Expect(slices.Collect(values), Equal([]int{1, 2})),
		)
	})

	t.Run("跳过后路径校验", func(t *testing.T) {
		skipped := xiter.Skip(src, 1)

		err := cmp.Every(cmp.Lt(2))(skipped)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)
			Then(t, "报错路径应基于 Skip 后的新序列索引",
				Expect(string(e.Pointer), Be(cmp.Eq("/1"))),
			)
		}
	})

	t.Run("跳过超过长度的情况", func(t *testing.T) {
		values := xiter.Skip(src, 10)

		Then(t, "结果应为空序列",
			Expect(slices.Collect(values), Be(cmp.Len[[]int](0))),
		)
	})
}
