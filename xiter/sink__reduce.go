package xiter

import (
	"iter"
)

// Reduce 从初始值开始累计序列中的元素。
func Reduce[T any, R any](seq iter.Seq[T], initial R, reducer func(r R, i T) R) R {
	ret := initial
	seq(func(v T) bool {
		ret = reducer(ret, v)
		return true
	})
	return ret
}
