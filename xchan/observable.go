package xchan

import (
	"errors"
)

// Completed 表示 observable 或 subscriber 以正常完成的方式结束。
var Completed = errors.New("completed")

// ValueObservable 暴露只读值通道。
type ValueObservable[T any] interface {
	Value() <-chan T
}

// ValueNotifier 暴露发送值的能力。
type ValueNotifier[T any] interface {
	Send(x T)
}

// Cancelable 暴露带原因的取消能力。
type Cancelable interface {
	CancelCause(err error)
}

// Observer 表示可读取值、感知结束并可取消的订阅者。
type Observer[T any] interface {
	ValueObservable[T]
	Cancelable
	Done() <-chan struct{}
	Err() error
}

// Observable 表示可产生 Observer 的源。
type Observable[T any] interface {
	Observe() Observer[T]
}

// NotifiableObserver 表示同时可接收值和对外观察值的 observer。
type NotifiableObserver[T any] interface {
	ValueNotifier[T]
	ValueObservable[T]
	Cancelable
	Done() <-chan struct{}
	Err() error
}

// Subscriber 表示可接收值并能被取消的下游消费者。
type Subscriber[T any] interface {
	ValueNotifier[T]
	Cancelable
	Done() <-chan struct{}
	Err() error
}

// ObservableFunc 将函数适配为 Observable。
type ObservableFunc[T any] func() Observer[T]

// Observe 调用底层函数并返回一个 Observer。
func (fn ObservableFunc[T]) Observe() Observer[T] {
	return fn()
}
