package lo

import (
	"github.com/samber/lo/internal/constraints"
)

// Range creates a slice of numbers (positive and/or negative) with given length.
// Play: https://go.dev/play/p/rho00R0WuHs
func Range(elementNum int) []int { _ = "STUB: not implemented"; return nil }

// RangeFrom creates a slice of numbers from start with specified length.
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) []T {
	_ = "STUB: not implemented"
	return nil
}

// RangeWithSteps creates a slice of numbers (positive and/or negative) progressing from start up to, but not including end.
// step set to zero will return an empty slice.
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) []T {
	_ = "STUB: not implemented"
	return nil
}

// Use math.Ceil instead of (count-1)/delta+1 because integer division
// fails for floats (e.g., 5.5/2.5=2.2 → ceil=3, not 2).

// Clamp clamps number within the inclusive lower and upper bounds.
// Play: https://go.dev/play/p/RU4lJNC2hlI
func Clamp[T constraints.Ordered](value, mIn, mAx T) T { _ = "STUB: not implemented"; return *new(T) }

// Sum sums the values in a collection. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/upfeJVqs4Bt
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// SumBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/Dz_a_7jN_ca
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByErr summarizes the values in a collection using the given return value from the iteration function.
// If the iteratee returns an error, iteration stops and the error is returned.
// If collection is empty 0 and nil error are returned.
func SumByErr[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// Product gets the product of the values in a collection. If collection is empty 1 is returned.
// Play: https://go.dev/play/p/2_kjM_smtAH
func Product[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// ProductBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 1 is returned.
// Play: https://go.dev/play/p/wadzrWr9Aer
func ProductBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// ProductByErr summarizes the values in a collection using the given return value from the iteration function.
// If the iteratee returns an error, iteration stops and the error is returned.
// If collection is empty 1 and nil error are returned.
func ProductByErr[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// Mean calculates the mean of a collection of numbers.
// Play: https://go.dev/play/p/tPURSuteUsP
func Mean[T constraints.Float | constraints.Integer](collection []T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MeanBy calculates the mean of a collection of numbers using the given return value from the iteration function.
// Play: https://go.dev/play/p/j7TsVwBOZ7P
func MeanBy[T any, R constraints.Float | constraints.Integer](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByErr calculates the mean of a collection of numbers using the given return value from the iteration function.
// If the iteratee returns an error, iteration stops and the error is returned.
// If collection is empty 0 and nil error are returned.
func MeanByErr[T any, R constraints.Float | constraints.Integer](collection []T, iteratee func(item T) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// Mode returns the mode (most frequent value) of a collection.
// If multiple values have the same highest frequency, then multiple values are returned.
// If the collection is empty, then the zero value of T is returned.
// Play: https://go.dev/play/p/PbiviqnV5zX
func Mode[T constraints.Integer | constraints.Float](collection []T) []T {
	_ = "STUB: not implemented"
	return nil
}
