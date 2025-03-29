package xiter

import "iter"

func Concat[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for x := range seq {
				if !yield(x) {
					return
				}
			}
		}
	}
}
