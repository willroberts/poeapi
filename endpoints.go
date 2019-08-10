package poeapi

import "fmt"

const (
	leaguesEndpoint     = "/leagues"
	leagueRulesEndpoint = "/league-rules"
	laddersEndpoint     = "/ladders"
	pvpMatchesEndpoint  = "/pvp-matches"
	stashTabsEndpoint   = "/public-stash-tabs"

	httpProtocol  = "http"
	httpsProtocol = "https"
)

// formatURL contructs a valid URL from the configured scheme (http or https),
// the configured host (generally api.pathofexile.com), and the requested
// endpoint.
func (c *client) formatURL(endpoint string) string {
	protocol := httpsProtocol
	if !c.useSSL {
		protocol = httpProtocol
	}

	return fmt.Sprintf("%s://%s%s", protocol, c.host, endpoint)
}
