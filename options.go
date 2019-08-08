package poeapi

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
		return ErrInvalidHost
	}
	if opts.RateLimit < 1 {
		return ErrInvalidRateLimit
	}
	if opts.StashTabRateLimit < 1 {
		return ErrInvalidStashTabRateLimit
	}
	return nil
}
