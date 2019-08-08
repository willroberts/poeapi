package poeapi

import "time"

type janklimiter struct {
	rateLimit         int
	stashTabRateLimit int

	lastRequest time.Time
}

// wait blocks execution until enough time has elasped since the last request.
func (j *janklimiter) wait(rate int) {
	interval := time.Duration(1000.0/rate) * time.Millisecond

	elapsed := time.Since(j.lastRequest)
	if elapsed < interval {
		time.Sleep(interval - elapsed)
	}

	j.lastRequest = time.Now()
}

func newJankLimiter(rateLimit, stashTabRateLimit int) *janklimiter {
	return &janklimiter{
		rateLimit:         rateLimit,
		stashTabRateLimit: stashTabRateLimit,
	}
}
