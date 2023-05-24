package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestMerge(t *testing.T) {
	seq0 := Of(0, 2, 4)
	seq1 := Of(1, 3, 5)

	merged := Merge(seq1, seq0)

	values := slices.Sorted(merged)
	testingx.Expect(t, values, testingx.Equal([]int{
		0, 1, 2, 3, 4, 5,
	}))
}
