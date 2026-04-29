package xchan

import (
	"context"
	"math/rand"
	"slices"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func FuzzNotifiableObserver(f *testing.F) {
	for range 10 {
		f.Add(rand.Intn(10000))
	}

	f.Fuzz(func(t *testing.T, n int) {
		if n < 0 {
			t.Skip()
		}
		if n > 256 {
			n = 256
		}

		x := NewNotifiableObserver[int]()

		// 模拟多协程监听 Done 信号
		for range 10 {
			go func() {
				<-x.Done()
			}()
		}

		// 异步发送数据并在结束时取消
		go func() {
			for i := range n {
				x.Send(i)
			}
			x.CancelCause(nil)
		}()

		// 消费数据
		values := slices.Collect(Values(context.Background(), x))

		Then(t, "收到的数据总量应与发送量一致",
			Expect(len(values), Equal(n)),
		)

		Then(t, "当 Observer 取消后，通道应当关闭",
			Expect(func() bool {
				_, ok := <-x.Done()
				return ok
			}(), Be(cmp.Eq(false))),

			Expect(func() bool {
				_, ok := <-x.Value()
				return ok
			}(), Be(cmp.Eq(false))),
		)
	})
}
