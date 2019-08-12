package poeapi

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultHost sets the hostname used by the client. This is set to
	// api.pathofexile.com, but may be overridden for testing purposes.
	DefaultHost = "api.pathofexile.com"

	// DefaultNinjaHost sets the hostname used to retrieve the latest stash
	// change ID from poe.ninja. This is set to "poe.ninja" but may be
	// overridden for testing purposes.
	DefaultNinjaHost = "poe.ninja"

	// DefaultRateLimit sets the rate limit for all endpoints except for the
	// stash endpoint. Most endpoints have a rate limit of 5 requests per
	// second. Tests performed with the ratetest program (cmd/ratetest) show
	// occasional failures at this rate, so we back down to 4 requests per
	// second by default to err on the side of caution.
	DefaultRateLimit = 4.0

	// DefaultStashRateLimit sets the rate limit for the stash endpoint.
	// The stash API has a rate limit of 1 request per second.
	DefaultStashRateLimit = 1.0

	// DefaultCacheSize sets the number of items which can be stores in the
	// in-memory LRU cache. A typical response from the API is around 500KB.
	// By default, allow around 200x500KB=100MB total cache memory usage.
	DefaultCacheSize = 200

	// DefaultRequestTimeout sets the time to wait before canceling HTTP
	// requests. Some endpoits take over 2-3s to respond, so we use 5s as a
	// a default.
	DefaultRequestTimeout = 5 * time.Second
)

// APIClient provides methods for interacting with the Path of Exile API.
type APIClient interface {
	// GetLadder sends multiple ladder requests to construct the entire ladder
	// for a given league in a single call. Ladders contain information about
	// the top-ranked characters in a given league. Up to 15,000 characters may
	// be returned.
	GetLadder(GetLadderOptions) (Ladder, error)

	// GetLeague retrieves all league (Standard, Hardcore, etc.) from the API.
	// Responses include information such as start and end times and rules for
	// the league.
	GetLeagues(GetLeaguesOptions) ([]League, error)

	// GetLeague retrieves a single league from the API by ID.
	GetLeague(GetLeagueOptions) (League, error)

	// GetLeagueRules retrieves all available league modifiers from the API.
	// These modifiers affect league mechanics, such as the 'Turbo' rule
	// granting increased attack, cast, and movement speed to monsters.
	GetLeagueRules() ([]LeagueRule, error)

	// GetLeagueRule retrieves a single rule from the API by ID.
	GetLeagueRule(GetLeagueRuleOptions) (LeagueRule, error)

	// GetPVPMatches retrieves past or upcoming PVP matches from the API.
	// Specific seasons may be requested in order to view past events.
	// Alternatively, not specifying a season returns all upcoming events.
	GetPVPMatches(GetPVPMatchesOptions) ([]PVPMatch, error)

	// GetStashes retrieves a batch of stashes from the trade API. Each response
	// contains a set of stashes which can be parsed for specific items.
	// Responses also include a "next change ID" which is used to request the
	// next set of stashes in chronological order (by publish time).
	GetStashes(GetStashOptions) (StashResponse, error)

	// GetLatestStashID retrieves the latest stash tab ID from poe.ninja. This
	// is helpful when building real-time trade applications, as not specifying
	// a stash ID starts from the beginning of time. This makes a single request
	// to poe.ninja's API, and caches the response to avoid subsequent traffic.
	GetLatestStashID() (string, error)
}

type client struct {
	httpClient *http.Client

	host      string
	ninjaHost string

	useSSL      bool
	useCache    bool
	useDNSCache bool

	limiter  *ratelimiter
	cache    *cache
	dnscache *dnscache
}

// NewAPIClient configures and returns an APIClient.
func NewAPIClient(opts ClientOptions) (APIClient, error) {
	if err := validateClientOptions(opts); err != nil {
		return nil, err
	}

	c := &client{
		host:        opts.Host,
		ninjaHost:   opts.NinjaHost,
		useSSL:      opts.UseSSL,
		useCache:    opts.UseCache,
		useDNSCache: opts.UseDNSCache,
		limiter:     newRateLimiter(opts.RateLimit, opts.StashRateLimit),
	}

	if opts.UseCache {
		cache, err := newCache(opts.CacheSize)
		if err != nil {
			return nil, err
		}
		c.cache = cache
	}

	if opts.UseDNSCache {
		c.dnscache = newDNSCache()
		c.httpClient = &http.Client{
			Transport: &http.Transport{
				// When a connection dials an address for the first time, if the
				// host is DefaultHost, resolve the IP using the local DNS
				// cache.
				Dial: func(proto, addr string) (net.Conn, error) {
					a := strings.Split(addr, ":")
					ip, err := c.dnscache.Get(a[0])
					if err != nil {
						return nil, err
					}
					newAddr := fmt.Sprintf("%s:%s", ip, a[1])
					return net.Dial(proto, newAddr)
				},
			},
			Timeout: opts.RequestTimeout,
		}
	} else {
		c.httpClient = &http.Client{Timeout: opts.RequestTimeout}
	}

	return c, nil
}

// ClientOptions contains settings for client initialization.
type ClientOptions struct {
	// The hostname used by the client.
	Host string

	// The hostname used for poe.ninja requests.
	NinjaHost string

	// Set to false if your network does not allow outbound HTTPS traffic.
	UseSSL bool

	// Set to false to always hit the API for requests instead of using a local
	// cache. Caching is always disabled for stash requests, since retrieving
	// a cached stash means we will never get a new change ID.
	UseCache bool

	// The number of items which can be stored in the cache. Most endpoints
	// have a response size up to 500KB or so.
	CacheSize int

	// Set to true to cache DNS resolution locally, speeding up subsequent
	// requests. Go's resolver does not cache by default.
	UseDNSCache bool

	// The number of requests per second for all API endpoints except the stash
	// tab endpoint. The API will ratelimit clients above 5rps.
	RateLimit float64

	// The number of requests per second for the stash endpoint. The API
	// will ratelimit clients above 1rps.
	StashRateLimit float64

	// Time to wait before canceling HTTP requests.
	RequestTimeout time.Duration
}

// DefaultClientOptions initializes the client with the most common settings.
var DefaultClientOptions = ClientOptions{
	Host:           DefaultHost,
	NinjaHost:      DefaultNinjaHost,
	UseSSL:         true,
	UseCache:       true,
	CacheSize:      DefaultCacheSize,
	UseDNSCache:    true,
	RateLimit:      DefaultRateLimit,
	StashRateLimit: DefaultStashRateLimit,
	RequestTimeout: DefaultRequestTimeout,
}

func validateClientOptions(opts ClientOptions) error {
	if _, ok := validHosts[opts.Host]; !ok {
		return ErrInvalidHost
	}
	if opts.NinjaHost == "" {
		return ErrInvalidNinjaHost
	}
	if opts.UseCache && opts.CacheSize < 1 {
		return ErrInvalidCacheSize
	}
	if opts.RateLimit < 0 {
		return ErrInvalidRateLimit
	}
	if opts.StashRateLimit < 0 {
		return ErrInvalidStashRateLimit
	}
	if opts.RequestTimeout < 1*time.Millisecond {
		return ErrInvalidRequestTimeout
	}
	return nil
}
