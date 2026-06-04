package mutable

// Filter overwrites collection's underlying array with the elements that satisfy
// predicate, in their original order, and returns a slice header whose length is the
// number kept. Use FilterI if you need the element's index.
//
// The caller's original slice variable still has its original length: only the returned
// slice has the new shorter length. Anything past that point in the original array is
// leftover from before the call (often duplicates of the last kept element). Either
// assign the result back, or re-slice with the returned length, if you want to discard
// the leftover.
// Play: https://go.dev/play/p/0jY3Z0B7O_5
func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// FilterI is like Filter but passes the element's index to predicate as well.
// See Filter for the in-place semantics (the caller's original slice keeps its length;
// only the returned slice has the new shorter length).
func FilterI[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// Map applies transform to each element of collection in place. Use MapI if you need
// the element's index.
// Play: https://go.dev/play/p/0jY3Z0B7O_5
func Map[T any, Slice ~[]T](collection Slice, transform func(item T) T) {
	_ = "STUB: not implemented"
	return
}

// MapI is like Map but passes the element's index to transform as well.
func MapI[T any, Slice ~[]T](collection Slice, transform func(item T, index int) T) {
	_ = "STUB: not implemented"
	return
}

// Shuffle returns a slice of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play: https://go.dev/play/p/2xb3WdLjeSJ
func Shuffle[T any, Slice ~[]T](collection Slice) { _ = "STUB: not implemented"; return }

// Reverse reverses a slice so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/O-M5pmCRgzV
func Reverse[T any, Slice ~[]T](collection Slice) { _ = "STUB: not implemented"; return }
