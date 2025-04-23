package xchan

import "errors"

var Completed = errors.New("completed")

type ValueObservable[T any] interface {
	Value() <-chan T
}

type ValueNotifier[T any] interface {
	Send(x T)
}

type Cancelable interface {
	CancelCause(err error)
}

type Observer[T any] interface {
	ValueObservable[T]
	Cancelable
	Done() <-chan struct{}
	Err() error
}

type Observable[T any] interface {
	Observe() Observer[T]
}

type NotifiableObserver[T any] interface {
	ValueNotifier[T]
	ValueObservable[T]
	Cancelable
	Done() <-chan struct{}
	Err() error
}

type Subscriber[T any] interface {
	ValueNotifier[T]
	Cancelable
	Done() <-chan struct{}
	Err() error
}

type ObservableFunc[T any] func() Observer[T]

func (fn ObservableFunc[T]) Observe() Observer[T] {
	return fn()
}
