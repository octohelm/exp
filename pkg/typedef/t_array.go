package typedef

import (
	"context"
)

func Array[T any](item Type[T]) Type[[]T] {
	return &arrayType[T]{
		Items: item,
	}
}

var _ interface {
	Validator[[]int]
} = &arrayType[int]{}

type arrayType[T any] struct {
	Items Type[T] `json:"items"`
}

func (t *arrayType[T]) Kind() string {
	return "array"
}

func (t *arrayType[T]) Validate(ctx context.Context, value any, optFns ...OptionFunc) (*[]T, error) {
	switch x := value.(type) {
	case []T:
		return &x, nil
	}
	return nil, InvalidType
}

func (t *arrayType[T]) Coerce(ctx context.Context, value any, optFns ...OptionFunc) (*[]T, error) {
	switch x := value.(type) {
	case []T:
		vv := make([]T, len(x))
		for i := range x {
			vv[i] = x[i]
		}
		return &vv, nil
	}
	return nil, InvalidType
}

func (t *arrayType[T]) EntryIter(ctx context.Context, v any, optFns ...OptionFunc) <-chan Entry {
	switch x := v.(type) {
	case []T:
		ch := make(chan Entry)

		go func() {
			defer close(ch)

			select {
			case <-ctx.Done():
				return
			default:
				for i := range x {
					ch <- Entry{
						Type:  t.Items,
						Key:   i,
						Value: x[i],
					}
				}
			}
		}()

		return ch
	}

	return nil
}
