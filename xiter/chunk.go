package xiter

import (
	"iter"
)

func Chunk[V any](seq iter.Seq[V], n int) iter.Seq[[]V] {
	if n < 1 {
		panic("cannot be less than 1")
	}

	return func(yield func([]V) bool) {
		count := 0
		chunk := make([]V, 0, n)

		emit := func() bool {
			if !yield(chunk) {
				return false
			}

			// reset if seq not break
			count = 0
			chunk = make([]V, 0, n)
			return true
		}

		for e := range seq {
			chunk = append(chunk, e)
			count++

			if count == n {
				if !emit() {
					return
				}
			}
		}

		if len(chunk) > 0 {
			if !emit() {
				return
			}
		}
	}
}
