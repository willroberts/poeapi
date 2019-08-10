// leaguetimer prints the time remaining for the current challenge league.
package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/willroberts/poeapi"
)

func main() {
	client, err := poeapi.NewAPIClient(poeapi.DefaultClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	leagues, err := client.GetLeagues(poeapi.GetLeaguesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var challenge poeapi.League
	var found bool
	for _, l := range leagues {
		if (l.EndTime != time.Time{}) {
			challenge = l
			found = true
		}
	}

	if !found {
		log.Fatal("failed to find challenge league")
	}

	h, m, s := timeUntil(challenge.EndTime)
	fmt.Printf("%s league has %d hours, %d minutes, and %d seconds remaining.",
		challenge.Name, h, m, s)
}

func timeUntil(t time.Time) (hours, minutes, seconds int) {
	totalSecondsRemaining := float64(time.Until(t) / time.Second)
	hoursRemaining, fractionalHours := math.Modf(totalSecondsRemaining / 3600)
	minutesRemaining, fractionalMinutes := math.Modf(fractionalHours * 60)
	secondsRemaining := fractionalMinutes * 60

	return int(hoursRemaining), int(minutesRemaining), int(secondsRemaining)
}
