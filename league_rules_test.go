package poeapi

import "testing"

func TestGetLeagueRules(t *testing.T) {
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for league rules test: %v", err)
	}

	_, err = c.GetLeagueRules()
	if err != nil {
		t.Fatalf("failed to get league rules: %v", err)
	}
}

func TestGetLeagueRulesRequestFailure(t *testing.T) {
	var (
		c = client{
			host:    "google.com",
			limiter: newRateLimiter(DefaultRateLimit, DefaultStashRateLimit),
		}
	)
	_, err := c.GetLeagueRules()
	if err != ErrNotFound {
		t.Fatal("failed to detect request error for league rules request")
	}
}

func TestParseLeagueRulesResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/league-rules.json")
	if err != nil {
		t.Fatalf("failed to read fixture for league rules test: %v", err)
	}

	_, err = parseLeagueRulesResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse league rules response: %v", err)
	}
}

func TestParseLeagueRulesResponseWithInvalidJSON(t *testing.T) {
	resp, err := loadFixture("fixtures/invalid.json")
	if err != nil {
		t.Fatalf("failed to read fixture for league rules test: %v", err)
	}

	_, err = parseLeagueRulesResponse(resp)
	if err == nil {
		t.Fatal("failed to detect invalid league rules json")
	}
}
