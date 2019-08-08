package poeapi

import (
	"encoding/json"
	"errors"
	"time"
)

func (c *client) GetAllLeagues() ([]League, error) {
	resp, err := c.getJSON(c.formatURL(leaguesEndpoint))
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
