//go:build go1.23

package it

import (
	"iter"
)

// Contains returns true if an element is present in a collection.
// Will iterate through the entire sequence if element is not found.
// Play: https://go.dev/play/p/1edj7hH3TS2
func Contains[T comparable](collection iter.Seq[T], element T) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsBy returns true if predicate function return true.
// Will iterate through the entire sequence if predicate never returns true.
func ContainsBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Every returns true if all elements of a subset are contained in a collection or if the subset is empty.
// Will iterate through the entire sequence if subset elements always match.
// Play: https://go.dev/play/p/rwM9Y353aIC
func Every[T comparable](collection iter.Seq[T], subset ...T) bool {
	_ = "STUB: not implemented"
	return false
}

// EveryBy returns true if the predicate returns true for all elements in the collection or if the collection is empty.
// Will iterate through the entire sequence if predicate never returns false.
func EveryBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Some returns true if at least 1 element of a subset is contained in a collection.
// If the subset is empty Some returns false.
// Will iterate through the entire sequence if subset elements never match.
// Play: https://go.dev/play/p/KmX-fXictQl
func Some[T comparable](collection iter.Seq[T], subset ...T) bool {
	_ = "STUB: not implemented"
	return false
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
// Will iterate through the entire sequence if predicate never returns true.
func SomeBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// None returns true if no element of a subset is contained in a collection or if the subset is empty.
// Will iterate through the entire sequence if subset elements never match.
// Play: https://go.dev/play/p/L7mm5S4a8Yo
func None[T comparable](collection iter.Seq[T], subset ...T) bool {
	_ = "STUB: not implemented"
	return false
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
// Will iterate through the entire sequence if predicate never returns true.
func NoneBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Intersect returns the intersection between given collections.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/kz3cGhGZZWF
func Intersect[T comparable, I ~func(func(T) bool)](lists ...I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// IntersectBy returns the intersection between given collections using a
// custom key selector function.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/X2nEvHC-lE2
func IntersectBy[T any, K comparable, I ~func(func(T) bool)](transform func(T) K, lists ...I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Union returns all distinct elements from given collections.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/ImIoFNpSUUB
func Union[T comparable, I ~func(func(T) bool)](lists ...I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Without returns a sequence excluding all given values.
// Will allocate a map large enough to hold all distinct excludes.
// Play: https://go.dev/play/p/LbN55AVBZ7h
func Without[T comparable, I ~func(func(T) bool)](collection I, exclude ...T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// WithoutBy filters a sequence by excluding elements whose extracted keys match any in the exclude list.
// Returns a sequence containing only the elements whose keys are not in the exclude list.
// Will allocate a map large enough to hold all distinct excludes.
// Play: https://go.dev/play/p/Hm734hnLnLI
func WithoutBy[T any, K comparable, I ~func(func(T) bool)](collection I, transform func(item T) K, exclude ...K) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// WithoutNth returns a sequence excluding the nth value.
// Will allocate a map large enough to hold all distinct nths.
// Play: https://go.dev/play/p/KGE7Lpsk18P
func WithoutNth[T comparable, I ~func(func(T) bool)](collection I, nths ...int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// ElementsMatch returns true if lists contain the same set of elements (including empty set).
// If there are duplicate elements, the number of occurrences in each list should match.
// The order of elements is not checked.
// Will iterate through each sequence before returning and allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/24SGQm1yMRe
func ElementsMatch[T comparable](list1, list2 iter.Seq[T]) bool {
	_ = "STUB: not implemented"
	return false
}

// ElementsMatchBy returns true if lists contain the same set of elements' keys (including empty set).
// If there are duplicate keys, the number of occurrences in each list should match.
// The order of elements is not checked.
// Will iterate through each sequence before returning and allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/I3vFrmQo43E
func ElementsMatchBy[T any, K comparable](list1, list2 iter.Seq[T], transform func(item T) K) bool {
	_ = "STUB: not implemented"
	return false
}
