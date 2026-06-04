package lo

// Contains returns true if an element is present in a collection.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func Contains[T comparable](collection []T, element T) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsBy returns true if predicate function return true.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Every returns true if all elements of a subset are contained in a collection or if the subset is empty.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func Every[T comparable](collection, subset []T) bool { _ = "STUB: not implemented"; return false }

// EveryBy returns true if the predicate returns true for all elements in the collection or if the collection is empty.
// Play: https://go.dev/play/p/dn1-vhHsq9x
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Some returns true if at least 1 element of a subset is contained in a collection.
// If the subset is empty Some returns false.
// Play: https://go.dev/play/p/Lj4ceFkeT9V
func Some[T comparable](collection, subset []T) bool { _ = "STUB: not implemented"; return false }

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
// Play: https://go.dev/play/p/DXF-TORBudx
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// None returns true if no element of a subset is contained in a collection or if the subset is empty.
// Play: https://go.dev/play/p/fye7JsmxzPV
func None[T comparable](collection, subset []T) bool { _ = "STUB: not implemented"; return false }

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
// Play: https://go.dev/play/p/O64WZ32H58S
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Intersect returns the intersection between collections.
// Play: https://go.dev/play/p/uuElL9X9e58
func Intersect[T comparable, Slice ~[]T](lists ...Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// IntersectBy returns the intersection between two collections using a custom key selector function.
// Play: https://go.dev/play/p/uWF8y2-zmtf
func IntersectBy[T any, K comparable, Slice ~[]T](transform func(T) K, lists ...Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// Difference returns the difference between two collections.
// The first value is the collection of elements absent from list2.
// The second value is the collection of elements absent from list1.
// Play: https://go.dev/play/p/pKE-JgzqRpz
func Difference[T comparable, Slice ~[]T](list1, list2 Slice) (Slice, Slice) {
	_ = "STUB: not implemented"
	return *new(Slice), *new(Slice)
}

// Union returns all distinct elements from given collections.
// result returns will not change the order of elements relatively.
// Play: https://go.dev/play/p/-hsqZNTH0ej
func Union[T comparable, Slice ~[]T](lists ...Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// UnionBy is like Union except that it accepts an iteratee which is invoked for each element of each collection
// to generate the criterion by which uniqueness is computed.
// Result values are chosen from the first collection in which the value occurs.
func UnionBy[T any, V comparable, Slice ~[]T](iteratee func(item T) V, lists ...Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// UnionByErr is like UnionBy except that it accepts an iteratee which can return an error.
// It returns the first error returned by the iteratee.
func UnionByErr[T any, V comparable, Slice ~[]T](iteratee func(item T) (V, error), lists ...Slice) (Slice, error) {
	_ = "STUB: not implemented"
	return *new(Slice), nil
}

// Without returns a slice excluding all given values.
// Play: https://go.dev/play/p/PcAVtYJsEsS
func Without[T comparable, Slice ~[]T](collection Slice, exclude ...T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// WithoutBy filters a slice by excluding elements whose extracted keys match any in the exclude list.
// Returns a new slice containing only the elements whose keys are not in the exclude list.
// Play: https://go.dev/play/p/VgWJOF01NbJ
func WithoutBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K, exclude ...K) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// WithoutByErr filters a slice by excluding elements whose extracted keys match any in the exclude list.
// It returns the first error returned by the iteratee.
func WithoutByErr[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) (K, error), exclude ...K) (Slice, error) {
	_ = "STUB: not implemented"
	return *new(Slice), nil
}

// WithoutEmpty returns a slice excluding zero values.
//
// Deprecated: Use lo.Compact instead.
// Play: https://go.dev/play/p/iZvYJWuniJm
func WithoutEmpty[T comparable, Slice ~[]T](collection Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// WithoutNth returns a slice excluding the nth value.
// Play: https://go.dev/play/p/5g3F9R2H1xL
func WithoutNth[T any, Slice ~[]T](collection Slice, nths ...int) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ElementsMatch returns true if lists contain the same set of elements (including empty set).
// If there are duplicate elements, the number of occurrences in each list should match.
// The order of elements is not checked.
// Play: https://go.dev/play/p/XWSEM4Ic_t0
func ElementsMatch[T comparable, Slice ~[]T](list1, list2 Slice) bool {
	_ = "STUB: not implemented"
	return false
}

// ElementsMatchBy returns true if lists contain the same set of elements' keys (including empty set).
// If there are duplicate keys, the number of occurrences in each list should match.
// The order of elements is not checked.
// Play: https://go.dev/play/p/XWSEM4Ic_t0
func ElementsMatchBy[T any, K comparable](list1, list2 []T, iteratee func(item T) K) bool {
	_ = "STUB: not implemented"
	return false
}
