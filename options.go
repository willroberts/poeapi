package poeapi

const (
	// API access should go to api.pathofexile.com. This may be overridden for
	// testing purposes.
	defaultHost = "api.pathofexile.com"

	// Most API endpoints have a rate limit of 5 requests per second. Tests
	// performed with the ratetest program (cmd/ratetest) show occasional
	// failures at this rate, so we back down to 4 requests per second by
	// default to err on the side of caution.
	defaultRateLimit = 4

	// The stash tab API has a rate limit of 1 request per second.
	defaultStashTabRateLimit = 1

	// A typical response from the public stash tabs API is around 3MB. By
	// default, allow around 50x3MB=150MB total cache memory usage.
	defaultCacheSize = 50
)

// ClientOptions contains settings for client initialization.
type ClientOptions struct {
	Host              string
	UseSSL            bool
	UseCache          bool
	CacheSize         int
	RateLimit         int
	StashTabRateLimit int
}

// DefaultOptions initializes the client with the most common settings.
var DefaultOptions = ClientOptions{
	Host:              defaultHost,
	UseSSL:            true,
	UseCache:          true,
	CacheSize:         defaultCacheSize,
	RateLimit:         defaultRateLimit,
	StashTabRateLimit: defaultStashTabRateLimit,
}

func validateOptions(opts ClientOptions) error {
	if opts.Host != defaultHost {
		return ErrInvalidHost
	}
	if opts.CacheSize < 1 {
		return ErrInvalidCacheSize
	}
	if opts.RateLimit < 1 {
		return ErrInvalidRateLimit
	}
	if opts.StashTabRateLimit < 1 {
		return ErrInvalidStashTabRateLimit
	}
	return nil
}
