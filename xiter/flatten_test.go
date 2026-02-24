package xiter

import (
	"errors"
	"slices"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"
)

func TestFlatten(t *testing.T) {
	src1 := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	src2 := Seq(func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	})

	flattenedSeq := Flatten(Of(src1, src2))

	t.Run("基本结果校验", func(t *testing.T) {
		Then(t, "扁平化后的结果应符合预期",
			Expect(slices.Collect(flattenedSeq), Equal([]int{
				0, 1, 2,
				0, 1, 2,
			})),
		)
	})

	t.Run("带有 Pointer 路径的元素校验", func(t *testing.T) {
		err := cmp.Every(cmp.Lt(2))(flattenedSeq)
		if err != nil {
			e, _ := errors.AsType[*cmp.ErrCheck](err)

			Then(t, "应精准定位到导致失败的扁平化下标",
				Expect(string(e.Pointer), Be(cmp.Eq("/2"))),
			)
		}
	})
}
