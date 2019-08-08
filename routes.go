package poeapi

import "fmt"

const (
	leaguesEndpoint     = "/leagues"
	laddersEndpoint     = "/ladders"
	stashTabsEndpoint   = "/public-stash-tabs"
	pvpMatchesEndpoint  = "/pvp-matches"
	leagueRulesEndpoint = "/league-rules"
)

// formatURL contructs a valid URL from the configured scheme (http or https),
// the configured host (generally api.pathofexile.com), and the requested
// endpoint.
func (c *client) formatURL(endpoint string) string {
	protocol := "https"
	if !c.useSSL {
		protocol = "http"
	}

	return fmt.Sprintf("%s://%s%s", protocol, c.host, endpoint)
}
