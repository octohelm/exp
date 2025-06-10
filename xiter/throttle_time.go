package xiter

import (
	"iter"
	"time"
)

func ThrottleTime[V any](seq iter.Seq[V], d time.Duration) iter.Seq[V] {
	return func(yield func(V) bool) {
		timer := time.NewTimer(d)
		defer timer.Stop()

		canEmit := true

		values := make(chan V)
		done := make(chan struct{})

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
					timer.Stop()
					return
				}
				if canEmit {
					if !yield(val) {
						return
					}
					canEmit = false
					timer.Reset(d)
				}
			case <-timer.C:
				canEmit = true
				timer.Reset(d)
			case <-done:
				return
			}
		}
	}
}
