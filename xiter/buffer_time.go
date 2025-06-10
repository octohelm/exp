package xiter

import (
	"iter"
	"time"
)

func BufferTime[V any](seq iter.Seq[V], d time.Duration) iter.Seq[[]V] {
	return func(yield func([]V) bool) {
		buffer := make([]V, 0)

		timer := time.NewTimer(d)
		defer timer.Stop()

		done := make(chan struct{})
		values := make(chan V)

		go func() {
			defer close(values)
			for v := range seq {
				values <- v
			}
			close(done)
		}()

		for {
			select {
			case val, ok := <-values:
				if !ok {
					if len(buffer) > 0 {
						if !yield(buffer) {
							return
						}
					}
					return
				}
				buffer = append(buffer, val)
				timer.Reset(d)
			case <-timer.C:

				if len(buffer) > 0 {
					if !yield(buffer) {
						return
					}

					buffer = make([]V, 0)
				}

				timer.Reset(d)

			case <-done:
				select {
				case <-timer.C:
					if len(buffer) > 0 {
						if !yield(buffer) {
							return
						}
					}
				default:
					if len(buffer) > 0 {
						if !yield(buffer) {
							return
						}
					}
				}
				return
			}
		}
	}
}
