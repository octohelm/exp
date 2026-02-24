package xiter_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestFilter(t *testing.T) {
	t.Run("过滤偶数", func(t *testing.T) {
		src := xiter.Seq(func(yield func(int) bool) {
			for i := range 10 {
				if !yield(i) {
					return
				}
			}
		})

		filteredSeq := xiter.Filter(src, func(x int) bool {
			return x%2 == 0
		})

		Then(t, "结果集应只包含偶数且顺序正确",
			Expect(slices.Collect(filteredSeq), Equal([]int{0, 2, 4, 6, 8})),
			Expect(filteredSeq, Be(cmp.Every(func(v int) error {
				if v%2 != 0 {
					return fmt.Errorf("期望偶数，实际得到 %d", v)
				}
				return nil
			}))),
		)
	})

	t.Run("边界测试", func(t *testing.T) {
		t.Run("全过滤", func(t *testing.T) {
			src := xiter.Of(1, 3, 5)
			filtered := xiter.Filter(src, func(x int) bool { return x%2 == 0 })

			Then(t, "不匹配时应返回空切片",
				Expect(slices.Collect(filtered), Be(cmp.Len[[]int](0))),
			)
		})

		t.Run("空输入", func(t *testing.T) {
			src := xiter.Of[int]()
			filtered := xiter.Filter(src, func(x int) bool { return true })

			Then(t, "空迭代器过滤后仍为空",
				Expect(slices.Collect(filtered), Be(cmp.Nil[[]int]())),
			)
		})
	})
}
