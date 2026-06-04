package lo

import (
	"time"

	"github.com/samber/lo/internal/constraints"
)

// IndexOf returns the index at which the first occurrence of a value is found in a slice or -1
// if the value cannot be found.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func IndexOf[T comparable](collection []T, element T) int { _ = "STUB: not implemented"; return 0 }

// LastIndexOf returns the index at which the last occurrence of a value is found in a slice or -1
// if the value cannot be found.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func LastIndexOf[T comparable](collection []T, element T) int { _ = "STUB: not implemented"; return 0 }

// HasPrefix returns true if the collection has the prefix.
// Play: https://go.dev/play/p/SrljzVDpMQM
func HasPrefix[T comparable](collection, prefix []T) bool { _ = "STUB: not implemented"; return false }

// HasSuffix returns true if the collection has the suffix.
// Play: https://go.dev/play/p/bJeLetQNAON
func HasSuffix[T comparable](collection, suffix []T) bool { _ = "STUB: not implemented"; return false }

// Find searches for an element in a slice based on a predicate. Returns element and true if element was found.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// FindErr searches for an element in a slice based on a predicate that can return an error.
// Returns the element and nil error if the element is found.
// Returns zero value and nil error if the element is not found.
// If the predicate returns an error, iteration stops immediately and returns zero value and the error.
// Play: https://go.dev/play/p/XK-qtpQWXJ9
func FindErr[T any](collection []T, predicate func(item T) (bool, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// FindIndexOf searches for an element in a slice based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Play: https://go.dev/play/p/XWSEM4Ic_t0
func FindIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	_ = "STUB: not implemented"
	return *new(T), 0, false
}

// FindLastIndexOf searches for the last element in a slice based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Play: https://go.dev/play/p/2VhPMiQvX-D
func FindLastIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	_ = "STUB: not implemented"
	return *new(T), 0, false
}

// FindOrElse searches for an element in a slice based on a predicate. Returns the element if found or a given fallback value otherwise.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func FindOrElse[T any](collection []T, fallback T, predicate func(item T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// FindKey returns the key of the first value matching.
// Play: https://go.dev/play/p/Bg0w1VDPYXx
func FindKey[K, V comparable](object map[K]V, value V) (K, bool) {
	_ = "STUB: not implemented"
	return *new(K), false
}

// FindKeyBy returns the key of the first element predicate returns true for.
// Play: https://go.dev/play/p/9IbiPElcyo8
func FindKeyBy[K comparable, V any](object map[K]V, predicate func(key K, value V) bool) (K, bool) {
	_ = "STUB: not implemented"
	return *new(K), false
}

// FindUniques returns a slice with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the collection.
// Play: https://go.dev/play/p/NV5vMK_2Z_n
func FindUniques[T comparable, Slice ~[]T](collection Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// FindUniquesBy returns a slice with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the slice. It accepts `iteratee` which is
// invoked for each element in the slice to generate the criterion by which uniqueness is computed.
// Play: https://go.dev/play/p/2vmxCs4kW_m
func FindUniquesBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// FindDuplicates returns a slice with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order they occur in the collection.
// Play: https://go.dev/play/p/muFgL_XBwoP
func FindDuplicates[T comparable, Slice ~[]T](collection Slice) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// FindDuplicatesBy returns a slice with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order they occur in the slice. It accepts `iteratee` which is
// invoked for each element in the slice to generate the criterion by which uniqueness is computed.
// Play: https://go.dev/play/p/LKdYdNHuGJG
func FindDuplicatesBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// FindDuplicatesByErr returns a slice with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order they occur in the slice. It accepts `iteratee` which is
// invoked for each element in the slice to generate the criterion by which uniqueness is computed.
// If the iteratee returns an error, iteration stops immediately and the error is returned with a nil slice.
// Play: https://go.dev/play/p/HiVILQqdFP0
func FindDuplicatesByErr[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) (U, error)) (Slice, error) {
	_ = "STUB: not implemented"
	return *new(Slice), nil
}

// First pass: identify duplicates

// Second pass: collect first occurrences of duplicates

// Min searches the minimum value of a collection.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/fJFLwpY8eMN
func Min[T constraints.Ordered](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinIndex searches the minimum value of a collection and the index of the minimum value.
// Returns (zero value, -1) when the collection is empty.
// Play: https://go.dev/play/p/RxAidik4p50
func MinIndex[T constraints.Ordered](collection []T) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// MinBy searches the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/-B1PsrHVnfx
func MinBy[T any](collection []T, less func(a, b T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MinByErr searches the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value and nil error when the collection is empty.
// If the comparison function returns an error, iteration stops and the error is returned.
// Play: https://go.dev/play/p/nvDYGS8q895
func MinByErr[T any](collection []T, less func(a, b T) (bool, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// MinIndexBy searches the minimum value of a collection using the given comparison function and the index of the minimum value.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
// Play: https://go.dev/play/p/zwwPRqWhnUY
func MinIndexBy[T any](collection []T, less func(a, b T) bool) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// MinIndexByErr searches the minimum value of a collection using the given comparison function and the index of the minimum value.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
// Comparison function can return an error to stop iteration immediately.
// Play: https://go.dev/play/p/MUqi_NvTKM1
func MinIndexByErr[T any](collection []T, less func(a, b T) (bool, error)) (T, int, error) {
	_ = "STUB: not implemented"
	return *new(T), 0, nil
}

// Earliest searches the minimum time.Time of a collection.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/pRyy0c6hsBs
func Earliest(times ...time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// EarliestBy searches the minimum time.Time of a collection using the given iteratee function.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/0XvCF6vuLXC
func EarliestBy[T any](collection []T, iteratee func(item T) time.Time) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// EarliestByErr searches the minimum time.Time of a collection using the given iteratee function.
// Returns zero value and nil error when the collection is empty.
// If the iteratee returns an error, iteration stops and the error is returned.
// Play: https://go.dev/play/p/zJUBUj7ANvq
func EarliestByErr[T any](collection []T, iteratee func(item T) (time.Time, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// Max searches the maximum value of a collection.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/wYvG8gRRFw-
func Max[T constraints.Ordered](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxIndex searches the maximum value of a collection and the index of the maximum value.
// Returns (zero value, -1) when the collection is empty.
// Play: https://go.dev/play/p/RFkB4Mzb1qt
func MaxIndex[T constraints.Ordered](collection []T) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// MaxBy searches the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value when the collection is empty.
//
// Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.
// See https://github.com/samber/lo/issues/129
//
// Play: https://go.dev/play/p/PJCc-ThrwX1
func MaxBy[T any](collection []T, greater func(a, b T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MaxByErr searches the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value and nil error when the collection is empty.
// If the comparison function returns an error, iteration stops and the error is returned.
//
// Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.
// See https://github.com/samber/lo/issues/129
//
// Play: https://go.dev/play/p/s-63-6_9zqM
func MaxByErr[T any](collection []T, greater func(a, b T) (bool, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// MaxIndexBy searches the maximum value of a collection using the given comparison function and the index of the maximum value.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
//
// Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.
// See https://github.com/samber/lo/issues/129
//
// Play: https://go.dev/play/p/5yd4W7pe2QJ
func MaxIndexBy[T any](collection []T, greater func(a, b T) bool) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// MaxIndexByErr searches the maximum value of a collection using the given comparison function and the index of the maximum value.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns (zero value, -1, nil) when the collection is empty.
// If the comparison function returns an error, iteration stops and the error is returned.
//
// Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.
// See https://github.com/samber/lo/issues/129
func MaxIndexByErr[T any](collection []T, greater func(a, b T) (bool, error)) (T, int, error) {
	_ = "STUB: not implemented"
	return *new(T), 0, nil
}

// Latest searches the maximum time.Time of a collection.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/dBfdf5s8s-Y
func Latest(times ...time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// LatestBy searches the maximum time.Time of a collection using the given iteratee function.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/p1HA8XumaMU
func LatestBy[T any](collection []T, iteratee func(item T) time.Time) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// LatestByErr searches the maximum time.Time of a collection using the given iteratee function.
// Returns zero value and nil error when the collection is empty.
// If the iteratee returns an error, iteration stops and the error is returned.
// Play: https://go.dev/play/p/WpBUptwnxuG
func LatestByErr[T any](collection []T, iteratee func(item T) (time.Time, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// First returns the first element of a collection and check for availability of the first element.
// Play: https://go.dev/play/p/94lu5X6_cbf
func First[T any](collection []T) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// FirstOrEmpty returns the first element of a collection or zero value if empty.
// Play: https://go.dev/play/p/i200n9wgrDA
func FirstOrEmpty[T any](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// FirstOr returns the first element of a collection or the fallback value if empty.
// Play: https://go.dev/play/p/x9CxQyRFXeZ
func FirstOr[T any](collection []T, fallback T) T { _ = "STUB: not implemented"; return *new(T) }

// Last returns the last element of a collection or error if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func Last[T any](collection []T) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// LastOrEmpty returns the last element of a collection or zero value if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func LastOrEmpty[T any](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// LastOr returns the last element of a collection or the fallback value if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func LastOr[T any](collection []T, fallback T) T { _ = "STUB: not implemented"; return *new(T) }

// Nth returns the element at index `nth` of collection. If `nth` is negative, the nth element
// from the end is returned. An error is returned when nth is out of slice bounds.
// Play: https://go.dev/play/p/mNFI9-kIZZ5
func Nth[T any, N constraints.Integer](collection []T, nth N) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func sliceNth[T any, N constraints.Integer](collection []T, nth N) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// NthOr returns the element at index `nth` of collection.
// If `nth` is negative, it returns the nth element from the end.
// If `nth` is out of slice bounds, it returns the fallback value instead of an error.
// Play: https://go.dev/play/p/njKcNhBBVsF
func NthOr[T any, N constraints.Integer](collection []T, nth N, fallback T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// NthOrEmpty returns the element at index `nth` of collection.
// If `nth` is negative, it returns the nth element from the end.
// If `nth` is out of slice bounds, it returns the zero value (empty value) for that type.
// Play: https://go.dev/play/p/sHoh88KWt6B
func NthOrEmpty[T any, N constraints.Integer](collection []T, nth N) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// randomIntGenerator is a function that should return a random integer in the range [0, n)
// where n is the argument passed to the randomIntGenerator.
type randomIntGenerator func(n int) int

// Sample returns a random item from collection.
// Play: https://go.dev/play/p/vCcSJbh5s6l
func Sample[T any](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SampleBy returns a random item from collection, using randomIntGenerator as the random index generator.
// Play: https://go.dev/play/p/HDmKmMgq0XN
func SampleBy[T any](collection []T, randomIntGenerator randomIntGenerator) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Samples returns N random unique items from collection.
// Play: https://go.dev/play/p/QYRD8aufD0C
func Samples[T any, Slice ~[]T](collection Slice, count int) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// SamplesBy returns N random unique items from collection, using randomIntGenerator as the random index generator.
// Play: https://go.dev/play/p/Dy9bGDhD_Gw
func SamplesBy[T any, Slice ~[]T](collection Slice, count int, randomIntGenerator randomIntGenerator) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// Removes index.
// It is faster to swap with last element and remove it.
