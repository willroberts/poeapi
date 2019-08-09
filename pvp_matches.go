package poeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetPVPMatchesOptions contains the request parameters for the PVP matches
// endpoint. All parameters are optional.
type GetPVPMatchesOptions struct {
	Type   string
	Season string
	Realm  string
}

// ToQueryParams converts PVP matches options to a URL parameter string.
func (opts GetPVPMatchesOptions) ToQueryParams() string {
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

// GetPVPMatches retrieves all PVP matches from the API.
func (c *client) GetPVPMatches(opts GetPVPMatchesOptions) ([]PVPMatch, error) {
	url := fmt.Sprintf("%s?%s", c.formatURL(pvpMatchesEndpoint),
		opts.ToQueryParams())
	fmt.Println(url)
	resp, err := c.get(url)
	if err != nil {
		return []PVPMatch{}, err
	}
	return parsePVPMatchesResponse(resp)
}

// parsePVPMatchesResponse unmarshals JSON from the API into local types.
func parsePVPMatchesResponse(resp string) ([]PVPMatch, error) {
	fmt.Println(resp)
	pvpMatches := make([]PVPMatch, 0)
	if err := json.Unmarshal([]byte(resp), &pvpMatches); err != nil {
		return []PVPMatch{}, err
	}
	return pvpMatches, nil
}
