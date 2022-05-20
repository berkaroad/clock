package clock

import (
	"sync"
	"sync/atomic"
	"time"
)

// NewForwardOnlyClock clock that starts at init unix time and will go forward reference clock.
//
// Use zero if initUnixTime is negative integer, else use 'referenceClock()'.
//
// Use 'NewSystemClock()' if referenceClock is nil.
func NewForwardOnlyClock(initUnixTime int64, referenceClock Clock) Clock {
	if referenceClock == nil {
		referenceClock = NewSystemClock()
	}
	now := referenceClock()
	if initUnixTime < 0 {
		if now < 0 {
			initUnixTime = 0
		} else {
			initUnixTime = now
		}
	}
	clock := &forwardOnlyClock{current: initUnixTime, refClock: referenceClock}
	offset := now - initUnixTime
	if offset > 0 {
		clock.offset = offset
	}
	clock.run()
	return clock.Clock()
}

type forwardOnlyClock struct {
	current  int64
	offset   int64
	refClock Clock
	once     sync.Once
}

func (c *forwardOnlyClock) Clock() Clock {
	return func() int64 {
		return c.current + c.offset
	}
}

func (c *forwardOnlyClock) run() {
	c.once.Do(func() {
		go func() {
			for {
				<-time.After(time.Second)
				atomic.AddInt64(&c.current, 1)
				go func() {
					lastOffset := atomic.LoadInt64(&c.offset)
					currentOffset := c.refClock() - atomic.LoadInt64(&c.current)
					if currentOffset > lastOffset {
						atomic.CompareAndSwapInt64(&c.offset, lastOffset, currentOffset)
					}
				}()
			}
		}()
	})
}
