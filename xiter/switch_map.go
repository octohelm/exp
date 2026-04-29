package xiter

import (
	"iter"
)

// SwitchMap 将每个元素映射为一个子序列，并按顺序展开输出。
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
