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

// APIClient provides methods for interacting with the Path of Exile API.
type APIClient interface {
	GetLeagues(GetLeaguesOptions) ([]League, error)
	GetCurrentChallengeLeague() (League, error)
	GetLeagueRules() ([]LeagueRule, error)
	GetLeagueRule(GetLeagueRuleOptions) (LeagueRule, error)
	GetLadder(GetLadderOptions) (Ladder, error)
	GetPVPMatches(GetPVPMatchesOptions) ([]PVPMatch, error)
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

// ClientOptions contains settings for client initialization.
type ClientOptions struct {
	Host              string
	UseSSL            bool
	UseCache          bool
	CacheSize         int
	RateLimit         int
	StashTabRateLimit int
}

// DefaultClientOptions initializes the client with the most common settings.
var DefaultClientOptions = ClientOptions{
	Host:              defaultHost,
	UseSSL:            true,
	UseCache:          true,
	CacheSize:         defaultCacheSize,
	RateLimit:         defaultRateLimit,
	StashTabRateLimit: defaultStashTabRateLimit,
}

func validateClientOptions(opts ClientOptions) error {
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
