package xiter

import (
	"iter"
	"slices"
)

// Of 从给定值创建一个 iter.Seq。
func Of[V any](values ...V) iter.Seq[V] {
	return slices.Values(values)
}

// Seq 返回传入的 iter.Seq，便于在链式调用中显式表达来源。
func Seq[V any](seq iter.Seq[V]) iter.Seq[V] {
	return seq
}
