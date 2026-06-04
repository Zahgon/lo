package lo

import (
	"sync"
	"time"
)

type debounce struct {
	after     time.Duration
	mu        *sync.Mutex
	timer     *time.Timer
	done      bool
	callbacks []func()
}

func (d *debounce) reset() { _ = "STUB: not implemented"; return }

// We need to lock the mutex here to avoid race conditions with 2 concurrent calls to reset()

func (d *debounce) cancel() { _ = "STUB: not implemented"; return }

// NewDebounce creates a debounced instance that delays invoking functions given until after wait milliseconds have elapsed.
// Play: https://go.dev/play/p/_IPY7ROzbMk
func NewDebounce(duration time.Duration, f ...func()) (func(), func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

type debounceByItem struct {
	mu    *sync.Mutex
	timer *time.Timer
	count int
}

type debounceBy[T comparable] struct {
	after     time.Duration
	mu        *sync.Mutex
	items     map[T]*debounceByItem
	callbacks []func(key T, count int)
}

func (d *debounceBy[T]) reset(key T) { _ = "STUB: not implemented"; return }

// We need to lock the mutex here to avoid race conditions with 2 concurrent calls to reset()

func (d *debounceBy[T]) cancel(key T) { _ = "STUB: not implemented"; return }

// NewDebounceBy creates a debounced instance for each distinct key, that delays invoking functions given until after wait milliseconds have elapsed.
// Play: https://go.dev/play/p/Izk7GEzZm2Q
func NewDebounceBy[T comparable](duration time.Duration, f ...func(key T, count int)) (func(key T), func(key T)) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempt invokes a function N times until it returns valid output. Returns either the caught error or nil.
// When the first argument is less than `1`, the function runs until a successful response is returned.
// Play: https://go.dev/play/p/3ggJZ2ZKcMj
func Attempt(maxIteration int, f func(index int) error) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// for retries >= 0 {

// AttemptWithDelay invokes a function N times until it returns valid output,
// with a pause between each call. Returns either the caught error or nil.
// When the first argument is less than `1`, the function runs until a successful
// response is returned.
// Play: https://go.dev/play/p/tVs6CygC7m1
func AttemptWithDelay(maxIteration int, delay time.Duration, f func(index int, duration time.Duration) error) (int, time.Duration, error) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), nil
}

// AttemptWhile invokes a function N times until it returns valid output.
// Returns either the caught error or nil, along with a bool value to determine
// whether the function should be invoked again. It will terminate the invoke
// immediately if the second return value is false. When the first
// argument is less than `1`, the function runs until a successful response is
// returned.
// Play: https://go.dev/play/p/1VS7HxlYMOG
func AttemptWhile(maxIteration int, f func(int) (error, bool)) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// for retries >= 0 {

// if shouldContinueInvoke is false, then return immediately

// AttemptWhileWithDelay invokes a function N times until it returns valid output,
// with a pause between each call. Returns either the caught error or nil, along
// with a bool value to determine whether the function should be invoked again.
// It will terminate the invoke immediately if the second return value is false.
// When the first argument is less than `1`, the function runs until a successful
// response is returned.
// Play: https://go.dev/play/p/mhufUjJfLEF
func AttemptWhileWithDelay(maxIteration int, delay time.Duration, f func(int, time.Duration) (error, bool)) (int, time.Duration, error) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), nil
}

// if shouldContinueInvoke is false, then return immediately

type transactionStep[T any] struct {
	exec       func(T) (T, error)
	onRollback func(T) T
}

// NewTransaction instantiate a new transaction.
// Play: https://go.dev/play/p/7B2o52wEQbj
func NewTransaction[T any]() *Transaction[T] { _ = "STUB: not implemented"; return nil }

// Transaction implements a Saga pattern.
type Transaction[T any] struct {
	steps []transactionStep[T]
}

// Then adds a step to the chain of callbacks. Returns the same Transaction.
// Play: https://go.dev/play/p/Qxrd7MGQGh1 https://go.dev/play/p/xrHb2_kMvTY
func (t *Transaction[T]) Then(exec func(T) (T, error), onRollback func(T) T) *Transaction[T] {
	_ = "STUB: not implemented"
	return nil
}

// Process runs the Transaction steps and rollbacks in case of errors.
// Play: https://go.dev/play/p/Qxrd7MGQGh1 https://go.dev/play/p/xrHb2_kMvTY
func (t *Transaction[T]) Process(state T) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// @TODO: single mutex per key?
type throttleBy[T comparable] struct {
	mu         *sync.Mutex
	timer      *time.Timer
	interval   time.Duration
	callbacks  []func(key T)
	countLimit int
	count      map[T]int
}

func (th *throttleBy[T]) throttledFunc(key T) { _ = "STUB: not implemented"; return }

func (th *throttleBy[T]) reset() { _ = "STUB: not implemented"; return }

// NewThrottle creates a throttled instance that invokes given functions only once in every interval.
// This returns 2 functions, First one is throttled function and Second one is a function to reset interval.
// Play: https://go.dev/play/p/qQn3fm8Z7jS
func NewThrottle(interval time.Duration, f ...func()) (throttle, reset func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewThrottleWithCount is NewThrottle with count limit, throttled function will be invoked count times in every interval.
// Play: https://go.dev/play/p/w5nc0MgWtjC
func NewThrottleWithCount(interval time.Duration, count int, f ...func()) (throttle, reset func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewThrottleBy creates a throttled instance that invokes given functions only once in every interval.
// This returns 2 functions, First one is throttled function and Second one is a function to reset interval.
// Play: https://go.dev/play/p/0Wv6oX7dHdC
func NewThrottleBy[T comparable](interval time.Duration, f ...func(key T)) (throttle func(key T), reset func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewThrottleByWithCount is NewThrottleBy with count limit, throttled function will be invoked count times in every interval.
// Play: https://go.dev/play/p/vQk3ECH7_EW
func NewThrottleByWithCount[T comparable](interval time.Duration, count int, f ...func(key T)) (throttle func(key T), reset func()) {
	_ = "STUB: not implemented"
	return nil, nil
}
