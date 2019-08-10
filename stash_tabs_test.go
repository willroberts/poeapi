package poeapi

import "testing"

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
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to creat eclient for stash test: %v", err)
	}
	_, err = c.GetStashes(GetStashOptions{})
	if err != nil {
		t.Fatalf("failed to get stashes: %v", err)
	}
}

func TestGetLatestChangeID(t *testing.T) {
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for latest change id test: %v", err)
	}
	_, err = c.GetLatestChangeID()
	if err != nil {
		t.Fatalf("failed to get latest change id: %v", err)
	}
}
