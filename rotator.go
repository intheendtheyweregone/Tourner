package main

import (
	"math/rand"
	"time"
)

// function yanked from stackoverflow, in the futue will make it fetch something from github
// with an updated list of my keeping.
// or custom lists via txt so that users can decide which instances they want to rotate between

func selectSite() string {
	sites := make([]string, 0)
	sites = append(sites,
		"https://l.trs.tn/api/upload",
		"https://take-me-to.space/api/upload",
		"https://safe.waifuhunter.club/api/upload",
	)

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	return sites[rand.Intn(len(sites))]

}
