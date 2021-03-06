package poeapi

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type requestFunc func(string) (string, error)

// Get is a helper function which includes caching and ratelimiting for outbound
// requests.
func (c *client) get(url string) (string, error) {
	return c.withCache(url, c.withRateLimit(url, c.getJSON))
}

// getJSON retrieves the given URL. It returns the JSON response as a string.
func (c *client) getJSON(url string) (string, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		// An error is returned if the Client's CheckRedirect function fails or
		// if there was an HTTP protocol error. A non-2xx response doesn't cause
		// an error.
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", parseError(resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *client) withCache(url string, fn requestFunc) (string, error) {
	if !c.useCache {
		return fn(url)
	}

	// Disable caching for stash endpoint.
	if strings.HasPrefix(url, c.formatURL(stashTabsEndpoint)) {
		return fn(url)
	}

	if cached, err := c.cache.Get(url); err == nil {
		return cached, nil
	}

	resp, err := fn(url)
	if err != nil {
		return "", err
	}

	c.cache.Set(url, resp)
	return resp, nil
}

func (c *client) withRateLimit(url string, fn requestFunc) requestFunc {
	if strings.HasPrefix(url, c.formatURL(stashTabsEndpoint)) {
		c.limiter.Wait(true)
		return fn
	}
	c.limiter.Wait(false)
	return fn
}

func parseError(statusCode int) error {
	switch statusCode {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusTooManyRequests:
		return ErrRateLimited
	case http.StatusInternalServerError:
		return ErrServerFailure
	default:
		return ErrUnknownFailure
	}
}
