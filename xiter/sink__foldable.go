package xiter

import "iter"

func Fold[T any](seq iter.Seq[T], reducer func(T, T) T) T {
	var prev T
	r := func(v1, v2 T) T { return v2 }
	seq(func(v T) bool {
		prev = r(prev, v)
		r = reducer
		return true
	})
	return prev
}

func Reduce[T any, R any](seq iter.Seq[T], initial R, reducer func(r R, i T) R) R {
	ret := initial
	seq(func(v T) bool {
		ret = reducer(ret, v)
		return true
	})
	return ret
}

func Count[T any](seq iter.Seq[T]) int {
	i := 0
	for _ = range seq {
		i++
	}
	return i
}
