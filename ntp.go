package clock

import (
	"time"

	"github.com/beevik/ntp"
)

// NewNTPClock returns clock from ntp server.
func NewNTPClock(ntpServers []string, opt ntp.QueryOptions) Clock {
	return func() int64 {
		var offset time.Duration
		for _, host := range ntpServers {
			resp, err := ntp.QueryWithOptions(host, opt)
			if err == nil {
				offset = resp.ClockOffset
				break
			}
		}
		return time.Now().Add(offset).Unix()
	}
}
