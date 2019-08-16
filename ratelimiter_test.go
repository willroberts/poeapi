package poeapi

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	var (
		rateLimit    = 10.0
		requestCount uint32
		testDuration = 2.0 // Seconds.
		r            = newRateLimiter(rateLimit, DefaultStashRateLimit)
	)

	timer := time.NewTimer(time.Duration(testDuration) * time.Second)
	defer timer.Stop()

	for i := 0; i < 25; i++ {
		go func() {
			r.Wait(false)
			atomic.AddUint32(&requestCount, 1)
		}()
	}
	<-timer.C

	if int(atomic.LoadUint32(&requestCount)) > int(rateLimit*testDuration) {
		t.Fatalf("ratelimiter failed: saw %d requests in %.1f seconds (expected %d)",
			requestCount, testDuration, int(rateLimit*testDuration))
	}
}
