package xiter_test

import (
	"context"
	"errors"
	"slices"
	"testing"

	"github.com/octohelm/exp/xiter"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestRecvContext(t *testing.T) {
	t.Run("Context 中止校验", func(t *testing.T) {
		c := make(chan int)
		// 注意：不要在这里直接 defer close(c)，防止异步发送者往已关闭的 chan 发送

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			defer close(c)
			for _, v := range []int{1, 2, 3, 4, 5} {
				select {
				case <-ctx.Done():
					return
				case c <- v:
					if v == 3 {
						cancel() // 发送到 3 之后取消
						return
					}
				}
			}
		}()

		values := slices.Collect(xiter.RecvContext(ctx, c))

		Then(t, "结果应在 context 取消处截断",
			Expect(values, Equal([]int{1, 2, 3})),
		)
	})

	t.Run("异步流的 Pointer 定位", func(t *testing.T) {
		c := make(chan int, 5)
		for _, v := range []int{10, 20, 30, 40} {
			c <- v
		}
		close(c)

		// 异步流转换为 Seq 并校验
		seq := xiter.RecvContext(context.Background(), c)

		// 校验所有异步收到的值都要小于 30
		err := cmp.Every(cmp.Lt(30))(seq)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)

			Then(t, "异步流产生的错误依然能精准定位",
				// 30 是第 3 个元素 (index 2)
				Expect(string(e.Pointer), Be(cmp.Eq("/2"))),
			)
		}
	})
}
