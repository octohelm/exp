package xiter

import "iter"

func Reduce[T any, R any](seq iter.Seq[T], initial R, reducer func(r R, i T) R) R {
	ret := initial
	seq(func(v T) bool {
		ret = reducer(ret, v)
		return true
	})
	return ret
}
