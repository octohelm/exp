package xiter

import (
	"iter"
)

// Fold 使用 reducer 把序列折叠为单个值。
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
