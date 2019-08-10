package poeapi

import "testing"

func TestGetLeague(t *testing.T) {
	client, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for league request: %v", err)
	}
	_, err = client.GetLeague(GetLeagueOptions{ID: "Standard"})
	if err != nil {
		t.Fatalf("failed to get all league: %v", err)
	}
}

func TestParseLeagueResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/league.json")
	if err != nil {
		t.Fatalf("failed to read fixture for league test: %v", err)
	}

	_, err = parseLeagueResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse league response: %v", err)
	}
}

func TestParseLeagueResponseWithInvalidJSON(t *testing.T) {
	resp, err := loadFixture("fixtures/invalid.json")
	if err != nil {
		t.Fatalf("failed to load fixture for league response parsing: %v", err)
	}

	_, err = parseLeagueResponse(resp)
	if err == nil {
		t.Fatal("failed to detect error in league response parsing")
	}
}

func TestGetLeagueWithRequestFailure(t *testing.T) {
	var (
		rateLimit = 1
		client    = client{
			host:    "google.com",
			limiter: newRateLimiter(rateLimit, rateLimit),
		}
	)
	_, err := client.GetLeague(GetLeagueOptions{ID: "Standard"})
	if err != ErrNotFound {
		t.Log("err:", err)
		t.Fatal("failed to detect request error for league request")
	}
}
