package typedef

import (
	"context"
)

func String() Type[string] {
	return &stringType{}
}

var _ interface {
	Validator[string]
} = &stringType{}

type stringType struct {
}

func (t *stringType) Kind() string {
	return "string"
}

func (t *stringType) Validate(ctx context.Context, v any, optFns ...OptionFunc) (*string, error) {
	switch x := v.(type) {
	case string:
		return &x, nil
	}
	return nil, InvalidType
}
