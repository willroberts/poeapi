package poeapi

import (
	"encoding/json"
	"errors"
	"time"
)

// APIClient provides methods for interacting with the Path of Exile API.
type APIClient interface {
	GetAllLeagues() ([]League, error)
	GetCurrentChallengeLeague() (League, error)
}

// ClientOptions contains settings for client initialization.
// TODO: Validate options somewhere.
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

// NewAPIClient configures and returns an APIClient.
func NewAPIClient(opts ClientOptions) APIClient {
	return &client{
		host:    opts.Host,
		useSSL:  opts.UseSSL,
		limiter: newJankLimiter(opts.RateLimit, opts.StashTabRateLimit),
	}
}

type client struct {
	host   string
	useSSL bool

	limiter *janklimiter
}

func (c *client) GetAllLeagues() ([]League, error) {
	leagues := make([]League, 0)

	resp, err := c.getJSON(c.formatURL(leaguesEndpoint))
	if err != nil {
		return leagues, err
	}

	if err := json.Unmarshal([]byte(resp), &leagues); err != nil {
		return leagues, err
	}

	return leagues, nil
}

func (c *client) GetCurrentChallengeLeague() (League, error) {
	leagues, err := c.GetAllLeagues()
	if err != nil {
		return League{}, err
	}

	// The challenge league is generally the fifth entry in the slice, after
	// Standard, Hardcore, SSF Standard, and SSF Hardcore.
	// It is the first entry with a non-nil EndTime value.
	for _, l := range leagues {
		if (l.EndTime != time.Time{}) {
			return l, nil
		}
	}
	return League{}, errors.New("failed to find challenge league")
}
