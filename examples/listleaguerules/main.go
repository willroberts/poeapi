// listleaguerules prints all league rules from the API.
package main

import (
	"fmt"
	"log"

	"github.com/willroberts/poeapi"
)

func main() {
	client, err := poeapi.NewAPIClient(poeapi.DefaultClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	rules, err := client.GetLeagueRules()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rules {
		fmt.Printf("%s: %s\n", r.Name, r.Description)
	}
}
