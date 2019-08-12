package poeapi

import (
	"testing"
)

func TestStashOptionsToQueryParams(t *testing.T) {
	opts := GetStashOptions{ID: "1234"}
	params := opts.toQueryParams()
	expected := "id=1234"
	if params != expected {
		t.Fatalf("failed to convert stash options to query params. expected %s, got %s",
			expected, params)
	}
}

func TestGetStash(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	_, err := c.GetStashes(GetStashOptions{})
	if err != nil {
		t.Fatalf("failed to get stashes: %v", err)
	}
}

func TestParseStashResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/stash.json")
	if err != nil {
		t.Fatalf("failed to load fixture: %v", err)
	}
	if _, err = parseStashResponse(resp); err != nil {
		t.Fatalf("failed to parse stash response: %v", err)
	}
}

func TestGetLatestStashID(t *testing.T) {
	c := client{
		ninjaHost:  testHost,
		useSSL:     false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}
	if _, err := c.GetLatestStashID(); err != nil {
		t.Fatalf("failed to get latest change id: %v", err)
	}
}

func TestGetLatestStashIDWithSSL(t *testing.T) {
	c := client{
		ninjaHost:  testHost,
		useSSL:     false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}
	if _, err := c.GetLatestStashID(); err != nil {
		t.Fatalf("failed to get latest change id: %v", err)
	}
}

func TestParseLatestChangeResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/latest-change.json")
	if err != nil {
		t.Fatalf("failed to load fixture: %v", err)
	}
	if _, err := parseLatestChangeResponse(resp); err != nil {
		t.Fatalf("failed to parse latest change: %v", err)
	}
}
