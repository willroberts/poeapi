package poeapi

// APIClient provides methods for interacting with the Path of Exile API.
type APIClient interface {
	GetAllLeagues() ([]League, error)
	GetCurrentChallengeLeague() (League, error)
}

type client struct {
	host   string
	useSSL bool

	limiter *ratelimiter
}

// NewAPIClient configures and returns an APIClient.
func NewAPIClient(opts ClientOptions) (APIClient, error) {
	if err := validateOptions(opts); err != nil {
		return nil, err
	}
	return &client{
		host:    opts.Host,
		useSSL:  opts.UseSSL,
		limiter: newRateLimiter(opts.RateLimit, opts.StashTabRateLimit),
	}, nil
}
