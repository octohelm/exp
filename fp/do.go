package fp

// Do 绑定一个额外参数，并将函数适配为 Operator。
func Do[Fn func(T, A) R, T, A, R any](fn Fn, a A) Operator[T, R] {
	return func(in T) R {
		return fn(in, a)
	}
}

// Do2 绑定两个额外参数，并将函数适配为 Operator。
func Do2[Fn func(T, A, B) R, T, A, B, R any](fn Fn, a A, b B) Operator[T, R] {
	return func(in T) R {
		return fn(in, a, b)
	}
}

// Do3 绑定三个额外参数，并将函数适配为 Operator。
func Do3[Fn func(T, A, B, C) R, T, A, B, C, R any](fn Fn, a A, b B, c C) Operator[T, R] {
	return func(in T) R {
		return fn(in, a, b, c)
	}
}

// Do4 绑定四个额外参数，并将函数适配为 Operator。
func Do4[Fn func(T, A, B, C, D) R, T, A, B, C, D, R any](fn Fn, a A, b B, c C, d D) Operator[T, R] {
	return func(in T) R {
		return fn(in, a, b, c, d)
	}
}
