// ratetest attempts to reverse engineer the rate limits for the PoE API.
// API latency can be anywhere from 100ms to over 100ms, so synchronous usage
// will rarely exceed 4-5 requests per second. For this reason, the test code is
// asynchronous to maximize the number of requests per second we can send.

// Results:
// API endpoints succeed reliably (for 100 sequential requests) at rates up to
// five per second.

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	initialRate int = 1
	maxRate     int = 10
	testCount   int = 10

	//testURL string = "https://api.pathofexile.com/public-stash-tabs" // 1rps.
	testURL string = "https://api.pathofexile.com/leagues/Standard" // 5rps.
	//testURL string = "https://api.pathofexile.com/ladders/Standard" // 5rps.
	//testURL string = "https://api.pathofexile.com/pvp-matches" // 5rps.
	//testURL string = "https://api.pathofexile.com/league-rules" // no limit.

	verbose bool = true
)

// Keys: Whole number of requests per second.
// Values: Milliseconds between each request.
var rateToInterval = map[int]float64{
	1:  1000,
	2:  500,
	3:  333,
	4:  250,
	5:  200,
	6:  167,
	7:  143,
	8:  125,
	9:  111,
	10: 100,
}

type limitfinder struct {
	rate      int
	successes int
	lock      sync.Mutex
}

func (lf *limitfinder) ParseCode(code int) bool {
	lf.lock.Lock()
	var limitFound bool

	if code == http.StatusOK {
		lf.successes++
	}
	if lf.successes == testCount {
		lf.rate++
		log.Println("rate limit:", lf.rate)
		if lf.rate == maxRate {
			limitFound = true
		}
		lf.successes = 0
	}
	if code == http.StatusTooManyRequests {
		limitFound = true
	}

	lf.lock.Unlock()
	return limitFound
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	doneCh := make(chan struct{}, 1)

	log.Println("sending requests to api.pathofexile.com")
	lf := limitfinder{rate: initialRate}
	log.Println("rate limit:", lf.rate)

	for {
		select {
		case <-sigCh:
			log.Println("ctrl-c received, exiting")
			return
		default:
			go func() {
				code, err := sendRequest()
				if err != nil {
					log.Fatal("fatal error:", err)
				}
				if limitFound := lf.ParseCode(code); limitFound {
					doneCh <- struct{}{}
				}
			}()

			select {
			case <-doneCh:
				log.Printf("rate limit reached at %d requests per second", lf.rate)
				return
			default:
			}

			sleepLen := time.Duration(rateToInterval[lf.rate]) * time.Millisecond
			time.Sleep(sleepLen)
		}
	}
}

func sendRequest() (int, error) {
	start := time.Now()

	resp, err := http.Get(testURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if verbose {
		log.Println("code:", resp.StatusCode, "latency:", time.Since(start))
	}

	return resp.StatusCode, nil
}
