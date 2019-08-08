package poeapi

// APIClient provides methods for interacting with the Path of Exile API.
type APIClient interface {
	GetAllLeagues() ([]League, error)
	GetCurrentChallengeLeague() (League, error)
}

type client struct {
	host     string
	useSSL   bool
	useCache bool

	limiter *ratelimiter
	cache   *cache
}

// NewAPIClient configures and returns an APIClient.
func NewAPIClient(opts ClientOptions) (APIClient, error) {
	if err := validateOptions(opts); err != nil {
		return nil, err
	}
	c := &client{
		host:     opts.Host,
		useSSL:   opts.UseSSL,
		useCache: opts.UseCache,
		limiter:  newRateLimiter(opts.RateLimit, opts.StashTabRateLimit),
	}

	if opts.UseCache {
		cache, err := newCache(opts.CacheSize)
		if err != nil {
			return nil, err
		}
		c.cache = cache
	}

	return c, nil
}
