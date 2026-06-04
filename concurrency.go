package lo

import (
	"context"
	"sync"
	"time"
)

type synchronize struct {
	locker sync.Locker
}

func (s *synchronize) Do(callback func()) { _ = "STUB: not implemented"; return }

// Synchronize wraps the underlying callback in a mutex. It receives an optional mutex.
// Play: https://go.dev/play/p/X3cqROSpQmu
func Synchronize(opt ...sync.Locker) *synchronize {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

// Async executes a function in a goroutine and returns the result in a channel.
// Play: https://go.dev/play/p/uo35gosuTLw
func Async[A any](f func() A) <-chan A { _ = "STUB: not implemented"; return nil }

// Async0 executes a function in a goroutine and returns a channel set once the function finishes.
// Play: https://go.dev/play/p/tNqf1cClG_o
func Async0(f func()) <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Async1 is an alias to Async.
// Play: https://go.dev/play/p/RBQWtIn4PsF
func Async1[A any](f func() A) <-chan A {
	_ = "STUB: not implemented"

	// Async2 has the same behavior as Async, but returns the 2 results as a tuple inside the channel.
	// Play: https://go.dev/play/p/5SzzDjssXOH
	return nil
}

func Async2[A, B any](f func() (A, B)) <-chan Tuple2[A, B] { _ = "STUB: not implemented"; return nil }

// Async3 has the same behavior as Async, but returns the 3 results as a tuple inside the channel.
// Play: https://go.dev/play/p/cZpZsDXNmlx
func Async3[A, B, C any](f func() (A, B, C)) <-chan Tuple3[A, B, C] {
	_ = "STUB: not implemented"
	return nil
}

// Async4 has the same behavior as Async, but returns the 4 results as a tuple inside the channel.
// Play: https://go.dev/play/p/9X5O2VrLzkR
func Async4[A, B, C, D any](f func() (A, B, C, D)) <-chan Tuple4[A, B, C, D] {
	_ = "STUB: not implemented"
	return nil
}

// Async5 has the same behavior as Async, but returns the 5 results as a tuple inside the channel.
// Play: https://go.dev/play/p/MqnUJpkmopA
func Async5[A, B, C, D, E any](f func() (A, B, C, D, E)) <-chan Tuple5[A, B, C, D, E] {
	_ = "STUB: not implemented"
	return nil
}

// Async6 has the same behavior as Async, but returns the 6 results as a tuple inside the channel.
// Play: https://go.dev/play/p/kM1X67JPdSP
func Async6[A, B, C, D, E, F any](f func() (A, B, C, D, E, F)) <-chan Tuple6[A, B, C, D, E, F] {
	_ = "STUB: not implemented"
	return nil
}

// WaitFor runs periodically until a condition is validated.
// Play: https://go.dev/play/p/t_wTDmubbK3
func WaitFor(condition func(i int) bool, timeout, heartbeatDelay time.Duration) (totalIterations int, elapsed time.Duration, conditionFound bool) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), false
}

// WaitForWithContext runs periodically until a condition is validated or context is canceled.
// Play: https://go.dev/play/p/t_wTDmubbK3
func WaitForWithContext(ctx context.Context, condition func(ctx context.Context, currentIteration int) bool, timeout, heartbeatDelay time.Duration) (totalIterations int, elapsed time.Duration, conditionFound bool) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), false
}
