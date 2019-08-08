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
		t.Fail()
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
	}
}
