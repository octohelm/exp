package typedef

import (
	"context"
)

type number interface {
	float32 | float64 | integer
}

func Number[T number]() Type[T] {
	return &numberType[T]{}
}

type numberType[T number] struct {
}

var _ interface {
	Validator[float64]
} = &numberType[float64]{}

func (t *numberType[T]) Kind() string {
	return "number"
}

func (t *numberType[T]) Validate(ctx context.Context, value any, optFns ...OptionFunc) (*T, error) {
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
	case float64:
		v := T(x)
		return &v, nil
	case float32:
		v := T(x)
		return &v, nil
	}
	return nil, InvalidType
}
