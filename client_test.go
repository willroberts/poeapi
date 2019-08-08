package poeapi

import "testing"

func TestNewAPIClient(t *testing.T) {
	_, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fail()
	}
}
