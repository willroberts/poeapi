package poeapi

import (
	"sync"
	"time"
)

const (
	// UnlimitedRate disables rate limiting when used as a rate limit.
	UnlimitedRate = 0
)

// ratelimiter uses blocking time.Sleep calls to prevent callers from sending
// requests too frequently. ratelimiter is threadsafe.
type ratelimiter struct {
	rateLimit      float64
	stashRateLimit float64

	lastRequest      time.Time
	lastStashRequest time.Time

	lock      sync.Mutex
	stashLock sync.Mutex
}

// Wait blocks execution until enough time has elasped since the last request.
func (r *ratelimiter) Wait(stash bool) {
	if stash {
		r.stashLock.Lock()
		defer r.stashLock.Unlock()
		r.waitLimit(r.stashRateLimit, r.lastStashRequest)
		r.lastStashRequest = time.Now()
		return
	}

	r.lock.Lock()
	defer r.lock.Unlock()
	r.waitLimit(r.rateLimit, r.lastRequest)
	r.lastRequest = time.Now()
}

func (r *ratelimiter) waitLimit(ratelimit float64, last time.Time) {
	if ratelimit == UnlimitedRate {
		return
	}
	interval := time.Duration(1000.0/ratelimit) * time.Millisecond
	elapsed := time.Since(last)
	if elapsed < interval {
		time.Sleep(interval - elapsed)
	}
}

func newRateLimiter(rateLimit, stashRateLimit float64) *ratelimiter {
	return &ratelimiter{
		rateLimit:      rateLimit,
		stashRateLimit: stashRateLimit,
	}
}
