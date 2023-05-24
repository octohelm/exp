package xiter

import (
	"iter"
	"slices"
)

func Seq[V any](seq iter.Seq[V]) iter.Seq[V] {
	return seq
}

func Of[V any](values ...V) iter.Seq[V] {
	return slices.Values(values)
}
