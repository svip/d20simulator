package main

import (
	"math/rand"
	"time"
	"fmt"
)

type simStats struct {
	runs int
	hits int
	rerolls int
}

func newStats() *simStats {
	return &simStats{}
}

type sim struct {
	regularStats *simStats
	newStats *simStats
	mod int
	dc int
}

func newSim(mod int, dc int) sim {
	return sim{
		newStats(),
		newStats(),
		mod,
		dc,
	}
}

func (s *sim) run(newRules bool, stats *simStats) {
	result := rand.Intn(20) + 1 // roll d20
	if newRules && result <= s.mod {
		stats.rerolls++
		result = rand.Intn(20) + 1 // re-roll d20
	}
	if result + s.mod >= s.dc {
		stats.hits++
	}
	stats.runs++
}

func (s *sim) RunOnce() {
	s.run(false, s.regularStats)
	s.run(true, s.newStats)
}

func (s *sim) PrintStats(title string) {
	fmt.Println(title)
	fmt.Printf("Ability mod: +%d, DC: %d\n", s.mod, s.dc)
	fmt.Printf("Standard rules: Runs: %d, Hits: %d, Percentage: %.2f %%\n", s.regularStats.runs, s.regularStats.hits,
			(float64(s.regularStats.hits)/float64(s.regularStats.runs))*100.0)
	fmt.Printf("Standard rules: Runs: %d, Hits: %d, Percentage: %.2f %% (re-rolls: %d)\n", s.newStats.runs, s.newStats.hits,
			(float64(s.newStats.hits)/float64(s.newStats.runs))*100.0, s.newStats.rerolls)
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	lowvhigh := newSim(2, 20)
	highvlow := newSim(5, 10)
	highvhigh := newSim(5, 20)
	lowvlow := newSim(2, 10)

	max := 10000

	for i := 0; i < max; i++ {
		lowvhigh.RunOnce()
		highvlow.RunOnce()
		highvhigh.RunOnce()
		lowvlow.RunOnce()
	}

	lowvhigh.PrintStats("Low vs. high")
	highvlow.PrintStats("High vs. low")
	highvhigh.PrintStats("High vs. high")
	lowvlow.PrintStats("Low vs. low")
}