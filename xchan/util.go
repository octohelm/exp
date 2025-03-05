package xchan

import (
	"context"
	"iter"

	"github.com/octohelm/exp/xiter"
)

func Observe[T any](ctx context.Context, o Observer[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		defer o.CancelCause(ctx.Err())

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-o.Value():
				if !ok || !yield(v) {
					return
				}
			}
		}
	}
}

func Values[T any](ctx context.Context, o Observable[T]) iter.Seq[T] {
	return xiter.RecvContext(ctx, o.Value())
}
