package poeapi

import "errors"

// ClientOptions contains settings for client initialization.
type ClientOptions struct {
	Host              string
	UseSSL            bool
	RateLimit         int
	StashTabRateLimit int
}

// DefaultOptions initializes the client with the most common settings.
var DefaultOptions = ClientOptions{
	Host:              "api.pathofexile.com",
	UseSSL:            true,
	RateLimit:         4,
	StashTabRateLimit: 1,
}

func validateOptions(opts ClientOptions) error {
	if opts.Host != "api.pathofexile.com" {
		return errors.New("unsupported API host")
	}
	if opts.RateLimit < 1 || opts.StashTabRateLimit < 1 {
		return errors.New("rate limits must be 1 or higher")
	}
	return nil
}
