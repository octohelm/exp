package typedef

import (
	"context"
)

func Optional[T any](t Type[T]) Type[T] {
	return &optionalType[T]{
		Type: t,
	}
}

var _ interface {
	Validator[any]
	Refiner[any]
} = &optionalType[any]{}

type optionalType[T any] struct {
	Type[T]
}

func (o *optionalType[T]) Kind() string {
	return o.Type.Kind()
}

func (o *optionalType[T]) Validate(ctx context.Context, value any, optFns ...OptionFunc) (*T, error) {
	if value == nil {
		return nil, nil
	}
	return Validate(ctx, o.Type, value, optFns...)
}

func (o *optionalType[T]) Refine(ctx context.Context, value any, optFns ...OptionFunc) (*T, error) {
	if value == nil {
		return nil, nil
	}
	return Refine(ctx, o.Type, value, optFns...)
}
