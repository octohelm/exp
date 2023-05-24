package typedef

import (
	"context"
)

func Boolean() Type[bool] {
	return &booleanType{}
}

var _ interface {
	Validator[bool]
} = &booleanType{}

type booleanType struct {
}

func (t *booleanType) Kind() string {
	return "boolean"
}

func (t *booleanType) Validate(ctx context.Context, v any, optFns ...OptionFunc) (*bool, error) {
	switch x := v.(type) {
	case bool:
		return &x, nil
	}
	return nil, InvalidType
}
