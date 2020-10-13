//Package gigasecond provides a method for adding a gigasecond to the time type
package gigasecond

import "time"

const gigasecond = 1e9 * time.Second

//AddGigasecond adds a gigasecond given the time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
