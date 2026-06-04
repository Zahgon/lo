package lo

// Keys creates a slice of the map keys.
// Play: https://go.dev/play/p/Uu11fHASqrU
func Keys[K comparable, V any](in ...map[K]V) []K { _ = "STUB: not implemented"; return nil }

// UniqKeys creates a slice of unique keys in the map.
// Play: https://go.dev/play/p/TPKAb6ILdHk
func UniqKeys[K comparable, V any](in ...map[K]V) []K { _ = "STUB: not implemented"; return nil }

// HasKey returns whether the given key exists.
// Play: https://go.dev/play/p/aVwubIvECqS
func HasKey[K comparable, V any](in map[K]V, key K) bool { _ = "STUB: not implemented"; return false }

// Values creates a slice of the map values.
// Play: https://go.dev/play/p/nnRTQkzQfF6
func Values[K comparable, V any](in ...map[K]V) []V { _ = "STUB: not implemented"; return nil }

// UniqValues creates a slice of unique values in the map.
// Play: https://go.dev/play/p/nf6bXMh7rM3
func UniqValues[K, V comparable](in ...map[K]V) []V { _ = "STUB: not implemented"; return nil }

// ValueOr returns the value of the given key or the fallback value if the key is not present.
// Play: https://go.dev/play/p/bAq9mHErB4V
func ValueOr[K comparable, V any](in map[K]V, key K, fallback V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

// PickBy returns same map type filtered by given predicate.
// Play: https://go.dev/play/p/kdg8GR_QMmf
func PickBy[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) bool) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// PickByErr returns same map type filtered by given predicate.
// It returns the first error returned by the predicate.
func PickByErr[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) (bool, error)) (Map, error) {
	_ = "STUB: not implemented"
	return *new(Map), nil
}

// PickByKeys returns same map type filtered by given keys.
// Play: https://go.dev/play/p/R1imbuci9qU
func PickByKeys[K comparable, V any, Map ~map[K]V](in Map, keys []K) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// PickByValues returns same map type filtered by given values.
// Play: https://go.dev/play/p/-_PPkSbO1Kc
func PickByValues[K, V comparable, Map ~map[K]V](in Map, values []V) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// OmitBy returns same map type filtered by given predicate.
// Play: https://go.dev/play/p/EtBsR43bdsd
func OmitBy[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) bool) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// OmitByErr returns same map type filtered by given predicate.
// It returns the first error returned by the predicate.
func OmitByErr[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) (bool, error)) (Map, error) {
	_ = "STUB: not implemented"
	return *new(Map), nil
}

// OmitByKeys returns same map type filtered by given keys.
// Play: https://go.dev/play/p/t1QjCrs-ysk
func OmitByKeys[K comparable, V any, Map ~map[K]V](in Map, keys []K) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// OmitByValues returns same map type filtered by given values.
// Play: https://go.dev/play/p/9UYZi-hrs8j
func OmitByValues[K, V comparable, Map ~map[K]V](in Map, values []V) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// Entries transforms a map into a slice of key/value pairs.
// Play: https://go.dev/play/p/_t4Xe34-Nl5
func Entries[K comparable, V any](in map[K]V) []Entry[K, V] { _ = "STUB: not implemented"; return nil }

// ToPairs transforms a map into a slice of key/value pairs.
// Alias of Entries().
// Play: https://go.dev/play/p/3Dhgx46gawJ
func ToPairs[K comparable, V any](in map[K]V) []Entry[K, V] {
	_ = "STUB: not implemented"

	// FromEntries transforms a slice of key/value pairs into a map.
	// Play: https://go.dev/play/p/oIr5KHFGCEN
	return nil
}

func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FromPairs transforms a slice of key/value pairs into a map.
// Alias of FromEntries().
// Play: https://go.dev/play/p/oIr5KHFGCEN
func FromPairs[K comparable, V any](entries []Entry[K, V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// Invert creates a map composed of the inverted keys and values. If map
// contains duplicate values, subsequent values overwrite property assignments
// of previous values.
// Play: https://go.dev/play/p/rFQ4rak6iA1
func Invert[K, V comparable](in map[K]V) map[V]K { _ = "STUB: not implemented"; return nil }

// Assign merges multiple maps from left to right.
// Play: https://go.dev/play/p/VhwfJOyxf5o
func Assign[K comparable, V any, Map ~map[K]V](maps ...Map) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

// ChunkEntries splits a map into a slice of elements in groups of length equal to its size. If the map cannot be split evenly,
// the final chunk will contain the remaining elements.
// Play: https://go.dev/play/p/X_YQL6mmoD-
func ChunkEntries[K comparable, V any](m map[K]V, size int) []map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// MapKeys manipulates map keys and transforms it to a map of another type.
// Play: https://go.dev/play/p/9_4WPIqOetJ
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) R) map[R]V {
	_ = "STUB: not implemented"
	return nil
}

// MapKeysErr manipulates map keys and transforms it to a map of another type.
// It returns the first error returned by the iteratee.
func MapKeysErr[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) (R, error)) (map[R]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MapValues manipulates map values and transforms it to a map of another type.
// Play: https://go.dev/play/p/T_8xAfvcf0W
func MapValues[K comparable, V, R any](in map[K]V, iteratee func(value V, key K) R) map[K]R {
	_ = "STUB: not implemented"
	return nil
}

// MapValuesErr manipulates map values and transforms it to a map of another type.
// It returns the first error returned by the iteratee.
func MapValuesErr[K comparable, V, R any](in map[K]V, iteratee func(value V, key K) (R, error)) (map[K]R, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MapEntries manipulates map entries and transforms it to a map of another type.
// Play: https://go.dev/play/p/VuvNQzxKimT
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2 {
	_ = "STUB: not implemented"
	return nil
}

// MapEntriesErr manipulates map entries and transforms it to a map of another type.
// It returns the first error returned by the iteratee.
func MapEntriesErr[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2, error)) (map[K2]V2, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MapToSlice transforms a map into a slice based on specified iteratee.
// Play: https://go.dev/play/p/4f5hbHyMf5h
func MapToSlice[K comparable, V, R any](in map[K]V, iteratee func(key K, value V) R) []R {
	_ = "STUB: not implemented"
	return nil
}

// MapToSliceErr transforms a map into a slice based on specified iteratee.
// It returns the first error returned by the iteratee.
func MapToSliceErr[K comparable, V, R any](in map[K]V, iteratee func(key K, value V) (R, error)) ([]R, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilterMapToSlice transforms a map into a slice based on specified iteratee.
// The iteratee returns a value and a boolean. If the boolean is true, the value is added to the result slice.
// If the boolean is false, the value is not added to the result slice.
// The order of the keys in the input map is not specified and the order of the keys in the output slice is not guaranteed.
// Play: https://go.dev/play/p/jgsD_Kil9pV
func FilterMapToSlice[K comparable, V, R any](in map[K]V, iteratee func(key K, value V) (R, bool)) []R {
	_ = "STUB: not implemented"
	return nil
}

// FilterMapToSliceErr transforms a map into a slice based on specified iteratee.
// The iteratee returns a value, a boolean, and an error. If the boolean is true, the value is added to the result slice.
// If the boolean is false, the value is not added to the result slice.
// If an error is returned, iteration stops immediately and returns the error.
// The order of the keys in the input map is not specified and the order of the keys in the output slice is not guaranteed.
// Play: https://go.dev/play/p/YjFEORLBWvk
func FilterMapToSliceErr[K comparable, V, R any](in map[K]V, iteratee func(key K, value V) (R, bool, error)) ([]R, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilterKeys transforms a map into a slice based on predicate returns true for specific elements.
// It is a mix of lo.Filter() and lo.Keys().
// Play: https://go.dev/play/p/OFlKXlPrBAe
func FilterKeys[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) []K {
	_ = "STUB: not implemented"
	return nil
}

// FilterValues transforms a map into a slice based on predicate returns true for specific elements.
// It is a mix of lo.Filter() and lo.Values().
// Play: https://go.dev/play/p/YVD5r_h-LX-
func FilterValues[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) []V {
	_ = "STUB: not implemented"
	return nil
}

// FilterKeysErr transforms a map into a slice of keys based on predicate that can return an error.
// It is a mix of lo.Filter() and lo.Keys() with error handling.
// If the predicate returns true, the key is added to the result slice.
// If the predicate returns an error, iteration stops immediately and returns the error.
// The order of the keys in the input map is not specified.
// Play: https://go.dev/play/p/j2gUQzCTu4t
func FilterKeysErr[K comparable, V any](in map[K]V, predicate func(key K, value V) (bool, error)) ([]K, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilterValuesErr transforms a map into a slice of values based on predicate that can return an error.
// It is a mix of lo.Filter() and lo.Values() with error handling.
// If the predicate returns true, the value is added to the result slice.
// If the predicate returns an error, iteration stops immediately and returns the error.
// The order of the keys in the input map is not specified.
// Play: https://go.dev/play/p/hKvHlqLzbdE
func FilterValuesErr[K comparable, V any](in map[K]V, predicate func(key K, value V) (bool, error)) ([]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
