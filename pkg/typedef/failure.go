package typedef

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var InvalidType = errors.New("invalid type")

func WrapFailure(err error, t Type[any], v any, optFns ...OptionFunc) error {
	o := newOption(optFns...)

	f := &Failure{
		Type:       t.Kind(),
		Value:      v,
		Path:       o.Path,
		Branch:     o.Branch,
		Msg:        o.Msg,
		Refinement: o.Refinement,
	}

	if f.Msg == "" {
		if errors.Is(err, InvalidType) {
			f.Msg = fmt.Sprintf("expect value of type %s, but receive %v", t.Kind(), v)
		} else {
			f.Msg = err.Error()
		}
	}

	return f
}

type Failure struct {
	Type       string
	Value      any
	Path       Path
	Branch     []any
	Msg        string
	Refinement string
}

func (f *Failure) Error() string {
	return fmt.Sprintf("%s: %s", f.Path, f.Msg)
}

func Wrap(left error, err error) error {
	if left == nil {
		return err
	}

	var failures Failures

	switch x := left.(type) {
	case Failures:
		failures = x
	case *Failure:
		failures = append(failures, x)
	}

	switch x := err.(type) {
	case Failures:
		failures = append(failures, x...)
	case *Failure:
		failures = append(failures, x)
	}

	return failures
}

type Failures []*Failure

func (failures Failures) Error() string {
	s := &strings.Builder{}
	for i := range failures {
		if i > 0 {
			s.WriteString("; ")
		}
		s.WriteString(failures[i].Error())
	}
	return s.String()
}
