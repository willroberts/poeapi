package poeapi

import (
	"io/ioutil"
	"net/http"
	"time"
)

const rateLimitTimeout = 2 * time.Second

func (c *client) getJSON(url string) (string, error) {
	var ratelimit int
	if url == c.formatURL(stashTabsEndpoint) {
		ratelimit = c.limiter.rateLimit
	} else {
		ratelimit = c.limiter.stashTabRateLimit
	}

	c.limiter.wait(ratelimit)
	resp, err := http.Get(url)
	if err != nil {
		// An error is returned if the Client's CheckRedirect function fails or
		// if there was an HTTP protocol error. A non-2xx response doesn't cause
		// an error.
		return "", err
	}
	defer resp.Body.Close()

	// Check status code for HTTP error handling.
	switch resp.StatusCode {
	case http.StatusOK:
		// Continue.
	case http.StatusBadRequest:
		return "", ErrBadRequest
	case http.StatusNotFound:
		return "", ErrNotFound
	case http.StatusTooManyRequests:
		return "", ErrRateLimited
	case http.StatusInternalServerError:
		return "", ErrServerFailure
	default:
		return "", ErrUnknownFailure
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
