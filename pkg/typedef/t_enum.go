package typedef

import (
	"context"
)

func Enum[T comparable](values []T) Type[T] {
	return &enumType[T]{
		Enum: values,
	}
}

var _ interface {
	Validator[any]
} = &enumType[any]{}

type enumType[T comparable] struct {
	Enum []T `json:"enum"`
}

func (t *enumType[T]) Kind() string {
	return "enum"
}

func (t *enumType[T]) Validate(ctx context.Context, v any, optFns ...OptionFunc) (*T, error) {
	for i := range t.Enum {
		if e := t.Enum[i]; e == v {
			return &e, nil
		}
	}
	return nil, InvalidType
}
