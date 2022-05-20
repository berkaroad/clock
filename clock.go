package clock

import (
	"sync"
	"time"
)

var globalClock Clock
var once sync.Once

// SetGlobal set global clock.
func SetGlobal(clock Clock) {
	if clock != nil {
		once.Do(func() {
			globalClock = clock
		})
	}
}

// Now returns global clock's current local time.
func Now() time.Time {
	clock := globalClock
	if clock == nil {
		return time.Now()
	}
	return clock.Now()
}

// Clock is a function that returns unix time.
type Clock func() int64

// Now returns clock's current local time.
func (c Clock) Now() time.Time {
	return time.Unix(c(), 0)
}

// NewSystemClock returns clock from os.
func NewSystemClock() Clock {
	return func() int64 {
		return time.Now().Unix()
	}
}
