package main

import (
	"math/rand"
	"time"
	"fmt"
)

type simStats struct {
	dc int
	runs int
	hits int
	rerolls int
}

func newStat(dc int) *simStats {
	return &simStats{dc, 0, 0, 0}
}

type sim struct {
	title string
	oldStats []*simStats
	newStats []*simStats
	mod int
}

func newSim(title string, mod int, basedc int, dctries int) *sim {
	var oldStats []*simStats
	var newStats []*simStats
	for i := 0; i < dctries; i++ {
		oldStats = append(oldStats, newStat(basedc+i))
		newStats = append(newStats, newStat(basedc+i))
	}
	return &sim{
		title,
		oldStats,
		newStats,
		mod,
	}
}

func (s *sim) run(newRules bool, stats *simStats) {
	result := rand.Intn(20) + 1 // roll d20
	if newRules && result <= s.mod {
		stats.rerolls++
		result = rand.Intn(20) + 1 // re-roll d20
	}
	if result + s.mod >= stats.dc {
		stats.hits++
	}
	stats.runs++
}

func (s *sim) RunOnce() {
	for _, simStat := range s.oldStats {
		s.run(false, simStat)
	}
	for _, simStat := range s.newStats {
		s.run(true, simStat)
	}
}

func (s *sim) PrintStats() {
	fmt.Println(s.title)
	fmt.Printf("Ability mod: +%d, DC: %d\n", s.mod, s.oldStats[0].dc)
	fmt.Printf("Standard rules: Runs: %d, Hits: %d, Percentage: %.2f %%\n", s.oldStats[0].runs, s.oldStats[0].hits,
			(float64(s.oldStats[0].hits)/float64(s.oldStats[0].runs))*100.0)
	fmt.Printf("New rules: Runs: %d, Hits: %d, Percentage: %.2f %% (re-rolls: %d)\n", s.newStats[1].runs, s.newStats[1].hits,
			(float64(s.newStats[1].hits)/float64(s.newStats[1].runs))*100.0, s.newStats[1].rerolls)
	fmt.Println()
}

type simContainer struct {
	sims []*sim
}

func newSimContainer() simContainer{
	return simContainer{}
}

func (c *simContainer) addSim(title string, mod int, basedc int, dctries int) {
	c.sims = append(c.sims, newSim(title, mod, basedc, dctries))
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

type point struct {
	x int
	y int
}

func (c *simContainer) DrawGraph() {
	width := 70
	pointWidth := 4
	height := 17
	for _, sim := range c.sims {
		fmt.Printf("%s: +%d    * = standard, # = new\n", sim.title, sim.mod)
		var oldPoints []point
		var newPoints []point
		for _, stat := range sim.oldStats {
			x := (stat.dc-5) * pointWidth
			y := height - int((float64(stat.hits)/float64(stat.runs))*float64(height))
			oldPoints = append(oldPoints, point{x,y})
		}
		for _, stat := range sim.newStats {
			x := (stat.dc-5) * pointWidth
			y := height - int((float64(stat.hits)/float64(stat.runs))*float64(height))
			newPoints = append(newPoints, point{x,y})
		}
		for h := 0; h < height; h++ {
			for w := 0; w < width; w+=pointWidth {
				foundPoint := false
				for _, p := range oldPoints {
					if p.x == w && p.y == h {
						fmt.Printf(" *")
						foundPoint = true
					}
				}
				if !foundPoint {
					for _, p := range newPoints {
						if p.x == w && p.y == h {
							fmt.Printf(" #")
							foundPoint = true
						}
					}
				}
				if !foundPoint {
					fmt.Printf("  ")
				}
				for i := 2; i < pointWidth; i++ {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		for _, stat := range sim.oldStats {
			dc := stat.dc
			if dc >= 10 {
				fmt.Printf("%d  ", dc)
			} else {
				fmt.Printf(" %d  ", dc)
			}
		}
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	container := newSimContainer()

	container.addSim("+0", 0, 5, 20)
	container.addSim("+1", 1, 5, 20)
	container.addSim("+2", 2, 5, 20)
	container.addSim("+3", 3, 5, 20)
	container.addSim("+4", 4, 5, 20)
	container.addSim("+5", 5, 5, 20)
	container.addSim("+6", 6, 5, 20)

	max := 10000

	for i := 0; i < max; i++ {
		container.RunAllOnce()
	}

	container.PrintAllStats()
	container.DrawGraph()
}