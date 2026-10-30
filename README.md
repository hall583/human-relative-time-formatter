# human-relative-time-formatter

Turns `time.Time` gaps into relative phrasing such as "3 hours ago" or "in 2 days", with an adjustable threshold beyond which output falls back to a calendar date.

```go
package main

import (
    "fmt"
    "time"

    human_relative_time_formatter "human-relative-time-formatter"
)

func main() {
    now := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
    f := human_relative_time_formatter.Formatter{Threshold: 7 * 24 * 3600}
    fmt.Println(f.Format(now, now.Add(-3*time.Hour)))  // "3 hours ago"
    fmt.Println(f.Format(now, now.Add(2*24*time.Hour))) // "in 2 days"
}
```

## Why

Relative time strings are useful in UIs and logs, but calendar-aware differences ("how many months?") are ambiguous: months vary in length, and daylight saving shifts make "24 hours ago" occasionally disagree with "yesterday." This library sidesteps that by bucketing gaps into fixed units—seconds, minutes, hours, days, weeks—so the output is deterministic and easy to reason about.

The `Threshold` field lets you cap how far back the relative phrasing goes. Set it to the number of seconds beyond which you would rather see a date; pass `0` to keep everything relative forever.

## Edge

Gaps under one second in either direction return `"just now"` rather than `"0 seconds ago"` or `"in 0 seconds"`, which avoids flicker near the present moment. Buckets round down: 119 seconds is "1 minute ago", not "2 minutes ago." The largest unit is weeks; there is no "months" bucket because its length is not fixed.
