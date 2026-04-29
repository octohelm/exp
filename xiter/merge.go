package xiter

import (
	"iter"
	"sync"
)

// Merge 并发消费多个序列，并按到达顺序合并输出。
func Merge[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	if len(seqs) == 0 {
		return func(yield func(T) bool) {
		}
	}

	if len(seqs) == 1 {
		return seqs[0]
	}

	return func(yield func(T) bool) {
		chValue := make(chan T)
		chDone := make(chan struct{})

		wg := &sync.WaitGroup{}

		for _, seq := range seqs {
			wg.Go(func() {
				for v := range seq {
					chValue <- v
				}
			})
		}

		go func() {
			wg.Wait()
			close(chDone)
		}()

		for {
			select {
			case <-chDone:
				return
			case v := <-chValue:
				if !yield(v) {
					return
				}

			}
		}
	}
}
