package poeapi

import (
	"log"
	"testing"
)

func TestGetPVPMatches(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	_, err := c.GetPVPMatches(GetPVPMatchesOptions{
		Type:   "season",
		Season: "EUPvPSeason1",
		Realm:  "pc",
	})
	if err != nil {
		t.Fatalf("failed to get pvp matches: %v", err)
	}
}

func TestValidateGetPVPMatchesOptions(t *testing.T) {
	opts := GetPVPMatchesOptions{
		Type:   "season",
		Season: "EUPvPSeason1",
		Realm:  "pc",
	}
	if err := validateGetPVPMatchesOptions(opts); err != nil {
		log.Fatalf("failed to validate pvp options: %v", err)
	}
}

func TestValidateGetPVPMatchesOptionsWithInvalidSeason(t *testing.T) {
	opts := GetPVPMatchesOptions{
		Type:   "season",
		Season: "",
		Realm:  "pc",
	}
	if err := validateGetPVPMatchesOptions(opts); err != ErrInvalidSeason {
		log.Fatal("failed to detect invalid pvp season")
	}
}

func TestValidateGetPVPMatchesOptionsWithInvalidRealm(t *testing.T) {
	opts := GetPVPMatchesOptions{
		Type:   "season",
		Season: "EUPvPSeason1",
		Realm:  "toaster",
	}
	if err := validateGetPVPMatchesOptions(opts); err != ErrInvalidRealm {
		log.Fatal("failed to detect invalid pvp realm")
	}
}

func TestGetPVPMatchesWithInvalidOptions(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	_, err := c.GetPVPMatches(GetPVPMatchesOptions{
		Type:   "season",
		Season: "",
		Realm:  "pc",
	})
	if err != ErrInvalidSeason {
		t.Fatal("failed to detect invalid season in pvp matches test")
	}
}

func TestParsePVPMatchesResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/pvp-matches.json")
	if err != nil {
		t.Fatalf("failed to read fixture for pvp matches test: %v", err)
	}

	_, err = parsePVPMatchesResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse pvp matches response: %v", err)
	}
}

func TestParsePVPMatchesResponseWithInvalidJSON(t *testing.T) {
	resp, err := loadFixture("fixtures/invalid.json")
	if err != nil {
		t.Fatalf("failed to read fixture for pvp matches test: %v", err)
	}

	_, err = parsePVPMatchesResponse(resp)
	if err == nil {
		t.Fatal("failed to detect invalid pvp matches json")
	}
}
