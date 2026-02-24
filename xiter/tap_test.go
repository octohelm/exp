package xiter_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestTap(t *testing.T) {
	src := xiter.Of(0, 1, 2)

	t.Run("副作用执行校验", func(t *testing.T) {
		count := 0
		tapped := xiter.Tap(src, func(e int) {
			count += e
		})

		// 触发迭代
		_ = slices.Collect(tapped)

		Then(t, "累加结果应符合预期",
			Expect(count, Be(cmp.Eq(3))),
		)
	})

	t.Run("断言过程中的副作用定位", func(t *testing.T) {
		observed := make([]int, 0)

		// 组合 Tap 和 Every
		tapped := xiter.Tap(src, func(e int) {
			observed = append(observed, e)
		})

		// 故意在 index 1 失败
		err := cmp.Every(cmp.Lt(1))(tapped)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)

			Then(t, "错误路径定位",
				Expect(string(e.Pointer), Be(cmp.Eq("/1"))),
			)

			Then(t, "副作用记录应止于失败处",
				Expect(observed, Equal([]int{0, 1})),
			)
		}
	})
}
