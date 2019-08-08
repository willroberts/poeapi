package poeapi

import "testing"

func TestValidateOptions(t *testing.T) {
	opts := DefaultOptions
	if err := validateOptions(opts); err != nil {
		t.Fatalf("failed to validate options: %v", err)
	}
}

func TestValidateOptionsInvalidHost(t *testing.T) {
	opts := ClientOptions{
		Host:              "google.com",
		RateLimit:         1,
		StashTabRateLimit: 1,
	}
	if err := validateOptions(opts); err != ErrInvalidHost {
		t.Fatal("failed to detect invalid host option")
	}
}

func TestValidateOptionsInvalidRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:              defaultHost,
		CacheSize:         defaultCacheSize,
		RateLimit:         0,
		StashTabRateLimit: defaultStashTabRateLimit,
	}
	if err := validateOptions(opts); err != ErrInvalidRateLimit {
		t.Fatal("failed to detect invalid rate limit option")
	}
}

func TestValidateOptionsInvalidStashTabRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:              defaultHost,
		CacheSize:         defaultCacheSize,
		RateLimit:         defaultRateLimit,
		StashTabRateLimit: 0,
	}
	if err := validateOptions(opts); err != ErrInvalidStashTabRateLimit {
		t.Fatal("failed to detect invalid stash tab rate limit option")
	}
}
