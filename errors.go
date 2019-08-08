package poeapi

import "errors"

// See https://www.pathofexile.com/developer/docs/api-errors.

// ErrBadRequest is raised when we have sent a malformed request to the API.
var ErrBadRequest = errors.New("bad request")

// ErrNotFound is raised when we have requested an invalid URL.
var ErrNotFound = errors.New("url not found")

// ErrRateLimited is raised when we exceed the API rate limits.
var ErrRateLimited = errors.New("rate limited")

// ErrServerFailure is raised when the API returns a 5xx response for any
// reason.
var ErrServerFailure = errors.New("server error")

// ErrUnknownFailure is raised when an undocumented status code is returned by
// the API. This should never occur, but is handled anyway.
var ErrUnknownFailure = errors.New("unknown server failure")

// ErrInvalidHost is raised when an unsupported hostname is provided.
var ErrInvalidHost = errors.New("invalid API host")

// ErrInvalidRateLimit is raised when the general rate limit is out of range.
var ErrInvalidRateLimit = errors.New("invalid rate limit")

// ErrInvalidStashTabRateLimit is raised when the stash tab rate limit is out of
// range.
var ErrInvalidStashTabRateLimit = errors.New("invalid stash tab rate limit")
