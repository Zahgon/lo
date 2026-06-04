package lo

import (
	"context"
	"time"
)

// DispatchingStrategy is a function that distributes messages to channels.
type DispatchingStrategy[T any] func(msg T, index uint64, channels []<-chan T) int

// ChannelDispatcher distributes messages from input channels into N child channels.
// Close events are propagated to children.
// Underlying channels can have a fixed buffer capacity or be unbuffered when cap is 0.
// Play: https://go.dev/play/p/UZGu2wVg3J2
func ChannelDispatcher[T any](stream <-chan T, count, channelBufferCap int, strategy DispatchingStrategy[T]) []<-chan T {
	_ = "STUB: not implemented"
	return nil
}

// propagate channel closing to children

func createChannels[T any](count, channelBufferCap int) []chan T {
	_ = "STUB: not implemented"
	return nil
}

func channelsToReadOnly[T any](children []chan T) []<-chan T { _ = "STUB: not implemented"; return nil }

func closeChannels[T any](children []chan T) { _ = "STUB: not implemented"; return }

func channelIsNotFull[T any](ch <-chan T) bool { _ = "STUB: not implemented"; return false }

// DispatchingStrategyRoundRobin distributes messages in a rotating sequential manner.
// If the channel capacity is exceeded, the next channel will be selected and so on.
// Play: https://go.dev/play/p/UZGu2wVg3J2
func DispatchingStrategyRoundRobin[T any](msg T, index uint64, channels []<-chan T) int {
	_ = "STUB: not implemented"
	return 0
}

// prevent CPU from burning 🔥

// DispatchingStrategyRandom distributes messages in a random manner.
// If the channel capacity is exceeded, another random channel will be selected and so on.
// Play: https://go.dev/play/p/GEyGn3TdGk4
func DispatchingStrategyRandom[T any](msg T, index uint64, channels []<-chan T) int {
	_ = "STUB: not implemented"
	return 0
}

// prevent CPU from burning 🔥

// DispatchingStrategyWeightedRandom distributes messages in a weighted manner.
// If the channel capacity is exceeded, another random channel will be selected and so on.
// Play: https://go.dev/play/p/v0eMh8NZG2L
func DispatchingStrategyWeightedRandom[T any](weights []int) DispatchingStrategy[T] {
	_ = "STUB: not implemented"
	return nil
}

// prevent CPU from burning 🔥

// DispatchingStrategyFirst distributes messages in the first non-full channel.
// If the capacity of the first channel is exceeded, the second channel will be selected and so on.
// Play: https://go.dev/play/p/OrJCvOmk42f
func DispatchingStrategyFirst[T any](msg T, index uint64, channels []<-chan T) int {
	_ = "STUB: not implemented"
	return 0
}

// prevent CPU from burning 🔥

// DispatchingStrategyLeast distributes messages in the emptiest channel.
// Play: https://go.dev/play/p/ypy0jrRcEe7
func DispatchingStrategyLeast[T any](msg T, index uint64, channels []<-chan T) int {
	_ = "STUB: not implemented"
	return 0
}

// DispatchingStrategyMost distributes messages in the fullest channel.
// If the channel capacity is exceeded, the next channel will be selected and so on.
// Play: https://go.dev/play/p/erHHone7rF9
func DispatchingStrategyMost[T any](msg T, index uint64, channels []<-chan T) int {
	_ = "STUB: not implemented"
	return 0
}

// SliceToChannel returns a read-only channel of collection elements.
// Play: https://go.dev/play/p/lIbSY3QmiEg
func SliceToChannel[T any](bufferSize int, collection []T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// ChannelToSlice returns a slice built from channel items. Blocks until channel closes.
// Play: https://go.dev/play/p/lIbSY3QmiEg
func ChannelToSlice[T any](ch <-chan T) []T { _ = "STUB: not implemented"; return nil }

// Generator implements the generator design pattern.
// Play: https://go.dev/play/p/lIbSY3QmiEg
//
// Deprecated: use "iter" package instead (Go >= 1.23).
func Generator[T any](bufferSize int, generator func(yield func(T))) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// WARNING: infinite loop

// Buffer creates a slice of n elements from a channel. Returns the slice and the slice length.
// @TODO: we should probably provide a helper that reuses the same buffer.
// Play: https://go.dev/play/p/gPQ-6xmcKQI
func Buffer[T any](ch <-chan T, size int) (collection []T, length int, readTime time.Duration, ok bool) {
	_ = "STUB: not implemented"
	return nil, 0, *new(time.Duration), false
}

// BufferWithContext creates a slice of n elements from a channel, with context. Returns the slice and the slice length.
// @TODO: we should probably provide a helper that reuses the same buffer.
// Play: https://go.dev/play/p/oRfOyJWK9YF
func BufferWithContext[T any](ctx context.Context, ch <-chan T, size int) (collection []T, length int, readTime time.Duration, ok bool) {
	_ = "STUB: not implemented"
	return nil, 0, *new(time.Duration), false
}

// BufferWithTimeout creates a slice of n elements from a channel, with timeout. Returns the slice and the slice length.
// Play: https://go.dev/play/p/sxyEM3koo4n
func BufferWithTimeout[T any](ch <-chan T, size int, timeout time.Duration) (collection []T, length int, readTime time.Duration, ok bool) {
	_ = "STUB: not implemented"
	return nil, 0, *new(time.Duration), false
}

// FanIn collects messages from multiple input channels into a single buffered channel.
// Output messages have no priority. When all upstream channels reach EOF, downstream channel closes.
// Play: https://go.dev/play/p/FH8Wq-T04Jb
func FanIn[T any](channelBufferCap int, upstreams ...<-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// Start an output goroutine for each input channel in upstreams.

// Start a goroutine to close out once all the output goroutines are done.

// FanOut broadcasts all the upstream messages to multiple downstream channels.
// When upstream channel reaches EOF, downstream channels close. If any downstream
// channels is full, broadcasting is paused.
// Play: https://go.dev/play/p/2LHxcjKX23L
func FanOut[T any](count, channelsBufferCap int, upstream <-chan T) []<-chan T {
	_ = "STUB: not implemented"
	return nil
}

// Close out once all the output goroutines are done.
