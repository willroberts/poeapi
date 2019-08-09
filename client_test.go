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
