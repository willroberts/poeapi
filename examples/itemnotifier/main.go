package main

import (
	"log"
	"time"

	"github.com/willroberts/poeapi"
)

var (
	targetItem   = "Kaom's Heart"
	targetLeague = "Standard"
)

func main() {
	client, err := poeapi.NewAPIClient(poeapi.ClientOptions{
		Host:           poeapi.DefaultHost,
		NinjaHost:      poeapi.DefaultNinjaHost,
		UseSSL:         true,
		UseCache:       true,
		CacheSize:      100,
		RateLimit:      poeapi.DefaultRateLimit,
		StashRateLimit: poeapi.DefaultStashRateLimit,
		RequestTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	latest, err := client.GetLatestStashID()
	if err != nil {
		log.Fatal(err)
	}

	for {
		stashes, err := client.GetStashes(poeapi.GetStashOptions{
			ID: latest,
		})
		if err != nil {
			log.Fatal(err)
		}

		if stashes.NextChangeID == latest {
			continue
		}
		latest = stashes.NextChangeID

		for _, s := range stashes.Stashes {
			for _, i := range s.Items {
				if i.Name == targetItem && i.League == targetLeague {
					if i.Note != "" {
						log.Printf("%s is selling %s for %s in %s league.",
							s.LastCharacterName, targetItem, i.Note, targetLeague)
					} else {
						log.Printf("%s is selling %s in %s league.",
							s.LastCharacterName, targetItem, targetLeague)
					}
				}
			}
		}
	}
}
