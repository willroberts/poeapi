package poeapi

import (
	"net/http"
	"sync"
	"testing"
)

func TestGetJSON(t *testing.T) {
	var (
		c = client{
			host:       testHost,
			limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
			httpClient: testClient,
		}
		url = c.formatURL(leaguesEndpoint)
	)
	_, err := c.getJSON(url)
	if err != nil {
		t.Fatalf("failed to get json: %v", err)
	}
}

func TestGetStashsJSON(t *testing.T) {
	var (
		c = client{
			host:       testHost,
			limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
			httpClient: testClient,
		}
		url = c.formatURL(stashTabsEndpoint)
	)
	_, err := c.getJSON(url)
	if err != nil {
		t.Fatalf("failed to get stash json: %v", err)
	}
}

func TestGetJSONWithInvalidProtocol(t *testing.T) {
	var (
		c = client{
			host:       testHost,
			limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
			httpClient: testClient,
		}
		url = "htps://127.0.0.1:8000"
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
			host:       testHost,
			limiter:    newRateLimiter(50, UnlimitedRate),
			httpClient: testClient,
		}
		url  = c.formatURL(rateLimitEndpoint)
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
			break
		}
	}
	if !rateLimited {
		t.Fatal("failed to handle rate-limited responses")
	}
}

func TestHandleServerError(t *testing.T) {
	c := client{
		host:       testHost,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	_, err := c.getJSON(c.formatURL(failureEndpoint))
	if err != ErrServerFailure {
		t.Fatal("failed to handle server error")
	}
}

func TestWithRateLimit(t *testing.T) {
	var (
		c = client{
			host:       testHost,
			limiter:    newRateLimiter(DefaultRateLimit, UnlimitedRate),
			httpClient: testClient,
		}
		url = "https://127.0.0.1:8000"
		fn  = func(s string) (string, error) { return s, nil }
	)
	_ = c.withRateLimit(url, fn)
}

func TestWithStashRateLimit(t *testing.T) {
	var (
		c = client{
			host:       testHost,
			useSSL:     true,
			limiter:    newRateLimiter(UnlimitedRate, DefaultStashRateLimit),
			httpClient: testClient,
		}
		url = "https://127.0.0.1:8000/public-stash-tabs"
		fn  = func(s string) (string, error) { return s, nil }
	)
	_ = c.withRateLimit(url, fn)
}

func TestParseCodeBadRequest(t *testing.T) {
	if err := parseError(http.StatusBadRequest); err != ErrBadRequest {
		t.Fatal("failed to detect bad request")
	}
}

func TestParseCodeUnknownFailure(t *testing.T) {
	if err := parseError(http.StatusTeapot); err != ErrUnknownFailure {
		t.Fatal("failed to detect unknown error")
	}
}

func TestCacheHelperWithStashURL(t *testing.T) {
	var (
		c = client{
			host:     testHost,
			useCache: true,
		}
		url = c.formatURL(stashTabsEndpoint)
		fn  = func(s string) (string, error) {
			return "", nil
		}
	)
	if _, err := c.withCache(url, fn); err != nil {
		t.Fatal("failed to bypass cache in decorated function")
	}
}

func TestCacheHelperWithErrorResult(t *testing.T) {
	cache, err := newCache(10)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}
	var (
		c = client{
			host:     testHost,
			useCache: true,
			cache:    cache,
		}
		url = c.formatURL(leaguesEndpoint)
		fn  = func(s string) (string, error) {
			return "", ErrUnknownFailure
		}
	)
	if _, err := c.withCache(url, fn); err != ErrUnknownFailure {
		t.Fatal("failed to detect error in decorated function")
	}
}
