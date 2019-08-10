package poeapi

import (
	"sync"
	"time"
)

// ratelimiter uses blocking time.Sleep calls to prevent callers from sending
// requests too frequently. ratelimiter is threadsafe.
type ratelimiter struct {
	rateLimit         int
	stashRateLimit int

	lastRequest      time.Time
	lastStashRequest time.Time
	lock             sync.Mutex
}

// wait blocks execution until enough time has elasped since the last request.
func (r *ratelimiter) wait(stash bool) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var interval time.Duration
	if stash {
		interval = time.Duration(1000.0/r.stashRateLimit) * time.Millisecond
		elapsed := time.Since(r.lastStashRequest)
		if elapsed < interval {
			time.Sleep(interval - elapsed)
		}
		r.lastStashRequest = time.Now()
	} else {
		interval = time.Duration(1000.0/r.rateLimit) * time.Millisecond
		elapsed := time.Since(r.lastRequest)
		if elapsed < interval {
			time.Sleep(interval - elapsed)
		}
		r.lastRequest = time.Now()
	}
}

func newRateLimiter(rateLimit, stashRateLimit int) *ratelimiter {
	return &ratelimiter{
		rateLimit:         rateLimit,
		stashRateLimit: stashRateLimit,
	}
}
