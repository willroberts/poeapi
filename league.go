package poeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetLeagueOptions contains the request parameters for the league endpoint.
// ID is a required option, and Realm is optional. The API allows callers to
// request the ladder alongside the league, but this is not implemented. Use
// the GetLadder() method instead.
type GetLeagueOptions struct {
	// The name of the league to retrieve.
	ID string

	// The realm of the ladder.
	// Valid options: 'pc', 'xbox', or 'sony'.
	Realm string
}

func (opts GetLeagueOptions) toQueryParams() string {
	u := url.Values{}
	if opts.Realm != "" {
		u.Add("realm", opts.Realm)
	}
	return u.Encode()
}

func validateGetLeagueOptions(opts GetLeagueOptions) error {
	if opts.ID == "" {
		return ErrInvalidLeagueID
	}
	if opts.Realm != "" {
		if _, ok := validRealms[opts.Realm]; !ok {
			return ErrInvalidRealm
		}
	}
	return nil
}

func (c *client) GetLeague(opts GetLeagueOptions) (League, error) {
	if err := validateGetLeagueOptions(opts); err != nil {
		return League{}, err
	}
	resp, err := c.get(fmt.Sprintf("%s/%s", c.formatURL(leaguesEndpoint), opts.ID))
	if err != nil {
		return League{}, err
	}
	return parseLeagueResponse(resp)
}

func parseLeagueResponse(resp string) (League, error) {
	league := League{}
	if err := json.Unmarshal([]byte(resp), &league); err != nil {
		return League{}, err
	}
	return league, nil
}
