package poeapi

import "time"

var (
	validHosts = map[string]struct{}{
		"api.pathofexile.com": struct{}{},
		"localhost:8000":      struct{}{},
	}

	validRealms = map[string]struct{}{
		"pc":   struct{}{},
		"xbox": struct{}{},
		"sony": struct{}{},
	}

	validLeagueTypes = map[string]struct{}{
		"main":   struct{}{},
		"event":  struct{}{},
		"season": struct{}{},
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

// PVPMatch represents a PVP event type.
type PVPMatch struct {
	ID            string    `json:"id"`
	Realm         string    `json:"realm"`
	StartTime     time.Time `json:"startAt"`
	EndTime       time.Time `json:"endAt"`
	LadderURL     string    `json:"url"`
	Description   string    `json:"description"`
	GlickoRatings bool      `json:"glickoRatings"`
	PVP           bool      `json:"pvp"`
	Style         string    `json:"style"`
	RegisterTime  time.Time `json:"registerAt"`
}

// StashResponse represents a response from the stash tab API which contains
// multiples stashes.
type StashResponse struct {
	NextChangeID string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
}

// Stash represents a single public stash tab.
type Stash struct {
	ID                string `json:"id"`
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	Index             string `json:"stash"`
	Type              string `json:"stashType"`
	Items             []Item `json:"items"`
	Public            bool   `json:"public"`
}

// Item represents an item listing in the stash tab API.
type Item struct {
	Corrupted         bool          `json:"corrupted"`
	ExplicitMods      []string      `json:"explicitMods"`
	FlavorText        []string      `json:"flavourText"`
	FrameType         int64         `json:"frameType"`
	Height            int64         `json:"h"`
	ID                string        `json:"id"`
	Icon              string        `json:"icon"`
	Identified        bool          `json:"identified"`
	ImplicitMods      []string      `json:"implicitMods"`
	InventoryID       string        `json:"inventoryId"`
	ItemLevel         int64         `json:"ilvl"`
	League            string        `json:"league"`
	LockedToCharacter bool          `json:"lockedToCharacter"`
	Name              string        `json:"name"`
	Note              string        `json:"note"`
	Properties        []Property    `json:"properties"`
	Requirements      []Requirement `json:"requirements"`
	SocketedItems     interface{}   `json:"socketedItems"`
	Sockets           []Socket      `json:"sockets"`
	Support           bool          `json:"support"`
	TalismanTier      int64         `json:"talismanTier"`
	TypeLine          string        `json:"typeLine"`
	Verified          bool          `json:"verified"`
	Width             int64         `json:"w"`
	XPosition         int64         `json:"x"`
	YPosition         int64         `json:"y"`
}

// Property represents a property of an Item.
type Property struct {
	DisplayMode int64           `json:"displayMode"`
	Name        string          `json:"name"`
	Values      [][]interface{} `json:"values"`
}

// Requirement represents an attribute requirement on an Item, such as strength.
type Requirement struct {
	DisplayMode int64           `json:"displayMode"`
	Name        string          `json:"name"`
	Values      [][]interface{} `json:"values"`
}

// Socket represents a socket on an item, including its color and links.
type Socket struct {
	Attribute string `json:"attr"`
	Group     int64  `json:"group"`
}
