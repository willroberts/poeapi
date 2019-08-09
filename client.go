package poeapi

// APIClient provides methods for interacting with the Path of Exile API.
type APIClient interface {
	GetAllLeagues() ([]League, error)
	GetCurrentChallengeLeague() (League, error)

	// GetLadder retrieves the entire leaderboard for a league. For ladders with
	// the maximum 15,000 entries, this function may take 15-20 seconds to run.
	GetLadder(GetLadderOptions) (Ladder, error)
}

// client is the implementation of the APIClient interface.
type client struct {
	host     string
	useSSL   bool
	useCache bool

	limiter *ratelimiter
	cache   *cache
}

// NewAPIClient configures and returns an APIClient.
func NewAPIClient(opts ClientOptions) (APIClient, error) {
	if err := validateClientOptions(opts); err != nil {
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
