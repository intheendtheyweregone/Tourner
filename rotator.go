package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Randomly select one of the slices inside of our map and then return it's contents
func pick(m map[string][]string) (string, []string) {
	k := rand.Intn(len(m))
	for s, x := range m {
		if k == 0 {
			return s, x
		}
		k--
	}
	panic("unreachable")
}

func selectSite() (string, string) {
	sites := make(map[string][]string)

	sites["LoliSafe"] = []string{"https://l.trs.tn/api/upload",
		"https://take-me-to.space/api/upload",
		"https://safe.waifuhunter.club/api/upload"}

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	hostType, p := pick(sites)
	randomIndex := rand.Intn(len(p))
	randomSitePicked := p[randomIndex]
	fmt.Printf("\n [Tourner] randomly selected: %v for upload", randomSitePicked)
	return hostType, randomSitePicked

}
