package poeapi

import "fmt"

const (
	leaguesEndpoint     = "/leagues"
	laddersEndpoint     = "/ladders"
	stashTabsEndpoint   = "/public-stash-tabs"
	pvpMatchesEndpoint  = "/pvp-matches"
	leagueRulesEndpoint = "/league-rules"
)

func (c *client) formatURL(endpoint string) string {
	protocol := "https"
	if !c.useSSL {
		protocol = "http"
	}

	return fmt.Sprintf("%s://%s%s", protocol, c.host, endpoint)
}
