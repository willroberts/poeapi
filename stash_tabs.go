package poeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// GetStashOptions ...
type GetStashOptions struct {
	// ID is the unique change ID containing a set of stashes. If ID is omitted,
	// the API will return the oldest stash tab possible.
	ID string
}

func (opts GetStashOptions) toQueryParams() string {
	u := url.Values{}
	if opts.ID != "" {
		u.Add("id", opts.ID)
	}
	return u.Encode()
}

// GetStashes ...
func (c *client) GetStashes(opts GetStashOptions) (StashResponse, error) {
	url := fmt.Sprintf("%s?%s", c.formatURL(stashTabsEndpoint),
		opts.toQueryParams())
	resp, err := c.get(url)
	if err != nil {
		return StashResponse{}, err
	}

	var s StashResponse
	if err := json.Unmarshal([]byte(resp), &s); err != nil {
		return StashResponse{}, err
	}

	return s, nil
}

type latestChange struct {
	ID string `json:"next_change_id"`
}

// GetLatestStashID retrieves the latest stash tab ID from poe.ninja, with
// caching and ratelimiting.
func (c *client) GetLatestStashID() (string, error) {
	// poe.ninja only supports HTTPS, but we allow HTTP for local testing.
	var url string
	if strings.HasPrefix(c.ninjaHost, "127.0.0.1") {
		url = fmt.Sprintf("http://%s/api/Data/GetStats", c.ninjaHost)
	} else {
		url = fmt.Sprintf("https://%s/api/Data/GetStats", c.ninjaHost)
	}

	resp, err := c.get(url)
	if err != nil {
		return "", err
	}

	fmt.Println(resp)

	var latest latestChange
	if err := json.Unmarshal([]byte(resp), &latest); err != nil {
		return "", err
	}

	return latest.ID, nil
}
