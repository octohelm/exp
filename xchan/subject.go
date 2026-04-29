package xchan

import (
	"sync"
	"sync/atomic"
)

// Subject 表示可向多个 subscriber 广播值的 observable。
type Subject[T any] struct {
	mu          sync.Mutex
	subscribers map[Subscriber[T]]struct{}
	done        atomic.Value
	err         error
}

var _ Subscriber[int] = &Subject[int]{}

// Err 返回 subject 的结束原因。
func (c *Subject[T]) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

// Done 返回 subject 的结束信号通道。
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

// CancelCause 结束 subject，并把结束原因传播给所有订阅者。
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

// Send 向所有当前订阅者广播一个值。
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

// Observe 创建一个新的 observer 并订阅当前 subject。
func (c *Subject[T]) Observe() Observer[T] {
	o := NewNotifiableObserver[T]()
	c.Subscribe(o)
	return o
}

// Subscribe 注册一个 subscriber，并在其结束时自动移除。
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
