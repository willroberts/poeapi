package poeapi

import (
	"encoding/json"
	"fmt"
)

// GetLeagueRuleOptions contains the request parameters for the league rules
// endpoint. The only parameter, ID, is required.
type GetLeagueRuleOptions struct {
	// The identifier of the league rule to retrieve.
	ID string
}

func validateLeagueRuleOptions(opts GetLeagueRuleOptions) error {
	if opts.ID == "" {
		return ErrInvalidLeagueRuleID
	}
	return nil
}

// GetLeagueRule retrieves a single league rule from the API.
func (c *client) GetLeagueRule(opts GetLeagueRuleOptions) (LeagueRule, error) {
	if err := validateLeagueRuleOptions(opts); err != nil {
		return LeagueRule{}, err
	}

	url := fmt.Sprintf("%s/%s", c.formatURL(leagueRulesEndpoint), opts.ID)
	resp, err := c.get(url)
	if err != nil {
		return LeagueRule{}, err
	}
	return parseLeagueRuleResponse(resp)
}

// parseLeagueRulesResponse unmarshals JSON from the API into local types.
func parseLeagueRuleResponse(resp string) (LeagueRule, error) {
	rule := LeagueRule{}
	if err := json.Unmarshal([]byte(resp), &rule); err != nil {
		return LeagueRule{}, err
	}
	return rule, nil
}
