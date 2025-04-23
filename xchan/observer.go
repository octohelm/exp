package xchan

import (
	"sync"
	"sync/atomic"
)

func NewNotifiableObserver[T any]() NotifiableObserver[T] {
	return &observer[T]{
		value: make(chan T),
	}
}

var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type observer[T any] struct {
	value chan T
	mu    sync.Mutex
	done  atomic.Value
	err   error
}

func (c *observer[T]) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *observer[T]) Send(v T) {
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}

	c.value <- v
	c.mu.Unlock()
}

func (c *observer[T]) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}

func (c *observer[T]) Value() <-chan T {
	return c.value
}

func (c *observer[T]) CancelCause(err error) {
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}

	if err == nil {
		err = Completed
	}

	c.err = err

	d, _ := c.done.Load().(chan struct{})
	if d == nil {
		c.done.Store(closedchan)
	} else {
		close(d)
	}

	close(c.value)
	c.mu.Unlock()
}
