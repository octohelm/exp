package typedef

import (
	"context"
	"github.com/pkg/errors"
)

func Object[T any](props T) Type[T] {
	switch x := any(props).(type) {
	case map[string]any:
		s := map[string]Type[any]{}

		for k := range x {
			if t, ok := x[k].(Type[any]); ok {
				s[k] = t
			} else {
				panic(errors.Errorf("object schema need value Type, but got %T", t))
			}
		}

		return &objectType[T]{
			props: ObjSchema[T](s),
		}
	}

	return &objectType[T]{
		//props: props,
	}
}

var _ interface {
	Validator[any]
	Coercer
	EntryIterator
} = &objectType[any]{}

type objectType[T any] struct {
	props ObjectType[T]
}

func (t *objectType[T]) Kind() string {
	return "object"
}

func (t *objectType[T]) Validate(ctx context.Context, v any, optFns ...OptionFunc) (*T, error) {
	switch x := v.(type) {
	case T:
		return &x, nil
	}
	return nil, InvalidType
}

func (t *objectType[T]) Coerce(ctx context.Context, value any, optFns ...OptionFunc) (any, error) {
	if x, ok := t.props.Is(value); ok {
		return *x, nil
	}
	return nil, InvalidType
}

func (t *objectType[T]) EntryIter(ctx context.Context, v any, optFns ...OptionFunc) <-chan Entry {
	if o, ok := ToObj(v); ok {
		ch := make(chan Entry)

		go func() {
			defer close(ch)

			select {
			case <-ctx.Done():
				return
			default:
				if t.props != nil {
					for key := range t.props.Keys(ctx) {
						propType := t.props.PropType(key)
						propValue, _ := o.Get(key)

						ch <- Entry{
							Value: propValue,
							Type:  propType,
							Key:   key,
						}
					}
				}
			}
		}()

		return ch
	}

	return nil
}

type Obj interface {
	Get(k string) (any, bool)
	Set(k string, value any)
}

func ToObj(v any) (Obj, bool) {
	switch x := v.(type) {
	case Obj:
		return x, true
	case map[string]any:
		return ObjValue(x), true
	}
	return nil, false
}

type ObjValue map[string]any

func (o ObjValue) Get(k string) (any, bool) {
	if v, ok := o[k]; ok {
		return v, ok
	}
	return nil, false
}

func (o ObjValue) Set(k string, value any) {
	o[k] = value
}

type ObjectType[T any] interface {
	Keys(ctx context.Context) <-chan string
	PropType(k string) Type[any]
	Is(v any) (*T, bool)
	New() *T
}

type ObjSchema[T any] map[string]Type[any]

func (o ObjSchema[T]) Is(v any) (*T, bool) {
	if x, ok := v.(T); ok {
		return &x, true
	}
	return nil, false
}

func (o ObjSchema[T]) New() *T {
	return new(T)
}

func (o ObjSchema[T]) Keys(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)

		select {
		case <-ctx.Done():
			return
		default:
			for k := range o {
				ch <- k
			}
		}
	}()
	return ch
}

func (o ObjSchema[T]) PropType(k string) Type[any] {
	if v, ok := o[k]; ok {
		return v
	}
	return Never()
}
