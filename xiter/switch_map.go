package xiter

import "iter"

func SwitchMap[I any, O any](seq iter.Seq[I], m func(e I) iter.Seq[O]) iter.Seq[O] {
	return func(yield func(O) bool) {
		for v := range seq {
			for vv := range m(v) {
				if !yield(vv) {
					return
				}
			}
		}
	}
}
