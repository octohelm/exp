package typedef

import (
	"context"
	"testing"

	"github.com/octohelm/exp/internal/testingutil"
)

func TestType(t *testing.T) {
	t.Run("Name", func(t *testing.T) {
		s := String()
		_, err := Validate(context.Background(), s, 2)
		testingutil.Expect(t, err, testingutil.NotBe[error](nil))
	})

	t.Run("Optional Name", func(t *testing.T) {
		s := Defaulted("default")(Optional(String()))
		v, err := Validate(context.Background(), s, nil)

		testingutil.Expect(t, err, testingutil.Be[error](nil))
		testingutil.Expect(t, *v, testingutil.Be("default"))
	})

	t.Run("Object", func(t *testing.T) {
		s := Object(map[string]any{
			"b": Defaulted(1)(Integer[int]()),
			"c": Optional(Number[float64]()),
			"options": Optional(
				Object(map[string]Type[any]{
					"label": String(),
					"value": Any(),
				}),
			),
		})

		v, err := Validate(context.Background(), s, map[string]any{})

		testingutil.Expect(t, err, testingutil.Be[error](nil))
		testingutil.Expect(t, *v, testingutil.Equal(map[string]any{
			"b": 1,
		}))
	})

	t.Run("ObjectWithStruct", func(t *testing.T) {
		type X struct {
			A int `json:"a,omitempty" default:"1"`
		}

		type O struct {
			X
			B string `json:"b,omitempty" default:"some"`
		}

		s := Object(O{})

		v, err := Validate(context.Background(), s, map[string]any{})
		testingutil.Expect(t, err, testingutil.Be[error](nil))
		testingutil.Expect(t, *v, testingutil.Equal(O{}))
	})
}

func Benchmark(b *testing.B) {
	s := Object(map[string]any{
		"a": Defaulted(1)(Integer[int]()),
	})

	for i := 0; i < b.N; i++ {
		_, _ = Validate(context.Background(), s, map[string]any{})
	}
}
