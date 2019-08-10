package poeapi

import "testing"

func TestGetLeagues(t *testing.T) {
	client, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for leagues request: %v", err)
	}
	_, err = client.GetLeagues(GetLeaguesOptions{})
	if err != nil {
		t.Fatalf("failed to get all leagues: %v", err)
	}
}

func TestParseLeaguesResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/leagues.json")
	if err != nil {
		t.Fatalf("failed to read fixture for leagues test: %v", err)
	}

	_, err = parseLeaguesResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse leagues response: %v", err)
	}
}

func TestParseLeaguesResponseWithInvalidJSON(t *testing.T) {
	resp, err := loadFixture("fixtures/invalid.json")
	if err != nil {
		t.Fatalf("failed to load fixture for leagues response parsing: %v", err)
	}

	_, err = parseLeaguesResponse(resp)
	if err == nil {
		t.Fatal("failed to detect error in leagues response parsing")
	}
}

func TestGetLeaguesWithRequestFailure(t *testing.T) {
	var (
		rateLimit = 1
		client    = client{
			host:    "google.com",
			limiter: newRateLimiter(rateLimit, rateLimit),
		}
	)
	_, err := client.GetLeagues(GetLeaguesOptions{})
	if err != ErrNotFound {
		t.Fatal("failed to detect request error for leagues request")
	}
}
