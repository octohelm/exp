package xchan

import (
	"iter"
	"sync"
	"sync/atomic"
)

type Subject[T any] struct {
	observers sync.Map
	closed    atomic.Bool
	err       atomic.Value
}

func (s *Subject[T]) CancelCause(err error) {
	if s.closed.Swap(true) {
		return
	}

	for o := range s.observer() {
		o.CancelCause(err)
	}

	return
}

func (s *Subject[T]) Send(value T) {
	if s.closed.Load() {
		return
	}

	for o := range s.observer() {
		if x, ok := o.(ValueNotifier[T]); ok {
			x.Send(value)
		}
	}
}

func (s *Subject[T]) observer() iter.Seq[Observer[T]] {
	return func(yield func(Observer[T]) bool) {
		for k := range s.observers.Range {
			if !yield(k.(Observer[T])) {
				return
			}
		}
	}
}

func (s *Subject[T]) Observe() Observer[T] {
	o := &observer[T]{}
	o.init()

	s.observers.Store(o, true)

	go func() {
		<-o.Done()

		s.observers.Delete(o)
	}()

	return o
}

type observer[T any] struct {
	value chan T
	done  chan struct{}

	closed atomic.Bool
	err    atomic.Value
}

func (c *observer[T]) init() {
	c.value = make(chan T)
	c.done = make(chan struct{})
}

func (c *observer[T]) Send(v T) {
	select {
	case <-c.done:
	case c.value <- v:
	}
}

func (c *observer[T]) Done() <-chan struct{} {
	return c.done
}

func (c *observer[T]) Value() <-chan T {
	return c.value
}

func (c *observer[T]) CancelCause(err error) {
	if c.closed.Swap(true) {
		return
	}
	if err == nil {
		err = Completed
	}
	c.err.Store(err)
	close(c.done)
}
