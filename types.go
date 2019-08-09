package poeapi

import "time"

// League represents a permanent or challenge league.
type League struct {
	Name         string       `json:"id" validate:"nonzero"`
	Realm        string       `json:"realm"`
	Description  string       `json:"description"`
	LadderURL    string       `json:"url"`
	StartTime    time.Time    `json:"startAt"`
	EndTime      time.Time    `json:"endAt"`
	DelveEnabled bool         `json:"delveEvent"`
	Rules        []LeagueRule `json:"rules"`
}

// LeagueRule represents a modifier placed on a league.
type LeagueRule struct {
	ID          string `json:"id" validate:"nonzero"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Ladder represents the leaderboard for a specific league.
type Ladder struct {
	TotalEntries int           `json:"total" validate:"nonzero"`
	Title        string        `json:"title"`
	StartTime    int           `json:"startTime"`
	Entries      []LadderEntry `json:"entries"`
}

// LadderEntry represents an entry on the ladder.
type LadderEntry struct {
	Online        bool      `json:"bool"`
	Rank          int       `json:"rank"`
	LabyrinthTime int       `json:"time"`
	Character     Character `json:"character"`
	Account       Account   `json:"account"`
}

// Character represents a player in a ladder entry.
type Character struct {
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Class      string `json:"class"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
}

// Account represents an account for a ladder entry.
type Account struct {
	Name       string     `json:"name"`
	Realm      string     `json:"realm"`
	Challenges Challenges `json:"challenges"`
}

// Challenges represents an account's completed challenges.
type Challenges struct {
	Total int `json:"total"`
}
