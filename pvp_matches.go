package poeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const seasonMatchType = "season"

// GetPVPMatchesOptions contains the request parameters for the PVP matches
// endpoint. All parameters are optional.
type GetPVPMatchesOptions struct {
	// The type of matches to find. Set to 'season' to search for a specific PVP
	// season, or leave blank to return all upcoming PVP matches.
	Type string

	// The name of the season for which to retrieve matches.
	Season string

	// The realm of PVP matches to retrieve.
	// Valid options: 'pc', 'xbox', or 'sony'.
	Realm string
}

// toQueryParams converts PVP matches options to a URL parameter string.
func (opts GetPVPMatchesOptions) toQueryParams() string {
	u := url.Values{}
	if opts.Type != "" {
		u.Add("type", opts.Type)
	}
	if opts.Season != "" {
		u.Add("season", opts.Season)
	}
	if opts.Realm != "" {
		u.Add("realm", opts.Realm)
	}
	return u.Encode()
}

func validateGetPVPMatchesOptions(opts GetPVPMatchesOptions) error {
	if opts.Type == seasonMatchType && opts.Season == "" {
		return ErrInvalidSeason
	}
	if opts.Realm != "" {
		if _, ok := validRealms[opts.Realm]; !ok {
			return ErrInvalidRealm
		}
	}
	return nil
}

// GetPVPMatches retrieves all PVP matches from the API.
func (c *client) GetPVPMatches(opts GetPVPMatchesOptions) ([]PVPMatch, error) {
	if err := validateGetPVPMatchesOptions(opts); err != nil {
		return []PVPMatch{}, err
	}
	url := fmt.Sprintf("%s?%s", c.formatURL(pvpMatchesEndpoint),
		opts.toQueryParams())
	resp, err := c.get(url)
	if err != nil {
		return []PVPMatch{}, err
	}
	return parsePVPMatchesResponse(resp)
}

// parsePVPMatchesResponse unmarshals JSON from the API into local types.
func parsePVPMatchesResponse(resp string) ([]PVPMatch, error) {
	pvpMatches := make([]PVPMatch, 0)
	if err := json.Unmarshal([]byte(resp), &pvpMatches); err != nil {
		return []PVPMatch{}, err
	}
	return pvpMatches, nil
}
