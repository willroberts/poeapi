package poeapi

import "testing"

func TestLadderOptionsToQueryParams(t *testing.T) {
	var (
		opts = GetLadderOptions{
			Realm:       "sony",
			Limit:       200,
			Offset:      200,
			Type:        "league",
			UniqueIDs:   true,
			AccountName: "testaccount",
		}
		expected = "accountName=testaccount&limit=200&offset=200&realm=sony&track=true&type=league"
	)

	params := opts.ToQueryParams()
	if params != expected {
		t.Fatalf("failed to convert ladder options to query params. expected %s, got %s",
			expected, params)
	}
}

func TestLabyrinthLadderOptionsToQueryParams(t *testing.T) {
	var (
		opts = GetLadderOptions{
			Realm:               "xbox",
			Limit:               200,
			Offset:              400,
			Type:                "labyrinth",
			UniqueIDs:           false,
			LabyrinthDifficulty: "Normal",
			LabyrinthStartTime:  1565283600,
		}
		expected = "difficulty=Normal&limit=200&offset=400&realm=xbox&start=1565283600&track=false&type=labyrinth"
	)

	params := opts.ToQueryParams()
	if params != expected {
		t.Fatalf("failed to convert ladder options to query params. expected %s, got %s",
			expected, params)
	}
}

func TestValidateLadderOptions(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Limit:  200,
		Offset: 0,
		Type:   "league",
	}
	if err := validateGetLadderOptions(opts); err != nil {
		t.Fatalf("failed to validate ladder options: %v", err)
	}
}

func TestValidateLadderOptionsWithMissingID(t *testing.T) {
	opts := GetLadderOptions{
		Realm:  "pc",
		Limit:  200,
		Offset: 0,
		Type:   "league",
	}
	if err := validateGetLadderOptions(opts); err != ErrMissingID {
		t.Fatalf("failed to detect missing id in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidRealm(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "testrealm",
		Limit:  200,
		Offset: 0,
		Type:   "league",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidRealm {
		t.Fatalf("failed to detect invalid realm in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidLimit(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Limit:  201,
		Offset: 0,
		Type:   "league",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLimit {
		t.Fatalf("failed to detect invalid limit in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidOffset(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Limit:  200,
		Offset: 15001,
		Type:   "league",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidOffset {
		t.Fatalf("failed to detect invalid offset in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidType(t *testing.T) {
	opts := GetLadderOptions{
		ID:     "Standard",
		Realm:  "pc",
		Limit:  200,
		Offset: 0,
		Type:   "testtype",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLadderType {
		t.Fatalf("failed to detect invalid type in ladder options")
	}
}

func TestValidateLadderOptionsWithInvalidDifficulty(t *testing.T) {
	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Limit:               200,
		Offset:              0,
		Type:                "labyrinth",
		LabyrinthDifficulty: "testdifficulty",
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidDifficulty {
		t.Fatalf("failed to detect invalid difficulty in ladder options")
	}
}

func TestValidateLadderOptionsWithNegativeStartTime(t *testing.T) {
	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Limit:               200,
		Offset:              0,
		Type:                "labyrinth",
		LabyrinthDifficulty: "Normal",
		LabyrinthStartTime:  -1,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLabyrinthStartTime {
		t.Fatalf("failed to detect negative start time in ladder options")
	}
}

func TestValidateLadderOptionsWithEarlyStartTime(t *testing.T) {
	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Limit:               200,
		Offset:              0,
		Type:                "labyrinth",
		LabyrinthDifficulty: "Normal",
		LabyrinthStartTime:  1400000000,
	}
	if err := validateGetLadderOptions(opts); err != ErrInvalidLabyrinthStartTime {
		t.Fatalf("failed to detect too-early start time in ladder options")
	}
}

func TestGetLadderPage(t *testing.T) {
	c := client{
		host:     defaultHost,
		useSSL:   true,
		useCache: false,
		limiter:  newRateLimiter(defaultRateLimit, defaultStashTabRateLimit),
	}

	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "xbox",
		Limit:               200,
		Type:                "labyrinth",
		UniqueIDs:           false,
		LabyrinthDifficulty: "Normal",
		LabyrinthStartTime:  1565283600,
	}

	_, err := c.getLadderPage(opts)
	if err != nil {
		t.Fatalf("failed to get ladder page: %v", err)
	}
}

func TestGetLadderPageFailure(t *testing.T) {
	c := client{
		host:     defaultHost,
		useSSL:   true,
		useCache: false,
		limiter:  newRateLimiter(defaultRateLimit, defaultStashTabRateLimit),
	}

	opts := GetLadderOptions{ID: "Nonexistent"}

	_, err := c.getLadderPage(opts)
	if err == nil {
		t.Fatal("failed to detect ladder request failure")
	}
}

func TestParseLadderResponse(t *testing.T) {
	resp, err := loadFixture("fixtures/ladders.json")
	if err != nil {
		t.Fatalf("failed to load fixture for ladder response test: %v", err)
	}

	_, err = parseLadderResponse(resp)
	if err != nil {
		t.Fatalf("failed to parse ladder response: %v", err)
	}
}

func TestParseInvalidLadderResponse(t *testing.T) {
	var resp = "{\"total\": \"invalid\"}"
	_, err := parseLadderResponse(resp)
	if err == nil {
		t.Fatal("failed to detect ladder parsing failure")
	}
}

func TestGetLadder(t *testing.T) {
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for ladder test: %v", err)
	}

	opts := GetLadderOptions{
		ID:                  "Standard",
		Realm:               "pc",
		Type:                "labyrinth",
		UniqueIDs:           false,
		LabyrinthDifficulty: "Cruel",
		LabyrinthStartTime:  1560186000,
	}

	_, err = c.GetLadder(opts)
	if err != nil {
		t.Fatalf("failed to get ladder page: %v", err)
	}
}

func TestGetSmallLadder(t *testing.T) {
	c, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client for small ladder test: %v", err)
	}

	opts := GetLadderOptions{
		ID:        "Standard",
		Realm:     "xbox",
		Type:      "labyrinth",
		UniqueIDs: false,
	}

	_, err = c.GetLadder(opts)
	if err != nil {
		t.Fatalf("failed to get small ladder page: %v", err)
	}
}
