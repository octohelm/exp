package xiter

import (
	"iter"
	"slices"
)

func Of[V any](values ...V) iter.Seq[V] {
	return slices.Values(values)
}

func Seq[V any](seq iter.Seq[V]) iter.Seq[V] {
	return seq
}
