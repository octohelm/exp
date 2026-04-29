package xiter_test

import (
	"slices"
	"testing"
	"time"

	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/exp/xiter"
)

func TestSinkAndChannelHelpers(t *testing.T) {
	t.Run("Count", func(t *testing.T) {
		Then(t, "Count 应返回元素数量",
			Expect(xiter.Count(xiter.Of(1, 2, 3, 4)), Equal(4)),
		)
	})

	t.Run("Fold", func(t *testing.T) {
		Then(t, "Fold 应按顺序折叠元素",
			Expect(xiter.Fold(xiter.Of(1, 2, 3, 4), func(a, b int) int { return a + b }), Equal(10)),
		)
	})

	t.Run("Reduce", func(t *testing.T) {
		Then(t, "Reduce 应从初始值开始累积",
			Expect(xiter.Reduce(xiter.Of(1, 2, 3, 4), 10, func(sum, v int) int { return sum + v }), Equal(20)),
		)
	})

	t.Run("Recv", func(t *testing.T) {
		values := make(chan int, 3)
		values <- 1
		values <- 2
		values <- 3
		close(values)

		Then(t, "Recv 应消费完整通道",
			Expect(slices.Collect(xiter.Recv(values)), Equal([]int{1, 2, 3})),
		)
	})
}

func TestDebounceTime(t *testing.T) {
	src := xiter.Seq(func(yield func(int) bool) {
		yield(1)
		time.Sleep(10 * time.Millisecond)
		yield(2)
		time.Sleep(50 * time.Millisecond)
		yield(3)
		time.Sleep(10 * time.Millisecond)
		yield(4)
	})

	Then(t, "DebounceTime 应只输出每个窗口中的最后一个值",
		Expect(slices.Collect(xiter.DebounceTime(src, 20*time.Millisecond)), Equal([]int{2, 4})),
	)
}
