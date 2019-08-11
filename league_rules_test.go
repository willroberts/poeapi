package poeapi

import "testing"

func TestGetLeagueRules(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	_, err := c.GetLeagueRules()
	if err != nil {
		t.Fatalf("failed to get league rules: %v", err)
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
