//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo/internal/constraints"
)

// Range creates a sequence of numbers (positive and/or negative) with given length.
// Play: https://go.dev/play/p/79QUZBa8Ukn
func Range(elementNum int) iter.Seq[int] { _ = "STUB: not implemented"; return nil }

// RangeFrom creates a sequence of numbers from start with specified length.
// Play: https://go.dev/play/p/WHP_NI5scj9
func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) iter.Seq[T] {
	_ = "STUB: not implemented"
	return nil
}

// RangeWithSteps creates a sequence of numbers (positive and/or negative) progressing from start up to, but not including end.
// step set to zero will return an empty sequence.
// Play: https://go.dev/play/p/qxm2YNLG0vT
func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) iter.Seq[T] {
	_ = "STUB: not implemented"
	return nil
}

// Sum sums the values in a collection. If collection is empty 0 is returned.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/nHbGFOEIeTa
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// SumBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 0 is returned.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/ZNiqXNMu5QP
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T], transform func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// Product gets the product of the values in a collection. If collection is empty 1 is returned.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/AOMCD1Yl5Bc
func Product[T constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// ProductBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 1 is returned.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/dgFCRJrlPHY
func ProductBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T], transform func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// Mean calculates the mean of a collection of numbers.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/Lez0CsvVRl_l
func Mean[T constraints.Float | constraints.Integer](collection iter.Seq[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MeanBy calculates the mean of a collection of numbers using the given return value from the iteration function.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/Ked4rpztH5Y
func MeanBy[T any, R constraints.Float | constraints.Integer](collection iter.Seq[T], transform func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// Mode returns the mode (most frequent value) of a collection.
// If multiple values have the same highest frequency, then multiple values are returned.
// If the collection is empty, then the zero value of T is returned.
// Will iterate through the entire sequence and allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/c_cmMMA5EhH
func Mode[T constraints.Integer | constraints.Float](collection iter.Seq[T]) []T {
	_ = "STUB: not implemented"
	return nil
}
