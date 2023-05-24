package fp

func Do[Fn func(T, A) R, T, A, R any](fn Fn, a A) Operator[T, R] {
	return func(in T) R {
		return fn(in, a)
	}
}

func Do2[Fn func(T, A, B) R, T, A, B, R any](fn Fn, a A, b B) Operator[T, R] {
	return func(in T) R {
		return fn(in, a, b)
	}
}

func Do3[Fn func(T, A, B, C) R, T, A, B, C, R any](fn Fn, a A, b B, c C) Operator[T, R] {
	return func(in T) R {
		return fn(in, a, b, c)
	}
}

func Do4[Fn func(T, A, B, C, D) R, T, A, B, C, D, R any](fn Fn, a A, b B, c C, d D) Operator[T, R] {
	return func(in T) R {
		return fn(in, a, b, c, d)
	}
}
