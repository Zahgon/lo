package lo

// T2 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T2[A, B any](a A, b B) Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

// T3 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] { _ = "STUB: not implemented"; return nil }

// T4 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// T5 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// T6 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	_ = "STUB: not implemented"
	return nil
}

// T7 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	_ = "STUB: not implemented"
	return nil
}

// T8 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	_ = "STUB: not implemented"
	return nil
}

// T9 creates a tuple from a list of values.
// Play: https://go.dev/play/p/IllL3ZO4BQm
func T9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	_ = "STUB: not implemented"
	return nil
}

// Unpack2 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack2[A, B any](tuple Tuple2[A, B]) (A, B) {
	_ = "STUB: not implemented"
	return *

	// Unpack3 returns values contained in a tuple.
	// Play: https://go.dev/play/p/xVP_k0kJ96W
	new(A), *new(B)
}

func Unpack3[A, B, C any](tuple Tuple3[A, B, C]) (A, B, C) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C)
}

// Unpack4 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack4[A, B, C, D any](tuple Tuple4[A, B, C, D]) (A, B, C, D) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D)
}

// Unpack5 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack5[A, B, C, D, E any](tuple Tuple5[A, B, C, D, E]) (A, B, C, D, E) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E)
}

// Unpack6 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack6[A, B, C, D, E, F any](tuple Tuple6[A, B, C, D, E, F]) (A, B, C, D, E, F) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F)
}

// Unpack7 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack7[A, B, C, D, E, F, G any](tuple Tuple7[A, B, C, D, E, F, G]) (A, B, C, D, E, F, G) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G)
}

// Unpack8 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack8[A, B, C, D, E, F, G, H any](tuple Tuple8[A, B, C, D, E, F, G, H]) (A, B, C, D, E, F, G, H) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H)
}

// Unpack9 returns values contained in a tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack9[A, B, C, D, E, F, G, H, I any](tuple Tuple9[A, B, C, D, E, F, G, H, I]) (A, B, C, D, E, F, G, H, I) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), *new(I)
}

// Zip2 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip2[A, B any](a []A, b []B) []Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

// Perf: separate loops per input slice improve CPU cache locality (each loop reads
// one contiguous memory region) and enable bounds-check elimination by the compiler.

// Zip3 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip3[A, B, C any](a []A, b []B, c []C) []Tuple3[A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

// Zip4 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip4[A, B, C, D any](a []A, b []B, c []C, d []D) []Tuple4[A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// Zip5 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip5[A, B, C, D, E any](a []A, b []B, c []C, d []D, e []E) []Tuple5[A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// Zip6 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip6[A, B, C, D, E, F any](a []A, b []B, c []C, d []D, e []E, f []F) []Tuple6[A, B, C, D, E, F] {
	_ = "STUB: not implemented"
	return nil
}

// Zip7 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip7[A, B, C, D, E, F, G any](a []A, b []B, c []C, d []D, e []E, f []F, g []G) []Tuple7[A, B, C, D, E, F, G] {
	_ = "STUB: not implemented"
	return nil
}

// Zip8 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip8[A, B, C, D, E, F, G, H any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H) []Tuple8[A, B, C, D, E, F, G, H] {
	_ = "STUB: not implemented"
	return nil
}

// Zip9 creates a slice of grouped elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaA6GaJTp
func Zip9[A, B, C, D, E, F, G, H, I any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, in []I) []Tuple9[A, B, C, D, E, F, G, H, I] {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy2 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/wlHur6yO8rR
func ZipBy2[A, B, Out any](a []A, b []B, iteratee func(a A, b B) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy3 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/j9maveOnSQX
func ZipBy3[A, B, C, Out any](a []A, b []B, c []C, iteratee func(a A, b B, c C) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy4 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/Y1eF2Ke0Ayz
func ZipBy4[A, B, C, D, Out any](a []A, b []B, c []C, d []D, iteratee func(a A, b B, c C, d D) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy5 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/SLynyalh5Oa
func ZipBy5[A, B, C, D, E, Out any](a []A, b []B, c []C, d []D, e []E, iteratee func(a A, b B, c C, d D, e E) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy6 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/IK6KVgw9e-S
func ZipBy6[A, B, C, D, E, F, Out any](a []A, b []B, c []C, d []D, e []E, f []F, iteratee func(a A, b B, c C, d D, e E, f F) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy7 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/4uW6a2vXh8w
func ZipBy7[A, B, C, D, E, F, G, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, iteratee func(a A, b B, c C, d D, e E, f F, g G) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy8 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/tk8xW7XzY4v
func ZipBy8[A, B, C, D, E, F, G, H, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, iteratee func(a A, b B, c C, d D, e E, f F, g G, h H) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipBy9 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/VGqjDmQ9YqX
func ZipBy9[A, B, C, D, E, F, G, H, I, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, iteratee func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// ZipByErr2 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr2[A, B, Out any](a []A, b []B, iteratee func(a A, b B) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr3 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr3[A, B, C, Out any](a []A, b []B, c []C, iteratee func(a A, b B, c C) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr4 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr4[A, B, C, D, Out any](a []A, b []B, c []C, d []D, iteratee func(a A, b B, c C, d D) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr5 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr5[A, B, C, D, E, Out any](a []A, b []B, c []C, d []D, e []E, iteratee func(a A, b B, c C, d D, e E) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr6 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr6[A, B, C, D, E, F, Out any](a []A, b []B, c []C, d []D, e []E, f []F, iteratee func(a A, b B, c C, d D, e E, f F) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr7 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr7[A, B, C, D, E, F, G, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, iteratee func(a A, b B, c C, d D, e E, f F, g G) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr8 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr8[A, B, C, D, E, F, G, H, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, iteratee func(a A, b B, c C, d D, e E, f F, g G, h H) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZipByErr9 creates a slice of transformed elements, the first of which contains the first elements
// of the given slices, the second of which contains the second elements of the given slices, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
// It returns the first error returned by the iteratee.
func ZipByErr9[A, B, C, D, E, F, G, H, I, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, iteratee func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unzip2 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/K-vG9tyD3Kf
func Unzip2[A, B any](tuples []Tuple2[A, B]) ([]A, []B) { _ = "STUB: not implemented"; return nil, nil }

// Unzip3 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip3[A, B, C any](tuples []Tuple3[A, B, C]) ([]A, []B, []C) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Unzip4 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip4[A, B, C, D any](tuples []Tuple4[A, B, C, D]) ([]A, []B, []C, []D) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Unzip5 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip5[A, B, C, D, E any](tuples []Tuple5[A, B, C, D, E]) ([]A, []B, []C, []D, []E) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

// Unzip6 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip6[A, B, C, D, E, F any](tuples []Tuple6[A, B, C, D, E, F]) ([]A, []B, []C, []D, []E, []F) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil
}

// Unzip7 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip7[A, B, C, D, E, F, G any](tuples []Tuple7[A, B, C, D, E, F, G]) ([]A, []B, []C, []D, []E, []F, []G) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil
}

// Unzip8 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip8[A, B, C, D, E, F, G, H any](tuples []Tuple8[A, B, C, D, E, F, G, H]) ([]A, []B, []C, []D, []E, []F, []G, []H) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil
}

// Unzip9 accepts a slice of grouped elements and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/ciHugugvaAW
func Unzip9[A, B, C, D, E, F, G, H, I any](tuples []Tuple9[A, B, C, D, E, F, G, H, I]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil, nil
}

// UnzipBy2 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/tN8yqaRZz0r
func UnzipBy2[In, A, B any](items []In, iteratee func(In) (a A, b B)) ([]A, []B) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnzipBy3 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/36ITO2DlQq1
func UnzipBy3[In, A, B, C any](items []In, iteratee func(In) (a A, b B, c C)) ([]A, []B, []C) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// UnzipBy4 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/zJ6qY1dD1rL
func UnzipBy4[In, A, B, C, D any](items []In, iteratee func(In) (a A, b B, c C, d D)) ([]A, []B, []C, []D) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// UnzipBy5 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/3f7jKkV9xZt
func UnzipBy5[In, A, B, C, D, E any](items []In, iteratee func(In) (a A, b B, c C, d D, e E)) ([]A, []B, []C, []D, []E) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

// UnzipBy6 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/8Y1b7tKu2pL
func UnzipBy6[In, A, B, C, D, E, F any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F)) ([]A, []B, []C, []D, []E, []F) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil
}

// UnzipBy7 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/7j1kLmVn3pM
func UnzipBy7[In, A, B, C, D, E, F, G any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, g G)) ([]A, []B, []C, []D, []E, []F, []G) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil
}

// UnzipBy8 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/1n2k3L4m5N6
func UnzipBy8[In, A, B, C, D, E, F, G, H any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, g G, h H)) ([]A, []B, []C, []D, []E, []F, []G, []H) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil
}

// UnzipBy9 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// Play: https://go.dev/play/p/7o8p9q0r1s2
func UnzipBy9[In, A, B, C, D, E, F, G, H, I any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, g G, h H, i I)) ([]A, []B, []C, []D, []E, []F, []G, []H, []I) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil, nil
}

// UnzipByErr2 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
// Play: https://go.dev/play/p/G2pyXQa1SUD
func UnzipByErr2[In, A, B any](items []In, iteratee func(In) (a A, b B, err error)) ([]A, []B, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// UnzipByErr3 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr3[In, A, B, C any](items []In, iteratee func(In) (a A, b B, c C, err error)) ([]A, []B, []C, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// UnzipByErr4 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr4[In, A, B, C, D any](items []In, iteratee func(In) (a A, b B, c C, d D, err error)) ([]A, []B, []C, []D, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

// UnzipByErr5 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr5[In, A, B, C, D, E any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, err error)) ([]A, []B, []C, []D, []E, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil
}

// UnzipByErr6 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr6[In, A, B, C, D, E, F any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, err error)) ([]A, []B, []C, []D, []E, []F, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil
}

// UnzipByErr7 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr7[In, A, B, C, D, E, F, G any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, g G, err error)) ([]A, []B, []C, []D, []E, []F, []G, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil
}

// UnzipByErr8 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr8[In, A, B, C, D, E, F, G, H any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, g G, h H, err error)) ([]A, []B, []C, []D, []E, []F, []G, []H, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil, nil
}

// UnzipByErr9 iterates over a collection and creates a slice regrouping the elements
// to their pre-zip configuration.
// It returns the first error returned by the iteratee.
func UnzipByErr9[In, A, B, C, D, E, F, G, H, I any](items []In, iteratee func(In) (a A, b B, c C, d D, e E, f F, g G, h H, i I, err error)) ([]A, []B, []C, []D, []E, []F, []G, []H, []I, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
}

// CrossJoin2 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/LSRL5DmdPag
func CrossJoin2[A, B any](listA []A, listB []B) []Tuple2[A, B] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin3 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/2WGeHyJj4fK
func CrossJoin3[A, B, C any](listA []A, listB []B, listC []C) []Tuple3[A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin4 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/6XhKjLmMnNp
func CrossJoin4[A, B, C, D any](listA []A, listB []B, listC []C, listD []D) []Tuple4[A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin5 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/7oPqRsTuVwX
func CrossJoin5[A, B, C, D, E any](listA []A, listB []B, listC []C, listD []D, listE []E) []Tuple5[A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin6 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/8yZ1aB2cD3e
func CrossJoin6[A, B, C, D, E, F any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F) []Tuple6[A, B, C, D, E, F] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin7 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/9f4g5h6i7j8
func CrossJoin7[A, B, C, D, E, F, G any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G) []Tuple7[A, B, C, D, E, F, G] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin8 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/0k1l2m3n4o5
func CrossJoin8[A, B, C, D, E, F, G, H any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H) []Tuple8[A, B, C, D, E, F, G, H] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoin9 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/6p7q8r9s0t1
func CrossJoin9[A, B, C, D, E, F, G, H, I any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, listI []I) []Tuple9[A, B, C, D, E, F, G, H, I] {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy2 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/VLy8iyrPN8X
func CrossJoinBy2[A, B, Out any](listA []A, listB []B, transform func(a A, b B) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy3 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/3z4y5x6w7v8
func CrossJoinBy3[A, B, C, Out any](listA []A, listB []B, listC []C, transform func(a A, b B, c C) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy4 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/8b9c0d1e2f3
func CrossJoinBy4[A, B, C, D, Out any](listA []A, listB []B, listC []C, listD []D, transform func(a A, b B, c C, d D) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy5 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/4g5h6i7j8k9
func CrossJoinBy5[A, B, C, D, E, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, transform func(a A, b B, c C, d D, e E) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy6 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/1l2m3n4o5p6
func CrossJoinBy6[A, B, C, D, E, F, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, transform func(a A, b B, c C, d D, e E, f F) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy7 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/7q8r9s0t1u2
func CrossJoinBy7[A, B, C, D, E, F, G, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, transform func(a A, b B, c C, d D, e E, f F, g G) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy8 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/3v4w5x6y7z8
func CrossJoinBy8[A, B, C, D, E, F, G, H, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, transform func(a A, b B, c C, d D, e E, f F, g G, h H) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinBy9 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// Play: https://go.dev/play/p/9a0b1c2d3e4
func CrossJoinBy9[A, B, C, D, E, F, G, H, I, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, listI []I, transform func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Out) []Out {
	_ = "STUB: not implemented"
	return nil
}

// CrossJoinByErr2 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr2[A, B, Out any](listA []A, listB []B, transform func(a A, b B) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr3 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr3[A, B, C, Out any](listA []A, listB []B, listC []C, transform func(a A, b B, c C) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr4 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr4[A, B, C, D, Out any](listA []A, listB []B, listC []C, listD []D, transform func(a A, b B, c C, d D) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr5 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr5[A, B, C, D, E, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, transform func(a A, b B, c C, d D, e E) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr6 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr6[A, B, C, D, E, F, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, transform func(a A, b B, c C, d D, e E, f F) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr7 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr7[A, B, C, D, E, F, G, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, transform func(a A, b B, c C, d D, e E, f F, g G) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr8 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr8[A, B, C, D, E, F, G, H, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, transform func(a A, b B, c C, d D, e E, f F, g G, h H) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossJoinByErr9 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The transform function
// is used to create the output values.
// Returns an empty list if a list is empty.
// It returns the first error returned by the transform function.
func CrossJoinByErr9[A, B, C, D, E, F, G, H, I, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, listI []I, transform func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (Out, error)) ([]Out, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
