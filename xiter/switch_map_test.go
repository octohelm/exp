package xiter_test

import (
	"errors"
	"iter"
	"slices"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/exp/xiter"
)

func TestSwitchMap(t *testing.T) {
	src := xiter.Of(0, 1, 2, 3, 4)

	mapped := xiter.SwitchMap(src, func(x int) iter.Seq[int] {
		if x%2 == 0 {
			return xiter.Of(x * 2)
		}
		return xiter.Of(x * x)
	})

	t.Run("基础映射逻辑校验", func(t *testing.T) {
		values := slices.Collect(mapped)
		Then(t, "转换后的值应符合条件分支逻辑",
			Expect(values, Equal([]int{0, 1, 4, 9, 8})),
		)
	})

	t.Run("展平后的错误定位", func(t *testing.T) {
		err := cmp.Every(cmp.Lt(9))(mapped)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)

			Then(t, "报错路径应指向展平后序列的全局索引 /3",
				Expect(string(e.Pointer), Be(cmp.Eq("/3"))),
			)
		}
	})

	t.Run("1 对多映射测试", func(t *testing.T) {
		multiMapped := xiter.SwitchMap(xiter.Of(1, 2), func(x int) iter.Seq[int] {
			return xiter.Of(x, x+10)
		})

		err := cmp.Every(cmp.Lt(10))(multiMapped)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)
			Then(t, "1 对多映射时的全局索引应正确",
				Expect(string(e.Pointer), Be(cmp.Eq("/1"))),
			)
		}
	})
}
