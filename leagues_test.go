package poeapi

import "testing"

func TestGetAllLeagues(t *testing.T) {
	client, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fail()
	}
	_, err = client.GetAllLeagues()
	if err != nil {
		t.Fail()
	}
}

func TestParseLeaguesResponse(t *testing.T) {
	t.Skip() // Coverage already provided in TestGetAllLeagues().
}

func TestParseLeaguesResponseWithInvalidJSON(t *testing.T) {
	resp := "{\"invalid_json\":true}"
	_, err := parseLeaguesResponse(resp)
	if err == nil {
		t.Fail()
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
		t.Fail()
	}
}

func TestGetCurrentChallengeLeague(t *testing.T) {
	client, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fail()
	}
	_, err = client.GetCurrentChallengeLeague()
	if err != nil {
		t.Fail()
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
		t.Fail()
	}
}
