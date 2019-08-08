package poeapi

import "time"

// ratelimiter uses blocking time.Sleep calls to prevent callers from sending
// requests too frequently.
type ratelimiter struct {
	rateLimit         int
	stashTabRateLimit int

	lastRequest time.Time
}

// wait blocks execution until enough time has elasped since the last request.
func (r *ratelimiter) wait(rate int) {
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
