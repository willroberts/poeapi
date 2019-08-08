package poeapi

import "testing"

func TestFormatURL(t *testing.T) {
	var (
		expected = "https://api.pathofexile.com/leagues"
		client   = client{
			host:   "api.pathofexile.com",
			useSSL: true,
		}
	)
	if client.formatURL(leaguesEndpoint) != expected {
		t.Fatalf("failed to format url: expected %s, got %s",
			expected, client.formatURL(leaguesEndpoint))
	}
}

func TestFormatHTTPURL(t *testing.T) {
	var (
		expected = "http://api.pathofexile.com/ladders"
		client   = client{
			host:   "api.pathofexile.com",
			useSSL: false,
		}
	)
	if client.formatURL(laddersEndpoint) != expected {
		t.Fail()
		t.Fatalf("failed to format url: expected %s, got %s",
			expected, client.formatURL(laddersEndpoint))
	}
}
