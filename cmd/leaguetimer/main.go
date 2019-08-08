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
	client, err := poeapi.NewAPIClient(poeapi.DefaultOptions)
	if err != nil {
		log.Fatal(err)
	}

	l, err := client.GetCurrentChallengeLeague()
	if err != nil {
		log.Fatal(err)
	}

	h, m, s := timeUntil(l.EndTime)
	fmt.Printf("%s league has %d hours, %d minutes, and %d seconds remaining.",
		l.Name, h, m, s)
}

func timeUntil(t time.Time) (hours, minutes, seconds int) {
	totalSecondsRemaining := float64(time.Until(t) / time.Second)
	hoursRemaining, fractionalHours := math.Modf(totalSecondsRemaining / 3600)
	minutesRemaining, fractionalMinutes := math.Modf(fractionalHours * 60)
	secondsRemaining := fractionalMinutes * 60

	return int(hoursRemaining), int(minutesRemaining), int(secondsRemaining)
}
