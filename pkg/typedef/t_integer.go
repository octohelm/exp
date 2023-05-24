package typedef

import (
	"context"
)

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Integer[T integer]() Type[T] {
	return &integerType[T]{}
}

var _ interface {
	Validator[int]
} = &integerType[int]{}

type integerType[T integer] struct {
}

func (t *integerType[T]) Kind() string {
	return "integer"
}

func (t *integerType[T]) Validate(ctx context.Context, value any, optFns ...OptionFunc) (*T, error) {
	switch x := value.(type) {
	case int:
		v := T(x)
		return &v, nil
	case int16:
		v := T(x)
		return &v, nil
	case int32:
		v := T(x)
		return &v, nil
	case int64:
		v := T(x)
		return &v, nil
	case uint:
		v := T(x)
		return &v, nil
	case uint16:
		v := T(x)
		return &v, nil
	case uint32:
		v := T(x)
		return &v, nil
	case uint64:
		v := T(x)
		return &v, nil
	}
	return nil, InvalidType
}
