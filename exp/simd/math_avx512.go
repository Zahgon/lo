//go:build go1.26 && goexperiment.simd && amd64

package simd

// AVX-512 (512-bit) SIMD sum functions - 64/32/16/8 lanes

// SumInt8x64 sums a slice of int8 using AVX-512 SIMD (Int8x64, 64 lanes).
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8x64[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt16x32 sums a slice of int16 using AVX-512 SIMD (Int16x32, 32 lanes).
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16x32[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt32x16 sums a slice of int32 using AVX-512 SIMD (Int32x16, 16 lanes).
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64x8 or handle overflow detection externally.
func SumInt32x16[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumInt64x8 sums a slice of int64 using AVX-512 SIMD (Int64x8, 8 lanes).
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64x8[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint8x64 sums a slice of uint8 using AVX-512 SIMD (Uint8x64, 64 lanes).
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8x64[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint16x32 sums a slice of uint16 using AVX-512 SIMD (Uint16x32, 32 lanes).
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16x32[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint32x16 sums a slice of uint32 using AVX-512 SIMD (Uint32x16, 16 lanes).
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64x8 or handle overflow detection externally.
func SumUint32x16[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumUint64x8 sums a slice of uint64 using AVX-512 SIMD (Uint64x8, 8 lanes).
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64x8[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat32x16 sums a slice of float32 using AVX-512 SIMD (Float32x16, 16 lanes).
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64x8.
func SumFloat32x16[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// SumFloat64x8 sums a slice of float64 using AVX-512 SIMD (Float64x8, 8 lanes).
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64x8[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt8x64 calculates the mean of a slice of int8 using AVX-512 SIMD
func MeanInt8x64[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt16x32 calculates the mean of a slice of int16 using AVX-512 SIMD
func MeanInt16x32[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt32x16 calculates the mean of a slice of int32 using AVX-512 SIMD
func MeanInt32x16[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanInt64x8 calculates the mean of a slice of int64 using AVX-512 SIMD
func MeanInt64x8[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint8x64 calculates the mean of a slice of uint8 using AVX-512 SIMD
func MeanUint8x64[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint16x32 calculates the mean of a slice of uint16 using AVX-512 SIMD
func MeanUint16x32[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint32x16 calculates the mean of a slice of uint32 using AVX-512 SIMD
func MeanUint32x16[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanUint64x8 calculates the mean of a slice of uint64 using AVX-512 SIMD
func MeanUint64x8[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat32x16 calculates the mean of a slice of float32 using AVX-512 SIMD
func MeanFloat32x16[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// MeanFloat64x8 calculates the mean of a slice of float64 using AVX-512 SIMD
func MeanFloat64x8[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// ClampInt8x64 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt8x64[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt16x32 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt16x32[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt32x16[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt64x2 clamps each element in collection between min and max values using AVX-512 SIMD.
// Int64x2 Min/Max operations in archsimd require AVX-512 (VPMAXSQ/VPMINSQ).
func ClampInt64x2[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint64x2 clamps each element in collection between min and max values using AVX-512 SIMD.
// Uint64x2 Min/Max operations in archsimd require AVX-512.
func ClampUint64x2[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampInt64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt64x8[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint8x64 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint8x64[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint16x32 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint16x32[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint32x16[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampUint64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint64x8[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampFloat32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampFloat32x16[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// ClampFloat64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampFloat64x8[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

// bearer:disable go_gosec_unsafe_unsafe

// MinInt8x64 finds the minimum value in a collection of int8 using AVX-512 SIMD
func MinInt8x64[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt16x32 finds the minimum value in a collection of int16 using AVX-512 SIMD
func MinInt16x32[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt32x16 finds the minimum value in a collection of int32 using AVX-512 SIMD
func MinInt32x16[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt64x2 finds the minimum value in a collection of int64 using AVX-512 SIMD.
// Int64x2 Min operations in archsimd require AVX-512.
func MinInt64x2[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint64x2 finds the minimum value in a collection of uint64 using AVX-512 SIMD.
// Uint64x2 Min operations in archsimd require AVX-512.
func MinUint64x2[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinInt64x8 finds the minimum value in a collection of int64 using AVX-512 SIMD
func MinInt64x8[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint8x64 finds the minimum value in a collection of uint8 using AVX-512 SIMD
func MinUint8x64[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint16x32 finds the minimum value in a collection of uint16 using AVX-512 SIMD
func MinUint16x32[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint32x16 finds the minimum value in a collection of uint32 using AVX-512 SIMD
func MinUint32x16[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinUint64x8 finds the minimum value in a collection of uint64 using AVX-512 SIMD
func MinUint64x8[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinFloat32x16 finds the minimum value in a collection of float32 using AVX-512 SIMD
func MinFloat32x16[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MinFloat64x8 finds the minimum value in a collection of float64 using AVX-512 SIMD
func MinFloat64x8[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find minimum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt8x64 finds the maximum value in a collection of int8 using AVX-512 SIMD
func MaxInt8x64[T ~int8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt16x32 finds the maximum value in a collection of int16 using AVX-512 SIMD
func MaxInt16x32[T ~int16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt32x16 finds the maximum value in a collection of int32 using AVX-512 SIMD
func MaxInt32x16[T ~int32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt64x2 finds the maximum value in a collection of int64 using AVX-512 SIMD.
// Int64x2 Max operations in archsimd require AVX-512.
func MaxInt64x2[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint64x2 finds the maximum value in a collection of uint64 using AVX-512 SIMD.
// Uint64x2 Max operations in archsimd require AVX-512.
func MaxUint64x2[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxInt64x8 finds the maximum value in a collection of int64 using AVX-512 SIMD
func MaxInt64x8[T ~int64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint8x64 finds the maximum value in a collection of uint8 using AVX-512 SIMD
func MaxUint8x64[T ~uint8](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint16x32 finds the maximum value in a collection of uint16 using AVX-512 SIMD
func MaxUint16x32[T ~uint16](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint32x16 finds the maximum value in a collection of uint32 using AVX-512 SIMD
func MaxUint32x16[T ~uint32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxUint64x8 finds the maximum value in a collection of uint64 using AVX-512 SIMD
func MaxUint64x8[T ~uint64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxFloat32x16 finds the maximum value in a collection of float32 using AVX-512 SIMD
func MaxFloat32x16[T ~float32](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// MaxFloat64x8 finds the maximum value in a collection of float64 using AVX-512 SIMD
func MaxFloat64x8[T ~float64](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// Find maximum in the vector (only if we processed any vectors)

// Handle remaining elements

// AVX-512 (512-bit) SIMD sumBy functions - 64/32/16/8 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD sum functions.

// SumByInt8x64 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt8x64[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt16x32 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt16x32[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt32x16 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt32x16[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByInt64x8 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt64x8[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint8x64 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint8x64[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint16x32 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint16x32[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint32x16 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint32x16[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByUint64x8 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint64x8[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat32x16 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByFloat32x16[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// SumByFloat64x8 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByFloat64x8[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// AVX-512 (512-bit) SIMD meanBy functions - 64/32/16/8 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD mean functions.

// MeanByInt8x64 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt8x64[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt16x32 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt16x32[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt32x16 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt32x16[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByInt64x8 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt64x8[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint8x64 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint8x64[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint16x32 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint16x32[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint32x16 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint32x16[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByUint64x8 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint64x8[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat32x16 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByFloat32x16[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// MeanByFloat64x8 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByFloat64x8[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	_ = "STUB: not implemented"
	return *new(R)
}
