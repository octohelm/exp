package xchan

import (
	"context"
	"iter"
)

func Subscribe[T any](ctx context.Context, source Observable[T], subscriber Subscriber[T]) {
	src := source.Observe()
	defer src.CancelCause(nil)
	defer subscriber.CancelCause(nil)

	for {
		select {
		case <-ctx.Done():
			return
		case <-subscriber.Done():
			return
		case v, ok := <-src.Value():
			if !ok {
				return
			}
			subscriber.Send(v)
		}
	}
}

func Values[T any](ctx context.Context, o ValueObservable[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
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

func Observe[T any](ctx context.Context, o Observer[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		defer o.CancelCause(ctx.Err())

		Values(ctx, o)(yield)
	}
}
