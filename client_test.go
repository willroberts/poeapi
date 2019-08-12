package poeapi

import (
	"testing"
	"time"
)

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
		Host:           "google.com",
		NinjaHost:      DefaultNinjaHost,
		RateLimit:      DefaultRateLimit,
		StashRateLimit: DefaultStashRateLimit,
		RequestTimeout: DefaultRequestTimeout,
	}
	if err := validateClientOptions(opts); err != ErrInvalidHost {
		t.Fatal("failed to detect invalid host option")
	}
}

func TestValidateOptionsInvalidNinjaHost(t *testing.T) {
	opts := ClientOptions{
		Host:           DefaultHost,
		NinjaHost:      "",
		RateLimit:      DefaultRateLimit,
		StashRateLimit: DefaultStashRateLimit,
		RequestTimeout: DefaultRequestTimeout,
	}
	if err := validateClientOptions(opts); err != ErrInvalidNinjaHost {
		t.Fatal("failed to detect invalid ninja host option")
	}
}

func TestValidateOptionsInvalidCacheSize(t *testing.T) {
	opts := ClientOptions{
		Host:           DefaultHost,
		NinjaHost:      DefaultNinjaHost,
		UseCache:       true,
		CacheSize:      0,
		RateLimit:      DefaultRateLimit,
		StashRateLimit: DefaultStashRateLimit,
		RequestTimeout: DefaultRequestTimeout,
	}
	if err := validateClientOptions(opts); err != ErrInvalidCacheSize {
		t.Fatal("failed to detect invalid cache size")
	}
}

func TestValidateOptionsInvalidRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:           DefaultHost,
		NinjaHost:      DefaultNinjaHost,
		CacheSize:      DefaultCacheSize,
		RateLimit:      -1,
		StashRateLimit: DefaultStashRateLimit,
		RequestTimeout: DefaultRequestTimeout,
	}
	if err := validateClientOptions(opts); err != ErrInvalidRateLimit {
		t.Fatal("failed to detect invalid rate limit option")
	}
}

func TestValidateOptionsInvalidStashRateLimit(t *testing.T) {
	opts := ClientOptions{
		Host:           DefaultHost,
		NinjaHost:      DefaultNinjaHost,
		CacheSize:      DefaultCacheSize,
		RateLimit:      DefaultRateLimit,
		StashRateLimit: -1,
		RequestTimeout: DefaultRequestTimeout,
	}
	if err := validateClientOptions(opts); err != ErrInvalidStashRateLimit {
		t.Fatal("failed to detect invalid stash rate limit option")
	}
}

func TestValidateOptionsInvalidRequestTimeout(t *testing.T) {
	opts := ClientOptions{
		Host:           DefaultHost,
		NinjaHost:      DefaultNinjaHost,
		CacheSize:      DefaultCacheSize,
		RateLimit:      DefaultRateLimit,
		StashRateLimit: DefaultStashRateLimit,
		RequestTimeout: 0 * time.Millisecond,
	}
	if err := validateClientOptions(opts); err != ErrInvalidRequestTimeout {
		t.Fatal("failed to detect invalid request timeout option")
	}
}

func TestClientDNSDialer(t *testing.T) {
	c, err := NewAPIClient(ClientOptions{
		Host:           testHost,
		NinjaHost:      testHost,
		UseSSL:         false,
		UseCache:       false,
		UseDNSCache:    true,
		RateLimit:      UnlimitedRate,
		StashRateLimit: UnlimitedRate,
		RequestTimeout: testTimeout,
	})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	if _, err := c.GetLeagueRules(); err != nil {
		t.Fatalf("failed to get http with dns caching: %v", err)
	}
}
