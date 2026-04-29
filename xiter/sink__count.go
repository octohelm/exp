package xiter

import (
	"iter"
)

// Count 统计序列中的元素数量。
func Count[T any](seq iter.Seq[T]) int {
	i := 0
	for range seq {
		i++
	}
	return i
}
