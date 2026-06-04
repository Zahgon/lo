package lo

import (
	"math"
	"regexp"
)

var (
	//nolint:revive
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)

	// bearer:disable go_lang_permissive_regex_validation
	splitWordReg = regexp.MustCompile(`([a-z])([A-Z0-9])|([a-zA-Z])([0-9])|([0-9])([a-zA-Z])|([A-Z])([A-Z])([a-z])`)
	// bearer:disable go_lang_permissive_regex_validation
	splitNumberLetterReg = regexp.MustCompile(`([0-9])([a-zA-Z])`)
	maximumCapacity      = math.MaxInt>>1 + 1
)

// RandomString return a random string.
// Play: https://go.dev/play/p/rRseOQVVum4
func RandomString(size int, charset []rune) string { _ = "STUB: not implemented"; return "" }

// see https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

// Edge case, because if the charset is a single character,
// it will panic below (divide by zero).
// -> https://github.com/samber/lo/issues/679

// Calculate the number of bits required to represent the charset,
// e.g., for 62 characters, it would need 6 bits (since 62 -> 64 = 2^6)

// Determine the corresponding bitmask,
// e.g., for 62 characters, the bitmask would be 111111.

// Available count, since xrand.Int64() returns a non-negative number, the first bit is fixed, so there are 63 random bits
// e.g., for 62 characters, this value is 10 (63 / 6).

// Generate the random string in a loop.

// Regenerate the random number if all available bits have been used

// Select a character from the charset

// Shift the bits to the right to prepare for the next character selection,
// e.g., for 62 characters, shift by 6 bits.

// Decrease the remaining number of uses for the current random number.

// nearestPowerOfTwo returns the nearest power of two.
func nearestPowerOfTwo(capacity int) int { _ = "STUB: not implemented"; return 0 }

// Substring extracts a substring from a string with Unicode character (rune) awareness.
// offset - starting position of the substring (can be positive, negative, or zero)
// length - number of characters to extract
// With positive offset, counting starts from the beginning of the string
// With negative offset, counting starts from the end of the string
// Play: https://go.dev/play/p/emzCC9zBjHu
func Substring[T ~string](str T, offset int, length uint) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// Validate UTF-8 and fix invalid sequences

// Convert to []rune to replicate behavior with duplicated �

// Remove null bytes from result

func substring[T ~string](str T, offset int, length uint) T {
	_ = "STUB: not implemented"

	// Empty length or offset beyond string bounds - return empty string
	return *new(T)
}

// Positive offset - count from the beginning

// Skip offset runes from the start

// If couldn't skip enough runes - string is shorter than offset

// If remaining string is shorter than or equal to length - return it entirely

// Otherwise proceed to trimming by length

// Zero offset or offset less than minus string length - start from beginning

// Count length runes from the start

// Negative offset - count from the end of string
// -len(str) < offset < 0
// Helper function to move backward through runes

// If offset is less than or equal to length - take from position to end

// Otherwise calculate start and end positions

// ChunkString returns a slice of strings split into groups of length size. If the string can't be split evenly,
// the final chunk will be the remaining characters.
// Play: https://go.dev/play/p/__FLTuJVz54
//
// Note: lo.ChunkString and lo.Chunk functions behave inconsistently for empty input: lo.ChunkString("", n) returns [""] instead of [].
// See https://github.com/samber/lo/issues/788
func ChunkString[T ~string](str T, size int) []T { _ = "STUB: not implemented"; return nil }

// RuneLength is an alias to utf8.RuneCountInString which returns the number of runes in string.
// Play: https://go.dev/play/p/BXT52mBk0zO
func RuneLength(str string) int { _ = "STUB: not implemented"; return 0 }

// PascalCase converts string to pascal case.
// Play: https://go.dev/play/p/uxER7XpRHLB
func PascalCase(str string) string { _ = "STUB: not implemented"; return "" }

// CamelCase converts string to camel case.
// Play: https://go.dev/play/p/4JNDzaMwXkm
func CamelCase(str string) string { _ = "STUB: not implemented"; return "" }

// KebabCase converts string to kebab case.
// Play: https://go.dev/play/p/ZBeMB4-pq45
func KebabCase(str string) string { _ = "STUB: not implemented"; return "" }

// SnakeCase converts string to snake case.
// Play: https://go.dev/play/p/ziB0V89IeVH
func SnakeCase(str string) string { _ = "STUB: not implemented"; return "" }

// Words splits string into a slice of its words.
// Play: https://go.dev/play/p/-f3VIQqiaVw
func Words(str string) []string { _ = "STUB: not implemented"; return nil }

// example: Int8Value => Int 8Value => Int 8 Value

// Capitalize converts the first character of string to upper case and the remaining to lower case.
// Play: https://go.dev/play/p/uLTZZQXqnsa
func Capitalize(str string) string { _ = "STUB: not implemented"; return "" }

// Ellipsis trims and truncates a string to a specified length in runes and appends an ellipsis
// if truncated. The length parameter counts Unicode code points (runes), not bytes, so multi-byte
// characters such as emoji or CJK ideographs are never split in the middle.
// Play: https://go.dev/play/p/qE93rgqe1TW
func Ellipsis(str string, length int) string { _ = "STUB: not implemented"; return "" }
