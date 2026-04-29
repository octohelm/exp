package xiter

import (
	"iter"
	"time"
)

// DebounceTime 在每个时间窗口内只输出最后一个元素。
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
