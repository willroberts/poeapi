/*
Package poeapi provides a client for the Path of Exile API.

Features include support for all API endpoints, thread-safe operations, built-in
caching and ratelimiting, and no reliance on external packages.

Usage:

	clientOpts := poeapi.ClientOptions{
	    Host:              "api.pathofexile.com", // The primary API domain.
	    NinjaHost:         "poe.ninja",           // Used to get latest stash ID.
    	UseSSL:            true,                  // Use HTTPS for requests.
    	UseCache:          true,                  // Enable the in-memory cache.
    	UseDNSCache:       true,                  // Enable the in-memory DNS resolution cache.
    	CacheSize:         200,                   // Number of items to store.
    	RateLimit:         4.0,                   // Requests per second.
    	StashRateLimit:    1.0,                   // Requests per second for trade API.
    	RequestTimeout:    5 * time.Second        // Time to wait before canceling requests.
	} // This is equivalent to poeapi.DefaultClientOptions.

	client, err := poeapi.NewAPIClient(clientOpts)
	if err != nil {
	    // Handle error.
	}

	ladder, err := client.GetLadder(poeapi.GetLadderOptions{
    	ID:    "SSF Hardcore",
    	Realm: "pc",
    	Type:  "league",
	})
	// Etc.

See the `APIClient` interface for the full description of methods available.
*/
package poeapi
