package xchan

import "errors"

type Observable[T any] interface {
	Value() <-chan T
}

type ValueNotifier[T any] interface {
	Send(x T)
}

type Cancelable interface {
	CancelCause(err error)
}

type Observer[T any] interface {
	Observable[T]
	Cancelable
	Done() <-chan struct{}
}

var Completed = errors.New("completed")
