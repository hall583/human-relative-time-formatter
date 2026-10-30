package human_relative_time_formatter

import (
	"fmt"
	"math"
	"time"
)

// Formatter converts a time.Time gap into a relative phrasing such as
// "3 hours ago" or "in 2 days".
//
// Design decision: this formatter bucketizes time into fixed-size units
// (seconds, minutes, hours, days, weeks) rather than calendar-aware
// differences. This keeps the output deterministic and avoids ambiguity
// around daylight-saving transitions and month-length variation.
//
// Threshold-adjustable: Set Threshold to the number of seconds beyond
// which the formatter falls back to a calendar-style "2006-01-02"
// rendering. Zero means "always relative" (no fallback).
type Formatter struct {
	// Threshold is the maximum gap in seconds for relative output.
	// Gaps at or beyond this value render as a calendar date instead.
	// Set to 0 to disable the fallback entirely.
	Threshold float64
}

const (
	secondsPerMinute = 60.0
	secondsPerHour   = 3600.0
	secondsPerDay    = 86400.0
	secondsPerWeek   = 604800.0
)

// Format returns a human-readable relative string describing the gap
// between reference and target.
//
// The reference time acts as "now". A target in the past yields
// "N units ago"; a target in the future yields "in N units".
// A gap smaller than one second in either direction returns "just now",
// which avoids flickering between "0 seconds ago" and "in 0 seconds"
// when the two times are effectively equal.
func (f Formatter) Format(reference, target time.Time) string {
	gap := target.Sub(reference)
	absSeconds := math.Abs(gap.Seconds())

	if f.Threshold > 0 && absSeconds >= f.Threshold {
		return target.Format("2006-01-02")
	}

	if absSeconds < 1 {
		return "just now"
	}

	direction := "ago"
	if gap > 0 {
		direction = "in"
	}

	n, unit := pickBucket(absSeconds)

	if direction == "in" {
		return fmt.Sprintf("in %d %s", n, unit)
	}
	return fmt.Sprintf("%d %s ago", n, unit)
}

// pickBucket converts an absolute second count into an integer magnitude
// and a singular/plural-aware unit name. It intentionally rounds down so
// the reported value never exceeds the true gap.
func pickBucket(absSeconds float64) (int, string) {
	switch {
	case absSeconds < secondsPerMinute:
		return int(absSeconds), pluralize(int(absSeconds), "second")
	case absSeconds < secondsPerHour:
		n := int(absSeconds / secondsPerMinute)
		return n, pluralize(n, "minute")
	case absSeconds < secondsPerDay:
		n := int(absSeconds / secondsPerHour)
		return n, pluralize(n, "hour")
	case absSeconds < secondsPerWeek:
		n := int(absSeconds / secondsPerDay)
		return n, pluralize(n, "day")
	default:
		n := int(absSeconds / secondsPerWeek)
		return n, pluralize(n, "week")
	}
}

// pluralize returns the unit with an "s" suffix when n is not 1.
func pluralize(n int, unit string) string {
	if n == 1 {
		return unit
	}
	return unit + "s"
}
