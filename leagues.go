package poeapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	mainLeagueType   = "main"
	eventLeagueType  = "event"
	seasonLeagueType = "season"
)

// GetLeaguesOptions contains the request parameters for the leagues endpoint.
// All parameters are optional.
type GetLeaguesOptions struct {
	// The type of leagues to retrieve.
	// Valid options: 'main', 'event', or 'season'.
	Type string

	// The realm of leagues to retrieve.
	// Valid options: 'pc', 'xbox', or 'sony'.
	Realm string

	// The name of the season to retrieve. Requires when Type is 'season'.
	Season string

	// Set to true to omit rules, registration time, and description from the
	// response.
	Compact bool

	// Number of leagues to retrieve. Defaults to 50, but can go up to 230 when
	// Compact is true.
	Limit int

	// Starting index for bulk league retrieval. Only needed when requesting
	// more than 50 leagues.
	// TODO: Abstract limit/offset behind client methods.
	Offset int
}

func (opts GetLeaguesOptions) toQueryParams() string {
	u := url.Values{}
	if opts.Type != "" {
		u.Add("type", opts.Type)
	}
	if opts.Realm != "" {
		u.Add("realm", opts.Realm)
	}
	if opts.Type == seasonLeagueType && opts.Season != "" {
		u.Add("season", opts.Season)
	}
	if opts.Compact {
		u.Add("compact", "1")
	}
	if opts.Limit != 0 {
		u.Add("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Offset != 0 {
		u.Add("offset", strconv.Itoa(opts.Offset))
	}
	return u.Encode()
}

func validateGetLeaguesOptions(opts GetLeaguesOptions) error {
	if opts.Type != "" {
		if _, ok := validLeagueTypes[opts.Type]; !ok {
			return ErrInvalidLeagueType
		}
	}
	if opts.Realm != "" {
		if _, ok := validRealms[opts.Realm]; !ok {
			return ErrInvalidRealm
		}
	}
	if opts.Type == seasonLeagueType && opts.Season == "" {
		return ErrInvalidSeason
	}
	if opts.Limit < 0 {
		return ErrInvalidLimit
	}
	if opts.Compact && opts.Limit > 230 {
		return ErrInvalidLimit
	}
	if !opts.Compact && opts.Limit > 50 {
		return ErrInvalidLimit
	}
	if opts.Offset < 0 {
		return ErrInvalidOffset
	}
	return nil
}

// GetLeagues retrieves all leagues (Standard, Hardcore, etc.) from the API.
func (c *client) GetLeagues(opts GetLeaguesOptions) ([]League, error) {
	if err := validateGetLeaguesOptions(opts); err != nil {
		return []League{}, err
	}
	resp, err := c.get(c.formatURL(leaguesEndpoint))
	if err != nil {
		return []League{}, err
	}
	return parseLeaguesResponse(resp)
}

func parseLeaguesResponse(resp string) ([]League, error) {
	leagues := make([]League, 0)
	if err := json.Unmarshal([]byte(resp), &leagues); err != nil {
		return []League{}, err
	}
	return leagues, nil
}
