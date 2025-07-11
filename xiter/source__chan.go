package xiter

import (
	"context"
	"iter"
)

func Recv[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

func RecvContext[T any](ctx context.Context, value <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-value:
				if !ok || !yield(v) {
					return
				}
			}
		}
	}
}
