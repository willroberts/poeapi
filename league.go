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
	ID    string
	Realm string // pc, xbox, or sony.
}

// ToQueryParams converts options to a URL query string.
func (opts GetLeagueOptions) ToQueryParams() string {
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

// GetLeague retrieves all league (Standard, Hardcore, etc.) from the API.
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

// parseLeagueResponse unmarshals JSON from the API into local types.
func parseLeagueResponse(resp string) (League, error) {
	league := League{}
	if err := json.Unmarshal([]byte(resp), &league); err != nil {
		return League{}, err
	}
	return league, nil
}
