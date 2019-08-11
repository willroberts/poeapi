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

func TestGetLatestStashID(t *testing.T) {
	c := client{
		host:       testHost,
		ninjaHost:  testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	_, err := c.GetLatestStashID()
	if err != nil {
		t.Fatalf("failed to get latest change id: %v", err)
	}
}
