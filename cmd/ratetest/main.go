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
	initialFreq  float64 = 0.5 // 4 RPS.
	minFreq      float64 = 0.1 // 10 RPS.
	successCount int     = 100

	// Testing suggests that /leagues/ URLs have no rate limit.
	//testURL string = "https://api.pathofexile.com/leagues/Standard" // 5rps.
	testURL string = "https://api.pathofexile.com/ladders/Standard" // 1/.15
)

type limitfinder struct {
	freq      float64
	successes int
	lock      sync.Mutex
}

func (lf *limitfinder) ParseCode(code int) bool {
	lf.lock.Lock()
	var limitFound bool

	if code == http.StatusOK {
		lf.successes++
	}
	if lf.successes == successCount {
		lf.freq = lf.freq - 0.05
		log.Println("new frequency:", lf.freq)
		if lf.freq < minFreq {
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
	lf := limitfinder{freq: initialFreq}

	log.Println("sending requests to api.pathofexile.com")
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
				log.Printf("testing ended at frequency %.2f", lf.freq)
				return
			default:
			}

			sleepLen := time.Duration(lf.freq*1000) * time.Millisecond
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
	log.Println("code:", resp.StatusCode, "latency:", time.Since(start))
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
