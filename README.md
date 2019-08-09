# poeapi

[![GoDoc Badge]][GoDoc]
[![GoReportCard Badge]][GoReportCard]
[![License Badge]][License]

A Go client for the Path of Exile API.

## Features

* Supports every endpoint of the Path of Exile API
* All operations are thread-safe
* Built-in, tunable rate limiting
* Built-in, tunable caching for responses
* 100% standard library code

## Examples

There are some examples in the `cmd` directory.

#### leaguetimer

This example retrieves the current challenge league from the API and prints how
much time it has remaining.

#### ladderstats

This example retrieves the ladder or leaderboard for a league, computes the
distribution of character classes, and prints the results.

## To Do

1. Implement all endpoints. GetLeague for single leagues, query params for all
   leagues endpoints. 
1. Use local HTTP server and fixtures to run tests without Internet access.
1. Open source the repo to enable GoDoc and Go Report Card
1. Add TravisCI to get a build passing badge.
1. Add CodeCov to get a coverage badge.

[GoDoc]: https://godoc.org/willroberts/poeapi
[GoDoc Badge]: https://godoc.org/willroberts/poeapi?status.svg
[GoReportCard]: https://goreportcard.com/report/github.com/willroberts/poeapi
[GoReportCard Badge]: https://goreportcard.com/badge/github.com/willroberts/poeapi
[License]: https://www.gnu.org/licenses/gpl-3.0
[License Badge]: https://img.shields.io/badge/License-GPLv3-blue.svg