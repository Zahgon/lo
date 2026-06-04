//go:build go1.26 && goexperiment.simd && amd64

package simd

// AVX2 (256-bit) SIMD sum functions - 32/16/8/4 lanes

// SumInt8x32 sums a slice of int8 using AVX2 SIMD (Int8x32, 32 lanes).
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8x32[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt16x16 sums a slice of int16 using AVX2 SIMD (Int16x16, 16 lanes).
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16x16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt32x8 sums a slice of int32 using AVX2 SIMD (Int32x8, 8 lanes).
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64x4 or handle overflow detection externally.
func SumInt32x8[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt64x4 sums a slice of int64 using AVX2 SIMD (Int64x4, 4 lanes).
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64x4[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint8x32 sums a slice of uint8 using AVX2 SIMD (Uint8x32, 32 lanes).
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8x32[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint16x16 sums a slice of uint16 using AVX2 SIMD (Uint16x16, 16 lanes).
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16x16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint32x8 sums a slice of uint32 using AVX2 SIMD (Uint32x8, 8 lanes).
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64x4 or handle overflow detection externally.
func SumUint32x8[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint64x4 sums a slice of uint64 using AVX2 SIMD (Uint64x4, 4 lanes).
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64x4[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat32x8 sums a slice of float32 using AVX2 SIMD (Float32x8, 8 lanes).
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64x4.
func SumFloat32x8[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat64x4 sums a slice of float64 using AVX2 SIMD (Float64x4, 4 lanes).
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64x4[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt8x32 calculates the mean of a slice of int8 using AVX2 SIMD
func MeanInt8x32[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt16x16 calculates the mean of a slice of int16 using AVX2 SIMD
func MeanInt16x16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt32x8 calculates the mean of a slice of int32 using AVX2 SIMD
func MeanInt32x8[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt64x4 calculates the mean of a slice of int64 using AVX2 SIMD
func MeanInt64x4[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint8x32 calculates the mean of a slice of uint8 using AVX2 SIMD
func MeanUint8x32[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint16x16 calculates the mean of a slice of uint16 using AVX2 SIMD
func MeanUint16x16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint32x8 calculates the mean of a slice of uint32 using AVX2 SIMD
func MeanUint32x8[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint64x4 calculates the mean of a slice of uint64 using AVX2 SIMD
func MeanUint64x4[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat32x8 calculates the mean of a slice of float32 using AVX2 SIMD
func MeanFloat32x8[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat64x4 calculates the mean of a slice of float64 using AVX2 SIMD
func MeanFloat64x4[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// ClampInt8x32 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt8x32[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt16x16 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt16x16[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt32x8 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt32x8[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt64x4 clamps each element in collection between min and max values using AVX2 SIMD and AVX-512 SIMD.
func ClampInt64x4[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint8x32 clamps each element in collection between min and max values using AVX2 SIMD
func ClampUint8x32[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint16x16 clamps each element in collection between min and max values using AVX2 SIMD
func ClampUint16x16[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint32x8 clamps each element in collection between min and max values using AVX2 SIMD
func ClampUint32x8[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint64x4 clamps each element in collection between min and max values using AVX2 SIMD and AVX-512 SIMD.
func ClampUint64x4[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampFloat32x8 clamps each element in collection between min and max values using AVX2 SIMD
func ClampFloat32x8[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampFloat64x4 clamps each element in collection between min and max values using AVX2 SIMD
func ClampFloat64x4[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// MinInt8x32 finds the minimum value in a collection of int8 using AVX2 SIMD
func MinInt8x32[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt16x16 finds the minimum value in a collection of int16 using AVX2 SIMD
func MinInt16x16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt32x8 finds the minimum value in a collection of int32 using AVX2 SIMD
func MinInt32x8[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt64x4 finds the minimum value in a collection of int64 using AVX2 SIMD
func MinInt64x4[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint8x32 finds the minimum value in a collection of uint8 using AVX2 SIMD
func MinUint8x32[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint16x16 finds the minimum value in a collection of uint16 using AVX2 SIMD
func MinUint16x16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint32x8 finds the minimum value in a collection of uint32 using AVX2 SIMD
func MinUint32x8[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint64x4 finds the minimum value in a collection of uint64 using AVX2 SIMD
func MinUint64x4[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinFloat32x8 finds the minimum value in a collection of float32 using AVX2 SIMD
func MinFloat32x8[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinFloat64x4 finds the minimum value in a collection of float64 using AVX2 SIMD
func MinFloat64x4[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt8x32 finds the maximum value in a collection of int8 using AVX2 SIMD
func MaxInt8x32[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt16x16 finds the maximum value in a collection of int16 using AVX2 SIMD
func MaxInt16x16[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt32x8 finds the maximum value in a collection of int32 using AVX2 SIMD
func MaxInt32x8[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt64x4 finds the maximum value in a collection of int64 using AVX2 SIMD
func MaxInt64x4[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint8x32 finds the maximum value in a collection of uint8 using AVX2 SIMD
func MaxUint8x32[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint16x16 finds the maximum value in a collection of uint16 using AVX2 SIMD
func MaxUint16x16[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint32x8 finds the maximum value in a collection of uint32 using AVX2 SIMD
func MaxUint32x8[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint64x4 finds the maximum value in a collection of uint64 using AVX2 SIMD
func MaxUint64x4[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxFloat32x8 finds the maximum value in a collection of float32 using AVX2 SIMD
func MaxFloat32x8[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxFloat64x4 finds the maximum value in a collection of float64 using AVX2 SIMD
func MaxFloat64x4[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// AVX2 (256-bit) SIMD sumBy functions - 32/16/8/4 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD sum functions.

// SumByInt8x32 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByInt8x32[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt16x16 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByInt16x16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt32x8 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByInt32x8[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt64x4 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByInt64x4[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint8x32 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByUint8x32[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint16x16 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByUint16x16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint32x8 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByUint32x8[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint64x4 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByUint64x4[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat32x8 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByFloat32x8[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat64x4 sums the values extracted by iteratee from a slice using AVX2 SIMD.
func SumByFloat64x4[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// AVX2 (256-bit) SIMD meanBy functions - 32/16/8/4 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD mean functions.

// MeanByInt8x32 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByInt8x32[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt16x16 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByInt16x16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt32x8 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByInt32x8[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt64x4 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByInt64x4[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint8x32 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByUint8x32[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint16x16 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByUint16x16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint32x8 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByUint32x8[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint64x4 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByUint64x4[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat32x8 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByFloat32x8[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat64x4 calculates the mean of values extracted by iteratee from a slice using AVX2 SIMD.
func MeanByFloat64x4[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}
