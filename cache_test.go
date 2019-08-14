package poeapi

import (
	"container/ring"
	"testing"
	"time"
)

func TestNewResponseCache(t *testing.T) {
	_, err := newResponseCache(1)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}
}

func TestInvalidCacheSize(t *testing.T) {
	_, err := newResponseCache(0)
	if err != ErrInvalidCacheSize {
		t.Fatal("failed to detect invalid cache size")
	}
}

func TestCacheEviction(t *testing.T) {
	cache, err := newResponseCache(10)
	if err != nil {
		t.Fatalf("failed to create cache for operations test: %v", err)
	}

	cache.Set("1", "A")
	cache.Set("2", "B")
	cache.Set("3", "C")
	cache.Set("4", "D")
	cache.Set("5", "E")
	cache.Set("6", "F")
	cache.Set("7", "G")
	cache.Set("8", "H")
	cache.Set("9", "I")

	val, err := cache.Get("5")
	if err != nil {
		t.Fatalf("failed to get from cache: %v", err)
	}
	if val != "E" {
		t.Fatalf("unexpected cache result: got %s, expected E", val)
	}

	cache.Set("foo", "foo")
	cache.Set("bar", "bar")

	val, err = cache.Get("1")
	if err != ErrNotFoundInCache {
		t.Fatalf("failed to get from cache: %v", err)
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
	cache, err := newResponseCache(10)
	if err != nil {
		t.Fatalf("failed to create cache for existing key test: %v", err)
	}
	cache.Set("foo", "bar")
	cache.Set("foo", "bar")
}

func TestDNSCacheResolve(t *testing.T) {
	var (
		host = "localhost"
		d    = newDNSCache()
	)
	if err := d.resolve(host); err != nil {
		t.Fatalf("failed to resolve in dns cache: %v", err)
	}
}

func TestDNSCacheGetUnresolved(t *testing.T) {
	var (
		host = "localhost"
		d    = newDNSCache()
	)
	_, err := d.Get(host)
	if err != nil {
		t.Fatalf("failed to get ip from dns cache: %v", err)
	}
}

func TestDNSCacheResolutionFailure(t *testing.T) {
	var (
		host = "11111"
		d    = newDNSCache()
	)
	if err := d.resolve(host); err == nil {
		t.Fatal("failed to detect dns resolution failure")
	}
}

func TestDNSCacheGetResolutionFailure(t *testing.T) {
	var (
		host = "11111"
		d    = newDNSCache()
	)
	if _, err := d.Get(host); err == nil {
		t.Fatal("failed to detect dns resolution failure")
	}
}

func TestDNSCacheInvalidEntry(t *testing.T) {
	var (
		host       = "test"
		d          = newDNSCache()
		entryCount = 3
	)
	d.IPs[host] = ring.New(entryCount)
	for i := 0; i < entryCount; i++ {
		d.IPs[host].Value = i
		d.IPs[host].Next()
	}
	if _, err := d.Get(host); err != ErrInvalidAddress {
		t.Fatal("failed to detect invalid dns cache entry")
	}
}
