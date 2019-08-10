// ladderstats prints the distribution of character classes in a ladder.
package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/willroberts/poeapi"
)

func main() {
	client, err := poeapi.NewAPIClient(poeapi.DefaultClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	opts := poeapi.GetLadderOptions{
		ID:        "SSF Hardcore",
		Realm:     "pc",
		UniqueIDs: false,
	}

	l, err := client.GetLadder(opts)
	if err != nil {
		log.Fatal(err)
	}

	dist := getDistributions(l.TotalEntries, countClasses(l))
	sorted := sortDistributions(dist)
	fmt.Println("Distribution of characters by class in ladder:")
	for _, entry := range sorted {
		fmt.Printf("    %s: %.2f%%\n", entry.key, entry.val*100)
	}
}

func countClasses(l poeapi.Ladder) map[string]int {
	count := make(map[string]int, 0)
	for _, e := range l.Entries {
		class := e.Character.Class
		val, ok := count[class]
		if !ok {
			count[class] = 1
			continue
		}
		count[class] = val + 1
	}
	return count
}

func getDistributions(total int, entries map[string]int) map[string]float64 {
	dist := make(map[string]float64, 0)
	for k, v := range entries {
		pct := float64(v) / float64(total)
		dist[k] = pct
	}
	return dist
}

func sortDistributions(d map[string]float64) classSet {
	set := make(classSet, len(d))
	i := 0
	for k, v := range d {
		set[i] = classDist{k, v}
		i++
	}
	sort.Sort(sort.Reverse(set))
	return set
}

// classSet is a custom type which allows us to sort a map by its values.
type classSet []classDist

func (cs classSet) Len() int           { return len(cs) }
func (cs classSet) Less(a, b int) bool { return cs[a].val < cs[b].val }
func (cs classSet) Swap(a, b int)      { cs[a], cs[b] = cs[b], cs[a] }

type classDist struct {
	key string
	val float64
}
