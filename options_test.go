package poeapi

import "testing"

func TestValidateOptions(t *testing.T) {
	opts := DefaultOptions
	if err := validateOptions(opts); err != nil {
		t.Fail()
	}
}

func TestValidateOptionsInvalidHost(t *testing.T) {
	opts := ClientOptions{
		Host:              "google.com",
		RateLimit:         1,
		StashTabRateLimit: 1,
	}
	if err := validateOptions(opts); err != ErrInvalidHost {
		t.Fail()
	}
}

func TestValidateOptionsInvalidRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:              "api.pathofexile.com",
		RateLimit:         0,
		StashTabRateLimit: 1,
	}
	if err := validateOptions(opts); err != ErrInvalidRateLimit {
		t.Fail()
	}
}

func TestValidateOptionsInvalidStashTabRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:              "api.pathofexile.com",
		RateLimit:         1,
		StashTabRateLimit: 0,
	}
	if err := validateOptions(opts); err != ErrInvalidStashTabRateLimit {
		t.Fail()
	}
}
