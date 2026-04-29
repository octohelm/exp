package xiter_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/exp/xiter"
)

func TestChunk(t *testing.T) {
	t.Run("基础分块测试", func(t *testing.T) {
		src := slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		chunked := xiter.Chunk(src, 3)

		Then(t, "数据应按指定大小分块，且处理末尾余数",
			Expect(slices.Collect(chunked), Equal([][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9},
			})),
		)
	})

	t.Run("边界条件", func(t *testing.T) {
		t.Run("空序列", func(t *testing.T) {
			src := slices.Values([]int{})
			chunked := xiter.Chunk(src, 5)

			Then(t, "空序列应返回空切片",
				Expect(slices.Collect(chunked), Be(cmp.Nil[[][]int]())),
			)
		})

		t.Run("分块大小大于总长度", func(t *testing.T) {
			src := slices.Values([]int{1, 2})
			chunked := xiter.Chunk(src, 10)

			Then(t, "应返回包含所有元素的单块",
				Expect(slices.Collect(chunked), Equal([][]int{{1, 2}})),
			)
		})
	})
}

func FuzzChunk(f *testing.F) {
	for range 10 {
		f.Add(rand.IntN(5000))
	}

	f.Fuzz(func(t *testing.T, total int) {
		if total < 0 {
			t.Skip()
		}

		src := func(yield func(int) bool) {
			for i := range total {
				if !yield(i) {
					return
				}
			}
		}

		chunkN := 50
		chunks := slices.Collect(xiter.Chunk(src, chunkN))

		Then(t, "验证分块逻辑完整性",
			ExpectMust(func() error {
				for i, chunk := range chunks {
					expectedSize := chunkN
					if i == len(chunks)-1 {
						expectedSize = total - chunkN*(len(chunks)-1)
					}

					if len(chunk) != expectedSize {
						t.Errorf("chunk %d: 期望长度 %d, 实际长度 %d", i, expectedSize, len(chunk))
					}
				}
				return nil
			}),
			Expect(len(slices.Concat(chunks...)), Equal(total)),
		)
	})
}
