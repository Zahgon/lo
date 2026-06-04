//go:build go1.26 && goexperiment.simd && amd64

package simd

// SumInt8 sums a slice of int8 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt16 sums a slice of int16 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt32 sums a slice of int32 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64 or handle overflow detection externally.
func SumInt32[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt64 sums a slice of int64 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint8 sums a slice of uint8 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint16 sums a slice of uint16 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint32 sums a slice of uint32 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64 or handle overflow detection externally.
func SumUint32[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint64 sums a slice of uint64 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat32 sums a slice of float32 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64.
func SumFloat32[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat64 sums a slice of float64 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt8 calculates the mean of a slice of int8 using the best available SIMD instruction set.
func MeanInt8[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt16 calculates the mean of a slice of int16 using the best available SIMD instruction set.
func MeanInt16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt32 calculates the mean of a slice of int32 using the best available SIMD instruction set.
func MeanInt32[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt64 calculates the mean of a slice of int64 using the best available SIMD instruction set.
func MeanInt64[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint8 calculates the mean of a slice of uint8 using the best available SIMD instruction set.
func MeanUint8[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint16 calculates the mean of a slice of uint16 using the best available SIMD instruction set.
func MeanUint16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint32 calculates the mean of a slice of uint32 using the best available SIMD instruction set.
func MeanUint32[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint64 calculates the mean of a slice of uint64 using the best available SIMD instruction set.
func MeanUint64[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat32 calculates the mean of a slice of float32 using the best available SIMD instruction set.
func MeanFloat32[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat64 calculates the mean of a slice of float64 using the best available SIMD instruction set.
func MeanFloat64[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinInt8 finds the minimum value in a collection of int8 using the best available SIMD instruction set.
func MinInt8[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinInt16 finds the minimum value in a collection of int16 using the best available SIMD instruction set.
func MinInt16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinInt32 finds the minimum value in a collection of int32 using the best available SIMD instruction set.
func MinInt32[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinInt64 finds the minimum value in a collection of int64 using the best available SIMD instruction set.
func MinInt64[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinInt64x2 requires AVX-512 (archsimd Int64x2.Min); use scalar fallback

// MinUint8 finds the minimum value in a collection of uint8 using the best available SIMD instruction set.
func MinUint8[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinUint16 finds the minimum value in a collection of uint16 using the best available SIMD instruction set.
func MinUint16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinUint32 finds the minimum value in a collection of uint32 using the best available SIMD instruction set.
func MinUint32[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinUint64 finds the minimum value in a collection of uint64 using the best available SIMD instruction set.
func MinUint64[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinUint64x2 requires AVX-512; use scalar fallback

// MinFloat32 finds the minimum value in a collection of float32 using the best available SIMD instruction set.
func MinFloat32[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MinFloat64 finds the minimum value in a collection of float64 using the best available SIMD instruction set.
func MinFloat64[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxInt8 finds the maximum value in a collection of int8 using the best available SIMD instruction set.
func MaxInt8[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxInt16 finds the maximum value in a collection of int16 using the best available SIMD instruction set.
func MaxInt16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxInt32 finds the maximum value in a collection of int32 using the best available SIMD instruction set.
func MaxInt32[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxInt64 finds the maximum value in a collection of int64 using the best available SIMD instruction set.
func MaxInt64[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxInt64x2 requires AVX-512; use scalar fallback

// MaxUint8 finds the maximum value in a collection of uint8 using the best available SIMD instruction set.
func MaxUint8[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxUint16 finds the maximum value in a collection of uint16 using the best available SIMD instruction set.
func MaxUint16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxUint32 finds the maximum value in a collection of uint32 using the best available SIMD instruction set.
func MaxUint32[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxUint64 finds the maximum value in a collection of uint64 using the best available SIMD instruction set.
func MaxUint64[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxUint64x2 requires AVX-512; use scalar fallback

// MaxFloat32 finds the maximum value in a collection of float32 using the best available SIMD instruction set.
func MaxFloat32[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxFloat64 finds the maximum value in a collection of float64 using the best available SIMD instruction set.
func MaxFloat64[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// ClampInt8 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt8[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampInt16 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt16[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampInt32 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt32[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampInt64 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt64[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampInt64x2 requires AVX-512; use scalar fallback

// ClampUint8 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint8[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampUint16 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint16[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampUint32 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint32[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampUint64 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint64[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampUint64x2 requires AVX-512; use scalar fallback

// ClampFloat32 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampFloat32[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// ClampFloat64 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampFloat64[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// SumByInt8 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt16 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt32 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumByInt64 or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt64 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
// Play: https://go.dev/play/p/TBD
func SumByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint8 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint16 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint32 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumByUint64 or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint64 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
// Play: https://go.dev/play/p/TBD
func SumByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat32 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumByFloat64.
// Play: https://go.dev/play/p/TBD
func SumByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat64 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
// Play: https://go.dev/play/p/TBD
func SumByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt8 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt16 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt32 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt64 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint8 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint16 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint32 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint64 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat32 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat64 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}
