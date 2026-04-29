package xiter

import (
	"context"
	"iter"
)

// Recv 将只读 channel 适配为 iter.Seq。
func Recv[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

// RecvContext 在上下文可取消的前提下消费 channel 并产出 iter.Seq。
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
