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

func TestGetLeagueWithInvalidOptions(t *testing.T) {
	client, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for league request: %v", err)
	}
	_, err = client.GetLeague(GetLeagueOptions{ID: ""})
	if err == nil {
		t.Fatal("failed to detect invalid options in league request")
	}
}

func TestGetLeagueWithRequestFailure(t *testing.T) {
	var (
		c = client{
			host:    "google.com",
			limiter: newRateLimiter(DefaultRateLimit, DefaultStashTabRateLimit),
		}
	)
	_, err := c.GetLeague(GetLeagueOptions{ID: "Standard"})
	if err != ErrNotFound {
		t.Log("err:", err)
		t.Fatal("failed to detect request error for league request")
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

func TestLeagueOptionsToQueryParams(t *testing.T) {
	opts := GetLeagueOptions{Realm: "test"}
	params := opts.toQueryParams()
	expected := "realm=test"
	if params != expected {
		t.Fatalf("failed to get query params from league options. expected %s, got %s",
			expected, params)
	}
}

func TestValidateGetLeagueOptions(t *testing.T) {
	opts := GetLeagueOptions{
		ID:    "test",
		Realm: "pc",
	}
	if err := validateGetLeagueOptions(opts); err != nil {
		t.Fatalf("failed to validate league options: %v", err)
	}
}

func TestValidateGetLeagueOptionsWithInvalidID(t *testing.T) {
	opts := GetLeagueOptions{
		ID:    "",
		Realm: "pc",
	}
	if err := validateGetLeagueOptions(opts); err != ErrInvalidLeagueID {
		t.Fatal("failed to detect invalid league id")
	}
}

func TestValidateGetLeagueOptionsWithInvalidRealm(t *testing.T) {
	opts := GetLeagueOptions{
		ID:    "test",
		Realm: "toaster",
	}
	if err := validateGetLeagueOptions(opts); err != ErrInvalidRealm {
		t.Fatalf("failed to detect invalid realm")
	}
}
