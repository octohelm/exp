package xiter

import (
	"iter"
)

// Tap 在不改变元素的前提下，对每个元素执行附加操作。
func Tap[V any](seq iter.Seq[V], tap func(e V)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			tap(v)
			if !yield(v) {
				return
			}
		}
	}
}
