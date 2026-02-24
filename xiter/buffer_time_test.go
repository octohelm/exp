package xiter

import (
	"slices"
	"testing"
	"time"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestBufferTime(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		src := Seq(func(yield func(int) bool) {
			for i := range 5 {
				if !yield(i) {
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
		})

		Then(t, "在 500ms 窗口内应收集所有元素",
			Expect(slices.Collect(BufferTime(src, 500*time.Millisecond)),
				Equal([][]int{{0, 1, 2, 3, 4}}),
			),
		)
	})

	t.Run("LongPause", func(t *testing.T) {
		src := Seq(func(yield func(int) bool) {
			yield(0)
			time.Sleep(50 * time.Millisecond)
			yield(1)
			time.Sleep(50 * time.Millisecond)
			yield(2)
			time.Sleep(200 * time.Millisecond) // 超过 100ms 窗口触发 flush
			yield(3)
			time.Sleep(50 * time.Millisecond)
			yield(4)
		})

		Then(t, "长间隔应导致分批次输出",
			Expect(slices.Collect(BufferTime(src, 100*time.Millisecond)),
				Equal([][]int{
					{0, 1, 2},
					{3, 4},
				}),
			),
		)
	})

	t.Run("BoundaryConditions", func(t *testing.T) {
		t.Run("EmptySource", func(t *testing.T) {
			src := Seq(func(yield func(int) bool) {
				// empty
			})

			Then(t, "空数据源应返回空切片",
				Expect(slices.Collect(BufferTime(src, 100*time.Millisecond)),
					Be(cmp.Nil[[][]int]()),
				),
			)
		})

		t.Run("ImmediateYield", func(t *testing.T) {
			src := Seq(func(yield func(int) bool) {
				yield(1)
			})

			Then(t, "单个元素且立即结束应返回单批次",
				Expect(slices.Collect(BufferTime(src, 10*time.Millisecond)),
					Equal([][]int{{1}}),
				),
			)
		})
	})
}
