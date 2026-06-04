//nolint:revive
package xtime

import (
	"time"
)

func NewFakeClock() *FakeClock { _ = "STUB: not implemented"; return nil }

func NewFakeClockAt(t time.Time) *FakeClock { _ = "STUB: not implemented"; return nil }

type FakeClock struct {
	_ noCopy

	// Not protected by a mutex. If a warning is thrown in your tests,
	// just disable parallel tests.
	time time.Time
}

func (c *FakeClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *FakeClock) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *FakeClock) Until(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *FakeClock) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }
