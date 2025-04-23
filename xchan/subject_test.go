package xchan

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"testing"
	"time"

	testingx "github.com/octohelm/x/testing"
)

func FuzzSubject(f *testing.F) {
	// f.Add(2, 2)

	for range 10 {
		f.Add(rand.Intn(1000), rand.Intn(100))
	}

	f.Fuzz(func(t *testing.T, n int, workerNumber int) {
		t.Run(fmt.Sprintf("worker %d with %d values", workerNumber, n), func(t *testing.T) {
			ret := &Subject[int]{}
			src := &Subject[int]{}

			wg := &WaitGroup{}
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

			wg.Go(func() {
				// defer send
				time.Sleep(time.Duration(n) * time.Millisecond)

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

			testingx.Expect(t, len(values), testingx.Be(n*workerNumber))
		})
	})
}

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Go(x func()) {
	w.Add(1)

	go func() {
		defer w.Done()

		x()
	}()
}
