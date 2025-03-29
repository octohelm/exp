package xiter

import (
	"slices"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestConcat(t *testing.T) {
	seq0 := Of(0, 2, 4)
	seq1 := Of(1, 3, 5)

	seq := Concat(seq0, seq1)

	values := slices.Collect(seq)
	testingx.Expect(t, values, testingx.Equal([]int{
		0, 2, 4, 1, 3, 5,
	}))
}
