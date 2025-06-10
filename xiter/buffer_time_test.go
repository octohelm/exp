package xiter

import (
	"iter"
	"slices"
	"testing"
	"time"

	testingx "github.com/octohelm/x/testing"
)

func TestBufferTime(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		src := Seq(func(yield func(int) bool) {
			for i := 0; i < 5; i++ {
				if !yield(i) {
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
		})

		buffered := BufferTime(src, 500*time.Millisecond)

		testingx.Expect(t, slices.Collect(buffered), testingx.Equal([][]int{
			{0, 1, 2, 3, 4},
		}))
	})

	t.Run("long pause", func(t *testing.T) {
		src := func() iter.Seq[int] {
			return func(yield func(int) bool) {
				yield(0)
				time.Sleep(50 * time.Millisecond)
				yield(1)
				time.Sleep(50 * time.Millisecond)
				yield(2)
				time.Sleep(200 * time.Millisecond) // Long pause
				yield(3)
				time.Sleep(50 * time.Millisecond)
				yield(4)
			}
		}()

		buffered := BufferTime(src, 100*time.Millisecond)

		testingx.Expect(t, slices.Collect(buffered), testingx.Equal([][]int{
			{0, 1, 2},
			{3, 4},
		}))
	})
}
