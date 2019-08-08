package poeapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const rateLimitTimeout = 5 * time.Second

func (c *apiClient) getJSON(url string) (string, error) {
	if err := c.limiter.Wait(); err != nil {
		log.Println("failed to get json:", err)
		return "", err
	}

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
		// Skip ahead.
	case http.StatusBadRequest:
		return "", ErrBadRequest
	case http.StatusNotFound:
		return "", ErrNotFound
	case http.StatusTooManyRequests:
		// Rate limiter is not working if we reach this block.
		// Back off and retry.
		time.Sleep(rateLimitTimeout)
		return c.getJSON(url)
	case http.StatusInternalServerError:
		// While 5xx errors should be retried, we leave this up to the caller.
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
