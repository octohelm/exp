package fp_test

import (
	"testing"

	. "github.com/octohelm/x/testing/v2"

	. "github.com/octohelm/exp/fp"
)

func TestPipeLongChainsAndDoHelpers(t *testing.T) {
	add := func(v, n int) int { return v + n }
	mul := func(v, n int) int { return v * n }
	sum3 := func(v, a, b, c int) int { return v + a + b + c }
	sum4 := func(v, a, b, c, d int) int { return v + a + b + c + d }

	t.Run("Do3", func(t *testing.T) {
		Then(t, "Do3 应绑定三个额外参数",
			Expect(Do3(sum3, 2, 3, 4)(1), Equal(10)),
		)
	})

	t.Run("Do4", func(t *testing.T) {
		Then(t, "Do4 应绑定四个额外参数",
			Expect(Do4(sum4, 2, 3, 4, 5)(1), Equal(15)),
		)
	})

	t.Run("Pipe5", func(t *testing.T) {
		Then(t, "Pipe5 应按顺序应用五个操作",
			Expect(Pipe5(1, Do(add, 1), Do(mul, 2), Do(add, 3), Do(mul, 4), Do(add, 5)), Equal(33)),
		)
	})

	t.Run("Pipe6", func(t *testing.T) {
		Then(t, "Pipe6 应按顺序应用六个操作",
			Expect(Pipe6(1, Do(add, 1), Do(mul, 2), Do(add, 3), Do(mul, 4), Do(add, 5), Do(mul, 6)), Equal(198)),
		)
	})

	t.Run("Pipe7", func(t *testing.T) {
		Then(t, "Pipe7 应按顺序应用七个操作",
			Expect(Pipe7(1, Do(add, 1), Do(mul, 2), Do(add, 3), Do(mul, 4), Do(add, 5), Do(mul, 6), Do(add, 7)), Equal(205)),
		)
	})

	t.Run("Pipe8", func(t *testing.T) {
		Then(t, "Pipe8 应按顺序应用八个操作",
			Expect(Pipe8(1, Do(add, 1), Do(mul, 2), Do(add, 3), Do(mul, 4), Do(add, 5), Do(mul, 6), Do(add, 7), Do(mul, 8)), Equal(1640)),
		)
	})

	t.Run("Pipe9", func(t *testing.T) {
		Then(t, "Pipe9 应按顺序应用九个操作",
			Expect(Pipe9(1, Do(add, 1), Do(mul, 2), Do(add, 3), Do(mul, 4), Do(add, 5), Do(mul, 6), Do(add, 7), Do(mul, 8), Do(add, 9)), Equal(1649)),
		)
	})
}
