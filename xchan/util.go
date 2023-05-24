package xchan

import (
	"context"
	"iter"

	"github.com/octohelm/exp/xiter"
)

func Values[T any](ctx context.Context, o Observable[T]) iter.Seq[T] {
	return xiter.RecvContext(ctx, o.Value())
}
