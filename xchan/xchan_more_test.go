package xchan

import (
	"context"
	"errors"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

type collectorSubscriber[T any] struct {
	mu     sync.Mutex
	values []T
	done   chan struct{}
	err    error
}

func newCollectorSubscriber[T any]() *collectorSubscriber[T] {
	return &collectorSubscriber[T]{done: make(chan struct{})}
}

func (c *collectorSubscriber[T]) Send(v T) {
	c.mu.Lock()
	c.values = append(c.values, v)
	c.mu.Unlock()
}

func (c *collectorSubscriber[T]) Done() <-chan struct{} {
	return c.done
}

func (c *collectorSubscriber[T]) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func (c *collectorSubscriber[T]) CancelCause(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err != nil {
		return
	}
	if err == nil {
		err = Completed
	}
	c.err = err
	close(c.done)
}

func (c *collectorSubscriber[T]) Values() []T {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]T(nil), c.values...)
}

func TestObservableFuncObserve(t *testing.T) {
	observer := NewNotifiableObserver[int]()

	fn := ObservableFunc[int](func() Observer[int] {
		return observer
	})

	go func() {
		observer.Send(1)
		observer.CancelCause(nil)
	}()

	Then(t, "ObservableFunc 应返回实际的观察者",
		Expect(slices.Collect(Values(context.Background(), fn.Observe())), Equal([]int{1})),
	)
}

func TestSubscribe(t *testing.T) {
	source := &Subject[int]{}
	subscriber := newCollectorSubscriber[int]()

	go func() {
		source.Send(1)
		source.Send(2)
		source.CancelCause(nil)
	}()

	Subscribe(context.Background(), source, subscriber)

	Then(t, "Subscribe 应把所有值转发给订阅者",
		Expect(subscriber.Values(), Equal([]int{1, 2})),
	)

	Then(t, "源完成后订阅者也应被关闭",
		Expect(subscriber.Err(), Equal(Completed)),
	)
}

func TestObserverAndSubjectLifecycle(t *testing.T) {
	t.Run("观察者", func(t *testing.T) {
		observer := NewNotifiableObserver[int]()

		Then(t, "新建 observer 的错误应为空",
			Expect(observer.Err(), Be(cmp.Nil[error]())),
		)

		done := observer.Done()
		observer.CancelCause(errors.New("停止"))

		Then(t, "取消后应保留错误并关闭 done/value",
			Expect(observer.Err().Error(), Equal("停止")),
			Expect(func() bool {
				select {
				case <-done:
					return true
				case <-time.After(100 * time.Millisecond):
					return false
				}
			}(), Be(cmp.Eq(true))),
			Expect(func() bool {
				_, ok := <-observer.Value()
				return ok
			}(), Be(cmp.Eq(false))),
		)

		observer.Send(1)
	})

	t.Run("主题", func(t *testing.T) {
		subject := &Subject[int]{}
		subscriber := newCollectorSubscriber[int]()

		subject.Subscribe(subscriber)
		done := subject.Done()
		subject.Send(1)
		subject.CancelCause(errors.New("已关闭"))
		subject.Send(2)

		Then(t, "主题取消后应保留错误并关闭 done",
			Expect(subject.Err().Error(), Equal("已关闭")),
			Expect(func() bool {
				select {
				case <-done:
					return true
				case <-time.After(100 * time.Millisecond):
					return false
				}
			}(), Be(cmp.Eq(true))),
			Expect(subscriber.Err().Error(), Equal("已关闭")),
			Expect(subscriber.Values(), Equal([]int{1})),
		)
	})
}
