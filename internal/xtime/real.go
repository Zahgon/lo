//nolint:revive
package xtime

import (
	"time"
)

func NewRealClock() *RealClock { _ = "STUB: not implemented"; return nil }

type RealClock struct {
	_ noCopy
}

func (c *RealClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *RealClock) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *RealClock) Until(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *RealClock) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }
