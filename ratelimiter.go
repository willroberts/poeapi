package poeapi

// Implementations of ratelimit, ratelimiter, and reservation sourced from the
// upcoming golang.org/x/time/rate package.

import (
	"errors"
	"math"
	"sync"
	"time"
)

const (
	infinite         = ratelimit(math.MaxFloat64)
	infiniteDuration = time.Duration(1<<63 - 1)
)

type ratelimit float64

func (limit ratelimit) durationFromTokens(tokens float64) time.Duration {
	seconds := tokens / float64(limit)
	return time.Nanosecond * time.Duration(1e9*seconds)
}

func (limit ratelimit) tokensFromDuration(d time.Duration) float64 {
	return d.Seconds() * float64(limit)
}

type ratelimiter struct {
	limit ratelimit
	burst int

	mu              sync.Mutex
	tokens          float64
	lastEvent       time.Time
	lastTokenUpdate time.Time
}

func newRatelimiter(r ratelimit, b int) *ratelimiter {
	return &ratelimiter{
		limit: r,
		burst: b,
	}
}

func (r *ratelimiter) advance(now time.Time) (time.Time, time.Time, float64) {
	last := r.lastTokenUpdate
	if now.Before(last) {
		last = now
	}

	maxElapsed := r.limit.durationFromTokens(float64(r.burst) - r.tokens)
	elapsed := now.Sub(last)
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	delta := r.limit.tokensFromDuration(elapsed)
	tokens := r.tokens + delta
	if burst := float64(r.burst); tokens > burst {
		tokens = burst
	}

	return now, last, tokens
}

func (r *ratelimiter) reserve(now time.Time, n int, maxFutureReserve time.Duration) reservation {
	r.mu.Lock()

	if r.limit == infinite {
		r.mu.Unlock()
		return reservation{
			ok:        true,
			lim:       r,
			tokens:    n,
			timeToAct: now,
		}
	}

	now, last, tokens := r.advance(now)
	tokens -= float64(1)

	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = r.limit.durationFromTokens(-tokens)
	}

	ok := n <= r.burst && waitDuration <= maxFutureReserve
	res := reservation{
		ok:    ok,
		lim:   r,
		limit: r.limit,
	}
	if ok {
		res.tokens = n
		res.timeToAct = now.Add(waitDuration)
	}
	if ok {
		r.lastTokenUpdate = now
		r.tokens = tokens
		r.lastEvent = res.timeToAct
	} else {
		r.lastTokenUpdate = last
	}

	r.mu.Unlock()
	return res
}

func (r *ratelimiter) Wait() error {
	if 1 > r.burst && r.limit != infinite {
		return errors.New("request exceeds burst limit")
	}

	now := time.Now()
	waitLimit := infiniteDuration

	res := r.reserve(now, 1, waitLimit)
	if !res.ok {
		return errors.New("context deadline exceeded")
	}

	delay := res.DelayFrom(now)
	if delay == 0 {
		return nil
	}

	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	}
}

type reservation struct {
	ok        bool
	lim       *ratelimiter
	tokens    int
	timeToAct time.Time
	limit     ratelimit
}

func (r *reservation) DelayFrom(now time.Time) time.Duration {
	if !r.ok {
		return infiniteDuration
	}
	delay := r.timeToAct.Sub(now)
	if delay < 0 {
		return 0
	}
	return delay
}
