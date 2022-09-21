# poeapi

[![License Badge]][License]
[![Build Badge]][Build]
[![CodeCov Badge]][CodeCov]
[![GoReportCard Badge]][GoReportCard]
[![GoDoc Badge]][GoDoc]

A Go client for the Path of Exile API.

## Features

* Supports every endpoint of the [Path of Exile API][API Docs]
* All operations are thread-safe
* Built-in, tunable rate limiting
* Built-in, tunable caching for responses
* No dependencies; 100% standard library code

## Usage

```go
clientOpts := poeapi.ClientOptions{
    Host:           "api.pathofexile.com", // The primary API domain.
    NinjaHost:      "poe.ninja",           // Used to get latest stash ID.
    UseSSL:         true,                  // Use HTTPS for requests.
    UseCache:       true,                  // Enable the in-memory cache.
    UseDNSCache:    true,                  // Enable the in-memory DNS resolution cache.
    CacheSize:      200,                   // Number of items to store.
    RateLimit:      4.0,                   // Requests per second.
    StashRateLimit: 1.0,                   // Requests per second for trade API.
    RequestTimeout: 5 * time.Second        // Time to wait before canceling requests.
} // This is equivalent to poeapi.DefaultClientOptions.

client, err := poeapi.NewAPIClient(clientOpts)
if err != nil {
    // Handle error.
}

ladder, err := client.GetLadder(poeapi.GetLadderOptions{
    ID:    "SSF Hardcore",
    Realm: "pc",
    Type:  "league",
})
// Etc.
```

## Interface

These are the methods available on the client's interface:

```go
// Method:                                 Returns:
GetLadder(poeapi.GetLadderOptions)         (poeapi.Ladder, error)
GetLeague(poeapi.GetLeagueOptions)         (poeapi.League, error)
GetLeagueRule(poeapi.GetLeagueRuleOptions) (poeapi.LeagueRule, error)
GetLeagueRules()                           ([]poeapi.LeagueRule, error)
GetLeagues(poeapi.GetLeaguesOptions)       ([]poeapi.League, error)
GetPVPMatches(poeapi.GetPVPMatchesOptions) ([]poeapi.PVPMatch, error)
GetStashes(poeapi.GetStashOptions)         (poeapi.StashResponse, error)
GetLatestStashID()                         (string, error)
```

See the [documentation][GoDoc] or [examples][Examples] for more usage information.

## Examples

There are several examples in the [examples][Examples] directory.

## Contributing

Contributions to the code, examples, or documentation are very welcome.

Found a problem? Have a question? [Open an issue][Issue]

Have a patch or a new feature? [Open a pull request][Pull Request]

[License]: https://www.gnu.org/licenses/gpl-3.0
[License Badge]: https://img.shields.io/badge/License-GPLv3-blue.svg
[Build]: https://github.com/willroberts/poeapi/actions/workflows/build.yaml
[Build Badge]: https://github.com/willroberts/poeapi/actions/workflows/build.yaml/badge.svg
[CodeCov]: https://codecov.io/gh/willroberts/poeapi
[CodeCov Badge]: https://codecov.io/gh/willroberts/poeapi/branch/main/graph/badge.svg
[GoReportCard]: https://goreportcard.com/report/github.com/willroberts/poeapi
[GoReportCard Badge]: https://goreportcard.com/badge/github.com/willroberts/poeapi
[GoDoc]: https://pkg.go.dev/github.com/willroberts/poeapi
[GoDoc Badge]: https://pkg.go.dev/badge/github.com/willroberts/poeapi

[API Docs]: https://www.pathofexile.com/developer/docs/reference
[Examples]: https://github.com/willroberts/poeapi/tree/main/examples
[Issue]: https://github.com/willroberts/poeapi/issues
[Pull Request]: https://github.com/willroberts/poeapi/pulls
