package poeapi

import "testing"

func TestGetLeagueRule(t *testing.T) {
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for league rule test: %v", err)
	}

	opts := GetLeagueRuleOptions{
		ID: "TurboMonsters",
	}

	_, err = c.GetLeagueRule(opts)
	if err != nil {
		t.Fatalf("failed to get league rule: %v", err)
	}
}

func TestParseLeagueRuleResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/league-rule.json")
	if err != nil {
		t.Fatalf("failed to read fixture for league rule test: %v", err)
	}

	_, err = parseLeagueRuleResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse league rule response: %v", err)
	}
}
