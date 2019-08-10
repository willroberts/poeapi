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

	earliestLabyrinthTime = 1456790400 // March 1, 2016: One day before 3.2.0.
)

var (
	validRealms = map[string]struct{}{
		"pc":   struct{}{},
		"xbox": struct{}{},
		"sony": struct{}{},
	}

	validLadderTypes = map[string]struct{}{
		"league":    struct{}{},
		"labyrinth": struct{}{},
		"pvp":       struct{}{},
	}

	validLabyrinthDifficulties = map[string]struct{}{
		"Normal":    struct{}{},
		"Cruel":     struct{}{},
		"Merciless": struct{}{},
		"Eternal":   struct{}{},
	}
)

// GetLadderOptions contains the request parameters for the ladder endpoint.
// All parameters are optional with the exception of ID.
type GetLadderOptions struct {
	// The name of the league whose ladder you want to retrieve.
	ID string

	// The realm of the ladder.
	// Valid options: 'pc', 'xbox', or 'sony'.
	Realm string

	// Number of results to retrieve. Default: 20. Maximum: 200.
	Limit int

	// Position at which to start retrieving results. Default: 0.
	Offset int

	// The type of league whose ladder you want to retrieve.
	// Valid options: 'league', 'pvp', or 'labyrinth'.
	Type string

	// Associate UUIDs with each character returned for tracking purposes.
	UniqueIDs bool

	// Only include the given account in results.
	AccountName string

	// Difficulty of the Labyrinth ladder to retrieve.
	// Valid options: 'Normal', 'Cruel', 'Merciless', or 'Eternal'.
	LabyrinthDifficulty string

	// Start time of the Labyrinth ladder to retrieve. This is a Unix timestamp.
	LabyrinthStartTime int
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

func validateGetLadderOptions(opts GetLadderOptions) error {
	if opts.ID == "" {
		return ErrMissingID
	}
	if _, ok := validRealms[opts.Realm]; opts.Realm != "" && !ok {
		return ErrInvalidRealm
	}
	if opts.Limit < 1 || opts.Limit > maxLimit {
		return ErrInvalidLimit
	}
	if opts.Offset < 0 || opts.Offset > maxLimit*maxPages {
		return ErrInvalidOffset
	}
	if _, ok := validLadderTypes[opts.Type]; opts.Type != "" && !ok {
		return ErrInvalidLadderType
	}
	if opts.Type == "labyrinth" {
		if opts.LabyrinthDifficulty != "" {
			if _, ok := validLabyrinthDifficulties[opts.LabyrinthDifficulty]; !ok {
				return ErrInvalidDifficulty
			}
		}
		if opts.LabyrinthStartTime < 0 {
			return ErrInvalidLabyrinthStartTime
		}
		if opts.LabyrinthStartTime > 0 && opts.LabyrinthStartTime < earliestLabyrinthTime {
			return ErrInvalidLabyrinthStartTime
		}
	}

	return nil
}

func (c *client) getLadderPage(opts GetLadderOptions) (Ladder, error) {
	if err := validateGetLadderOptions(opts); err != nil {
		return Ladder{}, err
	}
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
