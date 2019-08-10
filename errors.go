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

// ErrInvalidStashRateLimit is raised when the stash rate limit is out of
// range.
var ErrInvalidStashRateLimit = errors.New("invalid stash rate limit")

// ErrInvalidCacheSize is raised when the cache size is out of range.
var ErrInvalidCacheSize = errors.New("invalid cache size")

// ErrMissingID is raised when the user fails to provide a league ID for a
// ladder request.
var ErrMissingID = errors.New("missing league id")

// ErrInvalidRealm is raised when the provided realm is not pc, xbox, or sony.
var ErrInvalidRealm = errors.New("invalid realm")

// ErrInvalidLimit is raised when the page size limit is out of bounds.
var ErrInvalidLimit = errors.New("invalid limit")

// ErrInvalidOffset is raised when the page offset is out of bounds.
var ErrInvalidOffset = errors.New("invalid offset")

// ErrInvalidLadderType is raised when the provided type is not league, pvp, or
// labyrinth.
var ErrInvalidLadderType = errors.New("invalid ladder type")

// ErrInvalidDifficulty is raised when the provided difficulty is not Normal,
// Cruel, Merciless, or Eternal.
var ErrInvalidDifficulty = errors.New("invalid difficulty")

// ErrInvalidLabyrinthStartTime is raised when the provided Unix timestamp is
// below zero, above zero but earlier than the release of the Labyrinth, or
// is an invalid timestamp.
var ErrInvalidLabyrinthStartTime = errors.New("invalid labyrinth start time")

// ErrInvalidLeagueRuleID is raised when the ID is missing in a league rule
// request.
var ErrInvalidLeagueRuleID = errors.New("missing league rule id")

// ErrInvalidSeason is raised when requesting a season league with a missing
// season name.
var ErrInvalidSeason = errors.New("invalid season")

// ErrInvalidLeagueType is raised when requesting a league whose type is not
// supported by the API.
var ErrInvalidLeagueType = errors.New("invalid league type")

// ErrInvalidLeagueID is raised when the league ID is omitted from a league
// request.
var ErrInvalidLeagueID = errors.New("invalid league id")

// ErrInvalidStashID is raised when the stash ID is omitted from a stash
// request.
var ErrInvalidStashID = errors.New("invalid stash id")
