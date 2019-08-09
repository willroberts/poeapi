// leaguetimer prints the time remaining for the current challenge league.
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

	m, err := client.GetPVPMatches(poeapi.GetPVPMatchesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d upcoming PVP matches.", len(m))
}
