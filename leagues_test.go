package poeapi

import "testing"

func TestGetAllLeagues(t *testing.T) {
	client, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fail()
	}
	_, err = client.GetAllLeagues()
	if err != nil {
		t.Fail()
	}
}

func TestGetCurrentChallengeLeague(t *testing.T) {
	client, err := NewAPIClient(DefaultOptions)
	if err != nil {
		t.Fail()
	}
	_, err = client.GetCurrentChallengeLeague()
	if err != nil {
		t.Fail()
	}
}
