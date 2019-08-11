# poeapi

[![GoDoc Badge]][GoDoc]
[![GoReportCard Badge]][GoReportCard]
[![License Badge]][License]

A Go client for the Path of Exile API.

## Features

* Supports every endpoint of the [Path of Exile API](https://www.pathofexile.com/developer/docs/api)
* All operations are thread-safe
* Built-in, tunable rate limiting
* Built-in, tunable caching for responses
* No dependencies; 100% standard library code

## Usage

```go
clientOpts := poeapi.ClientOptions{
    Host:              "api.pathofexile.com",
    NinjaHost:         "poe.ninja",
    UseSSL:            true,
    UseCache:          true,
    CacheSize:         50,
    RateLimit:         4,
    StashRateLimit:    1,
} // This is equivalent to poeapi.DefaultClientOptions.

client, err := poeapi.NewAPIClient(clientOpts)
if err != nil {
    // ...
}

ladder, err := client.GetLadder(poeapi.GetLadderOptions{
    ID:    "SSF Hardcore",
    Realm: "pc",
    Type:  "league",
})
// ...
```

## Interface

These are the methods available on the client's interface:

```go
GetLeagues() ([]poeapi.League, error)
GetCurrentChallengeLeague() (poeapi.League, error)
GetLeagueRules() ([]poeapi.LeagueRule, error)
GetLeagueRule(poeapi.GetLeagueRuleOptions) (poeapi.LeagueRule, error)
GetLadder(poeapi.GetLadderOptions) (poeapi.Ladder, error)
GetPVPMatches(poeapi.GetPVPMatchesOptions) ([]poeapi.PVPMatch, error)
GetStashes(opts GetStashOptions) (StashResponse, error)
GetLatestStashID() (string, error)
```

See the [documentation](https://godoc.org/willroberts/poeapi) or examples for
more usage information.

## Examples

There are several examples in the `examples` directory.

#### leaguetimer

This example retrieves the current challenge league from the API and prints how
much time it has remaining.

#### listleaguerules

This example retrieves all league rules from the API and prints their names and
descriptions.

#### ladderstats

This example retrieves the ladder or leaderboard for a league, computes the
distribution of character classes, and prints the results.

#### upcomingmatches

This example retrieves all upcoming PVP matches from the API and prints how many
events are scheduled (hint: it's zero).

#### itemnotifier

This example searches in real time, until the user exits with Ctrl-C, for Kaom's
Heart in Standard league. When one is listed for sale, it prints the character
name and asking price (if there is one).

## To Do
1. Open source the repo to enable GoDoc and Go Report Card.
1. Add TravisCI to get a build passing badge.
1. Add CodeCov to get a coverage badge.
1. Improve documentation, test coverage, etc.

[GoDoc]: https://godoc.org/willroberts/poeapi
[GoDoc Badge]: https://godoc.org/willroberts/poeapi?status.svg
[GoReportCard]: https://goreportcard.com/report/github.com/willroberts/poeapi
[GoReportCard Badge]: https://goreportcard.com/badge/github.com/willroberts/poeapi
[License]: https://www.gnu.org/licenses/gpl-3.0
[License Badge]: https://img.shields.io/badge/License-GPLv3-blue.svg