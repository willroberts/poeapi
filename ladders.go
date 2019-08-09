package poeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"sync"
)

const (
	maxLimit = 200
	maxPages = 75 // 15000 % 200
)

// GetLadderOptions contains the request parameters for the ladder endpoint.
// All parameters are optional with the exception of ID.
type GetLadderOptions struct {
	// The name of the league whose ladder you want to retrieve.
	ID string `json:"id"`

	// The realm of the ladder.
	// Valid options: 'pc', 'xbox', or 'sony'.
	Realm string `json:"realm"`

	// Number of results to retrieve. Default: 20. Maximum: 200.
	Limit int `json:"limit"`

	// Position at which to start retrieving results. Default: 0.
	Offset int `json:"offset"`

	// The type of league whose ladder you want to retrieve.
	// Valid options: 'league', 'pvp', or 'labyrinth'.
	Type string `json:"type"`

	// Associate UUIDs with each character returned for tracking purposes.
	UniqueIDs bool `json:"track"`

	// Only include the given account in results.
	AccountName string `json:"accountName"`

	// Difficulty of the Labyrinth ladder to retrieve.
	// Valid options: 'Normal', 'Cruel', or 'Merciless'.
	LabyrinthDifficulty string `json:"difficulty"`

	// Start time of the Labyrinth ladder to retrieve. This is a Unix timestamp.
	LabyrinthStartTime int `json:"start"`
}

// ToQueryParams converts ladder options to a URL parameter string.
func (opts GetLadderOptions) ToQueryParams() string {
	u := url.Values{}
	if opts.Realm != "" {
		u.Add("realm", opts.Realm)
	}
	if opts.Limit != 0 {
		u.Add("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Offset != 0 {
		u.Add("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Type != "" {
		u.Add("type", opts.Type)
	}
	u.Add("track", strconv.FormatBool(opts.UniqueIDs))
	if opts.AccountName != "" && opts.Type == "league" {
		u.Add("accountName", opts.AccountName)
	}
	if opts.LabyrinthDifficulty != "" && opts.Type == "labyrinth" {
		u.Add("difficulty", opts.LabyrinthDifficulty)
	}
	if opts.LabyrinthStartTime != 0 && opts.Type == "labyrinth" {
		u.Add("start", strconv.Itoa(opts.LabyrinthStartTime))
	}
	return u.Encode()
}

func (c *client) getLadderPage(opts GetLadderOptions) (Ladder, error) {
	url := fmt.Sprintf("%s/%s?%s", c.formatURL(laddersEndpoint), opts.ID,
		opts.ToQueryParams())
	resp, err := c.get(url)
	if err != nil {
		return Ladder{}, err
	}
	return parseLadderResponse(resp)
}

func parseLadderResponse(resp string) (Ladder, error) {
	ladder := Ladder{}
	if err := json.Unmarshal([]byte(resp), &ladder); err != nil {
		return Ladder{}, err
	}
	return ladder, nil
}

func (c *client) GetLadder(opts GetLadderOptions) (Ladder, error) {
	entries := make([]LadderEntry, 0)
	opts.Limit = maxLimit

	// Make one initial request to determine the size of the ladder.
	first, err := c.getLadderPage(opts)
	if err != nil {
		return Ladder{}, err
	}
	ladderSize := first.TotalEntries
	if ladderSize <= maxLimit {
		return first, nil
	}

	// If there are entries remaining, make further requests.
	var (
		wg    sync.WaitGroup
		lock  sync.RWMutex
		errCh = make(chan error, maxPages)
	)
	for i := maxLimit; i < ladderSize; i += maxLimit {
		go func(offset int) {
			wg.Add(1)
			subOpts := opts
			subOpts.Offset = offset
			page, err := c.getLadderPage(subOpts)
			if err != nil {
				errCh <- err
			}
			lock.Lock()
			entries = append(entries, page.Entries...)
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	select {
	case err, ok := <-errCh:
		if ok {
			return Ladder{}, err
		}
	default:
		// Continue.
	}

	// Copy first page to get top-level values.
	ladder := first
	ladder.Entries = entries
	return ladder, nil
}
