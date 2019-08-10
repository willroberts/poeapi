package poeapi

import "testing"

func TestGetPVPMatches(t *testing.T) {
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for pvp matches test: %v", err)
	}

	opts := GetPVPMatchesOptions{
		Type:   "season",
		Season: "EUPvPSeason1",
	}
	_, err = c.GetPVPMatches(opts)
	if err != nil {
		t.Fatalf("failed to get pvp matches: %v", err)
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
