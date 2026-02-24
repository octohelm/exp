package xiter_test

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/octohelm/exp/xiter"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestMap(t *testing.T) {
	src := xiter.Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	t.Run("基础数值转换", func(t *testing.T) {
		mappedSeq := xiter.Map(src, func(x int) int {
			return x * x
		})

		Then(t, "结果应为平方值",
			Expect(slices.Collect(mappedSeq), Equal([]int{0, 1, 4})),
		)
	})

	t.Run("类型转换与深度断言", func(t *testing.T) {
		stringSeq := xiter.Map(src, func(x int) string {
			return "val_" + strconv.Itoa(x)
		})

		Then(t, "所有转换后的字符串都应符合格式",
			Expect(stringSeq, Be(cmp.Every(func(s string) error {
				if len(s) > 4 && s[:4] == "val_" {
					return nil
				}
				return fmt.Errorf("格式错误")
			}))),
		)
	})
}
