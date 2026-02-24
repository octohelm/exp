package xiter_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestTake(t *testing.T) {
	src := xiter.Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	t.Run("基础截取逻辑", func(t *testing.T) {
		values := xiter.Take(src, 2) // 预期 [0, 1]

		Then(t, "应仅产出前两个元素",
			Expect(slices.Collect(values), Equal([]int{0, 1})),
		)
	})

	t.Run("截取后的索引校验", func(t *testing.T) {
		taken := xiter.Take(src, 2)

		err := cmp.Every(cmp.Lt(1))(taken)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)

			Then(t, "报错路径应正确指向截取流内部的索引 /1",
				Expect(string(e.Pointer), Be(cmp.Eq("/1"))),
			)
		}
	})

	t.Run("截取 0 或负数", func(t *testing.T) {
		Then(t, "应返回空序列",
			Expect(slices.Collect(xiter.Take(src, 0)), Be(cmp.Len[[]int](0))),
			Expect(slices.Collect(xiter.Take(src, -1)), Be(cmp.Len[[]int](0))),
		)
	})
}
