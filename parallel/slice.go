package parallel

// Map manipulates a slice and transforms it to a slice of another type.
// `transform` is called in parallel. Result keep the same order.
// Play: https://go.dev/play/p/sCJaB3quRMC
func Map[T, R any](collection []T, transform func(item T, index int) R) []R {
	_ = "STUB: not implemented"
	return nil
}

// ForEach iterates over elements of collection and invokes callback for each element.
// `iteratee` is called in parallel.
// Play: https://go.dev/play/p/sCJaB3quRMC
func ForEach[T any](collection []T, callback func(item T, index int)) {
	_ = "STUB: not implemented"
	return
}

// Times invokes the iteratee n times, returning a slice of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is called in parallel.
// Play: https://go.dev/play/p/ZNnWNcJ4Au-
func Times[T any](count int, iteratee func(index int) T) []T { _ = "STUB: not implemented"; return nil }

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// The order of grouped values is determined by the order they occur in the collection.
// `iteratee` is called in parallel.
// Play: https://go.dev/play/p/EkyvA0gw4dj
func GroupBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) map[U]Slice {
	_ = "STUB: not implemented"
	return nil
}

// PartitionBy returns a slice of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// The order of groups is determined by their first appearance in the collection.
// `iteratee` is called in parallel.
// Play: https://go.dev/play/p/GwBQdMgx2nC
func PartitionBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K) []Slice {
	_ = "STUB: not implemented"
	return nil
}
