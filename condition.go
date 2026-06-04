package lo

// Ternary is a single line if/else statement.
// Take care to avoid dereferencing potentially nil pointers in your A/B expressions, because they are both evaluated. See TernaryF to avoid this problem.
// Play: https://go.dev/play/p/t-D7WBL44h2
func Ternary[T any](condition bool, ifOutput, elseOutput T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// TernaryF is a single line if/else statement whose options are functions.
// Play: https://go.dev/play/p/AO4VW20JoqM
func TernaryF[T any](condition bool, ifFunc, elseFunc func() T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Perf: value receivers (not pointer) allow the compiler to fully inline the entire
// If().ElseIf().Else() chain, eliminating all function call overhead.
type ifElse[T any] struct {
	result T
	done   bool
}

// If is a single line if/else statement.
// Play: https://go.dev/play/p/WSw3ApMxhyW
func If[T any](condition bool, result T) ifElse[T] {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

// IfF is a single line if/else statement whose options are functions.
// Play: https://go.dev/play/p/WSw3ApMxhyW
func IfF[T any](condition bool, resultF func() T) ifElse[T] {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

// ElseIf.
// Play: https://go.dev/play/p/WSw3ApMxhyW
func (i ifElse[T]) ElseIf(condition bool, result T) ifElse[T] {
	_ = "STUB: not implemented"
	return nil
}

// ElseIfF.
// Play: https://go.dev/play/p/WSw3ApMxhyW
func (i ifElse[T]) ElseIfF(condition bool, resultF func() T) ifElse[T] {
	_ = "STUB: not implemented"
	return nil
}

// Else.
// Play: https://go.dev/play/p/WSw3ApMxhyW
func (i ifElse[T]) Else(result T) T { _ = "STUB: not implemented"; return *new(T) }

// ElseF.
// Play: https://go.dev/play/p/WSw3ApMxhyW
func (i ifElse[T]) ElseF(resultF func() T) T { _ = "STUB: not implemented"; return *new(T) }

// Perf: value receivers (not pointer) allow the compiler to fully inline the entire
// Switch().Case().Default() chain, eliminating all function call overhead.
type switchCase[T comparable, R any] struct {
	predicate T
	result    R
	done      bool
}

// Switch is a pure functional switch/case/default statement.
// Play: https://go.dev/play/p/TGbKUMAeRUd
func Switch[T comparable, R any](predicate T) switchCase[T, R] {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

// Case.
// Play: https://go.dev/play/p/TGbKUMAeRUd
func (s switchCase[T, R]) Case(val T, result R) switchCase[T, R] {
	_ = "STUB: not implemented"
	return nil
}

// CaseF.
// Play: https://go.dev/play/p/TGbKUMAeRUd
func (s switchCase[T, R]) CaseF(val T, callback func() R) switchCase[T, R] {
	_ = "STUB: not implemented"
	return nil
}

// Default.
// Play: https://go.dev/play/p/TGbKUMAeRUd
func (s switchCase[T, R]) Default(result R) R { _ = "STUB: not implemented"; return *new(R) }

// DefaultF.
// Play: https://go.dev/play/p/TGbKUMAeRUd
func (s switchCase[T, R]) DefaultF(callback func() R) R { _ = "STUB: not implemented"; return *new(R) }
