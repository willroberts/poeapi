package poeapi

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	_, err := newCache(1)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}
}

func TestInvalidCacheSize(t *testing.T) {
	_, err := newCache(0)
	if err != ErrInvalidCacheSize {
		t.Fatal("failed to detect invalid cache size")
	}
}

func TestCacheEviction(t *testing.T) {
	cache, err := newCache(10)
	if err != nil {
		t.Fatalf("failed to create cache for operations test: %v", err)
	}

	cache.Add("1", "A")
	cache.Add("2", "B")
	cache.Add("3", "C")
	cache.Add("4", "D")
	cache.Add("5", "E")
	cache.Add("6", "F")
	cache.Add("7", "G")
	cache.Add("8", "H")
	cache.Add("9", "I")

	val := cache.Get("5")
	if val != "E" {
		t.Fatalf("unexpected cache result: got %s, expected E", val)
	}

	cache.Add("foo", "foo")
	cache.Add("bar", "bar")

	val = cache.Get("1")
	if val != "" {
		t.Fatalf("unexpected cache result: got %s, expected \"\"", val)
	}
}

func TestCacheLatency(t *testing.T) {
	c, err := NewAPIClient(ClientOptions{
		Host:           testHost,
		NinjaHost:      DefaultNinjaHost,
		UseSSL:         false,
		UseCache:       true,
		CacheSize:      DefaultCacheSize,
		RateLimit:      UnlimitedRate,
		StashRateLimit: UnlimitedRate,
		RequestTimeout: testTimeout,
	})
	if err != nil {
		t.Fatalf("failed to create client for latency test: %v", err)
	}

	_, err = c.GetLeague(GetLeagueOptions{ID: "Standard"})
	if err != nil {
		t.Fatalf("failed to get all leagues for latency test: %v", err)
	}

	start2 := time.Now()
	_, err = c.GetLeague(GetLeagueOptions{ID: "Standard"})
	if err != nil {
		t.Fatalf("failed to get all leagues for second latency test: %v", err)
	}
	duration2 := time.Since(start2)

	if duration2 > 1*time.Millisecond {
		t.Fatal("cache test took longer than 1ms")
	}
}

func TestCacheExistingKey(t *testing.T) {
	cache, err := newCache(10)
	if err != nil {
		t.Fatalf("failed to create cache for existing key test: %v", err)
	}
	cache.Add("foo", "bar")
	cache.Add("foo", "bar")
}
