package xchan

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func FuzzSubject(f *testing.F) {
	for range 10 {
		f.Add(rand.Intn(1000), rand.Intn(100))
	}

	f.Fuzz(func(t *testing.T, n int, workerNumber int) {
		if workerNumber <= 0 || n <= 0 {
			t.Skip()
		}

		t.Run(fmt.Sprintf("worker %d with %d values", workerNumber, n), func(t *testing.T) {
			ret := &Subject[int]{}
			src := &Subject[int]{}

			wg := &sync.WaitGroup{}

			// 消费者：每个 worker 期望从 src 观察并转发 n 个值到 ret
			for range workerNumber {
				wg.Go(func() {
					count := 0

					for x := range Observe(t.Context(), src.Observe()) {
						count++

						ret.Send(x)

						if count >= n {
							return
						}
					}
				})
			}

			// 生产者：延迟发送数据
			wg.Go(func() {
				// 模拟异步生产
				time.Sleep(time.Duration(n) * time.Microsecond)

				for i := range n {
					src.Send(i)
				}
			})

			go func() {
				wg.Wait()

				ret.CancelCause(nil)
			}()

			values := slices.Collect(
				Observe(t.Context(), ret.Observe()),
			)

			Then(t, "收到的总数据量应等于 worker 数乘以每个 worker 处理的量",
				Expect(len(values), Equal(n*workerNumber)),
			)

			Then(t, "结果集不应为 nil",
				Expect(values, Be(cmp.NotNil[[]int]())),
			)
		})
	})
}
