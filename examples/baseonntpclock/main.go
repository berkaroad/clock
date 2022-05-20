package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
	"github.com/berkaroad/clock"
)

func main() {
	var startTs int64 = 1655431380 // 2022-06-17 10:03:00
	var refClock clock.Clock = clock.NewNTPClock([]string{"pool.ntp.org"}, ntp.QueryOptions{Timeout: time.Second})
	clock.SetGlobal(clock.NewForwardOnlyClock(startTs, refClock))

	for i := 0; i < 100; i++ {
		<-time.After(time.Second)
		fmt.Printf("ntp clock: %v, global clock: %v\n", refClock.Now(), clock.Now())
	}
}
