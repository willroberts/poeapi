package poeapi

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	var (
		rateLimit    = 10
		requestCount = 0
		testDuration = 2 // Seconds.
		r            = newRateLimiter(rateLimit, rateLimit)
	)

	timer := time.NewTimer(time.Duration(testDuration) * time.Second)
	defer timer.Stop()

	for i := 0; i < 20; i++ {
		r.wait(rateLimit)
		requestCount++
	}
	<-timer.C

	if requestCount > (rateLimit * testDuration) {
		t.Fatalf("ratelimiter failed: saw %d requests in %d seconds (expected %d)",
			requestCount, testDuration, rateLimit*testDuration)
	}
}

func TestRateLimiterTooFast(t *testing.T) {
	var (
		rateLimit    = 1
		requestCount uint32
		testDuration = 5 // Seconds.
		r            = newRateLimiter(rateLimit, rateLimit)
	)

	timer := time.NewTimer(time.Duration(testDuration) * time.Second)
	defer timer.Stop()

	for i := 0; i < 10; i++ {
		go func() {
			r.wait(rateLimit)
			atomic.AddUint32(&requestCount, 1)
		}()
	}
	<-timer.C

	if int(atomic.LoadUint32(&requestCount)) > (rateLimit * testDuration) {
		t.Fatalf("ratelimiter failed: saw %d requests in %d seconds (expected %d)",
			requestCount, testDuration, rateLimit*testDuration)
	}
}
