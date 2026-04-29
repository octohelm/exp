package xchan

import (
	"context"
	"iter"
)

// Subscribe 持续把 source 的值转发到 subscriber，直到任一方结束。
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

// Values 将 ValueObservable 适配为 iter.Seq。
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

// Observe 将 Observer 适配为 iter.Seq，并在结束时传递上下文取消原因。
func Observe[T any](ctx context.Context, o Observer[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		defer o.CancelCause(ctx.Err())

		Values(ctx, o)(yield)
	}
}
