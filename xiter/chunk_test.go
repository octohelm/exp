package xiter

import (
	"math/rand/v2"
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestChunk(t *testing.T) {
	src := Seq(func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
		}
	})

	chunked := Chunk(src, 3)

	testingx.Expect(t, slices.Collect(chunked), testingx.Equal([][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}))
}

func FuzzChunk(f *testing.F) {
	for range 10 {
		f.Add(rand.IntN(50000))
	}

	f.Fuzz(func(t *testing.T, total int) {
		src := Seq(func(yield func(int) bool) {
			for i := range total {
				if !yield(i) {
					return
				}
			}
		})

		chunkN := 50

		chunked := Chunk(src, chunkN)
		chunks := slices.Collect(chunked)

		for i, chunk := range chunks {
			// last one
			if i == len(chunks)-1 {
				t.Run("last chunk", func(t *testing.T) {
					testingx.Expect(t, len(chunk), testingx.Equal(total-chunkN*(len(chunks)-1)))
				})
				continue
			}

			testingx.Expect(t, len(chunk), testingx.Equal(50))
		}
	})
}
