package typedef

import (
	"fmt"
	"strings"
)

type OptionFunc = func(o *option)

func WithMask(mask bool) OptionFunc {
	return func(o *option) {
		o.Mask = mask
	}
}

func WithPath(key any) OptionFunc {
	return func(o *option) {
		o.Path = append(o.Path, key)
	}
}

func WithMsg(msg string) OptionFunc {
	return func(o *option) {
		o.Msg = msg
	}
}

func WithRefinement(refinement string) OptionFunc {
	return func(o *option) {
		o.Refinement = refinement
	}
}

func WithBranch(value any) OptionFunc {
	return func(o *option) {
		o.Branch = append(o.Branch, value)
	}
}

func WithCoerce(coerce bool) OptionFunc {
	return func(o *option) {
		o.Coerce = coerce
	}
}

type option struct {
	Mask       bool
	Coerce     bool
	Branch     []any
	Path       Path
	Msg        string
	Refinement string
}

func newOption(optFns ...OptionFunc) *option {
	o := &option{
		Coerce: true,
	}

	for i := range optFns {
		optFns[i](o)
	}

	return o
}

type Path []any

func (p Path) String() string {
	b := &strings.Builder{}
	for i := range p {
		_, _ = fmt.Fprintf(b, "/%v", p[i])
	}
	return b.String()
}
