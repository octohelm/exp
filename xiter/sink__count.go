package xiter

import "iter"

func Count[T any](seq iter.Seq[T]) int {
	i := 0
	for _ = range seq {
		i++
	}
	return i
}
