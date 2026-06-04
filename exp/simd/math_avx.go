//go:build go1.26 && goexperiment.simd && amd64

package simd

// AVX (128-bit) SIMD sum functions - 16/8/4/2 lanes

// SumInt8x16 sums a slice of int8 using AVX SIMD (Int8x16, 16 lanes).
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8x16[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt16x8 sums a slice of int16 using AVX SIMD (Int16x8, 8 lanes).
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16x8[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt32x4 sums a slice of int32 using AVX SIMD (Int32x4, 4 lanes).
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64x2 or handle overflow detection externally.
func SumInt32x4[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt64x2 sums a slice of int64 using AVX SIMD (Int64x2, 2 lanes).
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64x2[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint8x16 sums a slice of uint8 using AVX SIMD (Uint8x16, 16 lanes).
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8x16[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint16x8 sums a slice of uint16 using AVX SIMD (Uint16x8, 8 lanes).
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16x8[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint32x4 sums a slice of uint32 using AVX SIMD (Uint32x4, 4 lanes).
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64x2 or handle overflow detection externally.
func SumUint32x4[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint64x2 sums a slice of uint64 using AVX SIMD (Uint64x2, 2 lanes).
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64x2[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat32x4 sums a slice of float32 using AVX SIMD (Float32x4, 4 lanes).
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64x2.
func SumFloat32x4[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat64x2 sums a slice of float64 using AVX SIMD (Float64x2, 2 lanes).
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64x2[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt8x16 calculates the mean of a slice of int8 using AVX SIMD
func MeanInt8x16[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt16x8 calculates the mean of a slice of int16 using AVX SIMD
func MeanInt16x8[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt32x4 calculates the mean of a slice of int32 using AVX SIMD
func MeanInt32x4[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt64x2 calculates the mean of a slice of int64 using AVX SIMD
func MeanInt64x2[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint8x16 calculates the mean of a slice of uint8 using AVX SIMD
func MeanUint8x16[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint16x8 calculates the mean of a slice of uint16 using AVX SIMD
func MeanUint16x8[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint32x4 calculates the mean of a slice of uint32 using AVX SIMD
func MeanUint32x4[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint64x2 calculates the mean of a slice of uint64 using AVX SIMD
func MeanUint64x2[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat32x4 calculates the mean of a slice of float32 using AVX SIMD
func MeanFloat32x4[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat64x2 calculates the mean of a slice of float64 using AVX SIMD
func MeanFloat64x2[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// ClampInt8x16 clamps each element in collection between min and max values using AVX SIMD
func ClampInt8x16[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt16x8 clamps each element in collection between min and max values using AVX SIMD
func ClampInt16x8[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt32x4 clamps each element in collection between min and max values using AVX SIMD
func ClampInt32x4[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint8x16 clamps each element in collection between min and max values using AVX SIMD
func ClampUint8x16[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint16x8 clamps each element in collection between min and max values using AVX SIMD
func ClampUint16x8[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint32x4 clamps each element in collection between min and max values using AVX SIMD
func ClampUint32x4[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampFloat32x4 clamps each element in collection between min and max values using AVX SIMD
func ClampFloat32x4[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampFloat64x2 clamps each element in collection between min and max values using AVX SIMD
func ClampFloat64x2[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// MinInt8x16 finds the minimum value in a collection of int8 using AVX SIMD
func MinInt8x16[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt16x8 finds the minimum value in a collection of int16 using AVX SIMD
func MinInt16x8[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt32x4 finds the minimum value in a collection of int32 using AVX SIMD
func MinInt32x4[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint8x16 finds the minimum value in a collection of uint8 using AVX SIMD
func MinUint8x16[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint16x8 finds the minimum value in a collection of uint16 using AVX SIMD
func MinUint16x8[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint32x4 finds the minimum value in a collection of uint32 using AVX SIMD
func MinUint32x4[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinFloat32x4 finds the minimum value in a collection of float32 using AVX SIMD
func MinFloat32x4[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinFloat64x2 finds the minimum value in a collection of float64 using AVX SIMD
func MinFloat64x2[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt8x16 finds the maximum value in a collection of int8 using AVX SIMD
func MaxInt8x16[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt16x8 finds the maximum value in a collection of int16 using AVX SIMD
func MaxInt16x8[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt32x4 finds the maximum value in a collection of int32 using AVX SIMD
func MaxInt32x4[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint8x16 finds the maximum value in a collection of uint8 using AVX SIMD
func MaxUint8x16[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint16x8 finds the maximum value in a collection of uint16 using AVX SIMD
func MaxUint16x8[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint32x4 finds the maximum value in a collection of uint32 using AVX SIMD
func MaxUint32x4[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxFloat32x4 finds the maximum value in a collection of float32 using AVX SIMD
func MaxFloat32x4[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxFloat64x2 finds the maximum value in a collection of float64 using AVX SIMD
func MaxFloat64x2[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// AVX (128-bit) SIMD sumBy functions - 16/8/4/2 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD sum functions.

// SumByInt8x16 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt8x16[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt16x8 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt16x8[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt32x4 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt32x4[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt64x2 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt64x2[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint8x16 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint8x16[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint16x8 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint16x8[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint32x4 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint32x4[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint64x2 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint64x2[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat32x4 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByFloat32x4[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat64x2 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByFloat64x2[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// AVX (128-bit) SIMD meanBy functions - 16/8/4/2 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD mean functions.

// MeanByInt8x16 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt8x16[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt16x8 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt16x8[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt32x4 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt32x4[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt64x2 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt64x2[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint8x16 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint8x16[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint16x8 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint16x8[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint32x4 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint32x4[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint64x2 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint64x2[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat32x4 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByFloat32x4[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat64x2 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByFloat64x2[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}
