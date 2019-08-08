package poeapi

import (
	"sync"
	"time"
)

// ratelimiter uses blocking time.Sleep calls to prevent callers from sending
// requests too frequently. ratelimiter is threadsafe.
type ratelimiter struct {
	rateLimit         int
	stashTabRateLimit int

	lastRequest time.Time
	lock        sync.Mutex
}

// wait blocks execution until enough time has elasped since the last request.
func (r *ratelimiter) wait(rate int) {
	r.lock.Lock()
	defer r.lock.Unlock()

	interval := time.Duration(1000.0/rate) * time.Millisecond

	elapsed := time.Since(r.lastRequest)
	if elapsed < interval {
		time.Sleep(interval - elapsed)
	}

	r.lastRequest = time.Now()
}

func newRateLimiter(rateLimit, stashTabRateLimit int) *ratelimiter {
	return &ratelimiter{
		rateLimit:         rateLimit,
		stashTabRateLimit: stashTabRateLimit,
	}
}
