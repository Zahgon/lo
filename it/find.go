//go:build go1.23

package it

import (
	"iter"
	"time"

	"github.com/samber/lo/internal/constraints"
)

// IndexOf returns the index at which the first occurrence of a value is found in a sequence or -1
// if the value cannot be found.
// Will iterate through the entire sequence if element is not found.
// Play: https://go.dev/play/p/1OZHU2yfb-m
func IndexOf[T comparable](collection iter.Seq[T], element T) int {
	_ = "STUB: not implemented"
	return 0
}

// LastIndexOf returns the index at which the last occurrence of a value is found in a sequence or -1
// if the value cannot be found.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/QPATR3VC5wT
func LastIndexOf[T comparable](collection iter.Seq[T], element T) int {
	_ = "STUB: not implemented"
	return 0
}

// HasPrefix returns true if the collection has the prefix.
// Will iterate at most the size of prefix.
// Play: https://go.dev/play/p/Fyj6uq-G5IH
func HasPrefix[T comparable](collection iter.Seq[T], prefix ...T) bool {
	_ = "STUB: not implemented"
	return false
}

// HasSuffix returns true if the collection has the suffix.
// Will iterate through the entire sequence and allocate a slice the size of suffix.
// Play: https://go.dev/play/p/r6bF9Rmq5S0
func HasSuffix[T comparable](collection iter.Seq[T], suffix ...T) bool {
	_ = "STUB: not implemented"
	return false
}

// Find searches for an element in a sequence based on a predicate. Returns element and true if element was found.
// Will iterate through the entire sequence if predicate never returns true.
// Play: https://go.dev/play/p/4w28pF_l58a
func Find[T any](collection iter.Seq[T], predicate func(item T) bool) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// FindIndexOf searches for an element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Will iterate through the entire sequence if predicate never returns true.
// Play: https://go.dev/play/p/ihchBAEkhXO
func FindIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool) {
	_ = "STUB: not implemented"
	return *new(T), 0, false
}

// FindLastIndexOf searches for the last element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/ezz6hXaC4Md
func FindLastIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool) {
	_ = "STUB: not implemented"
	return *new(T), 0, false
}

// FindOrElse searches for an element in a sequence based on a predicate. Returns the element if found or a given fallback value otherwise.
// Will iterate through the entire sequence if predicate never returns true.
// Play: https://go.dev/play/p/1harvaiGMfI
func FindOrElse[T any](collection iter.Seq[T], fallback T, predicate func(item T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// FindUniques returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the collection.
// Will iterate through the entire sequence before yielding and allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/O8dwXEbT56F
func FindUniques[T comparable, I ~func(func(T) bool)](collection I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// FindUniquesBy returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will iterate through the entire sequence before yielding and allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/TiwGIzeDuML
func FindUniquesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// FindDuplicates returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the collection.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/dw-VLQXKijT
func FindDuplicates[T comparable, I ~func(func(T) bool)](collection I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// FindDuplicatesBy returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/tm1tZdC93OH
func FindDuplicatesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Min search the minimum value of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/0VihyYEaM-M
func Min[T constraints.Ordered](collection iter.Seq[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MinIndex search the minimum value of a collection and the index of the minimum value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/70ncPxECj6l
func MinIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// MinBy search the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/J5koo8khN-g
func MinBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MinIndexBy search the minimum value of a collection using the given comparison function and the index of the minimum value.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/blldzWJpqVa
func MinIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// Earliest search the minimum time.Time of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/fI6_S10H7Py
func Earliest(times iter.Seq[time.Time]) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// EarliestBy search the minimum time.Time of a collection using the given transform function.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/y_Pf3Jmw-B4
func EarliestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Max searches the maximum value of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/C2ZtW2bsBZ6
func Max[T constraints.Ordered](collection iter.Seq[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MaxIndex searches the maximum value of a collection and the index of the maximum value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/zeu2wUvhl5e
func MaxIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// MaxBy search the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/yBhXFJb5oxC
func MaxBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// MaxIndexBy search the maximum value of a collection using the given comparison function and the index of the maximum value.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/MXyE6BTILjx
func MaxIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int) {
	_ = "STUB: not implemented"
	return *new(T), 0
}

// Latest search the maximum time.Time of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/r5Yq6ATSHoH
func Latest(times iter.Seq[time.Time]) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// LatestBy search the maximum time.Time of a collection using the given transform function.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/o_daRzHrDUU
func LatestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// First returns the first element of a collection and check for availability of the first element.
// Will iterate at most once.
// Play: https://go.dev/play/p/EhNyrc8jPfY
func First[T any](collection iter.Seq[T]) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// FirstOrEmpty returns the first element of a collection or zero value if empty.
// Will iterate at most once.
// Play: https://go.dev/play/p/NTUTgPCfevx
func FirstOrEmpty[T any](collection iter.Seq[T]) T { _ = "STUB: not implemented"; return *new(T) }

// FirstOr returns the first element of a collection or the fallback value if empty.
// Will iterate at most once.
// Play: https://go.dev/play/p/wGFXI5NHkE2
func FirstOr[T any](collection iter.Seq[T], fallback T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Last returns the last element of a collection or error if empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/eGZV-sSmn_Q
func Last[T any](collection iter.Seq[T]) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// LastOrEmpty returns the last element of a collection or zero value if empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/teODFK4YqM4
func LastOrEmpty[T any](collection iter.Seq[T]) T { _ = "STUB: not implemented"; return *new(T) }

// LastOr returns the last element of a collection or the fallback value if empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/HNubjW2Mrxs
func LastOr[T any](collection iter.Seq[T], fallback T) T { _ = "STUB: not implemented"; return *new(T) }

// Nth returns the element at index `nth` of collection. An error is returned when nth is out of bounds.
// Will iterate n times through the sequence.
// Play: https://go.dev/play/p/FqgCobsKqva
func Nth[T any, N constraints.Integer](collection iter.Seq[T], nth N) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func seqNth[T any, N constraints.Integer](collection iter.Seq[T], nth N) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// NthOr returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the fallback value instead of an error.
// Will iterate n times through the sequence.
// Play: https://go.dev/play/p/MNweuhpy4Ym
func NthOr[T any, N constraints.Integer](collection iter.Seq[T], nth N, fallback T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// NthOrEmpty returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the zero value (empty value) for that type.
// Will iterate n times through the sequence.
// Play: https://go.dev/play/p/pC0Zhu3EUhe
func NthOrEmpty[T any, N constraints.Integer](collection iter.Seq[T], nth N) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Sample returns a random item from collection.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/YDJVX0UXYDi
func Sample[T any](collection iter.Seq[T]) T { _ = "STUB: not implemented"; return *new(T) }

// SampleBy returns a random item from collection, using randomIntGenerator as the random index generator.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/QQooySxORib
func SampleBy[T any](collection iter.Seq[T], randomIntGenerator func(int) int) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Samples returns N random unique items from collection.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/GUTFx9LQ8pP
func Samples[T any, I ~func(func(T) bool)](collection I, count int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// SamplesBy returns N random unique items from collection, using randomIntGenerator as the random index generator.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/fX2FEtixrVG
func SamplesBy[T any, I ~func(func(T) bool)](collection I, count int, randomIntGenerator func(int) int) I {
	_ = "STUB: not implemented"
	return *new(I)
}
