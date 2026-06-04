package lo

// Partial returns new function that, when called, has its first argument set to the provided value.
// Play: https://go.dev/play/p/Sy1gAQiQZ3v
func Partial[T1, T2, R any](f func(a T1, b T2) R, arg1 T1) func(T2) R {
	_ = "STUB: not implemented"
	return nil
}

// Partial1 returns new function that, when called, has its first argument set to the provided value.
// Play: https://go.dev/play/p/D-ASTXCLBzw
func Partial1[T1, T2, R any](f func(T1, T2) R, arg1 T1) func(T2) R {
	_ = "STUB: not implemented"
	return nil

	// Partial2 returns new function that, when called, has its first argument set to the provided value.
	// Play: https://go.dev/play/p/-xiPjy4JChJ
}

func Partial2[T1, T2, T3, R any](f func(T1, T2, T3) R, arg1 T1) func(T2, T3) R {
	_ = "STUB: not implemented"
	return nil
}

// Partial3 returns new function that, when called, has its first argument set to the provided value.
// Play: https://go.dev/play/p/zWtSutpI26m
func Partial3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, arg1 T1) func(T2, T3, T4) R {
	_ = "STUB: not implemented"
	return nil
}

// Partial4 returns new function that, when called, has its first argument set to the provided value.
// Play: https://go.dev/play/p/kBrnnMTcJm0
func Partial4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, arg1 T1) func(T2, T3, T4, T5) R {
	_ = "STUB: not implemented"
	return nil
}

// Partial5 returns new function that, when called, has its first argument set to the provided value.
// Play: https://go.dev/play/p/7Is7K2y_VC3
func Partial5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, arg1 T1) func(T2, T3, T4, T5, T6) R {
	_ = "STUB: not implemented"
	return nil
}
