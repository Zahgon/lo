//go:build go1.26 && goexperiment.simd && amd64

package simd

// unsafeSliceInt8 converts a []T (where T ~int8) to []int8 via unsafe operations.
// This helper reduces code duplication and the risk of copy-paste errors.
//
//go:nosplit
func unsafeSliceInt8[T ~int8](collection []T, length uint) []int8 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceInt16 converts a []T (where T ~int16) to []int16 via unsafe operations.
//
//go:nosplit
func unsafeSliceInt16[T ~int16](collection []T, length uint) []int16 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceInt32 converts a []T (where T ~int32) to []int32 via unsafe operations.
//
//go:nosplit
func unsafeSliceInt32[T ~int32](collection []T, length uint) []int32 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceInt64 converts a []T (where T ~int64) to []int64 via unsafe operations.
//
//go:nosplit
func unsafeSliceInt64[T ~int64](collection []T, length uint) []int64 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceUint8 converts a []T (where T ~uint8) to []uint8 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint8[T ~uint8](collection []T, length uint) []uint8 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceUint16 converts a []T (where T ~uint16) to []uint16 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint16[T ~uint16](collection []T, length uint) []uint16 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceUint32 converts a []T (where T ~uint32) to []uint32 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint32[T ~uint32](collection []T, length uint) []uint32 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceUint64 converts a []T (where T ~uint64) to []uint64 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint64[T ~uint64](collection []T, length uint) []uint64 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceFloat32 converts a []T (where T ~float32) to []float32 via unsafe operations.
//
//go:nosplit
func unsafeSliceFloat32[T ~float32](collection []T, length uint) []float32 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// unsafeSliceFloat64 converts a []T (where T ~float64) to []float64 via unsafe operations.
//
//go:nosplit
func unsafeSliceFloat64[T ~float64](collection []T, length uint) []float64 {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}
