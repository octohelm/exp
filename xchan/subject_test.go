package xchan

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"testing"

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
		if n > 64 {
			n = 64
		}
		if workerNumber > 16 {
			workerNumber = 16
		}

		t.Run(fmt.Sprintf("%d 个协程处理 %d 个值", workerNumber, n), func(t *testing.T) {
			ret := &Subject[int]{}
			src := &Subject[int]{}
			retObserver := ret.Observe()

			wg := &sync.WaitGroup{}
			ready := &sync.WaitGroup{}
			ready.Add(workerNumber)

			// 消费者：每个 worker 期望从 src 观察并转发 n 个值到 ret
			for range workerNumber {
				wg.Go(func() {
					count := 0
					ob := src.Observe()
					ready.Done()

					for x := range Observe(t.Context(), ob) {
						count++

						ret.Send(x)

						if count >= n {
							return
						}
					}
				})
			}

			ready.Wait()

			// 生产者：延迟发送数据
			wg.Go(func() {
				for i := range n {
					src.Send(i)
				}

				src.CancelCause(nil)
			})

			go func() {
				wg.Wait()

				ret.CancelCause(nil)
			}()

			values := slices.Collect(
				Observe(t.Context(), retObserver),
			)

			Then(t, "收到的总数据量应等于协程数乘以每个协程处理的量",
				Expect(len(values), Equal(n*workerNumber)),
			)

			Then(t, "结果集不应为 nil",
				Expect(values, Be(cmp.NotNil[[]int]())),
			)
		})
	})
}
