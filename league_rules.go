package poeapi

import (
	"encoding/json"
)

// GetLeagueRules retrieves all league rules from the API.
func (c *client) GetLeagueRules() ([]LeagueRule, error) {
	resp, err := c.get(c.formatURL(leagueRulesEndpoint))
	if err != nil {
		return []LeagueRule{}, err
	}
	return parseLeagueRulesResponse(resp)
}

func parseLeagueRulesResponse(resp string) ([]LeagueRule, error) {
	rules := make([]LeagueRule, 0)
	if err := json.Unmarshal([]byte(resp), &rules); err != nil {
		return []LeagueRule{}, err
	}
	return rules, nil
}
