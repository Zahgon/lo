//nolint:revive
package xtime

import "time"

var clock Clock = &RealClock{}

func SetClock(c Clock) { _ = "STUB: not implemented"; return }

func Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func Since(t time.Time) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func Until(t time.Time) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func Sleep(d time.Duration) { _ = "STUB: not implemented"; return }

type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration
	Sleep(d time.Duration)
}
