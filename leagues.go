package poeapi

import (
	"encoding/json"
	"errors"
	"time"
)

// GetLeagues retrieves all leagues (Standard, Hardcore, etc.) from the API.
func (c *client) GetLeagues() ([]League, error) {
	resp, err := c.get(c.formatURL(leaguesEndpoint))
	if err != nil {
		return []League{}, err
	}
	return parseLeaguesResponse(resp)
}

// parseLeaguesResponse unmarshals JSON from the API into local types.
func parseLeaguesResponse(resp string) ([]League, error) {
	leagues := make([]League, 0)
	if err := json.Unmarshal([]byte(resp), &leagues); err != nil {
		return []League{}, err
	}
	return leagues, nil
}

// GetCurrentChallengeLeague retrieves all leagues and returns the first league
// with a time limit, which is generally the current challenge league.
func (c *client) GetCurrentChallengeLeague() (League, error) {
	leagues, err := c.GetLeagues()
	if err != nil {
		return League{}, err
	}

	for _, l := range leagues {
		if (l.EndTime != time.Time{}) {
			return l, nil
		}
	}
	return League{}, errors.New("failed to find challenge league")
}
