package poeapi

import "testing"

func TestGetAllLeagues(t *testing.T) {
	client, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for leagues request: %v", err)
	}
	_, err = client.GetAllLeagues()
	if err != nil {
		t.Fatalf("failed to get all leagues: %v", err)
	}
}

func TestParseLeaguesResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/leagues.json")
	if err != nil {
		t.Fatalf("failed to read fixture for pvp matches test: %v", err)
	}

	_, err = parseLeaguesResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse leagues response: %v", err)
	}
}

func TestParseLeaguesResponseWithInvalidJSON(t *testing.T) {
	resp := "{\"invalid_json\":true}"
	_, err := parseLeaguesResponse(resp)
	if err == nil {
		t.Fatal("failed to detect error in leagues response parsing")
	}
}

func TestGetAllLeaguesWithRequestFailure(t *testing.T) {
	var (
		rateLimit = 1
		client    = client{
			host:    "google.com",
			limiter: newRateLimiter(rateLimit, rateLimit),
		}
	)
	_, err := client.GetAllLeagues()
	if err != ErrNotFound {
		t.Fatal("failed to detect request error for leagues request")
	}
}

func TestGetCurrentChallengeLeague(t *testing.T) {
	client, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for challenge league request: %v", err)
	}
	_, err = client.GetCurrentChallengeLeague()
	if err != nil {
		t.Fatalf("failed to get current challenge league: %v", err)
	}
}

func TestGetCurrentChallengeLeagueWithGetFailure(t *testing.T) {
	var (
		rateLimit = 1
		client    = client{
			host:    "google.com",
			limiter: newRateLimiter(rateLimit, rateLimit),
		}
	)
	_, err := client.GetCurrentChallengeLeague()
	if err != ErrNotFound {
		t.Fatal("failed to detect request error for challenge leagues request")
	}
}
