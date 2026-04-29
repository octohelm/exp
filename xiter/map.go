package xiter

import (
	"iter"
)

// Map 将序列中的每个元素映射为新的值。
func Map[I any, O any](seq iter.Seq[I], m func(e I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for e := range seq {
			if !yield(m(e)) {
				return
			}
		}
	}
}
