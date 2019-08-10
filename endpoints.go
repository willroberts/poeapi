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
	if c.useSSL {
		return fmt.Sprintf("%s://%s%s", httpsProtocol, c.host, endpoint)
	}
	return fmt.Sprintf("%s://%s%s", httpProtocol, c.host, endpoint)
}
