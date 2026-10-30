package human_relative_time_formatter

import (
	"testing"
	"time"
)

func TestFormatFutureSeconds(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(45 * time.Second)
	got := Formatter{}.Format(ref, target)
	want := "in 45 seconds"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatPastSeconds(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(-30 * time.Second)
	got := Formatter{}.Format(ref, target)
	want := "30 seconds ago"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatOneMinuteSingular(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(-60 * time.Second)
	got := Formatter{}.Format(ref, target)
	want := "1 minute ago"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatFutureMinutes(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(5 * time.Minute)
	got := Formatter{}.Format(ref, target)
	want := "in 5 minutes"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatPastHours(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(-3 * time.Hour)
	got := Formatter{}.Format(ref, target)
	want := "3 hours ago"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatOneHourSingularFuture(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(1 * time.Hour)
	got := Formatter{}.Format(ref, target)
	want := "in 1 hour"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatFutureDays(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(48 * time.Hour)
	got := Formatter{}.Format(ref, target)
	want := "in 2 days"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatPastWeeks(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(-14 * 24 * time.Hour)
	got := Formatter{}.Format(ref, target)
	want := "2 weeks ago"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatJustNowNearZero(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(500 * time.Millisecond)
	got := Formatter{}.Format(ref, target)
	want := "just now"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatJustNowExactEqual(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	got := Formatter{}.Format(ref, ref)
	want := "just now"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatThresholdFallback(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(-400 * time.Hour) // ~16.6 days
	f := Formatter{Threshold: 3600 * 24 * 7} // 1 week
	got := f.Format(ref, target)
	want := target.Format("2006-01-02")
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatThresholdNotTriggeredWithinLimit(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	target := ref.Add(-3 * 24 * time.Hour)
	f := Formatter{Threshold: 3600 * 24 * 7} // 1 week
	got := f.Format(ref, target)
	want := "3 days ago"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFormatRoundsDown(t *testing.T) {
	ref := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	// 119 seconds rounds down to 1 minute, not up to 2.
	target := ref.Add(-119 * time.Second)
	got := Formatter{}.Format(ref, target)
	want := "1 minute ago"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
