//go:build go1.23

package it

import "iter"

// ChunkString returns a sequence of strings split into groups of length size. If the string can't be split evenly,
// the final chunk will be the remaining characters.
// Play: https://go.dev/play/p/Bcc5ixTQQoQ
//
// Note: it.ChunkString and it.Chunk functions behave inconsistently for empty input: it.ChunkString("", n) returns [""] instead of [].
// See https://github.com/samber/lo/issues/788
func ChunkString[T ~string](str T, size int) iter.Seq[T] { _ = "STUB: not implemented"; return nil }
