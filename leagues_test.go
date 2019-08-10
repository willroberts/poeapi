package poeapi

import "testing"

func TestLeaguesOptionsToQueryParams(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "season",
		Realm:   "pc",
		Season:  "testseason",
		Compact: true,
		Limit:   5,
		Offset:  5,
	}
	params := opts.toQueryParams()
	expected := "compact=1&limit=5&offset=5&realm=pc&season=testseason&type=season"
	if params != expected {
		t.Fatalf("failed to encode leagues query params: expected %s, got %s",
			expected, params)
	}
}

func TestValidateGetLeaguesOptions(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "main",
		Realm:   "pc",
		Compact: false,
		Limit:   0,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != nil {
		t.Fatalf("failed to validate leagues options: %v", err)
	}
}

func TestValidateGetLeaguesOptionsWithInvalidLeagueType(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "invalid",
		Realm:   "pc",
		Compact: false,
		Limit:   0,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidLeagueType {
		t.Fatal("failed to detect invalid league type in leagues options")
	}
}

func TestValidateGetLeaguesOptionsWithInvalidRealm(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "main",
		Realm:   "invalid",
		Compact: false,
		Limit:   0,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidRealm {
		t.Fatal("failed to detect invalid realm in leagues options")
	}
}

func TestValidateGetLeaguesOptionsWithInvalidSeason(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "season",
		Realm:   "pc",
		Season:  "",
		Compact: false,
		Limit:   0,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidSeason {
		t.Fatal("failed to detect invalid season in leagues options")
	}
}

func TestValidateGetLeaguesOptionsWithInvalidLimit(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "main",
		Realm:   "pc",
		Compact: false,
		Limit:   -1,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidLimit {
		t.Fatal("failed to detect invalid limit in leagues options")
	}
}

func TestValidateGetLeaguesOptionsWithInvalidCompactLimit(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "main",
		Realm:   "pc",
		Compact: true,
		Limit:   231,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidLimit {
		t.Fatal("failed to detect invalid limit in leagues options with compact=1")
	}
}

func TestValidateGetLeaguesOptionsWithInvalidNonCompactLimit(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "main",
		Realm:   "pc",
		Compact: false,
		Limit:   51,
		Offset:  0,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidLimit {
		t.Fatal("failed to detect invalid limit in leagues options with compact=0")
	}
}

func TestValidateGetLeaguesOptionsWithInvalidOffset(t *testing.T) {
	opts := GetLeaguesOptions{
		Type:    "main",
		Realm:   "pc",
		Compact: false,
		Limit:   0,
		Offset:  -1,
	}
	if err := validateGetLeaguesOptions(opts); err != ErrInvalidOffset {
		t.Fatal("failed to detect invalid offset in leagues options")
	}
}

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

func TestGetLeaguesWithInvalidOptions(t *testing.T) {
	client, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for leagues request: %v", err)
	}
	_, err = client.GetLeagues(GetLeaguesOptions{Realm: "toaster"})
	if err != ErrInvalidRealm {
		t.Fatal("failed to detect invalid options in leagues request")
	}
}

func TestGetLeaguesRequestFailure(t *testing.T) {
	var (
		c = client{
			host:    "google.com",
			limiter: newRateLimiter(DefaultRateLimit, DefaultStashTabRateLimit),
		}
	)
	_, err := c.GetLeagues(GetLeaguesOptions{})
	if err != ErrNotFound {
		t.Fatal("failed to detect request error for leagues request")
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
