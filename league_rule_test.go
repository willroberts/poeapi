package poeapi

import "testing"

func TestGetLeagueRule(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	opts := GetLeagueRuleOptions{ID: "TurboMonsters"}
	_, err := c.GetLeagueRule(opts)
	if err != nil {
		t.Fatalf("failed to get league rule: %v", err)
	}
}

func TestGetLeagueRuleWithInvalidOptions(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	opts := GetLeagueRuleOptions{ID: ""}
	_, err := c.GetLeagueRule(opts)
	if err != ErrInvalidLeagueRuleID {
		t.Fatal("failed to detect invalid league rule option in league rule request")
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

func TestParseLeagueRuleResponseWithInvalidJSON(t *testing.T) {
	resp, err := loadFixture("fixtures/invalid.json")
	if err != nil {
		t.Fatalf("failed to read fixture for league rule test: %v", err)
	}

	_, err = parseLeagueRuleResponse(resp)
	if err == nil {
		t.Fatal("failed to detect invalid league rule json")
	}
}

func TestValidateLeagueRuleOptions(t *testing.T) {
	opts := GetLeagueRuleOptions{ID: "test"}
	if err := validateLeagueRuleOptions(opts); err != nil {
		t.Fatalf("failed to validate league rule options: %v", err)
	}
}

func TestValidateLeagueRuleOptionsWithInvalidID(t *testing.T) {
	opts := GetLeagueRuleOptions{ID: ""}
	if err := validateLeagueRuleOptions(opts); err != ErrInvalidLeagueRuleID {
		t.Fatal("failed to detect invalid league rule option")
	}
}
