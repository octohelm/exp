package xiter

import (
	"iter"
	"time"
)

func DebounceTime[V any](seq iter.Seq[V], d time.Duration) iter.Seq[V] {
	return func(yield func(V) bool) {
		bufferedSeq := BufferTime(seq, d)

		for buffer := range bufferedSeq {
			if len(buffer) > 0 {
				lastValue := buffer[len(buffer)-1]
				if !yield(lastValue) {
					return
				}
			}
		}
	}
}
