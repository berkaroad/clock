package main

import (
	"fmt"
	"time"

	"github.com/berkaroad/clock"
)

func main() {
	var startTs int64 = 1655431380 // 2022-06-17 10:03:00
	var refClock clock.Clock = clock.NewSystemClock()
	clock.SetGlobal(clock.NewForwardOnlyClock(startTs, refClock))

	for i := 0; i < 100; i++ {
		<-time.After(time.Second)
		fmt.Printf("system clock: %v, global clock: %v\n", refClock.Now(), clock.Now())
	}
}
