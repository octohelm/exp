package fp

// Operator 表示接收 T 并返回 R 的单步变换。
type Operator[T any, R any] func(in T) R

// Pipe 依次执行一个操作并返回结果。
func Pipe[T, R any](
	in T,
	op1 Operator[T, R],
) R {
	return op1(in)
}

// Pipe2 依次执行两个操作并返回最终结果。
func Pipe2[T, A, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, R],
) R {
	return op2(op1(in))
}

// Pipe3 依次执行三个操作并返回最终结果。
func Pipe3[T, A, B, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, R],
) R {
	return op3(op2(op1(in)))
}

// Pipe4 依次执行四个操作并返回最终结果。
func Pipe4[T, A, B, C, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, C],
	op4 Operator[C, R],
) R {
	return op4(op3(op2(op1(in))))
}

// Pipe5 依次执行五个操作并返回最终结果。
func Pipe5[T, A, B, C, D, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, C],
	op4 Operator[C, D],
	op5 Operator[D, R],
) R {
	return op5(op4(op3(op2(op1(in)))))
}

// Pipe6 依次执行六个操作并返回最终结果。
func Pipe6[T, A, B, C, D, E, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, C],
	op4 Operator[C, D],
	op5 Operator[D, E],
	op6 Operator[E, R],
) R {
	return op6(op5(op4(op3(op2(op1(in))))))
}

// Pipe7 依次执行七个操作并返回最终结果。
func Pipe7[T, A, B, C, D, E, F, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, C],
	op4 Operator[C, D],
	op5 Operator[D, E],
	op6 Operator[E, F],
	op7 Operator[F, R],
) R {
	return op7(op6(op5(op4(op3(op2(op1(in)))))))
}

// Pipe8 依次执行八个操作并返回最终结果。
func Pipe8[T, A, B, C, D, E, F, G, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, C],
	op4 Operator[C, D],
	op5 Operator[D, E],
	op6 Operator[E, F],
	op7 Operator[F, G],
	op8 Operator[G, R],
) R {
	return op8(op7(op6(op5(op4(op3(op2(op1(in))))))))
}

// Pipe9 依次执行九个操作并返回最终结果。
func Pipe9[T, A, B, C, D, E, F, G, H, R any](
	in T,
	op1 Operator[T, A],
	op2 Operator[A, B],
	op3 Operator[B, C],
	op4 Operator[C, D],
	op5 Operator[D, E],
	op6 Operator[E, F],
	op7 Operator[F, G],
	op8 Operator[G, H],
	op9 Operator[H, R],
) R {
	return op9(op8(op7(op6(op5(op4(op3(op2(op1(in)))))))))
}
