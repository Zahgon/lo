//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo"
)

// SeqToChannel returns a read-only channel of collection elements.
// Play: https://go.dev/play/p/id3jqJPffT6
func SeqToChannel[T any](bufferSize int, collection iter.Seq[T]) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// SeqToChannel2 returns a read-only channel of collection elements.
// Play: https://go.dev/play/p/rpJdVnXUaG-
func SeqToChannel2[K, V any](bufferSize int, collection iter.Seq2[K, V]) <-chan lo.Tuple2[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// ChannelToSeq returns a sequence built from channels items. Blocks until channel closes.
// Play: https://go.dev/play/p/IXqSs2Ooqpm
func ChannelToSeq[T any](ch <-chan T) iter.Seq[T] { _ = "STUB: not implemented"; return nil }
