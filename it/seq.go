//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo"
	"github.com/samber/lo/internal/constraints"
)

// Length returns the length of collection.
// Will iterate through the entire sequence.
func Length[T any](collection iter.Seq[T]) int { _ = "STUB: not implemented"; return 0 }

// Drain consumes an entire sequence.
func Drain[T any](collection iter.Seq[T]) { _ = "STUB: not implemented"; return }

//nolint:revive

// Filter iterates over elements of collection, returning a sequence of all elements predicate returns true for.
// Play: https://go.dev/play/p/psenko2KKsX
func Filter[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// FilterI iterates over elements of collection, returning a sequence of all elements predicate returns true for.
// Play: https://go.dev/play/p/5fpdlQvdL-q
func FilterI[T any, I ~func(func(T) bool)](collection I, predicate func(item T, index int) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Map manipulates a sequence and transforms it to a sequence of another type.
// Play: https://go.dev/play/p/F_FkOP-9F9T
func Map[T, R any](collection iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// MapI manipulates a sequence and transforms it to a sequence of another type.
// Play: https://go.dev/play/p/6gqemRweL-r
func MapI[T, R any](collection iter.Seq[T], transform func(item T, index int) R) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// UniqMap manipulates a sequence and transforms it to a sequence of another type with unique values.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/7hrfnSxfL-s
func UniqMap[T any, R comparable](collection iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// UniqMapI manipulates a sequence and transforms it to a sequence of another type with unique values.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/8isgTsyfL-t
func UniqMapI[T any, R comparable](collection iter.Seq[T], transform func(item T, index int) R) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// FilterMap returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/Gxwu8j_TJuh
func FilterMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// FilterMapI returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/0XrQKOk-vw
func FilterMapI[T, R any](collection iter.Seq[T], callback func(item T, index int) (R, bool)) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMap manipulates a sequence and transforms and flattens it to a sequence of another type.
// The transform function can either return a sequence or a `nil`, and in the `nil` case
// no value is yielded.
// Play: https://go.dev/play/p/6toB9w2gpSy
func FlatMap[T, R any](collection iter.Seq[T], transform func(item T) iter.Seq[R]) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// FlatMapI manipulates a sequence and transforms and flattens it to a sequence of another type.
// The transform function can either return a sequence or a `nil`, and in the `nil` case
// no value is yielded.
// Play: https://go.dev/play/p/2ZtSMQm-xy
func FlatMapI[T, R any](collection iter.Seq[T], transform func(item T, index int) iter.Seq[R]) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/l5wWLATXclf
func Reduce[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// ReduceI reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/3AuTNRn-yz
func ReduceI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// ReduceLast is like Reduce except that it iterates over elements of collection in reverse.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/D2ZGZ2pN270
func ReduceLast[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// ReduceLastI is like Reduce except that it iterates over elements of collection in reverse.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/5CwPTPz-zb
func ReduceLastI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// ForEach iterates over elements of collection and invokes callback for each element.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/OvW9yNNYgsX
func ForEach[T any](collection iter.Seq[T], callback func(item T)) {
	_ = "STUB: not implemented"
	return
}

// ForEachI iterates over elements of collection and invokes callback for each element.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/6NhAE0-zm
func ForEachI[T any](collection iter.Seq[T], callback func(item T, index int)) {
	_ = "STUB: not implemented"
	return
}

// ForEachWhile iterates over elements of collection and invokes predicate for each element
// collection return value decide to continue or break, like do while().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/UsfPR_gmSs7
func ForEachWhile[T any](collection iter.Seq[T], predicate func(item T) bool) {
	_ = "STUB: not implemented"
	return
}

// ForEachWhileI iterates over elements of collection and invokes predicate for each element
// collection return value decide to continue or break, like do while().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/8PjCG2-zo
func ForEachWhileI[T any](collection iter.Seq[T], predicate func(item T, index int) bool) {
	_ = "STUB: not implemented"
	return
}

// Times invokes callback n times and returns a sequence of results.
// The transform is invoked with index as argument.
// Play: https://go.dev/play/p/0W4IRzQuCEc
func Times[T any](count int, callback func(index int) T) iter.Seq[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uniq returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/D-SenTW-ipj
func Uniq[T comparable, I ~func(func(T) bool)](collection I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// UniqBy returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/HKrt3AvwMTR
func UniqBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through transform.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/oRIakS89OYy
func GroupBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U][]T {
	_ = "STUB: not implemented"
	return nil
}

// GroupByMap returns an object composed of keys generated from the results of running each element of collection through transform.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/3UoHL7-zt
func GroupByMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K][]V {
	_ = "STUB: not implemented"
	return nil
}

// Chunk returns a sequence of elements split into groups of length size. If the sequence can't be split evenly,
// the final chunk will be the remaining elements.
// Play: https://go.dev/play/p/qo8esZ_L60Q
func Chunk[T any](collection iter.Seq[T], size int) iter.Seq[[]T] {
	_ = "STUB: not implemented"
	return nil
}

// Window creates a sequence of sliding windows of a given size.
// Each window overlaps with the previous one by size-1 elements.
// This is equivalent to Sliding(collection, size, 1).
// Play: https://go.dev/play/p/_1BzQYtKBhi
func Window[T any](collection iter.Seq[T], size int) iter.Seq[[]T] {
	_ = "STUB: not implemented"
	return nil
}

// Sliding creates a sequence of sliding windows of a given size with a given step.
// offset = step - size: offset == 0 means adjacent windows (no overlap/gap);
// offset < 0 means overlapping windows; offset > 0 means gaps between windows.
// Only full-size windows are yielded; a partial window at the end is not yielded.
// Play: https://go.dev/play/p/mzhO4CZeiik
func Sliding[T any](collection iter.Seq[T], size, step int) iter.Seq[[]T] {
	_ = "STUB: not implemented"
	return nil
}

// Adjacent windows: no overlap, no gap.

// Overlap: next window starts inside the current one; keep tail in buffer.

// offset > 0 (step > size)
// Gap: skip elements between windows.

// PartitionBy returns a sequence of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through transform.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/VxTx8mva28z
func PartitionBy[T any, K comparable](collection iter.Seq[T], transform func(item T) K) [][]T {
	_ = "STUB: not implemented"
	return nil
}

// Flatten returns a sequence a single level deep.
// Play: https://go.dev/play/p/CCklxuNk7Lm
func Flatten[T any, I ~func(func(T) bool)](collection []I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Concat returns a sequence of all the elements in iterators. Concat conserves the order of the elements.
// Play: https://go.dev/play/p/Fa0u7xT2JOR
func Concat[T any, I ~func(func(T) bool)](collection ...I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Interleave round-robin alternating input sequences and sequentially appending value at index into result.
// Will allocate a slice the size of collections.
// Play: https://go.dev/play/p/kNvnz4ClLgH
func Interleave[T any](collections ...iter.Seq[T]) iter.Seq[T] {
	_ = "STUB: not implemented"
	return nil
}

// Shuffle returns a sequence of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/8isgTsyfL-t
func Shuffle[T any, I ~func(func(T) bool)](collection I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Reverse reverses a sequence so that the first element becomes the last, the second element becomes the second to last, and so on.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/9jthUzgF-u
func Reverse[T any, I ~func(func(T) bool)](collection I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

//nolint:modernize // backward loop can't use slices.Backward (Go 1.24+), lo supports Go 1.18+

// Fill replaces elements of a sequence with `initial` value.
// Play: https://go.dev/play/p/mHShWq5ezMc
func Fill[T lo.Clonable[T], I ~func(func(T) bool)](collection I, initial T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Repeat builds a sequence with N copies of initial value.
// Play: https://go.dev/play/p/xs-aq0p_uDP
func Repeat[T lo.Clonable[T]](count int, initial T) iter.Seq[T] {
	_ = "STUB: not implemented"
	return nil
}

// RepeatBy builds a sequence with values returned by N calls of transform.
// Play: https://go.dev/play/p/i7BuZQBcUzZ
func RepeatBy[T any](count int, callback func(index int) T) iter.Seq[T] {
	_ = "STUB: not implemented"
	return nil
}

// KeyBy transforms a sequence to a map based on a pivot transform function.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/MMaHpzTqY0a
func KeyBy[K comparable, V any](collection iter.Seq[V], transform func(item V) K) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/ayCImLr_4im
func Associate[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// AssociateI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/5CwPTPz-zb
func AssociateI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// SeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Alias of Associate().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/6NhAE0-zm
func SeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// SeqToMapI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Alias of AssociateI().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/7OiBF1-zn
func SeqToMapI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FilterSeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// The third return value of the transform function is a boolean that indicates whether the key-value pair should be included in the map.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/8PjCG2-zo
func FilterSeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V, bool)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// FilterSeqToMapI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// The third return value of the transform function is a boolean that indicates whether the key-value pair should be included in the map.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/9QkDH3-zp
func FilterSeqToMapI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V, bool)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// Keyify returns a map with each unique element of the sequence as a key.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/aHOD29_l-rF
func Keyify[T comparable](collection iter.Seq[T]) map[T]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// Drop drops n elements from the beginning of a sequence.
// Play: https://go.dev/play/p/O1J1-uWc3z9
func Drop[T any, I ~func(func(T) bool)](collection I, n int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// DropLast drops n elements from the end of a sequence.
// Will allocate a slice of length n.
// Play: https://go.dev/play/p/-NzU5Px5Tp4
func DropLast[T any, I ~func(func(T) bool)](collection I, n int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// DropWhile drops elements from the beginning of a sequence while the predicate returns true.
// Play: https://go.dev/play/p/zSM8x08a9QD
func DropWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// DropLastWhile drops elements from the end of a sequence while the predicate returns true.
// Will allocate a slice large enough to hold the longest sequence of matching elements.
// Long input sequences of consecutive matches can cause excessive memory usage.
// Play: https://go.dev/play/p/qZ81Cq7R-Yt
func DropLastWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Take takes the first n elements from a sequence.
func Take[T any, I ~func(func(T) bool)](collection I, n int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TakeWhile takes elements from the beginning of a sequence while the predicate returns true.
func TakeWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// DropByIndex drops elements from a sequence by the index.
// Will allocate a map large enough to hold all distinct indexes.
// Play: https://go.dev/play/p/vPbrZYgiU4q
func DropByIndex[T any, I ~func(func(T) bool)](collection I, indexes ...int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TakeFilter filters elements and takes the first n elements that match the predicate.
// Equivalent to calling Take(Filter(...)), but more efficient as it stops after finding n matches.
// Play: https://go.dev/play/p/Db68Bhu4MCA
func TakeFilter[T any, I ~func(func(T) bool)](collection I, n int, predicate func(item T) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TakeFilterI filters elements and takes the first n elements that match the predicate.
// Equivalent to calling Take(FilterI(...)), but more efficient as it stops after finding n matches.
func TakeFilterI[T any, I ~func(func(T) bool)](collection I, n int, predicate func(item T, index int) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.
// Play: https://go.dev/play/p/IIQcknFhZnq
func Reject[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// RejectI is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.
// Play: https://go.dev/play/p/7YsLP1-zx
func RejectI[T any, I ~func(func(T) bool)](collection I, predicate func(item T, index int) bool) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// RejectMap is the opposite of FilterMap, this method returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/51jRHVYscgi
func RejectMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// RejectMapI is the opposite of FilterMap, this method returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/9jthUzgF-u
func RejectMapI[T, R any](collection iter.Seq[T], callback func(item T, index int) (R, bool)) iter.Seq[R] {
	_ = "STUB: not implemented"
	return nil
}

// Count counts the number of elements in the collection that equal value.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/UcJ-6cANwfY
func Count[T comparable](collection iter.Seq[T], value T) int { _ = "STUB: not implemented"; return 0 }

// CountBy counts the number of elements in the collection for which predicate is true.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/m6G0o3huCOG
func CountBy[T any](collection iter.Seq[T], predicate func(item T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

// CountValues counts the number of each element in the collection.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/PPBT4Fp-V3B
func CountValues[T comparable](collection iter.Seq[T]) map[T]int {
	_ = "STUB: not implemented"
	return nil
}

// CountValuesBy counts the number of each element returned from transform function.
// Is equivalent to chaining Map and CountValues.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/gnr_MPhYCHX
func CountValuesBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U]int {
	_ = "STUB: not implemented"
	return nil
}

// Subset returns a subset of a sequence from `offset` up to `length` elements.
// Will iterate at most offset+length times.
// Play: https://go.dev/play/p/r-6FdqOL28Z
func Subset[T any, I ~func(func(T) bool)](collection I, offset, length int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Slice returns a subset of a sequence from `start` up to, but not including `end`.
// Will iterate at most end times.
// Play: https://go.dev/play/p/jKIu1oPf5hK
func Slice[T any, I ~func(func(T) bool)](collection I, start, end int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Replace returns a sequence with the first n non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/aFXjeyf0KqV
func Replace[T comparable, I ~func(func(T) bool)](collection I, old, nEw T, n int) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// ReplaceAll returns a sequence with all non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/sOckhMvvwjc
func ReplaceAll[T comparable, I ~func(func(T) bool)](collection I, old, nEw T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// Compact returns a sequence of all non-zero elements.
// Play: https://go.dev/play/p/R_xT99w2QaU
func Compact[T comparable, I ~func(func(T) bool)](collection I) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// IsSorted checks if a sequence is sorted.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/o-BD4UOn-0U
func IsSorted[T constraints.Ordered](collection iter.Seq[T]) bool {
	_ = "STUB: not implemented"
	return false
}

// IsSortedBy checks if a sequence is sorted by transform.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/AfYOiGWa78T
func IsSortedBy[T any, K constraints.Ordered](collection iter.Seq[T], transform func(item T) K) bool {
	_ = "STUB: not implemented"
	return false
}

// Splice inserts multiple elements at index i. The helper is protected against overflow errors.
// Play: https://go.dev/play/p/TevQSvDKO_i
func Splice[T any, I ~func(func(T) bool)](collection I, index int, elements ...T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// CutPrefix returns collection without the provided leading prefix
// and reports whether it found the prefix.
// If collection doesn't start with prefix, CutPrefix returns collection, false.
// If prefix is empty, CutPrefix returns collection, true.
// Will iterate at most the size of separator before returning.
// Play: https://go.dev/play/p/bPnV39zVnAV
func CutPrefix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (after I, found bool) {
	_ = "STUB: not implemented" //nolint:gocyclo
	return *new(I), false
}

// CutSuffix returns collection without the provided ending suffix and reports
// whether it found the suffix. If collection doesn't end with suffix, CutSuffix returns collection, false.
// If suffix is empty, CutSuffix returns collection, true.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/CTRh9m1UHrZ
func CutSuffix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (before I, found bool) {
	_ = "STUB: not implemented"
	return *new(I), false
}

// Trim removes all the leading and trailing cutset from the collection.
// Will allocate a map large enough to hold all distinct cutset elements.
// Play: https://go.dev/play/p/k0VCcilk4V1
func Trim[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TrimFirst removes all the leading cutset from the collection.
// Will allocate a map large enough to hold all distinct cutset elements.
// Play: https://go.dev/play/p/4D4Ke5C5MwH
func TrimFirst[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TrimPrefix removes all the leading prefix from the collection.
// Play: https://go.dev/play/p/Pce4zSPnThY
func TrimPrefix[T comparable, I ~func(func(T) bool)](collection I, prefix []T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TrimLast removes all the trailing cutset from the collection.
// Will allocate a map large enough to hold all distinct cutset elements.
// Play: https://go.dev/play/p/GQLhnaeW0gd
func TrimLast[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// TrimSuffix removes all the trailing suffix from the collection.
// Play: https://go.dev/play/p/s9nwy9helEi
func TrimSuffix[T comparable, I ~func(func(T) bool)](collection I, suffix []T) I {
	_ = "STUB: not implemented"
	return *new(I)
}

// SeqToSeq2 converts a sequence into a sequence of key-value pairs, where the first
// element of iter.Seq2 is the index (starting from 0 and incrementing by 1 for each item).
// Play: https://go.dev/play/p/V5wL9xY8nQr
func SeqToSeq2[T any](in iter.Seq[T]) iter.Seq2[int, T] { _ = "STUB: not implemented"; return nil }

// Buffer returns a sequence of slices, each containing up to size items read from the channel.
// The last slice may be smaller if the channel closes before filling the buffer.
// Play: https://go.dev/play/p/zDZdcCA20ut
func Buffer[T any](seq iter.Seq[T], size int) iter.Seq[[]T] { _ = "STUB: not implemented"; return nil }

// keep pulling

// Buffer full, yield it

// allocate new buffer
// false = stop, true = continue

// Yield remaining partial buffer
