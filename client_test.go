package poeapi

import "testing"

func TestNewAPIClient(t *testing.T) {
	_, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fail()
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
		t.Fail()
	}
}