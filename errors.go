package lo

import (
	"fmt"
	"reflect"
)

const defaultAssertionFailureMessage = "assertion failed"

// Validate is a helper that creates an error when a condition is not met.
// Play: https://go.dev/play/p/vPyh51XpCBt
func Validate(ok bool, format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func messageFromMsgAndArgs(msgAndArgs ...any) string { _ = "STUB: not implemented"; return "" }

//nolint:errcheck,forcetypeassert

// MustChecker panics if err is error or false.
var MustChecker = func(err any, messageArgs ...any) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case bool:
		if !e {
			message := messageFromMsgAndArgs(messageArgs...)
			if message == "" {
				message = "not ok"
			}

			panic(message)
		}

	case error:
		message := messageFromMsgAndArgs(messageArgs...)
		if message != "" {
			panic(message + ": " + e.Error())
		}
		panic(e.Error())

	default:
		panic("must: invalid err type '" + reflect.TypeOf(err).Name() + "', should either be a bool or an error")
	}
}

// Must is a helper that wraps a call to a function returning a value and an error
// and panics if err is error or false.
// Play: https://go.dev/play/p/fOqtX5HudtN
func Must[T any](val T, err any, messageArgs ...any) T { _ = "STUB: not implemented"; return *new(T) }

// Must0 has the same behavior as Must, but callback returns no variable.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must0(err any, messageArgs ...any) { _ = "STUB: not implemented"; return }

// Must1 is an alias to Must.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must1[T any](val T, err any, messageArgs ...any) T { _ = "STUB: not implemented"; return *new(T) }

// Must2 has the same behavior as Must, but callback returns 2 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must2[T1, T2 any](val1 T1, val2 T2, err any, messageArgs ...any) (T1, T2) {
	_ = "STUB: not implemented"
	return *new(T1), *new(T2)
}

// Must3 has the same behavior as Must, but callback returns 3 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must3[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err any, messageArgs ...any) (T1, T2, T3) {
	_ = "STUB: not implemented"
	return *new(T1), *new(T2), *new(T3)
}

// Must4 has the same behavior as Must, but callback returns 4 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must4[T1, T2, T3, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err any, messageArgs ...any) (T1, T2, T3, T4) {
	_ = "STUB: not implemented"
	return *new(T1), *new(T2), *new(T3), *new(T4)
}

// Must5 has the same behavior as Must, but callback returns 5 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must5[T1, T2, T3, T4, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, err any, messageArgs ...any) (T1, T2, T3, T4, T5) {
	_ = "STUB: not implemented"
	return *new(T1), *new(T2), *new(T3), *new(T4), *new(T5)
}

// Must6 has the same behavior as Must, but callback returns 6 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must6[T1, T2, T3, T4, T5, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, err any, messageArgs ...any) (T1, T2, T3, T4, T5, T6) {
	_ = "STUB: not implemented"
	return *new(T1), *new(T2), *new(T3), *new(T4), *new(T5), *new(T6)
}

// Try calls the function and return false in case of error.
func Try(callback func() error) (ok bool) { _ = "STUB: not implemented"; return false }

// Try0 has the same behavior as Try, but callback returns no variable.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try0(callback func()) bool { _ = "STUB: not implemented"; return false }

// Try1 is an alias to Try.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try1(callback func() error) bool { _ = "STUB: not implemented"; return false }

// Try2 has the same behavior as Try, but callback returns 2 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try2[T any](callback func() (T, error)) bool { _ = "STUB: not implemented"; return false }

// Try3 has the same behavior as Try, but callback returns 3 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try3[T, R any](callback func() (T, R, error)) bool { _ = "STUB: not implemented"; return false }

// Try4 has the same behavior as Try, but callback returns 4 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try4[T, R, S any](callback func() (T, R, S, error)) bool {
	_ = "STUB: not implemented"
	return false
}

// Try5 has the same behavior as Try, but callback returns 5 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try5[T, R, S, Q any](callback func() (T, R, S, Q, error)) bool {
	_ = "STUB: not implemented"
	return false
}

// Try6 has the same behavior as Try, but callback returns 6 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try6[T, R, S, Q, U any](callback func() (T, R, S, Q, U, error)) bool {
	_ = "STUB: not implemented"
	return false
}

// TryOr has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr[A any](callback func() (A, error), fallbackA A) (A, bool) {
	_ = "STUB: not implemented"
	return *new(A), false
}

// TryOr1 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr1[A any](callback func() (A, error), fallbackA A) (A, bool) {
	_ = "STUB: not implemented"
	return *new(A), false
}

// TryOr2 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr2[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (A, B, bool) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), false
}

// TryOr3 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr3[A, B, C any](callback func() (A, B, C, error), fallbackA A, fallbackB B, fallbackC C) (A, B, C, bool) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), false
}

// TryOr4 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr4[A, B, C, D any](callback func() (A, B, C, D, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D) (A, B, C, D, bool) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), false
}

// TryOr5 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr5[A, B, C, D, E any](callback func() (A, B, C, D, E, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E) (A, B, C, D, E, bool) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), false
}

// TryOr6 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr6[A, B, C, D, E, F any](callback func() (A, B, C, D, E, F, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E, fallbackF F) (A, B, C, D, E, F, bool) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), false
}

// TryWithErrorValue has the same behavior as Try, but also returns value passed to panic.
// Play: https://go.dev/play/p/Kc7afQIT2Fs
func TryWithErrorValue(callback func() error) (errorValue any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// TryCatch has the same behavior as Try, but calls the catch function in case of error.
// Play: https://go.dev/play/p/PnOON-EqBiU
func TryCatch(callback func() error, catch func()) { _ = "STUB: not implemented"; return }

// TryCatchWithErrorValue has the same behavior as TryWithErrorValue, but calls the catch function in case of error.
// Play: https://go.dev/play/p/8Pc9gwX_GZO
func TryCatchWithErrorValue(callback func() error, catch func(any)) {
	_ = "STUB: not implemented"
	return
}

// ErrorsAs is a shortcut for errors.As(err, &&T).
// Play: https://go.dev/play/p/8wk5rH8UfrE
func ErrorsAs[T error](err error) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Assert does nothing when the condition is true, otherwise it panics with an optional message.
// Play: https://go.dev/play/p/Xv8LLKBMNwI
var Assert = func(condition bool, message ...string) {
	if condition {
		return
	}

	panicMessage := defaultAssertionFailureMessage
	if len(message) > 0 {
		panicMessage = fmt.Sprintf("%s: %s", defaultAssertionFailureMessage, message[0])
	}
	panic(panicMessage)
}

// Assertf does nothing when the condition is true, otherwise it panics with a formatted message.
// Play: https://go.dev/play/p/TVPEmVcyrdY
var Assertf = func(condition bool, format string, args ...any) {
	if condition {
		return
	}

	panicMessage := fmt.Sprintf("%s: %s", defaultAssertionFailureMessage, fmt.Sprintf(format, args...))
	panic(panicMessage)
}
