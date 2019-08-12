package poeapi

import "errors"

var (
	// ErrBadRequest is raised when we have sent a malformed request to the API.
	ErrBadRequest = errors.New("bad request")

	// ErrNotFound is raised when we have requested an invalid URL.
	ErrNotFound = errors.New("url not found")

	// ErrRateLimited is raised when we exceed the API rate limits.
	ErrRateLimited = errors.New("rate limited")

	// ErrServerFailure is raised when the API returns a 5xx response for any
	// reason.
	ErrServerFailure = errors.New("server error")

	// ErrUnknownFailure is raised when an undocumented status code is returned
	// by the API. This should never occur, but is handled anyway.
	ErrUnknownFailure = errors.New("unknown server failure")

	// ErrInvalidHost is raised when an unsupported hostname is provided.
	ErrInvalidHost = errors.New("invalid API host")

	// ErrInvalidNinjaHost is raised when the poe.ninja hostname is omitted.
	ErrInvalidNinjaHost = errors.New("invalid poe.ninja host")

	// ErrInvalidRateLimit is raised when the general rate limit is out of
	// range.
	ErrInvalidRateLimit = errors.New("invalid rate limit")

	// ErrInvalidStashRateLimit is raised when the stash rate limit is out of
	// range.
	ErrInvalidStashRateLimit = errors.New("invalid stash rate limit")

	// ErrInvalidRequestTimeout is raised when request timeout is too small.
	ErrInvalidRequestTimeout = errors.New("invalid request timeout")

	// ErrInvalidCacheSize is raised when the cache size is out of range.
	ErrInvalidCacheSize = errors.New("invalid cache size")

	// ErrInvalidAddress is raised when a malformed IP address is returned by
	// the DNS cache.
	ErrInvalidAddress = errors.New("invalid ip address")

	// ErrMissingID is raised when the user fails to provide a league ID for a
	// ladder request.
	ErrMissingID = errors.New("missing league id")

	// ErrInvalidRealm is raised when the provided realm is not pc, xbox, or
	// sony.
	ErrInvalidRealm = errors.New("invalid realm")

	// ErrInvalidLimit is raised when the page size limit is out of bounds.
	ErrInvalidLimit = errors.New("invalid limit")

	// ErrInvalidOffset is raised when the page offset is out of bounds.
	ErrInvalidOffset = errors.New("invalid offset")

	// ErrInvalidLadderType is raised when the provided type is not league, pvp,
	// or labyrinth.
	ErrInvalidLadderType = errors.New("invalid ladder type")

	// ErrInvalidDifficulty is raised when the provided difficulty is not Normal,
	// Cruel, Merciless, or Eternal.
	ErrInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidLabyrinthStartTime is raised when the provided Unix timestamp
	// is below zero, above zero but earlier than the release of the Labyrinth,
	// or is an invalid timestamp.
	ErrInvalidLabyrinthStartTime = errors.New("invalid labyrinth start time")

	// ErrInvalidLeagueRuleID is raised when the ID is missing in a league rule
	// request.
	ErrInvalidLeagueRuleID = errors.New("missing league rule id")

	// ErrInvalidSeason is raised when requesting a season league with a missing
	// season name.
	ErrInvalidSeason = errors.New("invalid season")

	// ErrInvalidLeagueType is raised when requesting a league whose type is not
	// supported by the API.
	ErrInvalidLeagueType = errors.New("invalid league type")

	// ErrInvalidLeagueID is raised when the league ID is omitted from a league
	// request.
	ErrInvalidLeagueID = errors.New("invalid league id")

	// ErrInvalidStashID is raised when the stash ID is omitted from a stash
	// request.
	ErrInvalidStashID = errors.New("invalid stash id")
)
