package poeapi

import "testing"

func TestNewAPIClient(t *testing.T) {
	_, err := NewAPIClient(DefaultClientOptions)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
}

func TestNewAPIClientWithInvalidOptions(t *testing.T) {
	var (
		opts = ClientOptions{
			Host: "google.com",
		}
	)
	_, err := NewAPIClient(opts)
	if err != ErrInvalidHost {
		t.Fatal("failed to detect invalid options")
	}
}

func TestValidateOptions(t *testing.T) {
	opts := DefaultClientOptions
	if err := validateClientOptions(opts); err != nil {
		t.Fatalf("failed to validate options: %v", err)
	}
}

func TestValidateOptionsInvalidHost(t *testing.T) {
	opts := ClientOptions{
		Host:              "google.com",
		RateLimit:         DefaultRateLimit,
		StashRateLimit: DefaultStashRateLimit,
	}
	if err := validateClientOptions(opts); err != ErrInvalidHost {
		t.Fatal("failed to detect invalid host option")
	}
}

func TestValidateOptionsInvalidCacheSize(t *testing.T) {
	opts := ClientOptions{
		Host:              DefaultHost,
		CacheSize:         0,
		RateLimit:         DefaultRateLimit,
		StashRateLimit: DefaultStashRateLimit,
	}
	if err := validateClientOptions(opts); err != ErrInvalidCacheSize {
		t.Fatal("failed to detect invalid cache size")
	}
}

func TestValidateOptionsInvalidRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:              DefaultHost,
		CacheSize:         DefaultCacheSize,
		RateLimit:         0,
		StashRateLimit: DefaultStashRateLimit,
	}
	if err := validateClientOptions(opts); err != ErrInvalidRateLimit {
		t.Fatal("failed to detect invalid rate limit option")
	}
}

func TestValidateOptionsInvalidStashRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:              DefaultHost,
		CacheSize:         DefaultCacheSize,
		RateLimit:         DefaultRateLimit,
		StashRateLimit: 0,
	}
	if err := validateClientOptions(opts); err != ErrInvalidStashRateLimit {
		t.Fatal("failed to detect invalid stash rate limit option")
	}
}
