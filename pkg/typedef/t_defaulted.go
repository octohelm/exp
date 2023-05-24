package typedef

import (
	"context"
)

func Defaulted[T any](v T) func(t Type[T]) Type[T] {
	return func(t Type[T]) Type[T] {
		return &defaultedType[T]{
			Type:         t,
			DefaultValue: v,
		}
	}
}

var _ interface {
	Coercer
} = &defaultedType[string]{}

type defaultedType[T any] struct {
	Type[T]
	DefaultValue T
}

func (o *defaultedType[T]) Kind() string {
	return o.Type.Kind()
}

func (o *defaultedType[T]) Coerce(ctx context.Context, value any, optFns ...OptionFunc) (any, error) {
	if value == nil {
		// FIXME handle empty value
		return o.DefaultValue, nil
	}
	switch x := value.(type) {
	case T:
		return x, nil
	}
	return nil, nil
}
