package poeapi

import "testing"

func TestGetAllLeagues(t *testing.T) {
	client, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fatalf("failed to create client for leagues request: %v", err)
	}
	_, err = client.GetAllLeagues()
	if err != nil {
		t.Fatalf("failed to get all leagues: %v", err)
	}
}

func TestParseLeaguesResponse(t *testing.T) {
	_, err := parseLeaguesResponse(leaguesResponseFixture)
	if err != nil {
		t.Fatalf("failed to parse leagues response: %v", err)
	}
}

func TestParseLeaguesResponseWithInvalidJSON(t *testing.T) {
	resp := "{\"invalid_json\":true}"
	_, err := parseLeaguesResponse(resp)
	if err == nil {
		t.Fatal("failed to detect error in leagues response parsing")
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
		t.Fatal("failed to detect request error for leagues request")
	}
}

func TestGetCurrentChallengeLeague(t *testing.T) {
	client, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fatalf("failed to create client for challenge league request: %v", err)
	}
	_, err = client.GetCurrentChallengeLeague()
	if err != nil {
		t.Fatalf("failed to get current challenge league: %v", err)
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
		t.Fatal("failed to detect request error for challenge leagues request")
	}
}

const (
	leaguesResponseFixture = `[{"id":"Standard","realm":"pc","description":"The default game mode.","url":"http:\/\/pathofexile.com\/forum\/view-thread\/71278","startAt":"2013-01-23T21:00:00Z","endAt":null,"delveEvent":true,"rules":[]},{"id":"Hardcore","realm":"pc","description":"A character killed in the Hardcore league is moved to the Standard league.","url":"http:\/\/pathofexile.com\/forum\/view-thread\/71276","startAt":"2013-01-23T21:00:00Z","endAt":null,"delveEvent":true,"rules":[{"id":"Hardcore","name":"Hardcore","description":"A character killed in Hardcore is moved to its parent league."}]},{"id":"SSF Standard","realm":"pc","description":"SSF Standard","url":"http:\/\/pathofexile.com\/forum\/view-thread\/1841357","startAt":"2013-01-23T21:00:00Z","endAt":null,"delveEvent":true,"rules":[{"id":"NoParties","name":"Solo","description":"You may not party in this league."}]},{"id":"SSF Hardcore","realm":"pc","description":"SSF Hardcore","url":"http:\/\/pathofexile.com\/forum\/view-thread\/1841353","startAt":"2013-01-23T21:00:00Z","endAt":null,"delveEvent":true,"rules":[{"id":"Hardcore","name":"Hardcore","description":"A character killed in Hardcore is moved to its parent league."},{"id":"NoParties","name":"Solo","description":"You may not party in this league."}]},{"id":"Legion","realm":"pc","description":"Journey to the Domain of Timeless Conflict to defeat historic Legions and their generals.\n\nThis is the default Path of Exile league.","registerAt":"2019-06-07T18:00:00Z","url":"http:\/\/pathofexile.com\/forum\/view-thread\/2515509","startAt":"2019-06-07T20:00:00Z","endAt":"2019-09-02T22:00:00Z","delveEvent":true,"rules":[]},{"id":"Hardcore Legion","realm":"pc","description":"Journey to the Domain of Timeless Conflict to defeat historic Legions and their generals.\n\nA character killed in Hardcore Legion becomes a Standard character.","registerAt":"2019-06-07T18:00:00Z","url":"http:\/\/pathofexile.com\/forum\/view-thread\/2515511","startAt":"2019-06-07T20:00:00Z","endAt":"2019-09-02T22:00:00Z","delveEvent":true,"rules":[{"id":"Hardcore","name":"Hardcore","description":"A character killed in Hardcore is moved to its parent league."}]},{"id":"SSF Legion","realm":"pc","description":"SSF Legion","registerAt":"2019-06-07T18:00:00Z","url":"http:\/\/pathofexile.com\/forum\/view-thread\/2515512","startAt":"2019-06-07T20:00:00Z","endAt":"2019-09-02T22:00:00Z","delveEvent":true,"rules":[{"id":"NoParties","name":"Solo","description":"You may not party in this league."}]},{"id":"SSF Legion HC","realm":"pc","description":"SSF HC Legion","registerAt":"2019-06-07T18:00:00Z","url":"http:\/\/pathofexile.com\/forum\/view-thread\/2515514","startAt":"2019-06-07T20:00:00Z","endAt":"2019-09-02T22:00:00Z","delveEvent":true,"rules":[{"id":"Hardcore","name":"Hardcore","description":"A character killed in Hardcore is moved to its parent league."},{"id":"NoParties","name":"Solo","description":"You may not party in this league."}]}]`
)
