package poeapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

const (
	repo              = "github.com/willroberts/poeapi"
	rateLimitEndpoint = "/rate-limit-me"
	failureEndpoint   = "/fail-me"
)

var (
	testTimeout = 200 * time.Millisecond
	testClient  = &http.Client{Timeout: testTimeout}
)

func init() {
	if err := startStubServer(); err != nil {
		log.Fatalf("failed to start test server: %v", err)
	}
}

func loadFixture(filename string) (string, error) {
	// GitHub Actions Workaround: load fixtures without GOPATH.
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		return string(b), nil
	}

	// Allow loading of fixtures on Windows systems.
	path := fmt.Sprintf("%s/src/%s/%s", os.Getenv("GOPATH"), repo, filename)
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "/", "\\")
	}

	b2, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b2), nil
}

type testHandler struct {
	ladderFixture       string
	leagueRuleFixture   string
	leagueRulesFixture  string
	leagueFixture       string
	leaguesFixture      string
	pvpMatchesFixture   string
	stashFixture        string
	latestChangeFixture string
}

func newTestHandler() (testHandler, error) {
	h := testHandler{}
	f, err := loadFixture("fixtures/ladder.json")
	if err != nil {
		return testHandler{}, err
	}
	h.ladderFixture = f
	f, err = loadFixture("fixtures/league-rule.json")
	if err != nil {
		return testHandler{}, err
	}
	h.leagueRuleFixture = f
	f, err = loadFixture("fixtures/league-rules.json")
	if err != nil {
		return testHandler{}, err
	}
	h.leagueRulesFixture = f
	f, err = loadFixture("fixtures/league.json")
	if err != nil {
		return testHandler{}, err
	}
	h.leagueFixture = f
	f, err = loadFixture("fixtures/leagues.json")
	if err != nil {
		return testHandler{}, err
	}
	h.leaguesFixture = f
	f, err = loadFixture("fixtures/pvp-matches.json")
	if err != nil {
		return testHandler{}, err
	}
	h.pvpMatchesFixture = f
	f, err = loadFixture("fixtures/stash.json")
	if err != nil {
		return testHandler{}, err
	}
	h.stashFixture = f
	f, err = loadFixture("fixtures/latest-change.json")
	if err != nil {
		return testHandler{}, err
	}
	h.latestChangeFixture = f
	return h, nil
}

func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ladders/Standard":
		w.Write([]byte(h.ladderFixture))
	case "/league-rules/TurboMonsters":
		w.Write([]byte(h.leagueRuleFixture))
	case "/league-rules":
		w.Write([]byte(h.leagueRulesFixture))
	case "/leagues/Standard":
		w.Write([]byte(h.leagueFixture))
	case "/leagues":
		w.Write([]byte(h.leaguesFixture))
	case "/pvp-matches":
		w.Write([]byte(h.pvpMatchesFixture))
	case "/public-stash-tabs":
		w.Write([]byte(h.stashFixture))
	case "/api/Data/GetStats":
		w.Write([]byte(h.latestChangeFixture))
	case rateLimitEndpoint:
		w.WriteHeader(http.StatusTooManyRequests)
	case failureEndpoint:
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func startStubServer() error {
	h, err := newTestHandler()
	if err != nil {
		return err
	}
	s := &http.Server{
		Addr:         testHost,
		Handler:      h,
		ReadTimeout:  testTimeout,
		WriteTimeout: testTimeout,
	}
	go func() {
		log.Println("starting local http server")
		if err := s.ListenAndServe(); err != nil {
			log.Println("http test server error:", err)
		}
	}()
	return nil
}

func TestStubServer(t *testing.T) {
	_, err := http.Get(fmt.Sprintf("http://%s/test", testHost))
	if err != nil {
		t.Fatalf("failed stub server test request: %v", err)
	}
}
