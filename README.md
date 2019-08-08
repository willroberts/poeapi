# poeapi

[![GoDoc Widget]][GoDoc]

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

## To Do

1. Implement all endpoints.
1. Use local HTTP server and fixtures to run tests without Internet access.

[GoDoc]: https://godoc.org/willroberts/poeapi
[GoDoc Widget]: https://godoc.org/willroberts/poeapi?status.svg