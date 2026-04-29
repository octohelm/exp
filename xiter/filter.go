package xiter

import (
	"iter"
)

// Filter 保留满足条件的元素。
func Filter[V any](seq iter.Seq[V], filter func(e V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for e := range seq {
			if filter(e) && !yield(e) {
				return
			}
		}
	}
}
