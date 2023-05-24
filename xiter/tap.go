package xiter

import "iter"

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
