package lo

// IsNil checks if a value is nil or if it's a reference type with a nil underlying value.
// Play: https://go.dev/play/p/P2sD0PMXw4F
func IsNil(x any) bool { _ = "STUB: not implemented"; return false }

//nolint:exhaustive

// IsNotNil checks if a value is not nil or if it's not a reference type with a nil underlying value.
// Play: https://go.dev/play/p/P2sD0PMXw4F
func IsNotNil(x any) bool {
	_ = "STUB: not implemented"

	// ToPtr returns a pointer copy of value.
	// Play: https://go.dev/play/p/P2sD0PMXw4F
	return false
}

func ToPtr[T any](x T) *T {
	_ = "STUB: not implemented"

	// Nil returns a nil pointer of type.
	// Play: https://go.dev/play/p/P2sD0PMXw4F
	return nil
}

func Nil[T any]() *T {
	_ = "STUB: not implemented"

	// EmptyableToPtr returns a pointer copy of value if it's nonzero.
	// Otherwise, returns nil pointer.
	// Play: https://go.dev/play/p/P2sD0PMXw4F
	return nil
}

func EmptyableToPtr[T any](x T) *T {
	_ = "STUB: not implemented"
	// 🤮
	return nil
}

// FromPtr returns the pointer value or empty.
// Play: https://go.dev/play/p/mhD9CwO3X0m
func FromPtr[T any](x *T) T { _ = "STUB: not implemented"; return *new(T) }

// FromPtrOr returns the pointer value or the fallback value.
// Play: https://go.dev/play/p/mhD9CwO3X0m
func FromPtrOr[T any](x *T, fallback T) T { _ = "STUB: not implemented"; return *new(T) }

// ToSlicePtr returns a slice of pointers to each value.
// Play: https://go.dev/play/p/P2sD0PMXw4F
func ToSlicePtr[T any](collection []T) []*T { _ = "STUB: not implemented"; return nil }

// FromSlicePtr returns a slice with the pointer values.
// Returns a zero value in case of a nil pointer element.
// Play: https://go.dev/play/p/lbunFvzlUDX
func FromSlicePtr[T any](collection []*T) []T { _ = "STUB: not implemented"; return nil }

// FromSlicePtrOr returns a slice with the pointer values or the fallback value.
// Play: https://go.dev/play/p/lbunFvzlUDX
func FromSlicePtrOr[T any](collection []*T, fallback T) []T { _ = "STUB: not implemented"; return nil }

// ToAnySlice returns a slice with all elements mapped to `any` type.
// Play: https://go.dev/play/p/P2sD0PMXw4F
func ToAnySlice[T any](collection []T) []any { _ = "STUB: not implemented"; return nil }

// FromAnySlice returns a slice with all elements mapped to a type.
// Returns false in case of type conversion failure.
// Play: https://go.dev/play/p/P2sD0PMXw4F
func FromAnySlice[T any](in []any) ([]T, bool) { _ = "STUB: not implemented"; return nil, false }

// Empty returns the zero value (https://go.dev/ref/spec#The_zero_value).
// Play: https://go.dev/play/p/P2sD0PMXw4F
func Empty[T any]() T {
	_ = "STUB: not implemented"
	return *

	// IsEmpty returns true if argument is a zero value.
	// Play: https://go.dev/play/p/P2sD0PMXw4F
	new(T)
}

func IsEmpty[T comparable](v T) bool { _ = "STUB: not implemented"; return false }

// IsNotEmpty returns true if argument is not a zero value.
// Play: https://go.dev/play/p/P2sD0PMXw4F
func IsNotEmpty[T comparable](v T) bool { _ = "STUB: not implemented"; return false }

// Coalesce returns the first non-empty arguments. Arguments must be comparable.
// Play: https://go.dev/play/p/Gyo9otyvFHH
func Coalesce[T comparable](values ...T) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// CoalesceOrEmpty returns the first non-empty arguments. Arguments must be comparable.
// Play: https://go.dev/play/p/Gyo9otyvFHH
func CoalesceOrEmpty[T comparable](v ...T) T { _ = "STUB: not implemented"; return *new(T) }

// CoalesceSlice returns the first non-zero slice.
// Play: https://go.dev/play/p/Gyo9otyvFHH
func CoalesceSlice[T any](v ...[]T) ([]T, bool) { _ = "STUB: not implemented"; return nil, false }

// CoalesceSliceOrEmpty returns the first non-zero slice.
// Play: https://go.dev/play/p/Gyo9otyvFHH
func CoalesceSliceOrEmpty[T any](v ...[]T) []T { _ = "STUB: not implemented"; return nil }

// CoalesceMap returns the first non-zero map.
// Play: https://go.dev/play/p/Gyo9otyvFHH
func CoalesceMap[K comparable, V any](v ...map[K]V) (map[K]V, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// CoalesceMapOrEmpty returns the first non-zero map.
// Play: https://go.dev/play/p/Gyo9otyvFHH
func CoalesceMapOrEmpty[K comparable, V any](v ...map[K]V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}
