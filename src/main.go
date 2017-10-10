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
	title string
	stats []*simStats
	mod int
	dc int
}

func newSim(title string, mod int, dc int) *sim {
	stats := []*simStats{newStats(), newStats()}
	return &sim{
		title,
		stats,
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
	s.run(false, s.stats[0])
	s.run(true, s.stats[1])
}

func (s *sim) PrintStats() {
	fmt.Println(s.title)
	fmt.Printf("Ability mod: +%d, DC: %d\n", s.mod, s.dc)
	fmt.Printf("Standard rules: Runs: %d, Hits: %d, Percentage: %.2f %%\n", s.stats[0].runs, s.stats[0].hits,
			(float64(s.stats[0].hits)/float64(s.stats[0].runs))*100.0)
	fmt.Printf("Standard rules: Runs: %d, Hits: %d, Percentage: %.2f %% (re-rolls: %d)\n", s.stats[1].runs, s.stats[1].hits,
			(float64(s.stats[1].hits)/float64(s.stats[1].runs))*100.0, s.stats[1].rerolls)
	fmt.Println()
}

type simContainer struct {
	sims []*sim
}

func newSimContainer() simContainer{
	return simContainer{}
}

func (c *simContainer) addSim(title string, mod int, dc int) {
	c.sims = append(c.sims, newSim(title, mod, dc))
}

func (c *simContainer) RunAllOnce() {
	for _, sim := range c.sims {
		sim.RunOnce()
	}
}

func (c *simContainer) PrintAllStats() {
	for _, sim := range c.sims {
		sim.PrintStats()
	}
}

type graphPoint struct {
	newRule bool
	x int
	y int
}

func (c *simContainer) DrawGraph() {
	/*width := 70
	height := 20
	for _, sim := range c.sims {

	}*/
}

func main() {
	rand.Seed(time.Now().UnixNano())

	container := newSimContainer()

	container.addSim("Low vs. high", 1, 20)
	container.addSim("High vs. low", 5, 10)
	container.addSim("High vs. high", 5, 20)
	container.addSim("Low vs. low", 1, 10)

	container.addSim("Long sword vs low", 6, 12)
	container.addSim("Long sword vs high", 6, 17)
	container.addSim("Javelin vs low", 5, 12)
	container.addSim("Javelin vs high", 5, 17)
	container.addSim("Quarterstaff vs low", 2, 12)
	container.addSim("Quarterstaff vs high", 2, 17)

	max := 10000

	for i := 0; i < max; i++ {
		container.RunAllOnce()
	}

	container.PrintAllStats()
	container.DrawGraph()
}