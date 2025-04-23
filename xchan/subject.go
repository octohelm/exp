package xchan

import (
	"sync"
	"sync/atomic"
)

type Subject[T any] struct {
	mu          sync.Mutex
	subscribers map[Subscriber[T]]struct{}
	done        atomic.Value
	err         error
}

var _ Subscriber[int] = &Subject[int]{}

func (c *Subject[T]) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *Subject[T]) Done() <-chan struct{} {
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

func (c *Subject[T]) CancelCause(err error) {
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

	for o := range c.subscribers {
		o.CancelCause(err)
	}
	c.subscribers = nil
	c.mu.Unlock()

	return
}

func (c *Subject[T]) Send(value T) {
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}

	for ob := range c.subscribers {
		ob.Send(value)
	}
	c.mu.Unlock()
}

func (c *Subject[T]) Observe() Observer[T] {
	o := NewNotifiableObserver[T]()
	c.Subscribe(o)
	return o
}

func (c *Subject[T]) Subscribe(o Subscriber[T]) {
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}

	if c.subscribers == nil {
		c.subscribers = map[Subscriber[T]]struct{}{}
	}
	c.subscribers[o] = struct{}{}
	c.mu.Unlock()

	go func() {
		<-o.Done()

		c.mu.Lock()
		delete(c.subscribers, o)
		c.mu.Unlock()
	}()
}
