// Package gigasecond helps to deal times in easy way using gigasenconds
package gigasecond

import "time"

const gigasecond = 1000000000

// AddGigasecond adds a gigasecond to the passed time and return the new time
func AddGigasecond(t time.Time) time.Time {
	return time.Unix(t.Unix() + gigasecond, 0)
}
