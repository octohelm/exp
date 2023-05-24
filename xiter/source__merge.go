package xiter

import (
	"cmp"
	"iter"
)

func Merge[T cmp.Ordered](seq1, seq2 iter.Seq[T]) iter.Seq[T] {
	return MergeFunc(seq1, seq2, cmp.Compare)
}

func MergeFunc[T any](seq1, seq2 iter.Seq[T], compare func(T, T) int) iter.Seq[T] {
	return func(yield func(T) bool) {
		p1, stop := iter.Pull(seq1)
		defer stop()
		p2, stop := iter.Pull(seq2)
		defer stop()

		v1, ok1 := p1()
		v2, ok2 := p2()
		for ok1 || ok2 {
			var c int
			if ok1 && ok2 {
				c = compare(v1, v2)
			}

			switch {
			case !ok2 || c < 0:
				if !yield(v1) {
					return
				}
				v1, ok1 = p1()
			case !ok1 || c > 0:
				if !yield(v2) {
					return
				}
				v2, ok2 = p2()
			default:
				if !yield(v1) || !yield(v2) {
					return
				}
				v1, ok1 = p1()
				v2, ok2 = p2()
			}
		}
	}
}
