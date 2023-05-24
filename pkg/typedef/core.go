package typedef

import "context"

type Type[T any] interface {
	Kind() string
}

type Coercer interface {
	Coerce(ctx context.Context, value any, optFns ...OptionFunc) (any, error)
}

type Validator[T any] interface {
	Validate(ctx context.Context, value any, optFns ...OptionFunc) (*T, error)
}

type Refiner[T any] interface {
	Refine(ctx context.Context, value any, optFns ...OptionFunc) (*T, error)
}

type EntryIterator interface {
	EntryIter(ctx context.Context, value any, optFns ...OptionFunc) <-chan Entry
}
