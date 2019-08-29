package poeapi

import "testing"

func TestLadderOptionstoQueryParams(t *testing.T) {
	var (
		opts = GetLadderOptions{
			Realm:       "sony",
			Type:        "league",
			UniqueIDs:   true,
			AccountName: "testaccount",
			limit:       200,
			offset:      200,
		}
		expected = "accountName=testaccount&limit=200&offset=200&realm=sony&track=true&type=league"
	)

	params := opts.toQueryParams()
	if params != expected {
		t.Fatalf("failed to convert ladder options to query params. expected %s, got %s",
			expected, params)
	}
}

func TestLabyrinthLadderOptionstoQueryParams(t *testing.T) {
	var (
		opts = GetLadderOptions{
			Realm:               "xbox",
			Type:                "labyrinth",
			UniqueIDs:           false,
			LabyrinthDifficulty: "Normal",
			LabyrinthStartTime:  1565283600,
			limit:               200,
			offset:              400,
		}
		expected = "difficulty=Normal&limit=200&offset=400&realm=xbox&start=1565283600&track=false&type=labyrinth"
	)

	params := opts.toQueryParams()
	if params != expected {
		t.Fatalf("failed to convert ladder options to query params. expected %s, got %s",
			expected, params)
	}
}

func TestValidateLadderOptions(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Type:   "league",
		limit:  200,
		offset: 0,
	}
	if err := validateGetLadderOptions(opts); err != nil {
		t.Fatalf("failed to validate ladder options: %v", err)
	}
}

func TestValidateLadderOptionsWithMissingID(t *testing.T) {
	opts := GetLadderOptions{
		Realm: "pc",
		Type:  "league",
	}
	if err := validateGetLadderOptions(opts); err != ErrMissingID {
		t.Fatalf("failed to detect missing id in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidRealm(t *testing.T) {
	opts := GetLadderOptions{
		ID:    "Standard",
		Realm: "testrealm",
		Type:  "league",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidRealm {
		t.Fatalf("failed to detect invalid realm in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidLimit(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Type:   "league",
		limit:  201,
		offset: 0,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLimit {
		t.Fatalf("failed to detect invalid limit in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidOffset(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Type:   "league",
		limit:  200,
		offset: 15001,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidOffset {
		t.Fatalf("failed to detect invalid offset in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidType(t *testing.T) {
	opts := GetLadderOptions{
		ID:    "Standard",
		Realm: "pc",
		Type:  "testtype",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLadderType {
		t.Fatalf("failed to detect invalid type in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidDifficulty(t *testing.T) {
	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Type:                "labyrinth",
		LabyrinthDifficulty: "testdifficulty",
		limit:               200,
		offset:              0,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidDifficulty {
		t.Fatalf("failed to detect invalid difficulty in ladder options")
	}
}

func TestValidateLadderOptionsWithNegativeStartTime(t *testing.T) {
	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Type:                "labyrinth",
		LabyrinthDifficulty: "Normal",
		LabyrinthStartTime:  -1,
		limit:               200,
		offset:              0,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLabyrinthStartTime {
		t.Fatalf("failed to detect negative start time in ladder options")
	}
}

func TestValidateLadderOptionsWithEarlyStartTime(t *testing.T) {
	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Type:                "labyrinth",
		LabyrinthDifficulty: "Normal",
		LabyrinthStartTime:  1400000000,
		limit:               200,
		offset:              0,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLabyrinthStartTime {
		t.Fatalf("failed to detect too-early start time in ladder options")
	}
}

func TestGetLadderPage(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "xbox",
		Type:                "labyrinth",
		UniqueIDs:           false,
		LabyrinthDifficulty: "Normal",
		LabyrinthStartTime:  1565283600,
		limit:               200,
	}

	_, err := c.getLadderPage(opts)
	if err != nil {
		t.Fatalf("failed to get ladder page: %v", err)
	}
}

func TestGetLadderPageFailure(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}
	_, err := c.getLadderPage(GetLadderOptions{
		ID:    "Nonexistent",
		limit: 1,
	})
	if err != ErrNotFound {
		t.Fatal("failed to detect ladder retrieval failure")
	}
}

func TestGetLadderRequestFailure(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}
	opts := GetLadderOptions{
		ID: "test",
	}
	_, err := c.GetLadder(opts)
	if err != ErrNotFound {
		t.Fatal("failed to detect ladder request failure")
	}
}

func TestGetLadderPageRequestFailure(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}
	opts := GetLadderOptions{
		ID:    "test",
		limit: 200,
	}
	_, err := c.getLadderPage(opts)
	if err != ErrNotFound {
		t.Fatal("failed to detect ladder request failure")
	}
}

func TestParseLadderResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/ladder.json")
	if err != nil {
		t.Fatalf("failed to load fixture for ladder response test: %v", err)
	}

	_, err = parseLadderResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse ladder response: %v", err)
	}
}

func TestParseInvalidLadderResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/invalid.json")
	if err != nil {
		t.Fatalf("failed to load fixture for ladder response test: %v", err)
	}

	_, err = parseLadderResponse(resp)
	if err == nil {
		t.Fatal("failed to detect ladder parsing failure")
	}
}

func TestGetLadder(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Type:                "labyrinth",
		UniqueIDs:           false,
		LabyrinthDifficulty: "Cruel",
		LabyrinthStartTime:  1560186000,
	}

	_, err := c.GetLadder(opts)
	if err != nil {
		t.Fatalf("failed to get ladder page: %v", err)
	}
}

func TestGetSmallLadder(t *testing.T) {
	c := client{
		host:       testHost,
		useSSL:     false,
		useCache:   false,
		limiter:    newRateLimiter(UnlimitedRate, UnlimitedRate),
		httpClient: testClient,
	}

	opts := GetLadderOptions{
		ID:        "Standard",
		Realm:     "xbox",
		Type:      "labyrinth",
		UniqueIDs: false,
	}

	_, err := c.GetLadder(opts)
	if err != nil {
		t.Fatalf("failed to get small ladder page: %v", err)
	}
}
