package poeapi

import (
	"sync"
	"testing"
)

func TestGetJSON(t *testing.T) {
	t.Skip() // Coverage already provided in other tests.
}

func TestGetStashTabsJSON(t *testing.T) {
	var (
		c = client{
			host:    defaultHost,
			limiter: newRateLimiter(defaultRateLimit, defaultStashTabRateLimit),
		}
		url = c.formatURL(stashTabsEndpoint)
	)
	_, err := c.getJSON(url)
	if err != nil {
		t.Fatalf("failed to get stash tabs json: %v", err)
	}
}

func TestGetJSONWithInvalidProtocol(t *testing.T) {
	var (
		c = client{
			limiter: newRateLimiter(defaultRateLimit, defaultStashTabRateLimit),
		}
		url = "htps://www.google.com"
	)
	_, err := c.getJSON(url)
	if err == nil {
		t.Fatal("failed to detect invalid http protocol")
	}
}

func TestGetJSONRateLimit(t *testing.T) {
	type errorCollector struct {
		set  []error
		lock sync.Mutex
	}
	var (
		c = client{
			host:    defaultHost,
			limiter: newRateLimiter(50, defaultStashTabRateLimit),
		}
		url  = c.formatURL(leaguesEndpoint)
		errs = errorCollector{
			set: make([]error, 0),
		}
	)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			_, err := c.getJSON(url)
			errs.lock.Lock()
			errs.set = append(errs.set, err)
			errs.lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	rateLimited := false
	for _, e := range errs.set {
		if e == ErrRateLimited {
			rateLimited = true
		}
	}
	if !rateLimited {
		t.Fatal("failed to handle rate-limited responses")
	}
}
