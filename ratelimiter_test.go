package poeapi

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	var (
		rateLimit    = 4
		requestCount = 0
		testDuration = 5 // Seconds.
	)
	r := newRateLimiter(rateLimit, rateLimit)
	timer := time.NewTimer(time.Duration(testDuration) * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		break
	default:
		requestCount++
		r.wait(rateLimit)
	}

	if requestCount > (rateLimit * testDuration) {
		t.Fail()
	}
}

func TestRateLimiterTooFast(t *testing.T) {
	var (
		rateLimit    = 1
		requestCount = 0
		testDuration = 5 // Seconds.
	)
	r := newRateLimiter(rateLimit, rateLimit)
	timer := time.NewTimer(time.Duration(testDuration) * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		break
	default:
		for i := 0; i < 5; i++ {
			go func() {
				r.wait(rateLimit)
			}()
			requestCount++
		}
	}

	if requestCount > (rateLimit * testDuration) {
		t.Fail()
	}
}
