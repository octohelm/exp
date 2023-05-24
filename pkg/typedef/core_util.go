package typedef

import (
	"context"
)

func Refine[T any](ctx context.Context, t Type[T], value any, optFns ...OptionFunc) (*T, error) {
	if v, ok := t.(Refiner[T]); ok {
		return v.Refine(ctx, value, optFns...)
	}
	switch x := value.(type) {
	case T:
		return &x, nil
	}
	return nil, InvalidType
}

func Validate[T any](ctx context.Context, t Type[T], value any, optFns ...OptionFunc) (*T, error) {
	o := newOption(optFns...)

	if o.Coerce {
		if coercer, ok := t.(Coercer); ok {
			ret, err := coercer.Coerce(ctx, value, optFns...)
			if err != nil {
				return nil, err
			}
			if ret != nil {
				value = ret
			}
		}
	}

	if v, ok := t.(Validator[T]); ok {
		ret, err := v.Validate(ctx, value, optFns...)
		if err != nil {
			return nil, err
		}
		if ret != nil {
			value = *ret
		}
	}

	if iter, ok := Iterate(ctx, t, value, optFns...); ok {
		var finalErr error

		for e := range iter {
			options := optFns[:]
			if e.Key != nil {
				options = append(options, WithPath(e.Key), WithBranch(e.Value))
			}

			valueValid, err := Validate[any](ctx, e.Type, e.Value, options...)
			if err != nil {
				finalErr = Wrap(finalErr, WrapFailure(err, e.Type, e.Value, options...))
				continue
			}

			if valueValid != nil {
				SetTo(value, e.Key, *valueValid)
			}
		}

		if finalErr != nil {
			return nil, finalErr
		}
	}

	if r, ok := t.(Refiner[T]); ok {
		return r.Refine(ctx, value, optFns...)
	}

	if value != nil {
		switch x := value.(type) {
		case *T:
			return x, nil
		case T:
			return &x, nil
		}
		return nil, InvalidType
	}

	return nil, nil
}

type Entry struct {
	Key   any
	Value any
	Type  Type[any]
}

func Iterate(ctx context.Context, t Type[any], value any, optFns ...OptionFunc) (<-chan Entry, bool) {
	if i, ok := t.(EntryIterator); ok {
		if e := i.EntryIter(ctx, value, optFns...); e != nil {
			return e, true
		}
	}
	return nil, false
}

func SetTo(v any, key any, value any) {
	if value == nil {
		return
	}

	switch x := v.(type) {
	case map[string]any:
		if k, ok := key.(string); ok {
			x[k] = value
		}
		return

	case []any:
		if k, ok := key.(int); ok {
			x[k] = value
		}
		return
	}
}
